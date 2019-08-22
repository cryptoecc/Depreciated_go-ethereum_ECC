package eccpow

import (
	"crypto/sha256"
	"github.com/Onther-Tech/go-ethereum/crypto"
	"math"
	"math/big"
	"math/rand"
	"sync"
	"time"

	"github.com/Onther-Tech/go-ethereum/common"
	"github.com/Onther-Tech/go-ethereum/consensus"
	"github.com/Onther-Tech/go-ethereum/core/types"
	"github.com/Onther-Tech/go-ethereum/metrics"
	"github.com/Onther-Tech/go-ethereum/rpc"
)

// sealTask wraps a seal block with relative result channel for remote sealer thread.
type sealTask struct {
	block   *types.Block
	results chan<- *types.Block
}

// mineResult wraps the pow solution parameters for the specified block.
type mineResult struct {
	nonce     types.BlockNonce
	mixDigest common.Hash
	hash      common.Hash

	errc chan error
}

// hashrate wraps the hash rate submitted by the remote sealer.
type hashrate struct {
	id   common.Hash
	ping time.Time
	rate uint64

	done chan struct{}
}

// sealWork wraps a seal work package for remote sealer.
type sealWork struct {
	errc chan error
	res  chan [4]string
}

// hasher is a repetitive hasher allowing the same hash data structures to be
// reused between hash runs instead of requiring new ones to be created.
//var hasher func(dest []byte, data []byte)

type ECC struct {
	H             [][]int
	col_in_row    [][]int
	row_in_col    [][]int
	hashVector    []int
	tmpHashVector []byte
	outputWord    []int
	LRqtl         [][]float64
	LRrtl         [][]float64
	LRpt          []float64
	LRft          []float64

	shared *ECC

	// Mining related fields
	rand     *rand.Rand    // Properly seeded random source for nonces
	threads  int           // Number of threads to mine on if mining
	update   chan struct{} // Notification channel to update mining parameters
	hashrate metrics.Meter // Meter tracking the average hashrate

	// Remote sealer related fields
	workCh       chan *sealTask   // Notification channel to push new work and relative result channel to remote sealer
	fetchWorkCh  chan *sealWork   // Channel used for remote sealer to fetch mining work
	submitWorkCh chan *mineResult // Channel used for remote sealer to submit their mining result
	fetchRateCh  chan chan uint64 // Channel used to gather submitted hash rate for local or remote sealer.
	submitRateCh chan *hashrate   // Channel used for remote sealer to submit their mining hashrate

	lock      sync.Mutex      // Ensures thread safety for the in-memory caches and mining fields
	closeOnce sync.Once       // Ensures exit channel will not be closed twice.
	exitCh    chan chan error // Notification channel to exiting backend threads

}

var (
	n    int
	wc   int
	wr   int
	seed int
	m    int

	// two256 is a big integer representing 2^256
	two256 = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), big.NewInt(0))
)

//const cross_err = 0.01

type (
	intMatrix   [][]int
	floatMatrix [][]float64
)

func (ecc *ECC) decoding() {
	maxIter := 20
	for i := 0; i < len(ecc.outputWord); i++ {
		ecc.outputWord[i] = 0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ecc.LRqtl[i][j] = 0
			ecc.LRrtl[i][j] = 0
		}
		//ecc.LRft[i] = math.Log((1 - cross_err) / cross_err * float64(ecc.hashVector[i]*2-1))
		if ecc.hashVector[i] == 0 {
			ecc.LRft[i] = 1.005033585350145
		} else {
			ecc.LRft[i] = -1.005033585350145
		}
	}

	for ind := 1; ind <= maxIter; ind++ {
		for t := 0; t < n; t++ {
			for m := 0; m < wc; m++ {
				temp3 := 0.0
				for mp := 0; mp < wc; mp++ {
					if mp != m {
						temp3 = infinityTest(temp3 + float64(ecc.LRrtl[t][ecc.row_in_col[mp][t]]))
					}
				}
				ecc.LRqtl[t][ecc.row_in_col[m][t]] = infinityTest(ecc.LRft[t] + float64(temp3))
			}
		}

		for k := 0; k < m; k++ {
			for l := 0; l < wr; l++ {
				temp3 := 0.0
				sign := 1
				for m := 0; m < wr; m++ {
					if m != l {
						temp3 = temp3 + funcF(math.Abs(ecc.LRqtl[ecc.col_in_row[m][k]][k]))
						temp_sign := 0
						if ecc.LRqtl[ecc.col_in_row[m][k]][k] > 0.0 {
							temp_sign = 1.0
						} else {
							temp_sign = -1.0
						}
						sign = sign * temp_sign
					}
				}
				magnitude := funcF(temp3)
				ecc.LRrtl[ecc.col_in_row[l][k]][k] = infinityTest(float64(sign) * magnitude)
			}
		}

		for m := 0; m < n; m++ {
			ecc.LRpt[m] = infinityTest(ecc.LRft[m])
			for k := 0; k < wc; k++ {
				ecc.LRpt[m] += ecc.LRrtl[m][ecc.row_in_col[k][m]]
				ecc.LRpt[m] = infinityTest(ecc.LRpt[m])
			}
		}
	}

	for i := 0; i < n; i++ {
		if ecc.LRpt[i] >= 0 {
			ecc.outputWord[i] = 1
		} else {
			ecc.outputWord[i] = 0
		}
	}
}

func generateSeed(prev_hash []byte) int {
	sum := 0
	i := 1
	for i < 31 {
		sum = sum + int(prev_hash[i])
		i++
	}
	seed = sum
	return sum
}

func (ecc *ECC) generateHashVector(headerWithNonce []byte) {
	if n <= 256 {
		hash := sha256.New()
		hash.Write(headerWithNonce)
		md := hash.Sum(nil)
		copy(ecc.tmpHashVector, md)
		//tmp := crypto.Keccak256(headerWithNonce)

	}

	for i := 0; i < n/8; i++ {
		decimal := int(ecc.tmpHashVector[i])
		for j := 7; j >= 0; j-- {
			ecc.hashVector[j+8*(i)] = decimal % 2
			ecc.outputWord[j+8*(i)] = decimal % 2
			decimal = decimal / 2
		}
	}

}

func (ecc *ECC) generateH() [][]int {
	//if ecc.H == nil{
	//	return false
	//}
	k := m / wc

	for i := 0; i < k; i++ {
		for j := i * wr; j < (i+1)*wr; j++ {
			ecc.H[i][j] = 1
		}
	}

	for i := 1; i < wc; i++ {
		colOrder := make([]int, n)
		for j := 0; j < n; j++ {
			colOrder[j] = j
		}

		rand.Seed(int64(seed))
		rand.Shuffle(len(colOrder), func(i, j int) { colOrder[i], colOrder[j] = colOrder[j], colOrder[i] })
		seed--

		for j := 0; j < n; j++ {
			index := colOrder[j]/wr + k*i
			ecc.H[index][j] = 1
		}
	}
	return ecc.H
}

func (ecc *ECC) generateQ() bool {
	row_index := 0
	col_index := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ecc.H[i][j] != 0 {
				ecc.col_in_row[col_index%wr][i] = j
				ecc.row_in_col[row_index/n][j] = i

				col_index++
				row_index++
			}
		}
	}
	return true
}

func (ecc *ECC) decision() bool {
	for i := 0; i < m; i++ {
		sum := 0
		for j := 0; j < wr; j++ {
			sum = sum + ecc.outputWord[ecc.col_in_row[j][i]]
		}
		if sum%2 != 0 {
			return false
		}
	}
	return true
}

func runLDPC(prev_hash []byte, cur_hash []byte, n int, wc int, wr int) (int, []byte) {
	//n := 24
	//wc := 3
	//wr := 6
	m := set_difficulty(n, wc, wr)

	ecc := ECC{
		H:             newIntMatrix(m, n),
		col_in_row:    newIntMatrix(wr, m),
		row_in_col:    newIntMatrix(wc, n),
		hashVector:    make([]int, n),
		tmpHashVector: make([]byte, n),
		outputWord:    make([]int, n),
		LRqtl:         newFloatMatrix(n, m),
		LRrtl:         newFloatMatrix(n, m),
		LRpt:          make([]float64, n),
		LRft:          make([]float64, n),
	}

	generateSeed(prev_hash)
	ecc.generateH()
	ecc.generateQ()

	nonce := 0
	hashWithNonce := make([]byte, len(cur_hash))
	for {
		cur_hash_ := string(cur_hash)
		nonce_ := string(nonce)

		hashAndNonce := cur_hash_ + nonce_
		hashWithNonce = []byte(hashAndNonce)
		ecc.generateHashVector(hashWithNonce)
		flag := ecc.decision()
		if flag == false {
			ecc.decoding()
			flag = ecc.decision()
		}
		if flag == true {
			break
		}
		nonce++
	}
	return nonce, crypto.Keccak256(hashWithNonce)
}

func set_difficulty(_n int, _wc int, _wr int) int {
	n = _n
	wc = _wc
	wr = _wr
	m = (int)(n * wc / wr)

	//New(n, wc, wr, m)
	return m
}

func infinityTest(x float64) float64 {
	if x >= 64.0 {
		return float64(64.0)
	} else if x <= -64.0 {
		return float64(-64.0)
	} else {
		return float64(x)
	}
}

//bigInfinity = 1000000
func funcF(x float64) float64 {
	if x >= 1000000 {
		return float64(1.0 / 1000000)
	} else if x <= (1.0 / 1000000) {
		return float64(1000000)
	} else {
		return float64(math.Log(math.Exp(x)+1)/math.Exp(x) - 1)
	}
}

func New(n int, wr int, wc int) *ECC {
	ecc := &ECC{
		H:             newIntMatrix(m, n),
		col_in_row:    newIntMatrix(wr, m),
		row_in_col:    newIntMatrix(wc, n),
		hashVector:    make([]int, n),
		tmpHashVector: make([]byte, n),
		outputWord:    make([]int, n),
		LRqtl:         newFloatMatrix(n, m),
		LRrtl:         newFloatMatrix(n, m),
		LRpt:          make([]float64, n),
		LRft:          make([]float64, n),
	}

	return ecc
}

func newIntMatrix(rows, cols int) intMatrix {
	m := intMatrix(make([][]int, rows))
	for i := range m {
		m[i] = make([]int, cols)
	}
	return m
}

func newFloatMatrix(rows, cols int) floatMatrix {
	m := floatMatrix(make([][]float64, rows))
	for i := range m {
		m[i] = make([]float64, cols)
	}
	return m
}

// NewShared creates a full sized ecc PoW shared between all requesters running
// in the same process.
//func NewShared() *ECC {
//	return &ecc{shared: sharedecc}
//}

func NewTester(notify []string, noverify bool) *ECC {
	ecc := &ECC{
		update:       make(chan struct{}),
		hashrate:     metrics.NewMeterForced(),
		workCh:       make(chan *sealTask),
		fetchWorkCh:  make(chan *sealWork),
		submitWorkCh: make(chan *mineResult),
		fetchRateCh:  make(chan chan uint64),
		submitRateCh: make(chan *hashrate),
		exitCh:       make(chan chan error),
	}
	go ecc.remote(notify, noverify)
	return ecc
}

// Close closes the exit channel to notify all backend threads exiting.
func (ecc *ECC) Close() error {
	var err error
	ecc.closeOnce.Do(func() {
		// Short circuit if the exit channel is not allocated.
		if ecc.exitCh == nil {
			return
		}
		errc := make(chan error)
		ecc.exitCh <- errc
		err = <-errc
		close(ecc.exitCh)
	})
	return err
}

// Threads returns the number of mining threads currently enabled. This doesn't
// necessarily mean that mining is running!
func (ecc *ECC) Threads() int {
	ecc.lock.Lock()
	defer ecc.lock.Unlock()

	return ecc.threads
}

// SetThreads updates the number of mining threads currently enabled. Calling
// this method does not start mining, only sets the thread count. If zero is
// specified, the miner will use all cores of the machine. Setting a thread
// count below zero is allowed and will cause the miner to idle, without any
// work being done.
func (ecc *ECC) SetThreads(threads int) {
	ecc.lock.Lock()
	defer ecc.lock.Unlock()

	// If we're running a shared PoW, set the thread count on that instead
	if ecc.shared != nil {
		ecc.shared.SetThreads(threads)
		return
	}
	// Update the threads and ping any running seal to pull in any changes
	ecc.threads = threads
	select {
	case ecc.update <- struct{}{}:
	default:
	}
}

// Hashrate implements PoW, returning the measured rate of the search invocations
// per second over the last minute.
// Note the returned hashrate includes local hashrate, but also includes the total
// hashrate of all remote miner.
func (ecc *ECC) Hashrate() float64 {
	// Short circuit if we are run the ecc in normal/test mode.

	var res = make(chan uint64, 1)

	select {
	case ecc.fetchRateCh <- res:
	case <-ecc.exitCh:
		// Return local hashrate only if ecc is stopped.
		return ecc.hashrate.Rate1()
	}

	// Gather total submitted hash rate of remote sealers.
	return ecc.hashrate.Rate1() + float64(<-res)
}

// APIs implements consensus.Engine, returning the user facing RPC APIs.
func (ecc *ECC) APIs(chain consensus.ChainReader) []rpc.API {
	// In order to ensure backward compatibility, we exposes ecc RPC APIs
	// to both eth and ecc namespaces.
	return []rpc.API{
		{
			Namespace: "eth",
			Version:   "1.0",
			Service:   &API{ecc},
			Public:    true,
		},
		{
			Namespace: "ecc",
			Version:   "1.0",
			Service:   &API{ecc},
			Public:    true,
		},
	}
}

//// SeedHash is the seed to use for generating a verification cache and the mining
//// dataset.
func SeedHash(block *types.Block) []byte {
	return block.ParentHash().Bytes()
}

package eccpow

import (
	"github.com/Onther-Tech/go-ethereum/crypto"
	"math"
	"math/rand"
	"unsafe"
)

// hasher is a repetitive hasher allowing the same hash data structures to be
// reused between hash runs instead of requiring new ones to be created.
type hasher func(dest []byte, data []byte)

var n int
var wc int
var wr int
var seed int
var H [][]int
var m int
var col_in_row [][]int
var row_in_col [][]int
var hashVector []int
var tmpHashVector []byte
var outputWord []int
var LRqtl [][]float64
var LRrtl [][]float64
var LRpt []float64
var LRft []float64
var cross_err = 0.01

func decoding() {
	maxIter := 20
	for i := 0; i < len(outputWord); i++ {
		outputWord[i] = 0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			LRqtl[i][j] = 0
			LRrtl[i][j] = 0
		}
		LRft[i] = math.Log((1 - cross_err) / cross_err * float64(hashVector[i]*2-1))
	}

	for ind := 1; ind <= maxIter; ind++ {
		for t := 0; t < n; t++ {
			for m := 0; m < wc; m++ {
				temp3 := 0
				for mp := 0; mp < wc; mp++ {
					if mp != m {
						a := LRrtl[t][row_in_col[mp][t]]
						b := float64(temp3) + a
						temp3 = infinityTest(float64(b))
					}
				}
				LRqtl[t][row_in_col[m][t]] = infinityTest(LRft[t] + float64(temp3))
			}
		}

		for k := 0; k < m; k++ {
			for l := 0; l < wr; l++ {
				temp3 := 0.0
				sign := 1
				for m := 0; m < wr; m++ {
					if m != l {
						temp3 = temp3 + funcF(math.Abs(LRqtl[col_in_row[m][k]][k]))
						temp_sign := 0
						if LRqtl[col_in_row[m][k]][k] > 0.0 {
							temp_sign = 1.0
						} else {
							temp_sign = -1.0
						}
						sign = sign * temp_sign
					}
				}
				magnitude := funcF(temp3)
				LRrtl[col_in_row[l][k]][k] = infinityTest(float64(sign) * magnitude)
			}
		}

		for m := 0; m < n; m++ {
			LRpt[m] = infinityTest(LRft[m])
			for k := 0; k < wc; k++ {
				LRpt[m] += LRrtl[m][row_in_col[k][m]]
				LRpt[m] = infinityTest(LRpt[m])
			}
		}
	}

	for i := 0; i < n; i++ {
		if LRpt[i] >= 0 {
			outputWord[i] = 1
		} else {
			outputWord[i] = 0
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

func generateHashVector(headerWithNonce []byte) {
	//inputSize := len(headerWithNonce)
	//hashVector := make([]byte, n)
	//tmpHashVector := make([]byte, 32)

	if n <= 256 {
		tmp := crypto.Keccak256(headerWithNonce)
		copy(tmpHashVector, tmp)
	}

	for i := 0; i < n/8; i++ {
		decimal := int(tmpHashVector[i])
		for j := 7; j >= 0; j-- {
			hashVector[j+8*(i)] = decimal % 2
		}
	}
	outputWord = hashVector
}

func generateH() bool {
	//if H == null{
	//	retrun false
	//}

	var k = m / wc

	for i := 0; i < k; 1++ {
		for j := i * wr; j < (i+1)*wr; j++ {
			H[i][j] = 1
		}
	}

	for i := 1; i < wc; i++ {
		colOrder := make([]int, n)
		for j := 0; j < n; j++ {
			colOrder[j] = j
		}
		seed--
		rand.Seed(int64(seed))

		val := make([]int, len(colOrder))
		for _, i := range rand.Perm(len(colOrder)) {
			val[i] = colOrder[i]
		}

		for j := 0; j < n; j++ {
			index := val[j]/wr + k*1
			H[index][j] = 1
		}
	}
	return true
}

func generateQ() bool {
	row_index := 0
	col_index := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			col_index++
			a := col_index % wr
			row_index++
			col_in_row[a][i] = j
			row_in_col[row_index/n][j] = i
		}
	}
	return true
}

func decision() bool {
	for i := 0; i < m; i++ {
		sum := 0
		for j := 0; j < wr; j++ {
			sum = sum + outputWord[col_in_row[j][i]]
		}
		if sum%2 != 0 {
			return false
		}
	}
	return true
}

func runLDPC(prev_hash []byte, cur_hash []byte) int {
	set_difficulty(24, 3, 6)
	generateSeed(prev_hash)
	generateH()
	generateQ()

	nonce := 0
	for {
		nonce_ := uint32(nonce)
		a := make([]byte, unsafe.Sizeof(nonce_))
		copy(a, *(*[]byte)(unsafe.Pointer(&nonce_)))

		generateHashVector()
		flag := decision()
		if flag == false {
			decoding()
			flag = decision()
		}
		if flag == true {
			break
		}
		nonce++
	}
	return nonce
}

func set_difficulty(_n int, _wc int, _wr int) bool {
	n = _n
	wc = _wc
	wr = _wr
	m = (int)(n * wc / wr)
	return true
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

//var bigInfinity = 1000000
func funcF(x float64) float64 {
	if x >= 1000000 {
		return float64(1.0 / 1000000)

	} else if x <= (1.0 / 1000000) {
		return float64(1000000)
	} else {
		return float64(math.Log(math.Exp(x)+1)/math.Exp(x) - 1)
	}
}

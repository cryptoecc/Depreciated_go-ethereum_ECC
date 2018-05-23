package downloader

import (
	"errors"
	"math"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/plasma/types"
)

const (
	measurementImpact = 0.1 // The impact a single measurement has on a peer's final throughput value.
)

var (
	errClosed            = errors.New("peer set is closed")
	errAlreadyRegistered = errors.New("peer is already registered")
	errNotRegistered     = errors.New("peer is not registered")
)

type peer interface {
	ID() []byte
	SendNewBlock(block *types.Block) error
	SendNewBlocks(blocks types.Blocks) error
	SendNewTransactions(txs []*types.Transaction) error
	RequestBlock(blkNum uint64) error
	RequestBlocks(blkNums []uint64) error
	RequestTransactions(hashes []common.Hash) error
	IsOperator() bool
}

type peerConnection struct {
	id string // Unique identifier of the peer

	blockIdle       int32   // Current block activity state of the peer (idle = 0, active = 1)
	blockThroughput float64 // Number of blocks (bodies) measured to be retrievable per second

	rtt time.Duration // Request round trip time to track responsiveness (QoS)

	blockStarted time.Time // Time instance when the last block (body) fetch was started
	peer         peer

	lacking map[common.Hash]struct{} // Set of hashes not to request (didn't have previously)

	log  log.Logger // Contextual logger to add extra infos to peer logs
	lock sync.RWMutex
}

func newPeerConnection(id string, peer peer, logger log.Logger) *peerConnection {
	return &peerConnection{
		id:   id,
		peer: peer,
		log:  logger,
	}
}

func (p *peerConnection) Reset() {
	p.lock.Lock()
	defer p.lock.Unlock()
	atomic.StoreInt32(&p.blockIdle, 0)

	p.blockThroughput = 0

	// p.lacking = make(map[common.Hash]struct{})
}

// BodyIdlePeers retrieves a flat list of all the currently body-idle peers within
// the active peer set, ordered by their reputation.
func (ps *peerSet) BodyIdlePeers() ([]*peerConnection, int) {
	idle := func(p *peerConnection) bool {
		return atomic.LoadInt32(&p.blockIdle) == 0
	}
	throughput := func(p *peerConnection) float64 {
		p.lock.RLock()
		defer p.lock.RUnlock()
		return p.blockThroughput
	}
	return ps.idlePeers(62, 64, idle, throughput)
}

// idlePeers retrieves a flat list of all currently idle peers satisfying the
// protocol version constraints, using the provided function to check idleness.
// The resulting set of peers are sorted by their measure throughput.
func (ps *peerSet) idlePeers(minProtocol, maxProtocol int, idleCheck func(*peerConnection) bool, throughput func(*peerConnection) float64) ([]*peerConnection, int) {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	idle, total := make([]*peerConnection, 0, len(ps.peers)), 0
	for _, p := range ps.peers {
		if idleCheck(p) {
			idle = append(idle, p)
		}
		total++
	}
	for i := 0; i < len(idle); i++ {
		for j := i + 1; j < len(idle); j++ {
			if throughput(idle[i]) < throughput(idle[j]) {
				idle[i], idle[j] = idle[j], idle[i]
			}
		}
	}
	return idle, total
}

// BlockCapacity retrieves the peers block download allowance based on its
// previously discovered throughput.
func (p *peerConnection) BlockCapacity(targetRTT time.Duration) int {
	p.lock.RLock()
	defer p.lock.RUnlock()

	return int(math.Min(1+math.Max(1, p.blockThroughput*float64(targetRTT)/float64(time.Second)), float64(MaxBlockFetch)))
}

// SetBlocksIdle sets the peer to idle, allowing it to execute new block retrieval
// requests. Its estimated block retrieval throughput is updated with that measured
// just now.
func (p *peerConnection) SetBlocksIdle(delivered int) {
	p.setIdle(p.blockStarted, delivered, &p.blockThroughput, &p.blockIdle)
}

// setIdle sets the peer to idle, allowing it to execute new retrieval requests.
// Its estimated retrieval throughput is updated with that measured just now.
func (p *peerConnection) setIdle(started time.Time, delivered int, throughput *float64, idle *int32) {
	// Irrelevant of the scaling, make sure the peer ends up idle
	defer atomic.StoreInt32(idle, 0)

	p.lock.Lock()
	defer p.lock.Unlock()

	// If nothing was delivered (hard timeout / unavailable data), reduce throughput to minimum
	if delivered == 0 {
		*throughput = 0
		return
	}
	// Otherwise update the throughput with a new measurement
	elapsed := time.Since(started) + 1 // +1 (ns) to ensure non-zero divisor
	measured := float64(delivered) / (float64(elapsed) / float64(time.Second))

	*throughput = (1-measurementImpact)*(*throughput) + measurementImpact*measured
	p.rtt = time.Duration((1-measurementImpact)*float64(p.rtt) + measurementImpact*float64(elapsed))

	p.log.Trace("Peer throughput measurements updated",
		"hps", p.blockThroughput,
		"miss", len(p.lacking), "rtt", p.rtt)
}

type peerSet struct {
	peers  map[string]*peerConnection
	lock   sync.RWMutex
	closed bool
}

func newPeerSet() *peerSet {
	return &peerSet{
		peers: make(map[string]*peerConnection),
	}
}

func (ps *peerSet) Len() int {
	return len(ps.peers)
}

func (ps *peerSet) Reset() {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	for _, peer := range ps.peers {
		peer.Reset()
	}

}
func (ps *peerSet) Register(p *peerConnection) error {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	if ps.closed {
		return errClosed
	}

	id := p.id

	if _, ok := ps.peers[id]; ok {
		return errAlreadyRegistered
	}

	ps.peers[id] = p
	return nil
}

func (ps *peerSet) Unregister(id string) error {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	if _, ok := ps.peers[id]; !ok {
		return errNotRegistered
	}
	delete(ps.peers, id)
	return nil
}

func (ps *peerSet) Operator() *peerConnection {
	for _, p := range ps.peers {
		if p.peer.IsOperator() {
			return p
		}
	}

	return nil
}

func (ps *peerSet) Peer(id string) *peerConnection {
	return ps.peers[id]
}

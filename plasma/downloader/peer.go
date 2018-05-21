package downloader

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/plasma/types"
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

	peer peer

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

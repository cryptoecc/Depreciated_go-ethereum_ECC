package plasma

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/plasma/types"
	set "gopkg.in/fatih/set.v0"
)

var (
	errClosed            = errors.New("peer set is closed")
	errAlreadyRegistered = errors.New("peer is already registered")
	errNotRegistered     = errors.New("peer is not registered")
)

const (
	maxKnownTxs      = 32768 // Maximum transactions hashes to keep in the known list (prevent DOS)
	maxKnownBlocks   = 1024  // Maximum block hashes to keep in the known list (prevent DOS)
	handshakeTimeout = 5 * time.Second
)

// Peer represents a plasma protocol peer connection.
type Peer struct {
	id string

	*p2p.Peer
	rw p2p.MsgReadWriter

	currentBlockNumber *big.Int
	knownTxs           *set.Set // Set of transaction hashes known to be known by this peer
	knownBlocks        *set.Set // Set of block hashes known to be known by this peer

	operator bool

	quit chan struct{}
}

// PeerInfo represents a short summary of the Ethereum sub-protocol metadata known
// about a connected peer.
type PeerInfo struct {
	CurrentBlockNumber uint64 `json:"currentBlockNumber"`
}

// newPeer creates a new plasma peer object, but does not run the handshake itself.
func newPeer(host *Plasma, remote *p2p.Peer, rw p2p.MsgReadWriter) *Peer {
	id := remote.ID()

	pubkey, err := remote.ID().Pubkey()
	peerAddress := crypto.PubkeyToAddress(*pubkey)

	if err != nil {
		return nil
	}

	operator := bytes.Equal(peerAddress.Bytes(), host.config.OperatorAddress.Bytes())

	return &Peer{
		Peer: remote,

		id: fmt.Sprintf("%x", id[:8]),
		rw: rw,

		knownTxs:    set.New(),
		knownBlocks: set.New(),

		operator: operator,

		quit: make(chan struct{}),
	}
}

// Info gathers and returns a collection of metadata known about a peer.
func (p *Peer) Info() *PeerInfo {
	return &PeerInfo{
		CurrentBlockNumber: p.currentBlockNumber.Uint64(),
	}
}

func (p *Peer) Log() log.Logger { return p.peer.Log() }

// start initiates the peer updater, periodically broadcasting the plasma packets
// into the network.
func (p *Peer) start() error {
	log.Trace("pls", "start", "peer", p.ID())
	return p.update()
}

// stop terminates the peer updater, stopping message forwarding to it.
func (p *Peer) stop() {
	close(p.quit)
	log.Trace("pls", "stop", "peer", p.ID())
}

// handshake sends the protocol initiation status message to the remote peer and
// verifies the remote status too.
func (p *Peer) handshake(config *Config) error {
	// Send the handshake status message asynchronously
	errc := make(chan error, 1)
	go func() {
		query := statusData{
			ProtocolVersion: ProtocolVersion,
			OperatorAddress: config.OperatorAddress,
			ContractAddress: config.ContractAddress,
			HighestEthBlock: 0, // TODO: read ethereum highest block
		}

		errc <- p2p.Send(p.rw, StatusCode, query)
	}()

	// Fetch the remote status packet and verify protocol match
	packet, err := p.rw.ReadMsg()

	if err != nil {
		return err
	}
	if packet.Code != StatusCode {
		return fmt.Errorf("peer [%x] sent packet %x before status packet", p.ID(), packet.Code)
	}
	var query statusData

	if err := packet.Decode(&query); err != nil {
		return err
	}

	if err != nil {
		return fmt.Errorf("peer [%x] sent bad status message: %v", p.ID(), err)
	}
	if query.ProtocolVersion != ProtocolVersion {
		return fmt.Errorf("peer [%x]: protocol version mismatch %d != %d", p.ID(), query.ProtocolVersion, ProtocolVersion)
	}

	// Wait until out own status is consumed too
	if err := <-errc; err != nil {
		return fmt.Errorf("peer [%x] failed to send status packet: %v", p.ID(), err)
	}

	return nil
}

// update broadcast and sync Plasma block with peer
/*
 * TODO: below update should be
 *  1. retrieve child block number recorded at Plasma contract in root chain
 *   - this loop should be in side of the Plasma's loop
 *  2. if host's blockchain is behind it, request the block to peers
 *  3. if the peer's blockchain is behind it, wait until the peer's request is arrived.
 *   - do nothing, just response
 */
func (p *Peer) update() error {
	for {
		select {
		case <-p.quit:
			break
		}
	}

	return nil
}

// mark marks an block known to the peer so that it won't be sent back.
func (peer *Peer) markBlock(block *types.Block) {
	peer.knownBlocks.Add(block.Data.BlockNumber.Uint64())
}

// mark marks an block known to the peer so that it won't be sent back.
func (peer *Peer) markTransaction(tx *types.Transaction) {
	peer.knownTxs.Add(tx.Hash())
}

// marked checks if an block is already known to the remote peer.
func (peer *Peer) markedBlock(blkNum uint64) bool {
	return peer.knownBlocks.Has(blkNum)
}

// marked checks if an block is already known to the remote peer.
func (peer *Peer) markedTransaction(hash common.Hash) bool {
	return peer.knownTxs.Has(hash)
}

// broadcast iterates over transaction pool
func (p *Peer) broadcast() error {
	return nil
}

// send operator info
func (p *Peer) SendOperator() error {
	return p2p.Send(p.rw, OperatorCode, []interface{}{p.host.config.OperatorNodeURL})
}

// send a single block
func (p *Peer) SendNewBlock(block *types.Block) error {
	return p2p.Send(p.rw, NewBlockCode, block)
}

// send a single block
func (p *Peer) SendNewBlocks(blocks []*types.Block) error {
	return p2p.Send(p.rw, NewBlocksCode, blocks)
}

// send transaction
func (p *Peer) SendNewTransactions(txs []*types.Transaction) error {
	return p2p.Send(p.rw, NewTransactionsCode, txs)
}

func (p *Peer) RequestBlock(blkNum uint64) error {
	p.Log().Info("Fetching plasma block", "blkNum", blkNum)

	return p2p.Send(p.rw, GetBlockCode, blkNum)
}

func (p *Peer) RequestBlocks(blkNums []uint64) error {
	p.Log().Info("Fetching plasma block", "blkNums", blkNums)

	return p2p.Send(p.rw, GetBlocksCode, blkNums)
}

func (p *Peer) ID() []byte {
	id := p.peer.ID()
	return id[:]
}

func errResp(code errCode, format string, v ...interface{}) error {
	return fmt.Errorf("%v - %v", code, fmt.Sprintf(format, v...))
}

type peerSet struct {
	peers  map[string]*Peer
	lock   sync.RWMutex
	closed bool
}

func newPeerSet() *peerSet {
	return &peerSet{
		peers: make(map[string]*Peer),
	}
}

// Register injects a new peer into the working set, or returns an error if the
// peer is already known.
func (ps *peerSet) Register(p *Peer) error {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	if ps.closed {
		return errClosed
	}

	id := p.id

	if _, ok := ps.peers[id]; ok {
		return errAlreadyRegistered
	}

	ps.peers[p.id] = p
	return nil
}

// Unregister removes a remote peer from the active set, disabling any further
// actions to/from that particular entity.
func (ps *peerSet) Unregister(id string) error {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	if _, ok := ps.peers[id]; !ok {
		return errNotRegistered
	}
	delete(ps.peers, id)
	return nil
}

// Peer retrieves the registered peer with the given id.
func (ps *peerSet) Peer(id string) *Peer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	return ps.peers[id]
}

// Len returns if the current number of peers in the set.
func (ps *peerSet) Len() int {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	return len(ps.peers)
}

// PeersWithoutBlock retrieves a list of peers that do not have a given block in
// their set of known hashes.
func (ps *peerSet) PeersWithoutBlock(blkNum uint64) []*Peer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	list := make([]*Peer, 0, len(ps.peers))
	for _, p := range ps.peers {
		if !p.markedBlock(blkNum) {
			list = append(list, p)
		}
	}
	return list
}

// PeersWithoutTransaction retrieves a list of peers that do not have a given tx in
// their set of known hashes.
func (ps *peerSet) PeersWithoutTransaction(hash common.Hash) []*Peer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	list := make([]*Peer, 0, len(ps.peers))
	for _, p := range ps.peers {
		if !p.markedTransaction(hash) {
			list = append(list, p)
		}
	}
	return list
}

// Operator returns operator pls peer
func (ps *peerSet) Operator() *Peer {
	ps.lock.RLock()
	defer ps.lock.RUnlock()

	for _, p := range ps.peers {
		if p.operator {
			return p
		}
	}
	return nil
}

// Close disconnects all peers.
// No new peers can be registered after Close has returned.
func (ps *peerSet) Close() {
	ps.lock.Lock()
	defer ps.lock.Unlock()

	for _, p := range ps.peers {
		p.peer.Disconnect(p2p.DiscQuitting)
	}
	ps.closed = true
}

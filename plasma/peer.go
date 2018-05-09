package plasma

import (
	"fmt"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/plasma/types"
	set "gopkg.in/fatih/set.v0"
)

// Peer represents a plasma protocol peer connection.
type Peer struct {
	host     *Plasma
	peer     *p2p.Peer
	rw       p2p.MsgReadWriter
	operator bool

	known *set.Set // Blocks already known by the peer to avoid wasting bandwidth

	quit chan struct{}
}

// newPeer creates a new plasma peer object, but does not run the handshake itself.
func newPeer(host *Plasma, remote *p2p.Peer, rw p2p.MsgReadWriter) *Peer {
	return &Peer{
		host:  host,
		peer:  remote,
		rw:    rw,
		known: set.New(),
		quit:  make(chan struct{}),
	}
}

// start initiates the peer updater, periodically broadcasting the plasma packets
// into the network.
func (p *Peer) start() {
	go p.update()
	log.Trace("pls", "start", "peer", p.ID())
}

// stop terminates the peer updater, stopping message forwarding to it.
func (p *Peer) stop() {
	close(p.quit)
	log.Trace("pls", "stop", "peer", p.ID())
}

// handshake sends the protocol initiation status message to the remote peer and
// verifies the remote status too.
func (p *Peer) handshake() error {
	// Send the handshake status message asynchronously
	errc := make(chan error, 1)
	go func() {
		query := statusData{
			ProtocolVersion: ProtocolVersion,
			OperatorAddress: p.host.config.OperatorAddress,
			ContractAddress: p.host.config.ContractAddress,
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
func (p *Peer) update() {
	for {
		select {
		case <-p.quit:
			break
		}
	}
}

// mark marks an block known to the peer so that it won't be sent back.
func (peer *Peer) markBlock(block *types.Block) {
	peer.known.Add(block.Hash())
}

// mark marks an block known to the peer so that it won't be sent back.
func (peer *Peer) markTransaction(tx *types.Transaction) {
	peer.known.Add(tx.Hash())
}

// marked checks if an block is already known to the remote peer.
func (peer *Peer) marked(block *types.Block) bool {
	return peer.known.Has(block.Hash())
}

// broadcast iterates over transaction pool
func (p *Peer) broadcast() error {
	return nil
}

func (p *Peer) ID() []byte {
	id := p.peer.ID()
	return id[:]
}

func errResp(code errCode, format string, v ...interface{}) error {
	return fmt.Errorf("%v - %v", code, fmt.Sprintf(format, v...))
}

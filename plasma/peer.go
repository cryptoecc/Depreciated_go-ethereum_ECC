package plasma

import (
	"fmt"
	// "time"

	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rlp"
	set "gopkg.in/fatih/set.v0"
)

// Peer represents a plasma protocol peer connection.
type Peer struct {
	host    *Plasma
	peer    *p2p.Peer
	rw      p2p.MsgReadWriter
	trusted bool

	known *set.Set // Blocks already known by the peer to avoid wasting bandwidth

	quit chan struct{}
}

// newPeer creates a new plasma peer object, but does not run the handshake itself.
func newPeer(host *Plasma, remote *p2p.Peer, rw p2p.MsgReadWriter) *Peer {
	return &Peer{
		host:    host,
		peer:    remote,
		rw:      rw,
		trusted: false,
		known:   set.New(),
		quit:    make(chan struct{}),
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
		errc <- p2p.Send(p.rw, statusCode, ProtocolVersion)
	}()
	// Fetch the remote status packet and verify protocol match
	packet, err := p.rw.ReadMsg()
	if err != nil {
		return err
	}
	if packet.Code != statusCode {
		return fmt.Errorf("peer [%x] sent packet %x before status packet", p.ID(), packet.Code)
	}
	s := rlp.NewStream(packet.Payload, uint64(packet.Size))
	peerVersion, err := s.Uint()
	if err != nil {
		return fmt.Errorf("peer [%x] sent bad status message: %v", p.ID(), err)
	}
	if peerVersion != ProtocolVersion {
		return fmt.Errorf("peer [%x]: protocol version mismatch %d != %d", p.ID(), peerVersion, ProtocolVersion)
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
func (peer *Peer) mark(block *Block) {
	peer.known.Add(block.Hash())
}

// marked checks if an block is already known to the remote peer.
func (peer *Peer) marked(block *Block) bool {
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

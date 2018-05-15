package plasma

import (
	// "encoding/json"
	// "errors"
	// "fmt"
	// "math"
	// "math/big"
	"sync"
	// "sync/atomic"
	// "time"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/core"
	// "github.com/ethereum/go-ethereum/core/types"
	// "github.com/ethereum/go-ethereum/eth/downloader"
	// "github.com/ethereum/go-ethereum/eth/fetcher"
	// "github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	// "github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	// "github.com/ethereum/go-ethereum/p2p/discover"
	// "github.com/ethereum/go-ethereum/params"
	// "github.com/ethereum/go-ethereum/rlp"
)

const (
	softResponseLimit = 2 * 1024 * 1024 // Target maximum size of returned blocks, headers or node data.
	estHeaderRlpSize  = 500             // Approximate size of an RLP encoded block header

	// txChanSize is the size of channel listening to TxPreEvent.
	// The number is referenced from the size of tx pool.
	txChanSize = 4096
)

type ProtocolManager struct {
	// txpool      txPool
	blockchain *BlockChain
	config     *Config
	maxPeers   int

	// downloader *downloader.Downloader
	// fetcher    *fetcher.Fetcher
	peers *peerSet

	Protocol p2p.Protocol

	eventMux *event.TypeMux
	// txCh          chan core.TxPreEvent
	// txSub         event.Subscription
	// minedBlockSub *event.TypeMuxSubscription

	// channels for fetcher, syncer, txsyncLoop
	newPeerCh chan *Peer
	// txsyncCh    chan *txsync
	quitSync    chan struct{}
	noMorePeers chan struct{}

	// wait group is used for graceful shutdowns during downloading
	// and processing
	wg sync.WaitGroup
}

// NewProtocolManager returns a new plasma protocol manager.
// The protocol manages peers capable with the plasma network.
// , txpool txPool, chaindb ethdb.Database
func NewProtocolManager(config *Config, mux *event.TypeMux, blockchain *BlockChain) (*ProtocolManager, error) {
	// func NewProtocolManager(config *Config, mode downloader.SyncMode, networkId uint64, mux *event.TypeMux, txpool txPool, engine consensus.Engine, blockchain *core.BlockChain, chaindb ethdb.Database) (*ProtocolManager, error) {
	// Create the protocol manager with the base fields
	manager := &ProtocolManager{
		eventMux:    mux,
		blockchain:  blockchain,
		config:      config,
		peers:       newPeerSet(),
		newPeerCh:   make(chan *Peer),
		noMorePeers: make(chan struct{}),
		// txsyncCh:    make(chan *txsync),
		quitSync: make(chan struct{}),
	}

	version := ProtocolVersion

	manager.Protocol = p2p.Protocol{
		Name:    ProtocolName,
		Version: ProtocolVersion,
		Length:  ProtocolLength,
		Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
			peer := manager.newPeer(int(version), p, rw)
			select {
			case manager.newPeerCh <- peer:
				manager.wg.Add(1)
				defer manager.wg.Done()
				return manager.handle(peer)
			case <-manager.quitSync:
				return p2p.DiscQuitting
			}
		},
		NodeInfo: func() interface{} {
			return manager.NodeInfo()
		},
		PeerInfo: func(id discover.NodeID) interface{} {
			if p := manager.peers.Peer(fmt.Sprintf("%x", id[:8])); p != nil {
				return p.Info()
			}
			return nil
		},
	}

	if len(manager.SubProtocols) == 0 {
		return nil, errIncompatibleConfig
	}
	// Construct the different synchronisation mechanisms
	manager.downloader = downloader.New(mode, chaindb, manager.eventMux, blockchain, nil, manager.removePeer)

	validator := func(header *types.Header) error {
		return engine.VerifyHeader(blockchain, header, true)
	}
	heighter := func() uint64 {
		return blockchain.CurrentBlock().NumberU64()
	}
	inserter := func(blocks types.Blocks) (int, error) {
		// If fast sync is running, deny importing weird blocks
		if atomic.LoadUint32(&manager.fastSync) == 1 {
			log.Warn("Discarded bad propagated block", "number", blocks[0].Number(), "hash", blocks[0].Hash())
			return 0, nil
		}
		atomic.StoreUint32(&manager.acceptTxs, 1) // Mark initial sync done on any fetcher import
		return manager.blockchain.InsertChain(blocks)
	}
	manager.fetcher = fetcher.New(blockchain.GetBlockByHash, validator, manager.BroadcastBlock, heighter, inserter, manager.removePeer)

	return manager, nil
}

func (pm *ProtocolManager) newPeer(p *p2p.Peer, rw p2p.MsgReadWriter) *Peer {
	return newPeer(pv, p, newMeteredMsgWriter(rw))
}

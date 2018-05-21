package plasma

import (
	// "encoding/json"
	// "errors"
	// "bytes"
	"fmt"
	// "math"
	"math/big"
	"sync"
	// "sync/atomic"
	// "time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/plasma/downloader"
	"github.com/ethereum/go-ethereum/plasma/types"
	// "github.com/ethereum/go-ethereum/eth/fetcher"
	// "github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	// "github.com/ethereum/go-ethereum/plasma/types"
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
	acceptTxs uint32 // Flag whether we're considered synchronised (enables transaction processing)

	// txpool      txPool
	blockchain *BlockChain
	config     *Config
	maxPeers   int
	operator   bool

	downloader *downloader.Downloader
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
	quitSync       chan struct{}
	noMorePeers    chan struct{}
	operatorNodeCh chan *Peer

	// wait group is used for graceful shutdowns during downloading
	// and processing
	wg sync.WaitGroup
}

// NewProtocolManager returns a new plasma protocol manager.
// The protocol manages peers capable with the plasma network.
// , txpool txPool, chaindb ethdb.Database
func NewProtocolManager(config *Config, operator bool, mux *event.TypeMux, blockchain *BlockChain, highter func() []uint64) *ProtocolManager {
	// func NewProtocolManager(config *Config, mode downloader.SyncMode, networkId uint64, mux *event.TypeMux, txpool txPool, engine consensus.Engine, blockchain *core.BlockChain, chaindb ethdb.Database) (*ProtocolManager, error) {
	// Create the protocol manager with the base fields

	manager := &ProtocolManager{
		eventMux:    mux,
		blockchain:  blockchain,
		config:      config,
		operator:    operator,
		peers:       newPeerSet(),
		newPeerCh:   make(chan *Peer, 1),
		noMorePeers: make(chan struct{}),
		// txsyncCh:    make(chan *txsync),
		operatorNodeCh: make(chan *Peer, 1),
		quitSync:       make(chan struct{}),
	}

	manager.Protocol = p2p.Protocol{
		Name:    ProtocolName,
		Version: ProtocolVersion,
		Length:  ProtocolLength,
		Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
			log.Info("[Plasma] protocol manager run...")
			peer := newPeer(p, rw, config)
			select {
			case manager.newPeerCh <- peer:
				manager.wg.Add(1)
				defer manager.wg.Done()
				return manager.handle(peer, config)
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

	// Construct the different synchronisation mechanisms
	dropPeer := manager.removePeer
	manager.downloader = downloader.New(blockchain, mux, highter, dropPeer)

	// validator := func(header *types.Header) error {
	// 	return engine.VerifyHeader(blockchain, header, true)
	// }

	// TODO: fetch deposit blocks
	// heighter := func() uint64 {
	// 	return blockchain.currentBlockNumber.Uint64()
	// }
	// inserter := func(blocks []*types.Block) (int, error) {
	// 	atomic.StoreUint32(&manager.acceptTxs, 1) // Mark initial sync done on any fetcher import
	// 	return manager.blockchain.AddBlocks(blocks)
	// }
	// manager.fetcher = fetcher.New(blockchain.GetBlockByHash, validator, manager.BroadcastBlock, heighter, inserter, manager.removePeer)

	return manager
}

func (pm *ProtocolManager) removePeer(id string) {
	// Short circuit if the peer was already removed
	peer := pm.peers.Peer(id)
	if peer == nil {
		return
	}
	log.Debug("Removing Plasma peer", "peer", id)

	// Unregister the peer from the downloader and Ethereum peer set
	// pm.downloader.UnregisterPeer(id)
	if err := pm.peers.Unregister(id); err != nil {
		log.Error("Peer removal failed", "peer", id, "err", err)
	}
	// Hard disconnect at the networking layer
	if peer != nil {
		peer.Peer.Disconnect(p2p.DiscUselessPeer)
	}
}

func (pm *ProtocolManager) Start(maxPeers int) {
	pm.maxPeers = maxPeers

	// broadcast transactions
	// pm.txCh = make(chan core.TxPreEvent, txChanSize)
	// pm.txSub = pm.txpool.SubscribeTxPreEvent(pm.txCh)
	// go pm.txBroadcastLoop()

	// broadcast mined blocks
	// pm.minedBlockSub = pm.eventMux.Subscribe(core.NewMinedBlockEvent{})
	// go pm.minedBroadcastLoop()

	// start sync handlers
	go pm.syncer()
	// go pm.txsyncLoop()
}

func (pm *ProtocolManager) Stop() {
	pm.peers.Close()
	close(pm.quitSync)
}

// handle is the callback invoked to manage the life cycle of an eth peer. When
// this function terminates, the peer is disconnected.
// XXX the life sycle of handle function is same as the p2p connection
func (pm *ProtocolManager) handle(peer *Peer, config *Config) error {
	log.Info("[Plasma] manager.handle()")
	// Check mas peer
	if pm.peers.Len() >= pm.maxPeers {
		log.Warn("[Plasma] too many peers")
		return p2p.DiscTooManyPeers
	}

	// pls handshake
	if err := peer.handshake(pm.blockchain.GetCurrentBlockNumber(), pm.config); err != nil {
		log.Warn("[Plasma] failed to handshake", "err", err)
		return err
	}

	peer.Log().Info("Plasma peer connected", "name", peer.Name(), "operator", peer.operator)

	// go func() {
	// 	ticker := time.NewTicker(5 * time.Second)
	//
	// 	for {
	// 		select {
	// 		case <-ticker.C:
	// 			peer.SendPing(1)
	// 		}
	// 	}
	// }()

	// if peer is operator, record it
	if peer.operator {
		config.OperatorNode = peer

		pm.operatorNodeCh <- peer
	}

	if err := pm.peers.Register(peer); err != nil {
		return err
	}
	defer pm.removePeer(peer.id)

	// // Register the peer in the downloader. If the downloader considers it banned, we disconnect
	if err := pm.downloader.RegisterPeer(peer.id, peer); err != nil {
		return err
	}
	// // Propagate existing transactions. new transactions appearing
	// // after this will be sent via broadcasts.
	// pm.syncTransactions(p)

	// main loop. handle incoming messages.
	for {
		if err := pm.handleMsg(peer); err != nil {
			peer.Log().Debug("Plasma message handling failed", "err", err)
			return err
		}
	}
}

// handle p2p message after handshake
func (pm *ProtocolManager) handleMsg(peer *Peer) error {
	rw := peer.rw

	packet, err := rw.ReadMsg()

	if err != nil {
		log.Warn("[Plasma] message loop failed", "peer", peer.id, "err", err)
		peer.stop()
		return err
	}

	log.Info("[Plasma] p2p message received", "code", packet.Code)

	switch packet.Code {
	case StatusCode:
		// this should not happen, but no need to panic; just ignore this message.
		log.Warn("unxepected status message received", "peer", peer.ID())

	case OperatorCode:
		// TODO: send operator node info for p2p conenction
		// if !pls.isOperator() && pls.config.OperatorNodeURL == "" {
		// 	var query operatorData
		// 	if err := packet.Decode(&query); err != nil {
		// 		return errResp(ErrDecode, "%v: %v", packet, err)
		// 	}
		//
		// 	pls.config.OperatorNodeURL = query.NodeURL
		// }

	case GetBlockCode: // request a single block
		var blkNum uint64

		if err := packet.Decode(&blkNum); err != nil {
			return errResp(ErrDecode, "%v: %v", packet, err)
		}

		log.Info("[Plasma] Querying a single block", "blkNum", blkNum)

		block, err := pm.blockchain.GetBlock(big.NewInt(int64(blkNum)))

		if err != nil {
			log.Warn("[Plasma] Failed to get block", "blkNum", blkNum, "err", err)
			return err
		}

		if block != nil {
			return peer.SendNewBlocks([]*types.Block{block})
		}

	case GetBlocksCode: // request batch of blocks
		var blkNums []uint64

		if err := packet.Decode(&blkNums); err != nil {
			return errResp(ErrDecode, "%v: %v", packet, err)
		}

		log.Info("[Plasma] Querying batch of blocks", "blkNums", blkNums)

		blocks := make(types.Blocks, 0, len(blkNums))

		for _, blkNum := range blkNums {
			block, err := pm.blockchain.GetBlock(big.NewInt(int64(blkNum)))

			if err != nil {
				log.Warn("[Plasma] Failed to get block", "blkNum", blkNum, "err", err)
				return err
			}

			if block != nil {
				blocks = append(blocks, block)
			}
		}

		if err := peer.SendNewBlocks(blocks); err != nil {
			log.Warn("[Plasma] Failed to send blocks to peer", "err", err)
			return err
		}

	case NewBlockCode: // received a block
		var block types.Block

		if err := packet.Decode(&block); err != nil {
			return errResp(ErrDecode, "%v: %v", packet, err)
		}

		log.Info("[Plasma] new block received", "hash", block.Hash(), "blkNum", block.Data.BlockNumber)

		if err := pm.downloader.DeliverBlock(peer.id, &block); err != nil {
			log.Warn("[Plasma] failed to add block", "hash", block.Hash(), "blkNum", block.Data.BlockNumber, "err", err)
			return err
		}

		log.Info("[Plasma] imported new block", "blkNum", block.Data.BlockNumber, "peer", peer.ID())

	case NewBlocksCode: // received batch of blocks
		var blocks types.Blocks

		if err := packet.Decode(&blocks); err != nil {
			return errResp(ErrDecode, "%v: %v", packet, err)
		}

		first := blocks[0]
		last := blocks[len(blocks)-1]

		log.Info("[Plasma] new batch of blocks received", "from", first.NumberU64(), "to", last.NumberU64())

		if err := pm.downloader.DeliverBlocks(peer.id, blocks); err != nil {
			log.Warn("[Plasma] failed to add blocks", "err", err)
			return err
		}

		log.Info("[Plasma] imported new blocks", "from", first.Data.BlockNumber, "to", last.Data.BlockNumber)

	case PingCode:
		var query pingData
		if err := packet.Decode(&query); err != nil {
			log.Info("[Plasma] Failed to decode ping", "err", err)
			return err
		}
		log.Info("[Plasma] ping received. send pong", "peer", peer.id, "query", query.Number)

		// get plasma block with number 1
		blk, _ := pm.blockchain.GetBlock(big1)

		if blk == nil {
			log.Info("[Plasma] No block with number 1. Do not pong")
		} else {
			payload := pongData{
				Block: blk,
			}

			if err := p2p.Send(peer.rw, PongCode, payload); err != nil {
				log.Info("[Plasma] Failed to send pong message", "err", err)
				return err
			}
		}

	case PongCode:
		var query pongData
		if err := packet.Decode(&query); err != nil {
			log.Info("[Plasma] Failed to decode pong", "err", err)
			return errResp(ErrDecode, "%v: %v", packet, err)
		}
		log.Info("[Plasma] pong received", "peer", peer.id, "blockNumber", query.Block.Data.BlockNumber)

	default:
		return errResp(ErrInvalidMsgCode, "%v", packet.Code)
	}

	return nil
}

// NodeInfo represents a short summary of the Plasma sub-protocol metadata
// known about the host peer.
type NodeInfo struct {
	CurrentBlockNumber uint64         `json:"currentBlockNumber"` // Current block number
	OperatorAddress    common.Address `json:"operatorAddress"`    // Operator Address
	ContractAddress    common.Address `json:"contractAddress"`    // Contract Address
}

// NodeInfo retrieves some protocol metadata about the running host node.
func (pm *ProtocolManager) NodeInfo() *NodeInfo {
	return &NodeInfo{
		CurrentBlockNumber: pm.blockchain.GetCurrentBlockNumber().Uint64(),
		OperatorAddress:    pm.config.OperatorAddress,
		ContractAddress:    pm.config.ContractAddress,
	}
}

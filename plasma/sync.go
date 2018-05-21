package plasma

import (
	// "math/rand"
	"sync/atomic"
	"time"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/plasma/types"

	"github.com/ethereum/go-ethereum/log"
	// "github.com/ethereum/go-ethereum/p2p/discover"
)

const (
	forceSyncCycle      = 10 * time.Second // Time interval to force syncs, even if few peers are available
	minDesiredPeerCount = 5                // Amount of peers desired to start syncing

	// This is the target size for the packs of transactions sent by txsyncLoop.
	// A pack can get larger than this if a single transactions exceeds this size.
	txsyncPackSize = 100 * 1024
)

// syncer is responsible for periodically synchronising with the network, both
// downloading hashes and blocks as well as handling the announcement handler.
func (pm *ProtocolManager) syncer() {
	if pm.operator {
		return
	}

	// Start and ensure cleanup of sync mechanisms
	// pm.fetcher.Start()
	// defer pm.fetcher.Stop()
	defer pm.downloader.Terminate()

	// Wait for different events to fire synchronisation operations
	forceSync := time.NewTicker(forceSyncCycle)
	defer forceSync.Stop()

	for {
		select {
		case <-pm.newPeerCh:
			log.Info("[Plasma] case <-pm.newPeerCh:")
			go pm.synchronise(pm.peers.Operator())

		case <-forceSync.C:
			log.Info("[Plasma] case <-forceSync.C:")
			// Force a sync even if not enough peers are present
			go pm.synchronise(pm.peers.Operator())

		case <-pm.noMorePeers:
			return
		}
	}
}

// synchronise tries to sync up our local block chain with a remote peer.
func (pm *ProtocolManager) synchronise(peer *Peer) {

	// Short circuit if no peers are available
	if peer == nil {
		log.Warn("[Plasma] Syncing without peer")
		return
	}

	pBlkNum := peer.currentBlockNumber

	// Run the sync cycle
	if err := pm.downloader.Synchronise(peer.id, pBlkNum); err != nil {
		log.Warn("[Plasma] Faild to syncrhonise", "err", err)
		return
	}

	atomic.StoreUint32(&pm.acceptTxs, 1) // Mark initial sync done
	// if head := pm.blockchain.GetCurrentBlock(); head != nil && head.NumberU64() > 0 {
	// 	// We've completed a sync cycle, notify all peers of new state. This path is
	// 	// essential in star-topology networks where a gateway node needs to notify
	// 	// all its out-of-date peers of the availability of a new block. This failure
	// 	// scenario will most often crop up in private and hackathon networks with
	// 	// degenerate connectivity, but it should be healthy for the mainnet too to
	// 	// more reliably update peers or the local TD state.
	// 	// go pm.BroadcastBlock(head, false)
	// }
}

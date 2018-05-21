package downloader

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/plasma/types"
)

var (
	MaxBlockFetch = 128 // Amount of blocks to be fetched per retrieval request

	maxResultsProcess = 2048 // Number of content download results to import at once into the chain
)

var (
	errBusy                    = errors.New("busy")
	errUnknownPeer             = errors.New("peer is unknown or unhealthy")
	errBadPeer                 = errors.New("action from bad peer ignored")
	errStallingPeer            = errors.New("peer is stalling")
	errNoOperator              = errors.New("Operator not connected yet")
	errNoPeers                 = errors.New("no peers to keep download active")
	errTimeout                 = errors.New("timeout")
	errEmptyHeaderSet          = errors.New("empty header set by peer")
	errPeersUnavailable        = errors.New("no peers available or all tried for download")
	errInvalidAncestor         = errors.New("retrieved ancestor is invalid")
	errInvalidChain            = errors.New("retrieved hash chain is invalid")
	errInvalidBlock            = errors.New("retrieved block is invalid")
	errInvalidBody             = errors.New("retrieved block body is invalid")
	errInvalidReceipt          = errors.New("retrieved receipt is invalid")
	errCancelBlockFetch        = errors.New("block download canceled (requested)")
	errCancelHeaderFetch       = errors.New("block header download canceled (requested)")
	errCancelBodyFetch         = errors.New("block body download canceled (requested)")
	errCancelReceiptFetch      = errors.New("receipt download canceled (requested)")
	errCancelStateFetch        = errors.New("state data download canceled (requested)")
	errCancelHeaderProcessing  = errors.New("header processing canceled (requested)")
	errCancelContentProcessing = errors.New("content processing canceled (requested)")
	errNoSyncActive            = errors.New("no sync active")
	errTooOld                  = errors.New("peer doesn't speak recent enough protocol version (need version >= 62)")
)

// interface for plasma.BlockChain
type BlockChain interface {
	GetCurrentBlock() *types.Block
	GetCurrentBlockNumber() *big.Int
	GetBlock(blkNum *big.Int) (*types.Block, error)
	GetTransaction(blkNum, txIndex *big.Int) (*types.Transaction, error)
	ApplyTransaction(tx *types.Transaction) error
	VerifyTransaction(tx *types.Transaction) error
	SubmitBlock(privKey *ecdsa.PrivateKey) (common.Hash, error)
	NewDeposit(amount *big.Int, depositor *common.Address, depositBlock *big.Int) (common.Hash, error)
	AddNewBlockListener(f func(blk *types.Block) error) error
	AddBlock(b *types.Block) (int, error)
	AddBlocks(b types.Blocks) (int, error)
}

type Downloader struct {
	blockchain BlockChain

	synchronising int32
	notified      int32
	committed     int32

	peers *peerSet
	queue *queue

	blockCh     chan *types.Block
	blockWakeCh chan bool

	mux *event.TypeMux

	// Cancellation and termination
	cancelPeer string         // Identifier of the peer currently being used as the master (cancel on drop)
	cancelCh   chan struct{}  // Channel to cancel mid-flight syncs
	cancelLock sync.RWMutex   // Lock to protect the cancel channel and peer in delivers
	cancelWg   sync.WaitGroup // Make sure all fetcher goroutines have exited.

	quitCh   chan struct{} // Quit channel to signal termination
	quitLock sync.RWMutex  // Lock to prevent double closes

	higheter func() []uint64
	dropPeer func(string)
}

func New(blockchain BlockChain, mux *event.TypeMux, higheter func() []uint64, dropPeer func(id string)) *Downloader {
	d := Downloader{
		blockchain:  blockchain,
		blockCh:     make(chan *types.Block, 1),
		blockWakeCh: make(chan bool, 1),
		quitCh:      make(chan struct{}),
		peers:       newPeerSet(),
		queue:       newQueue(),
		mux:         mux,
		higheter:    higheter,
		dropPeer:    dropPeer,
	}

	return &d
}

// Synchronising returns whether the downloader is currently retrieving blocks.
func (d *Downloader) Synchronising() bool {
	return atomic.LoadInt32(&d.synchronising) > 0
}

func (d *Downloader) AddBlkNums(blkNums []uint64) error {
	operator := d.peers.Operator()

	if operator == nil {
		return errNoOperator
	}

	if len(blkNums) > 0 {
		_, err := d.queue.enqueue(blkNums)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Downloader) RegisterPeer(id string, p peer) error {
	logger := log.New("peer", id)
	logger.Trace("[Plasma] Registering sync peer")

	if err := d.peers.Register(newPeerConnection(id, p, logger)); err != nil {
		logger.Error("[Plasma] Failed to register sync peer", "err", err)
		return err
	}

	return nil
}

// Terminate interrupts the downloader, canceling all pending operations.
// The downloader cannot be reused after calling Terminate.
func (d *Downloader) Terminate() {
	// Close the termination channel (make sure double close is allowed)
	d.quitLock.Lock()
	select {
	case <-d.quitCh:
	default:
		close(d.quitCh)
	}
	d.quitLock.Unlock()

	// Cancel any pending download requests
	d.Cancel()
}

// Cancel aborts all of the operations and waits for all download goroutines to
// finish before returning.
func (d *Downloader) Cancel() {
	d.cancel()
	d.cancelWg.Wait()
}

// cancel aborts all of the operations and resets the queue. However, cancel does
// not wait for the running download goroutines to finish. This method should be
// used when cancelling the downloads from inside the downloader.
func (d *Downloader) cancel() {
	// Close the current cancel channel
	d.cancelLock.Lock()
	if d.cancelCh != nil {
		select {
		case <-d.cancelCh:
			// Channel was already closed
		default:
			close(d.cancelCh)
		}
	}
	d.cancelLock.Unlock()
}

func (d *Downloader) Synchronise(id string, pBlkNum uint64) error {
	err := d.synchronise(id, pBlkNum)
	switch err {
	case nil:
		log.Info("[Plasma] syncronised")
	case errBusy:
		log.Info("[Plasma] already syncing process running")

	case errTimeout, errBadPeer, errStallingPeer,
		errEmptyHeaderSet, errPeersUnavailable, errTooOld,
		errInvalidAncestor, errInvalidChain:
		log.Warn("[Plasma] Synchronisation failed, dropping peer", "peer", id, "err", err)
		if d.dropPeer == nil {
			// The dropPeer method is nil when `--copydb` is used for a local copy.
			// Timeouts can occur if e.g. compaction hits at the wrong time, and can be ignored
			log.Warn("[Plasma] Downloader wants to drop peer, but peerdrop-function is not set", "peer", id)
		} else {
			d.dropPeer(id)
		}
	default:
		log.Warn("[Plasma] Synchronisation failed, retrying", "err", err)
	}
	return err
}

// synchronise will select the peer and use it for synchronising. If an empty string is given
// it will use the best peer possible and synchronize if its TD is higher than our own. If any of the
// checks fail an error will be returned. This method is synchronous
func (d *Downloader) synchronise(id string, pBlkNum uint64) error {
	// Make sure only one goroutine is ever allowed past this point at once
	if !atomic.CompareAndSwapInt32(&d.synchronising, 0, 1) {
		return errBusy
	}
	defer atomic.StoreInt32(&d.synchronising, 0)

	// Post a user notification of the sync (only once per session)
	if atomic.CompareAndSwapInt32(&d.notified, 0, 1) {
		log.Info("[Plasma] Block synchronisation started")
	}
	// Reset the queue, peer set and wake channels to clean any internal leftover state
	d.queue.Reset()
	d.peers.Reset()

	select {
	case <-d.blockCh:
	default:
	}

	select {
	case <-d.blockWakeCh:
	default:
	}

	// Create cancel channel for aborting mid-flight and mark the master peer
	d.cancelLock.Lock()
	d.cancelCh = make(chan struct{})
	d.cancelPeer = id
	d.cancelLock.Unlock()

	defer d.Cancel() // No matter what, we can't leave the cancel channel open

	// Retrieve the origin peer and initiate the downloading process
	p := d.peers.Peer(id)
	if p == nil {
		return errUnknownPeer
	}
	return d.syncWithPeer(p, pBlkNum)
}

// syncWithPeer starts a block synchronization based on the hash chain from the
// specified peer and head hash.
func (d *Downloader) syncWithPeer(p *peerConnection, pBlkNum uint64) (err error) {
	// TODO: uncomment?
	// d.mux.Post(PlsStartEvent{})
	// defer func() {
	// 	// reset on error
	// 	if err != nil {
	// 		d.mux.Post(PlsFailedEvent{err})
	// 	} else {
	// 		d.mux.Post(PlsDoneEvent{})
	// 	}
	// }()

	log.Info("[Plasma] Synchronising with the network", "peer", p.id, "blkNum", pBlkNum)
	defer func(start time.Time) {
		log.Info("[Plasma] Synchronisation terminated", "elapsed", time.Since(start))
	}(time.Now())

	// Look up the sync boundaries: the common ancestor and the target block
	blkNums := d.higheter()

	// Initiate the sync using a concurrent header and content retrieval algorithm

	fetchers := []func() error{
		func() error { return p.peer.RequestBlocks(blkNums) },
		func() error { return d.processSyncContent() },
	}

	return d.spawnSync(fetchers)
}

// spawnSync runs d.process and all given fetcher functions to completion in
// separate goroutines, returning the first error that appears.
func (d *Downloader) spawnSync(fetchers []func() error) error {
	errc := make(chan error, len(fetchers))
	d.cancelWg.Add(len(fetchers))
	for _, fn := range fetchers {
		fn := fn
		go func() { defer d.cancelWg.Done(); errc <- fn() }()
	}
	// Wait for the first error, then terminate the others.
	var err error
	for i := 0; i < len(fetchers); i++ {
		err = <-errc
		if i == len(fetchers)-1 {
			// Close the queue when all fetchers have exited.
			// This will cause the block processor to end when
			// it has processed the queue.
			d.queue.Close()
		}
		if err != nil {
			break
		}
	}
	d.queue.Close()
	d.Cancel()
	return err
}

func (d *Downloader) DeliverBlock(id string, block *types.Block) error {
	return d.queue.deliverBlocks(id, []*types.Block{block})
}

func (d *Downloader) DeliverBlocks(id string, blocks types.Blocks) error {
	return d.queue.deliverBlocks(id, blocks)
}

func (d *Downloader) processSyncContent() error {
	for {
		results := d.queue.Results(true)

		if len(results) == 0 {
			return nil
		}

		_, err := d.blockchain.AddBlocks(results)
		if err != nil {
			return err
		}
	}
}
func (d *Downloader) fillBlocks(blkNums []uint64) {
	operator := d.peers.Operator()

	d.queue.ReserveBlocks(operator)
}

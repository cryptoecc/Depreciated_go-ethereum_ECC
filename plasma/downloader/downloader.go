package downloader

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/plasma/types"
)

var (
	MaxBlockFetch = 128 // Amount of blocks to be fetched per retrieval request

	rttMinEstimate   = 2 * time.Second  // Minimum round-trip time to target for download requests
	rttMaxEstimate   = 20 * time.Second // Maximum round-trip time to target for download requests
	rttMinConfidence = 0.1              // Worse confidence factor in our estimated RTT value
	ttlScaling       = 3                // Constant scaling factor for RTT -> TTL conversion
	ttlLimit         = time.Minute      // Maximum TTL allowance to prevent reaching crazy timeouts

	BlockCheckFrequency = 100             // Verification frequency of the downloaded blocks
	BlockSafetyNet      = 2048            // Number of blocks to discard in case a chain violation is detected
	BlockForceVerify    = 24              // Number of blocks to verify before and after the pivot to accept it
	BlockContCheck      = 3 * time.Second // Time interval to check for block continuations during download
	fsMinFullBlocks     = 64              // Number of blocks to retrieve fully even in fast sync

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

	rttEstimate   uint64 // Round trip time to target for download requests
	rttConfidence uint64 // Confidence in the estimated RTT (unit: millionths to allow atomic ops)

	blockCh     chan dataPack
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
		rttEstimate:   uint64(rttMaxEstimate),
		rttConfidence: uint64(1000000),

		blockchain: blockchain,

		blockCh:     make(chan dataPack, 1),
		blockWakeCh: make(chan bool, 1),
		quitCh:      make(chan struct{}),

		peers: newPeerSet(),
		queue: newQueue(),

		mux: mux,

		higheter: higheter,
		dropPeer: dropPeer,
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
	log.Info("[Plasma] Synchronising with the network", "peer", p.id, "blkNum", pBlkNum)
	defer func(start time.Time) {
		log.Info("[Plasma] Synchronisation terminated", "elapsed", time.Since(start))
	}(time.Now())

	// Look up the blocks to fetch from rootchain contract
	blkNums := d.higheter()
	log.Info("[Plasma] blocks to fetch", "blkNums", blkNums)

	// Initiate the sync using a concurrent header and content retrieval algorithm

	fetchers := []func() error{
		func() error { return d.FetchBlocks(p, blkNums, "blocks") }, // data receiver: request blokcs
		func() error { return d.processSyncContent() },              // block processer: process fetched blocks
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
		if i == len(fetchers)-1 {
			// Close the queue when all data receivers have exited.
			// This will cause the block processor to end when
			// it has processed the queue.
			d.queue.Close()
		}

		if err = <-errc; err != nil {
			log.Warn("[Plasma] d.spawnSync() fetcher exited with err", "err", err)
			break
		}
	}
	d.queue.Close()
	d.Cancel()
	return err
}

func (d *Downloader) DeliverBlock(id string, block *types.Block) error {
	return d.deliver(id, d.blockCh, &blockPack{id, types.Blocks{block}}, blockInMeter, blockDropMeter)
}

func (d *Downloader) DeliverBlocks(id string, blocks types.Blocks) error {
	return d.deliver(id, d.blockCh, &blockPack{id, blocks}, blockInMeter, blockDropMeter)
}

// deliver injects a new batch of data received from a remote node.
func (d *Downloader) deliver(id string, destCh chan dataPack, packet dataPack, inMeter, dropMeter metrics.Meter) (err error) {
	// Update the delivery metrics for both good and failed deliveries
	inMeter.Mark(int64(packet.Items()))
	defer func() {
		if err != nil {
			dropMeter.Mark(int64(packet.Items()))
		}
	}()
	// Deliver or abort if the sync is canceled while queuing
	d.cancelLock.RLock()
	cancel := d.cancelCh
	d.cancelLock.RUnlock()
	if cancel == nil {
		return errNoSyncActive
	}
	select {
	case destCh <- packet:
		return nil
	case <-cancel:
		return errNoSyncActive
	}
}

func (d *Downloader) FetchBlocks(p *peerConnection, blkNums []uint64, kind string) error {
	// Short curcit for empty block fetching
	if len(blkNums) == 0 {
		return nil
	}

	// timer to finish
	var ttl time.Duration
	timeout := time.NewTimer(0) // timer to dump a non-responsive active peer
	<-timeout.C                 // timeout channel should be initially empty
	defer timeout.Stop()

	var (
		deliver = func(packet dataPack) (int, error) {
			pack := packet.(*blockPack)
			n, err := d.queue.deliverBlocks(pack.peerId, pack.blocks)

			return n, err
		}
		expire = func() map[string]int { return d.queue.ExpireBlocks(d.requestTTL()) }
		fetch  = func(p *peerConnection, blkNums []uint64) error {
			ttl = d.requestTTL()
			timeout.Reset(ttl)
			return p.peer.RequestBlocks(blkNums)
		}
		capacity = func(p *peerConnection) int { return p.BlockCapacity(d.requestRTT()) }
		setIdle  = func(p *peerConnection, accepted int) { p.SetBlocksIdle(accepted) }
		reserve  = d.queue.ReserveBlocks

		throttle = d.queue.ShouldThrottleBlocks
		pending  = d.queue.PendingBlocks
		inFlight = d.queue.InFlightBlocks
		// idle     = d.peers.BodyIdlePeers
	)

	// Push blocks to queue
	if err := d.AddBlkNums(blkNums); err != nil {
		return err
	}
	d.blockWakeCh <- false

	// Create a ticker to detect expired retrieval tasks
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	update := make(chan struct{}, 1)

	// Prepare the queue and fetch block parts until the block header fetcher's done
	finished := false

	for {
		select {
		case <-d.cancelCh:
			return errCancelBlockFetch

		case packet := <-d.blockCh:
			// If the peer was previously banned and failed to deliver its pack
			// in a reasonable time frame, ignore its message.
			if peer := d.peers.Peer(packet.PeerId()); peer != nil {
				// Deliver the received chunk of data and check chain validity
				accepted, err := deliver(packet)
				if err == errInvalidChain {
					return err
				}

				// Unless a peer delivered something completely else than requested (usually
				// caused by a timed out request which came through in the end), set it to
				// idle. If the delivery's stale, the peer should have already been idled.
				if err != errStaleDelivery {
					setIdle(peer, accepted)
				}

				// Issue a log to the user to see what's going on
				switch {
				case err == nil && packet.Items() == 0:
					peer.log.Trace("[Plasma] Requested data not delivered", "type", kind)
				case err == nil:
					peer.log.Trace("[Plasma] Delivered new batch of data", "type", kind, "count", packet.Stats())
				default:
					peer.log.Trace("[Plasma] Failed to deliver retrieved data", "type", kind, "err", err)
				}
			}
			// Blocks assembled, try to update the progress
			select {
			case update <- struct{}{}:
			default:
			}

		case cont := <-d.blockWakeCh:
			// The header fetcher sent a continuation flag, check if it's done
			if !cont {
				finished = true
			}
			// Headers arrive, try to update the progress
			select {
			case update <- struct{}{}:
			default:
			}

		case <-ticker.C:
			// Sanity check update the progress
			select {
			case update <- struct{}{}:
			default:
			}

		case <-update:
			log.Info("[Plasma] fetchBlocks case <-update:")
			// Short circuit if we lost all our peers
			if d.peers.Len() == 0 {
				log.Warn("[Plasma] FetchBlock <-update", "err", errNoPeers)
				return errNoPeers
			}
			log.Info("[Plasma] fetchBlocks case <-update:1")

			// Check for fetch request timeouts and demote the responsible peers
			for pid, fails := range expire() {
				log.Info("[Plasma] fetchBlocks case <-update:2")
				if peer := d.peers.Peer(pid); peer != nil {
					// If a lot of retrieval elements expired, we might have overestimated the remote peer or perhaps
					// ourselves. Only reset to minimal throughput but don't drop just yet. If even the minimal times
					// out that sync wise we need to get rid of the peer.
					//
					// The reason the minimum threshold is 2 is because the downloader tries to estimate the bandwidth
					// and latency of a peer separately, which requires pushing the measures capacity a bit and seeing
					// how response times reacts, to it always requests one more than the minimum (i.e. min 2).
					if fails > 2 {
						peer.log.Trace("Data delivery timed out", "type", kind)
						setIdle(peer, 0)
					} else {
						peer.log.Debug("Stalling delivery, dropping", "type", kind)
						if d.dropPeer == nil {
							// The dropPeer method is nil when `--copydb` is used for a local copy.
							// Timeouts can occur if e.g. compaction hits at the wrong time, and can be ignored
							peer.log.Warn("Downloader wants to drop peer, but peerdrop-function is not set", "peer", pid)
						} else {
							d.dropPeer(pid)
						}
					}
				}
			}

			log.Info("[Plasma] fetchBlocks case <-update:3", "!inFlight()", !inFlight(), "finished", finished)

			// If there's nothing more to fetch, wait or terminate
			if pending() == 0 {
				log.Info("[Plasma] fetchBlocks case <-update:4")
				if !inFlight() && finished {
					log.Debug("Data fetching completed", "type", kind)
					return nil
				}
				// TODO: activate?
				break
			}

			log.Info("[Plasma] fetchBlocks case <-update:5")

			// Send a download request to all idle peers, until throttled
			progressed, throttled, running := false, false, inFlight()
			operator := d.peers.Operator()

			if operator == nil {
				log.Warn("[Plasma] Operator not found to fetch")
				continue
			}

			// Short circuit if throttling activated
			if throttle() {
				throttled = true
				log.Warn("[Plasma] throttled...")
				continue
			}

			// Reserve a chunk of fetches for a peer. A nil can mean either that
			// no more headers are available, or that the peer is known not to
			// have them.
			log.Info("[Plasma] capacity(operator)", "capacity(operator)", capacity(operator))

			request, progress, err := reserve(operator, capacity(operator))
			if err != nil {
				log.Warn("[Plasma] failed to reserve to peer", "err", err)
				return err
			}

			if progress {
				progressed = true
			} else {
				log.Warn("[Plasma] No blocks to progress")
				break
			}

			if err := fetch(operator, request.BlkNums); err != nil {
				// Although we could try and make an attempt to fix this, this error really
				// means that we've double allocated a fetch task to a peer. If that is the
				// case, the internal state of the downloader and the queue is very wrong so
				// better hard crash and note the error instead of silently accumulating into
				// a much bigger issue.
				log.Warn("[Plasma]", "err", err)
				panic(fmt.Sprintf("[Plasma] %v: %s fetch assignment failed", operator, kind))
			}
			running = true
			log.Info("[Plasma] fetchBlocks case <-update:6")

			// Make sure that we have peers available for fetching. If all peers have been tried
			// and all failed throw an error
			if !progressed && !throttled && !running && pending() > 0 {
				log.Info("[Plasma] fetchBlocks case <-update:7")
				return errPeersUnavailable
			}
			log.Info("[Plasma] fetchBlocks case <-update:8")

		case <-timeout.C:
			if d.dropPeer == nil {
				// The dropPeer method is nil when `--copydb` is used for a local copy.
				// Timeouts can occur if e.g. compaction hits at the wrong time, and can be ignored
				p.log.Warn("[Peer] Downloader wants to drop peer, but peerdrop-function is not set", "peer", p.id)
				break
			}
			// Header retrieval timed out, consider the peer bad and drop
			p.log.Debug("[Peer]  block request timed out", "elapsed", ttl)
			d.dropPeer(p.id)

			// Finish the sync gracefully instead of dumping the gathered data though
			select {
			case d.blockWakeCh <- false:
			case <-d.cancelCh:
			}

			return errBadPeer
		}
	}

	return p.peer.RequestBlocks(blkNums)
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
	capacity := func(p *peerConnection) int { return p.BlockCapacity(d.requestRTT()) }

	operator := d.peers.Operator()

	d.queue.ReserveBlocks(operator, capacity(operator))
}

// requestRTT returns the current target round trip time for a download request
// to complete in.
//
// Note, the returned RTT is .9 of the actually estimated RTT. The reason is that
// the downloader tries to adapt queries to the RTT, so multiple RTT values can
// be adapted to, but smaller ones are preferred (stabler download stream).
func (d *Downloader) requestRTT() time.Duration {
	return time.Duration(atomic.LoadUint64(&d.rttEstimate)) * 9 / 10
}

// requestTTL returns the current timeout allowance for a single download request
// to finish under.
func (d *Downloader) requestTTL() time.Duration {
	var (
		rtt  = time.Duration(atomic.LoadUint64(&d.rttEstimate))
		conf = float64(atomic.LoadUint64(&d.rttConfidence)) / 1000000.0
	)
	ttl := time.Duration(ttlScaling) * time.Duration(float64(rtt)/conf)
	if ttl > ttlLimit {
		ttl = ttlLimit
	}
	return ttl
}

package downloader

import (
	"errors"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/metrics"
	"github.com/ethereum/go-ethereum/plasma/types"
	"gopkg.in/karalabe/cookiejar.v2/collections/prque"
)

var (
	blockCacheItems      = 8192             // Maximum number of blocks to cache before throttling the download
	blockCacheMemory     = 64 * 1024 * 1024 // Maximum amount of memory to use for block caching
	blockCacheSizeWeight = 0.1              // Multiplier to approximate the average block size based on past ones
)

var (
	errNoFetchesPending = errors.New("no fetches pending")
	errStaleDelivery    = errors.New("stale delivery")
)

type fetchRequest struct {
	Peer    *peerConnection
	BlkNums []uint64
	Time    time.Time
}

type queue struct {
	pq      *prque.Prque             // Priority queue of the block numbers to fetch
	pool    map[uint64]bool          // Requests to be fetched and processed
	pending map[string]*fetchRequest // Requests reserved to peer
	done    map[uint64]bool          // Requests processed and inserted into blockchain

	resultCache types.Blocks       // Downloaded but not yet delivered fetch results
	resultSize  common.StorageSize // Approximate size of a block (exponential moving average)
	lock        *sync.Mutex
	active      *sync.Cond
	closed      bool
}

func newQueue() *queue {
	lock := new(sync.Mutex)

	return &queue{
		pq:      prque.New(),
		pool:    make(map[uint64]bool),
		pending: make(map[string]*fetchRequest),
		done:    make(map[uint64]bool),

		resultCache: make(types.Blocks, 0),
		lock:        lock,
		active:      sync.NewCond(lock),
		closed:      false,
	}
}

func (q *queue) Reset() {
	q.lock.Lock()
	defer q.lock.Unlock()

	log.Info("[Plasma] Reset queue")

	q.closed = false

	q.pool = make(map[uint64]bool)
	q.pq.Reset()
	q.pending = make(map[string]*fetchRequest)
	q.done = make(map[uint64]bool)

	q.resultCache = make(types.Blocks, 0)
}

// Close marks the end of the sync, unblocking WaitResults.
// It may be called even if the queue is already closed.
func (q *queue) Close() {
	q.lock.Lock()
	q.closed = true
	q.lock.Unlock()
	log.Warn("[Plasma] q.Close() awake q.active")
	q.active.Broadcast()
}

func (q *queue) Closed() bool { return q.closed }

func (q *queue) enqueue(blkNums []uint64) ([]uint64, error) {
	log.Info("[Plasma] enqueueing...", "blkNums", blkNums)
	inserts := make([]uint64, 0, len(blkNums))
	for _, blkNum := range blkNums {

		if _, ok := q.pool[blkNum]; ok {
			log.Info("[Plasma] block is already in pending pool", "blkNum", blkNum)
			continue
		}

		if ok := q.done[blkNum]; ok {
			log.Info("[Plasma] block is already fetched", "blkNum", blkNum)
			continue
		}

		inserts = append(inserts, blkNum)

		q.pool[blkNum] = true
		q.pq.Push(blkNum, -float32(blkNum))
	}

	return inserts, nil
}

// ShouldThrottleBlocks checks if the download should be throttled (active block (body)
// fetches exceed block cache).
func (q *queue) ShouldThrottleBlocks() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.resultSlots(q.pending, q.done) <= 0
}

// PendingBlocks retrieves the number of block (body) requests pending for retrieval.
func (q *queue) PendingBlocks() int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.pq.Size()
}

// InFlightBlocks retrieves whether there are block fetch requests currently in
// flight.
func (q *queue) InFlightBlocks() bool {
	q.lock.Lock()
	defer q.lock.Unlock()

	log.Info("[Plasma] infliblocks ", "len(q.pending)", len(q.pending))

	return len(q.pending) > 0
}

// TODO: use count?
func (q *queue) ReserveBlocks(p *peerConnection, count int) (*fetchRequest, bool, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	// Short circuit if the peer's already downloading something (sanity check to
	// not corrupt state)
	if _, ok := q.pending[p.id]; ok {
		return nil, false, nil
	}

	// Short circuit if no item in priority queue
	if q.pq.Empty() {
		return nil, false, nil
	}

	blkNums := []uint64{}
	for i := 0; !q.pq.Empty() && i < count; i++ {
		blkNum := q.pq.PopItem().(uint64)
		log.Info("[Plasma] queue.ReserveBlocks popping", "blkNum", blkNum)
		blkNums = append(blkNums, blkNum)
	}

	request := &fetchRequest{
		Peer:    p,
		BlkNums: blkNums,
		Time:    time.Now(),
	}
	q.pending[p.id] = request

	return request, true, nil
}

// ExpireBlocks checks for in flight block body requests that exceeded a timeout
// allowance, canceling them and returning the responsible peers for penalisation.
func (q *queue) ExpireBlocks(timeout time.Duration) map[string]int {
	q.lock.Lock()
	defer q.lock.Unlock()

	return q.expire(timeout, q.pending, q.pq, blockTimeoutMeter)
}

// expire is the generic check that move expired tasks from a pending pool back
// into a task pool, returning all entities caught with expired tasks.
//
// Note, this method expects the queue lock to be already held. The
// reason the lock is not obtained in here is because the parameters already need
// to access the queue, so they already need a lock anyway.
func (q *queue) expire(timeout time.Duration, pendPool map[string]*fetchRequest, taskQueue *prque.Prque, timeoutMeter metrics.Meter) map[string]int {
	log.Info("[Plasma] queue.expire timeout", "timeout", timeout.String())
	// Iterate over the expired requests and return each to the queue
	expiries := make(map[string]int)
	for id, request := range pendPool {
		if time.Since(request.Time) > timeout {
			// Update the metrics with the timeout
			timeoutMeter.Mark(1)

			// Return any non satisfied requests to the pool
			if len(request.BlkNums) > 0 {
				for _, blkNum := range request.BlkNums {
					log.Info("[Plasma] queue.expire() pushing", "blkNum", blkNum)
					taskQueue.Push(blkNum, -float32(blkNum))
				}
			}

			// Add the peer to the expiry report along the the number of failed requests
			expiries[id] = len(request.BlkNums)
		}
	}
	// Remove the expired requests from the pending pool
	for id := range expiries {
		delete(pendPool, id)
	}
	return expiries
}

func (q *queue) deliverBlocks(id string, blocks types.Blocks) (int, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	// Assemble each of the results with their headers and retrieved data parts
	var (
		accepted int
		sending  types.Blocks
	)

	delete(q.pending, id)

	for _, block := range blocks {
		blkNum := block.NumberU64()

		q.done[blkNum] = true

		accepted++
		sending = append(sending, block)

		// Clean up a successful fetch
		delete(q.pool, blkNum)
	}

	q.resultCache = append(q.resultCache, sending...)

	if accepted > 0 {
		log.Info("[Plasma] awake q.active as blocks delivered")
		q.active.Signal()
	}

	return accepted, nil
}

// resultSlots calculates the number of results slots available for requests
// whilst adhering to both the item and the memory limit too of the results
// cache.
func (q *queue) resultSlots(pendPool map[string]*fetchRequest, donePool map[uint64]bool) int {
	return blockCacheItems - len(pendPool) - len(donePool)

	// Calculate the maximum length capped by the memory limit
	// limit := len(q.resultCache)
	// if common.StorageSize(len(q.resultCache))*q.resultSize > common.StorageSize(blockCacheMemory) {
	// 	limit = int((common.StorageSize(blockCacheMemory) + q.resultSize - 1) / q.resultSize)
	// }
	//
	// // Calculate the number of slots already finished
	// finished := 0
	// for _, result := range q.resultCache[:limit] {
	// 	if result == nil {
	// 		break
	// 	}
	//
	// 	if _, ok := donePool[result.NumberU64()]; ok {
	// 		finished++
	// 	}
	// }
	//
	// // Calculate the number of slots currently downloading
	// pending := 0
	// for _, request := range pendPool {
	// 	pending += len(request.BlkNums)
	// }
	// // Return the free slots to distribute
	// return limit - finished - pending
}

// Results retrieves and permanently removes a batch of fetch results from
// the cache. the result slice will be empty if the queue has been closed.
func (q *queue) Results(block bool) types.Blocks {
	q.lock.Lock()
	log.Info("[Plasma.queue] lock in Results()")
	defer q.lock.Unlock()

	// Count the number of items available for processing
	nproc := q.countProcessableItems()
	for nproc == 0 && !q.closed {
		if !block {
			return nil
		}

		log.Info("[Plasma] Wait until blocks are delivered")

		// TODO: return nil when sync is finished
		q.active.Wait()
		nproc = q.countProcessableItems()
		log.Info("[Plasma] delivered blocks", "nproc", nproc)
	}

	if nproc > maxResultsProcess {
		nproc = maxResultsProcess
	}

	results := make([]*types.Block, nproc)
	copy(results, q.resultCache[:nproc])

	if len(results) > 0 {
		for _, blk := range results {
			delete(q.done, blk.NumberU64())
		}

		// Delete the results from the cache and clear the tail.
		q.resultCache = q.resultCache[nproc:]
	}

	return results
}

// countProcessableItems counts the processable items.
func (q *queue) countProcessableItems() int {
	for i, result := range q.resultCache {
		if result == nil {
			return i
		}
	}

	return len(q.resultCache)
}

package downloader

import (
	"errors"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/log"
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
	pool    map[uint64]bool
	pq      *prque.Prque // [pls/1] Priority queue of the block numbers to fetch
	pending map[string]*fetchRequest
	done    map[uint64]bool

	resultCache types.Blocks // Downloaded but not yet delivered fetch results
	// resultSize   common.StorageSize // Approximate size of a block (exponential moving average)

	lock   *sync.Mutex
	active *sync.Cond
	closed bool
}

func newQueue() *queue {
	lock := new(sync.Mutex)

	return &queue{
		pool:        make(map[uint64]bool),
		pq:          prque.New(),
		pending:     make(map[string]*fetchRequest),
		done:        make(map[uint64]bool),
		resultCache: make(types.Blocks, 0),
		active:      sync.NewCond(lock),
		lock:        lock,
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
	log.Info("[Plasma] queue is closed")
	q.closed = true
	q.lock.Unlock()
	q.active.Broadcast()
}

func (q *queue) enqueue(blkNums []uint64) ([]uint64, error) {
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

func (q *queue) ReserveBlocks(p *peerConnection) *fetchRequest {
	q.lock.Lock()
	defer q.lock.Unlock()

	// Short circuit if the peer's already downloading something (sanity check to
	// not corrupt state)
	if _, ok := q.pending[p.id]; ok {
		return nil
	}

	blkNums := []uint64{}
	for !q.pq.Empty() {
		blkNum := q.pq.PopItem().(uint64)
		blkNums = append(blkNums, blkNum)
	}

	request := &fetchRequest{
		Peer:    p,
		BlkNums: blkNums,
		Time:    time.Now(),
	}
	q.pending[p.id] = request
	return request

}

func (q *queue) deliverBlocks(id string, blocks types.Blocks) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	// Assemble each of the results with their headers and retrieved data parts
	var (
		accepted int
		sending  types.Blocks
	)

	for _, block := range blocks {
		blkNum := block.NumberU64()

		delete(q.pending, id)
		q.done[blkNum] = true

		accepted++
		sending = append(sending, block)
	}

	q.resultCache = append(q.resultCache, sending...)

	if accepted > 0 {
		q.active.Signal()
	}

	return nil
}

// Results retrieves and permanently removes a batch of fetch results from
// the cache. the result slice will be empty if the queue has been closed.
func (q *queue) Results(block bool) types.Blocks {
	q.lock.Lock()
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
	// TODO: activate?
	for i, result := range q.resultCache {
		if result == nil {
			return i
		}
	}

	return len(q.resultCache)
}

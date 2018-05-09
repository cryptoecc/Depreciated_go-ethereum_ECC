package plasma

import (
	"errors"
	"time"

	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/plasma/types"
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

// fetchRequest is a currently running data retrieval operation.
type fetchRequest struct {
	Peer *Peer     // Peer to which the request was sent
	From uint64    // Plasma block number to fetch
	Time time.Time // Time when the request was made
}

// fetchResult is a struct collecting partial results from data fetchers until
// all outstanding pieces complete and the result as a whole can be processed.
type fetchResult struct {
	Pending      int    // Number of data fetches still pending
	Number       uint64 // Plasma block number to be fetched
	Transactions types.Transaction
}

type Queue struct {
	pls *Plasma
}

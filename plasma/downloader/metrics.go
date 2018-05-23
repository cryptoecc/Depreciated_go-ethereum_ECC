package downloader

import (
	"github.com/ethereum/go-ethereum/metrics"
)

var (
	blockInMeter      = metrics.NewRegisteredMeter("pls/downloader/blocks/in", nil)
	blockReqTimer     = metrics.NewRegisteredTimer("pls/downloader/blocks/req", nil)
	blockDropMeter    = metrics.NewRegisteredMeter("pls/downloader/blocks/drop", nil)
	blockTimeoutMeter = metrics.NewRegisteredMeter("pls/downloader/blocks/timeout", nil)
)

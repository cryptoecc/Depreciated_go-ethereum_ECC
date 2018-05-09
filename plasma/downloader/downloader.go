package downloader

import (
	"github.com/ethereum/go-ethereum/plasma"
)

type Downloader struct {
	blockchain *plasma.BlockChain
}

package downloader

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/plasma/types"
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
	AddBlock(b *types.Block) error
}

type Downloader struct {
	blockchain *BlockChain
}

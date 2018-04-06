package plasma

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

var (
	// Block Error
	invalidOperator = errors.New("sender is not operator")

	// Transaction error
	invalidSenderSignature       = errors.New("sender signature is invalid")
	spentTransactionOutput       = errors.New("transaction output is already spent")
	mismatchedTransactionAmounts = errors.New("sum of transaction intputs and outputs are not matched")
)

// BlockChain implements Plasma block chain service
type BlockChain struct {
	config              *Config
	blocks              []*Block
	currentBlock        *Block   // block not mined yet
	currentBlockNumber  *big.Int // block number of currentBlock
	pendingTransactions []*Transaction

	// Channels
	newBlock chan *Block
	quit     chan struct{}

	lock sync.RWMutex
}

// NewBlockChain creates BlockChain instance
func NewBlockChain(config *Config) *BlockChain {
	return &BlockChain{
		config:              config,
		blocks:              []*Block{nil}, // no block with block number 0
		currentBlock:        &Block{},
		currentBlockNumber:  big.NewInt(1),
		pendingTransactions: []*Transaction{},
		newBlock:            make(chan *Block, 1),
		quit:                make(chan struct{}),
	}
}

func (bc *BlockChain) getCurrentBlock() *Block {
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return bc.currentBlock
}

func (bc *BlockChain) getCurrentBlockNumber() *big.Int {
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return bc.currentBlockNumber
}

func (bc *BlockChain) getBlock(blkNum *big.Int) (*Block, error) {
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	if blkNum.Cmp(big.NewInt(0)) == 0 {
		return nil, errors.New("No block with block number 0")
	}

	if blkNum.Cmp(big.NewInt(int64(len(bc.blocks)))) >= 0 {
		return nil, errors.New("No block with block number " + blkNum.String())
	}

	return bc.blocks[blkNum.Int64()], nil
}

func (bc *BlockChain) getTransaction(blkNum, txIndex *big.Int) (*Transaction, error) {
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	if blkNum.Cmp(big.NewInt(0)) == 0 {
		return nil, errors.New("No block with block number 0")
	}

	if blkNum.Cmp(big.NewInt(int64(len(bc.blocks)))) >= 0 {
		return nil, errors.New("No block with block number " + blkNum.String())
	}

	if txIndex.Cmp(big.NewInt(int64(len(bc.blocks[blkNum.Int64()].transactionSet)))) > 0 {
		return nil, errors.New("No transaction with txindex" + txIndex.String())
	}

	return bc.blocks[blkNum.Int64()].transactionSet[txIndex.Int64()], nil
}

// TODO: broadcast new transaction to peers
func (bc *BlockChain) applyTransaction(tx *Transaction) error {
	if err := bc.verifyTransaction(tx); err != nil {
		log.Info("[Plasma Chain] Failed to verify transaction", "hash", tx.Hash(), "error", err)

		return err
	}

	bc.markUtxoSpent(tx.data.blkNum1, tx.data.txIndex1, tx.data.oIndex1)
	bc.markUtxoSpent(tx.data.blkNum2, tx.data.txIndex2, tx.data.oIndex2)

	bc.currentBlock.transactionSet = append(bc.currentBlock.transactionSet, tx)
	return nil
}

func (bc *BlockChain) verifyTransaction(tx *Transaction) error {
	outputAmounts := big.NewInt(0).Add(tx.data.amount1, tx.data.amount2)
	outputAmounts = big.NewInt(0).Add(outputAmounts, tx.data.fee)

	inputAmounts := big.NewInt(0)

	if tx.data.blkNum1.Cmp(big.NewInt(0)) > 0 {
		preTX, _ := bc.getTransaction(tx.data.blkNum1, tx.data.txIndex1)

		if err := verifyTxInput(tx, preTX, tx.data.blkNum1, tx.data.txIndex1, tx.data.oIndex1); err != nil {
			return err
		}

		inputAmount := preTX.data.amount1
		inputAmounts = big.NewInt(0).Add(inputAmounts, inputAmount)
	}

	if tx.data.blkNum2.Cmp(big.NewInt(0)) > 0 {
		preTX, _ := bc.getTransaction(tx.data.blkNum2, tx.data.txIndex2)

		if err := verifyTxInput(tx, preTX, tx.data.blkNum2, tx.data.txIndex2, tx.data.oIndex2); err != nil {
			return err
		}

		inputAmount := preTX.data.amount2
		inputAmounts = big.NewInt(0).Add(inputAmounts, inputAmount)
	}

	if inputAmounts.Cmp(outputAmounts) != 0 {
		return mismatchedTransactionAmounts
	}

	return nil
}

// verify UTXO can be spent
func verifyTxInput(tx, preTx *Transaction, blkNum, txIndex, oIndex *big.Int) error {
	sender, err := tx.Sender(oIndex)

	if err != nil {
		return err
	}

	var spent bool
	var owner *common.Address

	if oIndex.Cmp(big.NewInt(0)) == 0 {
		spent = tx.spent1
		owner = tx.data.newOwner1
	} else if oIndex.Cmp(big.NewInt(1)) == 0 {
		spent = tx.spent2
		owner = tx.data.newOwner2
	}

	if spent {
		return spentTransactionOutput
	}

	if sender != *owner {
		return invalidSenderSignature
	}

	return nil

}

func (bc *BlockChain) markUtxoSpent(blkNum, txIndex, oIndex *big.Int) {
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	if blkNum.Cmp(big.NewInt(0)) == 0 {
		return
	}

	if oIndex.Cmp(big.NewInt(0)) == 0 {
		bc.blocks[blkNum.Int64()].transactionSet[txIndex.Int64()].spent1 = true
	} else {
		bc.blocks[blkNum.Int64()].transactionSet[txIndex.Int64()].spent2 = true
	}
}

// submitBlock seals current block. Only operator can seal, broadcast to peers,
// and record it on root chain
// TODO: Check the submited block is correctly recorded
func (bc *BlockChain) submitBlock(privKey *ecdsa.PrivateKey) (common.Hash, error) {
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	b := bc.currentBlock

	if privKey == nil {
		privKey = bc.config.OperatorPrivateKey
	}

	_, err := b.Seal()
	if err != nil {
		return common.BytesToHash(nil), err
	}

	b.Sign(privKey)

	if sender, err := b.Sender(); err != nil {
		return common.BytesToHash(nil), err
	} else {
		if sender != bc.config.OperatorAddress {
			log.Warn("[Plasma chain] block sealer and plasma operator not matched", "sealer", sender, "operator", bc.config.OperatorAddress)
			return common.BytesToHash(nil), invalidOperator
		}
	}

	bc.currentBlock.blockNumber = big.NewInt(bc.currentBlockNumber.Int64())
	bc.blocks = append(bc.blocks, bc.currentBlock)
	bc.currentBlockNumber = big.NewInt(0).Add(bc.currentBlockNumber, big.NewInt(1))
	bc.currentBlock = &Block{}
	bc.newBlock <- b

	return b.Hash(), nil
}

func (bc *BlockChain) newDeposit(amount *big.Int, depositor *common.Address) (common.Hash, error) {
	bc.lock.Lock()
	defer bc.lock.Unlock()

	tx := NewTransaction(
		big0, big0, big0,
		big0, big0, big0,
		depositor, amount,
		&nullAddress, big0,
		big0)
	transactionSet := []*Transaction{tx}

	b := &Block{
		transactionSet: transactionSet,
		blockNumber:    big.NewInt(bc.currentBlockNumber.Int64()),
	}
	blkNum := *bc.currentBlockNumber

	_, err := b.Seal()
	if err != nil {
		return common.BytesToHash(nil), err
	}

	b.Sign(bc.config.OperatorPrivateKey)

	if sender, err := b.Sender(); err != nil {
		return common.BytesToHash(nil), err
	} else {
		if sender != bc.config.OperatorAddress {
			log.Warn("[Plasma chain] block sealer and plasma operator not matched", "sealer", sender, "operator", bc.config.OperatorAddress)
			return common.BytesToHash(nil), invalidOperator
		}
	}

	bc.blocks = append(bc.blocks, b)
	bc.currentBlockNumber = big.NewInt(0).Add(bc.currentBlockNumber, big.NewInt(1))
	bc.newBlock <- b

	log.Info("[Plasma Chain] New Deposit added", "blockNumber", blkNum.Uint64())

	return b.Hash(), nil
}

// TODO: use event.Feed if needed.
func (bc *BlockChain) addNewBlockListener(f func(blk *Block) error) error {
	for {
		select {
		case blk := <-bc.newBlock:
			if err := f(blk); err != nil {
				log.Info("[Plasma Chain] Faield to listen new block", err)
			}
		case <-bc.quit:
			return nil
		}
	}
}

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

	// Block / TX from peers
	invalidBlockNumber = errors.New("new block has invalid block number. plasma chain may not be synced")
)

// BlockChain implements Plasma block chain service
type BlockChain struct {
	config *Config
	blocks map[uint64]*Block

	// TODO: store to DB
	currentBlock        *Block   // block not mined yet
	currentBlockNumber  *big.Int // block number of currentBlock
	blockInterval       *big.Int // block submitted by operator
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
		blocks:              make(map[uint64]*Block),
		currentBlock:        &Block{},
		currentBlockNumber:  big.NewInt(1000),
		blockInterval:       big.NewInt(1000),
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

	b, ok := bc.blocks[blkNum.Uint64()]

	if !ok {
		return nil, errors.New("No block with block number " + blkNum.String())
	}

	return b, nil
}

func (bc *BlockChain) getTransaction(blkNum, txIndex *big.Int) (*Transaction, error) {
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	if blkNum.Cmp(big.NewInt(0)) == 0 {
		return nil, errors.New("No block with block number 0")
	}

	b, ok := bc.blocks[blkNum.Uint64()]

	if !ok {
		return nil, errors.New("No block with block number " + blkNum.String())
	}

	if txIndex.Cmp(big.NewInt(int64(len(b.data.TransactionSet)))) > 0 {
		return nil, errors.New("No transaction with tx index " + txIndex.String())
	}

	tx := b.data.TransactionSet[txIndex.Int64()]

	return tx, nil
}

// TODO: broadcast new transaction to peers
func (bc *BlockChain) applyTransaction(tx *Transaction) error {
	if err := bc.verifyTransaction(tx); err != nil {
		log.Info("[Plasma Chain] Failed to verify transaction", "hash", tx.Hash(), "error", err)

		return err
	}

	bc.markUtxoSpent(tx.data.BlkNum1, tx.data.TxIndex1, tx.data.OIndex1)
	bc.markUtxoSpent(tx.data.BlkNum2, tx.data.TxIndex2, tx.data.OIndex2)

	bc.currentBlock.data.TransactionSet = append(bc.currentBlock.data.TransactionSet, tx)
	return nil
}

func (bc *BlockChain) verifyTransaction(tx *Transaction) error {
	outputAmounts := big.NewInt(0)

	if tx.data.Amount1 != nil {
		outputAmounts = big.NewInt(0).Add(outputAmounts, tx.data.Amount1)
	}

	if tx.data.Amount2 != nil {
		outputAmounts = big.NewInt(0).Add(outputAmounts, tx.data.Amount2)
	}

	if tx.data.Fee != nil {
		outputAmounts = big.NewInt(0).Add(outputAmounts, tx.data.Fee)

	}

	inputAmounts := big.NewInt(0)

	if tx.data.BlkNum1.Cmp(big.NewInt(0)) > 0 {
		preTX, err := bc.getTransaction(tx.data.BlkNum1, tx.data.TxIndex1)

		if err != nil {
			return err
		}

		if err := verifyTxInput(tx, preTX, big0, tx.data.OIndex1); err != nil {
			return err
		}

		var inputAmount *big.Int
		if tx.data.OIndex1.Cmp(big0) == 0 {
			inputAmount = preTX.data.Amount1
		} else {
			inputAmount = preTX.data.Amount2
		}
		inputAmounts = big.NewInt(0).Add(inputAmounts, inputAmount)
	}

	if tx.data.BlkNum2.Cmp(big.NewInt(0)) > 0 {
		preTX, err := bc.getTransaction(tx.data.BlkNum2, tx.data.TxIndex2)

		if err != nil {
			return err
		}

		if err := verifyTxInput(tx, preTX, big1, tx.data.OIndex2); err != nil {
			return err
		}

		var inputAmount *big.Int
		if tx.data.OIndex1.Cmp(big0) == 0 {
			inputAmount = preTX.data.Amount1
		} else {
			inputAmount = preTX.data.Amount2
		}
		inputAmounts = big.NewInt(0).Add(inputAmounts, inputAmount)
	}

	if inputAmounts.Cmp(outputAmounts) != 0 {
		log.Info("[Plasma chain] mismatched amount", "input", inputAmounts, "output", outputAmounts)
		return mismatchedTransactionAmounts
	}

	return nil
}

// verify UTXO can be spent
func verifyTxInput(tx, preTx *Transaction, curOIndex, preOIndex *big.Int) error {
	sender, err := tx.Sender(curOIndex)

	if err != nil {
		return err
	}

	var spent bool
	var utxoOwner *common.Address

	if preOIndex.Cmp(big.NewInt(0)) == 0 {
		spent = preTx.spent1
		utxoOwner = preTx.data.NewOwner1
	} else if preOIndex.Cmp(big.NewInt(1)) == 0 {
		spent = preTx.spent2
		utxoOwner = preTx.data.NewOwner2
	}

	if spent {
		return spentTransactionOutput
	}

	if sender != *utxoOwner {
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
		bc.blocks[blkNum.Uint64()].data.TransactionSet[txIndex.Int64()].spent1 = true
	} else {
		bc.blocks[blkNum.Uint64()].data.TransactionSet[txIndex.Int64()].spent2 = true
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

	bc.currentBlock.data.BlockNumber = big.NewInt(bc.currentBlockNumber.Int64())
	bc.blocks[bc.currentBlockNumber.Uint64()] = bc.currentBlock
	bc.currentBlockNumber = big.NewInt(0).Add(bc.currentBlockNumber, bc.blockInterval)
	bc.currentBlock = &Block{}
	bc.newBlock <- b

	return b.Hash(), nil
}

// only operator can add deposit transaction
func (bc *BlockChain) newDeposit(amount *big.Int, depositor *common.Address, depositBlock *big.Int) (common.Hash, error) {
	bc.lock.Lock()
	defer bc.lock.Unlock()

	tx := NewTransaction(
		depositBlock, big0, big0,
		big0, big0, big0,
		depositor, amount,
		&nullAddress, big0,
		big0)

	transactionSet := []*Transaction{tx}

	b := NewBlock(depositBlock, transactionSet, nil)

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

	bc.blocks[depositBlock.Uint64()] = b
	bc.newBlock <- b

	log.Info("[Plasma Chain] New Deposit added", "depositBlock", depositBlock)

	return b.Hash(), nil
}

// TODO: use event.Feed if needed.
func (bc *BlockChain) addNewBlockListener(f func(blk *Block) error) error {
	for {
		select {
		case blk := <-bc.newBlock:
			if err := f(blk); err != nil {
				log.Info("[Plasma Chain] Faield to listen new block", "err", err)
			}
		case <-bc.quit:
			return nil
		}
	}
}

// add deposit block or synced block
func (bc *BlockChain) addBlock(b *Block) error {
	if bc.currentBlockNumber.Cmp(b.data.BlockNumber) != 0 {
		return invalidBlockNumber
	}

	bc.blocks[b.data.BlockNumber.Uint64()] = b
	bc.currentBlockNumber = big0.And(bc.currentBlockNumber, big1)

	// channel needed?
	// bc.newBlock <- b

	return nil
}

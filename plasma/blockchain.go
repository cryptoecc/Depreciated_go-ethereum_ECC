package plasma

import (
  "math/big"
  "sync"

  "github.com/ethereum/go-ethereum/common"
)

const (
  // Block Error
  invalidOperator = iota

  // Transaction error
  invalidSenderSignature
  spentTransactionOutput
  mismatchedTransactionAmounts
)

// BlockChain implements Plasma block chain service
type BlockChain struct {
  config              BlockChainConfig
  blocks              []*Block
  currentBlock        *Block
  currentBlockNumber  *big.Int
  pendingTransactions []*Transaction

  lock sync.RWMutex
}

// NewBlockChain creates BlockChain instance
func NewBlockChain(config BlockChainConfig) *BlockChain {
  return &BlockChain{
    config:              config,
    blocks:              make([]*Block),
    currentBlock:        &Block{},
    currentBlockNumber:  big.Int.NewInt(1),
    pendingTransactions: make([]*Transaction),
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

func (bc *BlockChain) getBlock(blkNum *big.Int) *Block {
  bc.lock.RLock()
  defer bc.lock.RUnlock()
  return bc.blocks[blkNum]
}

func (bc *BlockChain) getTransaction(blkNum, txIndex *big.Int) *Transaction {
  bc.lock.RLock()
  defer bc.lock.RUnlock()
  return bc.blocks[blkNum].transactionSet[txIndex]
}

func (bc *BlockChain) applyTransaction(tx *Transaction) error {
  if err := bc.verifyTransaction(tx); err != nil {
    return err
  }

  bc.markUtxoSpent(tx.blkNum1, tx.txIndex1, tx.oIndex1)
  bc.markUtxoSpent(tx.blkNum2, tx.txIndex2, tx.oIndex2)

  bc.currentBlock.transactionSet = append(bc.currentBlock.transactionSet, tx)
}

func (bc *BlockChain) verifyTransaction(tx *Transaction) error {
  outputAmounts := big.Int.NewInt(0).add(tx.amount1, tx.amount2)
  outputAmounts = big.Int.NewInt(0).add(outputAmounts, tx.fee)

  inputAmounts := big.Int.NewInt(0)

  if tx.blkNum1 != 0 {
    preTX := bc.getTransaction(tx.blkNum1, tx.txIndex1)

    if err := verifyTxInput(tx, pre, tx.blkNum1, tx.txIndex1, tx.oIndex1); err != nil {
      return err
    }

    inputAmount := preTX.amount1
    inputAmounts = big.Int.NewInt(0).add(inputAmounts, inputAmount)
  }

  if tx.blkNum2 != 0 {
    preTX := bc.getTransaction(tx.blkNum2, tx.txIndex2)

    if err := verifyTxInput(tx, pre, tx.blkNum2, tx.txIndex2, tx.oIndex2); err != nil {
      return err
    }

    inputAmount := preTX.amount2
    inputAmounts = big.Int.NewInt(0).add(inputAmounts, inputAmount)
  }

  if inputAmount != outputAmount {
    return mismatchedTransactionAmounts
  }

  return nil
}

// verify UTXO of preTX can be spent
func verifyTxInput(tx, preTx *Transaction, blkNum, txIndex, oIndex *big.Int) error {
  sender := tx.Sender(oIndex)

  var spent bool
  var owner common.Address

  if oIndex == 0 {
    spent = tx.spent1
    owner = tx.newOwner1
  } else if oIndex == 1 {
    spent = tx.spent2
    owner = tx.newOwner2
  }

  if spent {
    return spentTransactionOutput
  }

  if sender != owner {
    return invalidSenderSignature
  }

  return nil

}

func (bc *BlockChain) markUtxoSpent(blkNum, txIndex, oIndex *big.Int) {
  bc.lock.RLock()
  defer bc.lock.RUnlock()

  if blkNum == 0 {
    return
  }

  if oIndex == 0 {
    bc.blocks[blkNum].transactionSet[txIndex].spent1 = true
  } else {
    bc.blocks[blkNum].transactionSet[txIndex].spent2 = true
  }
}

func (bc *BlockChain) submitBlock(b *Block) {
  bc.lock.RLock()
  defer bc.lock.RUnlock()

  if bc.config.operatorAddress != b.Sender() {
    return invalidOperator
  }

  bc.blocks[bc.currentBlockNumber] = bc.currentBlock
  bc.currentBlockNumber = big.Int.NewInt(0).add(bc.currentBlockNumber, 1)
  bc.currentBlock = &Block{}
}

// read transaction with hash of `txHash` from root chain
func (bc *BlockChain) submitDeposit(txHash common.Hash) {
  bc.lock.RLock()
  defer bc.lock.RUnlock()

  tx := txHash
}

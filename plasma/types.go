package plasma

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/rlp"
)

// Block implements Plasma chiain block
type Block struct {
	transactionSet []*Transaction
	sig            []byte
}

func NewBlock() *Block {
	return &Block{}
}

// Hash returns sha3 hash of Block
func (b *Block) Hash() []byte {
	txHashes := make([][]byte, len(b.transactionSet))

	for _, tx := range b.transactionSet {
		txHashes = append(txHashes, tx.Hash())
	}

	h := sha3.NewKeccak256()
	rlp.Encode(h, txHashes)
	return h.Sum(nil)
}

func (b *Block) Sign(privKey *ecdsa.PrivateKey) error {
	sig, err := crypto.Sign(b.Hash(), privKey)

	if err != nil {
		return err
	}

	b.sig = sig
	return nil
}

// Sender returns address of block minder
func (b *Block) Sender() (common.Address, error) {
	return getSender(b.Hash(), b.sig)
}

// unsignTransaction only contains requierd fields for hash function
// TODO: change big.Int to uint64 to reduce txData size
type txData struct {
	blkNum1   *big.Int
	txIndex1  *big.Int
	oIndex1   *big.Int
	blkNum2   *big.Int
	txIndex2  *big.Int
	oIndex2   *big.Int
	newOwner1 *common.Address
	amount1   *big.Int
	newOwner2 *common.Address
	amount2   *big.Int
	fee       *big.Int
}

// Transaction implements Plasma chain transaction
type Transaction struct {
	data txData
	sig1 []byte
	sig2 []byte

	// whether TX output is spent or not
	spent1 bool
	spent2 bool
}

// NewTransaction creates Transaction instance
func NewTransaction(blkNum1, txIndex1, oIndex1, blkNum2, txIndex2, oIndex2 *big.Int, newOwner1 *common.Address, amount1 *big.Int, newOwner2 *common.Address, amount2, fee *big.Int) *Transaction {
	data := txData{
		blkNum1, txIndex1, oIndex1,
		blkNum2, txIndex2, oIndex2,
		newOwner1, amount1,
		newOwner2, amount2,
		fee,
	}

	return &Transaction{data, nil, nil, false, false}
}

// Hash returns sha3 hash of Transaction
func (tx *Transaction) Hash() []byte {
	h := sha3.NewKeccak256()
	rlp.Encode(h, tx.data)
	return h.Sum(nil)
}

// Sender returns owner address of TX input
func (tx *Transaction) Sender(oIndex *big.Int) (common.Address, error) {
	hash := tx.Hash()
	var sig []byte

	if oIndex.Cmp(big.NewInt(0)) == 0 {
		sig = tx.sig1
	} else if oIndex.Cmp(big.NewInt(1)) == 0 {
		sig = tx.sig2
	}

	return getSender(hash, sig)
}

func (tx *Transaction) Sign1(privKey *ecdsa.PrivateKey) error {
	sig, err := crypto.Sign(tx.Hash(), privKey)

	if err != nil {
		return err
	}

	tx.sig1 = sig
	return nil
}

func (tx *Transaction) Sign2(privKey *ecdsa.PrivateKey) error {
	sig, err := crypto.Sign(tx.Hash(), privKey)

	if err != nil {
		return err
	}

	tx.sig2 = sig
	return nil
}

func getSender(hash, sig []byte) (common.Address, error) {
	pubKeyBytes, err := crypto.Ecrecover(hash, sig)

	if err != nil {
		return common.Address{}, err
	}

	pubKey := crypto.ToECDSAPub(pubKeyBytes)
	return crypto.PubkeyToAddress(*pubKey), nil
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

func sign(hash []byte, privKey *ecdsa.PrivateKey) (sig []byte, err error) {
	sig, err = crypto.Sign(hash, privKey)
	return
}

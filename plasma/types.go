package plasma

import (
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

// Hash returns sha3 hash of Block
// TODO: implement this
func (b *Block) Hash() []byte {
	return nil
}

// Sender returns address of block minder
func (b *Block) Sender() (common.Address, error) {
	return getSender(b.Hash(), b.sig)
}

// Transaction implements Plasma chain transaction
type Transaction struct {
	// General TX Fields
	blkNum1   *big.Int
	txIndex1  *big.Int
	oIndex1   *big.Int
	blkNum2   *big.Int
	txIndex2  *big.Int
	oIndex2   *big.Int
	newOwner1 common.Address
	amount1   *big.Int
	newOwner2 common.Address
	amount2   *big.Int
	fee       *big.Int
	sig1      []byte
	sig2      []byte

	// whether TX output is spent or not
	spent1 bool
	spent2 bool
}

// Hash returns sha3 hash of Transaction
// TODO: implement this
func (tx *Transaction) Hash() []byte {
	return nil
}

// Sender returns sender of TX input
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

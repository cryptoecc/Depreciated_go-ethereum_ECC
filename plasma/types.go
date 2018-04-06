package plasma

import (
	"crypto/ecdsa"
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/plasma/merkle"
	"github.com/ethereum/go-ethereum/rlp"
)

// Block implements Plasma chiain block
type Block struct {
	blockNumber    *big.Int
	transactionSet []*Transaction
	merkle         *merkle.Merkle // TODO: store in DB with caching
	sig            []byte
}

type blockJSONData struct {
	hash         common.Hash `json:"hash"`
	transactions [][]byte    `json:"transactions"`
}

func (b *Block) ToRPCResponse() map[string]interface{} {
	var transactions []map[string]interface{}

	for _, tx := range b.transactionSet {
		transactions = append(transactions, tx.ToRPCResponse())
	}

	return map[string]interface{}{
		"hash":         b.merkle.Root(),
		"blockNumber":  b.blockNumber,
		"transactions": transactions,
	}
}

func NewBlock() *Block {
	return &Block{}
}

func (b *Block) Seal() (common.Hash, error) {
	var hashes []common.Hash

	for _, tx := range b.transactionSet {
		hashes = append(hashes, tx.Hash())
	}

	merkle, err := merkle.NewMerkle(16, hashes)

	b.merkle = merkle

	if err != nil {
		return common.HexToHash(""), err
	}

	return merkle.Root(), nil
}

// Hash returns sha3 hash of Block
func (b *Block) Hash() common.Hash {
	return b.merkle.Root()
}

func (b *Block) MarshalJSON() ([]byte, error) {
	var enc blockJSONData

	enc.hash = b.Hash()

	for _, tx := range b.transactionSet {
		txJSON, err := tx.MarshalJSON()

		if err != nil {
			return nil, err
		}

		enc.transactions = append(enc.transactions, txJSON)
	}

	return json.Marshal(&enc)
}

func (b *Block) Sign(privKey *ecdsa.PrivateKey) error {
	sig, err := crypto.Sign(b.Hash().Bytes(), privKey)

	if err != nil {
		return err
	}

	b.sig = sig
	return nil
}

// Sender returns address of block minder
func (b *Block) Sender() (common.Address, error) {
	return getSender(b.Hash().Bytes(), b.sig)
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

type txJSONData struct {
	blkNum1   *hexutil.Big    `json:"blkNum1"`
	txIndex1  *hexutil.Big    `json:"txIndex1"`
	oIndex1   *hexutil.Big    `json:"oIndex1"`
	blkNum2   *hexutil.Big    `json:"blkNum2"`
	txIndex2  *hexutil.Big    `json:"txIndex2"`
	oIndex2   *hexutil.Big    `json:"oIndex2"`
	newOwner1 *common.Address `json:"newOwner1"`
	amount1   *hexutil.Big    `json:"amount1"`
	newOwner2 *common.Address `json:"newOwner2"`
	amount2   *hexutil.Big    `json:"amount2"`
	fee       *hexutil.Big    `json:"fee"`
	sig1      hexutil.Bytes   `json:"sig1"`
	sig2      hexutil.Bytes   `json:"sig2"`
	spent1    *hexutil.Big    `json:"spent1"`
	spent2    *hexutil.Big    `json:"spent2"`
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
func (tx *Transaction) Hash() (h common.Hash) {
	d := sha3.NewKeccak256()
	rlp.Encode(d, tx.data)
	d.Sum(h[:0])

	return h
}

func (tx *Transaction) ToRPCResponse() map[string]interface{} {

	ret := map[string]interface{}{
		"blkNum1":   tx.data.blkNum1,
		"txIndex1":  tx.data.txIndex1,
		"oIndex1":   tx.data.oIndex1,
		"blkNum2":   tx.data.blkNum2,
		"txIndex2":  tx.data.txIndex2,
		"oIndex2":   tx.data.oIndex2,
		"newOwner1": tx.data.newOwner1,
		"amount1":   tx.data.amount1,
		"newOwner2": tx.data.newOwner2,
		"amount2":   tx.data.amount2,
		"fee":       tx.data.fee,
	}

	if len(tx.sig1) > 0 {
		ret["v1"] = tx.sig1[64]
		ret["r1"] = tx.sig1[0:32]
		ret["s1"] = tx.sig1[32:64]
	}

	if len(tx.sig2) > 0 {
		ret["v1"] = tx.sig2[64]
		ret["r1"] = tx.sig2[0:32]
		ret["s1"] = tx.sig2[32:64]
	}

	return ret
}

func (tx *Transaction) MarshalJSON() ([]byte, error) {

	var enc txJSONData

	enc.blkNum1 = (*hexutil.Big)(tx.data.blkNum1)
	enc.txIndex1 = (*hexutil.Big)(tx.data.txIndex1)
	enc.oIndex1 = (*hexutil.Big)(tx.data.oIndex1)
	enc.blkNum2 = (*hexutil.Big)(tx.data.blkNum2)
	enc.txIndex2 = (*hexutil.Big)(tx.data.txIndex2)
	enc.oIndex2 = (*hexutil.Big)(tx.data.oIndex2)
	enc.newOwner1 = tx.data.newOwner1
	enc.amount1 = (*hexutil.Big)(tx.data.amount1)
	enc.newOwner2 = tx.data.newOwner2
	enc.amount2 = (*hexutil.Big)(tx.data.amount2)
	enc.fee = (*hexutil.Big)(tx.data.fee)

	enc.sig1 = tx.sig1
	enc.sig2 = tx.sig2

	if tx.spent1 {
		enc.spent1 = (*hexutil.Big)(big1)
	} else {
		enc.spent1 = (*hexutil.Big)(big0)
	}
	if tx.spent2 {
		enc.spent2 = (*hexutil.Big)(big1)
	} else {
		enc.spent2 = (*hexutil.Big)(big0)
	}

	return json.Marshal(&enc)
}

// Sender returns owner address of TX input
func (tx *Transaction) Sender(oIndex *big.Int) (common.Address, error) {
	var sig []byte

	if oIndex.Cmp(big.NewInt(0)) == 0 {
		sig = tx.sig1
	} else if oIndex.Cmp(big.NewInt(1)) == 0 {
		sig = tx.sig2
	}

	return getSender(tx.Hash().Bytes(), sig)
}

func (tx *Transaction) Sign1(privKey *ecdsa.PrivateKey) error {
	sig, err := crypto.Sign(tx.Hash().Bytes(), privKey)

	if err != nil {
		return err
	}

	tx.sig1 = sig
	return nil
}

func (tx *Transaction) Sign2(privKey *ecdsa.PrivateKey) error {
	sig, err := crypto.Sign(tx.Hash().Bytes(), privKey)

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

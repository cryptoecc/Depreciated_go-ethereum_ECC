package types

import (
	"crypto/ecdsa"
	"encoding/json"
	"io"
	"math/big"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/plasma/merkle"
	"github.com/ethereum/go-ethereum/rlp"
)

type blockData struct {
	BlockNumber    *big.Int       `json:"blockNumber"`
	TransactionSet []*Transaction `json:"transactionSet"`
	Sig            []byte         `json:"sig"`
}

// Block implements Plasma chiain block
type Block struct {
	Data   blockData
	Merkle *merkle.Merkle `json:"merkle" rlp:"nil"` // TODO: store in DB with caching
	hash   atomic.Value
	size   atomic.Value
}

type Blocks []*Block

type blockJSONData struct {
	hash         common.Hash `json:"hash"`
	transactions [][]byte    `json:"transactions"`
}

func (b *Block) NumberU64() uint64 {
	return b.Data.BlockNumber.Uint64()
}

func (b *Block) ToRPCResponse() map[string]interface{} {
	var transactions []map[string]interface{}

	for _, tx := range b.Data.TransactionSet {
		transactions = append(transactions, tx.ToRPCResponse())
	}

	return map[string]interface{}{
		"hash":         b.Merkle.Root,
		"blockNumber":  b.Data.BlockNumber,
		"transactions": transactions,
	}
}

func NewBlock(blkNum *big.Int, txSet []*Transaction, sig []byte) *Block {
	data := blockData{
		BlockNumber:    blkNum,
		TransactionSet: txSet,
		Sig:            sig,
	}

	block := Block{
		Data: data,
	}

	return &block
}

func (b *Block) Seal() (common.Hash, error) {
	var hashes []common.Hash

	for _, tx := range b.Data.TransactionSet {
		hashes = append(hashes, tx.Hash())
	}

	merkle, err := merkle.NewMerkle(16, hashes)

	b.Merkle = merkle

	if err != nil {
		return common.HexToHash(""), err
	}

	return merkle.Root, nil
}

// Hash returns sha3 hash of Block
func (b *Block) Hash() common.Hash {
	if hash := b.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	v := b.Merkle.Root
	b.hash.Store(v)
	return v
}

// EncodeRLP implements rlp.Encoder
func (b *Block) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &b.Data)
}

// DecodeRLP implements rlp.Decoder
func (b *Block) DecodeRLP(s *rlp.Stream) error {
	_, size, _ := s.Kind()

	if err := s.Decode(&b.Data); err != nil {
		return err
	}

	b.size.Store(common.StorageSize(rlp.ListSize(size)))

	var hashes []common.Hash

	for _, tx := range b.Data.TransactionSet {
		hashes = append(hashes, tx.Hash())
	}

	merkle, err := merkle.NewMerkle(16, hashes)

	if err != nil {
		return err
	}

	b.Merkle = merkle

	return nil
}

func (b *Block) MarshalJSON() ([]byte, error) {
	log.Info("[Plasma types] Block.MarshalJSON()")
	var enc blockJSONData

	enc.hash = b.Hash()

	for _, tx := range b.Data.TransactionSet {
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

	b.Data.Sig = sig
	return nil
}

// Sender returns address of block minder
func (b *Block) Sender() (common.Address, error) {
	return getSender(b.Hash().Bytes(), b.Data.Sig)
}

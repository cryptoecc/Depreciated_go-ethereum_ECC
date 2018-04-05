package merkle

import (
	// "bytes"
	// "errors"
	// "math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/crypto/sha3"
)

func EmptyMerkleTreeRoot(depth uint16) common.Hash {
	root := EmptyHash

	for i := 0; i < int(depth); i++ {
		root = crypto.Keccak256Hash(root.Bytes(), root.Bytes())
	}

	return root
}

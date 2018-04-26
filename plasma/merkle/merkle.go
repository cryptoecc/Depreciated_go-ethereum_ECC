package merkle

import (
	"bytes"
	"errors"
	"fmt"
	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

var (
	invalidDepth     = errors.New("depth should be at least 1")
	invalidLeafCount = errors.New("num of leaves exceed max avaiable num with the depth")

	invalidMembership  = errors.New("leaf is not in the merkle tree")
	invalidProof       = errors.New("proof is not invalid for the merkle tree")
	invalidProofLength = errors.New("length of proof should be multiple of 32")

	invalidMerkle = errors.New("the number of hashes is more than depth permits")
)

type Node struct {
	Data  common.Hash
	Left  *Node
	Right *Node
}

func (n *Node) Bytes() []byte {
	var b []byte

	if n.Left != nil {
		b = append(b, n.Left.Bytes()...)
	}
	return n.Data.Bytes()
}

func NewNode(data common.Hash, left, right *Node) *Node {
	return &Node{
		Data:  data,
		Left:  left,
		Right: right,
	}
}

type Nodes []*Node

func NewNodes(hashes ...common.Hash) Nodes {
	var nodes Nodes

	for _, hash := range hashes {
		nodes = append(nodes, NewNode(hash, nil, nil))
	}

	return nodes
}

func (n Nodes) Hashes() []common.Hash {
	var hashes []common.Hash

	for _, node := range n {
		hashes = append(hashes, node.Data)
	}

	return hashes
}

func (n Nodes) Bytes() (b []byte) {
	for _, node := range n {
		b = append(b, node.Data.Bytes()...)
	}

	return b
}

type Tree [][]*Node

type Proof []common.Hash

func NewProof(bs ...[]byte) (Proof, error) {
	var proof Proof

	for _, b := range bs {
		if len(b)%32 != 0 {
			return nil, invalidProofLength
		}

		for i := 0; i < len(b); i += 32 {
			hash := common.BytesToHash(b[i : i+32])
			proof = append(proof, hash)
		}
	}

	return proof, nil
}

func (p Proof) Bytes() (b []byte) {
	for _, elem := range p {
		b = append(b, elem.Bytes()...)
	}

	return b
}

type Merkle struct {
	Root   common.Hash
	Depth  uint16
	Leaves Nodes
	Tree   Tree
}

var EmptyHash = common.StringToHash("")
var Empty32Bytes = EmptyHash.Bytes()
var EmptySha3Hash = crypto.Keccak256Hash(EmptyHash.Bytes())
var EmptyNode = NewNode(EmptyHash, nil, nil)

func NewMerkle(depth uint16, hashes []common.Hash) (*Merkle, error) {
	numLeaves := int(math.Pow(2, float64(depth)))
	numEmptyLeaves := numLeaves - len(hashes)

	merkle := Merkle{
		Depth: depth,
	}

	if len(hashes) > numLeaves {
		return nil, invalidMerkle
	}

	for _, hash := range hashes {
		node := NewNode(hash, nil, nil)

		merkle.Leaves = append(merkle.Leaves, node)
	}

	for i := 0; i < numEmptyLeaves; i++ {
		merkle.Leaves = append(merkle.Leaves, EmptyNode)
	}

	merkle.Tree = append(merkle.Tree, merkle.Leaves)
	merkle.Root = merkle.CreateTree(merkle.Leaves)

	return &merkle, nil
}

func (m *Merkle) CreateTree(leaves Nodes) common.Hash {
	if len(leaves) == 1 {
		return leaves[0].Data
	}

	d := sha3.NewKeccak256()

	nextLevel := len(leaves)
	var nextLeaves Nodes

	for i := 0; i < nextLevel; i += 2 {
		left, right := leaves[i], leaves[i+1]

		d.Reset()
		d.Write(left.Data.Bytes())
		d.Write(right.Data.Bytes())

		var combinedHash common.Hash

		d.Sum(combinedHash[:0])
		node := NewNode(combinedHash, left, right)
		nextLeaves = append(nextLeaves, node)
	}

	m.Tree = append(m.Tree, nextLeaves)

	return m.CreateTree(nextLeaves)
}

// CheckMembership prove hash exixts in the merkle tree
func (m *Merkle) CheckMembership(hash common.Hash, proofByte []byte) error {
	index := m.Index(hash)

	if index < 0 {
		return invalidMembership
	}

	proof, err := NewProof(proofByte)

	if err != nil {
		return err
	}

	computedHash := hash.Bytes()

	for i := 0; i < int(m.Depth); i++ {
		segment := proof[i]

		if index%2 == 0 {
			computedHash = crypto.Keccak256(computedHash, segment.Bytes())
		} else {
			computedHash = crypto.Keccak256(segment.Bytes(), computedHash)
		}

		index = index / 2
	}

	if !bytes.Equal(computedHash, m.Root.Bytes()) {
		return invalidProof
	}

	return nil
}

func (m *Merkle) GenerateProof(hash common.Hash) (Proof, error) {
	index := m.Index(hash)

	if index < 0 {
		return nil, invalidMembership
	}

	var proofBytes []byte

	for i := 0; i < int(m.Depth); i++ {
		var siblingIndex int

		if index%2 == 0 {
			siblingIndex = index + 1
		} else {
			siblingIndex = index - 1
		}

		index = index / 2

		proofBytes = append(proofBytes, m.Tree[i][siblingIndex].Data.Bytes()...)
	}

	return NewProof(proofBytes)
}

func (m *Merkle) Index(hash common.Hash) int {
	ret := -1

	for i, node := range m.Leaves {
		if bytes.Equal(node.Data.Bytes(), hash.Bytes()) {
			return i
		}
	}
	fmt.Print()

	return ret
}

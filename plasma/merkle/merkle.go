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
	data  common.Hash
	left  *Node
	right *Node
}

func (n *Node) Bytes() []byte {
	var b []byte

	if n.left != nil {
		b = append(b, n.left.Bytes()...)
	}
	return n.data.Bytes()
}

func NewNode(data common.Hash, left, right *Node) *Node {
	return &Node{
		data:  data,
		left:  left,
		right: right,
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
		hashes = append(hashes, node.data)
	}

	return hashes
}

func (n Nodes) Bytes() (b []byte) {
	for _, node := range n {
		b = append(b, node.data.Bytes()...)
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
	root   common.Hash
	depth  uint16
	leaves Nodes
	tree   Tree
}

var EmptyHash = common.StringToHash("")
var Empty32Bytes = EmptyHash.Bytes()
var EmptySha3Hash = crypto.Keccak256Hash(EmptyHash.Bytes())
var EmptyNode = NewNode(EmptyHash, nil, nil)

func NewMerkle(depth uint16, hashes []common.Hash) (*Merkle, error) {
	numLeaves := int(math.Pow(2, float64(depth)))
	numEmptyLeaves := numLeaves - len(hashes)

	merkle := Merkle{
		depth: depth,
	}

	if len(hashes) > numLeaves {
		return nil, invalidMerkle
	}

	for _, hash := range hashes {
		node := NewNode(hash, nil, nil)

		merkle.leaves = append(merkle.leaves, node)
	}

	for i := 0; i < numEmptyLeaves; i++ {
		merkle.leaves = append(merkle.leaves, EmptyNode)
	}

	merkle.tree = append(merkle.tree, merkle.leaves)
	merkle.root = merkle.CreateTree(merkle.leaves)

	return &merkle, nil
}

func (m *Merkle) CreateTree(leaves Nodes) common.Hash {
	if len(leaves) == 1 {
		return leaves[0].data
	}

	d := sha3.NewKeccak256()

	nextLevel := len(leaves)
	var nextLeaves Nodes

	for i := 0; i < nextLevel; i += 2 {
		left, right := leaves[i], leaves[i+1]

		d.Reset()
		d.Write(left.data.Bytes())
		d.Write(right.data.Bytes())

		var combinedHash common.Hash

		d.Sum(combinedHash[:0])
		node := NewNode(combinedHash, left, right)
		nextLeaves = append(nextLeaves, node)
	}

	m.tree = append(m.tree, nextLeaves)

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

	for i := 0; i < int(m.depth); i++ {
		segment := proof[i]

		if index%2 == 0 {
			computedHash = crypto.Keccak256(computedHash, segment.Bytes())
		} else {
			computedHash = crypto.Keccak256(segment.Bytes(), computedHash)
		}

		index = index / 2
	}

	if !bytes.Equal(computedHash, m.root.Bytes()) {
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

	for i := 0; i < int(m.depth); i++ {
		var siblingIndex int

		if index%2 == 0 {
			siblingIndex = index + 1
		} else {
			siblingIndex = index - 1
		}

		index = index / 2

		proofBytes = append(proofBytes, m.tree[i][siblingIndex].data.Bytes()...)
	}

	return NewProof(proofBytes)
}

func (m *Merkle) Index(hash common.Hash) int {
	ret := -1

	for i, node := range m.leaves {
		if bytes.Equal(node.data.Bytes(), hash.Bytes()) {
			return i
		}
	}
	fmt.Print()

	return ret
}

func (m *Merkle) Root() common.Hash {
	return m.root
}

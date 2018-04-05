package merkle

import (
	"bytes"
	"math"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestInitializeTreeWithoutLeaves(t *testing.T) {
	f := func(depth int) {
		merkle, _ := NewMerkle(uint16(depth), nil)

		leaveBytes := merkle.leaves.Bytes()
		emptyBytes := bytes.Repeat(Empty32Bytes,
			int(math.Pow(2, float64(depth))))

		if !bytes.Equal(emptyBytes, leaveBytes) {
			t.Fatal("emptyBytes and leaveBytes isn't matched")
		}
	}

	for depth := 1; depth <= 12; depth++ {
		f(depth)
	}
}

func TestInitializeTreeWithLeaves(t *testing.T) {
	leaves1 := NewNodes(
		common.HexToHash("a"),
		common.HexToHash("b"),
		common.HexToHash("c"),
	)

	leaves2 := NewNodes(
		common.HexToHash("a"),
		common.HexToHash("b"),
		common.HexToHash("c"),
		common.HexToHash("d"),
		common.HexToHash("e"),
	)

	f := func(depth int, nodes Nodes, numEmptyNodes int) {
		hashes := nodes.Hashes()

		merkle, _ := NewMerkle(uint16(depth), hashes)

		leaveBytes := merkle.leaves.Bytes()
		var targetBytes []byte

		for _, hash := range hashes {
			targetBytes = append(targetBytes, hash.Bytes()...)
		}

		emptyBytes := bytes.Repeat(Empty32Bytes, numEmptyNodes)

		targetBytes = append(targetBytes, emptyBytes...)

		if !bytes.Equal(leaveBytes, targetBytes) {
			t.Log("leaveBytes", leaveBytes)
			t.Log("targetBytes", targetBytes)
			t.Fatal("leaveBytes and targetBytes isn't matched")
		}
	}

	f(2, leaves1, 1)
	f(3, leaves2, 3)
}

func TestInitializeTreeWithLeavesMoreThanDepth(t *testing.T) {
	depth := uint16(1)
	hash := common.HexToHash("dead")
	hashes := []common.Hash{hash, hash, hash}
	_, err := NewMerkle(depth, hashes)

	if err != invalidMerkle {
		t.Fatal("invalid error", err)
	}
}

func TestEmptyRoot(t *testing.T) {
	root1Bytes := append(Empty32Bytes, Empty32Bytes...)
	root1 := crypto.Keccak256Hash(root1Bytes)

	root2Bytes := append(root1.Bytes(), root1.Bytes()...)
	root2 := crypto.Keccak256Hash(root2Bytes)

	root3Bytes := append(root2.Bytes(), root2.Bytes()...)
	root3 := crypto.Keccak256Hash(root3Bytes)

	root3_2 := EmptyMerkleTreeRoot(3)

	root16 := EmptyMerkleTreeRoot(16)

	merkle1, _ := NewMerkle(1, nil)
	merkle2, _ := NewMerkle(2, nil)
	merkle3, _ := NewMerkle(3, nil)
	merkle16, _ := NewMerkle(16, nil)

	if root1.Hex() != merkle1.root.Hex() {
		t.Fatal("root1 has not matched")
	}
	if root2.Hex() != merkle2.root.Hex() {
		t.Fatal("root2 has not matched")
	}
	if root3.Hex() != merkle3.root.Hex() {
		t.Fatal("root3 has not matched")
	}
	if root3.Hex() != root3_2.Hex() {
		t.Fatal("root3_2 has not matched")
	}
	if root16.Hex() != merkle16.root.Hex() {
		t.Fatal("root16 has not matched")
	}
}

func TestCheckMembership(t *testing.T) {
	// depth 1
	leaf1 := common.HexToHash("ff00000000000000000000000000000000000000000000000000000000000001")
	leaf2 := common.HexToHash("ff00000000000000000000000000000000000000000000000000000000000002")
	leaf3 := common.HexToHash("ff00000000000000000000000000000000000000000000000000000000000003")
	leaf4 := common.HexToHash("ff00000000000000000000000000000000000000000000000000000000000004")

	// depth 2
	leaf12Root := crypto.Keccak256Hash(leaf1.Bytes(), leaf2.Bytes())
	leaf34Root := crypto.Keccak256Hash(leaf3.Bytes(), leaf4.Bytes())

	// depth 3
	root := crypto.Keccak256Hash(leaf12Root.Bytes(), leaf34Root.Bytes())

	// depth 16
	zeroHash := EmptyMerkleTreeRoot(2)
	zeroHashesBytes := zeroHash.Bytes()

	for i := 0; i < 13; i++ {
		root = crypto.Keccak256Hash(root.Bytes(), zeroHash.Bytes())
		newZeroHash := crypto.Keccak256Hash(zeroHash.Bytes(), zeroHash.Bytes())
		zeroHash = newZeroHash
		zeroHashesBytes = append(zeroHashesBytes, zeroHash.Bytes()...)
	}

	leaf1Proof, _ := NewProof(leaf2.Bytes(), leaf34Root.Bytes(), zeroHashesBytes)
	leaf2Proof, _ := NewProof(leaf1.Bytes(), leaf34Root.Bytes(), zeroHashesBytes)
	leaf3Proof, _ := NewProof(leaf4.Bytes(), leaf12Root.Bytes(), zeroHashesBytes)
	leaf4Proof, _ := NewProof(leaf3.Bytes(), leaf12Root.Bytes(), zeroHashesBytes)

	merkle, err := NewMerkle(16, []common.Hash{leaf1, leaf2, leaf3, leaf4})

	if err != nil {
		t.Error(err)
	}

	if err := merkle.CheckMembership(leaf1, leaf1Proof.Bytes()); err != nil {
		t.Fatal(err)
	}
	if err := merkle.CheckMembership(leaf2, leaf2Proof.Bytes()); err != nil {
		t.Fatal(err)
	}
	if err := merkle.CheckMembership(leaf3, leaf3Proof.Bytes()); err != nil {
		t.Fatal(err)
	}
	if err := merkle.CheckMembership(leaf4, leaf4Proof.Bytes()); err != nil {
		t.Fatal(err)
	}
}

func TestGenerateProof(t *testing.T) {
	leaf1 := common.HexToHash("ff00000000000000000000000000000000000000000000000000000000000001")
	leaf2 := common.HexToHash("ff00000000000000000000000000000000000000000000000000000000000002")
	leaf3 := common.HexToHash("ff00000000000000000000000000000000000000000000000000000000000003")
	leaf4 := common.HexToHash("ff00000000000000000000000000000000000000000000000000000000000004")

	merkle, err := NewMerkle(16, []common.Hash{leaf1, leaf2, leaf3, leaf4})

	if err != nil {
		t.Error(err)
	}

	leaf1Proof, _ := merkle.GenerateProof(leaf1)
	leaf2Proof, _ := merkle.GenerateProof(leaf2)
	leaf3Proof, _ := merkle.GenerateProof(leaf3)
	leaf4Proof, _ := merkle.GenerateProof(leaf4)

	if err := merkle.CheckMembership(leaf1, leaf1Proof.Bytes()); err != nil {
		t.Fatal(err)
	}
	if err := merkle.CheckMembership(leaf2, leaf2Proof.Bytes()); err != nil {
		t.Fatal(err)
	}
	if err := merkle.CheckMembership(leaf3, leaf3Proof.Bytes()); err != nil {
		t.Fatal(err)
	}
	if err := merkle.CheckMembership(leaf4, leaf4Proof.Bytes()); err != nil {
		t.Fatal(err)
	}
}

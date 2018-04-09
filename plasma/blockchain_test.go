package plasma

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var bc *BlockChain
var keys []*ecdsa.PrivateKey
var addrs []*common.Address

func TestMain(t *testing.T) {
	initialize(t)

	// deposit 1
	amount := big.NewInt(1000)
	owner1, owner1Key := addrs[0], keys[0]

	if _, err := bc.newDeposit(amount, owner1); err != nil {
		t.Fatal("Failed to create new deposit", err)
	}

	// deposit transaction 1
	b1 := <-bc.newBlock
	tx1 := NewTransaction(
		big1, big0, big0,
		big0, big0, big0,
		owner1, big.NewInt(1000),
		&nullAddress, big0,
		big.NewInt(0))

	if b1.transactionSet[0].Hash() != tx1.Hash() {
		t.Fatal("tx1 is not included into block1")
	}
	if sender, _ := b1.Sender(); sender != bc.config.OperatorAddress {
		t.Fatal("b1 sender and operator address mismatched")
	}

	// deposit 2
	owner2, owner2Key := addrs[1], keys[1]

	if _, err := bc.newDeposit(big.NewInt(1000), owner2); err != nil {
		t.Fatal("Failed to create new deposit", err)
	}

	// deposit transaction 2
	b2 := <-bc.newBlock
	tx2 := NewTransaction(
		big.NewInt(2), big0, big0,
		big0, big0, big0,
		owner2, big.NewInt(1000),
		&nullAddress, big0,
		big.NewInt(0))

	if sender, _ := b1.Sender(); sender != bc.config.OperatorAddress {
		t.Fatal("b2 sender and operator address mismatched")
	}
	if b2.transactionSet[0].Hash() != tx2.Hash() {
		t.Fatal("tx2 is not included into block2")
	}

	// apply 1st transaction
	tx3 := NewTransaction(
		big.NewInt(1), big0, big0,
		big0, big0, big0,
		owner1, big.NewInt(900),
		&nullAddress, big0,
		big.NewInt(100))

	tx3.Sign1(owner1Key)

	if err := bc.applyTransaction(tx3); err != nil {
		t.Fatal("Failed to apply transaction 3", "error", err)
	}

	// apply 2nd transaction
	tx4 := NewTransaction(
		big.NewInt(2), big0, big0,
		big0, big0, big0,
		owner1, big.NewInt(800),
		&nullAddress, big0,
		big.NewInt(200))

	tx4.Sign1(owner2Key)

	if err := bc.applyTransaction(tx4); err != nil {
		t.Fatal("Failed to apply transaction 4", "error", err)
	}

	// submit tx 3, 4
	_, err := bc.submitBlock(bc.config.OperatorPrivateKey)
	if err != nil {
		t.Fatal("Failed to submit block 3")
	}
	b3 := <-bc.newBlock
	if b3.transactionSet[0].Hash() != tx3.Hash() {
		t.Fatal("tx3 is not included into block3")
	}
	if b3.transactionSet[1].Hash() != tx4.Hash() {
		t.Fatal("tx4 is not included into block3")
	}

	// apply 3rd, 4th transaction
	tx5 := NewTransaction(
		big.NewInt(3), big0, big0,
		big.NewInt(3), big1, big0,
		owner1, big.NewInt(1700),
		&nullAddress, big0,
		big.NewInt(0))

	tx5.Sign1(owner1Key)
	tx5.Sign2(owner1Key)

	if err := bc.applyTransaction(tx5); err != nil {
		t.Fatal("Failed to apply transaction 5", "error", err)
	}
}

func initialize(t *testing.T) {
	config := DefaultConfig
	config.OperatorPrivateKey, _ = crypto.HexToECDSA("9cd69f009ac86203e54ec50e3686de95ff6126d3b30a19f926a0fe9323c17181")
	bc = NewBlockChain(&config)

	keyStrs := []string{
		"abf82ff96b463e9d82b83cb9bb450fe87e6166d4db6d7021d0c71d7e960d5abe",
		"dcb7118c9946a39cd40b661e0d368e4afcc3cc48d21aa750d8164ca2e44961c4",
		"2d7aaa9b78d759813448eb26483284cd5e4344a17dede2ab7f062f0757113a28",
		"0e5c6904f09186a0cfe945da201e9d9f0443e07d9e795a9d26cc5cbb882874ac",
		"7f60d75be8f8833a47311c001adbc3771784c52ea9115200a516e3f050c3bc2b",
		"949dbd0607598c41478b32c27da65ab550d54246922fa8978a8c1b9e901e06a6",
		"87a3c9405478581d513a16075038e5869d02311371b757f7163200a09f230f18",
		"e5faea48461ef5a0b78839573073e5a2f579155bf7a4cceb15e49b41963af6a3",
		"ccfb970ed6f3bb68a15d87a67071da16544c918cf978dc41906e686326bb953d",
		"27a3706e23375353aabc8da00d59db6795abae3036dee967103088c8f15e5335",
	}

	for _, keyStr := range keyStrs {
		key, _ := crypto.HexToECDSA(keyStr)
		addr := crypto.PubkeyToAddress(key.PublicKey)

		keys = append(keys, key)
		addrs = append(addrs, &addr)
	}

}

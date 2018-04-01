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

	// deposit
	amount := big.NewInt(1000)
	owner, ownerKey := addrs[0], keys[0]

	if err := bc.newDeposit(amount, owner); err != nil {
		t.Fatal("Failed to create new deposit", err)
	}

	// apply
	tx := NewTransaction(
		big1, big0, big0,
		big0, big0, big0,
		owner, big.NewInt(900),
		&nullAddress, big0,
		big.NewInt(100))

	tx.Sign1(ownerKey)

	if err := bc.applyTransaction(tx); err != nil {
		t.Fatal("Failed to apply transact", "error", err)
	}
}

func initialize(t *testing.T) {
	bc = NewBlockChain(&DefaultConfig)

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

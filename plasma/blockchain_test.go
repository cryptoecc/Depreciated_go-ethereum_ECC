package plasma

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/plasma/types"
)

var bc *BlockChain
var keys []*ecdsa.PrivateKey
var addrs []*common.Address

// to create BlockChain
func verifyBlockMock(block *types.Block) error {
	return nil
}

func TestMain(t *testing.T) {
	initialize(t)

	// deposit 1
	blkNum1 := big.NewInt(1)
	amount := big.NewInt(1000)
	owner1, owner1Key := addrs[0], keys[0]

	if _, err := bc.NewDeposit(amount, owner1, blkNum1); err != nil {
		t.Fatal("Failed to create new deposit", err)
	}

	// deposit transaction 1
	b1 := <-bc.newBlock
	tx1 := types.NewTransaction(
		blkNum1, big0, big0,
		big0, big0, big0,
		owner1, amount,
		&nullAddress, big0,
		big0)

	if b1.Data.TransactionSet[0].Hash() != tx1.Hash() {
		t.Fatal("tx1 is not included into block1")
	}
	if b1.Data.TransactionSet[0].Data.BlkNum1.Cmp(blkNum1) != 0 {
		t.Fatal("tx1 has wrong block number")
	}
	if sender, _ := b1.Sender(); sender != bc.config.OperatorAddress {
		t.Fatal("b1 sender and operator address mismatched")
	}
	if tx, err := bc.GetTransaction(big1, big0); err != nil {
		t.Fatal("bc failed to get tx1", err)
	} else if tx.Hash() != tx1.Hash() {
		t.Fatal("what's wrong")
	}

	// deposit 2
	blkNum2 := big.NewInt(2)
	owner2, owner2Key := addrs[1], keys[1]

	if _, err := bc.NewDeposit(big.NewInt(1000), owner2, blkNum2); err != nil {
		t.Fatal("Failed to create new deposit", err)
	}

	// deposit transaction 2
	b2 := <-bc.newBlock
	tx2 := types.NewTransaction(
		blkNum2, big0, big0,
		big0, big0, big0,
		owner2, big.NewInt(1000),
		&nullAddress, big0,
		big.NewInt(0))

	if sender, _ := b1.Sender(); sender != bc.config.OperatorAddress {
		t.Fatal("b2 sender and operator address mismatched")
	}
	if b2.Data.TransactionSet[0].Data.BlkNum1.Cmp(blkNum2) != 0 {
		t.Fatal("tx2 has wrong block number")
	}
	if b2.Data.TransactionSet[0].Hash() != tx2.Hash() {
		t.Fatal("tx2 is not included into block2")
	}
	if tx, err := bc.GetTransaction(big.NewInt(2), big0); err != nil {
		t.Fatal("bc failed to get tx2", err)
	} else if tx.Hash() != tx2.Hash() {
		t.Fatal("what's wrong")
	}

	// apply 1st transaction
	tx3 := types.NewTransaction(
		big1, big0, big0,
		big0, big0, big0,
		owner1, big.NewInt(900),
		&nullAddress, big0,
		big.NewInt(100))

	tx3.Sign1(owner1Key)

	if err := bc.ApplyTransaction(tx3); err != nil {
		t.Log(tx3.Data)
		t.Fatal("Failed to apply transaction 3", "error", err)
	}

	// apply 2nd transaction
	tx4 := types.NewTransaction(
		big.NewInt(2), big0, big0,
		big0, big0, big0,
		owner1, big.NewInt(800),
		&nullAddress, big0,
		big.NewInt(200))

	tx4.Sign1(owner2Key)

	if err := bc.ApplyTransaction(tx4); err != nil {
		t.Fatal("Failed to apply transaction 4", "error", err)
	}

	// submit tx 3, 4
	blkNum3 := big.NewInt(1000)
	_, err := bc.SubmitBlock(bc.config.OperatorPrivateKey)
	if err != nil {
		t.Fatal("Failed to submit block 3")
	}
	b3 := <-bc.newBlock

	if b3.Data.BlockNumber.Cmp(blkNum3) != 0 {
		t.Fatal("b3 has wrong block number", b3.Data.BlockNumber)
	}

	if b3.Data.TransactionSet[0].Hash() != tx3.Hash() {
		t.Fatal("tx3 is not included into block3")
	}
	if b3.Data.TransactionSet[1].Hash() != tx4.Hash() {
		t.Fatal("tx4 is not included into block3")
	}

	// apply 3rd, 4th transaction
	tx5 := types.NewTransaction(
		big.NewInt(1000), big0, big0,
		big.NewInt(1000), big1, big0,
		owner1, big.NewInt(1700),
		&nullAddress, big0,
		big.NewInt(0))

	tx5.Sign1(owner1Key)
	tx5.Sign2(owner1Key)

	if err := bc.ApplyTransaction(tx5); err != nil {
		t.Fatal("Failed to apply transaction 5", "error", err)
	}
}

func initialize(t *testing.T) {
	config := DefaultConfig
	config.OperatorPrivateKey, _ = crypto.HexToECDSA("9cd69f009ac86203e54ec50e3686de95ff6126d3b30a19f926a0fe9323c17181")
	bc = NewBlockChain(&config, verifyBlockMock)

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

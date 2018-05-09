package types

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestHashEquality(t *testing.T) {
	addr1 := common.StringToAddress("a")
	addr2 := common.StringToAddress("b")

	tx1 := NewTransaction(
		big0, big0, big0,
		big0, big0, big0,
		&nullAddress, big0,
		&nullAddress, big0,
		big0)

	tx2 := NewTransaction(
		big1, big0, big0,
		big0, big0, big0,
		&nullAddress, big0,
		&nullAddress, big0,
		big0)

	tx3 := NewTransaction(
		big1, big1, big0,
		big0, big0, big0,
		&nullAddress, big0,
		&nullAddress, big0,
		big0)

	tx4 := NewTransaction(
		big1, big1, big1,
		big1, big0, big0,
		&nullAddress, big0,
		&nullAddress, big0,
		big0)

	tx5 := NewTransaction(
		big1, big1, big1,
		big1, big1, big0,
		&nullAddress, big0,
		&nullAddress, big0,
		big0)

	tx6 := NewTransaction(
		big1, big1, big1,
		big1, big1, big1,
		&nullAddress, big0,
		&nullAddress, big0,
		big0)

	tx7 := NewTransaction(
		big1, big1, big1,
		big1, big1, big1,
		&addr1, big0,
		&nullAddress, big0,
		big0)

	tx8 := NewTransaction(
		big1, big1, big1,
		big1, big1, big1,
		&addr1, big1,
		&nullAddress, big0,
		big0)

	tx9 := NewTransaction(
		big1, big1, big1,
		big1, big1, big1,
		&addr1, big1,
		&addr2, big0,
		big0)

	tx10 := NewTransaction(
		big1, big1, big1,
		big1, big1, big1,
		&addr1, big1,
		&addr2, big1,
		big0)

	tx11 := NewTransaction(
		big1, big1, big1,
		big1, big1, big1,
		&addr1, big1,
		&addr2, big1,
		big1)

	txs := []*Transaction{tx1, tx2, tx3, tx4, tx5, tx6, tx7, tx8, tx9, tx10, tx11}
	captions := []string{"tx1", "tx2", "tx3", "tx4", "tx5", "tx6", "tx7", "tx8", "tx9", "tx10", "tx11"}

	for i := 0; i < len(txs); i++ {
		for j := 0; j < len(txs); j++ {
			if i != j {
				testTxHash(t, txs[i], txs[j], captions[i], captions[j])
			}
		}
	}
}

func testTxHash(t *testing.T, tx1, tx2 *Transaction, caption1, caption2 string) {
	if tx1.Hash() == tx2.Hash() {
		t.Fatalf("%s and %s have same hash with %s", caption1, caption2, tx1.Hash().Hex())
	}
}

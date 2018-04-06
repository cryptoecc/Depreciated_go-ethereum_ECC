package plasma

import (
	"context"
	// "encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	// "sync"
	// "time"
)

// PlasmaAPI provides an API to access Plasma related information.
type PlasmaAPI struct {
}

// PlasmaOperatorAPI provides an API for plasma operator.
type PlasmaOperatorAPI struct {
}

type PublicPlasmaAPI struct {
	pls *Plasma
}

type TxArgs struct {
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
	from1     *common.Address // sign input 1
	from2     *common.Address // sign input 2
}

func NewPlasmaAPI() PlasmaAPI {
	return PlasmaAPI{}
}

func NewPlasmaOperatorAPI() PlasmaOperatorAPI {
	return PlasmaOperatorAPI{}
}

func NewPublicPlasmaAPI(pls *Plasma) *PublicPlasmaAPI {
	api := &PublicPlasmaAPI{
		pls: pls,
	}
	return api
}

// Version returns the Whisper sub-protocol version.
func (api *PublicPlasmaAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}

func (api *PublicPlasmaAPI) Deposit(ctx context.Context) string {
	return ""
}

func (api *PublicPlasmaAPI) CurrentBlockNumber(ctx context.Context) uint64 {
	return api.pls.blockchain.getCurrentBlockNumber().Uint64()
}

func (api *PublicPlasmaAPI) SubmitBlock(ctx context.Context) (common.Hash, error) {
	return api.pls.blockchain.submitBlock(api.pls.config.OperatorPrivateKey)
}

func (api *PublicPlasmaAPI) ApplyTransaction(ctx context.Context, args TxArgs) (common.Hash, error) {
	tx := NewTransaction(
		args.blkNum1, args.txIndex1, args.oIndex1,
		args.blkNum2, args.txIndex2, args.oIndex2,
		args.newOwner1, args.amount1,
		args.newOwner2, args.amount2,
		args.fee)

	txHash := tx.Hash().Bytes()

	if args.from1 != nil {
		sig, err := api.pls.sign(txHash, *args.from1)

		if err != nil {
			return common.BytesToHash(nil), err
		}

		tx.sig1 = sig
	}

	if args.from2 != nil {
		sig, err := api.pls.sign(txHash, *args.from2)

		if err != nil {
			return common.BytesToHash(nil), err
		}

		tx.sig2 = sig
	}

	log.Info("[Plasma API] apply transaction", "tx", tx)

	return tx.Hash(), api.pls.blockchain.applyTransaction(tx)
}

func (api *PublicPlasmaAPI) GetBlock(ctx context.Context, blkNum *big.Int) (map[string]interface{}, error) {
	b, err := api.pls.blockchain.getBlock(blkNum)

	if err != nil {
		return nil, err
	}

	return b.ToRPCResponse(), nil
}

func (api *PublicPlasmaAPI) GetTransaction(ctx context.Context, blkNum, txIndex *big.Int) (map[string]interface{}, error) {
	tx, err := api.pls.blockchain.getTransaction(blkNum, txIndex)

	if err != nil {
		return nil, err
	}

	return tx.ToRPCResponse(), nil
}

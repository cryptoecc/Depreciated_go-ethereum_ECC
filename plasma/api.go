package plasma

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/plasma/types"
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
	BlkNum1   *hexutil.Big    `json:"blkNum1"`
	TxIndex1  *hexutil.Big    `json:"txIndex1"`
	OIndex1   *hexutil.Big    `json:"oIndex1"`
	BlkNum2   *hexutil.Big    `json:"blkNum2"`
	TxIndex2  *hexutil.Big    `json:"txIndex2"`
	OIndex2   *hexutil.Big    `json:"oIndex2"`
	NewOwner1 *common.Address `json:"newOwner1"`
	Amount1   *hexutil.Big    `json:"amount1"`
	NewOwner2 *common.Address `json:"newOwner2"`
	Amount2   *hexutil.Big    `json:"amount2"`
	Fee       *hexutil.Big    `json:"fee"`
	From1     *common.Address `json:"from1"` // sign input 1
	From2     *common.Address `json:"from2"` // sign input 2
}

// setDefaults is a helper function that fills in default values for unspecified tx fields.
func (args *TxArgs) setDefaults(ctx context.Context, b Backend) error {
	log.Info("[Plasma api] TxArgs.setDefaults", "args.BlkNum1", args.BlkNum1)
	return nil
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

	if args.BlkNum1 == nil || args.BlkNum2 == nil {
		return common.Hash{}, errors.New("Failed to read arguments")
	}

	if args.BlkNum1.ToInt().Cmp(big0) > 0 && args.From1 == nil {
		return common.Hash{}, errors.New("Transaction input 1 should be signed by owner")
	}

	if args.BlkNum2.ToInt().Cmp(big0) > 0 && args.From2 == nil {
		return common.Hash{}, errors.New("Transaction input 2 should be signed by owner")
	}

	tx := types.NewTransaction(
		args.BlkNum1.ToInt(), args.TxIndex1.ToInt(), args.OIndex1.ToInt(),
		args.BlkNum2.ToInt(), args.TxIndex2.ToInt(), args.OIndex2.ToInt(),
		args.NewOwner1, args.Amount1.ToInt(),
		args.NewOwner2, args.Amount2.ToInt(),
		args.Fee.ToInt())

	txHash := tx.Hash().Bytes()

	if args.From1 != nil {
		sig, err := api.pls.sign(txHash, *args.From1)

		if err != nil {
			return common.Hash{}, err
		}

		tx.SetSig1(sig)
	}

	if args.From2 != nil {
		sig, err := api.pls.sign(txHash, *args.From2)

		if err != nil {
			return common.Hash{}, err
		}

		tx.SetSig2(sig)
	}

	log.Info("[Plasma API] apply transaction", "txhash", tx.Hash(), "args", args)

	return tx.Hash(), api.pls.blockchain.applyTransaction(tx)
}

func (api *PublicPlasmaAPI) GetBlock(ctx context.Context, BlkNum *hexutil.Big) (map[string]interface{}, error) {
	b, err := api.pls.blockchain.getBlock(BlkNum.ToInt())

	if err != nil {
		return nil, err
	}

	return b.ToRPCResponse(), nil
}

func (api *PublicPlasmaAPI) GetTransaction(ctx context.Context, BlkNum, TxIndex *big.Int) (map[string]interface{}, error) {
	tx, err := api.pls.blockchain.getTransaction(BlkNum, TxIndex)

	if err != nil {
		return nil, err
	}

	return tx.ToRPCResponse(), nil
}

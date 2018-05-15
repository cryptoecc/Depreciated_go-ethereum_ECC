package plasma

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/plasma/types"
)

// Official short name of the protocol used during capability negotiation.
const ProtocolName = "pls"
const ProtocolVersion = 1

const ProtocolMaxMsgSize = 10 * 1024 * 1024 // Maximum cap on the size of a protocol message

// pls protocol message codes
/**
 * TODO
 *  1. exchange node's block number, tx hash info
 *  2. mark the node has block or tx
 *  3. broadcast node blocks or txs it doesn't have
 */
const (
	// Messages for blockchain
	StatusCode          = 0x00 // used by plasma protocol
	NewBlockCode        = 0x01 // broadcast a single block
	NewBlocksCode       = 0x02 // broadcast batch of block
	NewTransactionsCode = 0x03 // broadcast TX
	GetBlockCode        = 0x04 // request a single block
	GetBlocksCode       = 0x05 // request batch of blocks

	// Messages for node info
	OperatorCode = 0x07 // operator node info
	PingCode     = 0x08 // ping
	PongCode     = 0x09 // pong

	// Number of implemented messages
	ProtocolLength = 0x10
)

type errCode int

const (
	ErrMsgTooLarge = iota
	ErrDecode
	ErrInvalidMsgCode
	ErrProtocolVersionMismatch
	ErrOperatorAddressMismatch
	ErrContractAddressMismatch
	ErrNoStatusMsg
	ErrExtraStatusMsg
	ErrSuspendedPeer
)

func (e errCode) String() string {
	return errorToString[int(e)]
}

// XXX change once legacy code is out
var errorToString = map[int]string{
	ErrMsgTooLarge:             "Message too long",
	ErrDecode:                  "Invalid message",
	ErrInvalidMsgCode:          "Invalid message code",
	ErrProtocolVersionMismatch: "Protocol version mismatch",
	ErrOperatorAddressMismatch: "Operator mismatch",
	ErrContractAddressMismatch: "Rootchain contract mismatch",
	ErrNoStatusMsg:             "No status message",
	ErrExtraStatusMsg:          "Extra status message",
	ErrSuspendedPeer:           "Suspended peer",
}

var (
	big0        = big.NewInt(0)
	big1        = big.NewInt(1)
	nullAddress = common.HexToAddress("")
)

// statusData is the network packet for the status message.
type statusData struct {
	ProtocolVersion uint64
	OperatorAddress common.Address
	ContractAddress common.Address
	HighestEthBlock uint64
}

type operatorData struct {
	NodeURL string
}

type newBlockData struct {
	Block *types.Block
}

type pingData struct {
	Number uint64
}

type pongData struct {
	Block *types.Block
}

package plasma

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	ProtocolVersion    = uint64(1)
	ProtocolVersionStr = "1"
	ProtocolName       = "pls"

	statusCode           = 0x01 // used by plasma protocol
	operatorCode         = 0x02 // operator node info
	newBlockCode         = 0x03 // broadcast block
	newTransactionCode   = 0x04 // broadcast TX
	getBlockCode         = 0x05 // request block by hash or number
	numberOfMessageCodes = 0x06

	MaxMessageSize        = uint32(10 * 1024 * 1024) // maximum accepted size of a message.
	DefaultMaxMessageSize = uint32(1024 * 1024)
)

type errCode int

const (
	ErrMsgTooLarge = iota
	ErrDecode
	ErrInvalidMsgCode
	ErrProtocolVersionMismatch
)

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

type getBlockData struct {
	Number uint64
	Hash   common.Hash
}

type newBlockData struct {
	Block *Block
}

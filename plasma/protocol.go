package plasma

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	ProtocolVersion    = uint64(1)
	ProtocolVersionStr = "1"
	ProtocolName       = "pls"

	statusCode           = 0 // used by plasma protocol
	messagesCode         = 1 // normal whisper message
	p2pCode              = 2 // peer-to-peer message (to be consumed by the peer, but not forwarded any further)
	p2pRequestCode       = 3 // peer-to-peer message, used by Dapp protocol
	NumberOfMessageCodes = 64

	MaxMessageSize        = uint32(10 * 1024 * 1024) // maximum accepted size of a message.
	DefaultMaxMessageSize = uint32(1024 * 1024)

	expirationCycle   = time.Second
	transmissionCycle = 300 * time.Millisecond
)

var (
	big0        = big.NewInt(0)
	big1        = big.NewInt(1)
	nullAddress = common.HexToAddress("")
)

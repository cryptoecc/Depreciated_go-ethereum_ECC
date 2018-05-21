package plasma

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

// Config represents the configuration state of a plasma node.
type Config struct {
	MaxPeers int

	// Address of plasma contract on root chain
	ContractAddress common.Address

	// Address of plasma operator
	// TODO: load this address from Plasma contract on Ethereum network
	OperatorAddress common.Address

	// If node is operator, specify the private key
	// TODO: chagne to nodePrivateKey
	OperatorPrivateKey *ecdsa.PrivateKey

	// TODO: assign this field
	IsOperator bool

	// Plasma operator node
	OperatorNode    *Peer
	OperatorNodeURL string
	OperatorNodeID  string // TODO: replace above 2 fields

	// BlockChain specific configs
	DataDir string
	OnDisk  bool // disk or memory
}

var DefaultConfig = Config{
	MaxPeers:        25,
	ContractAddress: common.StringToAddress("0x0"),
	OperatorAddress: params.PlasmaOperatorAddress,
}

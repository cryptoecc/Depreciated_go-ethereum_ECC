package plasma

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

// Config represents the configuration state of a plasma node.
type Config struct {
	// Address of plasma contract on root chain
	ContractAddress common.Address

	// Address of plasma operator
	// TODO: load this address from Plasma contract on Ethereum network
	OperatorAddress common.Address

	// If this node is operator, specify the private key
	OperatorPrivateKey *ecdsa.PrivateKey
}

// BlockChainConfig represents BlockChain specific configuration.
type BlockChainConfig struct {
	Config

	DataDir string
	OnDisk  bool // disk or memory
}

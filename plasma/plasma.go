package plasma

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/contracts/plasma/contract"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/rpc"
)

// Plasma implements the Plasma full node service
type Plasma struct {
	config   *Config
	protocol p2p.Protocol
	context  context.Context

	// Channels
	quit   chan bool
	client chan *rpc.Client

	// Handlers
	backend    *ethclient.Client // actual rpc backend
	blockchain *BlockChain
	downloader *Downloader

	eventMux       *event.TypeMux
	accountManager *accounts.Manager

	ApiBackend *Backend

	lock sync.RWMutex
}

func New(config *Config) *Plasma {
	if config == nil {
		config = &DefaultConfig
	}

	pls := &Plasma{
		config:  config,
		context: context.Background(),
	}

	pls.client = make(chan *rpc.Client)
	pls.quit = make(chan bool)

	pls.protocol = p2p.Protocol{
		Name:    ProtocolName,
		Version: uint(ProtocolVersion),
		Length:  NumberOfMessageCodes,
		Run:     pls.HandlePeer,
		NodeInfo: func() interface{} {
			return map[string]interface{}{
				"version":      ProtocolVersionStr,
				"currentBlock": pls.CurrentBlockNumber(),
				"numPeers":     len(pls.getPeers()),
			}
		},
	}

	return pls
}

func (pls *Plasma) RegisterRpcClient(rpcClient *rpc.Client) {
	if rpcClient == nil {
		log.Warn("Cannot register nil RPC client to Plasma")
	} else {
		pls.client <- rpcClient
	}
}

// Start implements node.Service, starting the background data propagation thread
// of the Plasma protocol.
func (pls *Plasma) Start(*p2p.Server) error {
	go pls.run()

	log.Info("Plama started", "version", ProtocolVersionStr)
	return nil
}

// Stop implements node.Service, stopping the background data propagation thread
// of the Plasma protocol.
func (pls *Plasma) Stop() error {
	close(pls.quit)
	return nil
}

// Protocols implements node.Service, retrieving the P2P protocols the service wishes to start.
func (pls *Plasma) Protocols() []p2p.Protocol {
	return []p2p.Protocol{pls.protocol}
}

// Version returns the plasma sub-protocols version number.
func (pls *Plasma) Version() uint {
	return pls.protocol.Version
}

// APIs implements node.Service, retrieving the list of RPC descriptors the service provides
func (pls *Plasma) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: ProtocolName,
			Version:   ProtocolVersionStr,
			Service:   NewPublicPlasmaAPI(pls),
			Public:    true,
		},
	}
}

// HandlePeer is called by the underlying P2P layer when the plasma sub-protocol
// connection is negotiated.
func (pls *Plasma) HandlePeer(peer *p2p.Peer, rw p2p.MsgReadWriter) error {
	return nil
}

func (pls *Plasma) CurrentBlockNumber() uint64 {
	return pls.blockchain.getCurrentBlockNumber().Uint64()
}

func (pls *Plasma) getPeers() []*discover.Node {
	return nil
}

func (pls *Plasma) run() {
	select {
	case rpcClient := <-pls.client:
		pls.backend = ethclient.NewClient(rpcClient)
		log.Info("RPC Client attached")
	case <-pls.quit:
		return
	}

	if err := pls.initialize(); err != nil {
		log.Info("Plasma failed to initialize: %v", err)
	}

	log.Info("Plasma initialized and running")

loop:
	for {
		switch {

		case <-pls.quit:
			break loop

		default:
			if err := pls.checkNextBlock(); err != nil {
				log.Info("Plasma failed to fetch next block: %v", err)
			}
		}
	}
}

// TODO: Load contract instnace. If operator, deploy plasma contract.
func (pls *Plasma) initialize() error {
	// deploy plasma contract
	// TODO: check if contract is deployed

	deployed, err := pls.checkContractDepoyed()

	if err != nil {
		return err
	}

	if deployed {
		log.Info("Plasma contract is already deployed", "address", pls.config.ContractAddress)
	} else {
		transactOpts := bind.NewKeyedTransactor(pls.config.OperatorPrivateKey)

		addr, _, _, err := contract.DeployRootChain(transactOpts, pls.backend)

		if err != nil {
			return err
		}

		log.Info("Plasma contract deployed", "txhash", addr)
		// TODO: fetch deployed contract address
	}

	return nil
}

// TODO: read next block from Plasma contract
// If new block is submitted to the contract, request raw block data to operator
func (pls *Plasma) checkNextBlock() error {
	return nil
}

func (pls *Plasma) checkContractDepoyed() (bool, error) {
	// nil to recent block
	code, err := pls.backend.CodeAt(pls.context, pls.config.ContractAddress, nil)
	if err != nil {
		return false, err
	} else {
		return len(code) > 0, nil
	}
}

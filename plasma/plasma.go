package plasma

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/contracts/plasma/contract"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rpc"
)

// Plasma implements the Plasma full node service
type Plasma struct {
	config    *Config
	protocol  p2p.Protocol
	context   context.Context
	rootchain *contract.RootChain

	// Channels
	quit        chan bool
	backendChan chan *ethclient.Client

	// Handlers
	server     *p2p.Server
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

	pls.backendChan = make(chan *ethclient.Client)
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

	pls.blockchain = NewBlockChain(config)

	return pls
}

func (pls *Plasma) RegisterRpcClient(rpcClient *rpc.Client) {
	if rpcClient == nil {
		log.Warn("[Plasma] Cannot register nil RPC client to Plasma")
	} else {
		pls.backendChan <- ethclient.NewClient(rpcClient)
	}
}

// RegisterClient register endpoint of ethereum jsonrpc for Plasma single node
func (pls *Plasma) RegisterClient(backend *ethclient.Client) {
	if backend == nil {
		log.Warn("[Plasma] Cannot register nil endpoint to Plasma")
	} else {
		pls.backendChan <- backend
	}
}

// Start implements node.Service, starting the background data propagation thread
// of the Plasma protocol.
func (pls *Plasma) Start(server *p2p.Server) error {
	pls.server = server

	if pls.isOperator() {
		pls.config.OperatorNode = server.Self()
	}

	go pls.run()

	log.Info("[Plasma] node started", "version", ProtocolVersionStr)
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
	case backend := <-pls.backendChan:
		pls.backend = backend
		log.Info("[Plasma] Ethereum jsonrpc backend attached")
	case <-pls.quit:
		return
	}

	if err := pls.initialize(); err != nil {
		log.Info("[Plasma] Failed to initialize", err)
	}

	log.Info("[Plasma] node initialized and running")

loop:
	for {
		switch {

		case <-pls.quit:
			break loop

		}
	}
}

// TODO: If operator, deploy or load contract. If not operator, load contract
func (pls *Plasma) initialize() error {
	// deploy or load plasma contract
	deployed, err := pls.checkContractDepoyed()

	if err != nil {
		return err
	}

	if deployed {
		rootchain, err := contract.NewRootChain(pls.config.ContractAddress, pls.backend)
		if err != nil {
			return err
		}

		pls.rootchain = rootchain

		log.Info("[Plasma] Contract is already deployed", "address", pls.config.ContractAddress)
	} else {
		if !pls.isOperator() {
			return fmt.Errorf("[Plasma] Contract is not deployed yet at", pls.config.ContractAddress)
		}

		transactOpts := bind.NewKeyedTransactor(pls.config.OperatorPrivateKey)

		address, tx, rootchain, err := contract.DeployRootChain(transactOpts, pls.backend)

		if err != nil {
			return err
		}

		pls.config.ContractAddress = address
		pls.rootchain = rootchain
		log.Info("[Plasma] Contract deployed", "hash", tx.Hash(), "contract", address)
	}

	// run deposit listener
	err = pls.listenDeposit()
	if err != nil {
		return err
	}

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

func (pls *Plasma) isOperator() bool {
	operatorAddress := crypto.PubkeyToAddress(pls.config.OperatorPrivateKey.PublicKey)

	return operatorAddress == params.PlasmaOperatorAddress
}

// watch Deposit event
func (pls *Plasma) listenDeposit() error {
	filterer, err := contract.NewRootChainFilterer(pls.config.ContractAddress, pls.backend)

	if err != nil {
		return err
	}

	// TODO: If plasma node had stopped previously, read event from last parent block to stoped
	watchOpts := bind.WatchOpts{
		Context: pls.context,
		Start:   nil,
	}

	sink := make(chan *contract.RootChainDeposit)

	sub, err := filterer.WatchDeposit(&watchOpts, sink)

	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case deposit := <-sink:
				if deposit != nil {
					log.Info("[Plasma] New deposit on plasma contract", "depositor", deposit.Depositor, "amount", deposit.Amount)
					if err := pls.blockchain.newDeposit(deposit.Amount, &deposit.Depositor); err != nil {
						log.Warn("[Plasma] Failed to add new deposit from rootchain", err)
					}
				}

			case <-pls.quit:
				sub.Unsubscribe()
				return
			case err := <-sub.Err():
				log.Warn("[Plasma] Deposit subscription error", err)
				sub.Unsubscribe()
				return
			}
		}
	}()

	return nil
}

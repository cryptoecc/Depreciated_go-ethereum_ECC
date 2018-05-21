package plasma

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/plasma/contract"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/plasma/types"
	"github.com/ethereum/go-ethereum/rpc"
)

var invalidBlockHash = errors.New("Wrong block hash")

// Plasma implements the Plasma full node service
type Plasma struct {
	config   *Config
	protocol p2p.Protocol
	context  context.Context

	// RootCHain contract binding
	rootchain *contract.RootChain

	// options to send ethereum transcation
	transactOpts *bind.TransactOpts

	// Channels
	quit          chan bool
	backendCh     chan *ethclient.Client
	rootchainCh   chan struct{}
	rootchainLoad bool // checks rootchain contract instance is loaded

	eventMux *event.TypeMux

	// Peers
	peers *peerSet

	// Handlers
	protocolManager *ProtocolManager
	server          *p2p.Server
	backend         *ethclient.Client // actual rpc backend
	blockchain      *BlockChain       // Plasma blockchain

	accountManager *accounts.Manager // node account manager

	ApiBackend *Backend // pls api backend

	lock     sync.RWMutex
	peerLock sync.RWMutex
}

// New creates Plasma instance
func New(ctx *node.ServiceContext, config *Config) *Plasma {
	if config == nil {
		config = &DefaultConfig
	}

	pls := &Plasma{
		config:  config,
		context: context.Background(),

		backendCh:   make(chan *ethclient.Client),
		rootchainCh: make(chan struct{}, 1),
		quit:        make(chan bool),

		eventMux: ctx.EventMux,

		peers:          newPeerSet(),
		accountManager: ctx.AccountManager,
	}

	verifier := func(block *types.Block) error { return pls.verifyBlock(block) }
	higheter := func() []uint64 { return pls.higheter() }

	blockchain := NewBlockChain(config, verifier)
	pls.blockchain = blockchain

	pls.protocolManager = NewProtocolManager(config, pls.isOperator(), pls.eventMux, blockchain, higheter)

	// TODO: sync pm.protocol
	// pls.protocol = p2p.Protocol{
	// 	Name:    ProtocolName,
	// 	Version: ProtocolVersion,
	// 	Length:  ProtocolLength,
	// 	Run:     pls.HandlePeer,
	// 	NodeInfo: func() interface{} {
	// 		return map[string]interface{}{
	// 			"version":      ProtocolVersion,
	// 			"currentBlock": pls.CurrentBlockNumber(),
	// 			"numPeers":     len(pls.getPeers()),
	// 		}
	// 	},
	// 	PeerInfo: func(id discover.NodeID) interface{} {
	// 		return id
	// 	},
	// }

	return pls
}

// RegisterRpcClient takes node's rpc client to register ethclient as plasma's ethereum backend
func (pls *Plasma) RegisterRpcClient(rpcClient *rpc.Client) {
	if rpcClient == nil {
		log.Warn("[Plasma] Cannot register nil RPC client to Plasma")
	} else {
		pls.backendCh <- ethclient.NewClient(rpcClient)
	}
}

// RegisterClient registers endpoint of ethereum jsonrpc for Plasma single node
func (pls *Plasma) RegisterClient(backend *ethclient.Client) {
	if backend == nil {
		log.Warn("[Plasma] Cannot register nil endpoint to Plasma")
	} else {
		pls.backendCh <- backend
	}
}

// Start implements node.Service, starting the background data propagation thread
// of the Plasma protocol.
func (pls *Plasma) Start(server *p2p.Server) error {
	pls.server = server

	pls.protocolManager.Start(pls.config.MaxPeers)
	go pls.run()

	log.Info("[Plasma] node started", "version", ProtocolVersion, "operator", pls.config.OperatorAddress, "contract", pls.config.ContractAddress)
	return nil
}

// Stop implements node.Service, stopping the background data propagation thread
// of the Plasma protocol.
func (pls *Plasma) Stop() error {
	pls.protocolManager.Stop()
	close(pls.quit)
	return nil
}

// Protocols implements node.Service, retrieving the P2P protocols the service wishes to start.
func (pls *Plasma) Protocols() []p2p.Protocol {
	return []p2p.Protocol{pls.protocolManager.Protocol}
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
			Version:   string(ProtocolVersion),
			Service:   NewPublicPlasmaAPI(pls),
			Public:    true,
		},
	}
}

func (pls *Plasma) run() {
	select {
	case backend := <-pls.backendCh:
		pls.backend = backend
		log.Info("[Plasma] Ethereum jsonrpc backend attached")
	case <-pls.quit:
		return
	}

	log.Info("[Plasma] node initializing")
	if err := pls.initialize(); err != nil {
		log.Info("[Plasma] Failed to initialize", "err", err)
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
	// wait until ethereum is synced
	if err := pls.waitEthSynced(); err != nil {
		log.Info("[Plasma] Failed to wait Eth syncing", "err", err)
		return err
	}
	log.Info("[Plasma] Eth is synced. check contract deployment")

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
			return fmt.Errorf("[Plasma] Contract is not deployed yet at", "contract", pls.config.ContractAddress)
		}

		pls.transactOpts = bind.NewKeyedTransactor(pls.config.OperatorPrivateKey)

		address, tx, rootchain, err := contract.DeployRootChain(pls.transactOpts, pls.backend)

		if err != nil {
			return err
		}

		pls.config.ContractAddress = address
		pls.rootchain = rootchain
		log.Info("[Plasma] Contract deployed", "hash", tx.Hash(), "contract", address)
	}

	pls.rootchainCh <- struct{}{}
	pls.rootchainLoad = true

	// run deposit listener
	err = pls.listenDeposit()
	if err != nil {
		return err
	}

	// run submit listener
	err = pls.listenSubmit()
	if err != nil {
		return err
	}

	if pls.isOperator() {
		if err := pls.addSubmitListener(); err != nil {
			return err
		}
	}

	return nil
}

// waitEthSynced waits until all of ethereum blockchain data is downloaded
func (pls *Plasma) waitEthSynced() error {
	// Assume operator node run only after ethereum synced
	if pls.isOperator() {
		return nil
	}

	log.Info("[Plasma] wait until operator peer is connedted")
	// Wait operator peer is conencted
	<-pls.protocolManager.operatorNodeCh
	log.Info("[Plasma] operator peer is connedted")

	// Wait until syncing is finished
	for {
		time.Sleep(time.Second * 10)
		result, err := pls.backend.SyncProgress(pls.context)

		fmt.Println(result, err)

		if err != nil {
			return err
		}

		if result != nil {
			log.Info("[Plasma] wait until eth is synced", "current", result.CurrentBlock, "highest", result.HighestBlock)
		}

		if result == nil {
			log.Info("[Plasma] Ethereum is synced")
			return nil
		}
	}
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
	if pls.config.OperatorPrivateKey == nil {
		return false
	}

	return params.PlasmaOperatorAddress == crypto.PubkeyToAddress(pls.config.OperatorPrivateKey.PublicKey)
}

// watch Deposit event
func (pls *Plasma) listenDeposit() error {
	filterer, err := contract.NewRootChainFilterer(pls.config.ContractAddress, pls.backend)

	if err != nil {
		return err
	}

	// TODO: If plasma node had stopped previously, read event from last parent block when node stopped
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

					if _, err := pls.blockchain.NewDeposit(deposit.Amount, &deposit.Depositor, deposit.DepositBlock); err != nil {
						log.Warn("[Plasma] Failed to add new deposit from rootchain", "err", err)
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

// watch Submit event
func (pls *Plasma) listenSubmit() error {
	if pls.isOperator() {
		return nil
	}

	filterer, err := contract.NewRootChainFilterer(pls.config.ContractAddress, pls.backend)

	if err != nil {
		return err
	}

	// TODO: If plasma node had stopped previously, read event from last parent block when node stopped
	watchOpts := bind.WatchOpts{
		Context: pls.context,
		Start:   nil,
	}

	sink := make(chan *contract.RootChainSubmit)

	sub, err := filterer.WatchSubmit(&watchOpts, sink)

	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case submitBlock := <-sink:
				if submitBlock != nil {
					log.Info("[Plasma] New block is submitted", "blkNum", submitBlock.SubmitBlock, "root", submitBlock.Root)

					// TODO: enqueue request
					operator := pls.protocolManager.peers.Operator()
					blkNum := submitBlock.SubmitBlock.Uint64()
					if err := operator.RequestBlock(blkNum); err != nil {
						log.Warn("[Plasma] Failed to request submit-block to operator", "err", err)
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

// addSubmitListener send new sealded block to root chain
func (pls *Plasma) addSubmitListener() error {
	listener := func(blk *types.Block) error {
		if len(blk.Data.TransactionSet) == 1 && blk.Data.TransactionSet[0].Data.BlkNum1.Cmp(blk.Data.BlockNumber) == 0 {
			log.Info("[Plasma] New deposit block added", "hash", blk.Hash(), "blkNum", blk.Data.BlockNumber)
			return nil
		}

		tx, err := pls.rootchain.SubmitBlock(pls.transactOpts, blk.Hash())

		if err != nil {
			log.Info("[Plasma] Failed to submimt new block", "hash", blk.Hash(), "err", err)
		} else {
			log.Info("[Plasma] Submimt new block", "blkhash", blk.Hash(), "txhash", tx.Hash())
		}
		return nil
	}

	return pls.blockchain.AddNewBlockListener(listener)
}

// sign any bytes from unlocked ethereum account
func (pls *Plasma) sign(hash []byte, from common.Address) ([]byte, error) {
	// Look up the wallet containing the requested signer
	account := accounts.Account{Address: from}
	wallet, err := pls.accountManager.Find(account)
	if err != nil {
		return nil, err
	}

	return wallet.SignHash(account, hash)
}

// verify block hash on root
func (pls *Plasma) verifyBlock(block *types.Block) error {
	callOpts := bind.CallOpts{
		Pending: false,
		Context: pls.context,
	}

	res, err := pls.rootchain.ChildChain(&callOpts, block.Data.BlockNumber)

	if err != nil {
		return err
	}

	root1 := block.Hash().Bytes()
	root2 := res.Root[:]

	if !bytes.Equal(root1, root2) {
		log.Warn("[Plasma] wrong depsoit block hash", "root1", root1, "root2", root2)
		return invalidBlockHash
	}

	return nil
}

// get block numbers not downloaded yet
func (pls *Plasma) higheter() []uint64 {
	// wait rootchain load
	if !pls.rootchainLoad {
		<-pls.rootchainCh
	}

	callOpts := bind.CallOpts{
		Pending: false,
		Context: pls.context,
	}

	localBlkNumBig := pls.blockchain.GetCurrentBlockNumber()
	remoteBlkNumBig, err := pls.rootchain.CurrentChildBlock(&callOpts)

	if err != nil {
		return nil
	}

	log.Info("[Plasma] Checking plasma block", "localBlkNum", localBlkNumBig, "remoteBlkNum", remoteBlkNumBig)

	localBlkNum := localBlkNumBig.Uint64()
	remoteBlkNum := remoteBlkNumBig.Uint64()

	childBlockInterval := pls.blockchain.blockInterval.Uint64()
	var epochs, blocksToRequest []uint64

	for blkNum := localBlkNum - childBlockInterval; blkNum <= remoteBlkNum-childBlockInterval; blkNum += childBlockInterval {
		log.Info("[Plasma] adding epochs", "blkNum", blkNum)
		epochs = append(epochs, blkNum)
	}

	for _, blkNum := range epochs {
		// submit-block
		if blkNum > 0 {
			log.Info("[Plasma] add submit block request queue", "blkNum", blkNum)
			blocksToRequest = append(blocksToRequest, blkNum)
		}

		// deposit-block
		var i uint64
		for i = 1; i < childBlockInterval; i++ {
			depositBlockNumber := blkNum + i

			res, err := pls.rootchain.ChildChain(&callOpts, big.NewInt(int64(depositBlockNumber)))

			isEmpty := true

			for i := 0; i < len(res.Root); i++ {
				if res.Root[i] != 0x0 {
					isEmpty = false
					break
				}
			}

			if err != nil || isEmpty {
				log.Info("[Plasma] stop", "err", err, "isEmpty", isEmpty)
				break
			}

			log.Info("[Plasma] add deposit block request queue", "blkNum", depositBlockNumber)
			blocksToRequest = append(blocksToRequest, depositBlockNumber)
		}
	}

	return blocksToRequest
}

package plasma

import (
	"bytes"
	"context"
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
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discover"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/plasma/types"
	"github.com/ethereum/go-ethereum/rpc"
)

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
	quit           chan bool
	backendCh      chan *ethclient.Client
	operatorNodeCh chan *Peer

	// Peers
	peers map[*Peer]struct{} // Set of currently active peers

	// Handlers
	server     *p2p.Server
	backend    *ethclient.Client // actual rpc backend
	blockchain *BlockChain       // Plasma blockchain
	downloader *Downloader       // Plasma downloader (TODO: implements this)

	eventMux       *event.TypeMux
	accountManager *accounts.Manager // node account manager

	ApiBackend *Backend // pls api backend

	lock     sync.RWMutex
	peerLock sync.RWMutex
}

// New creates Plasma instance
func New(config *Config, accountManager *accounts.Manager) *Plasma {
	if config == nil {
		config = &DefaultConfig
	}

	pls := &Plasma{
		config:         config,
		context:        context.Background(),
		accountManager: accountManager,
	}

	pls.backendCh = make(chan *ethclient.Client)
	pls.quit = make(chan bool)
	pls.peers = make(map[*Peer]struct{})
	pls.operatorNodeCh = make(chan *Peer, 1)

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
		PeerInfo: func(id discover.NodeID) interface{} {
			return id
		},
	}

	pls.blockchain = NewBlockChain(config)

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

	go pls.run()

	log.Info("[Plasma] node started", "version", ProtocolVersionStr, "operator", pls.config.OperatorAddress, "contract", pls.config.ContractAddress)
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
// XXX the life sycle of HandlePeer function is same as the p2p connection
func (pls *Plasma) HandlePeer(remote *p2p.Peer, rw p2p.MsgReadWriter) error {
	peer := newPeer(pls, remote, rw)

	if _, ok := pls.peers[peer]; ok {
		return nil
	}

	// registger peer
	pls.peerLock.Lock()
	pls.peers[peer] = struct{}{}
	pls.peerLock.Unlock()

	// unregister peer
	defer func() {
		pls.peerLock.Lock()
		delete(pls.peers, peer)
		pls.peerLock.Unlock()
	}()

	// pls handshake
	if err := peer.handshake(); err != nil {
		return err
	}

	// if peer is operator, record it
	peerPubkeyByte := peer.peer.ID()
	peerPubKey, _ := peerPubkeyByte.Pubkey()
	peerAddress := crypto.PubkeyToAddress(*peerPubKey)

	log.Info("[Plasma] peer connected", "localAddr", peer.peer.LocalAddr(), "remoteAddr", peer.peer.RemoteAddr(), "id", peer.peer.ID().String())

	if peerAddress == pls.config.OperatorAddress {
		pls.config.OperatorNode = peer

		var buffer bytes.Buffer

		buffer.WriteString("enode://")
		buffer.WriteString(peerPubkeyByte.String())
		buffer.WriteString("@")
		buffer.WriteString(peer.peer.RemoteAddr().String())

		pls.config.OperatorNodeURL = buffer.String()
		log.Info("[Plasma] operator node url", "url", pls.config.OperatorNodeURL)

		// TODO: should use channel?
		pls.operatorNodeCh <- peer
	}

	peer.start()
	defer peer.stop()

	return pls.handlePeer(peer)
}

// CurrentBlockNumber returns currnt block number + 1 on plasma chain
func (pls *Plasma) CurrentBlockNumber() uint64 {
	return pls.blockchain.getCurrentBlockNumber().Uint64()
}

func (pls *Plasma) getPeers() []*discover.Node {
	return nil
}

func (pls *Plasma) run() {
	select {
	case backend := <-pls.backendCh:
		pls.backend = backend
		log.Info("[Plasma] Ethereum jsonrpc backend attached")
	case <-pls.quit:
		return
	}

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

	// wait until plasma chain is synced
	if err := pls.waitPlsSynced(); err != nil {
		log.Warn("[Plasma] Failed to wait PLS syncing", "err", err)
		return err
	}

	// run deposit listener
	err = pls.listenDeposit()
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

	// Wait operator peer is conencted
	<-pls.operatorNodeCh

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

func (pls *Plasma) waitPlsSynced() error {
	// operator doesn't have to be synced.
	if pls.isOperator() {
		return nil
	}

	callOpts := bind.CallOpts{
		Pending: false,
		Context: pls.context,
	}

	localBlkNumBig := pls.blockchain.getCurrentBlockNumber()
	remoteBlkNumBig, err := pls.rootchain.CurrentChildBlock(&callOpts)

	if err != nil {
		return nil
	}

	log.Info("[Plasma] Checking plasma block", "localBlkNum", localBlkNumBig, "remoteBlkNum", remoteBlkNumBig)

	localBlkNum := localBlkNumBig.Uint64()
	remoteBlkNum := remoteBlkNumBig.Uint64()

	childBlockInterval := pls.blockchain.blockInterval.Uint64()
	var epochs, blocksToRequest []*big.Int

	for blkNum := localBlkNum - childBlockInterval; blkNum <= remoteBlkNum-childBlockInterval; blkNum += childBlockInterval {
		log.Info("[Plasma] adding epochs", "blkNum", blkNum)
		epochs = append(epochs, big.NewInt(int64(blkNum)))
	}

	for _, blkNum := range epochs {
		// submit-block
		if blkNum.Cmp(big0) > 0 {
			log.Info("[Plasma] add submit block request queue", "blkNum", blkNum)
			blocksToRequest = append(blocksToRequest, blkNum)
		}

		// deposit-block
		var i int64
		for i = 1; i < int64(childBlockInterval); i++ {
			depositBlockNumber := big0.Add(blkNum, big.NewInt(i))

			res, err := pls.rootchain.ChildChain(&callOpts, depositBlockNumber)
			log.Info("[Plasma] testing deposit block", "blkNum", depositBlockNumber)

			if err != nil || res.CreatedAt == nil {
				log.Info("[Plasma] stop", "err", err, "createdAt", res.CreatedAt, "root", res.Root)
				break
			}

			log.Info("[Plasma] add deposit block request queue", "blkNum", depositBlockNumber)
			blocksToRequest = append(blocksToRequest, depositBlockNumber)
		}
	}

	for _, blkNum := range blocksToRequest {
		query := getBlockData{
			Number: blkNum.Uint64(),
		}

		log.Info("[Plasma] requset block from operator", "blockNumber", blkNum)
		p2p.Send(pls.config.OperatorNode.rw, GetBlockCode, query)
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
	if pls.config.OperatorPrivateKey == nil {
		return false
	}

	return params.PlasmaOperatorAddress == crypto.PubkeyToAddress(pls.config.OperatorPrivateKey.PublicKey)
}

// watch Deposit event
// TODO: only operator need to listen it.
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

					if pls.isOperator() {
						// operator seal new deposit block
						if _, err := pls.blockchain.newDeposit(deposit.Amount, &deposit.Depositor, deposit.DepositBlock); err != nil {
							log.Warn("[Plasma] Failed to add new deposit from rootchain", "err", err)
						}
					} else {
						// other node request deposit block
						// TODO: use request pool. operator will send the block before expected request arrived.
						if pls.config.OperatorNode == nil {
							log.Warn("[Plasma] Operator Node is nil")
							continue
						}
						query := getBlockData{
							Number: pls.CurrentBlockNumber(),
						}

						p2p.Send(pls.config.OperatorNode.rw, GetBlockCode, query)
						log.Info("[Plasma] requesting new block", "number", query.Number)

						// packet, err := pls.config.OperatorNode.rw.ReadMsg()
						//
						// if err == nil {
						// 	log.Warn("[Plasma] Failed to fetch deposit block", "err", err)
						// 	continue
						// }
						//
						// if packet.Code != NewBlockCode {
						// 	log.Warn("[Plasma] Client expected to receive new deposit block.", "code", packet.Code)
						// 	continue
						// }
						//
						// var newBlockQuery newBlockData
						// if err := packet.Decode(&newBlockQuery); err != nil {
						// 	log.Warn("[Plasma] Failed to decode new block data", "err", err)
						// 	continue
						// }
						//
						// pls.blockchain.addBlock(newBlockQuery.Block)
						// log.Info("[Plasma] New deposit block fetched", "hash", newBlockQuery.Block.Hash())
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

	return pls.blockchain.addNewBlockListener(listener)
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

// handlePeer handles p2p message after handshake
func (pls *Plasma) handlePeer(peer *Peer) error {
	rw := peer.rw

	go func(peer *Peer) {
		for {
			time.Sleep(time.Second * 1)
			query := pingData{
				Number: 10,
			}

			if err := p2p.Send(peer.rw, PingCode, query); err != nil {
				log.Info("[Plasma] Failed to send pong message", "err", err)
			}
		}
	}(peer)

	for {
		packet, err := rw.ReadMsg()

		if err != nil {
			log.Warn("[Plasma] message loop", "peer", peer.ID(), "err", err)
			return err
		}

		log.Info("[Plasma] p2p message received", "code", packet.Code)

		switch packet.Code {
		case StatusCode:
			// this should not happen, but no need to panic; just ignore this message.
			log.Warn("unxepected status message received", "peer", peer.ID())
		case OperatorCode:
			if !pls.isOperator() && pls.config.OperatorNodeURL == "" {
				var query operatorData
				if err := packet.Decode(&query); err != nil {
					return errResp(ErrDecode, "%v: %v", packet, err)
				}

				pls.config.OperatorNodeURL = query.NodeURL
			}
		case GetBlockCode: // request block
			var query getBlockData

			if err := packet.Decode(&query); err != nil {
				return errResp(ErrDecode, "%v: %v", packet, err)
			}

			block, err := pls.blockchain.getBlock(big.NewInt(int64(query.Number)))

			if err != nil {
				return errResp(ErrDecode, "%v: %v", packet, err)
			}

			payload := newBlockData{
				Block: block,
			}

			err = p2p.Send(peer.rw, NewBlockCode, payload)

			if err != nil {
				log.Warn("[Plasma] Failed to send block", "blkNum", query.Number, "hash", query.Hash)
			}
		case NewBlockCode: // response block
			var payload newBlockData

			if err := packet.Decode(&payload); err != nil {
				return errResp(ErrDecode, "%v: %v", packet, err)
			}

			rawblock := payload.Block
			log.Info("[Plasma] new block received", "hash", rawblock.Hash(), "blkNum", rawblock.Data.BlockNumber)

			if err := pls.blockchain.addBlock(rawblock); err != nil {
				return errResp(ErrDecode, "%v: %v", packet, err)
			}

			log.Info("[Plasma] imported new block", "hash", rawblock.Hash(), "peer", peer.ID())
		case PingCode:
			var query pingData
			if err := packet.Decode(&query); err != nil {
				log.Info("[Plasma] Failed to decode ping", "err", err)
				return errResp(ErrDecode, "%v: %v", packet, err)
			}
			log.Info("[Plasma] ping received. send pong", "peer", peer.ID(), "query", query.Number)

			// get plasma block with number 1
			blk, _ := peer.host.blockchain.getBlock(big1)

			if blk == nil {
				log.Info("[Plasma] No block with number 1. Do not pong")
			} else {
				payload := pongData{
					Block: blk,
				}

				if err := p2p.Send(peer.rw, PongCode, payload); err != nil {
					log.Info("[Plasma] Failed to send pong message", "err", err)
					return errResp(ErrDecode, "%v: %v", packet, err)
				}
			}

		case PongCode:
			log.Info("[Plasma] pong received", "peer", peer.ID())

			var query pongData
			if err := packet.Decode(&query); err != nil {
				log.Info("[Plasma] Failed to decode pong", "err", err)
				return errResp(ErrDecode, "%v: %v", packet, err)
			}
			log.Info("[Plasma] pong received", "peer", peer.ID(), "blockNumber", query.Block.Data.BlockNumber, "hash", query.Block.Hash())

		}
	}
}

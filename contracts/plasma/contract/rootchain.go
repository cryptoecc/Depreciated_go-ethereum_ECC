// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// ByteUtilsABI is the input ABI used to generate the binding from.
const ByteUtilsABI = "[]"

// ByteUtilsBin is the compiled bytecode used for deploying new contracts.
const ByteUtilsBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820404e9bea48bd5ed6854c499932789d08a7721b793a5f3e96eb0feb351fc289de0029`

// DeployByteUtils deploys a new Ethereum contract, binding an instance of ByteUtils to it.
func DeployByteUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ByteUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(ByteUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ByteUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ByteUtils{ByteUtilsCaller: ByteUtilsCaller{contract: contract}, ByteUtilsTransactor: ByteUtilsTransactor{contract: contract}, ByteUtilsFilterer: ByteUtilsFilterer{contract: contract}}, nil
}

// ByteUtils is an auto generated Go binding around an Ethereum contract.
type ByteUtils struct {
	ByteUtilsCaller     // Read-only binding to the contract
	ByteUtilsTransactor // Write-only binding to the contract
	ByteUtilsFilterer   // Log filterer for contract events
}

// ByteUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ByteUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ByteUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ByteUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ByteUtilsSession struct {
	Contract     *ByteUtils        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ByteUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ByteUtilsCallerSession struct {
	Contract *ByteUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ByteUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ByteUtilsTransactorSession struct {
	Contract     *ByteUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ByteUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ByteUtilsRaw struct {
	Contract *ByteUtils // Generic contract binding to access the raw methods on
}

// ByteUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ByteUtilsCallerRaw struct {
	Contract *ByteUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ByteUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ByteUtilsTransactorRaw struct {
	Contract *ByteUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewByteUtils creates a new instance of ByteUtils, bound to a specific deployed contract.
func NewByteUtils(address common.Address, backend bind.ContractBackend) (*ByteUtils, error) {
	contract, err := bindByteUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ByteUtils{ByteUtilsCaller: ByteUtilsCaller{contract: contract}, ByteUtilsTransactor: ByteUtilsTransactor{contract: contract}, ByteUtilsFilterer: ByteUtilsFilterer{contract: contract}}, nil
}

// NewByteUtilsCaller creates a new read-only instance of ByteUtils, bound to a specific deployed contract.
func NewByteUtilsCaller(address common.Address, caller bind.ContractCaller) (*ByteUtilsCaller, error) {
	contract, err := bindByteUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ByteUtilsCaller{contract: contract}, nil
}

// NewByteUtilsTransactor creates a new write-only instance of ByteUtils, bound to a specific deployed contract.
func NewByteUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ByteUtilsTransactor, error) {
	contract, err := bindByteUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ByteUtilsTransactor{contract: contract}, nil
}

// NewByteUtilsFilterer creates a new log filterer instance of ByteUtils, bound to a specific deployed contract.
func NewByteUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ByteUtilsFilterer, error) {
	contract, err := bindByteUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ByteUtilsFilterer{contract: contract}, nil
}

// bindByteUtils binds a generic wrapper to an already deployed contract.
func bindByteUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ByteUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ByteUtils *ByteUtilsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ByteUtils.Contract.ByteUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ByteUtils *ByteUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ByteUtils.Contract.ByteUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ByteUtils *ByteUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ByteUtils.Contract.ByteUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ByteUtils *ByteUtilsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ByteUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ByteUtils *ByteUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ByteUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ByteUtils *ByteUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ByteUtils.Contract.contract.Transact(opts, method, params...)
}

// ECRecoveryABI is the input ABI used to generate the binding from.
const ECRecoveryABI = "[]"

// ECRecoveryBin is the compiled bytecode used for deploying new contracts.
const ECRecoveryBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582099916461c2b7f1ddb0d2c2c353c82025cfa13e0be4455ea58d10ea24c99fb9d90029`

// DeployECRecovery deploys a new Ethereum contract, binding an instance of ECRecovery to it.
func DeployECRecovery(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECRecovery, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ECRecoveryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}, ECRecoveryFilterer: ECRecoveryFilterer{contract: contract}}, nil
}

// ECRecovery is an auto generated Go binding around an Ethereum contract.
type ECRecovery struct {
	ECRecoveryCaller     // Read-only binding to the contract
	ECRecoveryTransactor // Write-only binding to the contract
	ECRecoveryFilterer   // Log filterer for contract events
}

// ECRecoveryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECRecoveryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECRecoveryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECRecoveryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoverySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECRecoverySession struct {
	Contract     *ECRecovery       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECRecoveryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECRecoveryCallerSession struct {
	Contract *ECRecoveryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ECRecoveryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECRecoveryTransactorSession struct {
	Contract     *ECRecoveryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ECRecoveryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ECRecoveryRaw struct {
	Contract *ECRecovery // Generic contract binding to access the raw methods on
}

// ECRecoveryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECRecoveryCallerRaw struct {
	Contract *ECRecoveryCaller // Generic read-only contract binding to access the raw methods on
}

// ECRecoveryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECRecoveryTransactorRaw struct {
	Contract *ECRecoveryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewECRecovery creates a new instance of ECRecovery, bound to a specific deployed contract.
func NewECRecovery(address common.Address, backend bind.ContractBackend) (*ECRecovery, error) {
	contract, err := bindECRecovery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}, ECRecoveryFilterer: ECRecoveryFilterer{contract: contract}}, nil
}

// NewECRecoveryCaller creates a new read-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryCaller(address common.Address, caller bind.ContractCaller) (*ECRecoveryCaller, error) {
	contract, err := bindECRecovery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryCaller{contract: contract}, nil
}

// NewECRecoveryTransactor creates a new write-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryTransactor(address common.Address, transactor bind.ContractTransactor) (*ECRecoveryTransactor, error) {
	contract, err := bindECRecovery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryTransactor{contract: contract}, nil
}

// NewECRecoveryFilterer creates a new log filterer instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryFilterer(address common.Address, filterer bind.ContractFilterer) (*ECRecoveryFilterer, error) {
	contract, err := bindECRecovery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryFilterer{contract: contract}, nil
}

// bindECRecovery binds a generic wrapper to an already deployed contract.
func bindECRecovery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecovery *ECRecoveryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecovery.Contract.ECRecoveryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecovery *ECRecoveryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecovery.Contract.ECRecoveryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecovery *ECRecoveryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecovery.Contract.ECRecoveryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECRecovery *ECRecoveryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ECRecovery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECRecovery *ECRecoveryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECRecovery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECRecovery *ECRecoveryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECRecovery.Contract.contract.Transact(opts, method, params...)
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
const MathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582093c6f0865f73208d3868408da47a0b254e78fda541b25402612a9fe3198b9e4a0029`

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// MerkleABI is the input ABI used to generate the binding from.
const MerkleABI = "[]"

// MerkleBin is the compiled bytecode used for deploying new contracts.
const MerkleBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058208f32eaf9099bad41eec69ce6a07a742c04a638d63dd3f50d92b4e18db98b2df70029`

// DeployMerkle deploys a new Ethereum contract, binding an instance of Merkle to it.
func DeployMerkle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Merkle, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerkleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Merkle{MerkleCaller: MerkleCaller{contract: contract}, MerkleTransactor: MerkleTransactor{contract: contract}, MerkleFilterer: MerkleFilterer{contract: contract}}, nil
}

// Merkle is an auto generated Go binding around an Ethereum contract.
type Merkle struct {
	MerkleCaller     // Read-only binding to the contract
	MerkleTransactor // Write-only binding to the contract
	MerkleFilterer   // Log filterer for contract events
}

// MerkleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleSession struct {
	Contract     *Merkle           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleCallerSession struct {
	Contract *MerkleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MerkleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleTransactorSession struct {
	Contract     *MerkleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleRaw struct {
	Contract *Merkle // Generic contract binding to access the raw methods on
}

// MerkleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleCallerRaw struct {
	Contract *MerkleCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleTransactorRaw struct {
	Contract *MerkleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkle creates a new instance of Merkle, bound to a specific deployed contract.
func NewMerkle(address common.Address, backend bind.ContractBackend) (*Merkle, error) {
	contract, err := bindMerkle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Merkle{MerkleCaller: MerkleCaller{contract: contract}, MerkleTransactor: MerkleTransactor{contract: contract}, MerkleFilterer: MerkleFilterer{contract: contract}}, nil
}

// NewMerkleCaller creates a new read-only instance of Merkle, bound to a specific deployed contract.
func NewMerkleCaller(address common.Address, caller bind.ContractCaller) (*MerkleCaller, error) {
	contract, err := bindMerkle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleCaller{contract: contract}, nil
}

// NewMerkleTransactor creates a new write-only instance of Merkle, bound to a specific deployed contract.
func NewMerkleTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTransactor, error) {
	contract, err := bindMerkle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTransactor{contract: contract}, nil
}

// NewMerkleFilterer creates a new log filterer instance of Merkle, bound to a specific deployed contract.
func NewMerkleFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleFilterer, error) {
	contract, err := bindMerkle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleFilterer{contract: contract}, nil
}

// bindMerkle binds a generic wrapper to an already deployed contract.
func bindMerkle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merkle *MerkleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Merkle.Contract.MerkleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merkle *MerkleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merkle.Contract.MerkleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merkle *MerkleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merkle.Contract.MerkleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Merkle *MerkleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Merkle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Merkle *MerkleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Merkle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Merkle *MerkleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Merkle.Contract.contract.Transact(opts, method, params...)
}

// PriorityQueueABI is the input ABI used to generate the binding from.
const PriorityQueueABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"minChild\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"k\",\"type\":\"uint256\"}],\"name\":\"insert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"delMin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// PriorityQueueBin is the compiled bytecode used for deploying new contracts.
const PriorityQueueBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a03191633178155604080516020810190915290815261003e9060019081610049565b5060006002556100b6565b828054828255906000526020600020908101928215610089579160200282015b82811115610089578251829060ff16905591602001919060010190610069565b50610095929150610099565b5090565b6100b391905b80821115610095576000815560010161009f565b90565b6105a4806100c56000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632dcdcd0c811461007157806390b5561d1461009b578063b07576ac146100b5578063bda1504b146100ca578063d6362e97146100df575b600080fd5b34801561007d57600080fd5b506100896004356100f4565b60408051918252519081900360200190f35b3480156100a757600080fd5b506100b36004356101c3565b005b3480156100c157600080fd5b5061008961023d565b3480156100d657600080fd5b506100896102f5565b3480156100eb57600080fd5b506100896102fb565b600060025461011e600161011260028661031c90919063ffffffff16565b9063ffffffff61035216565b111561013c5761013582600263ffffffff61031c16565b90506101be565b60016101538161011285600263ffffffff61031c16565b8154811061015d57fe5b600091825260209091200154600161017c84600263ffffffff61031c16565b8154811061018657fe5b906000526020600020015410156101a85761013582600263ffffffff61031c16565b610135600161011284600263ffffffff61031c16565b919050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146101e757600080fd5b60018054808201825560008290527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60182905560025461022c9163ffffffff61035216565b600281905561023a90610361565b50565b60008054819073ffffffffffffffffffffffffffffffffffffffff16331461026457600080fd5b600180548190811061027257fe5b90600052602060002001549050600160025481548110151561029057fe5b90600052602060002001546001808154811015156102aa57fe5b906000526020600020018190555060016002548154811015156102c957fe5b60009182526020822001556002546102e890600163ffffffff61046e16565b6002556101be6001610480565b60025481565b600060018081548110151561030c57fe5b9060005260206000200154905090565b60008083151561032f576000915061034b565b5082820282848281151561033f57fe5b041461034757fe5b8091505b5092915050565b60008282018381101561034757fe5b60005b600061037783600263ffffffff61056116565b111561046a57600161039083600263ffffffff61056116565b8154811061039a57fe5b90600052602060002001546001838154811015156103b457fe5b906000526020600020015410156104525760016103d883600263ffffffff61056116565b815481106103e257fe5b906000526020600020015490506001828154811015156103fe57fe5b600091825260209091200154600161041d84600263ffffffff61056116565b8154811061042757fe5b90600052602060002001819055508060018381548110151561044557fe5b6000918252602090912001555b61046382600263ffffffff61056116565b9150610364565b5050565b60008282111561047a57fe5b50900390565b6000805b60025461049b60028561031c90919063ffffffff16565b1161055c576104a9836100f4565b91506001828154811015156104ba57fe5b90600052602060002001546001848154811015156104d457fe5b906000526020600020015411156105545760018054849081106104f357fe5b9060005260206000200154905060018281548110151561050f57fe5b906000526020600020015460018481548110151561052957fe5b90600052602060002001819055508060018381548110151561054757fe5b6000918252602090912001555b819250610484565b505050565b600080828481151561056f57fe5b049493505050505600a165627a7a72305820425fce4bd38c396d9ceff70720d973c1c1a162925bbfd36f619accc618499b230029`

// DeployPriorityQueue deploys a new Ethereum contract, binding an instance of PriorityQueue to it.
func DeployPriorityQueue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PriorityQueue, error) {
	parsed, err := abi.JSON(strings.NewReader(PriorityQueueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PriorityQueueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PriorityQueue{PriorityQueueCaller: PriorityQueueCaller{contract: contract}, PriorityQueueTransactor: PriorityQueueTransactor{contract: contract}, PriorityQueueFilterer: PriorityQueueFilterer{contract: contract}}, nil
}

// PriorityQueue is an auto generated Go binding around an Ethereum contract.
type PriorityQueue struct {
	PriorityQueueCaller     // Read-only binding to the contract
	PriorityQueueTransactor // Write-only binding to the contract
	PriorityQueueFilterer   // Log filterer for contract events
}

// PriorityQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriorityQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriorityQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriorityQueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriorityQueueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PriorityQueueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriorityQueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PriorityQueueSession struct {
	Contract     *PriorityQueue    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PriorityQueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PriorityQueueCallerSession struct {
	Contract *PriorityQueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PriorityQueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PriorityQueueTransactorSession struct {
	Contract     *PriorityQueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PriorityQueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type PriorityQueueRaw struct {
	Contract *PriorityQueue // Generic contract binding to access the raw methods on
}

// PriorityQueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PriorityQueueCallerRaw struct {
	Contract *PriorityQueueCaller // Generic read-only contract binding to access the raw methods on
}

// PriorityQueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PriorityQueueTransactorRaw struct {
	Contract *PriorityQueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPriorityQueue creates a new instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueue(address common.Address, backend bind.ContractBackend) (*PriorityQueue, error) {
	contract, err := bindPriorityQueue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriorityQueue{PriorityQueueCaller: PriorityQueueCaller{contract: contract}, PriorityQueueTransactor: PriorityQueueTransactor{contract: contract}, PriorityQueueFilterer: PriorityQueueFilterer{contract: contract}}, nil
}

// NewPriorityQueueCaller creates a new read-only instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueueCaller(address common.Address, caller bind.ContractCaller) (*PriorityQueueCaller, error) {
	contract, err := bindPriorityQueue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PriorityQueueCaller{contract: contract}, nil
}

// NewPriorityQueueTransactor creates a new write-only instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*PriorityQueueTransactor, error) {
	contract, err := bindPriorityQueue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PriorityQueueTransactor{contract: contract}, nil
}

// NewPriorityQueueFilterer creates a new log filterer instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueueFilterer(address common.Address, filterer bind.ContractFilterer) (*PriorityQueueFilterer, error) {
	contract, err := bindPriorityQueue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PriorityQueueFilterer{contract: contract}, nil
}

// bindPriorityQueue binds a generic wrapper to an already deployed contract.
func bindPriorityQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriorityQueueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriorityQueue *PriorityQueueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriorityQueue.Contract.PriorityQueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriorityQueue *PriorityQueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriorityQueue.Contract.PriorityQueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriorityQueue *PriorityQueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriorityQueue.Contract.PriorityQueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PriorityQueue *PriorityQueueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PriorityQueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PriorityQueue *PriorityQueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriorityQueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PriorityQueue *PriorityQueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PriorityQueue.Contract.contract.Transact(opts, method, params...)
}

// CurrentSize is a free data retrieval call binding the contract method 0xbda1504b.
//
// Solidity: function currentSize() constant returns(uint256)
func (_PriorityQueue *PriorityQueueCaller) CurrentSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriorityQueue.contract.Call(opts, out, "currentSize")
	return *ret0, err
}

// CurrentSize is a free data retrieval call binding the contract method 0xbda1504b.
//
// Solidity: function currentSize() constant returns(uint256)
func (_PriorityQueue *PriorityQueueSession) CurrentSize() (*big.Int, error) {
	return _PriorityQueue.Contract.CurrentSize(&_PriorityQueue.CallOpts)
}

// CurrentSize is a free data retrieval call binding the contract method 0xbda1504b.
//
// Solidity: function currentSize() constant returns(uint256)
func (_PriorityQueue *PriorityQueueCallerSession) CurrentSize() (*big.Int, error) {
	return _PriorityQueue.Contract.CurrentSize(&_PriorityQueue.CallOpts)
}

// GetMin is a free data retrieval call binding the contract method 0xd6362e97.
//
// Solidity: function getMin() constant returns(uint256)
func (_PriorityQueue *PriorityQueueCaller) GetMin(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriorityQueue.contract.Call(opts, out, "getMin")
	return *ret0, err
}

// GetMin is a free data retrieval call binding the contract method 0xd6362e97.
//
// Solidity: function getMin() constant returns(uint256)
func (_PriorityQueue *PriorityQueueSession) GetMin() (*big.Int, error) {
	return _PriorityQueue.Contract.GetMin(&_PriorityQueue.CallOpts)
}

// GetMin is a free data retrieval call binding the contract method 0xd6362e97.
//
// Solidity: function getMin() constant returns(uint256)
func (_PriorityQueue *PriorityQueueCallerSession) GetMin() (*big.Int, error) {
	return _PriorityQueue.Contract.GetMin(&_PriorityQueue.CallOpts)
}

// MinChild is a free data retrieval call binding the contract method 0x2dcdcd0c.
//
// Solidity: function minChild(i uint256) constant returns(uint256)
func (_PriorityQueue *PriorityQueueCaller) MinChild(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PriorityQueue.contract.Call(opts, out, "minChild", i)
	return *ret0, err
}

// MinChild is a free data retrieval call binding the contract method 0x2dcdcd0c.
//
// Solidity: function minChild(i uint256) constant returns(uint256)
func (_PriorityQueue *PriorityQueueSession) MinChild(i *big.Int) (*big.Int, error) {
	return _PriorityQueue.Contract.MinChild(&_PriorityQueue.CallOpts, i)
}

// MinChild is a free data retrieval call binding the contract method 0x2dcdcd0c.
//
// Solidity: function minChild(i uint256) constant returns(uint256)
func (_PriorityQueue *PriorityQueueCallerSession) MinChild(i *big.Int) (*big.Int, error) {
	return _PriorityQueue.Contract.MinChild(&_PriorityQueue.CallOpts, i)
}

// DelMin is a paid mutator transaction binding the contract method 0xb07576ac.
//
// Solidity: function delMin() returns(uint256)
func (_PriorityQueue *PriorityQueueTransactor) DelMin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PriorityQueue.contract.Transact(opts, "delMin")
}

// DelMin is a paid mutator transaction binding the contract method 0xb07576ac.
//
// Solidity: function delMin() returns(uint256)
func (_PriorityQueue *PriorityQueueSession) DelMin() (*types.Transaction, error) {
	return _PriorityQueue.Contract.DelMin(&_PriorityQueue.TransactOpts)
}

// DelMin is a paid mutator transaction binding the contract method 0xb07576ac.
//
// Solidity: function delMin() returns(uint256)
func (_PriorityQueue *PriorityQueueTransactorSession) DelMin() (*types.Transaction, error) {
	return _PriorityQueue.Contract.DelMin(&_PriorityQueue.TransactOpts)
}

// Insert is a paid mutator transaction binding the contract method 0x90b5561d.
//
// Solidity: function insert(k uint256) returns()
func (_PriorityQueue *PriorityQueueTransactor) Insert(opts *bind.TransactOpts, k *big.Int) (*types.Transaction, error) {
	return _PriorityQueue.contract.Transact(opts, "insert", k)
}

// Insert is a paid mutator transaction binding the contract method 0x90b5561d.
//
// Solidity: function insert(k uint256) returns()
func (_PriorityQueue *PriorityQueueSession) Insert(k *big.Int) (*types.Transaction, error) {
	return _PriorityQueue.Contract.Insert(&_PriorityQueue.TransactOpts, k)
}

// Insert is a paid mutator transaction binding the contract method 0x90b5561d.
//
// Solidity: function insert(k uint256) returns()
func (_PriorityQueue *PriorityQueueTransactorSession) Insert(k *big.Int) (*types.Transaction, error) {
	return _PriorityQueue.Contract.Insert(&_PriorityQueue.TransactOpts, k)
}

// RLPABI is the input ABI used to generate the binding from.
const RLPABI = "[]"

// RLPBin is the compiled bytecode used for deploying new contracts.
const RLPBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582067e8a9f5cc7866a98f3d922e324c9321111dde60053dd7e16210d6b3ef66ebfd0029`

// DeployRLP deploys a new Ethereum contract, binding an instance of RLP to it.
func DeployRLP(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLP, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RLPBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLP{RLPCaller: RLPCaller{contract: contract}, RLPTransactor: RLPTransactor{contract: contract}, RLPFilterer: RLPFilterer{contract: contract}}, nil
}

// RLP is an auto generated Go binding around an Ethereum contract.
type RLP struct {
	RLPCaller     // Read-only binding to the contract
	RLPTransactor // Write-only binding to the contract
	RLPFilterer   // Log filterer for contract events
}

// RLPCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPSession struct {
	Contract     *RLP              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPCallerSession struct {
	Contract *RLPCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RLPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPTransactorSession struct {
	Contract     *RLPTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPRaw struct {
	Contract *RLP // Generic contract binding to access the raw methods on
}

// RLPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPCallerRaw struct {
	Contract *RLPCaller // Generic read-only contract binding to access the raw methods on
}

// RLPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPTransactorRaw struct {
	Contract *RLPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLP creates a new instance of RLP, bound to a specific deployed contract.
func NewRLP(address common.Address, backend bind.ContractBackend) (*RLP, error) {
	contract, err := bindRLP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLP{RLPCaller: RLPCaller{contract: contract}, RLPTransactor: RLPTransactor{contract: contract}, RLPFilterer: RLPFilterer{contract: contract}}, nil
}

// NewRLPCaller creates a new read-only instance of RLP, bound to a specific deployed contract.
func NewRLPCaller(address common.Address, caller bind.ContractCaller) (*RLPCaller, error) {
	contract, err := bindRLP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPCaller{contract: contract}, nil
}

// NewRLPTransactor creates a new write-only instance of RLP, bound to a specific deployed contract.
func NewRLPTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPTransactor, error) {
	contract, err := bindRLP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPTransactor{contract: contract}, nil
}

// NewRLPFilterer creates a new log filterer instance of RLP, bound to a specific deployed contract.
func NewRLPFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPFilterer, error) {
	contract, err := bindRLP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPFilterer{contract: contract}, nil
}

// bindRLP binds a generic wrapper to an already deployed contract.
func bindRLP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLP *RLPRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLP.Contract.RLPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLP *RLPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLP.Contract.RLPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLP *RLPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLP.Contract.RLPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLP *RLPCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLP *RLPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLP *RLPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLP.Contract.contract.Transact(opts, method, params...)
}

// RLPEncodeABI is the input ABI used to generate the binding from.
const RLPEncodeABI = "[]"

// RLPEncodeBin is the compiled bytecode used for deploying new contracts.
const RLPEncodeBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a7230582015d1a0ed9a6f6607caa58325f14da18c8d93566ccf4d068d4dfe9eb33c40d4d60029`

// DeployRLPEncode deploys a new Ethereum contract, binding an instance of RLPEncode to it.
func DeployRLPEncode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RLPEncode, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPEncodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RLPEncodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RLPEncode{RLPEncodeCaller: RLPEncodeCaller{contract: contract}, RLPEncodeTransactor: RLPEncodeTransactor{contract: contract}, RLPEncodeFilterer: RLPEncodeFilterer{contract: contract}}, nil
}

// RLPEncode is an auto generated Go binding around an Ethereum contract.
type RLPEncode struct {
	RLPEncodeCaller     // Read-only binding to the contract
	RLPEncodeTransactor // Write-only binding to the contract
	RLPEncodeFilterer   // Log filterer for contract events
}

// RLPEncodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPEncodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPEncodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RLPEncodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPEncodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RLPEncodeSession struct {
	Contract     *RLPEncode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RLPEncodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RLPEncodeCallerSession struct {
	Contract *RLPEncodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RLPEncodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RLPEncodeTransactorSession struct {
	Contract     *RLPEncodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RLPEncodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type RLPEncodeRaw struct {
	Contract *RLPEncode // Generic contract binding to access the raw methods on
}

// RLPEncodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RLPEncodeCallerRaw struct {
	Contract *RLPEncodeCaller // Generic read-only contract binding to access the raw methods on
}

// RLPEncodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RLPEncodeTransactorRaw struct {
	Contract *RLPEncodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRLPEncode creates a new instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncode(address common.Address, backend bind.ContractBackend) (*RLPEncode, error) {
	contract, err := bindRLPEncode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLPEncode{RLPEncodeCaller: RLPEncodeCaller{contract: contract}, RLPEncodeTransactor: RLPEncodeTransactor{contract: contract}, RLPEncodeFilterer: RLPEncodeFilterer{contract: contract}}, nil
}

// NewRLPEncodeCaller creates a new read-only instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeCaller(address common.Address, caller bind.ContractCaller) (*RLPEncodeCaller, error) {
	contract, err := bindRLPEncode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeCaller{contract: contract}, nil
}

// NewRLPEncodeTransactor creates a new write-only instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPEncodeTransactor, error) {
	contract, err := bindRLPEncode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeTransactor{contract: contract}, nil
}

// NewRLPEncodeFilterer creates a new log filterer instance of RLPEncode, bound to a specific deployed contract.
func NewRLPEncodeFilterer(address common.Address, filterer bind.ContractFilterer) (*RLPEncodeFilterer, error) {
	contract, err := bindRLPEncode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RLPEncodeFilterer{contract: contract}, nil
}

// bindRLPEncode binds a generic wrapper to an already deployed contract.
func bindRLPEncode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPEncodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPEncode *RLPEncodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLPEncode.Contract.RLPEncodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPEncode *RLPEncodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPEncode.Contract.RLPEncodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPEncode *RLPEncodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPEncode.Contract.RLPEncodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RLPEncode *RLPEncodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RLPEncode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RLPEncode *RLPEncodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RLPEncode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RLPEncode *RLPEncodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RLPEncode.Contract.contract.Transact(opts, method, params...)
}

// RootChainABI is the input ABI used to generate the binding from.
const RootChainABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"nextWeekOldChildBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exitIds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"utxoPos\",\"type\":\"uint256\"},{\"name\":\"txBytes\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"sigs\",\"type\":\"bytes\"}],\"name\":\"startExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cUtxoPos\",\"type\":\"uint256\"},{\"name\":\"eUtxoPos\",\"type\":\"uint256\"},{\"name\":\"txBytes\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"sigs\",\"type\":\"bytes\"},{\"name\":\"confirmationSig\",\"type\":\"bytes\"}],\"name\":\"challengeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exits\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"utxoPos\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"childBlockInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weekOldBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"recentBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentChildBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getChildChain\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentDepositBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"submitBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDepositBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeExits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"priority\",\"type\":\"uint256\"}],\"name\":\"getExit\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"childChain\",\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"created_at\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"depositBlock\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"submitBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"Submit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"exitor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"utxoPos\",\"type\":\"uint256\"}],\"name\":\"Exit\",\"type\":\"event\"}]"

// RootChainBin is the compiled bytecode used for deploying new contracts.
const RootChainBin = `0x608060405234801561001057600080fd5b5060048054600160a060020a031916331790556103e8600981905560055560016006819055600855610040610082565b604051809103906000f08015801561005c573d6000803e3d6000fd5b5060038054600160a060020a031916600160a060020a0392909216919091179055610093565b60405161066980620021c983390190565b61212680620000a36000396000f3006080604052600436106100e25763ffffffff60e060020a600035041663107e58f181146100e757806316b38840146101115780631c91a6b91461012957806332773ba314610205578063342de1791461031f57806338a9e0bc1461035f5780634237b5f3146103745780636f84b695146103895780637a95f1e81461039e57806385444de3146103b3578063a98c7f2c146103e4578063baa47694146103f9578063bcd5926114610411578063bf7e214f14610426578063c6ab44cd14610457578063d0e30db01461046c578063e60f1ff114610474578063f95643b11461048c575b600080fd5b3480156100f357600080fd5b506100ff6004356104a4565b60408051918252519081900360200190f35b34801561011d57600080fd5b506100ff6004356104e1565b34801561013557600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261020395833595369560449491939091019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506104f39650505050505050565b005b34801561021157600080fd5b50604080516020600460443581810135601f810184900484028501840190955284845261020394823594602480359536959460649492019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a99988101979196509182019450925082915084018382808284375094975061092b9650505050505050565b34801561032b57600080fd5b50610337600435610ae4565b60408051600160a060020a039094168452602084019290925282820152519081900360600190f35b34801561036b57600080fd5b506100ff610b10565b34801561038057600080fd5b506100ff610b16565b34801561039557600080fd5b506100ff610b1c565b3480156103aa57600080fd5b506100ff610b22565b3480156103bf57600080fd5b506103cb600435610b28565b6040805192835260208301919091528051918290030190f35b3480156103f057600080fd5b506100ff610b42565b34801561040557600080fd5b50610203600435610b48565b34801561041d57600080fd5b506100ff610c91565b34801561043257600080fd5b5061043b610cb5565b60408051600160a060020a039092168252519081900360200190f35b34801561046357600080fd5b506100ff610cc4565b6102036110f0565b34801561048057600080fd5b506103376004356112ca565b34801561049857600080fd5b506103cb6004356112f8565b6009546000906104db906104cf60016104c3868463ffffffff61131116565b9063ffffffff61132d16565b9063ffffffff61134316565b92915050565b60026020526000908152604090205481565b606060008060008060008060005b6105144262093a8063ffffffff61136e16565b60085460009081526020819052604090206001015410156105a757600854600090815260208190526040902060010154151561058a576000806105586008546104a4565b81526020019081526020016000206001015460001415610577576105a7565b6105826008546104a4565b6008556105a2565b60085461059e90600163ffffffff61132d16565b6008555b610501565b6105c1600b6105b58d611380565b9063ffffffff6113d316565b9750633b9aca00808d049750612710908d066000898152602081905260409020548a519290910497506127108802633b9aca008a028f0303965094506106229089906006600289020190811061061357fe5b9060200190602002015161147b565b600160a060020a0316331461063657600080fd5b8a6040518082805190602001908083835b602083106106665780518252601f199092019160209182019101610647565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390209250826106a38a600060826114c7565b6040518281528151602080830191908401908083835b602083106106d85780518252601f1990920191602091820191016106b9565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405180910390209150610746838561072e8b600081518110151561071f57fe5b9060200190602002015161150e565b6107408c600381518110151561071f57fe5b8d611556565b151561075157600080fd5b6107638287868d63ffffffff6116a516565b151561076e57600080fd5b60085487101561079d57610796600854888e81151561078957fe5b049063ffffffff61134316565b90506107a0565b508a5b60008c815260026020526040902054156107b957600080fd5b60008c81526002602052604080822083905560035481517f90b5561d000000000000000000000000000000000000000000000000000000008152600481018590529151600160a060020a03909116926390b5561d926024808201939182900301818387803b15801561082a57600080fd5b505af115801561083e573d6000803e3d6000fd5b505050506060604051908101604052806108658a8860020260060181518110151561061357fe5b600160a060020a0316815260200161088a8a8860020260070181518110151561071f57fe5b815260209081018e9052600083815260018083526040918290208451815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0390911617815584840151918101919091559281015160029093019290925581513381529081018e905281517f22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631929181900390910190a1505050505050505050505050565b633b9aca0080870460009081526020818152604080832054898452600283528184205491518951612710968d069690960495919492938392839283928d9282918401908083835b602083106109915780518252601f199092019160209182019101610972565b51815160001960209485036101000a01908116901991909116179052604080519490920184900384208085528482018d905282519485900390920184208285528f51929a5098508995508e945083810192508401908083835b60208310610a095780518252601f1990920191602091820191016109ea565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912060008c815260019092529290205491965050600160a060020a03169350610a5e92508591508a9050611731565b600160a060020a03828116911614610a7557600080fd5b610a878288888d63ffffffff6116a516565b1515610a9257600080fd5b505050600091825250600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff191681559283018490556002928301849055998352529687209690965550505050505050565b6001602081905260009182526040909120805491810154600290910154600160a060020a039092169183565b60095481565b60085481565b60075481565b60055481565b600090815260208190526040902080546001909101549091565b60065481565b600454600160a060020a03163314610b5f57600080fd5b610b724262093a8063ffffffff61136e16565b6008546000908152602081905260409020600101541015610c05576008546000908152602081905260409020600101541515610be857600080610bb66008546104a4565b81526020019081526020016000206001015460001415610bd557610c05565b610be06008546104a4565b600855610c00565b600854610bfc90600163ffffffff61132d16565b6008555b610b5f565b60408051808201825282815242602080830191825260058054600090815280835285902093518455915160019093019290925554825190815290810183905281517f03e03b35d08c09e37050f76f1a914a334810af275603f96ed8ad6f2cc3617eba929181900390910190a1600954600554610c869163ffffffff61132d16565b600555506001600655565b6000610cb06006546104c360095460055461136e90919063ffffffff16565b905090565b600454600160a060020a031681565b600080610ccf612097565b6000805b610ce64262093a8063ffffffff61136e16565b6008546000908152602081905260409020600101541015610d79576008546000908152602081905260409020600101541515610d5c57600080610d2a6008546104a4565b81526020019081526020016000206001015460001415610d4957610d79565b610d546008546104a4565b600855610d74565b600854610d7090600163ffffffff61132d16565b6008555b610cd3565b610d8c426212750063ffffffff61136e16565b935060016000600360009054906101000a9004600160a060020a0316600160a060020a031663d6362e976040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610de557600080fd5b505af1158015610df9573d6000803e3d6000fd5b505050506040513d6020811015610e0f57600080fd5b50518152602081810192909252604090810160002081516060810183528154600160a060020a0316815260018201549381019390935260020154908201819052909350610e6690633b9aca0063ffffffff61131116565b91505b60008281526020819052604090206001015484118015610f115750600354604080517fbda1504b0000000000000000000000000000000000000000000000000000000081529051600092600160a060020a03169163bda1504b91600480830192602092919082900301818787803b158015610ee357600080fd5b505af1158015610ef7573d6000803e3d6000fd5b505050506040513d6020811015610f0d57600080fd5b5051115b156110e95782516020840151604051600160a060020a039092169181156108fc0291906000818181858888f19350505050158015610f53573d6000803e3d6000fd5b50600360009054906101000a9004600160a060020a0316600160a060020a031663b07576ac6040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610fa757600080fd5b505af1158015610fbb573d6000803e3d6000fd5b505050506040513d6020811015610fd157600080fd5b50516000818152600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff191681558084018590556002908101859055888201518552825280842084905560035481517fd6362e9700000000000000000000000000000000000000000000000000000000815291519596509294600160a060020a039093169263d6362e979260048084019391929182900301818787803b15801561107a57600080fd5b505af115801561108e573d6000803e3d6000fd5b505050506040513d60208110156110a457600080fd5b50518152602081810192909252604090810160002081516060810183528154600160a060020a0316815260018201549381019390935260020154908201529250610e69565b5050505090565b60006060600080600060095460065410151561110b57600080fd5b611113610c91565b945061111e85611806565b9350836040518082805190602001908083835b602083106111505780518252601f199092019160209182019101611131565b51815160209384036101000a600019018019909216911617905260408051929094018290038220608280845260c0840190955295509093508301905061104080388339019050506040518281528151602080830191908401908083835b602083106111cc5780518252601f1990920191602091820191016111ad565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405180910390209150600090505b6010811015611241576040805192835260208084018590528151938490038201842085855290840194909452805192839003019091209190600101611201565b604080518082018252838152426020808301918252600089815290819052929092209051815590516001918201556006546112819163ffffffff61132d16565b6006556040805133815234602082015280820187905290517f90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a159181900360600190a15050505050565b6000818152600160208190526040909120805491810154600290910154600160a060020a0390921693909250565b6000602081905290815260409020805460019091015482565b600080828481151561131f57fe5b0490508091505b5092915050565b60008282018381101561133c57fe5b9392505050565b6000808315156113565760009150611326565b5082820282848281151561136657fe5b041461133c57fe5b60008282111561137a57fe5b50900390565b6113886120c2565b815160008115156113ae57604080518082019091526000808252602082015292506113cc565b60208401905060408051908101604052808281526020018381525092505b5050919050565b60606113dd6120d9565b60006113e8856119d1565b15156113f357600080fd5b8360405190808252806020026020018201604052801561142d57816020015b61141a6120c2565b8152602001906001900390816114125790505b50925061143985611a03565b91505b61144582611a3c565b156114735761145382611a5b565b838281518110151561146157fe5b6020908102909101015260010161143c565b505092915050565b600080600061148984611a9d565b151561149457600080fd5b61149d84611ac6565b9092509050601481146114af57600080fd5b50516c01000000000000000000000000900492915050565b604051606090601f831680820180850187830187015b818310156114f55780518352602092830192016114dd565b5050848352601f01601f19166040525090509392505050565b600080600061151c84611a9d565b151561152757600080fd5b61153084611ac6565b91509150602081111561154257600080fd5b806020036101000a82510492505050919050565b60006060806060600080600060606041895181151561157157fe5b061580156115825750610104895111155b151561158d57600080fd5b61159a89600060416114c7565b96506115a8896041806114c7565b95506115b789608260416114c7565b604080518f8152602081018f9052815190819003909101902090955093508a1580156115e1575089155b1561160b576115f08486611731565b600160a060020a031633600160a060020a0316149750611695565b600192506001915060008b1115611646576116268486611731565b600160a060020a03166116398e89611731565b600160a060020a03161492505b60008a11156116885761165c8960c360416114c7565b90506116688482611731565b600160a060020a031661167b8e88611731565b600160a060020a03161491505b8280156116925750815b97505b5050505050505095945050505050565b60008060008084516102001415156116bc57600080fd5b5086905060205b6102008111611723578481015192506002870615156116fa5760408051928352602083018490528051928390030190912090611715565b60408051848152602081019390935280519283900301909120905b6002870496506020016116c3565b509390931495945050505050565b6000806000808451604114151561174b57600093506117fd565b50505060208201516040830151606084015160001a601b60ff8216101561177057601b015b8060ff16601b1415801561178857508060ff16601c14155b1561179657600093506117fd565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af11580156117f0573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b60408051600b8082526101808201909252606091600091839182918291816020015b606081526020019060019003908161182857505060408051601480825281830190925291945060208201610280803883395050604080516000815260208101909152919350611878919050611b43565b9050600093505b81518410156118bc57815160009083908690811061189957fe5b906020010190600160f860020a031916908160001a90535060019093019261187f565b6118c582611b43565b91506118d86118d387611baf565b611b43565b8360008151811015156118e757fe5b60209081029091010152600193505b60068410156119235780838581518110151561190e57fe5b602090810290910101526001909301926118f6565b61192f6118d333611baf565b83600681518110151561193e57fe5b602090810290910101526119546118d334611baf565b83600781518110151561196357fe5b60209081029091010152825182908490600890811061197e57fe5b60209081029091010152825181908490600990811061199957fe5b60209081029091010152825181908490600a9081106119b457fe5b602090810290910101526119c783611c71565b9695505050505050565b6000808260200151600014156119ea57600091506119fd565b8260000151905060c0815160001a101591505b50919050565b611a0b6120d9565b6000611a16836119d1565b1515611a2157600080fd5b611a2a83611c96565b83519383529092016020820152919050565b6000611a466120c2565b50508051602080820151915192015191011190565b611a636120c2565b600080611a6f84611a3c565b156100e25783602001519150611a8482611d16565b82845260208085018290528382019086015290506113cc565b600080826020015160001415611ab657600091506119fd565b5050515160c060009190911a1090565b6000806000806000611ad786611a9d565b1515611ae257600080fd5b8551805160001a935091506080831015611b025781945060019350611b3b565b60b8831015611b205760018660200151039350816001019450611b3b565b60b78303905080600187602001510303935080820160010194505b505050915091565b60608082516001148015611b7657506080836000815181101515611b6357fe5b016020015160f860020a90819004810204105b15611ba2576040805160018082528183019092529060208083019080388339508593506104db92505050565b61133c83608060b7611da3565b60606000805b602082108015611bde5750838260208110611bcc57fe5b1a60f860020a0260f860020a90046000145b15611bee57816001019150611bb5565b816020036040519080825280601f01601f191660200182016040528015611c1f578160200160208202803883390190505b5092505b60208210156113cc57838260208110611c3857fe5b1a60f860020a028382815181101515611c4d57fe5b906020010190600160f860020a031916908160001a90535060019182019101611c23565b6060806060611c7f84611f44565b9150611c8e8260c060f7611da3565b949350505050565b6000806000836020015160001415611cb157600092506113cc565b50508151805160001a906080821015611ccd57600092506113cc565b60b8821080611ce8575060c08210158015611ce8575060f882105b15611cf657600192506113cc565b60c0821015611d0b5760b519820192506113cc565b5060f5190192915050565b8051600090811a6080811015611d2f57600191506119fd565b60b8811015611d4457607e19810191506119fd565b60c0811015611d6d57600183015160b76020839003016101000a9004810160b5190191506119fd565b60f8811015611d825760be19810191506119fd565b6001929092015160f76020849003016101000a900490910160f51901919050565b8251606090602085019082906000908160015b8083811515611dc157fe5b0415611dd65760019091019061010002611db6565b60378311611e5b57826001016040519080825280601f01601f191660200182016040528015611e0f578160200160208202803883390190505b509450828960ff160160f860020a02856000815181101515611e2d57fe5b906020010190600160f860020a031916908160001a905350602185019350611e56848785612053565b611f36565b8282600101016040519080825280601f01601f191660200182016040528015611e8e578160200160208202803883390190505b509450818860ff160160f860020a02856000815181101515611eac57fe5b906020010190600160f860020a031916908160001a905350600190505b818111611f23576101008183036101000a84811515611ee457fe5b04811515611eee57fe5b0660f860020a028582815181101515611f0357fe5b906020010190600160f860020a031916908160001a905350600101611ec9565b8160218601019350611f36848785612053565b509298975050505050505050565b60606000806060600060606000875160001415611f71576040805160008152602081019091529650612048565b600094505b8751851015611fa9578785815181101515611f8d57fe5b9060200190602002015151860195508480600101955050611f76565b856040519080825280601f01601f191660200182016040528015611fd7578160200160208202803883390190505b509350602084019250600094505b8751851015612044578785815181101515611ffc57fe5b90602001906020020151915060208201905061201a83828451612053565b878581518110151561202857fe5b9060200190602002015151830192508480600101955050611fe5565b8396505b505050505050919050565b60005b60208210612078578251845260209384019390920191601f1990910190612056565b50905182516020929092036101000a6000190180199091169116179052565b6060604051908101604052806000600160a060020a0316815260200160008152602001600081525090565b604080518082019091526000808252602082015290565b6060604051908101604052806120ed6120c2565b81526020016000815250905600a165627a7a72305820a49810acc81b98d23b1554f4125adf71f2af8dedc8c4ba21cfa3640dca05e9b30029608060405234801561001057600080fd5b5060008054600160a060020a03191633178155604080516020810190915290815261003e9060019081610049565b5060006002556100b6565b828054828255906000526020600020908101928215610089579160200282015b82811115610089578251829060ff16905591602001919060010190610069565b50610095929150610099565b5090565b6100b391905b80821115610095576000815560010161009f565b90565b6105a4806100c56000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632dcdcd0c811461007157806390b5561d1461009b578063b07576ac146100b5578063bda1504b146100ca578063d6362e97146100df575b600080fd5b34801561007d57600080fd5b506100896004356100f4565b60408051918252519081900360200190f35b3480156100a757600080fd5b506100b36004356101c3565b005b3480156100c157600080fd5b5061008961023d565b3480156100d657600080fd5b506100896102f5565b3480156100eb57600080fd5b506100896102fb565b600060025461011e600161011260028661031c90919063ffffffff16565b9063ffffffff61035216565b111561013c5761013582600263ffffffff61031c16565b90506101be565b60016101538161011285600263ffffffff61031c16565b8154811061015d57fe5b600091825260209091200154600161017c84600263ffffffff61031c16565b8154811061018657fe5b906000526020600020015410156101a85761013582600263ffffffff61031c16565b610135600161011284600263ffffffff61031c16565b919050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146101e757600080fd5b60018054808201825560008290527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60182905560025461022c9163ffffffff61035216565b600281905561023a90610361565b50565b60008054819073ffffffffffffffffffffffffffffffffffffffff16331461026457600080fd5b600180548190811061027257fe5b90600052602060002001549050600160025481548110151561029057fe5b90600052602060002001546001808154811015156102aa57fe5b906000526020600020018190555060016002548154811015156102c957fe5b60009182526020822001556002546102e890600163ffffffff61046e16565b6002556101be6001610480565b60025481565b600060018081548110151561030c57fe5b9060005260206000200154905090565b60008083151561032f576000915061034b565b5082820282848281151561033f57fe5b041461034757fe5b8091505b5092915050565b60008282018381101561034757fe5b60005b600061037783600263ffffffff61056116565b111561046a57600161039083600263ffffffff61056116565b8154811061039a57fe5b90600052602060002001546001838154811015156103b457fe5b906000526020600020015410156104525760016103d883600263ffffffff61056116565b815481106103e257fe5b906000526020600020015490506001828154811015156103fe57fe5b600091825260209091200154600161041d84600263ffffffff61056116565b8154811061042757fe5b90600052602060002001819055508060018381548110151561044557fe5b6000918252602090912001555b61046382600263ffffffff61056116565b9150610364565b5050565b60008282111561047a57fe5b50900390565b6000805b60025461049b60028561031c90919063ffffffff16565b1161055c576104a9836100f4565b91506001828154811015156104ba57fe5b90600052602060002001546001848154811015156104d457fe5b906000526020600020015411156105545760018054849081106104f357fe5b9060005260206000200154905060018281548110151561050f57fe5b906000526020600020015460018481548110151561052957fe5b90600052602060002001819055508060018381548110151561054757fe5b6000918252602090912001555b819250610484565b505050565b600080828481151561056f57fe5b049493505050505600a165627a7a72305820425fce4bd38c396d9ceff70720d973c1c1a162925bbfd36f619accc618499b230029`

// DeployRootChain deploys a new Ethereum contract, binding an instance of RootChain to it.
func DeployRootChain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RootChain, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootChainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}, RootChainFilterer: RootChainFilterer{contract: contract}}, nil
}

// RootChain is an auto generated Go binding around an Ethereum contract.
type RootChain struct {
	RootChainCaller     // Read-only binding to the contract
	RootChainTransactor // Write-only binding to the contract
	RootChainFilterer   // Log filterer for contract events
}

// RootChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainSession struct {
	Contract     *RootChain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainCallerSession struct {
	Contract *RootChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RootChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainTransactorSession struct {
	Contract     *RootChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RootChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainRaw struct {
	Contract *RootChain // Generic contract binding to access the raw methods on
}

// RootChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainCallerRaw struct {
	Contract *RootChainCaller // Generic read-only contract binding to access the raw methods on
}

// RootChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainTransactorRaw struct {
	Contract *RootChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChain creates a new instance of RootChain, bound to a specific deployed contract.
func NewRootChain(address common.Address, backend bind.ContractBackend) (*RootChain, error) {
	contract, err := bindRootChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}, RootChainFilterer: RootChainFilterer{contract: contract}}, nil
}

// NewRootChainCaller creates a new read-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainCaller(address common.Address, caller bind.ContractCaller) (*RootChainCaller, error) {
	contract, err := bindRootChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainCaller{contract: contract}, nil
}

// NewRootChainTransactor creates a new write-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainTransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainTransactor, error) {
	contract, err := bindRootChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainTransactor{contract: contract}, nil
}

// NewRootChainFilterer creates a new log filterer instance of RootChain, bound to a specific deployed contract.
func NewRootChainFilterer(address common.Address, filterer bind.ContractFilterer) (*RootChainFilterer, error) {
	contract, err := bindRootChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootChainFilterer{contract: contract}, nil
}

// bindRootChain binds a generic wrapper to an already deployed contract.
func bindRootChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.RootChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_RootChain *RootChainCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_RootChain *RootChainSession) Authority() (common.Address, error) {
	return _RootChain.Contract.Authority(&_RootChain.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_RootChain *RootChainCallerSession) Authority() (common.Address, error) {
	return _RootChain.Contract.Authority(&_RootChain.CallOpts)
}

// ChildBlockInterval is a free data retrieval call binding the contract method 0x38a9e0bc.
//
// Solidity: function childBlockInterval() constant returns(uint256)
func (_RootChain *RootChainCaller) ChildBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "childBlockInterval")
	return *ret0, err
}

// ChildBlockInterval is a free data retrieval call binding the contract method 0x38a9e0bc.
//
// Solidity: function childBlockInterval() constant returns(uint256)
func (_RootChain *RootChainSession) ChildBlockInterval() (*big.Int, error) {
	return _RootChain.Contract.ChildBlockInterval(&_RootChain.CallOpts)
}

// ChildBlockInterval is a free data retrieval call binding the contract method 0x38a9e0bc.
//
// Solidity: function childBlockInterval() constant returns(uint256)
func (_RootChain *RootChainCallerSession) ChildBlockInterval() (*big.Int, error) {
	return _RootChain.Contract.ChildBlockInterval(&_RootChain.CallOpts)
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, created_at uint256)
func (_RootChain *RootChainCaller) ChildChain(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Root      [32]byte
	CreatedAt *big.Int
}, error) {
	ret := new(struct {
		Root      [32]byte
		CreatedAt *big.Int
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "childChain", arg0)
	return *ret, err
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, created_at uint256)
func (_RootChain *RootChainSession) ChildChain(arg0 *big.Int) (struct {
	Root      [32]byte
	CreatedAt *big.Int
}, error) {
	return _RootChain.Contract.ChildChain(&_RootChain.CallOpts, arg0)
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, created_at uint256)
func (_RootChain *RootChainCallerSession) ChildChain(arg0 *big.Int) (struct {
	Root      [32]byte
	CreatedAt *big.Int
}, error) {
	return _RootChain.Contract.ChildChain(&_RootChain.CallOpts, arg0)
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) CurrentChildBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentChildBlock")
	return *ret0, err
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_RootChain *RootChainSession) CurrentChildBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentChildBlock(&_RootChain.CallOpts)
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentChildBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentChildBlock(&_RootChain.CallOpts)
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) CurrentDepositBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentDepositBlock")
	return *ret0, err
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_RootChain *RootChainSession) CurrentDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentDepositBlock(&_RootChain.CallOpts)
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentDepositBlock(&_RootChain.CallOpts)
}

// ExitIds is a free data retrieval call binding the contract method 0x16b38840.
//
// Solidity: function exitIds( uint256) constant returns(uint256)
func (_RootChain *RootChainCaller) ExitIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "exitIds", arg0)
	return *ret0, err
}

// ExitIds is a free data retrieval call binding the contract method 0x16b38840.
//
// Solidity: function exitIds( uint256) constant returns(uint256)
func (_RootChain *RootChainSession) ExitIds(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.ExitIds(&_RootChain.CallOpts, arg0)
}

// ExitIds is a free data retrieval call binding the contract method 0x16b38840.
//
// Solidity: function exitIds( uint256) constant returns(uint256)
func (_RootChain *RootChainCallerSession) ExitIds(arg0 *big.Int) (*big.Int, error) {
	return _RootChain.Contract.ExitIds(&_RootChain.CallOpts, arg0)
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits( uint256) constant returns(owner address, amount uint256, utxoPos uint256)
func (_RootChain *RootChainCaller) Exits(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner   common.Address
	Amount  *big.Int
	UtxoPos *big.Int
}, error) {
	ret := new(struct {
		Owner   common.Address
		Amount  *big.Int
		UtxoPos *big.Int
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "exits", arg0)
	return *ret, err
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits( uint256) constant returns(owner address, amount uint256, utxoPos uint256)
func (_RootChain *RootChainSession) Exits(arg0 *big.Int) (struct {
	Owner   common.Address
	Amount  *big.Int
	UtxoPos *big.Int
}, error) {
	return _RootChain.Contract.Exits(&_RootChain.CallOpts, arg0)
}

// Exits is a free data retrieval call binding the contract method 0x342de179.
//
// Solidity: function exits( uint256) constant returns(owner address, amount uint256, utxoPos uint256)
func (_RootChain *RootChainCallerSession) Exits(arg0 *big.Int) (struct {
	Owner   common.Address
	Amount  *big.Int
	UtxoPos *big.Int
}, error) {
	return _RootChain.Contract.Exits(&_RootChain.CallOpts, arg0)
}

// GetChildChain is a free data retrieval call binding the contract method 0x85444de3.
//
// Solidity: function getChildChain(blockNumber uint256) constant returns(bytes32, uint256)
func (_RootChain *RootChainCaller) GetChildChain(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _RootChain.contract.Call(opts, out, "getChildChain", blockNumber)
	return *ret0, *ret1, err
}

// GetChildChain is a free data retrieval call binding the contract method 0x85444de3.
//
// Solidity: function getChildChain(blockNumber uint256) constant returns(bytes32, uint256)
func (_RootChain *RootChainSession) GetChildChain(blockNumber *big.Int) ([32]byte, *big.Int, error) {
	return _RootChain.Contract.GetChildChain(&_RootChain.CallOpts, blockNumber)
}

// GetChildChain is a free data retrieval call binding the contract method 0x85444de3.
//
// Solidity: function getChildChain(blockNumber uint256) constant returns(bytes32, uint256)
func (_RootChain *RootChainCallerSession) GetChildChain(blockNumber *big.Int) ([32]byte, *big.Int, error) {
	return _RootChain.Contract.GetChildChain(&_RootChain.CallOpts, blockNumber)
}

// GetDepositBlock is a free data retrieval call binding the contract method 0xbcd59261.
//
// Solidity: function getDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) GetDepositBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getDepositBlock")
	return *ret0, err
}

// GetDepositBlock is a free data retrieval call binding the contract method 0xbcd59261.
//
// Solidity: function getDepositBlock() constant returns(uint256)
func (_RootChain *RootChainSession) GetDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.GetDepositBlock(&_RootChain.CallOpts)
}

// GetDepositBlock is a free data retrieval call binding the contract method 0xbcd59261.
//
// Solidity: function getDepositBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) GetDepositBlock() (*big.Int, error) {
	return _RootChain.Contract.GetDepositBlock(&_RootChain.CallOpts)
}

// GetExit is a free data retrieval call binding the contract method 0xe60f1ff1.
//
// Solidity: function getExit(priority uint256) constant returns(address, uint256, uint256)
func (_RootChain *RootChainCaller) GetExit(opts *bind.CallOpts, priority *big.Int) (common.Address, *big.Int, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _RootChain.contract.Call(opts, out, "getExit", priority)
	return *ret0, *ret1, *ret2, err
}

// GetExit is a free data retrieval call binding the contract method 0xe60f1ff1.
//
// Solidity: function getExit(priority uint256) constant returns(address, uint256, uint256)
func (_RootChain *RootChainSession) GetExit(priority *big.Int) (common.Address, *big.Int, *big.Int, error) {
	return _RootChain.Contract.GetExit(&_RootChain.CallOpts, priority)
}

// GetExit is a free data retrieval call binding the contract method 0xe60f1ff1.
//
// Solidity: function getExit(priority uint256) constant returns(address, uint256, uint256)
func (_RootChain *RootChainCallerSession) GetExit(priority *big.Int) (common.Address, *big.Int, *big.Int, error) {
	return _RootChain.Contract.GetExit(&_RootChain.CallOpts, priority)
}

// NextWeekOldChildBlock is a free data retrieval call binding the contract method 0x107e58f1.
//
// Solidity: function nextWeekOldChildBlock(value uint256) constant returns(uint256)
func (_RootChain *RootChainCaller) NextWeekOldChildBlock(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "nextWeekOldChildBlock", value)
	return *ret0, err
}

// NextWeekOldChildBlock is a free data retrieval call binding the contract method 0x107e58f1.
//
// Solidity: function nextWeekOldChildBlock(value uint256) constant returns(uint256)
func (_RootChain *RootChainSession) NextWeekOldChildBlock(value *big.Int) (*big.Int, error) {
	return _RootChain.Contract.NextWeekOldChildBlock(&_RootChain.CallOpts, value)
}

// NextWeekOldChildBlock is a free data retrieval call binding the contract method 0x107e58f1.
//
// Solidity: function nextWeekOldChildBlock(value uint256) constant returns(uint256)
func (_RootChain *RootChainCallerSession) NextWeekOldChildBlock(value *big.Int) (*big.Int, error) {
	return _RootChain.Contract.NextWeekOldChildBlock(&_RootChain.CallOpts, value)
}

// RecentBlock is a free data retrieval call binding the contract method 0x6f84b695.
//
// Solidity: function recentBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) RecentBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "recentBlock")
	return *ret0, err
}

// RecentBlock is a free data retrieval call binding the contract method 0x6f84b695.
//
// Solidity: function recentBlock() constant returns(uint256)
func (_RootChain *RootChainSession) RecentBlock() (*big.Int, error) {
	return _RootChain.Contract.RecentBlock(&_RootChain.CallOpts)
}

// RecentBlock is a free data retrieval call binding the contract method 0x6f84b695.
//
// Solidity: function recentBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) RecentBlock() (*big.Int, error) {
	return _RootChain.Contract.RecentBlock(&_RootChain.CallOpts)
}

// WeekOldBlock is a free data retrieval call binding the contract method 0x4237b5f3.
//
// Solidity: function weekOldBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) WeekOldBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "weekOldBlock")
	return *ret0, err
}

// WeekOldBlock is a free data retrieval call binding the contract method 0x4237b5f3.
//
// Solidity: function weekOldBlock() constant returns(uint256)
func (_RootChain *RootChainSession) WeekOldBlock() (*big.Int, error) {
	return _RootChain.Contract.WeekOldBlock(&_RootChain.CallOpts)
}

// WeekOldBlock is a free data retrieval call binding the contract method 0x4237b5f3.
//
// Solidity: function weekOldBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) WeekOldBlock() (*big.Int, error) {
	return _RootChain.Contract.WeekOldBlock(&_RootChain.CallOpts)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x32773ba3.
//
// Solidity: function challengeExit(cUtxoPos uint256, eUtxoPos uint256, txBytes bytes, proof bytes, sigs bytes, confirmationSig bytes) returns()
func (_RootChain *RootChainTransactor) ChallengeExit(opts *bind.TransactOpts, cUtxoPos *big.Int, eUtxoPos *big.Int, txBytes []byte, proof []byte, sigs []byte, confirmationSig []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "challengeExit", cUtxoPos, eUtxoPos, txBytes, proof, sigs, confirmationSig)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x32773ba3.
//
// Solidity: function challengeExit(cUtxoPos uint256, eUtxoPos uint256, txBytes bytes, proof bytes, sigs bytes, confirmationSig bytes) returns()
func (_RootChain *RootChainSession) ChallengeExit(cUtxoPos *big.Int, eUtxoPos *big.Int, txBytes []byte, proof []byte, sigs []byte, confirmationSig []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeExit(&_RootChain.TransactOpts, cUtxoPos, eUtxoPos, txBytes, proof, sigs, confirmationSig)
}

// ChallengeExit is a paid mutator transaction binding the contract method 0x32773ba3.
//
// Solidity: function challengeExit(cUtxoPos uint256, eUtxoPos uint256, txBytes bytes, proof bytes, sigs bytes, confirmationSig bytes) returns()
func (_RootChain *RootChainTransactorSession) ChallengeExit(cUtxoPos *big.Int, eUtxoPos *big.Int, txBytes []byte, proof []byte, sigs []byte, confirmationSig []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeExit(&_RootChain.TransactOpts, cUtxoPos, eUtxoPos, txBytes, proof, sigs, confirmationSig)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_RootChain *RootChainTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_RootChain *RootChainSession) Deposit() (*types.Transaction, error) {
	return _RootChain.Contract.Deposit(&_RootChain.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_RootChain *RootChainTransactorSession) Deposit() (*types.Transaction, error) {
	return _RootChain.Contract.Deposit(&_RootChain.TransactOpts)
}

// FinalizeExits is a paid mutator transaction binding the contract method 0xc6ab44cd.
//
// Solidity: function finalizeExits() returns(uint256)
func (_RootChain *RootChainTransactor) FinalizeExits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "finalizeExits")
}

// FinalizeExits is a paid mutator transaction binding the contract method 0xc6ab44cd.
//
// Solidity: function finalizeExits() returns(uint256)
func (_RootChain *RootChainSession) FinalizeExits() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeExits(&_RootChain.TransactOpts)
}

// FinalizeExits is a paid mutator transaction binding the contract method 0xc6ab44cd.
//
// Solidity: function finalizeExits() returns(uint256)
func (_RootChain *RootChainTransactorSession) FinalizeExits() (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeExits(&_RootChain.TransactOpts)
}

// StartExit is a paid mutator transaction binding the contract method 0x1c91a6b9.
//
// Solidity: function startExit(utxoPos uint256, txBytes bytes, proof bytes, sigs bytes) returns()
func (_RootChain *RootChainTransactor) StartExit(opts *bind.TransactOpts, utxoPos *big.Int, txBytes []byte, proof []byte, sigs []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startExit", utxoPos, txBytes, proof, sigs)
}

// StartExit is a paid mutator transaction binding the contract method 0x1c91a6b9.
//
// Solidity: function startExit(utxoPos uint256, txBytes bytes, proof bytes, sigs bytes) returns()
func (_RootChain *RootChainSession) StartExit(utxoPos *big.Int, txBytes []byte, proof []byte, sigs []byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, utxoPos, txBytes, proof, sigs)
}

// StartExit is a paid mutator transaction binding the contract method 0x1c91a6b9.
//
// Solidity: function startExit(utxoPos uint256, txBytes bytes, proof bytes, sigs bytes) returns()
func (_RootChain *RootChainTransactorSession) StartExit(utxoPos *big.Int, txBytes []byte, proof []byte, sigs []byte) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, utxoPos, txBytes, proof, sigs)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(root bytes32) returns()
func (_RootChain *RootChainTransactor) SubmitBlock(opts *bind.TransactOpts, root [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitBlock", root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(root bytes32) returns()
func (_RootChain *RootChainSession) SubmitBlock(root [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitBlock(&_RootChain.TransactOpts, root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(root bytes32) returns()
func (_RootChain *RootChainTransactorSession) SubmitBlock(root [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitBlock(&_RootChain.TransactOpts, root)
}

// RootChainDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the RootChain contract.
type RootChainDepositIterator struct {
	Event *RootChainDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootChainDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootChainDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootChainDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainDeposit represents a Deposit event raised by the RootChain contract.
type RootChainDeposit struct {
	Depositor    common.Address
	Amount       *big.Int
	DepositBlock *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(depositor address, amount uint256, depositBlock uint256)
func (_RootChain *RootChainFilterer) FilterDeposit(opts *bind.FilterOpts) (*RootChainDepositIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &RootChainDepositIterator{contract: _RootChain.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x90890809c654f11d6e72a28fa60149770a0d11ec6c92319d6ceb2bb0a4ea1a15.
//
// Solidity: event Deposit(depositor address, amount uint256, depositBlock uint256)
func (_RootChain *RootChainFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *RootChainDeposit) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainDeposit)
				if err := _RootChain.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RootChainExitIterator is returned from FilterExit and is used to iterate over the raw logs and unpacked data for Exit events raised by the RootChain contract.
type RootChainExitIterator struct {
	Event *RootChainExit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootChainExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainExit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootChainExit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootChainExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainExit represents a Exit event raised by the RootChain contract.
type RootChainExit struct {
	Exitor  common.Address
	UtxoPos *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExit is a free log retrieval operation binding the contract event 0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631.
//
// Solidity: event Exit(exitor address, utxoPos uint256)
func (_RootChain *RootChainFilterer) FilterExit(opts *bind.FilterOpts) (*RootChainExitIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "Exit")
	if err != nil {
		return nil, err
	}
	return &RootChainExitIterator{contract: _RootChain.contract, event: "Exit", logs: logs, sub: sub}, nil
}

// WatchExit is a free log subscription operation binding the contract event 0x22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631.
//
// Solidity: event Exit(exitor address, utxoPos uint256)
func (_RootChain *RootChainFilterer) WatchExit(opts *bind.WatchOpts, sink chan<- *RootChainExit) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "Exit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainExit)
				if err := _RootChain.contract.UnpackLog(event, "Exit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RootChainSubmitIterator is returned from FilterSubmit and is used to iterate over the raw logs and unpacked data for Submit events raised by the RootChain contract.
type RootChainSubmitIterator struct {
	Event *RootChainSubmit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RootChainSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainSubmit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RootChainSubmit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RootChainSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainSubmit represents a Submit event raised by the RootChain contract.
type RootChainSubmit struct {
	SubmitBlock *big.Int
	Root        [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmit is a free log retrieval operation binding the contract event 0x03e03b35d08c09e37050f76f1a914a334810af275603f96ed8ad6f2cc3617eba.
//
// Solidity: event Submit(submitBlock uint256, root bytes32)
func (_RootChain *RootChainFilterer) FilterSubmit(opts *bind.FilterOpts) (*RootChainSubmitIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "Submit")
	if err != nil {
		return nil, err
	}
	return &RootChainSubmitIterator{contract: _RootChain.contract, event: "Submit", logs: logs, sub: sub}, nil
}

// WatchSubmit is a free log subscription operation binding the contract event 0x03e03b35d08c09e37050f76f1a914a334810af275603f96ed8ad6f2cc3617eba.
//
// Solidity: event Submit(submitBlock uint256, root bytes32)
func (_RootChain *RootChainFilterer) WatchSubmit(opts *bind.WatchOpts, sink chan<- *RootChainSubmit) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "Submit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainSubmit)
				if err := _RootChain.contract.UnpackLog(event, "Submit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820117122f6316c1af0e5aa769cc3004989356f14473ece36926dcb37b714a9c40b0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// ValidateABI is the input ABI used to generate the binding from.
const ValidateABI = "[]"

// ValidateBin is the compiled bytecode used for deploying new contracts.
const ValidateBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820c5033f318829ed0b65ba08d3bb349a6dafd33cead49fb8837a57f10c0fdbaed60029`

// DeployValidate deploys a new Ethereum contract, binding an instance of Validate to it.
func DeployValidate(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validate, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidateABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validate{ValidateCaller: ValidateCaller{contract: contract}, ValidateTransactor: ValidateTransactor{contract: contract}, ValidateFilterer: ValidateFilterer{contract: contract}}, nil
}

// Validate is an auto generated Go binding around an Ethereum contract.
type Validate struct {
	ValidateCaller     // Read-only binding to the contract
	ValidateTransactor // Write-only binding to the contract
	ValidateFilterer   // Log filterer for contract events
}

// ValidateCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidateSession struct {
	Contract     *Validate         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidateCallerSession struct {
	Contract *ValidateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ValidateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidateTransactorSession struct {
	Contract     *ValidateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ValidateRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidateRaw struct {
	Contract *Validate // Generic contract binding to access the raw methods on
}

// ValidateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidateCallerRaw struct {
	Contract *ValidateCaller // Generic read-only contract binding to access the raw methods on
}

// ValidateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidateTransactorRaw struct {
	Contract *ValidateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidate creates a new instance of Validate, bound to a specific deployed contract.
func NewValidate(address common.Address, backend bind.ContractBackend) (*Validate, error) {
	contract, err := bindValidate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validate{ValidateCaller: ValidateCaller{contract: contract}, ValidateTransactor: ValidateTransactor{contract: contract}, ValidateFilterer: ValidateFilterer{contract: contract}}, nil
}

// NewValidateCaller creates a new read-only instance of Validate, bound to a specific deployed contract.
func NewValidateCaller(address common.Address, caller bind.ContractCaller) (*ValidateCaller, error) {
	contract, err := bindValidate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidateCaller{contract: contract}, nil
}

// NewValidateTransactor creates a new write-only instance of Validate, bound to a specific deployed contract.
func NewValidateTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidateTransactor, error) {
	contract, err := bindValidate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidateTransactor{contract: contract}, nil
}

// NewValidateFilterer creates a new log filterer instance of Validate, bound to a specific deployed contract.
func NewValidateFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidateFilterer, error) {
	contract, err := bindValidate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidateFilterer{contract: contract}, nil
}

// bindValidate binds a generic wrapper to an already deployed contract.
func bindValidate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validate *ValidateRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validate.Contract.ValidateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validate *ValidateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validate.Contract.ValidateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validate *ValidateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validate.Contract.ValidateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validate *ValidateCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validate *ValidateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validate *ValidateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validate.Contract.contract.Transact(opts, method, params...)
}

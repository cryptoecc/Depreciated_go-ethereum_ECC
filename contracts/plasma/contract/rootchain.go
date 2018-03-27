// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ByteUtilsABI is the input ABI used to generate the binding from.
const ByteUtilsABI = "[]"

// ByteUtilsBin is the compiled bytecode used for deploying new contracts.
const ByteUtilsBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a723058204a605024beb6ecc9662a7f3f5e025ea3818b253ea8df9435b0bdd5a833b3b2970029`

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
	return address, tx, &ByteUtils{ByteUtilsCaller: ByteUtilsCaller{contract: contract}, ByteUtilsTransactor: ByteUtilsTransactor{contract: contract}}, nil
}

// ByteUtils is an auto generated Go binding around an Ethereum contract.
type ByteUtils struct {
	ByteUtilsCaller     // Read-only binding to the contract
	ByteUtilsTransactor // Write-only binding to the contract
}

// ByteUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ByteUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ByteUtilsTransactor struct {
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
	contract, err := bindByteUtils(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ByteUtils{ByteUtilsCaller: ByteUtilsCaller{contract: contract}, ByteUtilsTransactor: ByteUtilsTransactor{contract: contract}}, nil
}

// NewByteUtilsCaller creates a new read-only instance of ByteUtils, bound to a specific deployed contract.
func NewByteUtilsCaller(address common.Address, caller bind.ContractCaller) (*ByteUtilsCaller, error) {
	contract, err := bindByteUtils(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ByteUtilsCaller{contract: contract}, nil
}

// NewByteUtilsTransactor creates a new write-only instance of ByteUtils, bound to a specific deployed contract.
func NewByteUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ByteUtilsTransactor, error) {
	contract, err := bindByteUtils(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ByteUtilsTransactor{contract: contract}, nil
}

// bindByteUtils binds a generic wrapper to an already deployed contract.
func bindByteUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ByteUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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
const ECRecoveryBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a723058204c8022d04ec69ed553df56173729f4ac1fa1ec4ca8c7e907bb3f77ca4babc2ca0029`

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
	return address, tx, &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}}, nil
}

// ECRecovery is an auto generated Go binding around an Ethereum contract.
type ECRecovery struct {
	ECRecoveryCaller     // Read-only binding to the contract
	ECRecoveryTransactor // Write-only binding to the contract
}

// ECRecoveryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ECRecoveryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECRecoveryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECRecoveryTransactor struct {
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
	contract, err := bindECRecovery(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECRecovery{ECRecoveryCaller: ECRecoveryCaller{contract: contract}, ECRecoveryTransactor: ECRecoveryTransactor{contract: contract}}, nil
}

// NewECRecoveryCaller creates a new read-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryCaller(address common.Address, caller bind.ContractCaller) (*ECRecoveryCaller, error) {
	contract, err := bindECRecovery(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryCaller{contract: contract}, nil
}

// NewECRecoveryTransactor creates a new write-only instance of ECRecovery, bound to a specific deployed contract.
func NewECRecoveryTransactor(address common.Address, transactor bind.ContractTransactor) (*ECRecoveryTransactor, error) {
	contract, err := bindECRecovery(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ECRecoveryTransactor{contract: contract}, nil
}

// bindECRecovery binds a generic wrapper to an already deployed contract.
func bindECRecovery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ECRecoveryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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
const MathBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a723058200a9085802bea0d65c038ad1f718c545603a9d0d99320d153e66feaf49753beae0029`

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
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
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
	contract, err := bindMath(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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
const MerkleBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a72305820e0156048576727685577b61abbe74f7ed02aa98a8f94f2e6351c86d37aefd8e60029`

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
	return address, tx, &Merkle{MerkleCaller: MerkleCaller{contract: contract}, MerkleTransactor: MerkleTransactor{contract: contract}}, nil
}

// Merkle is an auto generated Go binding around an Ethereum contract.
type Merkle struct {
	MerkleCaller     // Read-only binding to the contract
	MerkleTransactor // Write-only binding to the contract
}

// MerkleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTransactor struct {
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
	contract, err := bindMerkle(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Merkle{MerkleCaller: MerkleCaller{contract: contract}, MerkleTransactor: MerkleTransactor{contract: contract}}, nil
}

// NewMerkleCaller creates a new read-only instance of Merkle, bound to a specific deployed contract.
func NewMerkleCaller(address common.Address, caller bind.ContractCaller) (*MerkleCaller, error) {
	contract, err := bindMerkle(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleCaller{contract: contract}, nil
}

// NewMerkleTransactor creates a new write-only instance of Merkle, bound to a specific deployed contract.
func NewMerkleTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTransactor, error) {
	contract, err := bindMerkle(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MerkleTransactor{contract: contract}, nil
}

// bindMerkle binds a generic wrapper to an already deployed contract.
func bindMerkle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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
const PriorityQueueBin = `0x6060604052341561000f57600080fd5b60008054600160a060020a03191633600160a060020a03161790556020604051908101604052600081526100469060019081610051565b5060006002556100be565b828054828255906000526020600020908101928215610091579160200282015b82811115610091578251829060ff16905591602001919060010190610071565b5061009d9291506100a1565b5090565b6100bb91905b8082111561009d57600081556001016100a7565b90565b6105db806100cd6000396000f30060606040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632dcdcd0c811461007157806390b5561d14610099578063b07576ac146100b1578063bda1504b146100c4578063d6362e97146100d7575b600080fd5b341561007c57600080fd5b6100876004356100ea565b60405190815260200160405180910390f35b34156100a457600080fd5b6100af6004356101ba565b005b34156100bc57600080fd5b610087610227565b34156100cf57600080fd5b6100876102ec565b34156100e257600080fd5b6100876102f2565b6000600254610114600161010860028661031590919063ffffffff16565b9063ffffffff61034b16565b11156101325761012b82600263ffffffff61031516565b90506101b5565b60016101498161010885600263ffffffff61031516565b8154811061015357fe5b600091825260209091200154600161017284600263ffffffff61031516565b8154811061017c57fe5b906000526020600020900154101561019f5761012b82600263ffffffff61031516565b61012b600161010884600263ffffffff61031516565b919050565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146101e257600080fd5b600180548082016101f38382610576565b50600091825260209091200181905560025461021690600163ffffffff61034b16565b60028190556102249061035a565b50565b6000805481903373ffffffffffffffffffffffffffffffffffffffff90811691161461025257600080fd5b600180548190811061026057fe5b9060005260206000209001549050600160025481548110151561027f57fe5b90600052602060002090015460018081548110151561029a57fe5b6000918252602090912001556002546001805490919081106102b857fe5b60009182526020822001556002546102d790600163ffffffff61046916565b6002556102e4600161047b565b8091505b5090565b60025481565b600060018081548110151561030357fe5b90600052602060002090015490505b90565b6000808315156103285760009150610344565b5082820282848281151561033857fe5b041461034057fe5b8091505b5092915050565b60008282018381101561034057fe5b60005b600061037083600263ffffffff61055f16565b111561046557600161038983600263ffffffff61055f16565b8154811061039357fe5b9060005260206000209001546001838154811015156103ae57fe5b906000526020600020900154101561044d5760016103d383600263ffffffff61055f16565b815481106103dd57fe5b90600052602060002090015490506001828154811015156103fa57fe5b600091825260209091200154600161041984600263ffffffff61055f16565b8154811061042357fe5b600091825260209091200155600180548291908490811061044057fe5b6000918252602090912001555b61045e82600263ffffffff61055f16565b915061035d565b5050565b60008282111561047557fe5b50900390565b6000805b60025461049660028561031590919063ffffffff16565b1161055a576104a4836100ea565b91506001828154811015156104b557fe5b9060005260206000209001546001848154811015156104d057fe5b90600052602060002090015411156105525760018054849081106104f057fe5b906000526020600020900154905060018281548110151561050d57fe5b90600052602060002090015460018481548110151561052857fe5b600091825260209091200155600180548291908490811061054557fe5b6000918252602090912001555b81925061047f565b505050565b600080828481151561056d57fe5b04949350505050565b81548183558181151161055a5760008381526020902061055a91810190830161031291905b808211156102e8576000815560010161059b5600a165627a7a72305820a19df8a356e5aa336313e2f2c05c7d43b27a533f8d2c532f56b4e36c163fff380029`

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
	return address, tx, &PriorityQueue{PriorityQueueCaller: PriorityQueueCaller{contract: contract}, PriorityQueueTransactor: PriorityQueueTransactor{contract: contract}}, nil
}

// PriorityQueue is an auto generated Go binding around an Ethereum contract.
type PriorityQueue struct {
	PriorityQueueCaller     // Read-only binding to the contract
	PriorityQueueTransactor // Write-only binding to the contract
}

// PriorityQueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type PriorityQueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PriorityQueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PriorityQueueTransactor struct {
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
	contract, err := bindPriorityQueue(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PriorityQueue{PriorityQueueCaller: PriorityQueueCaller{contract: contract}, PriorityQueueTransactor: PriorityQueueTransactor{contract: contract}}, nil
}

// NewPriorityQueueCaller creates a new read-only instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueueCaller(address common.Address, caller bind.ContractCaller) (*PriorityQueueCaller, error) {
	contract, err := bindPriorityQueue(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PriorityQueueCaller{contract: contract}, nil
}

// NewPriorityQueueTransactor creates a new write-only instance of PriorityQueue, bound to a specific deployed contract.
func NewPriorityQueueTransactor(address common.Address, transactor bind.ContractTransactor) (*PriorityQueueTransactor, error) {
	contract, err := bindPriorityQueue(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PriorityQueueTransactor{contract: contract}, nil
}

// bindPriorityQueue binds a generic wrapper to an already deployed contract.
func bindPriorityQueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PriorityQueueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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
const RLPBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a7230582070ef0cf6157ea46fe46f797baf97692abf7245caf2e8f179bbbf92e643ef5dd50029`

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
	return address, tx, &RLP{RLPCaller: RLPCaller{contract: contract}, RLPTransactor: RLPTransactor{contract: contract}}, nil
}

// RLP is an auto generated Go binding around an Ethereum contract.
type RLP struct {
	RLPCaller     // Read-only binding to the contract
	RLPTransactor // Write-only binding to the contract
}

// RLPCaller is an auto generated read-only Go binding around an Ethereum contract.
type RLPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RLPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RLPTransactor struct {
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
	contract, err := bindRLP(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RLP{RLPCaller: RLPCaller{contract: contract}, RLPTransactor: RLPTransactor{contract: contract}}, nil
}

// NewRLPCaller creates a new read-only instance of RLP, bound to a specific deployed contract.
func NewRLPCaller(address common.Address, caller bind.ContractCaller) (*RLPCaller, error) {
	contract, err := bindRLP(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RLPCaller{contract: contract}, nil
}

// NewRLPTransactor creates a new write-only instance of RLP, bound to a specific deployed contract.
func NewRLPTransactor(address common.Address, transactor bind.ContractTransactor) (*RLPTransactor, error) {
	contract, err := bindRLP(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RLPTransactor{contract: contract}, nil
}

// bindRLP binds a generic wrapper to an already deployed contract.
func bindRLP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RLPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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

// RootChainABI is the input ABI used to generate the binding from.
const RootChainABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exitIds\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"utxoPos\",\"type\":\"uint256\"},{\"name\":\"txBytes\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"sigs\",\"type\":\"bytes\"}],\"name\":\"startExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"cUtxoPos\",\"type\":\"uint256\"},{\"name\":\"eUtxoPos\",\"type\":\"uint256\"},{\"name\":\"txBytes\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"sigs\",\"type\":\"bytes\"},{\"name\":\"confirmationSig\",\"type\":\"bytes\"}],\"name\":\"challengeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exits\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"utxoPos\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weekOldBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"blknum\",\"type\":\"uint256\"}],\"name\":\"submitBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"recentBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentChildBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getChildChain\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txBytes\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalizeExits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"priority\",\"type\":\"uint256\"}],\"name\":\"getExit\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"childChain\",\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"created_at\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"exitor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"utxoPos\",\"type\":\"uint256\"}],\"name\":\"Exit\",\"type\":\"event\"}]"

// RootChainBin is the compiled bytecode used for deploying new contracts.
const RootChainBin = `0x6060604052341561000f57600080fd5b60048054600160a060020a03191633600160a060020a03161790556001600555610037610072565b604051809103906000f080151561004d57600080fd5b60038054600160a060020a031916600160a060020a0392909216919091179055610082565b6040516106a880611aba83390190565b611a29806100916000396000f3006060604052600436106100b65763ffffffff60e060020a60003504166316b3884081146100bb5780631c91a6b9146100e357806332773ba3146101bf578063342de179146102e15780634237b5f31461032d57806346ab67cb146103405780636f84b695146103595780637a95f1e81461036c57806385444de31461037f57806398b1e06a146103ad578063bf7e214f146103f3578063c6ab44cd14610422578063e60f1ff114610435578063f95643b11461044b575b600080fd5b34156100c657600080fd5b6100d1600435610461565b60405190815260200160405180910390f35b34156100ee57600080fd5b6101bd600480359060446024803590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061047395505050505050565b005b34156101ca57600080fd5b6101bd600480359060248035919060649060443590810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284378201915050505050509190803590602001908201803590602001908080601f01602080910402602001604051908101604052818152929190602084018383808284375094965061086e95505050505050565b34156102ec57600080fd5b6102f7600435610a30565b6040518084600160a060020a0316600160a060020a03168152602001838152602001828152602001935050505060405180910390f35b341561033857600080fd5b6100d1610a5c565b341561034b57600080fd5b6101bd600435602435610a62565b341561036457600080fd5b6100d1610b45565b341561037757600080fd5b6100d1610b4b565b341561038a57600080fd5b610395600435610b51565b60405191825260208201526040908101905180910390f35b6101bd60046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610b6b95505050505050565b34156103fe57600080fd5b610406610de1565b604051600160a060020a03909116815260200160405180910390f35b341561042d57600080fd5b6100d1610df0565b341561044057600080fd5b6102f760043561117e565b341561045657600080fd5b6103956004356111ac565b60026020526000908152604090205481565b61047b611993565b6000806000806000806000805b61049b4262093a8063ffffffff6111c516565b60075460009081526020819052604090206001015410156104f25760075460009081526020819052604090206001015415156104d6576104f2565b6007546104ea90600163ffffffff6111d716565b600755610488565b61050c600b6105008e6111f1565b9063ffffffff61124516565b9850633b9aca00808e049850612710908e0660008a81526020819052604090205491900497506127108802633b9aca008a028f03039650945061056989600660028902018151811061055a57fe5b906020019060200201516112fb565b600160a060020a031633600160a060020a031614151561058857600080fd5b8b6040518082805190602001908083835b602083106105b85780518252601f199092019160209182019101610599565b6001836020036101000a038019825116818451161790925250505091909101925060409150505180910390209350836105f48b60006082611348565b6040518281526020810182805190602001908083835b602083106106295780518252601f19909201916020918201910161060a565b6001836020036101000a038019825116818451161790925250505091909101935060409250505051809103902092506106778960008151811061066857fe5b9060200190602002015161139e565b6106878a60038151811061066857fe5b633b9aca000201915061069c8486848d6113e5565b15156106a757600080fd5b6106b98388878e63ffffffff61156116565b15156106c457600080fd5b6007548810156106f3576106ec600754898f8115156106df57fe5b049063ffffffff6115ee16565b90506106f6565b508b5b60008d8152600260205260409020541561070f57600080fd5b60008d81526002602052604090819020829055600354600160a060020a0316906390b5561d9083905160e060020a63ffffffff84160281526004810191909152602401600060405180830381600087803b151561076b57600080fd5b6102c65a03f1151561077c57600080fd5b5050506060604051908101604052806107a08b896002026006018151811061055a57fe5b600160a060020a031681526020016107c38b896002026007018151811061066857fe5b815260209081018f905260008381526001909152604090208151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0391909116178155602082015181600101556040820151600290910155507f22d324652c93739755cf4581508b60875ebdd78c20c0cff5cf8e23452b299631338e604051600160a060020a03909216825260208201526040908101905180910390a150505050505050505050505050565b633b9aca0080870460009081526020818152604080832054898452600290925280832054612710948b0694909404939192908190819081908b90518082805190602001908083835b602083106108d55780518252601f1990920191602091820191016108b6565b6001836020036101000a03801982511681845116179092525050509190910192506040915050518091039020935083866040519182526020820152604090810190518091039020925083896040518281526020810182805190602001908083835b602083106109555780518252601f199092019160209182019101610936565b6001836020036101000a038019825116818451161790925250505091909101935060409250505051908190039020600086815260016020526040902054909250600160a060020a031690506109aa8389611619565b600160a060020a038281169116146109c157600080fd5b6109d38288888d63ffffffff61156116565b15156109de57600080fd5b505050600091825250600160208181526040808420805473ffffffffffffffffffffffffffffffffffffffff191681559283018490556002928301849055998352529687209690965550505050505050565b6001602081905260009182526040909120805491810154600290910154600160a060020a039092169183565b60075481565b60045433600160a060020a03908116911614610a7d57600080fd5b610a904262093a8063ffffffff6111c516565b6007546000908152602081905260409020600101541015610ae7576007546000908152602081905260409020600101541515610acb57610ae7565b600754610adf90600163ffffffff6111d716565b600755610a7d565b6005548114610af557600080fd5b6040805190810160409081528382524260208084019190915260055460009081529081905220815181556020820151600191820155600554610b3e92509063ffffffff6111d716565b6005555050565b60065481565b60055481565b600090815260208190526040902080546001909101549091565b610b73611993565b6000806000610b86600b610500876111f1565b93508351600b14610b9657600080fd5b6006831015610bc357610bae84848151811061066857fe5b15610bb857600080fd5b600190920191610b96565b34610bd48560078151811061066857fe5b14610bde57600080fd5b610bee8460098151811061066857fe5b15610bf857600080fd5b846040518082805190602001908083835b60208310610c285780518252601f199092019160209182019101610c09565b6001836020036101000a038019825116818451161790925250505091909101925060409150505180910390206082604051805910610c635750595b818152601f19601f830116810160200160405290506040518281526020810182805190602001908083835b60208310610cad5780518252601f199092019160209182019101610c8e565b6001836020036101000a03801982511681845116179092525050509190910193506040925050505180910390209050600092505b6010831015610d2c5780826040519182526020820152604090810190518091039020905081826040519182526020820152604090810190519081900390206001909301929150610ce1565b6040805190810160409081528282524260208084019190915260055460009081529081905220815181556020820151600191820155600554610d7592509063ffffffff6111d716565b6005557fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c610da98560068151811061055a57fe5b610db98660078151811061066857fe5b604051600160a060020a03909216825260208201526040908101905180910390a15050505050565b600454600160a060020a031681565b600080610dfb6119a5565b6000805b610e124262093a8063ffffffff6111c516565b6007546000908152602081905260409020600101541015610e69576007546000908152602081905260409020600101541515610e4d57610e69565b600754610e6190600163ffffffff6111d716565b600755610dff565b610e7c426212750063ffffffff6111c516565b600354909450600190600090600160a060020a031663d6362e9782604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610ecc57600080fd5b6102c65a03f11515610edd57600080fd5b505050604051805190508152602001908152602001600020606060405190810160409081528254600160a060020a03168252600183015460208301526002909201549181019182529350610f3e90633b9aca0090519063ffffffff6116f916565b91505b6000828152602081905260409020600101548490108015610fc45750600354600090600160a060020a031663bda1504b82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610fa757600080fd5b6102c65a03f11515610fb857600080fd5b50505060405180519050115b15611177578251600160a060020a03166108fc84602001519081150290604051600060405180830381858888f19350505050151561100157600080fd5b600354600160a060020a031663b07576ac6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561104957600080fd5b6102c65a03f1151561105a57600080fd5b50505060405180516000818152600160208190526040808320805473ffffffffffffffffffffffffffffffffffffffff19168155918201839055600291820183905592945092509085015181526020019081526020016000206000905560016000600360009054906101000a9004600160a060020a0316600160a060020a031663d6362e976000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561111657600080fd5b6102c65a03f1151561112757600080fd5b505050604051805190508152602001908152602001600020606060405190810160409081528254600160a060020a0316825260018301546020830152600290920154918101919091529250610f41565b5050505090565b6000818152600160208190526040909120805491810154600290910154600160a060020a0390921693909250565b6000602081905290815260409020805460019091015482565b6000828211156111d157fe5b50900390565b6000828201838110156111e657fe5b8091505b5092915050565b6111f96119c5565b60008083519150811515611222576040805190810160405260008082526020820152925061123e565b5060208301604080519081016040528181526020810183905292505b5050919050565b61124d611993565b6112556119dc565b600061126085611710565b151561126b57600080fd5b836040518059106112795750595b9080825280602002602001820160405280156112af57816020015b61129c6119c5565b8152602001906001900390816112945790505b5092506112bb8561173d565b91505b6112c782611776565b156112f3576112d582611799565b8382815181106112e157fe5b602090810290910101526001016112be565b505092915050565b6000806000611309846117db565b151561131457600080fd5b61131d84611805565b90925090506014811461132f57600080fd5b6c01000000000000000000000000825104949350505050565b611350611993565b611358611993565b6040519050601f831680820184810186838901015b8183101561138557805183526020928301920161136d565b5050848352601f01601f19166040525090509392505050565b60008060006113ac846117db565b15156113b757600080fd5b6113c084611805565b9150915060208111156113d257600080fd5b806020036101000a825104949350505050565b60006113ef611993565b6113f7611993565b6113ff611993565b6000611409611993565b6000806041895181151561141957fe5b0615801561142a5750610104895111155b151561143557600080fd5b6114428960006041611348565b965061145089604180611348565b955061145f8960826041611348565b94508b8b6040519182526020820152604090810190518091039020935089600014156114aa5761148f8486611619565b600160a060020a031633600160a060020a0316149750611552565b633b9aca008a10156114e4576114c08486611619565b600160a060020a03166114d38d89611619565b600160a060020a0316149750611552565b6114f18960c36041611348565b92506114fd8486611619565b600160a060020a03166115108d89611619565b600160a060020a03161491506115268484611619565b600160a060020a03166115398d88611619565b600160a060020a031614905081801561154f5750805b97505b50505050505050949350505050565b60008060008084516102001461157657600080fd5b5086905060205b61020081116115e0578085015192506002870615156115b6578183604051918252602082015260409081019051809103902091506115d2565b8282604051918252602082015260409081019051809103902091505b60028704965060200161157d565b509390931495945050505050565b60008083151561160157600091506111ea565b5082820282848281151561161157fe5b04146111e657fe5b600080600080845160411461163157600093506116f0565b6020850151925060408501519150606085015160001a9050601b8160ff16101561165957601b015b8060ff16601b1415801561167157508060ff16601c14155b1561167f57600093506116f0565b6001868285856040516000815260200160405260006040516020015260405193845260ff90921660208085019190915260408085019290925260608401929092526080909201915160208103908084039060008661646e5a03f115156116e457600080fd5b50506020604051035193505b50505092915050565b600080828481151561170757fe5b04949350505050565b600080826020015115156117275760009150611737565b8251905060c0815160001a101591505b50919050565b6117456119dc565b600061175083611710565b151561175b57600080fd5b61176483611882565b83519383529092016020820152919050565b60006117806119c5565b8251905080602001518151018360200151109392505050565b6117a16119c5565b6000806117ad84611776565b156100b657836020015191506117c282611901565b828452602080850182905283820190860152905061123e565b600080826020015115156117f25760009150611737565b8251905060c0815160001a109392505050565b6000806000806000611816866117db565b151561182157600080fd5b85519150815160001a92506080831015611841578194506001935061187a565b60b883101561185f576001866020015103935081600101945061187a565b5060b619820180600160208801510303935080820160010194505b505050915091565b60008060008360200151151561189b576000925061123e565b83519050805160001a915060808210156118b8576000925061123e565b60b88210806118d3575060c082101580156118d3575060f882105b156118e1576001925061123e565b60c08210156118f65760b5198201925061123e565b5060f5190192915050565b600080825160001a9050608081101561191d5760019150611737565b60b881101561193257607e1981019150611737565b60c081101561195c5760b78103806020036101000a60018501510480820160010193505050611737565b60f88110156119715760be1981019150611737565b60f78103806020036101000a6001850151048082016001019350505050919050565b60206040519081016040526000815290565b606060405190810160409081526000808352602083018190529082015290565b604080519081016040526000808252602082015290565b6060604051908101604052806119f06119c5565b81526020016000815250905600a165627a7a72305820d6b5f4d9063cc98896ba2d21f6253108498195e6d50e011d151ee6f46d991a4500296060604052341561000f57600080fd5b60008054600160a060020a03191633600160a060020a03161790556020604051908101604052600081526100469060019081610051565b5060006002556100be565b828054828255906000526020600020908101928215610091579160200282015b82811115610091578251829060ff16905591602001919060010190610071565b5061009d9291506100a1565b5090565b6100bb91905b8082111561009d57600081556001016100a7565b90565b6105db806100cd6000396000f30060606040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632dcdcd0c811461007157806390b5561d14610099578063b07576ac146100b1578063bda1504b146100c4578063d6362e97146100d7575b600080fd5b341561007c57600080fd5b6100876004356100ea565b60405190815260200160405180910390f35b34156100a457600080fd5b6100af6004356101ba565b005b34156100bc57600080fd5b610087610227565b34156100cf57600080fd5b6100876102ec565b34156100e257600080fd5b6100876102f2565b6000600254610114600161010860028661031590919063ffffffff16565b9063ffffffff61034b16565b11156101325761012b82600263ffffffff61031516565b90506101b5565b60016101498161010885600263ffffffff61031516565b8154811061015357fe5b600091825260209091200154600161017284600263ffffffff61031516565b8154811061017c57fe5b906000526020600020900154101561019f5761012b82600263ffffffff61031516565b61012b600161010884600263ffffffff61031516565b919050565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146101e257600080fd5b600180548082016101f38382610576565b50600091825260209091200181905560025461021690600163ffffffff61034b16565b60028190556102249061035a565b50565b6000805481903373ffffffffffffffffffffffffffffffffffffffff90811691161461025257600080fd5b600180548190811061026057fe5b9060005260206000209001549050600160025481548110151561027f57fe5b90600052602060002090015460018081548110151561029a57fe5b6000918252602090912001556002546001805490919081106102b857fe5b60009182526020822001556002546102d790600163ffffffff61046916565b6002556102e4600161047b565b8091505b5090565b60025481565b600060018081548110151561030357fe5b90600052602060002090015490505b90565b6000808315156103285760009150610344565b5082820282848281151561033857fe5b041461034057fe5b8091505b5092915050565b60008282018381101561034057fe5b60005b600061037083600263ffffffff61055f16565b111561046557600161038983600263ffffffff61055f16565b8154811061039357fe5b9060005260206000209001546001838154811015156103ae57fe5b906000526020600020900154101561044d5760016103d383600263ffffffff61055f16565b815481106103dd57fe5b90600052602060002090015490506001828154811015156103fa57fe5b600091825260209091200154600161041984600263ffffffff61055f16565b8154811061042357fe5b600091825260209091200155600180548291908490811061044057fe5b6000918252602090912001555b61045e82600263ffffffff61055f16565b915061035d565b5050565b60008282111561047557fe5b50900390565b6000805b60025461049660028561031590919063ffffffff16565b1161055a576104a4836100ea565b91506001828154811015156104b557fe5b9060005260206000209001546001848154811015156104d057fe5b90600052602060002090015411156105525760018054849081106104f057fe5b906000526020600020900154905060018281548110151561050d57fe5b90600052602060002090015460018481548110151561052857fe5b600091825260209091200155600180548291908490811061054557fe5b6000918252602090912001555b81925061047f565b505050565b600080828481151561056d57fe5b04949350505050565b81548183558181151161055a5760008381526020902061055a91810190830161031291905b808211156102e8576000815560010161059b5600a165627a7a72305820a19df8a356e5aa336313e2f2c05c7d43b27a533f8d2c532f56b4e36c163fff380029`

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
	return address, tx, &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}}, nil
}

// RootChain is an auto generated Go binding around an Ethereum contract.
type RootChain struct {
	RootChainCaller     // Read-only binding to the contract
	RootChainTransactor // Write-only binding to the contract
}

// RootChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainTransactor struct {
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
	contract, err := bindRootChain(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}}, nil
}

// NewRootChainCaller creates a new read-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainCaller(address common.Address, caller bind.ContractCaller) (*RootChainCaller, error) {
	contract, err := bindRootChain(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainCaller{contract: contract}, nil
}

// NewRootChainTransactor creates a new write-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainTransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainTransactor, error) {
	contract, err := bindRootChain(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RootChainTransactor{contract: contract}, nil
}

// bindRootChain binds a generic wrapper to an already deployed contract.
func bindRootChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, created_at uint256)
func (_RootChain *RootChainCaller) ChildChain(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Root       [32]byte
	Created_at *big.Int
}, error) {
	ret := new(struct {
		Root       [32]byte
		Created_at *big.Int
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "childChain", arg0)
	return *ret, err
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, created_at uint256)
func (_RootChain *RootChainSession) ChildChain(arg0 *big.Int) (struct {
	Root       [32]byte
	Created_at *big.Int
}, error) {
	return _RootChain.Contract.ChildChain(&_RootChain.CallOpts, arg0)
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, created_at uint256)
func (_RootChain *RootChainCallerSession) ChildChain(arg0 *big.Int) (struct {
	Root       [32]byte
	Created_at *big.Int
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

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(txBytes bytes) returns()
func (_RootChain *RootChainTransactor) Deposit(opts *bind.TransactOpts, txBytes []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "deposit", txBytes)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(txBytes bytes) returns()
func (_RootChain *RootChainSession) Deposit(txBytes []byte) (*types.Transaction, error) {
	return _RootChain.Contract.Deposit(&_RootChain.TransactOpts, txBytes)
}

// Deposit is a paid mutator transaction binding the contract method 0x98b1e06a.
//
// Solidity: function deposit(txBytes bytes) returns()
func (_RootChain *RootChainTransactorSession) Deposit(txBytes []byte) (*types.Transaction, error) {
	return _RootChain.Contract.Deposit(&_RootChain.TransactOpts, txBytes)
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

// SubmitBlock is a paid mutator transaction binding the contract method 0x46ab67cb.
//
// Solidity: function submitBlock(root bytes32, blknum uint256) returns()
func (_RootChain *RootChainTransactor) SubmitBlock(opts *bind.TransactOpts, root [32]byte, blknum *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitBlock", root, blknum)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0x46ab67cb.
//
// Solidity: function submitBlock(root bytes32, blknum uint256) returns()
func (_RootChain *RootChainSession) SubmitBlock(root [32]byte, blknum *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitBlock(&_RootChain.TransactOpts, root, blknum)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0x46ab67cb.
//
// Solidity: function submitBlock(root bytes32, blknum uint256) returns()
func (_RootChain *RootChainTransactorSession) SubmitBlock(root [32]byte, blknum *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitBlock(&_RootChain.TransactOpts, root, blknum)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a72305820793ad75a41d277276db656c6290ddabf98fffbbf89e3f6b9d46e96ef15b6bdb70029`

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
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
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
	contract, err := bindSafeMath(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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
const ValidateBin = `0x60606040523415600e57600080fd5b603580601b6000396000f3006060604052600080fd00a165627a7a72305820ece3262425fe387813b8a61683081fb1e300e8a5335c6464347948657706559d0029`

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
	return address, tx, &Validate{ValidateCaller: ValidateCaller{contract: contract}, ValidateTransactor: ValidateTransactor{contract: contract}}, nil
}

// Validate is an auto generated Go binding around an Ethereum contract.
type Validate struct {
	ValidateCaller     // Read-only binding to the contract
	ValidateTransactor // Write-only binding to the contract
}

// ValidateCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidateTransactor struct {
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
	contract, err := bindValidate(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validate{ValidateCaller: ValidateCaller{contract: contract}, ValidateTransactor: ValidateTransactor{contract: contract}}, nil
}

// NewValidateCaller creates a new read-only instance of Validate, bound to a specific deployed contract.
func NewValidateCaller(address common.Address, caller bind.ContractCaller) (*ValidateCaller, error) {
	contract, err := bindValidate(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ValidateCaller{contract: contract}, nil
}

// NewValidateTransactor creates a new write-only instance of Validate, bound to a specific deployed contract.
func NewValidateTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidateTransactor, error) {
	contract, err := bindValidate(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ValidateTransactor{contract: contract}, nil
}

// bindValidate binds a generic wrapper to an already deployed contract.
func bindValidate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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

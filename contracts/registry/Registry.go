// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registry

import (
	"math/big"
	"strings"

	"github.com/klaytn/klaytn"
	"github.com/klaytn/klaytn/accounts/abi/bind"
	"github.com/klaytn/klaytn/blockchain/types"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = klaytn.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)
// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"activationBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Activated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"}],\"name\":\"ConstructContract\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"deprecateBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Deprecated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"prevName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"replaceBlockNumber\",\"type\":\"uint256\"}],\"name\":\"Replaced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"UpdateGovernance\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"activationBlockNumber\",\"type\":\"uint256\"}],\"name\":\"activate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"constructContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"contractNames\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"deprecateBlockNumber\",\"type\":\"uint256\"}],\"name\":\"deprecate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllContractNames\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getContractIfActive\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"enumIRegistry.State\",\"name\":\"state\",\"type\":\"uint8\"}],\"name\":\"readAllContractsAtGivenState\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"activation\",\"type\":\"bool\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"activationBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deprecateBlockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prevName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"replaceBlockNumber\",\"type\":\"uint256\"}],\"name\":\"replace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"stateAt\",\"outputs\":[{\"internalType\":\"enumIRegistry.State\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newGovernance\",\"type\":\"address\"}],\"name\":\"updateGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

const RegistryBinRuntime = "0x608060"

const RegistryContractAddress = ""

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// ContractNames is a free data retrieval call binding the contract method 0x3ca6bb92.
//
// Solidity: function contractNames(uint256 ) view returns(string)
func (_Registry *RegistryCaller) ContractNames(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	ret0 := new(string)
	out := ret0
	err := _Registry.contract.Call(opts, out, "contractNames", arg0)
	return *ret0, err
}

// ContractNames is a free data retrieval call binding the contract method 0x3ca6bb92.
//
// Solidity: function contractNames(uint256 ) view returns(string)
func (_Registry *RegistrySession) ContractNames(arg0 *big.Int) (string, error) {
	return _Registry.Contract.ContractNames(&_Registry.CallOpts, arg0)
}

// ContractNames is a free data retrieval call binding the contract method 0x3ca6bb92.
//
// Solidity: function contractNames(uint256 ) view returns(string)
func (_Registry *RegistryCallerSession) ContractNames(arg0 *big.Int) (string, error) {
	return _Registry.Contract.ContractNames(&_Registry.CallOpts, arg0)
}

// GetAllContractNames is a free data retrieval call binding the contract method 0xd7a9b816.
//
// Solidity: function getAllContractNames() view returns(string[])
func (_Registry *RegistryCaller) GetAllContractNames(opts *bind.CallOpts) ([]string, error) {
	ret0 := new([]string)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getAllContractNames")
	return *ret0, err
}

// GetAllContractNames is a free data retrieval call binding the contract method 0xd7a9b816.
//
// Solidity: function getAllContractNames() view returns(string[])
func (_Registry *RegistrySession) GetAllContractNames() ([]string, error) {
	return _Registry.Contract.GetAllContractNames(&_Registry.CallOpts)
}

// GetAllContractNames is a free data retrieval call binding the contract method 0xd7a9b816.
//
// Solidity: function getAllContractNames() view returns(string[])
func (_Registry *RegistryCallerSession) GetAllContractNames() ([]string, error) {
	return _Registry.Contract.GetAllContractNames(&_Registry.CallOpts)
}

// GetContractIfActive is a free data retrieval call binding the contract method 0x35b5cfbb.
//
// Solidity: function getContractIfActive(string name) view returns(address)
func (_Registry *RegistryCaller) GetContractIfActive(opts *bind.CallOpts, name string) (common.Address, error) {
	ret0 := new(common.Address)
	out := ret0
	err := _Registry.contract.Call(opts, out, "getContractIfActive", name)
	return *ret0, err
}

// GetContractIfActive is a free data retrieval call binding the contract method 0x35b5cfbb.
//
// Solidity: function getContractIfActive(string name) view returns(address)
func (_Registry *RegistrySession) GetContractIfActive(name string) (common.Address, error) {
	return _Registry.Contract.GetContractIfActive(&_Registry.CallOpts, name)
}

// GetContractIfActive is a free data retrieval call binding the contract method 0x35b5cfbb.
//
// Solidity: function getContractIfActive(string name) view returns(address)
func (_Registry *RegistryCallerSession) GetContractIfActive(name string) (common.Address, error) {
	return _Registry.Contract.GetContractIfActive(&_Registry.CallOpts, name)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Registry *RegistryCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	ret0 := new(common.Address)
	out := ret0
	err := _Registry.contract.Call(opts, out, "governance")
	return *ret0, err
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Registry *RegistrySession) Governance() (common.Address, error) {
	return _Registry.Contract.Governance(&_Registry.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Registry *RegistryCallerSession) Governance() (common.Address, error) {
	return _Registry.Contract.Governance(&_Registry.CallOpts)
}

// ReadAllContractsAtGivenState is a free data retrieval call binding the contract method 0xfa96b37f.
//
// Solidity: function readAllContractsAtGivenState(uint256 blockNumber, uint8 state) view returns(string[], address[])
func (_Registry *RegistryCaller) ReadAllContractsAtGivenState(opts *bind.CallOpts, blockNumber *big.Int, state uint8) ([]string, []common.Address, error) {
	var (
		ret0 = new([]string)
		ret1 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Registry.contract.Call(opts, out, "readAllContractsAtGivenState", blockNumber, state)

	return *ret0, *ret1, err
}

// ReadAllContractsAtGivenState is a free data retrieval call binding the contract method 0xfa96b37f.
//
// Solidity: function readAllContractsAtGivenState(uint256 blockNumber, uint8 state) view returns(string[], address[])
func (_Registry *RegistrySession) ReadAllContractsAtGivenState(blockNumber *big.Int, state uint8) ([]string, []common.Address, error) {
	return _Registry.Contract.ReadAllContractsAtGivenState(&_Registry.CallOpts, blockNumber, state)
}

// ReadAllContractsAtGivenState is a free data retrieval call binding the contract method 0xfa96b37f.
//
// Solidity: function readAllContractsAtGivenState(uint256 blockNumber, uint8 state) view returns(string[], address[])
func (_Registry *RegistryCallerSession) ReadAllContractsAtGivenState(blockNumber *big.Int, state uint8) ([]string, []common.Address, error) {
	return _Registry.Contract.ReadAllContractsAtGivenState(&_Registry.CallOpts, blockNumber, state)
}

// Registry is a free data retrieval call binding the contract method 0x92a296c9.
//
// Solidity: function registry(string ) view returns(address addr, uint256 activationBlockNumber, uint256 deprecateBlockNumber)
func (_Registry *RegistryCaller) Registry(opts *bind.CallOpts, arg0 string) (struct {
	Addr                  common.Address
	ActivationBlockNumber *big.Int
	DeprecateBlockNumber  *big.Int
}, error,
) {
	ret := new(struct {
		Addr                  common.Address
		ActivationBlockNumber *big.Int
		DeprecateBlockNumber  *big.Int
	})
	out := ret
	err := _Registry.contract.Call(opts, out, "registry", arg0)
	return *ret, err
}

// Registry is a free data retrieval call binding the contract method 0x92a296c9.
//
// Solidity: function registry(string ) view returns(address addr, uint256 activationBlockNumber, uint256 deprecateBlockNumber)
func (_Registry *RegistrySession) Registry(arg0 string) (struct {
	Addr                  common.Address
	ActivationBlockNumber *big.Int
	DeprecateBlockNumber  *big.Int
}, error) {
	return _Registry.Contract.Registry(&_Registry.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0x92a296c9.
//
// Solidity: function registry(string ) view returns(address addr, uint256 activationBlockNumber, uint256 deprecateBlockNumber)
func (_Registry *RegistryCallerSession) Registry(arg0 string) (struct {
	Addr                  common.Address
	ActivationBlockNumber *big.Int
	DeprecateBlockNumber  *big.Int
}, error) {
	return _Registry.Contract.Registry(&_Registry.CallOpts, arg0)
}

// StateAt is a free data retrieval call binding the contract method 0x05e4f16e.
//
// Solidity: function stateAt(string name, uint256 blockNumber) view returns(uint8)
func (_Registry *RegistryCaller) StateAt(opts *bind.CallOpts, name string, blockNumber *big.Int) (uint8, error) {
	ret0 := new(uint8)
	out := ret0
	err := _Registry.contract.Call(opts, out, "stateAt", name, blockNumber)
	return *ret0, err
}

// StateAt is a free data retrieval call binding the contract method 0x05e4f16e.
//
// Solidity: function stateAt(string name, uint256 blockNumber) view returns(uint8)
func (_Registry *RegistrySession) StateAt(name string, blockNumber *big.Int) (uint8, error) {
	return _Registry.Contract.StateAt(&_Registry.CallOpts, name, blockNumber)
}

// StateAt is a free data retrieval call binding the contract method 0x05e4f16e.
//
// Solidity: function stateAt(string name, uint256 blockNumber) view returns(uint8)
func (_Registry *RegistryCallerSession) StateAt(name string, blockNumber *big.Int) (uint8, error) {
	return _Registry.Contract.StateAt(&_Registry.CallOpts, name, blockNumber)
}

// Activate is a paid mutator transaction binding the contract method 0x51afb6a3.
//
// Solidity: function activate(string name, uint256 activationBlockNumber) returns()
func (_Registry *RegistryTransactor) Activate(opts *bind.TransactOpts, name string, activationBlockNumber *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "activate", name, activationBlockNumber)
}

// Activate is a paid mutator transaction binding the contract method 0x51afb6a3.
//
// Solidity: function activate(string name, uint256 activationBlockNumber) returns()
func (_Registry *RegistrySession) Activate(name string, activationBlockNumber *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Activate(&_Registry.TransactOpts, name, activationBlockNumber)
}

// Activate is a paid mutator transaction binding the contract method 0x51afb6a3.
//
// Solidity: function activate(string name, uint256 activationBlockNumber) returns()
func (_Registry *RegistryTransactorSession) Activate(name string, activationBlockNumber *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Activate(&_Registry.TransactOpts, name, activationBlockNumber)
}

// ConstructContract is a paid mutator transaction binding the contract method 0xe14c672d.
//
// Solidity: function constructContract(address _governance) returns()
func (_Registry *RegistryTransactor) ConstructContract(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "constructContract", _governance)
}

// ConstructContract is a paid mutator transaction binding the contract method 0xe14c672d.
//
// Solidity: function constructContract(address _governance) returns()
func (_Registry *RegistrySession) ConstructContract(_governance common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ConstructContract(&_Registry.TransactOpts, _governance)
}

// ConstructContract is a paid mutator transaction binding the contract method 0xe14c672d.
//
// Solidity: function constructContract(address _governance) returns()
func (_Registry *RegistryTransactorSession) ConstructContract(_governance common.Address) (*types.Transaction, error) {
	return _Registry.Contract.ConstructContract(&_Registry.TransactOpts, _governance)
}

// Deprecate is a paid mutator transaction binding the contract method 0x4d0e6606.
//
// Solidity: function deprecate(string name, uint256 deprecateBlockNumber) returns()
func (_Registry *RegistryTransactor) Deprecate(opts *bind.TransactOpts, name string, deprecateBlockNumber *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "deprecate", name, deprecateBlockNumber)
}

// Deprecate is a paid mutator transaction binding the contract method 0x4d0e6606.
//
// Solidity: function deprecate(string name, uint256 deprecateBlockNumber) returns()
func (_Registry *RegistrySession) Deprecate(name string, deprecateBlockNumber *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Deprecate(&_Registry.TransactOpts, name, deprecateBlockNumber)
}

// Deprecate is a paid mutator transaction binding the contract method 0x4d0e6606.
//
// Solidity: function deprecate(string name, uint256 deprecateBlockNumber) returns()
func (_Registry *RegistryTransactorSession) Deprecate(name string, deprecateBlockNumber *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Deprecate(&_Registry.TransactOpts, name, deprecateBlockNumber)
}

// Register is a paid mutator transaction binding the contract method 0x60d7a278.
//
// Solidity: function register(string name, address addr, bool activation) returns()
func (_Registry *RegistryTransactor) Register(opts *bind.TransactOpts, name string, addr common.Address, activation bool) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "register", name, addr, activation)
}

// Register is a paid mutator transaction binding the contract method 0x60d7a278.
//
// Solidity: function register(string name, address addr, bool activation) returns()
func (_Registry *RegistrySession) Register(name string, addr common.Address, activation bool) (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts, name, addr, activation)
}

// Register is a paid mutator transaction binding the contract method 0x60d7a278.
//
// Solidity: function register(string name, address addr, bool activation) returns()
func (_Registry *RegistryTransactorSession) Register(name string, addr common.Address, activation bool) (*types.Transaction, error) {
	return _Registry.Contract.Register(&_Registry.TransactOpts, name, addr, activation)
}

// Replace is a paid mutator transaction binding the contract method 0xe47f1160.
//
// Solidity: function replace(string prevName, string newName, uint256 replaceBlockNumber) returns()
func (_Registry *RegistryTransactor) Replace(opts *bind.TransactOpts, prevName string, newName string, replaceBlockNumber *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "replace", prevName, newName, replaceBlockNumber)
}

// Replace is a paid mutator transaction binding the contract method 0xe47f1160.
//
// Solidity: function replace(string prevName, string newName, uint256 replaceBlockNumber) returns()
func (_Registry *RegistrySession) Replace(prevName string, newName string, replaceBlockNumber *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Replace(&_Registry.TransactOpts, prevName, newName, replaceBlockNumber)
}

// Replace is a paid mutator transaction binding the contract method 0xe47f1160.
//
// Solidity: function replace(string prevName, string newName, uint256 replaceBlockNumber) returns()
func (_Registry *RegistryTransactorSession) Replace(prevName string, newName string, replaceBlockNumber *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.Replace(&_Registry.TransactOpts, prevName, newName, replaceBlockNumber)
}

// UpdateGovernance is a paid mutator transaction binding the contract method 0xb2561263.
//
// Solidity: function updateGovernance(address _newGovernance) returns()
func (_Registry *RegistryTransactor) UpdateGovernance(opts *bind.TransactOpts, _newGovernance common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "updateGovernance", _newGovernance)
}

// UpdateGovernance is a paid mutator transaction binding the contract method 0xb2561263.
//
// Solidity: function updateGovernance(address _newGovernance) returns()
func (_Registry *RegistrySession) UpdateGovernance(_newGovernance common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UpdateGovernance(&_Registry.TransactOpts, _newGovernance)
}

// UpdateGovernance is a paid mutator transaction binding the contract method 0xb2561263.
//
// Solidity: function updateGovernance(address _newGovernance) returns()
func (_Registry *RegistryTransactorSession) UpdateGovernance(_newGovernance common.Address) (*types.Transaction, error) {
	return _Registry.Contract.UpdateGovernance(&_Registry.TransactOpts, _newGovernance)
}

// RegistryActivatedIterator is returned from FilterActivated and is used to iterate over the raw logs and unpacked data for Activated events raised by the Registry contract.
type RegistryActivatedIterator struct {
	Event *RegistryActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryActivated)
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
		it.Event = new(RegistryActivated)
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
func (it *RegistryActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryActivated represents a Activated event raised by the Registry contract.
type RegistryActivated struct {
	Name                  string
	ActivationBlockNumber *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterActivated is a free log retrieval operation binding the contract event 0x46ca43c7152dee1b324184afa3e3edec374f29a3880b6189ec2106c94c3f4b5d.
//
// Solidity: event Activated(string name, uint256 indexed activationBlockNumber)
func (_Registry *RegistryFilterer) FilterActivated(opts *bind.FilterOpts, activationBlockNumber []*big.Int) (*RegistryActivatedIterator, error) {

	var activationBlockNumberRule []interface{}
	for _, activationBlockNumberItem := range activationBlockNumber {
		activationBlockNumberRule = append(activationBlockNumberRule, activationBlockNumberItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Activated", activationBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &RegistryActivatedIterator{contract: _Registry.contract, event: "Activated", logs: logs, sub: sub}, nil
}

// WatchActivated is a free log subscription operation binding the contract event 0x46ca43c7152dee1b324184afa3e3edec374f29a3880b6189ec2106c94c3f4b5d.
//
// Solidity: event Activated(string name, uint256 indexed activationBlockNumber)
func (_Registry *RegistryFilterer) WatchActivated(opts *bind.WatchOpts, sink chan<- *RegistryActivated, activationBlockNumber []*big.Int) (event.Subscription, error) {

	var activationBlockNumberRule []interface{}
	for _, activationBlockNumberItem := range activationBlockNumber {
		activationBlockNumberRule = append(activationBlockNumberRule, activationBlockNumberItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Activated", activationBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryActivated)
				if err := _Registry.contract.UnpackLog(event, "Activated", log); err != nil {
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

// ParseActivated is a log parse operation binding the contract event 0x46ca43c7152dee1b324184afa3e3edec374f29a3880b6189ec2106c94c3f4b5d.
//
// Solidity: event Activated(string name, uint256 indexed activationBlockNumber)
func (_Registry *RegistryFilterer) ParseActivated(log types.Log) (*RegistryActivated, error) {
	event := new(RegistryActivated)
	if err := _Registry.contract.UnpackLog(event, "Activated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryConstructContractIterator is returned from FilterConstructContract and is used to iterate over the raw logs and unpacked data for ConstructContract events raised by the Registry contract.
type RegistryConstructContractIterator struct {
	Event *RegistryConstructContract // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryConstructContractIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryConstructContract)
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
		it.Event = new(RegistryConstructContract)
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
func (it *RegistryConstructContractIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryConstructContractIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryConstructContract represents a ConstructContract event raised by the Registry contract.
type RegistryConstructContract struct {
	Governance common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConstructContract is a free log retrieval operation binding the contract event 0x15755ba3b1deec9fcf75c3db96f2d79482964aea97c6da27ceca7411850e4054.
//
// Solidity: event ConstructContract(address indexed governance)
func (_Registry *RegistryFilterer) FilterConstructContract(opts *bind.FilterOpts, governance []common.Address) (*RegistryConstructContractIterator, error) {

	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ConstructContract", governanceRule)
	if err != nil {
		return nil, err
	}
	return &RegistryConstructContractIterator{contract: _Registry.contract, event: "ConstructContract", logs: logs, sub: sub}, nil
}

// WatchConstructContract is a free log subscription operation binding the contract event 0x15755ba3b1deec9fcf75c3db96f2d79482964aea97c6da27ceca7411850e4054.
//
// Solidity: event ConstructContract(address indexed governance)
func (_Registry *RegistryFilterer) WatchConstructContract(opts *bind.WatchOpts, sink chan<- *RegistryConstructContract, governance []common.Address) (event.Subscription, error) {

	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ConstructContract", governanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryConstructContract)
				if err := _Registry.contract.UnpackLog(event, "ConstructContract", log); err != nil {
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

// ParseConstructContract is a log parse operation binding the contract event 0x15755ba3b1deec9fcf75c3db96f2d79482964aea97c6da27ceca7411850e4054.
//
// Solidity: event ConstructContract(address indexed governance)
func (_Registry *RegistryFilterer) ParseConstructContract(log types.Log) (*RegistryConstructContract, error) {
	event := new(RegistryConstructContract)
	if err := _Registry.contract.UnpackLog(event, "ConstructContract", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryDeprecatedIterator is returned from FilterDeprecated and is used to iterate over the raw logs and unpacked data for Deprecated events raised by the Registry contract.
type RegistryDeprecatedIterator struct {
	Event *RegistryDeprecated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryDeprecated)
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
		it.Event = new(RegistryDeprecated)
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
func (it *RegistryDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryDeprecated represents a Deprecated event raised by the Registry contract.
type RegistryDeprecated struct {
	Name                 string
	DeprecateBlockNumber *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDeprecated is a free log retrieval operation binding the contract event 0x058d0e8a83e7a5228ca32808c88378cac5ff2eb4b3ef45a109be564227420c05.
//
// Solidity: event Deprecated(string name, uint256 indexed deprecateBlockNumber)
func (_Registry *RegistryFilterer) FilterDeprecated(opts *bind.FilterOpts, deprecateBlockNumber []*big.Int) (*RegistryDeprecatedIterator, error) {

	var deprecateBlockNumberRule []interface{}
	for _, deprecateBlockNumberItem := range deprecateBlockNumber {
		deprecateBlockNumberRule = append(deprecateBlockNumberRule, deprecateBlockNumberItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Deprecated", deprecateBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &RegistryDeprecatedIterator{contract: _Registry.contract, event: "Deprecated", logs: logs, sub: sub}, nil
}

// WatchDeprecated is a free log subscription operation binding the contract event 0x058d0e8a83e7a5228ca32808c88378cac5ff2eb4b3ef45a109be564227420c05.
//
// Solidity: event Deprecated(string name, uint256 indexed deprecateBlockNumber)
func (_Registry *RegistryFilterer) WatchDeprecated(opts *bind.WatchOpts, sink chan<- *RegistryDeprecated, deprecateBlockNumber []*big.Int) (event.Subscription, error) {

	var deprecateBlockNumberRule []interface{}
	for _, deprecateBlockNumberItem := range deprecateBlockNumber {
		deprecateBlockNumberRule = append(deprecateBlockNumberRule, deprecateBlockNumberItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Deprecated", deprecateBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryDeprecated)
				if err := _Registry.contract.UnpackLog(event, "Deprecated", log); err != nil {
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

// ParseDeprecated is a log parse operation binding the contract event 0x058d0e8a83e7a5228ca32808c88378cac5ff2eb4b3ef45a109be564227420c05.
//
// Solidity: event Deprecated(string name, uint256 indexed deprecateBlockNumber)
func (_Registry *RegistryFilterer) ParseDeprecated(log types.Log) (*RegistryDeprecated, error) {
	event := new(RegistryDeprecated)
	if err := _Registry.contract.UnpackLog(event, "Deprecated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Registry contract.
type RegistryRegisteredIterator struct {
	Event *RegistryRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryRegistered)
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
		it.Event = new(RegistryRegistered)
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
func (it *RegistryRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryRegistered represents a Registered event raised by the Registry contract.
type RegistryRegistered struct {
	Name string
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x50f74ca45caac8020b8d891bd13ea5a2d79564986ee6a839f0d914896388322d.
//
// Solidity: event Registered(string name, address indexed addr)
func (_Registry *RegistryFilterer) FilterRegistered(opts *bind.FilterOpts, addr []common.Address) (*RegistryRegisteredIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Registered", addrRule)
	if err != nil {
		return nil, err
	}
	return &RegistryRegisteredIterator{contract: _Registry.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x50f74ca45caac8020b8d891bd13ea5a2d79564986ee6a839f0d914896388322d.
//
// Solidity: event Registered(string name, address indexed addr)
func (_Registry *RegistryFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *RegistryRegistered, addr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Registered", addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryRegistered)
				if err := _Registry.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x50f74ca45caac8020b8d891bd13ea5a2d79564986ee6a839f0d914896388322d.
//
// Solidity: event Registered(string name, address indexed addr)
func (_Registry *RegistryFilterer) ParseRegistered(log types.Log) (*RegistryRegistered, error) {
	event := new(RegistryRegistered)
	if err := _Registry.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryReplacedIterator is returned from FilterReplaced and is used to iterate over the raw logs and unpacked data for Replaced events raised by the Registry contract.
type RegistryReplacedIterator struct {
	Event *RegistryReplaced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryReplacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryReplaced)
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
		it.Event = new(RegistryReplaced)
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
func (it *RegistryReplacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryReplacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryReplaced represents a Replaced event raised by the Registry contract.
type RegistryReplaced struct {
	PrevName           string
	NewName            string
	ReplaceBlockNumber *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterReplaced is a free log retrieval operation binding the contract event 0xe56cea40c3f094f642625199671fa50eeee44fd97d16ce800adddce75019e956.
//
// Solidity: event Replaced(string prevName, string newName, uint256 indexed replaceBlockNumber)
func (_Registry *RegistryFilterer) FilterReplaced(opts *bind.FilterOpts, replaceBlockNumber []*big.Int) (*RegistryReplacedIterator, error) {

	var replaceBlockNumberRule []interface{}
	for _, replaceBlockNumberItem := range replaceBlockNumber {
		replaceBlockNumberRule = append(replaceBlockNumberRule, replaceBlockNumberItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "Replaced", replaceBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &RegistryReplacedIterator{contract: _Registry.contract, event: "Replaced", logs: logs, sub: sub}, nil
}

// WatchReplaced is a free log subscription operation binding the contract event 0xe56cea40c3f094f642625199671fa50eeee44fd97d16ce800adddce75019e956.
//
// Solidity: event Replaced(string prevName, string newName, uint256 indexed replaceBlockNumber)
func (_Registry *RegistryFilterer) WatchReplaced(opts *bind.WatchOpts, sink chan<- *RegistryReplaced, replaceBlockNumber []*big.Int) (event.Subscription, error) {

	var replaceBlockNumberRule []interface{}
	for _, replaceBlockNumberItem := range replaceBlockNumber {
		replaceBlockNumberRule = append(replaceBlockNumberRule, replaceBlockNumberItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "Replaced", replaceBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryReplaced)
				if err := _Registry.contract.UnpackLog(event, "Replaced", log); err != nil {
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

// ParseReplaced is a log parse operation binding the contract event 0xe56cea40c3f094f642625199671fa50eeee44fd97d16ce800adddce75019e956.
//
// Solidity: event Replaced(string prevName, string newName, uint256 indexed replaceBlockNumber)
func (_Registry *RegistryFilterer) ParseReplaced(log types.Log) (*RegistryReplaced, error) {
	event := new(RegistryReplaced)
	if err := _Registry.contract.UnpackLog(event, "Replaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryUpdateGovernanceIterator is returned from FilterUpdateGovernance and is used to iterate over the raw logs and unpacked data for UpdateGovernance events raised by the Registry contract.
type RegistryUpdateGovernanceIterator struct {
	Event *RegistryUpdateGovernance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  klaytn.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RegistryUpdateGovernanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryUpdateGovernance)
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
		it.Event = new(RegistryUpdateGovernance)
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
func (it *RegistryUpdateGovernanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryUpdateGovernanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryUpdateGovernance represents a UpdateGovernance event raised by the Registry contract.
type RegistryUpdateGovernance struct {
	NewGovernance common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateGovernance is a free log retrieval operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address indexed newGovernance)
func (_Registry *RegistryFilterer) FilterUpdateGovernance(opts *bind.FilterOpts, newGovernance []common.Address) (*RegistryUpdateGovernanceIterator, error) {

	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "UpdateGovernance", newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &RegistryUpdateGovernanceIterator{contract: _Registry.contract, event: "UpdateGovernance", logs: logs, sub: sub}, nil
}

// WatchUpdateGovernance is a free log subscription operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address indexed newGovernance)
func (_Registry *RegistryFilterer) WatchUpdateGovernance(opts *bind.WatchOpts, sink chan<- *RegistryUpdateGovernance, newGovernance []common.Address) (event.Subscription, error) {

	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "UpdateGovernance", newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryUpdateGovernance)
				if err := _Registry.contract.UnpackLog(event, "UpdateGovernance", log); err != nil {
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

// ParseUpdateGovernance is a log parse operation binding the contract event 0x8d55d160c0009eb3d739442df0a3ca089ed64378bfac017e7ddad463f9815b87.
//
// Solidity: event UpdateGovernance(address indexed newGovernance)
func (_Registry *RegistryFilterer) ParseUpdateGovernance(log types.Log) (*RegistryUpdateGovernance, error) {
	event := new(RegistryUpdateGovernance)
	if err := _Registry.contract.UnpackLog(event, "UpdateGovernance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

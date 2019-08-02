// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package userGroup

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AddrToBoolABI is the input ABI used to generate the binding from.
const AddrToBoolABI = "[]"

// AddrToBoolBin is the compiled bytecode used for deploying new contracts.
const AddrToBoolBin = `0x6104b7610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100875760003560e01c806398b3fbd71161006557806398b3fbd71461012b578063cd06a7891461016c578063d89a75b81461018f578063da3f9426146101d557610087565b80633c6039cf1461008c578063448e4174146100d95780639601640a14610108575b600080fd5b81801561009857600080fd5b506100c5600480360360408110156100af57600080fd5b50803590602001356001600160a01b0316610201565b604080519115158252519081900360200190f35b6100f6600480360360208110156100ef57600080fd5b50356102a0565b60408051918252519081900360200190f35b6100f66004803603604081101561011e57600080fd5b50803590602001356102aa565b81801561013757600080fd5b506100c56004803603606081101561014e57600080fd5b508035906001600160a01b03602082013516906040013515156102fa565b6100c56004803603604081101561018257600080fd5b50803590602001356103c0565b6101b2600480360360408110156101a557600080fd5b50803590602001356103cb565b604080516001600160a01b03909316835290151560208301528051918290030190f35b6100c5600480360360408110156101eb57600080fd5b50803590602001356001600160a01b0316610410565b6001600160a01b0381166000908152602083905260408120548061022957600091505061029a565b6001600160a01b03831660009081526020859052604081209081556001908101805460ff191690558085018054600019840190811061026457fe5b60009182526020909120018054911515600160a01b0260ff60a01b19909216919091179055505060028201805460001901905560015b92915050565b600061029a826000195b60010160005b6001830154821080156102e457508260010182815481106102cd57fe5b600091825260209091200154600160a01b900460ff165b156102f4576001909101906102b0565b50919050565b6001600160a01b038216600090815260208490526040812080546001909101805460ff191684151517905580156103355760019150506103b9565b60018086018054916103499190830161042e565b6001600160a01b038516600090815260208790526040902060018083019091558601805491925085918390811061037c57fe5b6000918252602082200180546001600160a01b0319166001600160a01b039390931692909217909155600286018054600101905591506103b99050565b9392505050565b600191909101541190565b6000808360010183815481106103dd57fe5b60009182526020808320909101546001600160a01b03168083529590526040902060010154939460ff9094169392505050565b6001600160a01b031660009081526020919091526040902054151590565b81548183558181111561045257600083815260209020610452918101908301610457565b505050565b61047f91905b8082111561047b5780546001600160a81b031916815560010161045d565b5090565b9056fea265627a7a72305820b0290c518bba8c4019a32511e7d2b37f9ded8fc15be6c516ada3e900e8328a2b64736f6c634300050a0032`

// DeployAddrToBool deploys a new Ethereum contract, binding an instance of AddrToBool to it.
func DeployAddrToBool(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddrToBool, error) {
	parsed, err := abi.JSON(strings.NewReader(AddrToBoolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddrToBoolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddrToBool{AddrToBoolCaller: AddrToBoolCaller{contract: contract}, AddrToBoolTransactor: AddrToBoolTransactor{contract: contract}, AddrToBoolFilterer: AddrToBoolFilterer{contract: contract}}, nil
}

// AddrToBool is an auto generated Go binding around an Ethereum contract.
type AddrToBool struct {
	AddrToBoolCaller     // Read-only binding to the contract
	AddrToBoolTransactor // Write-only binding to the contract
	AddrToBoolFilterer   // Log filterer for contract events
}

// AddrToBoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddrToBoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrToBoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddrToBoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrToBoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddrToBoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrToBoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddrToBoolSession struct {
	Contract     *AddrToBool       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddrToBoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddrToBoolCallerSession struct {
	Contract *AddrToBoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AddrToBoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddrToBoolTransactorSession struct {
	Contract     *AddrToBoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AddrToBoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddrToBoolRaw struct {
	Contract *AddrToBool // Generic contract binding to access the raw methods on
}

// AddrToBoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddrToBoolCallerRaw struct {
	Contract *AddrToBoolCaller // Generic read-only contract binding to access the raw methods on
}

// AddrToBoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddrToBoolTransactorRaw struct {
	Contract *AddrToBoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddrToBool creates a new instance of AddrToBool, bound to a specific deployed contract.
func NewAddrToBool(address common.Address, backend bind.ContractBackend) (*AddrToBool, error) {
	contract, err := bindAddrToBool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddrToBool{AddrToBoolCaller: AddrToBoolCaller{contract: contract}, AddrToBoolTransactor: AddrToBoolTransactor{contract: contract}, AddrToBoolFilterer: AddrToBoolFilterer{contract: contract}}, nil
}

// NewAddrToBoolCaller creates a new read-only instance of AddrToBool, bound to a specific deployed contract.
func NewAddrToBoolCaller(address common.Address, caller bind.ContractCaller) (*AddrToBoolCaller, error) {
	contract, err := bindAddrToBool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddrToBoolCaller{contract: contract}, nil
}

// NewAddrToBoolTransactor creates a new write-only instance of AddrToBool, bound to a specific deployed contract.
func NewAddrToBoolTransactor(address common.Address, transactor bind.ContractTransactor) (*AddrToBoolTransactor, error) {
	contract, err := bindAddrToBool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddrToBoolTransactor{contract: contract}, nil
}

// NewAddrToBoolFilterer creates a new log filterer instance of AddrToBool, bound to a specific deployed contract.
func NewAddrToBoolFilterer(address common.Address, filterer bind.ContractFilterer) (*AddrToBoolFilterer, error) {
	contract, err := bindAddrToBool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddrToBoolFilterer{contract: contract}, nil
}

// bindAddrToBool binds a generic wrapper to an already deployed contract.
func bindAddrToBool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddrToBoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddrToBool *AddrToBoolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddrToBool.Contract.AddrToBoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddrToBool *AddrToBoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddrToBool.Contract.AddrToBoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddrToBool *AddrToBoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddrToBool.Contract.AddrToBoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddrToBool *AddrToBoolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddrToBool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddrToBool *AddrToBoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddrToBool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddrToBool *AddrToBoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddrToBool.Contract.contract.Transact(opts, method, params...)
}

// UserGroupABI is the input ABI used to generate the binding from.
const UserGroupABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"address\"},{\"name\":\"permisssion\",\"type\":\"uint8\"}],\"name\":\"andOrupdateUser\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deleted\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"address\"}],\"name\":\"getUserPerssion\",\"outputs\":[{\"name\":\"read\",\"type\":\"bool\"},{\"name\":\"update\",\"type\":\"bool\"},{\"name\":\"change\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"changeStatusOfUserGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"key\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"read\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"update\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"change\",\"type\":\"bool\"}],\"name\":\"_updateUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"_changeStatusOfUserGroup\",\"type\":\"event\"}]"

// UserGroupBin is the compiled bytecode used for deploying new contracts.
const UserGroupBin = `0x60806040526000805460ff60a01b1916905534801561001d57600080fd5b506040516109463803806109468339818101604052602081101561004057600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556108d4806100726000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806302d05d3f1461005c5780634c767d0f146100805780636b35f7c1146100c3578063878f70c9146100cb5780639c4e9f7c14610113575b600080fd5b610064610134565b604080516001600160a01b039092168252519081900360200190f35b6100af6004803603604081101561009657600080fd5b5080356001600160a01b0316906020013560ff16610143565b604080519115158252519081900360200190f35b6100af6104a1565b6100f1600480360360208110156100e157600080fd5b50356001600160a01b03166104b1565b6040805193151584529115156020840152151582820152519081900360600190f35b6101326004803603602081101561012957600080fd5b5035151561074b565b005b6000546001600160a01b031681565b60408051636d1fca1360e11b815260076004820152336024820152905160009173__$1a5b05563ffbb1c6a29a7456f190d0d702$__9163da3f942691604480820192602092909190829003018186803b15801561019f57600080fd5b505af41580156101b3573d6000803e3d6000fd5b505050506040513d60208110156101c957600080fd5b5051806101e057506000546001600160a01b031633145b61021b5760405162461bcd60e51b815260040180806020018281038252602d815260200180610873602d913960400191505060405180910390fd5b60078260ff16118061022b575060005b156102385750600061049b565b600060018084161461024b57600061024e565b60015b90506000600180851614610263576000610266565b60015b9050600060018086161461027b57600061027e565b60015b604080516398b3fbd760e01b8152600160048201526001600160a01b03891660248201528515156044820152905191925073__$1a5b05563ffbb1c6a29a7456f190d0d702$__916398b3fbd791606480820192602092909190829003018186803b1580156102eb57600080fd5b505af41580156102ff573d6000803e3d6000fd5b505050506040513d602081101561031557600080fd5b5050604080516398b3fbd760e01b81526004818101526001600160a01b03881660248201528315156044820152905173__$1a5b05563ffbb1c6a29a7456f190d0d702$__916398b3fbd7916064808301926020929190829003018186803b15801561037f57600080fd5b505af4158015610393573d6000803e3d6000fd5b505050506040513d60208110156103a957600080fd5b5050604080516398b3fbd760e01b8152600760048201526001600160a01b03881660248201528215156044820152905173__$1a5b05563ffbb1c6a29a7456f190d0d702$__916398b3fbd7916064808301926020929190829003018186803b15801561041457600080fd5b505af4158015610428573d6000803e3d6000fd5b505050506040513d602081101561043e57600080fd5b5050604080516001600160a01b0388168152841515602082015283151581830152821515606082015290517f9894bb533b126d657a7922559cc34561135edb194cf395b47e1c19e826056bae9181900360800190a1600193505050505b92915050565b600054600160a01b900460ff1681565b60008054819081906001600160a01b03858116911614156104da57506001915081905080610744565b600054600160a01b900460ff161515600114156105005750600191508190506000610744565b60408051636d1fca1360e11b8152600160048201526001600160a01b0386166024820152905173__$1a5b05563ffbb1c6a29a7456f190d0d702$__9163da3f9426916044808301926020929190829003018186803b15801561056157600080fd5b505af4158015610575573d6000803e3d6000fd5b505050506040513d602081101561058b57600080fd5b50511515600114156105be576001600160a01b0384166000908152600160208190526040909120015460ff1692506105c3565b600092505b60408051636d1fca1360e11b81526004818101526001600160a01b0386166024820152905173__$1a5b05563ffbb1c6a29a7456f190d0d702$__9163da3f9426916044808301926020929190829003018186803b15801561062357600080fd5b505af4158015610637573d6000803e3d6000fd5b505050506040513d602081101561064d57600080fd5b505115156001141561067f576001600160a01b03841660009081526004602052604090206001015460ff169150610684565b600091505b60408051636d1fca1360e11b8152600760048201526001600160a01b0386166024820152905173__$1a5b05563ffbb1c6a29a7456f190d0d702$__9163da3f9426916044808301926020929190829003018186803b1580156106e557600080fd5b505af41580156106f9573d6000803e3d6000fd5b505050506040513d602081101561070f57600080fd5b505115156001141561074057506001600160a01b03831660009081526007602052604090206001015460ff16610744565b5060005b9193909250565b60408051636d1fca1360e11b815260076004820152336024820152905173__$1a5b05563ffbb1c6a29a7456f190d0d702$__9163da3f9426916044808301926020929190829003018186803b1580156107a357600080fd5b505af41580156107b7573d6000803e3d6000fd5b505050506040513d60208110156107cd57600080fd5b5051806107e457506000546001600160a01b031633145b61081f5760405162461bcd60e51b815260040180806020018281038252602d815260200180610873602d913960400191505060405180910390fd5b60008054821515600160a01b810260ff60a01b199092169190911790915560408051918252517f65539928111859040d829c780a5ecc7a052a3c160109013679271b2e5c46d88c9181900360200190a15056fe4f6e6c79207065726d69737373696f6e20757365722063616e2063616c6c20746869732066756e6374696f6e2ea265627a7a723058203d7ab454d7dc5c7892eeb0fa4062e76ee7f873099cbcd12e4ef08a36f7599cc864736f6c634300050a0032`

// DeployUserGroup deploys a new Ethereum contract, binding an instance of UserGroup to it.
func DeployUserGroup(auth *bind.TransactOpts, backend bind.ContractBackend, sender common.Address) (common.Address, *types.Transaction, *UserGroup, error) {
	parsed, err := abi.JSON(strings.NewReader(UserGroupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserGroupBin), backend, sender)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserGroup{UserGroupCaller: UserGroupCaller{contract: contract}, UserGroupTransactor: UserGroupTransactor{contract: contract}, UserGroupFilterer: UserGroupFilterer{contract: contract}}, nil
}

// UserGroup is an auto generated Go binding around an Ethereum contract.
type UserGroup struct {
	UserGroupCaller     // Read-only binding to the contract
	UserGroupTransactor // Write-only binding to the contract
	UserGroupFilterer   // Log filterer for contract events
}

// UserGroupCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserGroupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserGroupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserGroupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserGroupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserGroupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserGroupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserGroupSession struct {
	Contract     *UserGroup        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserGroupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserGroupCallerSession struct {
	Contract *UserGroupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// UserGroupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserGroupTransactorSession struct {
	Contract     *UserGroupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UserGroupRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserGroupRaw struct {
	Contract *UserGroup // Generic contract binding to access the raw methods on
}

// UserGroupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserGroupCallerRaw struct {
	Contract *UserGroupCaller // Generic read-only contract binding to access the raw methods on
}

// UserGroupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserGroupTransactorRaw struct {
	Contract *UserGroupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserGroup creates a new instance of UserGroup, bound to a specific deployed contract.
func NewUserGroup(address common.Address, backend bind.ContractBackend) (*UserGroup, error) {
	contract, err := bindUserGroup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserGroup{UserGroupCaller: UserGroupCaller{contract: contract}, UserGroupTransactor: UserGroupTransactor{contract: contract}, UserGroupFilterer: UserGroupFilterer{contract: contract}}, nil
}

// NewUserGroupCaller creates a new read-only instance of UserGroup, bound to a specific deployed contract.
func NewUserGroupCaller(address common.Address, caller bind.ContractCaller) (*UserGroupCaller, error) {
	contract, err := bindUserGroup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserGroupCaller{contract: contract}, nil
}

// NewUserGroupTransactor creates a new write-only instance of UserGroup, bound to a specific deployed contract.
func NewUserGroupTransactor(address common.Address, transactor bind.ContractTransactor) (*UserGroupTransactor, error) {
	contract, err := bindUserGroup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserGroupTransactor{contract: contract}, nil
}

// NewUserGroupFilterer creates a new log filterer instance of UserGroup, bound to a specific deployed contract.
func NewUserGroupFilterer(address common.Address, filterer bind.ContractFilterer) (*UserGroupFilterer, error) {
	contract, err := bindUserGroup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserGroupFilterer{contract: contract}, nil
}

// bindUserGroup binds a generic wrapper to an already deployed contract.
func bindUserGroup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserGroupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserGroup *UserGroupRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserGroup.Contract.UserGroupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserGroup *UserGroupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserGroup.Contract.UserGroupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserGroup *UserGroupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserGroup.Contract.UserGroupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserGroup *UserGroupCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserGroup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserGroup *UserGroupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserGroup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserGroup *UserGroupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserGroup.Contract.contract.Transact(opts, method, params...)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_UserGroup *UserGroupCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserGroup.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_UserGroup *UserGroupSession) Creator() (common.Address, error) {
	return _UserGroup.Contract.Creator(&_UserGroup.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_UserGroup *UserGroupCallerSession) Creator() (common.Address, error) {
	return _UserGroup.Contract.Creator(&_UserGroup.CallOpts)
}

// Deleted is a free data retrieval call binding the contract method 0x6b35f7c1.
//
// Solidity: function deleted() constant returns(bool)
func (_UserGroup *UserGroupCaller) Deleted(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UserGroup.contract.Call(opts, out, "deleted")
	return *ret0, err
}

// Deleted is a free data retrieval call binding the contract method 0x6b35f7c1.
//
// Solidity: function deleted() constant returns(bool)
func (_UserGroup *UserGroupSession) Deleted() (bool, error) {
	return _UserGroup.Contract.Deleted(&_UserGroup.CallOpts)
}

// Deleted is a free data retrieval call binding the contract method 0x6b35f7c1.
//
// Solidity: function deleted() constant returns(bool)
func (_UserGroup *UserGroupCallerSession) Deleted() (bool, error) {
	return _UserGroup.Contract.Deleted(&_UserGroup.CallOpts)
}

// GetUserPerssion is a free data retrieval call binding the contract method 0x878f70c9.
//
// Solidity: function getUserPerssion(address key) constant returns(bool read, bool update, bool change)
func (_UserGroup *UserGroupCaller) GetUserPerssion(opts *bind.CallOpts, key common.Address) (struct {
	Read   bool
	Update bool
	Change bool
}, error) {
	ret := new(struct {
		Read   bool
		Update bool
		Change bool
	})
	out := ret
	err := _UserGroup.contract.Call(opts, out, "getUserPerssion", key)
	return *ret, err
}

// GetUserPerssion is a free data retrieval call binding the contract method 0x878f70c9.
//
// Solidity: function getUserPerssion(address key) constant returns(bool read, bool update, bool change)
func (_UserGroup *UserGroupSession) GetUserPerssion(key common.Address) (struct {
	Read   bool
	Update bool
	Change bool
}, error) {
	return _UserGroup.Contract.GetUserPerssion(&_UserGroup.CallOpts, key)
}

// GetUserPerssion is a free data retrieval call binding the contract method 0x878f70c9.
//
// Solidity: function getUserPerssion(address key) constant returns(bool read, bool update, bool change)
func (_UserGroup *UserGroupCallerSession) GetUserPerssion(key common.Address) (struct {
	Read   bool
	Update bool
	Change bool
}, error) {
	return _UserGroup.Contract.GetUserPerssion(&_UserGroup.CallOpts, key)
}

// AndOrupdateUser is a paid mutator transaction binding the contract method 0x4c767d0f.
//
// Solidity: function andOrupdateUser(address key, uint8 permisssion) returns(bool res)
func (_UserGroup *UserGroupTransactor) AndOrupdateUser(opts *bind.TransactOpts, key common.Address, permisssion uint8) (*types.Transaction, error) {
	return _UserGroup.contract.Transact(opts, "andOrupdateUser", key, permisssion)
}

// AndOrupdateUser is a paid mutator transaction binding the contract method 0x4c767d0f.
//
// Solidity: function andOrupdateUser(address key, uint8 permisssion) returns(bool res)
func (_UserGroup *UserGroupSession) AndOrupdateUser(key common.Address, permisssion uint8) (*types.Transaction, error) {
	return _UserGroup.Contract.AndOrupdateUser(&_UserGroup.TransactOpts, key, permisssion)
}

// AndOrupdateUser is a paid mutator transaction binding the contract method 0x4c767d0f.
//
// Solidity: function andOrupdateUser(address key, uint8 permisssion) returns(bool res)
func (_UserGroup *UserGroupTransactorSession) AndOrupdateUser(key common.Address, permisssion uint8) (*types.Transaction, error) {
	return _UserGroup.Contract.AndOrupdateUser(&_UserGroup.TransactOpts, key, permisssion)
}

// ChangeStatusOfUserGroup is a paid mutator transaction binding the contract method 0x9c4e9f7c.
//
// Solidity: function changeStatusOfUserGroup(bool status) returns()
func (_UserGroup *UserGroupTransactor) ChangeStatusOfUserGroup(opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	return _UserGroup.contract.Transact(opts, "changeStatusOfUserGroup", status)
}

// ChangeStatusOfUserGroup is a paid mutator transaction binding the contract method 0x9c4e9f7c.
//
// Solidity: function changeStatusOfUserGroup(bool status) returns()
func (_UserGroup *UserGroupSession) ChangeStatusOfUserGroup(status bool) (*types.Transaction, error) {
	return _UserGroup.Contract.ChangeStatusOfUserGroup(&_UserGroup.TransactOpts, status)
}

// ChangeStatusOfUserGroup is a paid mutator transaction binding the contract method 0x9c4e9f7c.
//
// Solidity: function changeStatusOfUserGroup(bool status) returns()
func (_UserGroup *UserGroupTransactorSession) ChangeStatusOfUserGroup(status bool) (*types.Transaction, error) {
	return _UserGroup.Contract.ChangeStatusOfUserGroup(&_UserGroup.TransactOpts, status)
}

// UserGroupChangeStatusOfUserGroupIterator is returned from FilterChangeStatusOfUserGroup and is used to iterate over the raw logs and unpacked data for ChangeStatusOfUserGroup events raised by the UserGroup contract.
type UserGroupChangeStatusOfUserGroupIterator struct {
	Event *UserGroupChangeStatusOfUserGroup // Event containing the contract specifics and raw log

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
func (it *UserGroupChangeStatusOfUserGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserGroupChangeStatusOfUserGroup)
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
		it.Event = new(UserGroupChangeStatusOfUserGroup)
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
func (it *UserGroupChangeStatusOfUserGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserGroupChangeStatusOfUserGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserGroupChangeStatusOfUserGroup represents a ChangeStatusOfUserGroup event raised by the UserGroup contract.
type UserGroupChangeStatusOfUserGroup struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChangeStatusOfUserGroup is a free log retrieval operation binding the contract event 0x65539928111859040d829c780a5ecc7a052a3c160109013679271b2e5c46d88c.
//
// Solidity: event _changeStatusOfUserGroup(bool status)
func (_UserGroup *UserGroupFilterer) FilterChangeStatusOfUserGroup(opts *bind.FilterOpts) (*UserGroupChangeStatusOfUserGroupIterator, error) {

	logs, sub, err := _UserGroup.contract.FilterLogs(opts, "_changeStatusOfUserGroup")
	if err != nil {
		return nil, err
	}
	return &UserGroupChangeStatusOfUserGroupIterator{contract: _UserGroup.contract, event: "_changeStatusOfUserGroup", logs: logs, sub: sub}, nil
}

// WatchChangeStatusOfUserGroup is a free log subscription operation binding the contract event 0x65539928111859040d829c780a5ecc7a052a3c160109013679271b2e5c46d88c.
//
// Solidity: event _changeStatusOfUserGroup(bool status)
func (_UserGroup *UserGroupFilterer) WatchChangeStatusOfUserGroup(opts *bind.WatchOpts, sink chan<- *UserGroupChangeStatusOfUserGroup) (event.Subscription, error) {

	logs, sub, err := _UserGroup.contract.WatchLogs(opts, "_changeStatusOfUserGroup")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserGroupChangeStatusOfUserGroup)
				if err := _UserGroup.contract.UnpackLog(event, "_changeStatusOfUserGroup", log); err != nil {
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

// UserGroupUpdateUserIterator is returned from FilterUpdateUser and is used to iterate over the raw logs and unpacked data for UpdateUser events raised by the UserGroup contract.
type UserGroupUpdateUserIterator struct {
	Event *UserGroupUpdateUser // Event containing the contract specifics and raw log

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
func (it *UserGroupUpdateUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserGroupUpdateUser)
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
		it.Event = new(UserGroupUpdateUser)
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
func (it *UserGroupUpdateUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserGroupUpdateUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserGroupUpdateUser represents a UpdateUser event raised by the UserGroup contract.
type UserGroupUpdateUser struct {
	Key    common.Address
	Read   bool
	Update bool
	Change bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateUser is a free log retrieval operation binding the contract event 0x9894bb533b126d657a7922559cc34561135edb194cf395b47e1c19e826056bae.
//
// Solidity: event _updateUser(address key, bool read, bool update, bool change)
func (_UserGroup *UserGroupFilterer) FilterUpdateUser(opts *bind.FilterOpts) (*UserGroupUpdateUserIterator, error) {

	logs, sub, err := _UserGroup.contract.FilterLogs(opts, "_updateUser")
	if err != nil {
		return nil, err
	}
	return &UserGroupUpdateUserIterator{contract: _UserGroup.contract, event: "_updateUser", logs: logs, sub: sub}, nil
}

// WatchUpdateUser is a free log subscription operation binding the contract event 0x9894bb533b126d657a7922559cc34561135edb194cf395b47e1c19e826056bae.
//
// Solidity: event _updateUser(address key, bool read, bool update, bool change)
func (_UserGroup *UserGroupFilterer) WatchUpdateUser(opts *bind.WatchOpts, sink chan<- *UserGroupUpdateUser) (event.Subscription, error) {

	logs, sub, err := _UserGroup.contract.WatchLogs(opts, "_updateUser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserGroupUpdateUser)
				if err := _UserGroup.contract.UnpackLog(event, "_updateUser", log); err != nil {
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

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package userGroups

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

// AddrToMappingABI is the input ABI used to generate the binding from.
const AddrToMappingABI = "[]"

// AddrToMappingBin is the compiled bytecode used for deploying new contracts.
const AddrToMappingBin = `0x6113c7610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100f45760003560e01c8063852579dd11610096578063e2bfd5ef11610070578063e2bfd5ef146104de578063e9f8aa87146105af578063f6023580146105d2578063f7bdc7511461022d576100f4565b8063852579dd14610344578063955a356c146103f8578063ce655e261461041b576100f4565b80633e391f6f116100d25780633e391f6f1461022d5780635bf71d021461025057806373197327146102fb57806379f0eb9e14610327576100f4565b8063038dac3a146100f957806314091c51146101c557806323d8ae96146101f4575b600080fd5b81801561010557600080fd5b506101b16004803603604081101561011c57600080fd5b81359190810190604081016020820135600160201b81111561013d57600080fd5b82018360208201111561014f57600080fd5b803590602001918460018302840111600160201b8311171561017057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610618945050505050565b604080519115158252519081900360200190f35b6101e2600480360360208110156101db57600080fd5b503561075b565b60408051918252519081900360200190f35b81801561020057600080fd5b506101b16004803603604081101561021757600080fd5b50803590602001356001600160a01b0316610769565b6101b16004803603604081101561024357600080fd5b5080359060200135610964565b6101b16004803603604081101561026657600080fd5b81359190810190604081016020820135600160201b81111561028757600080fd5b82018360208201111561029957600080fd5b803590602001918460018302840111600160201b831117156102ba57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061096f945050505050565b6101b16004803603604081101561031157600080fd5b50803590602001356001600160a01b03166109df565b6101e26004803603602081101561033d57600080fd5b50356109fd565b6103676004803603604081101561035a57600080fd5b5080359060200135610a0b565b6040518080602001836001600160a01b03166001600160a01b03168152602001828103825284818151815260200191508051906020019080838360005b838110156103bc5781810151838201526020016103a4565b50505050905090810190601f1680156103e95780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6101e26004803603604081101561040e57600080fd5b5080359060200135610b32565b81801561042757600080fd5b506101b16004803603606081101561043e57600080fd5b81359190810190604081016020820135600160201b81111561045f57600080fd5b82018360208201111561047157600080fd5b803590602001918460018302840111600160201b8311171561049257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b03169150610b829050565b8180156104ea57600080fd5b506101b16004803603608081101561050157600080fd5b8135916001600160a01b0360208201351691810190606081016040820135600160201b81111561053057600080fd5b82018360208201111561054257600080fd5b803590602001918460018302840111600160201b8311171561056357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b03169150610d499050565b6101e2600480360360408110156105c557600080fd5b508035906020013561115a565b6105f5600480360360408110156105e857600080fd5b50803590602001356111a5565b604080516001600160a01b03909316835260208301919091528051918290030190f35b60008083600001836040518082805190602001908083835b6020831061064f5780518252601f199092019160209182019101610630565b51815160001960209485036101000a019081169019919091161790529201948552506040519384900301909220549250505080610690576000915050610755565b83600001836040518082805190602001908083835b602083106106c45780518252601f1990920191602091820191016106a5565b51815160001960209485036101000a81019182169119929092161790915293909101958652604051958690030190942060008155600190810180546001600160a01b0319169055888101805491955093509085019150811061072257fe5b600091825260209091206002918202016001908101805460ff191693151593909317909255850180546000190190559150505b92915050565b600061075582600019610b32565b6001600160a01b03811660009081526020839052604081205480610791576000915050610755565b6001600160a01b03831660009081526020859052604081206107b5906001016109fd565b90505b6001600160a01b03841660009081526020869052604090206107dd9060010182610964565b156108dd576001600160a01b03841660009081526020869052604090206002810180546108b09260010191908490811061081357fe5b60009182526020918290206002918202018054604080516001831615610100026000190190921693909304601f8101859004850282018501909352828152929091908301828280156108a65780601f1061087b576101008083540402835291602001916108a6565b820191906000526020600020905b81548152906001019060200180831161088957829003601f168201915b5050505050610618565b506001600160a01b03841660009081526020869052604090206108d6906001018261115a565b90506107b8565b506001600160a01b038316600090815260208590526040812081815590600182018161090c60028501826111e7565b600282016000905550505050600184600101600183038154811061092c57fe5b60009182526020909120018054911515600160a01b0260ff60a01b199092169190911790555050506002018054600019019055600190565b600191909101541190565b60008083600001836040518082805190602001908083835b602083106109a65780518252601f199092019160209182019101610987565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220549290921195945050505050565b6001600160a01b031660009081526020919091526040902054151590565b60006107558260001961115a565b60606000836001018381548110610a1e57fe5b60009182526020918290206002918202018054604080516001831615610100026000190190921693909304601f810185900485028201850190935282815292909190830182828015610ab15780601f10610a8657610100808354040283529160200191610ab1565b820191906000526020600020905b815481529060010190602001808311610a9457829003601f168201915b5050505050915083600001826040518082805190602001908083835b60208310610aec5780518252601f199092019160209182019101610acd565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206001015493966001600160a01b039094169550929350505050565b60010160005b600183015482108015610b6c5750826001018281548110610b5557fe5b600091825260209091200154600160a01b900460ff165b15610b7c57600190910190610b38565b50919050565b60008084600001846040518082805190602001908083835b60208310610bb95780518252601f199092019160209182019101610b9a565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101842054885190955087948a9450899350918291908401908083835b60208310610c1c5780518252601f199092019160209182019101610bfd565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922060010180546001600160a01b0319166001600160a01b03949094169390931790925550508015610c7b576001915050610d42565b6001808601805491610c8f9190830161120b565b90508060010185600001856040518082805190602001908083835b60208310610cc95780518252601f199092019160209182019101610caa565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092909255505060018501805485919083908110610d0d57fe5b90600052602060002090600202016000019080519060200190610d3192919061123c565b505050600283018054600101905560005b9392505050565b6001600160a01b038316600090815260208581526040808320805491518651929385936001909301928892918291908401908083835b60208310610d9e5780518252601f199092019160209182019101610d7f565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382018520546001600160a01b038c1660009081528d8452919091208a519196508995600190910194508a9350918291908401908083835b60208310610e1c5780518252601f199092019160209182019101610dfd565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922060010180546001600160a01b0319166001600160a01b03949094169390931790925550508115610fab578015610e8257600192505050611152565b6001600160a01b0386166000908152602088905260409020600201805490610ead906001830161120b565b6001600160a01b038716600090815260208981526040918290209151885193945060018086019493019289928291908401908083835b60208310610f025780518252601f199092019160209182019101610ee3565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094209490945550506001600160a01b0388166000908152918990529020600201805486919083908110610f5c57fe5b90600052602060002090600202016000019080519060200190610f8092919061123c565b505050506001600160a01b038316600090815260208590526040812060030180546001019055611152565b6001600160a01b0386166000908152602088905260409020600201805490610fd6906001830161120b565b6001600160a01b038716600090815260208981526040918290209151885193945060018086019493019289928291908401908083835b6020831061102b5780518252601f19909201916020918201910161100c565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094209490945550506001600160a01b038816600090815291899052902060020180548691908390811061108557fe5b906000526020600020906002020160000190805190602001906110a992919061123c565b506001600160a01b0386166000908152602088905260409020600301805460019081019091558088018054916110e1919083016112ba565b6001600160a01b038716600090815260208990526040902060018083019091558801805491935087918490811061111457fe5b6000918252602082200180546001600160a01b0319166001600160a01b03939093169290921790915560028801805460010190559250611152915050565b949350505050565b60010160005b600183015482108015611195575082600101828154811061117d57fe5b600091825260209091206001600290920201015460ff165b15610b7c57600190910190611160565b6000808360010183815481106111b757fe5b60009182526020808320909101546001600160a01b03168083529590526040902093946001949094019392505050565b508054600082556002029060005260206000209081019061120891906112de565b50565b8154818355818111156112375760020281600202836000526020600020918201910161123791906112de565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061127d57805160ff19168380011785556112aa565b828001600101855582156112aa579182015b828111156112aa57825182559160200191906001019061128f565b506112b6929150611310565b5090565b8154818355818111156112375760008381526020902061123791810190830161132a565b61130d91905b808211156112b65760006112f8828261134e565b5060018101805460ff191690556002016112e4565b90565b61130d91905b808211156112b65760008155600101611316565b61130d91905b808211156112b65780546001600160a81b0319168155600101611330565b50805460018160011615610100020316600290046000825580601f106113745750611208565b601f016020900490600052602060002090810190611208919061131056fea265627a7a72305820b7fab7f7dffaeddea94665a4530ecd5372da347a8192a2b96026d849ccad362164736f6c634300050a0032`

// DeployAddrToMapping deploys a new Ethereum contract, binding an instance of AddrToMapping to it.
func DeployAddrToMapping(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddrToMapping, error) {
	parsed, err := abi.JSON(strings.NewReader(AddrToMappingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddrToMappingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddrToMapping{AddrToMappingCaller: AddrToMappingCaller{contract: contract}, AddrToMappingTransactor: AddrToMappingTransactor{contract: contract}, AddrToMappingFilterer: AddrToMappingFilterer{contract: contract}}, nil
}

// AddrToMapping is an auto generated Go binding around an Ethereum contract.
type AddrToMapping struct {
	AddrToMappingCaller     // Read-only binding to the contract
	AddrToMappingTransactor // Write-only binding to the contract
	AddrToMappingFilterer   // Log filterer for contract events
}

// AddrToMappingCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddrToMappingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrToMappingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddrToMappingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrToMappingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddrToMappingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddrToMappingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddrToMappingSession struct {
	Contract     *AddrToMapping    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddrToMappingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddrToMappingCallerSession struct {
	Contract *AddrToMappingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AddrToMappingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddrToMappingTransactorSession struct {
	Contract     *AddrToMappingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AddrToMappingRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddrToMappingRaw struct {
	Contract *AddrToMapping // Generic contract binding to access the raw methods on
}

// AddrToMappingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddrToMappingCallerRaw struct {
	Contract *AddrToMappingCaller // Generic read-only contract binding to access the raw methods on
}

// AddrToMappingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddrToMappingTransactorRaw struct {
	Contract *AddrToMappingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddrToMapping creates a new instance of AddrToMapping, bound to a specific deployed contract.
func NewAddrToMapping(address common.Address, backend bind.ContractBackend) (*AddrToMapping, error) {
	contract, err := bindAddrToMapping(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddrToMapping{AddrToMappingCaller: AddrToMappingCaller{contract: contract}, AddrToMappingTransactor: AddrToMappingTransactor{contract: contract}, AddrToMappingFilterer: AddrToMappingFilterer{contract: contract}}, nil
}

// NewAddrToMappingCaller creates a new read-only instance of AddrToMapping, bound to a specific deployed contract.
func NewAddrToMappingCaller(address common.Address, caller bind.ContractCaller) (*AddrToMappingCaller, error) {
	contract, err := bindAddrToMapping(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddrToMappingCaller{contract: contract}, nil
}

// NewAddrToMappingTransactor creates a new write-only instance of AddrToMapping, bound to a specific deployed contract.
func NewAddrToMappingTransactor(address common.Address, transactor bind.ContractTransactor) (*AddrToMappingTransactor, error) {
	contract, err := bindAddrToMapping(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddrToMappingTransactor{contract: contract}, nil
}

// NewAddrToMappingFilterer creates a new log filterer instance of AddrToMapping, bound to a specific deployed contract.
func NewAddrToMappingFilterer(address common.Address, filterer bind.ContractFilterer) (*AddrToMappingFilterer, error) {
	contract, err := bindAddrToMapping(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddrToMappingFilterer{contract: contract}, nil
}

// bindAddrToMapping binds a generic wrapper to an already deployed contract.
func bindAddrToMapping(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddrToMappingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddrToMapping *AddrToMappingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddrToMapping.Contract.AddrToMappingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddrToMapping *AddrToMappingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddrToMapping.Contract.AddrToMappingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddrToMapping *AddrToMappingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddrToMapping.Contract.AddrToMappingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddrToMapping *AddrToMappingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddrToMapping.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddrToMapping *AddrToMappingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddrToMapping.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddrToMapping *AddrToMappingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddrToMapping.Contract.contract.Transact(opts, method, params...)
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

// UserGroupsABI is the input ABI used to generate the binding from.
const UserGroupsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"updateUserGroupWithNewGroup\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"userGroupCreator\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"value\",\"type\":\"address\"}],\"name\":\"updateUserGroupByAddress\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"creater\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"deleteGroupByName\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerToGroups\",\"outputs\":[{\"name\":\"size\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getGroupCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getGroupByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deleteOwnerRec\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"address\"}],\"name\":\"groups\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"add\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"address\"}],\"name\":\"_updateUserGroup\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"_deleteGroupByName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"_deleteOwnerRec\",\"type\":\"event\"}]"

// UserGroupsBin is the compiled bytecode used for deploying new contracts.
const UserGroupsBin = `0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055611d23806100326000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806375e2f29c1161006657806375e2f29c1461022b5780638e933ef7146102cf57806396a2f896146102e9578063d4eeccce1461030f578063f02b1b4e146103cc57610093565b806316a94bbf146100985780632bc07c5a146101505780633847f37b1461017457806345653a6d14610223575b600080fd5b61013c600480360360208110156100ae57600080fd5b810190602081018135600160201b8111156100c857600080fd5b8201836020820111156100da57600080fd5b803590602001918460018302840111600160201b831117156100fb57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506103d4945050505050565b604080519115158252519081900360200190f35b6101586105ae565b604080516001600160a01b039092168252519081900360200190f35b61013c6004803603604081101561018a57600080fd5b810190602081018135600160201b8111156101a457600080fd5b8201836020820111156101b657600080fd5b803590602001918460018302840111600160201b831117156101d757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b031691506105f19050565b6101586107ba565b61013c6004803603602081101561024157600080fd5b810190602081018135600160201b81111561025b57600080fd5b82018360208201111561026d57600080fd5b803590602001918460018302840111600160201b8311171561028e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506107c9945050505050565b6102d7610bdd565b60408051918252519081900360200190f35b6102d7600480360360208110156102ff57600080fd5b50356001600160a01b0316610be3565b61033b6004803603604081101561032557600080fd5b506001600160a01b038135169060200135610c9f565b6040518080602001836001600160a01b03166001600160a01b03168152602001828103825284818151815260200191508051906020019080838360005b83811015610390578181015183820152602001610378565b50505050905090810190601f1680156103bd5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b61013c610e51565b6000806103df6105ae565b60405163e2bfd5ef60e01b815260016004820181815233602484018190526001600160a01b038516606485015260806044850190815288516084860152885195965073__$11cae51210796eb8f20e6110d091dfb08c$__9563e2bfd5ef9592938a938993919260a490910190602086019080838360005b8381101561046e578181015183820152602001610456565b50505050905090810190601f16801561049b5780820380516001836020036101000a031916815260200191505b509550505050505060206040518083038186803b1580156104bb57600080fd5b505af41580156104cf573d6000803e3d6000fd5b505050506040513d60208110156104e557600080fd5b505060408051338082526001600160a01b038416928201929092526060602082810182815287519284019290925286517f9b4a33bb61ca8ea8f298d84955207994837a4921e3ed0741753f5ca9a3456d9d9493889387939192909160808401919086019080838360005b8381101561056757818101518382015260200161054f565b50505050905090810190601f1680156105945780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a160019150505b919050565b6000336040516105bd9061139b565b6001600160a01b03909116815260405190819003602001906000f0801580156105ea573d6000803e3d6000fd5b5090505b90565b60405163e2bfd5ef60e01b815260016004820181815233602484018190526001600160a01b038516606485015260806044850190815286516084860152865160009573__$11cae51210796eb8f20e6110d091dfb08c$__9563e2bfd5ef959094938a938a9360a4019060208601908083838e5b8381101561067c578181015183820152602001610664565b50505050905090810190601f1680156106a95780820380516001836020036101000a031916815260200191505b509550505050505060206040518083038186803b1580156106c957600080fd5b505af41580156106dd573d6000803e3d6000fd5b505050506040513d60208110156106f357600080fd5b505060408051338082526001600160a01b038516928201929092526060602082810182815287519284019290925286517f9b4a33bb61ca8ea8f298d84955207994837a4921e3ed0741753f5ca9a3456d9d9493889388939192909160808401919086019080838360005b8381101561077557818101518382015260200161075d565b50505050905090810190601f1680156107a25780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a150600192915050565b6000546001600160a01b031681565b60408051637319732760e01b815260016004820152336024820152905160009173__$11cae51210796eb8f20e6110d091dfb08c$__91637319732791604480820192602092909190829003018186803b15801561082557600080fd5b505af4158015610839573d6000803e3d6000fd5b505050506040513d602081101561084f57600080fd5b505161085d575060006105a9565b3360009081526001602081815260408084208151632dfb8e8160e11b81529301600484018181526024850192835287516044860152875173__$11cae51210796eb8f20e6110d091dfb08c$__96635bf71d029693958a95939460640192908601918190849084905b838110156108dd5781810151838201526020016108c5565b50505050905090810190601f16801561090a5780820380516001836020036101000a031916815260200191505b50935050505060206040518083038186803b15801561092857600080fd5b505af415801561093c573d6000803e3d6000fd5b505050506040513d602081101561095257600080fd5b5051610960575060006105a9565b3360009081526001602081815260408084209051865191909301928692909182918401908083835b602083106109a75780518252601f199092019160209182019101610988565b51815160209384036101000a600019018019909216911617905292019485525060408051948590039091018420600190810154632713a7df60e21b8652600486019190915290516001600160a01b0390911694508493849350639c4e9f7c925060248082019260009290919082900301818387803b158015610a2857600080fd5b505af1158015610a3c573d6000803e3d6000fd5b505033600090815260016020818152604080842081516301c6d61d60e11b8152930160048401818152602485019283528b5160448601528b5173__$11cae51210796eb8f20e6110d091dfb08c$__985063038dac3a975091958c95919460649092019291860191908190849084905b83811015610ac3578181015183820152602001610aab565b50505050905090810190601f168015610af05780820380516001836020036101000a031916815260200191505b50935050505060206040518083038186803b158015610b0e57600080fd5b505af4158015610b22573d6000803e3d6000fd5b505050506040513d6020811015610b3857600080fd5b505060408051602080825286518282015286517f0d21241cbf9584f265935b748e8240b0397928e543a3d915fa70471f722c8ce3938893928392918301919085019080838360005b83811015610b98578181015183820152602001610b80565b50505050905090810190601f168015610bc55780820380516001836020036101000a031916815260200191505b509250505060405180910390a16001925050506105a9565b60035481565b60408051637319732760e01b8152600160048201526001600160a01b0383166024820152905160009173__$11cae51210796eb8f20e6110d091dfb08c$__91637319732791604480820192602092909190829003018186803b158015610c4857600080fd5b505af4158015610c5c573d6000803e3d6000fd5b505050506040513d6020811015610c7257600080fd5b5051610c80575060006105a9565b506001600160a01b031660009081526001602052604090206003015490565b60408051637319732760e01b8152600160048201526001600160a01b0384166024820152905160609160009173__$11cae51210796eb8f20e6110d091dfb08c$__916373197327916044808301926020929190829003018186803b158015610d0657600080fd5b505af4158015610d1a573d6000803e3d6000fd5b505050506040513d6020811015610d3057600080fd5b5051610d4d57505060408051602081019091526000808252610e4a565b6001600160a01b0384166000908152600160208190526040808320815163852579dd60e01b815292016004830152602482018690525173__$11cae51210796eb8f20e6110d091dfb08c$__9263852579dd9260448082019391829003018186803b158015610dba57600080fd5b505af4158015610dce573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040908152811015610df757600080fd5b810190808051600160201b811115610e0e57600080fd5b82016020810184811115610e2157600080fd5b8151600160201b811182820187101715610e3a57600080fd5b5050602090910151909450925050505b9250929050565b60408051637319732760e01b815260016004820152336024820152905160009173__$11cae51210796eb8f20e6110d091dfb08c$__91637319732791604480820192602092909190829003018186803b158015610ead57600080fd5b505af4158015610ec1573d6000803e3d6000fd5b505050506040513d6020811015610ed757600080fd5b5051610ee5575060006105ee565b3360009081526001602081815260408084208151633cf875cf60e11b8152930160048401819052905190939273__$11cae51210796eb8f20e6110d091dfb08c$__926379f0eb9e92602480840193829003018186803b158015610f4757600080fd5b505af4158015610f5b573d6000803e3d6000fd5b505050506040513d6020811015610f7157600080fd5b505190505b73__$11cae51210796eb8f20e6110d091dfb08c$__633e391f6f83836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b158015610fcd57600080fd5b505af4158015610fe1573d6000803e3d6000fd5b505050506040513d6020811015610ff757600080fd5b505115611312576060600073__$11cae51210796eb8f20e6110d091dfb08c$__63852579dd85856040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561105957600080fd5b505af415801561106d573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604090815281101561109657600080fd5b810190808051600160201b8111156110ad57600080fd5b820160208101848111156110c057600080fd5b8151600160201b8111828201871017156110d957600080fd5b50506020918201513360009081526001808552604080832081516301c6d61d60e11b815292016004830181815260248401928352865160448501528651969b5094995073__$11cae51210796eb8f20e6110d091dfb08c$__985063038dac3a9750958a95919360649093019290860191908190849084905b83811015611169578181015183820152602001611151565b50505050905090810190601f1680156111965780820380516001836020036101000a031916815260200191505b50935050505060206040518083038186803b1580156111b457600080fd5b505af41580156111c8573d6000803e3d6000fd5b505050506040513d60208110156111de57600080fd5b50506040805133808252602082810184815286519484019490945285517f9064673ca868275dc97f9fe659c8421d0fa89b07a8578773dd156336d186c3d3949293879390929091606084019185019080838360005b8381101561124b578181015183820152602001611233565b50505050905090810190601f1680156112785780820380516001836020036101000a031916815260200191505b50935050505060405180910390a1505073__$11cae51210796eb8f20e6110d091dfb08c$__63e9f8aa8783836040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b1580156112df57600080fd5b505af41580156112f3573d6000803e3d6000fd5b505050506040513d602081101561130957600080fd5b50519050610f76565b50604080516311ec574b60e11b815260016004820152336024820152905173__$11cae51210796eb8f20e6110d091dfb08c$__916323d8ae96916044808301926020929190829003018186803b15801561136b57600080fd5b505af415801561137f573d6000803e3d6000fd5b505050506040513d602081101561139557600080fd5b50505090565b610946806113a98339019056fe60806040526000805460ff60a01b1916905534801561001d57600080fd5b506040516109463803806109468339818101604052602081101561004057600080fd5b5051600080546001600160a01b039092166001600160a01b03199092169190911790556108d4806100726000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c806302d05d3f1461005c5780634c767d0f146100805780636b35f7c1146100c3578063878f70c9146100cb5780639c4e9f7c14610113575b600080fd5b610064610134565b604080516001600160a01b039092168252519081900360200190f35b6100af6004803603604081101561009657600080fd5b5080356001600160a01b0316906020013560ff16610143565b604080519115158252519081900360200190f35b6100af6104a1565b6100f1600480360360208110156100e157600080fd5b50356001600160a01b03166104b1565b6040805193151584529115156020840152151582820152519081900360600190f35b6101326004803603602081101561012957600080fd5b5035151561074b565b005b6000546001600160a01b031681565b60408051636d1fca1360e11b815260076004820152336024820152905160009173__$1a5b05563ffbb1c6a29a7456f190d0d702$__9163da3f942691604480820192602092909190829003018186803b15801561019f57600080fd5b505af41580156101b3573d6000803e3d6000fd5b505050506040513d60208110156101c957600080fd5b5051806101e057506000546001600160a01b031633145b61021b5760405162461bcd60e51b815260040180806020018281038252602d815260200180610873602d913960400191505060405180910390fd5b60078260ff16118061022b575060005b156102385750600061049b565b600060018084161461024b57600061024e565b60015b90506000600180851614610263576000610266565b60015b9050600060018086161461027b57600061027e565b60015b604080516398b3fbd760e01b8152600160048201526001600160a01b03891660248201528515156044820152905191925073__$1a5b05563ffbb1c6a29a7456f190d0d702$__916398b3fbd791606480820192602092909190829003018186803b1580156102eb57600080fd5b505af41580156102ff573d6000803e3d6000fd5b505050506040513d602081101561031557600080fd5b5050604080516398b3fbd760e01b81526004818101526001600160a01b03881660248201528315156044820152905173__$1a5b05563ffbb1c6a29a7456f190d0d702$__916398b3fbd7916064808301926020929190829003018186803b15801561037f57600080fd5b505af4158015610393573d6000803e3d6000fd5b505050506040513d60208110156103a957600080fd5b5050604080516398b3fbd760e01b8152600760048201526001600160a01b03881660248201528215156044820152905173__$1a5b05563ffbb1c6a29a7456f190d0d702$__916398b3fbd7916064808301926020929190829003018186803b15801561041457600080fd5b505af4158015610428573d6000803e3d6000fd5b505050506040513d602081101561043e57600080fd5b5050604080516001600160a01b0388168152841515602082015283151581830152821515606082015290517f9894bb533b126d657a7922559cc34561135edb194cf395b47e1c19e826056bae9181900360800190a1600193505050505b92915050565b600054600160a01b900460ff1681565b60008054819081906001600160a01b03858116911614156104da57506001915081905080610744565b600054600160a01b900460ff161515600114156105005750600191508190506000610744565b60408051636d1fca1360e11b8152600160048201526001600160a01b0386166024820152905173__$1a5b05563ffbb1c6a29a7456f190d0d702$__9163da3f9426916044808301926020929190829003018186803b15801561056157600080fd5b505af4158015610575573d6000803e3d6000fd5b505050506040513d602081101561058b57600080fd5b50511515600114156105be576001600160a01b0384166000908152600160208190526040909120015460ff1692506105c3565b600092505b60408051636d1fca1360e11b81526004818101526001600160a01b0386166024820152905173__$1a5b05563ffbb1c6a29a7456f190d0d702$__9163da3f9426916044808301926020929190829003018186803b15801561062357600080fd5b505af4158015610637573d6000803e3d6000fd5b505050506040513d602081101561064d57600080fd5b505115156001141561067f576001600160a01b03841660009081526004602052604090206001015460ff169150610684565b600091505b60408051636d1fca1360e11b8152600760048201526001600160a01b0386166024820152905173__$1a5b05563ffbb1c6a29a7456f190d0d702$__9163da3f9426916044808301926020929190829003018186803b1580156106e557600080fd5b505af41580156106f9573d6000803e3d6000fd5b505050506040513d602081101561070f57600080fd5b505115156001141561074057506001600160a01b03831660009081526007602052604090206001015460ff16610744565b5060005b9193909250565b60408051636d1fca1360e11b815260076004820152336024820152905173__$1a5b05563ffbb1c6a29a7456f190d0d702$__9163da3f9426916044808301926020929190829003018186803b1580156107a357600080fd5b505af41580156107b7573d6000803e3d6000fd5b505050506040513d60208110156107cd57600080fd5b5051806107e457506000546001600160a01b031633145b61081f5760405162461bcd60e51b815260040180806020018281038252602d815260200180610873602d913960400191505060405180910390fd5b60008054821515600160a01b810260ff60a01b199092169190911790915560408051918252517f65539928111859040d829c780a5ecc7a052a3c160109013679271b2e5c46d88c9181900360200190a15056fe4f6e6c79207065726d69737373696f6e20757365722063616e2063616c6c20746869732066756e6374696f6e2ea265627a7a723058203d7ab454d7dc5c7892eeb0fa4062e76ee7f873099cbcd12e4ef08a36f7599cc864736f6c634300050a0032a265627a7a723058205a103931fb651a85ab24f7ef202d01292d6c5066a48fc69451f4739cfcaa74f064736f6c634300050a0032`

// DeployUserGroups deploys a new Ethereum contract, binding an instance of UserGroups to it.
func DeployUserGroups(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserGroups, error) {
	parsed, err := abi.JSON(strings.NewReader(UserGroupsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserGroupsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserGroups{UserGroupsCaller: UserGroupsCaller{contract: contract}, UserGroupsTransactor: UserGroupsTransactor{contract: contract}, UserGroupsFilterer: UserGroupsFilterer{contract: contract}}, nil
}

// UserGroups is an auto generated Go binding around an Ethereum contract.
type UserGroups struct {
	UserGroupsCaller     // Read-only binding to the contract
	UserGroupsTransactor // Write-only binding to the contract
	UserGroupsFilterer   // Log filterer for contract events
}

// UserGroupsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserGroupsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserGroupsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserGroupsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserGroupsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserGroupsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserGroupsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserGroupsSession struct {
	Contract     *UserGroups       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserGroupsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserGroupsCallerSession struct {
	Contract *UserGroupsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// UserGroupsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserGroupsTransactorSession struct {
	Contract     *UserGroupsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// UserGroupsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserGroupsRaw struct {
	Contract *UserGroups // Generic contract binding to access the raw methods on
}

// UserGroupsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserGroupsCallerRaw struct {
	Contract *UserGroupsCaller // Generic read-only contract binding to access the raw methods on
}

// UserGroupsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserGroupsTransactorRaw struct {
	Contract *UserGroupsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserGroups creates a new instance of UserGroups, bound to a specific deployed contract.
func NewUserGroups(address common.Address, backend bind.ContractBackend) (*UserGroups, error) {
	contract, err := bindUserGroups(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserGroups{UserGroupsCaller: UserGroupsCaller{contract: contract}, UserGroupsTransactor: UserGroupsTransactor{contract: contract}, UserGroupsFilterer: UserGroupsFilterer{contract: contract}}, nil
}

// NewUserGroupsCaller creates a new read-only instance of UserGroups, bound to a specific deployed contract.
func NewUserGroupsCaller(address common.Address, caller bind.ContractCaller) (*UserGroupsCaller, error) {
	contract, err := bindUserGroups(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserGroupsCaller{contract: contract}, nil
}

// NewUserGroupsTransactor creates a new write-only instance of UserGroups, bound to a specific deployed contract.
func NewUserGroupsTransactor(address common.Address, transactor bind.ContractTransactor) (*UserGroupsTransactor, error) {
	contract, err := bindUserGroups(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserGroupsTransactor{contract: contract}, nil
}

// NewUserGroupsFilterer creates a new log filterer instance of UserGroups, bound to a specific deployed contract.
func NewUserGroupsFilterer(address common.Address, filterer bind.ContractFilterer) (*UserGroupsFilterer, error) {
	contract, err := bindUserGroups(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserGroupsFilterer{contract: contract}, nil
}

// bindUserGroups binds a generic wrapper to an already deployed contract.
func bindUserGroups(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserGroupsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserGroups *UserGroupsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserGroups.Contract.UserGroupsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserGroups *UserGroupsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserGroups.Contract.UserGroupsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserGroups *UserGroupsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserGroups.Contract.UserGroupsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserGroups *UserGroupsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserGroups.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserGroups *UserGroupsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserGroups.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserGroups *UserGroupsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserGroups.Contract.contract.Transact(opts, method, params...)
}

// Creater is a free data retrieval call binding the contract method 0x45653a6d.
//
// Solidity: function creater() constant returns(address)
func (_UserGroups *UserGroupsCaller) Creater(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserGroups.contract.Call(opts, out, "creater")
	return *ret0, err
}

// Creater is a free data retrieval call binding the contract method 0x45653a6d.
//
// Solidity: function creater() constant returns(address)
func (_UserGroups *UserGroupsSession) Creater() (common.Address, error) {
	return _UserGroups.Contract.Creater(&_UserGroups.CallOpts)
}

// Creater is a free data retrieval call binding the contract method 0x45653a6d.
//
// Solidity: function creater() constant returns(address)
func (_UserGroups *UserGroupsCallerSession) Creater() (common.Address, error) {
	return _UserGroups.Contract.Creater(&_UserGroups.CallOpts)
}

// GetGroupByIndex is a free data retrieval call binding the contract method 0xd4eeccce.
//
// Solidity: function getGroupByIndex(address owner, uint256 index) constant returns(string, address)
func (_UserGroups *UserGroupsCaller) GetGroupByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (string, common.Address, error) {
	var (
		ret0 = new(string)
		ret1 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _UserGroups.contract.Call(opts, out, "getGroupByIndex", owner, index)
	return *ret0, *ret1, err
}

// GetGroupByIndex is a free data retrieval call binding the contract method 0xd4eeccce.
//
// Solidity: function getGroupByIndex(address owner, uint256 index) constant returns(string, address)
func (_UserGroups *UserGroupsSession) GetGroupByIndex(owner common.Address, index *big.Int) (string, common.Address, error) {
	return _UserGroups.Contract.GetGroupByIndex(&_UserGroups.CallOpts, owner, index)
}

// GetGroupByIndex is a free data retrieval call binding the contract method 0xd4eeccce.
//
// Solidity: function getGroupByIndex(address owner, uint256 index) constant returns(string, address)
func (_UserGroups *UserGroupsCallerSession) GetGroupByIndex(owner common.Address, index *big.Int) (string, common.Address, error) {
	return _UserGroups.Contract.GetGroupByIndex(&_UserGroups.CallOpts, owner, index)
}

// GetGroupCount is a free data retrieval call binding the contract method 0x96a2f896.
//
// Solidity: function getGroupCount(address owner) constant returns(uint256)
func (_UserGroups *UserGroupsCaller) GetGroupCount(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UserGroups.contract.Call(opts, out, "getGroupCount", owner)
	return *ret0, err
}

// GetGroupCount is a free data retrieval call binding the contract method 0x96a2f896.
//
// Solidity: function getGroupCount(address owner) constant returns(uint256)
func (_UserGroups *UserGroupsSession) GetGroupCount(owner common.Address) (*big.Int, error) {
	return _UserGroups.Contract.GetGroupCount(&_UserGroups.CallOpts, owner)
}

// GetGroupCount is a free data retrieval call binding the contract method 0x96a2f896.
//
// Solidity: function getGroupCount(address owner) constant returns(uint256)
func (_UserGroups *UserGroupsCallerSession) GetGroupCount(owner common.Address) (*big.Int, error) {
	return _UserGroups.Contract.GetGroupCount(&_UserGroups.CallOpts, owner)
}

// OwnerToGroups is a free data retrieval call binding the contract method 0x8e933ef7.
//
// Solidity: function ownerToGroups() constant returns(uint256 size)
func (_UserGroups *UserGroupsCaller) OwnerToGroups(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _UserGroups.contract.Call(opts, out, "ownerToGroups")
	return *ret0, err
}

// OwnerToGroups is a free data retrieval call binding the contract method 0x8e933ef7.
//
// Solidity: function ownerToGroups() constant returns(uint256 size)
func (_UserGroups *UserGroupsSession) OwnerToGroups() (*big.Int, error) {
	return _UserGroups.Contract.OwnerToGroups(&_UserGroups.CallOpts)
}

// OwnerToGroups is a free data retrieval call binding the contract method 0x8e933ef7.
//
// Solidity: function ownerToGroups() constant returns(uint256 size)
func (_UserGroups *UserGroupsCallerSession) OwnerToGroups() (*big.Int, error) {
	return _UserGroups.Contract.OwnerToGroups(&_UserGroups.CallOpts)
}

// DeleteGroupByName is a paid mutator transaction binding the contract method 0x75e2f29c.
//
// Solidity: function deleteGroupByName(string name) returns(bool res)
func (_UserGroups *UserGroupsTransactor) DeleteGroupByName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _UserGroups.contract.Transact(opts, "deleteGroupByName", name)
}

// DeleteGroupByName is a paid mutator transaction binding the contract method 0x75e2f29c.
//
// Solidity: function deleteGroupByName(string name) returns(bool res)
func (_UserGroups *UserGroupsSession) DeleteGroupByName(name string) (*types.Transaction, error) {
	return _UserGroups.Contract.DeleteGroupByName(&_UserGroups.TransactOpts, name)
}

// DeleteGroupByName is a paid mutator transaction binding the contract method 0x75e2f29c.
//
// Solidity: function deleteGroupByName(string name) returns(bool res)
func (_UserGroups *UserGroupsTransactorSession) DeleteGroupByName(name string) (*types.Transaction, error) {
	return _UserGroups.Contract.DeleteGroupByName(&_UserGroups.TransactOpts, name)
}

// DeleteOwnerRec is a paid mutator transaction binding the contract method 0xf02b1b4e.
//
// Solidity: function deleteOwnerRec() returns(bool res)
func (_UserGroups *UserGroupsTransactor) DeleteOwnerRec(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserGroups.contract.Transact(opts, "deleteOwnerRec")
}

// DeleteOwnerRec is a paid mutator transaction binding the contract method 0xf02b1b4e.
//
// Solidity: function deleteOwnerRec() returns(bool res)
func (_UserGroups *UserGroupsSession) DeleteOwnerRec() (*types.Transaction, error) {
	return _UserGroups.Contract.DeleteOwnerRec(&_UserGroups.TransactOpts)
}

// DeleteOwnerRec is a paid mutator transaction binding the contract method 0xf02b1b4e.
//
// Solidity: function deleteOwnerRec() returns(bool res)
func (_UserGroups *UserGroupsTransactorSession) DeleteOwnerRec() (*types.Transaction, error) {
	return _UserGroups.Contract.DeleteOwnerRec(&_UserGroups.TransactOpts)
}

// UpdateUserGroupByAddress is a paid mutator transaction binding the contract method 0x3847f37b.
//
// Solidity: function updateUserGroupByAddress(string name, address value) returns(bool res)
func (_UserGroups *UserGroupsTransactor) UpdateUserGroupByAddress(opts *bind.TransactOpts, name string, value common.Address) (*types.Transaction, error) {
	return _UserGroups.contract.Transact(opts, "updateUserGroupByAddress", name, value)
}

// UpdateUserGroupByAddress is a paid mutator transaction binding the contract method 0x3847f37b.
//
// Solidity: function updateUserGroupByAddress(string name, address value) returns(bool res)
func (_UserGroups *UserGroupsSession) UpdateUserGroupByAddress(name string, value common.Address) (*types.Transaction, error) {
	return _UserGroups.Contract.UpdateUserGroupByAddress(&_UserGroups.TransactOpts, name, value)
}

// UpdateUserGroupByAddress is a paid mutator transaction binding the contract method 0x3847f37b.
//
// Solidity: function updateUserGroupByAddress(string name, address value) returns(bool res)
func (_UserGroups *UserGroupsTransactorSession) UpdateUserGroupByAddress(name string, value common.Address) (*types.Transaction, error) {
	return _UserGroups.Contract.UpdateUserGroupByAddress(&_UserGroups.TransactOpts, name, value)
}

// UpdateUserGroupWithNewGroup is a paid mutator transaction binding the contract method 0x16a94bbf.
//
// Solidity: function updateUserGroupWithNewGroup(string name) returns(bool res)
func (_UserGroups *UserGroupsTransactor) UpdateUserGroupWithNewGroup(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _UserGroups.contract.Transact(opts, "updateUserGroupWithNewGroup", name)
}

// UpdateUserGroupWithNewGroup is a paid mutator transaction binding the contract method 0x16a94bbf.
//
// Solidity: function updateUserGroupWithNewGroup(string name) returns(bool res)
func (_UserGroups *UserGroupsSession) UpdateUserGroupWithNewGroup(name string) (*types.Transaction, error) {
	return _UserGroups.Contract.UpdateUserGroupWithNewGroup(&_UserGroups.TransactOpts, name)
}

// UpdateUserGroupWithNewGroup is a paid mutator transaction binding the contract method 0x16a94bbf.
//
// Solidity: function updateUserGroupWithNewGroup(string name) returns(bool res)
func (_UserGroups *UserGroupsTransactorSession) UpdateUserGroupWithNewGroup(name string) (*types.Transaction, error) {
	return _UserGroups.Contract.UpdateUserGroupWithNewGroup(&_UserGroups.TransactOpts, name)
}

// UserGroupCreator is a paid mutator transaction binding the contract method 0x2bc07c5a.
//
// Solidity: function userGroupCreator() returns(address addr)
func (_UserGroups *UserGroupsTransactor) UserGroupCreator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserGroups.contract.Transact(opts, "userGroupCreator")
}

// UserGroupCreator is a paid mutator transaction binding the contract method 0x2bc07c5a.
//
// Solidity: function userGroupCreator() returns(address addr)
func (_UserGroups *UserGroupsSession) UserGroupCreator() (*types.Transaction, error) {
	return _UserGroups.Contract.UserGroupCreator(&_UserGroups.TransactOpts)
}

// UserGroupCreator is a paid mutator transaction binding the contract method 0x2bc07c5a.
//
// Solidity: function userGroupCreator() returns(address addr)
func (_UserGroups *UserGroupsTransactorSession) UserGroupCreator() (*types.Transaction, error) {
	return _UserGroups.Contract.UserGroupCreator(&_UserGroups.TransactOpts)
}

// UserGroupsDeleteGroupByNameIterator is returned from FilterDeleteGroupByName and is used to iterate over the raw logs and unpacked data for DeleteGroupByName events raised by the UserGroups contract.
type UserGroupsDeleteGroupByNameIterator struct {
	Event *UserGroupsDeleteGroupByName // Event containing the contract specifics and raw log

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
func (it *UserGroupsDeleteGroupByNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserGroupsDeleteGroupByName)
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
		it.Event = new(UserGroupsDeleteGroupByName)
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
func (it *UserGroupsDeleteGroupByNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserGroupsDeleteGroupByNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserGroupsDeleteGroupByName represents a DeleteGroupByName event raised by the UserGroups contract.
type UserGroupsDeleteGroupByName struct {
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeleteGroupByName is a free log retrieval operation binding the contract event 0x0d21241cbf9584f265935b748e8240b0397928e543a3d915fa70471f722c8ce3.
//
// Solidity: event _deleteGroupByName(string name)
func (_UserGroups *UserGroupsFilterer) FilterDeleteGroupByName(opts *bind.FilterOpts) (*UserGroupsDeleteGroupByNameIterator, error) {

	logs, sub, err := _UserGroups.contract.FilterLogs(opts, "_deleteGroupByName")
	if err != nil {
		return nil, err
	}
	return &UserGroupsDeleteGroupByNameIterator{contract: _UserGroups.contract, event: "_deleteGroupByName", logs: logs, sub: sub}, nil
}

// WatchDeleteGroupByName is a free log subscription operation binding the contract event 0x0d21241cbf9584f265935b748e8240b0397928e543a3d915fa70471f722c8ce3.
//
// Solidity: event _deleteGroupByName(string name)
func (_UserGroups *UserGroupsFilterer) WatchDeleteGroupByName(opts *bind.WatchOpts, sink chan<- *UserGroupsDeleteGroupByName) (event.Subscription, error) {

	logs, sub, err := _UserGroups.contract.WatchLogs(opts, "_deleteGroupByName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserGroupsDeleteGroupByName)
				if err := _UserGroups.contract.UnpackLog(event, "_deleteGroupByName", log); err != nil {
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

// UserGroupsDeleteOwnerRecIterator is returned from FilterDeleteOwnerRec and is used to iterate over the raw logs and unpacked data for DeleteOwnerRec events raised by the UserGroups contract.
type UserGroupsDeleteOwnerRecIterator struct {
	Event *UserGroupsDeleteOwnerRec // Event containing the contract specifics and raw log

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
func (it *UserGroupsDeleteOwnerRecIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserGroupsDeleteOwnerRec)
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
		it.Event = new(UserGroupsDeleteOwnerRec)
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
func (it *UserGroupsDeleteOwnerRecIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserGroupsDeleteOwnerRecIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserGroupsDeleteOwnerRec represents a DeleteOwnerRec event raised by the UserGroups contract.
type UserGroupsDeleteOwnerRec struct {
	Owner common.Address
	Name  string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDeleteOwnerRec is a free log retrieval operation binding the contract event 0x9064673ca868275dc97f9fe659c8421d0fa89b07a8578773dd156336d186c3d3.
//
// Solidity: event _deleteOwnerRec(address owner, string name)
func (_UserGroups *UserGroupsFilterer) FilterDeleteOwnerRec(opts *bind.FilterOpts) (*UserGroupsDeleteOwnerRecIterator, error) {

	logs, sub, err := _UserGroups.contract.FilterLogs(opts, "_deleteOwnerRec")
	if err != nil {
		return nil, err
	}
	return &UserGroupsDeleteOwnerRecIterator{contract: _UserGroups.contract, event: "_deleteOwnerRec", logs: logs, sub: sub}, nil
}

// WatchDeleteOwnerRec is a free log subscription operation binding the contract event 0x9064673ca868275dc97f9fe659c8421d0fa89b07a8578773dd156336d186c3d3.
//
// Solidity: event _deleteOwnerRec(address owner, string name)
func (_UserGroups *UserGroupsFilterer) WatchDeleteOwnerRec(opts *bind.WatchOpts, sink chan<- *UserGroupsDeleteOwnerRec) (event.Subscription, error) {

	logs, sub, err := _UserGroups.contract.WatchLogs(opts, "_deleteOwnerRec")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserGroupsDeleteOwnerRec)
				if err := _UserGroups.contract.UnpackLog(event, "_deleteOwnerRec", log); err != nil {
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

// UserGroupsUpdateUserGroupIterator is returned from FilterUpdateUserGroup and is used to iterate over the raw logs and unpacked data for UpdateUserGroup events raised by the UserGroups contract.
type UserGroupsUpdateUserGroupIterator struct {
	Event *UserGroupsUpdateUserGroup // Event containing the contract specifics and raw log

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
func (it *UserGroupsUpdateUserGroupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserGroupsUpdateUserGroup)
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
		it.Event = new(UserGroupsUpdateUserGroup)
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
func (it *UserGroupsUpdateUserGroupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserGroupsUpdateUserGroupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserGroupsUpdateUserGroup represents a UpdateUserGroup event raised by the UserGroups contract.
type UserGroupsUpdateUserGroup struct {
	Add   common.Address
	Name  string
	Value common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdateUserGroup is a free log retrieval operation binding the contract event 0x9b4a33bb61ca8ea8f298d84955207994837a4921e3ed0741753f5ca9a3456d9d.
//
// Solidity: event _updateUserGroup(address add, string name, address value)
func (_UserGroups *UserGroupsFilterer) FilterUpdateUserGroup(opts *bind.FilterOpts) (*UserGroupsUpdateUserGroupIterator, error) {

	logs, sub, err := _UserGroups.contract.FilterLogs(opts, "_updateUserGroup")
	if err != nil {
		return nil, err
	}
	return &UserGroupsUpdateUserGroupIterator{contract: _UserGroups.contract, event: "_updateUserGroup", logs: logs, sub: sub}, nil
}

// WatchUpdateUserGroup is a free log subscription operation binding the contract event 0x9b4a33bb61ca8ea8f298d84955207994837a4921e3ed0741753f5ca9a3456d9d.
//
// Solidity: event _updateUserGroup(address add, string name, address value)
func (_UserGroups *UserGroupsFilterer) WatchUpdateUserGroup(opts *bind.WatchOpts, sink chan<- *UserGroupsUpdateUserGroup) (event.Subscription, error) {

	logs, sub, err := _UserGroups.contract.WatchLogs(opts, "_updateUserGroup")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserGroupsUpdateUserGroup)
				if err := _UserGroups.contract.UnpackLog(event, "_updateUserGroup", log); err != nil {
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

// UserGroupsGroupsIterator is returned from FilterGroups and is used to iterate over the raw logs and unpacked data for Groups events raised by the UserGroups contract.
type UserGroupsGroupsIterator struct {
	Event *UserGroupsGroups // Event containing the contract specifics and raw log

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
func (it *UserGroupsGroupsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserGroupsGroups)
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
		it.Event = new(UserGroupsGroups)
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
func (it *UserGroupsGroupsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserGroupsGroupsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserGroupsGroups represents a Groups event raised by the UserGroups contract.
type UserGroupsGroups struct {
	Name  string
	Value common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGroups is a free log retrieval operation binding the contract event 0xc2819d715e7e323f52575996fddf7784aa001fb74803c4fff38371cb694edd8c.
//
// Solidity: event groups(string name, address value)
func (_UserGroups *UserGroupsFilterer) FilterGroups(opts *bind.FilterOpts) (*UserGroupsGroupsIterator, error) {

	logs, sub, err := _UserGroups.contract.FilterLogs(opts, "groups")
	if err != nil {
		return nil, err
	}
	return &UserGroupsGroupsIterator{contract: _UserGroups.contract, event: "groups", logs: logs, sub: sub}, nil
}

// WatchGroups is a free log subscription operation binding the contract event 0xc2819d715e7e323f52575996fddf7784aa001fb74803c4fff38371cb694edd8c.
//
// Solidity: event groups(string name, address value)
func (_UserGroups *UserGroupsFilterer) WatchGroups(opts *bind.WatchOpts, sink chan<- *UserGroupsGroups) (event.Subscription, error) {

	logs, sub, err := _UserGroups.contract.WatchLogs(opts, "groups")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserGroupsGroups)
				if err := _UserGroups.contract.UnpackLog(event, "groups", log); err != nil {
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

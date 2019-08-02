// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dictionary

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

// DictionaryABI is the input ABI used to generate the binding from.
const DictionaryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"filename\",\"type\":\"string\"}],\"name\":\"getFileInfo\",\"outputs\":[{\"name\":\"isExit\",\"type\":\"bool\"},{\"name\":\"fname\",\"type\":\"string\"},{\"name\":\"fileType\",\"type\":\"uint256\"},{\"name\":\"createTime\",\"type\":\"uint256\"},{\"name\":\"changeTime\",\"type\":\"uint256\"},{\"name\":\"accessControl\",\"type\":\"address\"},{\"name\":\"fSize\",\"type\":\"uint256\"},{\"name\":\"isRec\",\"type\":\"bool\"},{\"name\":\"subDic\",\"type\":\"address\"},{\"name\":\"fileHash\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"name\":\"getNextFile\",\"outputs\":[{\"name\":\"r_keyIndex\",\"type\":\"uint256\"},{\"name\":\"fname\",\"type\":\"string\"},{\"name\":\"fileType\",\"type\":\"uint256\"},{\"name\":\"createTime\",\"type\":\"uint256\"},{\"name\":\"changeTime\",\"type\":\"uint256\"},{\"name\":\"accessControl\",\"type\":\"address\"},{\"name\":\"fSize\",\"type\":\"uint256\"},{\"name\":\"isRec\",\"type\":\"bool\"},{\"name\":\"subDic\",\"type\":\"address\"},{\"name\":\"fileHash\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"changetime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"address\"}],\"name\":\"getUserPerssion\",\"outputs\":[{\"name\":\"read\",\"type\":\"bool\"},{\"name\":\"update\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accessType\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"uG\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"createtime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newName\",\"type\":\"string\"},{\"name\":\"newFileType\",\"type\":\"uint256\"},{\"name\":\"newUG\",\"type\":\"address\"},{\"name\":\"newFileSize\",\"type\":\"uint256\"},{\"name\":\"newIsRec\",\"type\":\"bool\"},{\"name\":\"newsubDic\",\"type\":\"address\"},{\"name\":\"newFileBlockHash\",\"type\":\"string\"}],\"name\":\"addOrUpdateFile\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"filename\",\"type\":\"string\"}],\"name\":\"removeFile\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newUG\",\"type\":\"address\"}],\"name\":\"updateAccessControl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// DictionaryBin is the compiled bytecode used for deploying new contracts.
const DictionaryBin = `0x60806040526000805460ff60a01b1916905534801561001d57600080fd5b50604051611b3d380380611b3d8339818101604052602081101561004057600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055611acb806100726000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063b55ec18a11610071578063b55ec18a1461039e578063b8dd1b82146103ba578063e0175ec5146103c2578063e7106101146103ca578063f1afe04d1461051d578063f89488db146105c1576100a9565b806302d05d3f146100ae5780630cc9c17b146100d2578063147b6195146102a3578063703f7f3e14610343578063878f70c91461035d575b600080fd5b6100b66105e9565b604080516001600160a01b039092168252519081900360200190f35b610176600480360360208110156100e857600080fd5b810190602081018135600160201b81111561010257600080fd5b82018360208201111561011457600080fd5b803590602001918460018302840111600160201b8311171561013557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506105f8945050505050565b604080518b1515815290810189905260608101889052608081018790526001600160a01b0380871660a083015260c0820186905284151560e0830152831661010082015261014060208083018281528c51928401929092528b51610120840191610160850191908e019080838360005b838110156101fe5781810151838201526020016101e6565b50505050905090810190601f16801561022b5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561025e578181015183820152602001610246565b50505050905090810190601f16801561028b5780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060405180910390f35b6102c0600480360360208110156102b957600080fd5b50356108cc565b604080518b815290810189905260608101889052608081018790526001600160a01b0380871660a083015260c0820186905284151560e0830152831661010082015261014060208083018281528c51928401929092528b51610120840191610160850191908e0190808383600083156101fe5781810151838201526020016101e6565b61034b610d31565b60408051918252519081900360200190f35b6103836004803603602081101561037357600080fd5b50356001600160a01b0316610d37565b60408051921515835290151560208301528051918290030190f35b6103a6610e0b565b604080519115158252519081900360200190f35b6100b6610e1b565b61034b610e2a565b6103a6600480360360e08110156103e057600080fd5b810190602081018135600160201b8111156103fa57600080fd5b82018360208201111561040c57600080fd5b803590602001918460018302840111600160201b8311171561042d57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092958435956001600160a01b036020870135811696604081013596506060810135151595506080810135909116935060c081019060a00135600160201b8111156104a957600080fd5b8201836020820111156104bb57600080fd5b803590602001918460018302840111600160201b831117156104dc57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610e30945050505050565b6103a66004803603602081101561053357600080fd5b810190602081018135600160201b81111561054d57600080fd5b82018360208201111561055f57600080fd5b803590602001918460018302840111600160201b8311171561058057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611137945050505050565b6105e7600480360360208110156105d757600080fd5b50356001600160a01b03166112f7565b005b6000546001600160a01b031681565b60006060600080600080600080600060606106116118d0565b60036000018c6040518082805190602001908083835b602083106106465780518252601f199092019160209182019101610627565b518151600019602094850361010090810a8201928316921993909316919091179092529490920196875260408051978890038201882060019081018054610140601f60029483161590990290960116919091049586018390049092028801830190526101208701848152909550869450928592509084018282801561070c5780601f106106e15761010080835404028352916020019161070c565b820191906000526020600020905b8154815290600101906020018083116106ef57829003601f168201915b50505091835250506001828101546020808401919091526002808501546040808601919091526003860154606086015260048601546001600160a01b039081166080870152600587015460a0870152600687015460ff8116151560c08801526101009081900490911660e08701526007870180548351968116158302600019011693909304601f81018590048502860185019092528185529094019390918301828280156107fb5780601f106107d0576101008083540402835291602001916107fb565b820191906000526020600020905b8154815290600101906020018083116107de57829003601f168201915b5050505050815250509050600060036000018d6040518082805190602001908083835b6020831061083d5780518252601f19909201916020918201910161081e565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000015411816000015182602001518360400151846060015185608001518660a001518760c001518860e001518961010001518898508090509a509a509a509a509a509a509a509a509a509a50505b9193959799509193959799565b60006060600080600080600080600060605b6004548b1080156109105750600480548c9081106108f857fe5b600091825260209091206001600290920201015460ff165b15610920576001909a01996108de565b6040805160048054606060208202840181018552938301818152610a2e946003938593909160009085015b82821015610a165760008481526020908190206040805160028681029093018054600181161561010002600019011693909304601f8101859004909402810160609081018352918101848152909384928491908401828280156109ef5780601f106109c4576101008083540402835291602001916109ef565b820191906000526020600020905b8154815290600101906020018083116109d257829003601f168201915b505050918352505060019182015460ff16151560209182015291835292909201910161094b565b5050505081526020016002820154815250508c611872565b1515600114156108bf57606060036001018c81548110610a4a57fe5b60009182526020918290206002918202018054604080516001831615610100026000190190921693909304601f810185900485028201850190935282815292909190830182828015610add5780601f10610ab257610100808354040283529160200191610add565b820191906000526020600020905b815481529060010190602001808311610ac057829003601f168201915b50505050509050610aec6118d0565b6003600001826040518082805190602001908083835b60208310610b215780518252601f199092019160209182019101610b02565b518151600019602094850361010090810a8201928316921993909316919091179092529490920196875260408051978890038201882060019081018054610140601f600294831615909902909601169190910495860183900490920288018301905261012087018481529095508694509285925090840182828015610be75780601f10610bbc57610100808354040283529160200191610be7565b820191906000526020600020905b815481529060010190602001808311610bca57829003601f168201915b50505091835250506001828101546020808401919091526002808501546040808601919091526003860154606086015260048601546001600160a01b039081166080870152600587015460a0870152600687015460ff8116151560c08801526101009081900490911660e08701526007870180548351968116158302600019011693909304601f8101859004850286018501909252818552909401939091830182828015610cd65780601f10610cab57610100808354040283529160200191610cd6565b820191906000526020600020905b815481529060010190602001808311610cb957829003601f168201915b50505050508152505090508c816000015182602001518360400151846060015185608001518660a001518760c001518860e001518961010001518898508090509b509b509b509b509b509b509b509b509b509b5050506108bf565b60025481565b600080548190600160a01b900460ff16610d5657506001905080610e06565b6000546001600160a01b0384811691161415610d7757506001905080610e06565b6006546040805163878f70c960e01b81526001600160a01b03868116600483015291516000938493849391169163878f70c991602480820192606092909190829003018186803b158015610dca57600080fd5b505afa158015610dde573d6000803e3d6000fd5b505050506040513d6060811015610df457600080fd5b50805160209091015190955093505050505b915091565b600054600160a01b900460ff1681565b6006546001600160a01b031681565b60015481565b6000806000610e3e33610d37565b90925090506001811515141561112557600060036000018b6040518082805190602001908083835b60208310610e855780518252601f199092019160209182019101610e66565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390206000015490506040518061012001604052808c81526020018b81526020014281526020014281526020018815156001151514610ef4578a610f01565b6006546001600160a01b03165b6001600160a01b031681526020018981526020018815158152602001876001600160a01b031681526020018681525060036000018c6040518082805190602001908083835b60208310610f655780518252601f199092019160209182019101610f46565b51815160209384036101000a600019018019909216911617905292019485525060405193849003810190932084518051600190920194610faa94508593500190611930565b506020828101516001830155604083015160028301556060830151600383015560808301516004830180546001600160a01b0319166001600160a01b0392831617905560a0840151600584015560c084015160068401805460e087015160ff1990911692151592909217610100600160a81b0319166101009290931682029290921790915583015180516110449260078501920190611930565b50508115905061105a576001935050505061112c565b600480549061106c90600183016119ae565b90508060010160036000018c6040518082805190602001908083835b602083106110a75780518252601f199092019160209182019101611088565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220929092555050600480548c9190839081106110e957fe5b9060005260206000209060020201600001908051906020019061110d929190611930565b5050600580546001908101909155925061112c915050565b6000925050505b979650505050505050565b600080600061114533610d37565b9092509050600181151514156112eb5760006003600001856040518082805190602001908083835b6020831061118c5780518252601f19909201916020918201910161116d565b51815160001960209485036101000a0190811690199190911617905292019485525060405193849003019092205492505050806111cf57600093505050506112f2565b6003600001856040518082805190602001908083835b602083106112045780518252601f1990920191602091820191016111e5565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922060008082559092509050600182018161124882826119df565b60006001830181905560028301819055600383018190556004830180546001600160a01b0319169055600583018190556006830180546001600160a81b03191690556112989060078401906119df565b505050506001600360010160018303815481106112b157fe5b600091825260209091206002909102016001908101805460ff1916921515929092179091556005805460001901905593506112f292505050565b6000925050505b919050565b6000546001600160a01b031633146113405760405162461bcd60e51b8152600401808060200182810382526024815260200180611a736024913960400191505060405180910390fd5b6000805460ff60a01b1916600160a01b178155600680546001600160a01b0384166001600160a01b0319909116179055604080516004805460606020820284018101855293830181815261147c9460039385939091889085015b828210156114655760008481526020908190206040805160028681029093018054600181161561010002600019011693909304601f81018590049094028101606090810183529181018481529093849284919084018282801561143e5780601f106114135761010080835404028352916020019161143e565b820191906000526020600020905b81548152906001019060200180831161142157829003601f168201915b505050918352505060019182015460ff16151560209182015291835292909201910161139a565b505050508152602001600282015481525050611879565b90505b604080516004805460606020820284018101855293830181815261158d946003938593909160009085015b828210156115755760008481526020908190206040805160028681029093018054600181161561010002600019011693909304601f81018590049094028101606090810183529181018481529093849284919084018282801561154e5780601f106115235761010080835404028352916020019161154e565b820191906000526020600020905b81548152906001019060200180831161153157829003601f168201915b505050918352505060019182015460ff1615156020918201529183529290920191016114aa565b50505050815260200160028201548152505082611872565b1561186e576060600360010182815481106115a457fe5b60009182526020918290206002918202018054604080516001831615610100026000190190921693909304601f8101859004850282018501909352828152929091908301828280156116375780601f1061160c57610100808354040283529160200191611637565b820191906000526020600020905b81548152906001019060200180831161161a57829003601f168201915b5050505050905060006003600001826040518082805190602001908083835b602083106116755780518252601f199092019160209182019101611656565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922060078101546001918201945060ff161515141591506117579050576006546004820180546001600160a01b0319166001600160a01b0390921691909117905560018101546117575760068101546040805163f89488db60e01b81526001600160a01b03878116600483015291516101009093049091169163f89488db9160248082019260009290919082900301818387803b15801561173e57600080fd5b505af1158015611752573d6000803e3d6000fd5b505050505b50506040805160048054606060208202840181018552938301818152611867946003938593909160009085015b8282101561184f5760008481526020908190206040805160028681029093018054600181161561010002600019011693909304601f8101859004909402810160609081018352918101848152909384928491908401828280156118285780601f106117fd57610100808354040283529160200191611828565b820191906000526020600020905b81548152906001019060200180831161180b57829003601f168201915b505050918352505060019182015460ff161515602091820152918352929092019101611784565b5050505081526020016002820154815250508261188d565b905061147f565b5050565b9051511190565b60006118878260001961188d565b92915050565b60010160005b825151821080156118ba575082518051839081106118ad57fe5b6020026020010151602001515b156118ca57600190910190611893565b50919050565b6040518061012001604052806060815260200160008152602001600081526020016000815260200160006001600160a01b031681526020016000815260200160001515815260200160006001600160a01b03168152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061197157805160ff191683800117855561199e565b8280016001018555821561199e579182015b8281111561199e578251825591602001919060010190611983565b506119aa929150611a26565b5090565b8154818355818111156119da576002028160020283600052602060002091820191016119da9190611a43565b505050565b50805460018160011615610100020316600290046000825580601f10611a055750611a23565b601f016020900490600052602060002090810190611a239190611a26565b50565b611a4091905b808211156119aa5760008155600101611a2c565b90565b611a4091905b808211156119aa576000611a5d82826119df565b5060018101805460ff19169055600201611a4956fe4f6e6c792063726561746f722063616e2063616c6c20746869732066756e6374696f6e2ea265627a7a72305820e6e6ac4cb21c28d199339c2f5f05c702e7137177b44d6fcfd169f3051cfd305764736f6c634300050a0032`

// DeployDictionary deploys a new Ethereum contract, binding an instance of Dictionary to it.
func DeployDictionary(auth *bind.TransactOpts, backend bind.ContractBackend, sender common.Address) (common.Address, *types.Transaction, *Dictionary, error) {
	parsed, err := abi.JSON(strings.NewReader(DictionaryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DictionaryBin), backend, sender)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dictionary{DictionaryCaller: DictionaryCaller{contract: contract}, DictionaryTransactor: DictionaryTransactor{contract: contract}, DictionaryFilterer: DictionaryFilterer{contract: contract}}, nil
}

// Dictionary is an auto generated Go binding around an Ethereum contract.
type Dictionary struct {
	DictionaryCaller     // Read-only binding to the contract
	DictionaryTransactor // Write-only binding to the contract
	DictionaryFilterer   // Log filterer for contract events
}

// DictionaryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DictionaryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DictionaryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DictionaryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DictionaryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DictionaryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DictionarySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DictionarySession struct {
	Contract     *Dictionary       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DictionaryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DictionaryCallerSession struct {
	Contract *DictionaryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DictionaryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DictionaryTransactorSession struct {
	Contract     *DictionaryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DictionaryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DictionaryRaw struct {
	Contract *Dictionary // Generic contract binding to access the raw methods on
}

// DictionaryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DictionaryCallerRaw struct {
	Contract *DictionaryCaller // Generic read-only contract binding to access the raw methods on
}

// DictionaryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DictionaryTransactorRaw struct {
	Contract *DictionaryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDictionary creates a new instance of Dictionary, bound to a specific deployed contract.
func NewDictionary(address common.Address, backend bind.ContractBackend) (*Dictionary, error) {
	contract, err := bindDictionary(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dictionary{DictionaryCaller: DictionaryCaller{contract: contract}, DictionaryTransactor: DictionaryTransactor{contract: contract}, DictionaryFilterer: DictionaryFilterer{contract: contract}}, nil
}

// NewDictionaryCaller creates a new read-only instance of Dictionary, bound to a specific deployed contract.
func NewDictionaryCaller(address common.Address, caller bind.ContractCaller) (*DictionaryCaller, error) {
	contract, err := bindDictionary(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DictionaryCaller{contract: contract}, nil
}

// NewDictionaryTransactor creates a new write-only instance of Dictionary, bound to a specific deployed contract.
func NewDictionaryTransactor(address common.Address, transactor bind.ContractTransactor) (*DictionaryTransactor, error) {
	contract, err := bindDictionary(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DictionaryTransactor{contract: contract}, nil
}

// NewDictionaryFilterer creates a new log filterer instance of Dictionary, bound to a specific deployed contract.
func NewDictionaryFilterer(address common.Address, filterer bind.ContractFilterer) (*DictionaryFilterer, error) {
	contract, err := bindDictionary(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DictionaryFilterer{contract: contract}, nil
}

// bindDictionary binds a generic wrapper to an already deployed contract.
func bindDictionary(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DictionaryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dictionary *DictionaryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dictionary.Contract.DictionaryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dictionary *DictionaryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dictionary.Contract.DictionaryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dictionary *DictionaryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dictionary.Contract.DictionaryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dictionary *DictionaryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dictionary.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dictionary *DictionaryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dictionary.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dictionary *DictionaryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dictionary.Contract.contract.Transact(opts, method, params...)
}

// AccessType is a free data retrieval call binding the contract method 0xb55ec18a.
//
// Solidity: function accessType() constant returns(bool)
func (_Dictionary *DictionaryCaller) AccessType(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dictionary.contract.Call(opts, out, "accessType")
	return *ret0, err
}

// AccessType is a free data retrieval call binding the contract method 0xb55ec18a.
//
// Solidity: function accessType() constant returns(bool)
func (_Dictionary *DictionarySession) AccessType() (bool, error) {
	return _Dictionary.Contract.AccessType(&_Dictionary.CallOpts)
}

// AccessType is a free data retrieval call binding the contract method 0xb55ec18a.
//
// Solidity: function accessType() constant returns(bool)
func (_Dictionary *DictionaryCallerSession) AccessType() (bool, error) {
	return _Dictionary.Contract.AccessType(&_Dictionary.CallOpts)
}

// Changetime is a free data retrieval call binding the contract method 0x703f7f3e.
//
// Solidity: function changetime() constant returns(uint256)
func (_Dictionary *DictionaryCaller) Changetime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dictionary.contract.Call(opts, out, "changetime")
	return *ret0, err
}

// Changetime is a free data retrieval call binding the contract method 0x703f7f3e.
//
// Solidity: function changetime() constant returns(uint256)
func (_Dictionary *DictionarySession) Changetime() (*big.Int, error) {
	return _Dictionary.Contract.Changetime(&_Dictionary.CallOpts)
}

// Changetime is a free data retrieval call binding the contract method 0x703f7f3e.
//
// Solidity: function changetime() constant returns(uint256)
func (_Dictionary *DictionaryCallerSession) Changetime() (*big.Int, error) {
	return _Dictionary.Contract.Changetime(&_Dictionary.CallOpts)
}

// Createtime is a free data retrieval call binding the contract method 0xe0175ec5.
//
// Solidity: function createtime() constant returns(uint256)
func (_Dictionary *DictionaryCaller) Createtime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dictionary.contract.Call(opts, out, "createtime")
	return *ret0, err
}

// Createtime is a free data retrieval call binding the contract method 0xe0175ec5.
//
// Solidity: function createtime() constant returns(uint256)
func (_Dictionary *DictionarySession) Createtime() (*big.Int, error) {
	return _Dictionary.Contract.Createtime(&_Dictionary.CallOpts)
}

// Createtime is a free data retrieval call binding the contract method 0xe0175ec5.
//
// Solidity: function createtime() constant returns(uint256)
func (_Dictionary *DictionaryCallerSession) Createtime() (*big.Int, error) {
	return _Dictionary.Contract.Createtime(&_Dictionary.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Dictionary *DictionaryCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dictionary.contract.Call(opts, out, "creator")
	return *ret0, err
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Dictionary *DictionarySession) Creator() (common.Address, error) {
	return _Dictionary.Contract.Creator(&_Dictionary.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() constant returns(address)
func (_Dictionary *DictionaryCallerSession) Creator() (common.Address, error) {
	return _Dictionary.Contract.Creator(&_Dictionary.CallOpts)
}

// GetFileInfo is a free data retrieval call binding the contract method 0x0cc9c17b.
//
// Solidity: function getFileInfo(string filename) constant returns(bool isExit, string fname, uint256 fileType, uint256 createTime, uint256 changeTime, address accessControl, uint256 fSize, bool isRec, address subDic, string fileHash)
func (_Dictionary *DictionaryCaller) GetFileInfo(opts *bind.CallOpts, filename string) (struct {
	IsExit        bool
	Fname         string
	FileType      *big.Int
	CreateTime    *big.Int
	ChangeTime    *big.Int
	AccessControl common.Address
	FSize         *big.Int
	IsRec         bool
	SubDic        common.Address
	FileHash      string
}, error) {
	ret := new(struct {
		IsExit        bool
		Fname         string
		FileType      *big.Int
		CreateTime    *big.Int
		ChangeTime    *big.Int
		AccessControl common.Address
		FSize         *big.Int
		IsRec         bool
		SubDic        common.Address
		FileHash      string
	})
	out := ret
	err := _Dictionary.contract.Call(opts, out, "getFileInfo", filename)
	return *ret, err
}

// GetFileInfo is a free data retrieval call binding the contract method 0x0cc9c17b.
//
// Solidity: function getFileInfo(string filename) constant returns(bool isExit, string fname, uint256 fileType, uint256 createTime, uint256 changeTime, address accessControl, uint256 fSize, bool isRec, address subDic, string fileHash)
func (_Dictionary *DictionarySession) GetFileInfo(filename string) (struct {
	IsExit        bool
	Fname         string
	FileType      *big.Int
	CreateTime    *big.Int
	ChangeTime    *big.Int
	AccessControl common.Address
	FSize         *big.Int
	IsRec         bool
	SubDic        common.Address
	FileHash      string
}, error) {
	return _Dictionary.Contract.GetFileInfo(&_Dictionary.CallOpts, filename)
}

// GetFileInfo is a free data retrieval call binding the contract method 0x0cc9c17b.
//
// Solidity: function getFileInfo(string filename) constant returns(bool isExit, string fname, uint256 fileType, uint256 createTime, uint256 changeTime, address accessControl, uint256 fSize, bool isRec, address subDic, string fileHash)
func (_Dictionary *DictionaryCallerSession) GetFileInfo(filename string) (struct {
	IsExit        bool
	Fname         string
	FileType      *big.Int
	CreateTime    *big.Int
	ChangeTime    *big.Int
	AccessControl common.Address
	FSize         *big.Int
	IsRec         bool
	SubDic        common.Address
	FileHash      string
}, error) {
	return _Dictionary.Contract.GetFileInfo(&_Dictionary.CallOpts, filename)
}

// GetNextFile is a free data retrieval call binding the contract method 0x147b6195.
//
// Solidity: function getNextFile(uint256 keyIndex) constant returns(uint256 r_keyIndex, string fname, uint256 fileType, uint256 createTime, uint256 changeTime, address accessControl, uint256 fSize, bool isRec, address subDic, string fileHash)
func (_Dictionary *DictionaryCaller) GetNextFile(opts *bind.CallOpts, keyIndex *big.Int) (struct {
	RKeyIndex     *big.Int
	Fname         string
	FileType      *big.Int
	CreateTime    *big.Int
	ChangeTime    *big.Int
	AccessControl common.Address
	FSize         *big.Int
	IsRec         bool
	SubDic        common.Address
	FileHash      string
}, error) {
	ret := new(struct {
		RKeyIndex     *big.Int
		Fname         string
		FileType      *big.Int
		CreateTime    *big.Int
		ChangeTime    *big.Int
		AccessControl common.Address
		FSize         *big.Int
		IsRec         bool
		SubDic        common.Address
		FileHash      string
	})
	out := ret
	err := _Dictionary.contract.Call(opts, out, "getNextFile", keyIndex)
	return *ret, err
}

// GetNextFile is a free data retrieval call binding the contract method 0x147b6195.
//
// Solidity: function getNextFile(uint256 keyIndex) constant returns(uint256 r_keyIndex, string fname, uint256 fileType, uint256 createTime, uint256 changeTime, address accessControl, uint256 fSize, bool isRec, address subDic, string fileHash)
func (_Dictionary *DictionarySession) GetNextFile(keyIndex *big.Int) (struct {
	RKeyIndex     *big.Int
	Fname         string
	FileType      *big.Int
	CreateTime    *big.Int
	ChangeTime    *big.Int
	AccessControl common.Address
	FSize         *big.Int
	IsRec         bool
	SubDic        common.Address
	FileHash      string
}, error) {
	return _Dictionary.Contract.GetNextFile(&_Dictionary.CallOpts, keyIndex)
}

// GetNextFile is a free data retrieval call binding the contract method 0x147b6195.
//
// Solidity: function getNextFile(uint256 keyIndex) constant returns(uint256 r_keyIndex, string fname, uint256 fileType, uint256 createTime, uint256 changeTime, address accessControl, uint256 fSize, bool isRec, address subDic, string fileHash)
func (_Dictionary *DictionaryCallerSession) GetNextFile(keyIndex *big.Int) (struct {
	RKeyIndex     *big.Int
	Fname         string
	FileType      *big.Int
	CreateTime    *big.Int
	ChangeTime    *big.Int
	AccessControl common.Address
	FSize         *big.Int
	IsRec         bool
	SubDic        common.Address
	FileHash      string
}, error) {
	return _Dictionary.Contract.GetNextFile(&_Dictionary.CallOpts, keyIndex)
}

// GetUserPerssion is a free data retrieval call binding the contract method 0x878f70c9.
//
// Solidity: function getUserPerssion(address key) constant returns(bool read, bool update)
func (_Dictionary *DictionaryCaller) GetUserPerssion(opts *bind.CallOpts, key common.Address) (struct {
	Read   bool
	Update bool
}, error) {
	ret := new(struct {
		Read   bool
		Update bool
	})
	out := ret
	err := _Dictionary.contract.Call(opts, out, "getUserPerssion", key)
	return *ret, err
}

// GetUserPerssion is a free data retrieval call binding the contract method 0x878f70c9.
//
// Solidity: function getUserPerssion(address key) constant returns(bool read, bool update)
func (_Dictionary *DictionarySession) GetUserPerssion(key common.Address) (struct {
	Read   bool
	Update bool
}, error) {
	return _Dictionary.Contract.GetUserPerssion(&_Dictionary.CallOpts, key)
}

// GetUserPerssion is a free data retrieval call binding the contract method 0x878f70c9.
//
// Solidity: function getUserPerssion(address key) constant returns(bool read, bool update)
func (_Dictionary *DictionaryCallerSession) GetUserPerssion(key common.Address) (struct {
	Read   bool
	Update bool
}, error) {
	return _Dictionary.Contract.GetUserPerssion(&_Dictionary.CallOpts, key)
}

// UG is a free data retrieval call binding the contract method 0xb8dd1b82.
//
// Solidity: function uG() constant returns(address)
func (_Dictionary *DictionaryCaller) UG(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dictionary.contract.Call(opts, out, "uG")
	return *ret0, err
}

// UG is a free data retrieval call binding the contract method 0xb8dd1b82.
//
// Solidity: function uG() constant returns(address)
func (_Dictionary *DictionarySession) UG() (common.Address, error) {
	return _Dictionary.Contract.UG(&_Dictionary.CallOpts)
}

// UG is a free data retrieval call binding the contract method 0xb8dd1b82.
//
// Solidity: function uG() constant returns(address)
func (_Dictionary *DictionaryCallerSession) UG() (common.Address, error) {
	return _Dictionary.Contract.UG(&_Dictionary.CallOpts)
}

// AddOrUpdateFile is a paid mutator transaction binding the contract method 0xe7106101.
//
// Solidity: function addOrUpdateFile(string newName, uint256 newFileType, address newUG, uint256 newFileSize, bool newIsRec, address newsubDic, string newFileBlockHash) returns(bool res)
func (_Dictionary *DictionaryTransactor) AddOrUpdateFile(opts *bind.TransactOpts, newName string, newFileType *big.Int, newUG common.Address, newFileSize *big.Int, newIsRec bool, newsubDic common.Address, newFileBlockHash string) (*types.Transaction, error) {
	return _Dictionary.contract.Transact(opts, "addOrUpdateFile", newName, newFileType, newUG, newFileSize, newIsRec, newsubDic, newFileBlockHash)
}

// AddOrUpdateFile is a paid mutator transaction binding the contract method 0xe7106101.
//
// Solidity: function addOrUpdateFile(string newName, uint256 newFileType, address newUG, uint256 newFileSize, bool newIsRec, address newsubDic, string newFileBlockHash) returns(bool res)
func (_Dictionary *DictionarySession) AddOrUpdateFile(newName string, newFileType *big.Int, newUG common.Address, newFileSize *big.Int, newIsRec bool, newsubDic common.Address, newFileBlockHash string) (*types.Transaction, error) {
	return _Dictionary.Contract.AddOrUpdateFile(&_Dictionary.TransactOpts, newName, newFileType, newUG, newFileSize, newIsRec, newsubDic, newFileBlockHash)
}

// AddOrUpdateFile is a paid mutator transaction binding the contract method 0xe7106101.
//
// Solidity: function addOrUpdateFile(string newName, uint256 newFileType, address newUG, uint256 newFileSize, bool newIsRec, address newsubDic, string newFileBlockHash) returns(bool res)
func (_Dictionary *DictionaryTransactorSession) AddOrUpdateFile(newName string, newFileType *big.Int, newUG common.Address, newFileSize *big.Int, newIsRec bool, newsubDic common.Address, newFileBlockHash string) (*types.Transaction, error) {
	return _Dictionary.Contract.AddOrUpdateFile(&_Dictionary.TransactOpts, newName, newFileType, newUG, newFileSize, newIsRec, newsubDic, newFileBlockHash)
}

// RemoveFile is a paid mutator transaction binding the contract method 0xf1afe04d.
//
// Solidity: function removeFile(string filename) returns(bool res)
func (_Dictionary *DictionaryTransactor) RemoveFile(opts *bind.TransactOpts, filename string) (*types.Transaction, error) {
	return _Dictionary.contract.Transact(opts, "removeFile", filename)
}

// RemoveFile is a paid mutator transaction binding the contract method 0xf1afe04d.
//
// Solidity: function removeFile(string filename) returns(bool res)
func (_Dictionary *DictionarySession) RemoveFile(filename string) (*types.Transaction, error) {
	return _Dictionary.Contract.RemoveFile(&_Dictionary.TransactOpts, filename)
}

// RemoveFile is a paid mutator transaction binding the contract method 0xf1afe04d.
//
// Solidity: function removeFile(string filename) returns(bool res)
func (_Dictionary *DictionaryTransactorSession) RemoveFile(filename string) (*types.Transaction, error) {
	return _Dictionary.Contract.RemoveFile(&_Dictionary.TransactOpts, filename)
}

// UpdateAccessControl is a paid mutator transaction binding the contract method 0xf89488db.
//
// Solidity: function updateAccessControl(address newUG) returns()
func (_Dictionary *DictionaryTransactor) UpdateAccessControl(opts *bind.TransactOpts, newUG common.Address) (*types.Transaction, error) {
	return _Dictionary.contract.Transact(opts, "updateAccessControl", newUG)
}

// UpdateAccessControl is a paid mutator transaction binding the contract method 0xf89488db.
//
// Solidity: function updateAccessControl(address newUG) returns()
func (_Dictionary *DictionarySession) UpdateAccessControl(newUG common.Address) (*types.Transaction, error) {
	return _Dictionary.Contract.UpdateAccessControl(&_Dictionary.TransactOpts, newUG)
}

// UpdateAccessControl is a paid mutator transaction binding the contract method 0xf89488db.
//
// Solidity: function updateAccessControl(address newUG) returns()
func (_Dictionary *DictionaryTransactorSession) UpdateAccessControl(newUG common.Address) (*types.Transaction, error) {
	return _Dictionary.Contract.UpdateAccessControl(&_Dictionary.TransactOpts, newUG)
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

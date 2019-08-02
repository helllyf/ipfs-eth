// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package router

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
const DictionaryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newName\",\"type\":\"string\"},{\"name\":\"newFileType\",\"type\":\"uint256\"},{\"name\":\"newUG\",\"type\":\"address\"},{\"name\":\"newFileSize\",\"type\":\"uint256\"},{\"name\":\"newIsRec\",\"type\":\"bool\"},{\"name\":\"newsubDic\",\"type\":\"address\"},{\"name\":\"newFileBlockHash\",\"type\":\"uint256\"}],\"name\":\"addFile\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delIndex\",\"type\":\"uint256\"}],\"name\":\"removeFile\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delIndex\",\"type\":\"uint256\"},{\"name\":\"newName\",\"type\":\"string\"},{\"name\":\"newFileType\",\"type\":\"uint256\"},{\"name\":\"newUG\",\"type\":\"address\"},{\"name\":\"newFileSize\",\"type\":\"uint256\"},{\"name\":\"newIsRec\",\"type\":\"bool\"},{\"name\":\"newsubDic\",\"type\":\"address\"},{\"name\":\"newFileBlockHash\",\"type\":\"uint256\"}],\"name\":\"updateFile\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"updateDicInfo\",\"outputs\":[{\"name\":\"res\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newUG\",\"type\":\"address\"}],\"name\":\"updateAccessControl\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// DictionaryBin is the compiled bytecode used for deploying new contracts.
const DictionaryBin = `0x608060405260006002556003805460ff1916905534801561001f57600080fd5b50604051610c62380380610c628339818101604052602081101561004257600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610bee806100746000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806302d05d3f1461006757806392a01a021461008b5780639926b03f1461016f578063d98f6d391461018c578063f13ab19614610264578063f89488db1461030a575b600080fd5b61006f610332565b604080516001600160a01b039092168252519081900360200190f35b61015b600480360360e08110156100a157600080fd5b8101906020810181356401000000008111156100bc57600080fd5b8201836020820111156100ce57600080fd5b803590602001918460018302840111640100000000831117156100f057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350506001600160a01b0360208301358116926040810135925060608101351515916080820135169060a00135610341565b604080519115158252519081900360200190f35b61015b6004803603602081101561018557600080fd5b50356104fd565b61015b60048036036101008110156101a357600080fd5b813591908101906040810160208201356401000000008111156101c557600080fd5b8201836020820111156101d757600080fd5b803590602001918460018302840111640100000000831117156101f957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350506001600160a01b0360208301358116926040810135925060608101351515916080820135169060a001356105de565b61015b6004803603602081101561027a57600080fd5b81019060208101813564010000000081111561029557600080fd5b8201836020820111156102a757600080fd5b803590602001918460018302840111640100000000831117156102c957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610789945050505050565b6103306004803603602081101561032057600080fd5b50356001600160a01b03166107c6565b005b6000546001600160a01b031681565b600080600061034f33610984565b9092509050600181151514156104eb57610367610a32565b6040518061012001604052808c81526020018b8152602001428152602001428152602001881515600115151461039d578a6103aa565b6007546001600160a01b03165b6001600160a01b03908116825260208083018c90528a15156040840152908916606083015260809091018790529082526000828201819052600680546001810180835591909252835180518051929486946009027ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f01938492610431928492910190610a53565b50602082810151600183810191909155604084015160028401556060840151600384015560808401516004840180546001600160a01b0319166001600160a01b0392831617905560a0850151600585015560c085015160068501805460e088015160ff1991821693151593909317610100600160a81b031916610100939094168302939093179055909401516007909301929092559390930151600892909201805490911691151591909117905594506104f29350505050565b6000925050505b979650505050505050565b600080600061050b33610984565b9092509050600181151514156105d2576006848154811061052857fe5b600091825260208220600990910201906105428282610ad1565b506000600182810182905560028301829055600383018290556004830180546001600160a01b031916905560058301829055600680840180546001600160a81b031916905560079093019190915581549091908690811061059f57fe5b906000526020600020906009020160080160006101000a81548160ff0219169083151502179055506001925050506105d9565b6000925050505b919050565b60008060006105ec33610984565b9150915060068b815481106105fd57fe5b600091825260209091206008600990920201015460ff161515600114156106295760009250505061077d565b60018115151415610776576040518061012001604052808b81526020018a8152602001428152602001428152602001871515600115151461066a5789610677565b6007546001600160a01b03165b6001600160a01b031681526020018881526020018715158152602001866001600160a01b031681526020018581525060068c815481106106b357fe5b6000918252602091829020835180516009909302909101926106da92849290910190610a53565b506020820151600182810191909155604083015160028301556060830151600383015560808301516004830180546001600160a01b0319166001600160a01b0392831617905560a0840151600584015560c084015160068401805460e087015160ff1990911692151592909217610100600160a81b0319166101009290931682029290921790915590920151600790910155925061077d915050565b6000925050505b98975050505050505050565b600080600061079733610984565b9092509050600181151514156105d25783516107ba906001906020870190610a53565b506001925050506105d9565b6000546001600160a01b0316331461080f5760405162461bcd60e51b8152600401808060200182810382526024815260200180610b966024913960400191505060405180910390fd5b6003805460ff19166001179055600780546001600160a01b0383166001600160a01b031990911617905560005b600654811015610980576006818154811061085357fe5b600091825260209091206006600990920201015460ff1615156001141561097857600754600680546001600160a01b03909216918390811061089157fe5b906000526020600020906009020160000160040160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600681815481106108d657fe5b90600052602060002090600902016000016001015460001415610978576006818154811061090057fe5b600091825260208220600660099092020101546040805163f89488db60e01b81526001600160a01b0386811660048301529151610100909304919091169263f89488db9260248084019382900301818387803b15801561095f57600080fd5b505af1158015610973573d6000803e3d6000fd5b505050505b60010161083c565b5050565b600354600090819060ff1661099e57506001905080610a2d565b6007546040805163878f70c960e01b81526001600160a01b03868116600483015291516000938493849391169163878f70c991602480820192606092909190829003018186803b1580156109f157600080fd5b505afa158015610a05573d6000803e3d6000fd5b505050506040513d6060811015610a1b57600080fd5b50805160209091015190955093505050505b915091565b604051806101400160405280610a46610b18565b8152600060209091015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a9457805160ff1916838001178555610ac1565b82800160010185558215610ac1579182015b82811115610ac1578251825591602001919060010190610aa6565b50610acd929150610b78565b5090565b50805460018160011615610100020316600290046000825580601f10610af75750610b15565b601f016020900490600052602060002090810190610b159190610b78565b50565b6040518061012001604052806060815260200160008152602001600081526020016000815260200160006001600160a01b031681526020016000815260200160001515815260200160006001600160a01b03168152602001600081525090565b610b9291905b80821115610acd5760008155600101610b7e565b9056fe4f6e6c792063726561746f722063616e2063616c6c20746869732066756e6374696f6e2ea265627a7a72305820a1dd2c6e5fde4f3706fbcc4b644af5f007704ae68a04c595efea096e321b12c164736f6c634300050a0032`

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

// AddFile is a paid mutator transaction binding the contract method 0x92a01a02.
//
// Solidity: function addFile(string newName, uint256 newFileType, address newUG, uint256 newFileSize, bool newIsRec, address newsubDic, uint256 newFileBlockHash) returns(bool res)
func (_Dictionary *DictionaryTransactor) AddFile(opts *bind.TransactOpts, newName string, newFileType *big.Int, newUG common.Address, newFileSize *big.Int, newIsRec bool, newsubDic common.Address, newFileBlockHash *big.Int) (*types.Transaction, error) {
	return _Dictionary.contract.Transact(opts, "addFile", newName, newFileType, newUG, newFileSize, newIsRec, newsubDic, newFileBlockHash)
}

// AddFile is a paid mutator transaction binding the contract method 0x92a01a02.
//
// Solidity: function addFile(string newName, uint256 newFileType, address newUG, uint256 newFileSize, bool newIsRec, address newsubDic, uint256 newFileBlockHash) returns(bool res)
func (_Dictionary *DictionarySession) AddFile(newName string, newFileType *big.Int, newUG common.Address, newFileSize *big.Int, newIsRec bool, newsubDic common.Address, newFileBlockHash *big.Int) (*types.Transaction, error) {
	return _Dictionary.Contract.AddFile(&_Dictionary.TransactOpts, newName, newFileType, newUG, newFileSize, newIsRec, newsubDic, newFileBlockHash)
}

// AddFile is a paid mutator transaction binding the contract method 0x92a01a02.
//
// Solidity: function addFile(string newName, uint256 newFileType, address newUG, uint256 newFileSize, bool newIsRec, address newsubDic, uint256 newFileBlockHash) returns(bool res)
func (_Dictionary *DictionaryTransactorSession) AddFile(newName string, newFileType *big.Int, newUG common.Address, newFileSize *big.Int, newIsRec bool, newsubDic common.Address, newFileBlockHash *big.Int) (*types.Transaction, error) {
	return _Dictionary.Contract.AddFile(&_Dictionary.TransactOpts, newName, newFileType, newUG, newFileSize, newIsRec, newsubDic, newFileBlockHash)
}

// RemoveFile is a paid mutator transaction binding the contract method 0x9926b03f.
//
// Solidity: function removeFile(uint256 delIndex) returns(bool res)
func (_Dictionary *DictionaryTransactor) RemoveFile(opts *bind.TransactOpts, delIndex *big.Int) (*types.Transaction, error) {
	return _Dictionary.contract.Transact(opts, "removeFile", delIndex)
}

// RemoveFile is a paid mutator transaction binding the contract method 0x9926b03f.
//
// Solidity: function removeFile(uint256 delIndex) returns(bool res)
func (_Dictionary *DictionarySession) RemoveFile(delIndex *big.Int) (*types.Transaction, error) {
	return _Dictionary.Contract.RemoveFile(&_Dictionary.TransactOpts, delIndex)
}

// RemoveFile is a paid mutator transaction binding the contract method 0x9926b03f.
//
// Solidity: function removeFile(uint256 delIndex) returns(bool res)
func (_Dictionary *DictionaryTransactorSession) RemoveFile(delIndex *big.Int) (*types.Transaction, error) {
	return _Dictionary.Contract.RemoveFile(&_Dictionary.TransactOpts, delIndex)
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

// UpdateDicInfo is a paid mutator transaction binding the contract method 0xf13ab196.
//
// Solidity: function updateDicInfo(string newName) returns(bool res)
func (_Dictionary *DictionaryTransactor) UpdateDicInfo(opts *bind.TransactOpts, newName string) (*types.Transaction, error) {
	return _Dictionary.contract.Transact(opts, "updateDicInfo", newName)
}

// UpdateDicInfo is a paid mutator transaction binding the contract method 0xf13ab196.
//
// Solidity: function updateDicInfo(string newName) returns(bool res)
func (_Dictionary *DictionarySession) UpdateDicInfo(newName string) (*types.Transaction, error) {
	return _Dictionary.Contract.UpdateDicInfo(&_Dictionary.TransactOpts, newName)
}

// UpdateDicInfo is a paid mutator transaction binding the contract method 0xf13ab196.
//
// Solidity: function updateDicInfo(string newName) returns(bool res)
func (_Dictionary *DictionaryTransactorSession) UpdateDicInfo(newName string) (*types.Transaction, error) {
	return _Dictionary.Contract.UpdateDicInfo(&_Dictionary.TransactOpts, newName)
}

// UpdateFile is a paid mutator transaction binding the contract method 0xd98f6d39.
//
// Solidity: function updateFile(uint256 delIndex, string newName, uint256 newFileType, address newUG, uint256 newFileSize, bool newIsRec, address newsubDic, uint256 newFileBlockHash) returns(bool res)
func (_Dictionary *DictionaryTransactor) UpdateFile(opts *bind.TransactOpts, delIndex *big.Int, newName string, newFileType *big.Int, newUG common.Address, newFileSize *big.Int, newIsRec bool, newsubDic common.Address, newFileBlockHash *big.Int) (*types.Transaction, error) {
	return _Dictionary.contract.Transact(opts, "updateFile", delIndex, newName, newFileType, newUG, newFileSize, newIsRec, newsubDic, newFileBlockHash)
}

// UpdateFile is a paid mutator transaction binding the contract method 0xd98f6d39.
//
// Solidity: function updateFile(uint256 delIndex, string newName, uint256 newFileType, address newUG, uint256 newFileSize, bool newIsRec, address newsubDic, uint256 newFileBlockHash) returns(bool res)
func (_Dictionary *DictionarySession) UpdateFile(delIndex *big.Int, newName string, newFileType *big.Int, newUG common.Address, newFileSize *big.Int, newIsRec bool, newsubDic common.Address, newFileBlockHash *big.Int) (*types.Transaction, error) {
	return _Dictionary.Contract.UpdateFile(&_Dictionary.TransactOpts, delIndex, newName, newFileType, newUG, newFileSize, newIsRec, newsubDic, newFileBlockHash)
}

// UpdateFile is a paid mutator transaction binding the contract method 0xd98f6d39.
//
// Solidity: function updateFile(uint256 delIndex, string newName, uint256 newFileType, address newUG, uint256 newFileSize, bool newIsRec, address newsubDic, uint256 newFileBlockHash) returns(bool res)
func (_Dictionary *DictionaryTransactorSession) UpdateFile(delIndex *big.Int, newName string, newFileType *big.Int, newUG common.Address, newFileSize *big.Int, newIsRec bool, newsubDic common.Address, newFileBlockHash *big.Int) (*types.Transaction, error) {
	return _Dictionary.Contract.UpdateFile(&_Dictionary.TransactOpts, delIndex, newName, newFileType, newUG, newFileSize, newIsRec, newsubDic, newFileBlockHash)
}

// RouterABI is the input ABI used to generate the binding from.
const RouterABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"show\",\"outputs\":[{\"name\":\"dic\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"dicCreator\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dic\",\"type\":\"address\"}],\"name\":\"insert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"insertAuto\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// RouterBin is the compiled bytecode used for deploying new contracts.
const RouterBin = `0x608060405234801561001057600080fd5b50610e06806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063287a61e414610051578063849263d914610093578063bc902ad21461009b578063cc39c1fa146100c3575b600080fd5b6100776004803603602081101561006757600080fd5b50356001600160a01b03166100cb565b604080516001600160a01b039092168252519081900360200190f35b6100776100e9565b6100c1600480360360208110156100b157600080fd5b50356001600160a01b031661012b565b005b6100c161015a565b6001600160a01b039081166000908152602081905260409020541690565b6000336040516100f890610162565b6001600160a01b03909116815260405190819003602001906000f080158015610125573d6000803e3d6000fd5b50905090565b33600090815260208190526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b61012b6100e9565b610c62806101708339019056fe608060405260006002556003805460ff1916905534801561001f57600080fd5b50604051610c62380380610c628339818101604052602081101561004257600080fd5b5051600080546001600160a01b039092166001600160a01b0319909216919091179055610bee806100746000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806302d05d3f1461006757806392a01a021461008b5780639926b03f1461016f578063d98f6d391461018c578063f13ab19614610264578063f89488db1461030a575b600080fd5b61006f610332565b604080516001600160a01b039092168252519081900360200190f35b61015b600480360360e08110156100a157600080fd5b8101906020810181356401000000008111156100bc57600080fd5b8201836020820111156100ce57600080fd5b803590602001918460018302840111640100000000831117156100f057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350506001600160a01b0360208301358116926040810135925060608101351515916080820135169060a00135610341565b604080519115158252519081900360200190f35b61015b6004803603602081101561018557600080fd5b50356104fd565b61015b60048036036101008110156101a357600080fd5b813591908101906040810160208201356401000000008111156101c557600080fd5b8201836020820111156101d757600080fd5b803590602001918460018302840111640100000000831117156101f957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350506001600160a01b0360208301358116926040810135925060608101351515916080820135169060a001356105de565b61015b6004803603602081101561027a57600080fd5b81019060208101813564010000000081111561029557600080fd5b8201836020820111156102a757600080fd5b803590602001918460018302840111640100000000831117156102c957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610789945050505050565b6103306004803603602081101561032057600080fd5b50356001600160a01b03166107c6565b005b6000546001600160a01b031681565b600080600061034f33610984565b9092509050600181151514156104eb57610367610a32565b6040518061012001604052808c81526020018b8152602001428152602001428152602001881515600115151461039d578a6103aa565b6007546001600160a01b03165b6001600160a01b03908116825260208083018c90528a15156040840152908916606083015260809091018790529082526000828201819052600680546001810180835591909252835180518051929486946009027ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f01938492610431928492910190610a53565b50602082810151600183810191909155604084015160028401556060840151600384015560808401516004840180546001600160a01b0319166001600160a01b0392831617905560a0850151600585015560c085015160068501805460e088015160ff1991821693151593909317610100600160a81b031916610100939094168302939093179055909401516007909301929092559390930151600892909201805490911691151591909117905594506104f29350505050565b6000925050505b979650505050505050565b600080600061050b33610984565b9092509050600181151514156105d2576006848154811061052857fe5b600091825260208220600990910201906105428282610ad1565b506000600182810182905560028301829055600383018290556004830180546001600160a01b031916905560058301829055600680840180546001600160a81b031916905560079093019190915581549091908690811061059f57fe5b906000526020600020906009020160080160006101000a81548160ff0219169083151502179055506001925050506105d9565b6000925050505b919050565b60008060006105ec33610984565b9150915060068b815481106105fd57fe5b600091825260209091206008600990920201015460ff161515600114156106295760009250505061077d565b60018115151415610776576040518061012001604052808b81526020018a8152602001428152602001428152602001871515600115151461066a5789610677565b6007546001600160a01b03165b6001600160a01b031681526020018881526020018715158152602001866001600160a01b031681526020018581525060068c815481106106b357fe5b6000918252602091829020835180516009909302909101926106da92849290910190610a53565b506020820151600182810191909155604083015160028301556060830151600383015560808301516004830180546001600160a01b0319166001600160a01b0392831617905560a0840151600584015560c084015160068401805460e087015160ff1990911692151592909217610100600160a81b0319166101009290931682029290921790915590920151600790910155925061077d915050565b6000925050505b98975050505050505050565b600080600061079733610984565b9092509050600181151514156105d25783516107ba906001906020870190610a53565b506001925050506105d9565b6000546001600160a01b0316331461080f5760405162461bcd60e51b8152600401808060200182810382526024815260200180610b966024913960400191505060405180910390fd5b6003805460ff19166001179055600780546001600160a01b0383166001600160a01b031990911617905560005b600654811015610980576006818154811061085357fe5b600091825260209091206006600990920201015460ff1615156001141561097857600754600680546001600160a01b03909216918390811061089157fe5b906000526020600020906009020160000160040160006101000a8154816001600160a01b0302191690836001600160a01b03160217905550600681815481106108d657fe5b90600052602060002090600902016000016001015460001415610978576006818154811061090057fe5b600091825260208220600660099092020101546040805163f89488db60e01b81526001600160a01b0386811660048301529151610100909304919091169263f89488db9260248084019382900301818387803b15801561095f57600080fd5b505af1158015610973573d6000803e3d6000fd5b505050505b60010161083c565b5050565b600354600090819060ff1661099e57506001905080610a2d565b6007546040805163878f70c960e01b81526001600160a01b03868116600483015291516000938493849391169163878f70c991602480820192606092909190829003018186803b1580156109f157600080fd5b505afa158015610a05573d6000803e3d6000fd5b505050506040513d6060811015610a1b57600080fd5b50805160209091015190955093505050505b915091565b604051806101400160405280610a46610b18565b8152600060209091015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a9457805160ff1916838001178555610ac1565b82800160010185558215610ac1579182015b82811115610ac1578251825591602001919060010190610aa6565b50610acd929150610b78565b5090565b50805460018160011615610100020316600290046000825580601f10610af75750610b15565b601f016020900490600052602060002090810190610b159190610b78565b50565b6040518061012001604052806060815260200160008152602001600081526020016000815260200160006001600160a01b031681526020016000815260200160001515815260200160006001600160a01b03168152602001600081525090565b610b9291905b80821115610acd5760008155600101610b7e565b9056fe4f6e6c792063726561746f722063616e2063616c6c20746869732066756e6374696f6e2ea265627a7a72305820a1dd2c6e5fde4f3706fbcc4b644af5f007704ae68a04c595efea096e321b12c164736f6c634300050a0032a265627a7a7230582086f0a8466f57ea8984437e103f79f3c1f45645919dc75a05b17b3f80607c77b464736f6c634300050a0032`

// DeployRouter deploys a new Ethereum contract, binding an instance of Router to it.
func DeployRouter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Router, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RouterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RouterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// Show is a free data retrieval call binding the contract method 0x287a61e4.
//
// Solidity: function show(address sender) constant returns(address dic)
func (_Router *RouterCaller) Show(opts *bind.CallOpts, sender common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Router.contract.Call(opts, out, "show", sender)
	return *ret0, err
}

// Show is a free data retrieval call binding the contract method 0x287a61e4.
//
// Solidity: function show(address sender) constant returns(address dic)
func (_Router *RouterSession) Show(sender common.Address) (common.Address, error) {
	return _Router.Contract.Show(&_Router.CallOpts, sender)
}

// Show is a free data retrieval call binding the contract method 0x287a61e4.
//
// Solidity: function show(address sender) constant returns(address dic)
func (_Router *RouterCallerSession) Show(sender common.Address) (common.Address, error) {
	return _Router.Contract.Show(&_Router.CallOpts, sender)
}

// DicCreator is a paid mutator transaction binding the contract method 0x849263d9.
//
// Solidity: function dicCreator() returns(address addr)
func (_Router *RouterTransactor) DicCreator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "dicCreator")
}

// DicCreator is a paid mutator transaction binding the contract method 0x849263d9.
//
// Solidity: function dicCreator() returns(address addr)
func (_Router *RouterSession) DicCreator() (*types.Transaction, error) {
	return _Router.Contract.DicCreator(&_Router.TransactOpts)
}

// DicCreator is a paid mutator transaction binding the contract method 0x849263d9.
//
// Solidity: function dicCreator() returns(address addr)
func (_Router *RouterTransactorSession) DicCreator() (*types.Transaction, error) {
	return _Router.Contract.DicCreator(&_Router.TransactOpts)
}

// Insert is a paid mutator transaction binding the contract method 0xbc902ad2.
//
// Solidity: function insert(address dic) returns()
func (_Router *RouterTransactor) Insert(opts *bind.TransactOpts, dic common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "insert", dic)
}

// Insert is a paid mutator transaction binding the contract method 0xbc902ad2.
//
// Solidity: function insert(address dic) returns()
func (_Router *RouterSession) Insert(dic common.Address) (*types.Transaction, error) {
	return _Router.Contract.Insert(&_Router.TransactOpts, dic)
}

// Insert is a paid mutator transaction binding the contract method 0xbc902ad2.
//
// Solidity: function insert(address dic) returns()
func (_Router *RouterTransactorSession) Insert(dic common.Address) (*types.Transaction, error) {
	return _Router.Contract.Insert(&_Router.TransactOpts, dic)
}

// InsertAuto is a paid mutator transaction binding the contract method 0xcc39c1fa.
//
// Solidity: function insertAuto() returns()
func (_Router *RouterTransactor) InsertAuto(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "insertAuto")
}

// InsertAuto is a paid mutator transaction binding the contract method 0xcc39c1fa.
//
// Solidity: function insertAuto() returns()
func (_Router *RouterSession) InsertAuto() (*types.Transaction, error) {
	return _Router.Contract.InsertAuto(&_Router.TransactOpts)
}

// InsertAuto is a paid mutator transaction binding the contract method 0xcc39c1fa.
//
// Solidity: function insertAuto() returns()
func (_Router *RouterTransactorSession) InsertAuto() (*types.Transaction, error) {
	return _Router.Contract.InsertAuto(&_Router.TransactOpts)
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

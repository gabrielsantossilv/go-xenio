// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package xenio

import (
	"math/big"
	"strings"
	"github.com/xenioplatform/go-xenio/common"
	"github.com/xenioplatform/go-xenio/core/types"
	"github.com/xenioplatform/go-xenio/accounts/abi"
	"github.com/xenioplatform/go-xenio/accounts/abi/bind"
)

// XNOUsersABI is the input ABI used to generate the binding from.
const XNOUsersABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_isServer\",\"type\":\"bool\"},{\"name\":\"_country\",\"type\":\"bytes32\"},{\"name\":\"_state\",\"type\":\"bytes32\"}],\"name\":\"registerNewUser\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_imageURL\",\"type\":\"string\"},{\"name\":\"_SHA256notaryHash\",\"type\":\"bytes32\"}],\"name\":\"addLogo\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllGamersAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"userAddress\",\"type\":\"address\"}],\"name\":\"getGameImages\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllUsersAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllServersAddresses\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"SHA256notaryHash\",\"type\":\"bytes32\"}],\"name\":\"getImage\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllImages\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"gameAddress\",\"type\":\"address\"}],\"name\":\"removeUser\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"imgHash\",\"type\":\"bytes32\"}],\"name\":\"removeImage\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllGamers\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bool[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllUsers\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bool[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllServers\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bool[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"},{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]"

// XNOUsersBin is the compiled bytecode used for deploying new contracts.
const XNOUsersBin = `0x606060405260008054600160a060020a033316600160a060020a0319909116179055611375806100306000396000f300606060405236156100c25763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630ed32a9881146100cf5780632594677514610104578063322d3b011461015757806334c16694146101bd5780633a5aeec3146101ee57806355622f13146102015780636ced1ae9146102145780637b7284fc146102a957806398575188146102bc5780639f6c9358146102db578063af600d19146102f1578063e2842d79146104b1578063eb043172146104c4575b34156100cd57600080fd5b005b34156100da57600080fd5b6100f060043560243515156044356064356104d7565b604051901515815260200160405180910390f35b341561010f57600080fd5b6100f060046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650509335935061056a92505050565b341561016257600080fd5b61016a610629565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156101a9578082015183820152602001610191565b505050509050019250505060405180910390f35b34156101c857600080fd5b6101dc600160a060020a0360043516610747565b60405190815260200160405180910390f35b34156101f957600080fd5b61016a610765565b341561020c57600080fd5b61016a6107ce565b341561021f57600080fd5b61022a60043561088a565b6040518080602001838152602001828103825284818151815260200191508051906020019080838360005b8381101561026d578082015183820152602001610255565b50505050905090810190601f16801561029a5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b34156102b457600080fd5b61016a610957565b34156102c757600080fd5b6100f0600160a060020a03600435166109b5565b34156102e657600080fd5b6100f0600435610a18565b34156102fc57600080fd5b610304610a5c565b6040518080602001806020018060200180602001806020018060200187810387528d818151815260200191508051906020019060200280838360005b83811015610358578082015183820152602001610340565b5050505090500187810386528c818151815260200191508051906020019060200280838360005b8381101561039757808201518382015260200161037f565b5050505090500187810385528b818151815260200191508051906020019060200280838360005b838110156103d65780820151838201526020016103be565b5050505090500187810384528a818151815260200191508051906020019060200280838360005b838110156104155780820151838201526020016103fd565b50505050905001878103835289818151815260200191508051906020019060200280838360005b8381101561045457808201518382015260200161043c565b50505050905001878103825288818151815260200191508051906020019060200280838360005b8381101561049357808201518382015260200161047b565b505050509050019c5050505050505050505050505060405180910390f35b34156104bc57600080fd5b610304610cfa565b34156104cf57600080fd5b610304610f73565b60006104e16111fa565b858152841515602080830191909152604080830186905260608301859052600160a060020a033316600090815260019092529020819081518155602082015160018201805460ff191691151591909117905560408201516002820155606082015160038201556080820151600482015560a0820151600590910155506001915050949350505050565b60003383511561057d5760009150610622565b60008381526003602052604090205460026101006001831615026000190190911604156105ad5760009150610622565b60048054600181016105bf838261122f565b50600091825260208083209190910185905584825260039052604090208480516105ed929160200190611258565b50600083815260036020908152604080832042600191820155600160a060020a03851684529182905290912060040184905591505b5092915050565b6106316112d6565b6000805b6002548110156106e6576001600060028381548110151561065257fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206001015460ff1615156106de578154829060018101610695838261122f565b916000526020600020900160006002848154811015156106b157fe5b60009182526020909120015482546101009290920a600160a060020a039182168102910219909116179055505b600101610635565b8180548060200260200160405190810160405280929190818152602001828054801561073b57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161071d575b50505050509250505090565b600160a060020a031660009081526001602052604090206004015490565b61076d6112d6565b60028054806020026020016040519081016040528092919081815260200182805480156107c357602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116107a5575b505050505090505b90565b6107d66112d6565b6000805b6002548110156106e657600160006002838154811015156107f757fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190206001015460ff1615610882578154829060018101610839838261122f565b9160005260206000209001600060028481548110151561085557fe5b60009182526020909120015482546101009290920a600160a060020a039182168102910219909116179055505b6001016107da565b6108926112d6565b600082815260036020908152604080832060018082015482549294909385936002610100948216159490940260001901169290920491601f83018290048202909101905190810160405280929190818152602001828054600181600116156101000203166002900480156109475780601f1061091c57610100808354040283529160200191610947565b820191906000526020600020905b81548152906001019060200180831161092a57829003601f168201915b5050505050915091509150915091565b61095f6112d6565b60048054806020026020016040519081016040528092919081815260200182805480156107c357602002820191906000526020600020905b81548152600190910190602001808311610997575050505050905090565b6000805433600160a060020a039081169116146109d157600080fd5b50600160a060020a031660009081526001602081905260408220828155808201805460ff191690556002810183905560038101839055600481018390556005019190915590565b6000805433600160a060020a03908116911614610a3457600080fd5b600082815260036020526040812090610a4d82826112e8565b50600060019182015592915050565b610a646112d6565b610a6c6112d6565b610a746112d6565b610a7c6112d6565b610a846112d6565b610a8c6112d6565b6000610a966112d6565b610a9e6112d6565b610aa66112d6565b610aae6112d6565b610ab66112d6565b610abe6112d6565b6000610ac86111fa565b600254985088604051805910610adb5750595b9080825280602002602001820160405250975088604051805910610afc5750595b9080825280602002602001820160405250965088604051805910610b1d5750595b9080825280602002602001820160405250955088604051805910610b3e5750595b9080825280602002602001820160405250945088604051805910610b5f5750595b9080825280602002602001820160405250935088604051805910610b805750595b90808252806020026020018201604052509250600091505b600254821015610ce45760016000600284815481101515610bb557fe5b6000918252602080832090910154600160a060020a0316835282019290925260409081019091209060c09051908101604090815282548252600183015460ff16151560208301908152600284015491830191909152600383015460608301526004830154608083015260059092015460a08201529150511515610cd9578051888381518110610c4057fe5b6020908102909101810191909152810151878381518110610c5d57fe5b9115156020928302909101909101526040810151868381518110610c7d57fe5b602090810290910101526060810151858381518110610c9857fe5b602090810290910101526080810151848381518110610cb357fe5b6020908102909101015260a0810151838381518110610cce57fe5b602090810290910101525b600190910190610b98565b50959d949c50929a509098509650945092505050565b610d026112d6565b610d0a6112d6565b610d126112d6565b610d1a6112d6565b610d226112d6565b610d2a6112d6565b6000610d346112d6565b610d3c6112d6565b610d446112d6565b610d4c6112d6565b610d546112d6565b610d5c6112d6565b6000610d666111fa565b600254985088604051805910610d795750595b9080825280602002602001820160405250975088604051805910610d9a5750595b9080825280602002602001820160405250965088604051805910610dbb5750595b9080825280602002602001820160405250955088604051805910610ddc5750595b9080825280602002602001820160405250945088604051805910610dfd5750595b9080825280602002602001820160405250935088604051805910610e1e5750595b90808252806020026020018201604052509250600091505b88821015610ce45760016000600284815481101515610e5157fe5b6000918252602080832090910154600160a060020a0316835282019290925260409081019091209060c09051908101604090815282548252600183015460ff1615156020830152600283015490820152600382015460608201526004820154608082015260059091015460a082015290508051888381518110610ed057fe5b6020908102909101810191909152810151878381518110610eed57fe5b9115156020928302909101909101526040810151868381518110610f0d57fe5b602090810290910101526060810151858381518110610f2857fe5b602090810290910101526080810151848381518110610f4357fe5b6020908102909101015260a0810151838381518110610f5e57fe5b60209081029091010152600190910190610e36565b610f7b6112d6565b610f836112d6565b610f8b6112d6565b610f936112d6565b610f9b6112d6565b610fa36112d6565b6000610fad6112d6565b610fb56112d6565b610fbd6112d6565b610fc56112d6565b610fcd6112d6565b610fd56112d6565b6000610fdf6111fa565b600254985088604051805910610ff25750595b90808252806020026020018201604052509750886040518059106110135750595b90808252806020026020018201604052509650886040518059106110345750595b90808252806020026020018201604052509550886040518059106110555750595b90808252806020026020018201604052509450886040518059106110765750595b90808252806020026020018201604052509350886040518059106110975750595b90808252806020026020018201604052509250600091505b600254821015610ce457600160006002848154811015156110cc57fe5b6000918252602080832090910154600160a060020a0316835282019290925260409081019091209060c09051908101604090815282548252600183015460ff16151560208301908152600284015491830191909152600383015460608301526004830154608083015260059092015460a0820152915051156111ef57805188838151811061115657fe5b602090810290910181019190915281015187838151811061117357fe5b911515602092830290910190910152604081015186838151811061119357fe5b6020908102909101015260608101518583815181106111ae57fe5b6020908102909101015260808101518483815181106111c957fe5b6020908102909101015260a08101518383815181106111e457fe5b602090810290910101525b6001909101906110af565b60c06040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a082015290565b8154818355818115116112535760008381526020902061125391810190830161132f565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061129957805160ff19168380011785556112c6565b828001600101855582156112c6579182015b828111156112c65782518255916020019190600101906112ab565b506112d292915061132f565b5090565b60206040519081016040526000815290565b50805460018160011615610100020316600290046000825580601f1061130e575061132c565b601f01602090049060005260206000209081019061132c919061132f565b50565b6107cb91905b808211156112d257600081556001016113355600a165627a7a723058201f9420a61bf9d68a653d812e550a628bc692f5d73d4f61b5aa4be12b135172930029`

// DeployXNOUsers deploys a new Ethereum contract, binding an instance of XNOUsers to it.
func DeployXNOUsers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *XNOUsers, error) {
	parsed, err := abi.JSON(strings.NewReader(XNOUsersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(XNOUsersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &XNOUsers{XNOUsersCaller: XNOUsersCaller{contract: contract}, XNOUsersTransactor: XNOUsersTransactor{contract: contract}}, nil
}

// XNOUsers is an auto generated Go binding around an Ethereum contract.
type XNOUsers struct {
	XNOUsersCaller     // Read-only binding to the contract
	XNOUsersTransactor // Write-only binding to the contract
}

// XNOUsersCaller is an auto generated read-only Go binding around an Ethereum contract.
type XNOUsersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XNOUsersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type XNOUsersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// XNOUsersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type XNOUsersSession struct {
	Contract     *XNOUsers         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// XNOUsersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type XNOUsersCallerSession struct {
	Contract *XNOUsersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// XNOUsersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type XNOUsersTransactorSession struct {
	Contract     *XNOUsersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// XNOUsersRaw is an auto generated low-level Go binding around an Ethereum contract.
type XNOUsersRaw struct {
	Contract *XNOUsers // Generic contract binding to access the raw methods on
}

// XNOUsersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type XNOUsersCallerRaw struct {
	Contract *XNOUsersCaller // Generic read-only contract binding to access the raw methods on
}

// XNOUsersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type XNOUsersTransactorRaw struct {
	Contract *XNOUsersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewXNOUsers creates a new instance of XNOUsers, bound to a specific deployed contract.
func NewXNOUsers(address common.Address, backend bind.ContractBackend) (*XNOUsers, error) {
	contract, err := bindXNOUsers(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &XNOUsers{XNOUsersCaller: XNOUsersCaller{contract: contract}, XNOUsersTransactor: XNOUsersTransactor{contract: contract}}, nil
}

// NewXNOUsersCaller creates a new read-only instance of XNOUsers, bound to a specific deployed contract.
func NewXNOUsersCaller(address common.Address, caller bind.ContractCaller) (*XNOUsersCaller, error) {
	contract, err := bindXNOUsers(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &XNOUsersCaller{contract: contract}, nil
}

// NewXNOUsersTransactor creates a new write-only instance of XNOUsers, bound to a specific deployed contract.
func NewXNOUsersTransactor(address common.Address, transactor bind.ContractTransactor) (*XNOUsersTransactor, error) {
	contract, err := bindXNOUsers(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &XNOUsersTransactor{contract: contract}, nil
}

// bindXNOUsers binds a generic wrapper to an already deployed contract.
func bindXNOUsers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(XNOUsersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XNOUsers *XNOUsersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XNOUsers.Contract.XNOUsersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XNOUsers *XNOUsersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XNOUsers.Contract.XNOUsersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XNOUsers *XNOUsersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XNOUsers.Contract.XNOUsersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_XNOUsers *XNOUsersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _XNOUsers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_XNOUsers *XNOUsersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _XNOUsers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_XNOUsers *XNOUsersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _XNOUsers.Contract.contract.Transact(opts, method, params...)
}

// GetAllGamers is a free data retrieval call binding the contract method 0xaf600d19.
//
// Solidity: function getAllGamers() constant returns(bytes32[], bool[], bytes32[], bytes32[], bytes32[], bytes32[])
func (_XNOUsers *XNOUsersCaller) GetAllGamers(opts *bind.CallOpts) ([][32]byte, []bool, [][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
	var (
		ret0 = new([][32]byte)
		ret1 = new([]bool)
		ret2 = new([][32]byte)
		ret3 = new([][32]byte)
		ret4 = new([][32]byte)
		ret5 = new([][32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _XNOUsers.contract.Call(opts, out, "getAllGamers")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetAllGamers is a free data retrieval call binding the contract method 0xaf600d19.
//
// Solidity: function getAllGamers() constant returns(bytes32[], bool[], bytes32[], bytes32[], bytes32[], bytes32[])
func (_XNOUsers *XNOUsersSession) GetAllGamers() ([][32]byte, []bool, [][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
	return _XNOUsers.Contract.GetAllGamers(&_XNOUsers.CallOpts)
}

// GetAllGamers is a free data retrieval call binding the contract method 0xaf600d19.
//
// Solidity: function getAllGamers() constant returns(bytes32[], bool[], bytes32[], bytes32[], bytes32[], bytes32[])
func (_XNOUsers *XNOUsersCallerSession) GetAllGamers() ([][32]byte, []bool, [][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
	return _XNOUsers.Contract.GetAllGamers(&_XNOUsers.CallOpts)
}

// GetAllGamersAddresses is a free data retrieval call binding the contract method 0x322d3b01.
//
// Solidity: function getAllGamersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersCaller) GetAllGamersAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _XNOUsers.contract.Call(opts, out, "getAllGamersAddresses")
	return *ret0, err
}

// GetAllGamersAddresses is a free data retrieval call binding the contract method 0x322d3b01.
//
// Solidity: function getAllGamersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersSession) GetAllGamersAddresses() ([]common.Address, error) {
	return _XNOUsers.Contract.GetAllGamersAddresses(&_XNOUsers.CallOpts)
}

// GetAllGamersAddresses is a free data retrieval call binding the contract method 0x322d3b01.
//
// Solidity: function getAllGamersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersCallerSession) GetAllGamersAddresses() ([]common.Address, error) {
	return _XNOUsers.Contract.GetAllGamersAddresses(&_XNOUsers.CallOpts)
}

// GetAllImages is a free data retrieval call binding the contract method 0x7b7284fc.
//
// Solidity: function getAllImages() constant returns(bytes32[])
func (_XNOUsers *XNOUsersCaller) GetAllImages(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _XNOUsers.contract.Call(opts, out, "getAllImages")
	return *ret0, err
}

// GetAllImages is a free data retrieval call binding the contract method 0x7b7284fc.
//
// Solidity: function getAllImages() constant returns(bytes32[])
func (_XNOUsers *XNOUsersSession) GetAllImages() ([][32]byte, error) {
	return _XNOUsers.Contract.GetAllImages(&_XNOUsers.CallOpts)
}

// GetAllImages is a free data retrieval call binding the contract method 0x7b7284fc.
//
// Solidity: function getAllImages() constant returns(bytes32[])
func (_XNOUsers *XNOUsersCallerSession) GetAllImages() ([][32]byte, error) {
	return _XNOUsers.Contract.GetAllImages(&_XNOUsers.CallOpts)
}

// GetAllServers is a free data retrieval call binding the contract method 0xeb043172.
//
// Solidity: function getAllServers() constant returns(bytes32[], bool[], bytes32[], bytes32[], bytes32[], bytes32[])
func (_XNOUsers *XNOUsersCaller) GetAllServers(opts *bind.CallOpts) ([][32]byte, []bool, [][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
	var (
		ret0 = new([][32]byte)
		ret1 = new([]bool)
		ret2 = new([][32]byte)
		ret3 = new([][32]byte)
		ret4 = new([][32]byte)
		ret5 = new([][32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _XNOUsers.contract.Call(opts, out, "getAllServers")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetAllServers is a free data retrieval call binding the contract method 0xeb043172.
//
// Solidity: function getAllServers() constant returns(bytes32[], bool[], bytes32[], bytes32[], bytes32[], bytes32[])
func (_XNOUsers *XNOUsersSession) GetAllServers() ([][32]byte, []bool, [][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
	return _XNOUsers.Contract.GetAllServers(&_XNOUsers.CallOpts)
}

// GetAllServers is a free data retrieval call binding the contract method 0xeb043172.
//
// Solidity: function getAllServers() constant returns(bytes32[], bool[], bytes32[], bytes32[], bytes32[], bytes32[])
func (_XNOUsers *XNOUsersCallerSession) GetAllServers() ([][32]byte, []bool, [][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
	return _XNOUsers.Contract.GetAllServers(&_XNOUsers.CallOpts)
}

// GetAllServersAddresses is a free data retrieval call binding the contract method 0x55622f13.
//
// Solidity: function getAllServersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersCaller) GetAllServersAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _XNOUsers.contract.Call(opts, out, "getAllServersAddresses")
	return *ret0, err
}

// GetAllServersAddresses is a free data retrieval call binding the contract method 0x55622f13.
//
// Solidity: function getAllServersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersSession) GetAllServersAddresses() ([]common.Address, error) {
	return _XNOUsers.Contract.GetAllServersAddresses(&_XNOUsers.CallOpts)
}

// GetAllServersAddresses is a free data retrieval call binding the contract method 0x55622f13.
//
// Solidity: function getAllServersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersCallerSession) GetAllServersAddresses() ([]common.Address, error) {
	return _XNOUsers.Contract.GetAllServersAddresses(&_XNOUsers.CallOpts)
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() constant returns(bytes32[], bool[], bytes32[], bytes32[], bytes32[], bytes32[])
func (_XNOUsers *XNOUsersCaller) GetAllUsers(opts *bind.CallOpts) ([][32]byte, []bool, [][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
	var (
		ret0 = new([][32]byte)
		ret1 = new([]bool)
		ret2 = new([][32]byte)
		ret3 = new([][32]byte)
		ret4 = new([][32]byte)
		ret5 = new([][32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _XNOUsers.contract.Call(opts, out, "getAllUsers")
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() constant returns(bytes32[], bool[], bytes32[], bytes32[], bytes32[], bytes32[])
func (_XNOUsers *XNOUsersSession) GetAllUsers() ([][32]byte, []bool, [][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
	return _XNOUsers.Contract.GetAllUsers(&_XNOUsers.CallOpts)
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() constant returns(bytes32[], bool[], bytes32[], bytes32[], bytes32[], bytes32[])
func (_XNOUsers *XNOUsersCallerSession) GetAllUsers() ([][32]byte, []bool, [][32]byte, [][32]byte, [][32]byte, [][32]byte, error) {
	return _XNOUsers.Contract.GetAllUsers(&_XNOUsers.CallOpts)
}

// GetAllUsersAddresses is a free data retrieval call binding the contract method 0x3a5aeec3.
//
// Solidity: function getAllUsersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersCaller) GetAllUsersAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _XNOUsers.contract.Call(opts, out, "getAllUsersAddresses")
	return *ret0, err
}

// GetAllUsersAddresses is a free data retrieval call binding the contract method 0x3a5aeec3.
//
// Solidity: function getAllUsersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersSession) GetAllUsersAddresses() ([]common.Address, error) {
	return _XNOUsers.Contract.GetAllUsersAddresses(&_XNOUsers.CallOpts)
}

// GetAllUsersAddresses is a free data retrieval call binding the contract method 0x3a5aeec3.
//
// Solidity: function getAllUsersAddresses() constant returns(address[])
func (_XNOUsers *XNOUsersCallerSession) GetAllUsersAddresses() ([]common.Address, error) {
	return _XNOUsers.Contract.GetAllUsersAddresses(&_XNOUsers.CallOpts)
}

// GetGameImages is a free data retrieval call binding the contract method 0x34c16694.
//
// Solidity: function getGameImages(userAddress address) constant returns(bytes32)
func (_XNOUsers *XNOUsersCaller) GetGameImages(opts *bind.CallOpts, userAddress common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _XNOUsers.contract.Call(opts, out, "getGameImages", userAddress)
	return *ret0, err
}

// GetGameImages is a free data retrieval call binding the contract method 0x34c16694.
//
// Solidity: function getGameImages(userAddress address) constant returns(bytes32)
func (_XNOUsers *XNOUsersSession) GetGameImages(userAddress common.Address) ([32]byte, error) {
	return _XNOUsers.Contract.GetGameImages(&_XNOUsers.CallOpts, userAddress)
}

// GetGameImages is a free data retrieval call binding the contract method 0x34c16694.
//
// Solidity: function getGameImages(userAddress address) constant returns(bytes32)
func (_XNOUsers *XNOUsersCallerSession) GetGameImages(userAddress common.Address) ([32]byte, error) {
	return _XNOUsers.Contract.GetGameImages(&_XNOUsers.CallOpts, userAddress)
}

// GetImage is a free data retrieval call binding the contract method 0x6ced1ae9.
//
// Solidity: function getImage(SHA256notaryHash bytes32) constant returns(string, uint256)
func (_XNOUsers *XNOUsersCaller) GetImage(opts *bind.CallOpts, SHA256notaryHash [32]byte) (string, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _XNOUsers.contract.Call(opts, out, "getImage", SHA256notaryHash)
	return *ret0, *ret1, err
}

// GetImage is a free data retrieval call binding the contract method 0x6ced1ae9.
//
// Solidity: function getImage(SHA256notaryHash bytes32) constant returns(string, uint256)
func (_XNOUsers *XNOUsersSession) GetImage(SHA256notaryHash [32]byte) (string, *big.Int, error) {
	return _XNOUsers.Contract.GetImage(&_XNOUsers.CallOpts, SHA256notaryHash)
}

// GetImage is a free data retrieval call binding the contract method 0x6ced1ae9.
//
// Solidity: function getImage(SHA256notaryHash bytes32) constant returns(string, uint256)
func (_XNOUsers *XNOUsersCallerSession) GetImage(SHA256notaryHash [32]byte) (string, *big.Int, error) {
	return _XNOUsers.Contract.GetImage(&_XNOUsers.CallOpts, SHA256notaryHash)
}

// AddLogo is a paid mutator transaction binding the contract method 0x25946775.
//
// Solidity: function addLogo(_imageURL string, _SHA256notaryHash bytes32) returns(success bool)
func (_XNOUsers *XNOUsersTransactor) AddLogo(opts *bind.TransactOpts, _imageURL string, _SHA256notaryHash [32]byte) (*types.Transaction, error) {
	return _XNOUsers.contract.Transact(opts, "addLogo", _imageURL, _SHA256notaryHash)
}

// AddLogo is a paid mutator transaction binding the contract method 0x25946775.
//
// Solidity: function addLogo(_imageURL string, _SHA256notaryHash bytes32) returns(success bool)
func (_XNOUsers *XNOUsersSession) AddLogo(_imageURL string, _SHA256notaryHash [32]byte) (*types.Transaction, error) {
	return _XNOUsers.Contract.AddLogo(&_XNOUsers.TransactOpts, _imageURL, _SHA256notaryHash)
}

// AddLogo is a paid mutator transaction binding the contract method 0x25946775.
//
// Solidity: function addLogo(_imageURL string, _SHA256notaryHash bytes32) returns(success bool)
func (_XNOUsers *XNOUsersTransactorSession) AddLogo(_imageURL string, _SHA256notaryHash [32]byte) (*types.Transaction, error) {
	return _XNOUsers.Contract.AddLogo(&_XNOUsers.TransactOpts, _imageURL, _SHA256notaryHash)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0x0ed32a98.
//
// Solidity: function registerNewUser(_id bytes32, _isServer bool, _country bytes32, _state bytes32) returns(success bool)
func (_XNOUsers *XNOUsersTransactor) RegisterNewUser(opts *bind.TransactOpts, _id [32]byte, _isServer bool, _country [32]byte, _state [32]byte) (*types.Transaction, error) {
	return _XNOUsers.contract.Transact(opts, "registerNewUser", _id, _isServer, _country, _state)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0x0ed32a98.
//
// Solidity: function registerNewUser(_id bytes32, _isServer bool, _country bytes32, _state bytes32) returns(success bool)
func (_XNOUsers *XNOUsersSession) RegisterNewUser(_id [32]byte, _isServer bool, _country [32]byte, _state [32]byte) (*types.Transaction, error) {
	return _XNOUsers.Contract.RegisterNewUser(&_XNOUsers.TransactOpts, _id, _isServer, _country, _state)
}

// RegisterNewUser is a paid mutator transaction binding the contract method 0x0ed32a98.
//
// Solidity: function registerNewUser(_id bytes32, _isServer bool, _country bytes32, _state bytes32) returns(success bool)
func (_XNOUsers *XNOUsersTransactorSession) RegisterNewUser(_id [32]byte, _isServer bool, _country [32]byte, _state [32]byte) (*types.Transaction, error) {
	return _XNOUsers.Contract.RegisterNewUser(&_XNOUsers.TransactOpts, _id, _isServer, _country, _state)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x9f6c9358.
//
// Solidity: function removeImage(imgHash bytes32) returns(success bool)
func (_XNOUsers *XNOUsersTransactor) RemoveImage(opts *bind.TransactOpts, imgHash [32]byte) (*types.Transaction, error) {
	return _XNOUsers.contract.Transact(opts, "removeImage", imgHash)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x9f6c9358.
//
// Solidity: function removeImage(imgHash bytes32) returns(success bool)
func (_XNOUsers *XNOUsersSession) RemoveImage(imgHash [32]byte) (*types.Transaction, error) {
	return _XNOUsers.Contract.RemoveImage(&_XNOUsers.TransactOpts, imgHash)
}

// RemoveImage is a paid mutator transaction binding the contract method 0x9f6c9358.
//
// Solidity: function removeImage(imgHash bytes32) returns(success bool)
func (_XNOUsers *XNOUsersTransactorSession) RemoveImage(imgHash [32]byte) (*types.Transaction, error) {
	return _XNOUsers.Contract.RemoveImage(&_XNOUsers.TransactOpts, imgHash)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(gameAddress address) returns(success bool)
func (_XNOUsers *XNOUsersTransactor) RemoveUser(opts *bind.TransactOpts, gameAddress common.Address) (*types.Transaction, error) {
	return _XNOUsers.contract.Transact(opts, "removeUser", gameAddress)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(gameAddress address) returns(success bool)
func (_XNOUsers *XNOUsersSession) RemoveUser(gameAddress common.Address) (*types.Transaction, error) {
	return _XNOUsers.Contract.RemoveUser(&_XNOUsers.TransactOpts, gameAddress)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(gameAddress address) returns(success bool)
func (_XNOUsers *XNOUsersTransactorSession) RemoveUser(gameAddress common.Address) (*types.Transaction, error) {
	return _XNOUsers.Contract.RemoveUser(&_XNOUsers.TransactOpts, gameAddress)
}
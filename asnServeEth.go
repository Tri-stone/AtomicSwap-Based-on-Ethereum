// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// AsnEthABI is the input ABI used to generate the binding from.
const AsnEthABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint64\"}],\"name\":\"getTransactionByid\",\"outputs\":[{\"name\":\"hashSecret\",\"type\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"refundStartTime\",\"type\":\"uint256\"},{\"name\":\"redeemEndTime\",\"type\":\"uint256\"},{\"name\":\"settler\",\"type\":\"address\"},{\"name\":\"secret\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint64\"},{\"name\":\"_secret\",\"type\":\"string\"}],\"name\":\"redeem\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint64\"}],\"name\":\"getHashSecretByid\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint64\"}],\"name\":\"getSecretByid\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint64\"}],\"name\":\"refund\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint64\"},{\"name\":\"_hashSecretX\",\"type\":\"bytes32\"},{\"name\":\"_receiver\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_expirationSender\",\"type\":\"uint256\"},{\"name\":\"_expirationReceiver\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"_hashSecretX\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_expirationSender\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_expirationRecerver\",\"type\":\"uint256\"}],\"name\":\"Fund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_id\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_receiver\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_secret\",\"type\":\"string\"}],\"name\":\"Redeem\",\"type\":\"event\"}]"

// AsnEth is an auto generated Go binding around an Ethereum contract.
type AsnEth struct {
	AsnEthCaller     // Read-only binding to the contract
	AsnEthTransactor // Write-only binding to the contract
	AsnEthFilterer   // Log filterer for contract events
}

// AsnEthCaller is an auto generated read-only Go binding around an Ethereum contract.
type AsnEthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AsnEthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AsnEthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AsnEthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AsnEthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AsnEthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AsnEthSession struct {
	Contract     *AsnEth           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AsnEthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AsnEthCallerSession struct {
	Contract *AsnEthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AsnEthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AsnEthTransactorSession struct {
	Contract     *AsnEthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AsnEthRaw is an auto generated low-level Go binding around an Ethereum contract.
type AsnEthRaw struct {
	Contract *AsnEth // Generic contract binding to access the raw methods on
}

// AsnEthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AsnEthCallerRaw struct {
	Contract *AsnEthCaller // Generic read-only contract binding to access the raw methods on
}

// AsnEthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AsnEthTransactorRaw struct {
	Contract *AsnEthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAsnEth creates a new instance of AsnEth, bound to a specific deployed contract.
func NewAsnEth(address common.Address, backend bind.ContractBackend) (*AsnEth, error) {
	contract, err := bindAsnEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AsnEth{AsnEthCaller: AsnEthCaller{contract: contract}, AsnEthTransactor: AsnEthTransactor{contract: contract}, AsnEthFilterer: AsnEthFilterer{contract: contract}}, nil
}

// NewAsnEthCaller creates a new read-only instance of AsnEth, bound to a specific deployed contract.
func NewAsnEthCaller(address common.Address, caller bind.ContractCaller) (*AsnEthCaller, error) {
	contract, err := bindAsnEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AsnEthCaller{contract: contract}, nil
}

// NewAsnEthTransactor creates a new write-only instance of AsnEth, bound to a specific deployed contract.
func NewAsnEthTransactor(address common.Address, transactor bind.ContractTransactor) (*AsnEthTransactor, error) {
	contract, err := bindAsnEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AsnEthTransactor{contract: contract}, nil
}

// NewAsnEthFilterer creates a new log filterer instance of AsnEth, bound to a specific deployed contract.
func NewAsnEthFilterer(address common.Address, filterer bind.ContractFilterer) (*AsnEthFilterer, error) {
	contract, err := bindAsnEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AsnEthFilterer{contract: contract}, nil
}

// bindAsnEth binds a generic wrapper to an already deployed contract.
func bindAsnEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AsnEthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AsnEth *AsnEthRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AsnEth.Contract.AsnEthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AsnEth *AsnEthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AsnEth.Contract.AsnEthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AsnEth *AsnEthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AsnEth.Contract.AsnEthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AsnEth *AsnEthCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AsnEth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AsnEth *AsnEthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AsnEth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AsnEth *AsnEthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AsnEth.Contract.contract.Transact(opts, method, params...)
}

// ContractName is a free data retrieval call binding the contract method 0x75d0c0dc.
//
// Solidity: function contractName() constant returns(string)
func (_AsnEth *AsnEthCaller) ContractName(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AsnEth.contract.Call(opts, out, "contractName")
	return *ret0, err
}

// ContractName is a free data retrieval call binding the contract method 0x75d0c0dc.
//
// Solidity: function contractName() constant returns(string)
func (_AsnEth *AsnEthSession) ContractName() (string, error) {
	return _AsnEth.Contract.ContractName(&_AsnEth.CallOpts)
}

// ContractName is a free data retrieval call binding the contract method 0x75d0c0dc.
//
// Solidity: function contractName() constant returns(string)
func (_AsnEth *AsnEthCallerSession) ContractName() (string, error) {
	return _AsnEth.Contract.ContractName(&_AsnEth.CallOpts)
}

// GetHashSecretByid is a free data retrieval call binding the contract method 0x96518f5e.
//
// Solidity: function getHashSecretByid(_id uint64) constant returns(bytes32)
func (_AsnEth *AsnEthCaller) GetHashSecretByid(opts *bind.CallOpts, _id uint64) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AsnEth.contract.Call(opts, out, "getHashSecretByid", _id)
	return *ret0, err
}

// GetHashSecretByid is a free data retrieval call binding the contract method 0x96518f5e.
//
// Solidity: function getHashSecretByid(_id uint64) constant returns(bytes32)
func (_AsnEth *AsnEthSession) GetHashSecretByid(_id uint64) ([32]byte, error) {
	return _AsnEth.Contract.GetHashSecretByid(&_AsnEth.CallOpts, _id)
}

// GetHashSecretByid is a free data retrieval call binding the contract method 0x96518f5e.
//
// Solidity: function getHashSecretByid(_id uint64) constant returns(bytes32)
func (_AsnEth *AsnEthCallerSession) GetHashSecretByid(_id uint64) ([32]byte, error) {
	return _AsnEth.Contract.GetHashSecretByid(&_AsnEth.CallOpts, _id)
}

// GetSecretByid is a free data retrieval call binding the contract method 0xc639d48a.
//
// Solidity: function getSecretByid(_id uint64) constant returns(string)
func (_AsnEth *AsnEthCaller) GetSecretByid(opts *bind.CallOpts, _id uint64) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AsnEth.contract.Call(opts, out, "getSecretByid", _id)
	return *ret0, err
}

// GetSecretByid is a free data retrieval call binding the contract method 0xc639d48a.
//
// Solidity: function getSecretByid(_id uint64) constant returns(string)
func (_AsnEth *AsnEthSession) GetSecretByid(_id uint64) (string, error) {
	return _AsnEth.Contract.GetSecretByid(&_AsnEth.CallOpts, _id)
}

// GetSecretByid is a free data retrieval call binding the contract method 0xc639d48a.
//
// Solidity: function getSecretByid(_id uint64) constant returns(string)
func (_AsnEth *AsnEthCallerSession) GetSecretByid(_id uint64) (string, error) {
	return _AsnEth.Contract.GetSecretByid(&_AsnEth.CallOpts, _id)
}

// GetTransactionByid is a free data retrieval call binding the contract method 0x750f83b3.
//
// Solidity: function getTransactionByid(_id uint64) constant returns(hashSecret bytes32, sender address, receiver address, token address, amount uint256, refundStartTime uint256, redeemEndTime uint256, settler address, secret string)
func (_AsnEth *AsnEthCaller) GetTransactionByid(opts *bind.CallOpts, _id uint64) (struct {
	HashSecret      [32]byte
	Sender          common.Address
	Receiver        common.Address
	Token           common.Address
	Amount          *big.Int
	RefundStartTime *big.Int
	RedeemEndTime   *big.Int
	Settler         common.Address
	Secret          string
}, error) {
	ret := new(struct {
		HashSecret      [32]byte
		Sender          common.Address
		Receiver        common.Address
		Token           common.Address
		Amount          *big.Int
		RefundStartTime *big.Int
		RedeemEndTime   *big.Int
		Settler         common.Address
		Secret          string
	})
	out := ret
	err := _AsnEth.contract.Call(opts, out, "getTransactionByid", _id)
	return *ret, err
}

// GetTransactionByid is a free data retrieval call binding the contract method 0x750f83b3.
//
// Solidity: function getTransactionByid(_id uint64) constant returns(hashSecret bytes32, sender address, receiver address, token address, amount uint256, refundStartTime uint256, redeemEndTime uint256, settler address, secret string)
func (_AsnEth *AsnEthSession) GetTransactionByid(_id uint64) (struct {
	HashSecret      [32]byte
	Sender          common.Address
	Receiver        common.Address
	Token           common.Address
	Amount          *big.Int
	RefundStartTime *big.Int
	RedeemEndTime   *big.Int
	Settler         common.Address
	Secret          string
}, error) {
	return _AsnEth.Contract.GetTransactionByid(&_AsnEth.CallOpts, _id)
}

// GetTransactionByid is a free data retrieval call binding the contract method 0x750f83b3.
//
// Solidity: function getTransactionByid(_id uint64) constant returns(hashSecret bytes32, sender address, receiver address, token address, amount uint256, refundStartTime uint256, redeemEndTime uint256, settler address, secret string)
func (_AsnEth *AsnEthCallerSession) GetTransactionByid(_id uint64) (struct {
	HashSecret      [32]byte
	Sender          common.Address
	Receiver        common.Address
	Token           common.Address
	Amount          *big.Int
	RefundStartTime *big.Int
	RedeemEndTime   *big.Int
	Settler         common.Address
	Secret          string
}, error) {
	return _AsnEth.Contract.GetTransactionByid(&_AsnEth.CallOpts, _id)
}

// Fund is a paid mutator transaction binding the contract method 0xdcc9843f.
//
// Solidity: function fund(_id uint64, _hashSecretX bytes32, _receiver address, _token address, _amount uint256, _expirationSender uint256, _expirationReceiver uint256) returns(bool)
func (_AsnEth *AsnEthTransactor) Fund(opts *bind.TransactOpts, _id uint64, _hashSecretX [32]byte, _receiver common.Address, _token common.Address, _amount *big.Int, _expirationSender *big.Int, _expirationReceiver *big.Int) (*types.Transaction, error) {
	return _AsnEth.contract.Transact(opts, "fund", _id, _hashSecretX, _receiver, _token, _amount, _expirationSender, _expirationReceiver)
}

// Fund is a paid mutator transaction binding the contract method 0xdcc9843f.
//
// Solidity: function fund(_id uint64, _hashSecretX bytes32, _receiver address, _token address, _amount uint256, _expirationSender uint256, _expirationReceiver uint256) returns(bool)
func (_AsnEth *AsnEthSession) Fund(_id uint64, _hashSecretX [32]byte, _receiver common.Address, _token common.Address, _amount *big.Int, _expirationSender *big.Int, _expirationReceiver *big.Int) (*types.Transaction, error) {
	return _AsnEth.Contract.Fund(&_AsnEth.TransactOpts, _id, _hashSecretX, _receiver, _token, _amount, _expirationSender, _expirationReceiver)
}

// Fund is a paid mutator transaction binding the contract method 0xdcc9843f.
//
// Solidity: function fund(_id uint64, _hashSecretX bytes32, _receiver address, _token address, _amount uint256, _expirationSender uint256, _expirationReceiver uint256) returns(bool)
func (_AsnEth *AsnEthTransactorSession) Fund(_id uint64, _hashSecretX [32]byte, _receiver common.Address, _token common.Address, _amount *big.Int, _expirationSender *big.Int, _expirationReceiver *big.Int) (*types.Transaction, error) {
	return _AsnEth.Contract.Fund(&_AsnEth.TransactOpts, _id, _hashSecretX, _receiver, _token, _amount, _expirationSender, _expirationReceiver)
}

// Redeem is a paid mutator transaction binding the contract method 0x90dd93c5.
//
// Solidity: function redeem(_id uint64, _secret string) returns(bool)
func (_AsnEth *AsnEthTransactor) Redeem(opts *bind.TransactOpts, _id uint64, _secret string) (*types.Transaction, error) {
	return _AsnEth.contract.Transact(opts, "redeem", _id, _secret)
}

// Redeem is a paid mutator transaction binding the contract method 0x90dd93c5.
//
// Solidity: function redeem(_id uint64, _secret string) returns(bool)
func (_AsnEth *AsnEthSession) Redeem(_id uint64, _secret string) (*types.Transaction, error) {
	return _AsnEth.Contract.Redeem(&_AsnEth.TransactOpts, _id, _secret)
}

// Redeem is a paid mutator transaction binding the contract method 0x90dd93c5.
//
// Solidity: function redeem(_id uint64, _secret string) returns(bool)
func (_AsnEth *AsnEthTransactorSession) Redeem(_id uint64, _secret string) (*types.Transaction, error) {
	return _AsnEth.Contract.Redeem(&_AsnEth.TransactOpts, _id, _secret)
}

// Refund is a paid mutator transaction binding the contract method 0xd7194ccb.
//
// Solidity: function refund(_id uint64) returns(bool)
func (_AsnEth *AsnEthTransactor) Refund(opts *bind.TransactOpts, _id uint64) (*types.Transaction, error) {
	return _AsnEth.contract.Transact(opts, "refund", _id)
}

// Refund is a paid mutator transaction binding the contract method 0xd7194ccb.
//
// Solidity: function refund(_id uint64) returns(bool)
func (_AsnEth *AsnEthSession) Refund(_id uint64) (*types.Transaction, error) {
	return _AsnEth.Contract.Refund(&_AsnEth.TransactOpts, _id)
}

// Refund is a paid mutator transaction binding the contract method 0xd7194ccb.
//
// Solidity: function refund(_id uint64) returns(bool)
func (_AsnEth *AsnEthTransactorSession) Refund(_id uint64) (*types.Transaction, error) {
	return _AsnEth.Contract.Refund(&_AsnEth.TransactOpts, _id)
}

// AsnEthFundIterator is returned from FilterFund and is used to iterate over the raw logs and unpacked data for Fund events raised by the AsnEth contract.
type AsnEthFundIterator struct {
	Event *AsnEthFund // Event containing the contract specifics and raw log

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
func (it *AsnEthFundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AsnEthFund)
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
		it.Event = new(AsnEthFund)
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
func (it *AsnEthFundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AsnEthFundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AsnEthFund represents a Fund event raised by the AsnEth contract.
type AsnEthFund struct {
	Id                 uint64
	HashSecretX        [32]byte
	Sender             common.Address
	Receiver           common.Address
	Token              common.Address
	Amount             *big.Int
	ExpirationSender   *big.Int
	ExpirationRecerver *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterFund is a free log retrieval operation binding the contract event 0xe9d36e435972ddda8ccc5b594a8acc9de53455f522c71e1e6a1d07467315221d.
//
// Solidity: e Fund(_id uint64, _hashSecretX bytes32, _sender address, _receiver address, _token address, _amount uint256, _expirationSender uint256, _expirationRecerver uint256)
func (_AsnEth *AsnEthFilterer) FilterFund(opts *bind.FilterOpts) (*AsnEthFundIterator, error) {

	logs, sub, err := _AsnEth.contract.FilterLogs(opts, "Fund")
	if err != nil {
		return nil, err
	}
	return &AsnEthFundIterator{contract: _AsnEth.contract, event: "Fund", logs: logs, sub: sub}, nil
}

// WatchFund is a free log subscription operation binding the contract event 0xe9d36e435972ddda8ccc5b594a8acc9de53455f522c71e1e6a1d07467315221d.
//
// Solidity: e Fund(_id uint64, _hashSecretX bytes32, _sender address, _receiver address, _token address, _amount uint256, _expirationSender uint256, _expirationRecerver uint256)
func (_AsnEth *AsnEthFilterer) WatchFund(opts *bind.WatchOpts, sink chan<- *AsnEthFund) (event.Subscription, error) {

	logs, sub, err := _AsnEth.contract.WatchLogs(opts, "Fund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AsnEthFund)
				if err := _AsnEth.contract.UnpackLog(event, "Fund", log); err != nil {
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

// AsnEthRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the AsnEth contract.
type AsnEthRedeemIterator struct {
	Event *AsnEthRedeem // Event containing the contract specifics and raw log

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
func (it *AsnEthRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AsnEthRedeem)
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
		it.Event = new(AsnEthRedeem)
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
func (it *AsnEthRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AsnEthRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AsnEthRedeem represents a Redeem event raised by the AsnEth contract.
type AsnEthRedeem struct {
	Id       uint64
	Sender   common.Address
	Receiver common.Address
	Token    common.Address
	Amount   *big.Int
	Secret   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0x15dd052f5ae162ad19865274f89f200be2a118dc46c6ae12c7d611acd8414ddd.
//
// Solidity: e Redeem(_id uint64, _sender address, _receiver address, _token address, _amount uint256, _secret string)
func (_AsnEth *AsnEthFilterer) FilterRedeem(opts *bind.FilterOpts) (*AsnEthRedeemIterator, error) {

	logs, sub, err := _AsnEth.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &AsnEthRedeemIterator{contract: _AsnEth.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0x15dd052f5ae162ad19865274f89f200be2a118dc46c6ae12c7d611acd8414ddd.
//
// Solidity: e Redeem(_id uint64, _sender address, _receiver address, _token address, _amount uint256, _secret string)
func (_AsnEth *AsnEthFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *AsnEthRedeem) (event.Subscription, error) {

	logs, sub, err := _AsnEth.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AsnEthRedeem)
				if err := _AsnEth.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// AsnEthRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the AsnEth contract.
type AsnEthRefundIterator struct {
	Event *AsnEthRefund // Event containing the contract specifics and raw log

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
func (it *AsnEthRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AsnEthRefund)
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
		it.Event = new(AsnEthRefund)
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
func (it *AsnEthRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AsnEthRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AsnEthRefund represents a Refund event raised by the AsnEth contract.
type AsnEthRefund struct {
	Id       uint64
	Sender   common.Address
	Receiver common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0x49c6d7ac84fbfe83159f726c600335499f4fec0c93380d92a3c888f61b12066b.
//
// Solidity: e Refund(_id uint64, _sender address, _receiver address, _token address, _amount uint256)
func (_AsnEth *AsnEthFilterer) FilterRefund(opts *bind.FilterOpts) (*AsnEthRefundIterator, error) {

	logs, sub, err := _AsnEth.contract.FilterLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return &AsnEthRefundIterator{contract: _AsnEth.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0x49c6d7ac84fbfe83159f726c600335499f4fec0c93380d92a3c888f61b12066b.
//
// Solidity: e Refund(_id uint64, _sender address, _receiver address, _token address, _amount uint256)
func (_AsnEth *AsnEthFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *AsnEthRefund) (event.Subscription, error) {

	logs, sub, err := _AsnEth.contract.WatchLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AsnEthRefund)
				if err := _AsnEth.contract.UnpackLog(event, "Refund", log); err != nil {
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

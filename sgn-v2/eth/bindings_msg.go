// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MsgDataTypesBridgeTransferParams is an auto generated low-level Go binding around an user-defined struct.
type MsgDataTypesBridgeTransferParams struct {
	Request []byte
	Sigs    [][]byte
	Signers []common.Address
	Powers  []*big.Int
}

// MsgDataTypesMsgWithTransferExecutionParams is an auto generated low-level Go binding around an user-defined struct.
type MsgDataTypesMsgWithTransferExecutionParams struct {
	Message  []byte
	Transfer MsgDataTypesTransferInfo
	Sigs     [][]byte
	Signers  []common.Address
	Powers   []*big.Int
}

// MsgDataTypesRouteInfo is an auto generated low-level Go binding around an user-defined struct.
type MsgDataTypesRouteInfo struct {
	Sender     common.Address
	Receiver   common.Address
	SrcChainId uint64
	SrcTxHash  [32]byte
}

// MsgDataTypesRouteInfo2 is an auto generated low-level Go binding around an user-defined struct.
type MsgDataTypesRouteInfo2 struct {
	Sender     []byte
	Receiver   common.Address
	SrcChainId uint64
	SrcTxHash  [32]byte
}

// MsgDataTypesTransferInfo is an auto generated low-level Go binding around an user-defined struct.
type MsgDataTypesTransferInfo struct {
	T          uint8
	Sender     common.Address
	Receiver   common.Address
	Token      common.Address
	Amount     *big.Int
	Wdseq      uint64
	SrcChainId uint64
	RefId      [32]byte
	SrcTxHash  [32]byte
}

// IMessageReceiverAppMetaData contains all meta data concerning the IMessageReceiverApp contract.
var IMessageReceiverAppMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_sender\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// IMessageReceiverAppABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageReceiverAppMetaData.ABI instead.
var IMessageReceiverAppABI = IMessageReceiverAppMetaData.ABI

// IMessageReceiverApp is an auto generated Go binding around an Ethereum contract.
type IMessageReceiverApp struct {
	IMessageReceiverAppCaller     // Read-only binding to the contract
	IMessageReceiverAppTransactor // Write-only binding to the contract
	IMessageReceiverAppFilterer   // Log filterer for contract events
}

// IMessageReceiverAppCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageReceiverAppCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageReceiverAppTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageReceiverAppTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageReceiverAppFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageReceiverAppFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageReceiverAppSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageReceiverAppSession struct {
	Contract     *IMessageReceiverApp // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMessageReceiverAppCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageReceiverAppCallerSession struct {
	Contract *IMessageReceiverAppCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IMessageReceiverAppTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageReceiverAppTransactorSession struct {
	Contract     *IMessageReceiverAppTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IMessageReceiverAppRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageReceiverAppRaw struct {
	Contract *IMessageReceiverApp // Generic contract binding to access the raw methods on
}

// IMessageReceiverAppCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageReceiverAppCallerRaw struct {
	Contract *IMessageReceiverAppCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageReceiverAppTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageReceiverAppTransactorRaw struct {
	Contract *IMessageReceiverAppTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageReceiverApp creates a new instance of IMessageReceiverApp, bound to a specific deployed contract.
func NewIMessageReceiverApp(address common.Address, backend bind.ContractBackend) (*IMessageReceiverApp, error) {
	contract, err := bindIMessageReceiverApp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageReceiverApp{IMessageReceiverAppCaller: IMessageReceiverAppCaller{contract: contract}, IMessageReceiverAppTransactor: IMessageReceiverAppTransactor{contract: contract}, IMessageReceiverAppFilterer: IMessageReceiverAppFilterer{contract: contract}}, nil
}

// NewIMessageReceiverAppCaller creates a new read-only instance of IMessageReceiverApp, bound to a specific deployed contract.
func NewIMessageReceiverAppCaller(address common.Address, caller bind.ContractCaller) (*IMessageReceiverAppCaller, error) {
	contract, err := bindIMessageReceiverApp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageReceiverAppCaller{contract: contract}, nil
}

// NewIMessageReceiverAppTransactor creates a new write-only instance of IMessageReceiverApp, bound to a specific deployed contract.
func NewIMessageReceiverAppTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageReceiverAppTransactor, error) {
	contract, err := bindIMessageReceiverApp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageReceiverAppTransactor{contract: contract}, nil
}

// NewIMessageReceiverAppFilterer creates a new log filterer instance of IMessageReceiverApp, bound to a specific deployed contract.
func NewIMessageReceiverAppFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageReceiverAppFilterer, error) {
	contract, err := bindIMessageReceiverApp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageReceiverAppFilterer{contract: contract}, nil
}

// bindIMessageReceiverApp binds a generic wrapper to an already deployed contract.
func bindIMessageReceiverApp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMessageReceiverAppABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageReceiverApp *IMessageReceiverAppRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageReceiverApp.Contract.IMessageReceiverAppCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageReceiverApp *IMessageReceiverAppRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.IMessageReceiverAppTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageReceiverApp *IMessageReceiverAppRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.IMessageReceiverAppTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageReceiverApp *IMessageReceiverAppCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageReceiverApp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageReceiverApp *IMessageReceiverAppTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageReceiverApp *IMessageReceiverAppTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.contract.Transact(opts, method, params...)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x063ce4e5.
//
// Solidity: function executeMessage(bytes _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactor) ExecuteMessage(opts *bind.TransactOpts, _sender []byte, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.contract.Transact(opts, "executeMessage", _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x063ce4e5.
//
// Solidity: function executeMessage(bytes _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppSession) ExecuteMessage(_sender []byte, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessage(&_IMessageReceiverApp.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x063ce4e5.
//
// Solidity: function executeMessage(bytes _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactorSession) ExecuteMessage(_sender []byte, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessage(&_IMessageReceiverApp.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactor) ExecuteMessage0(opts *bind.TransactOpts, _sender common.Address, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.contract.Transact(opts, "executeMessage0", _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppSession) ExecuteMessage0(_sender common.Address, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessage0(&_IMessageReceiverApp.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactorSession) ExecuteMessage0(_sender common.Address, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessage0(&_IMessageReceiverApp.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.contract.Transact(opts, "executeMessageWithTransfer", _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransfer(&_IMessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactorSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransfer(&_IMessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactor) ExecuteMessageWithTransferFallback(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.contract.Transact(opts, "executeMessageWithTransferFallback", _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransferFallback(&_IMessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactorSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransferFallback(&_IMessageReceiverApp.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.contract.Transact(opts, "executeMessageWithTransferRefund", _token, _amount, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransferRefund(&_IMessageReceiverApp.TransactOpts, _token, _amount, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_IMessageReceiverApp *IMessageReceiverAppTransactorSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _IMessageReceiverApp.Contract.ExecuteMessageWithTransferRefund(&_IMessageReceiverApp.TransactOpts, _token, _amount, _message, _executor)
}

// MessageBusMetaData contains all meta data concerning the MessageBus contract.
var MessageBusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"_sigsVerifier\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidityBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridgeV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVaultV2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"CallReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMsgDataTypes.MsgType\",\"name\":\"msgType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumMsgDataTypes.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBase\",\"type\":\"uint256\"}],\"name\":\"FeeBaseUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"}],\"name\":\"FeePerByteUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidityBridge\",\"type\":\"address\"}],\"name\":\"LiquidityBridgeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Message\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Message2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTransferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"MessageWithTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMsgDataTypes.MsgType\",\"name\":\"msgType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"name\":\"NeedRetry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegBridge\",\"type\":\"address\"}],\"name\":\"PegBridgeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegBridgeV2\",\"type\":\"address\"}],\"name\":\"PegBridgeV2Updated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegVault\",\"type\":\"address\"}],\"name\":\"PegVaultUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegVaultV2\",\"type\":\"address\"}],\"name\":\"PegVaultV2Updated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"calcFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.RouteInfo\",\"name\":\"_route\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.RouteInfo2\",\"name\":\"_route\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executedMessages\",\"outputs\":[{\"internalType\":\"enumMsgDataTypes.TxStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePerByte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridgeV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVaultV2\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridgeV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVaultV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preExecuteMessageGasUsage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.BridgeTransferParams\",\"name\":\"_transferParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.MsgWithTransferExecutionParams\",\"name\":\"_msgParams\",\"type\":\"tuple\"}],\"name\":\"refundAndExecuteMsg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_srcBridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_srcTransferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFeeBase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFeePerByte\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setLiquidityBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridgeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVaultV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_usage\",\"type\":\"uint256\"}],\"name\":\"setPreExecuteMessageGasUsage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigsVerifier\",\"outputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.BridgeTransferParams\",\"name\":\"_transferParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.MsgWithTransferExecutionParams\",\"name\":\"_msgParams\",\"type\":\"tuple\"}],\"name\":\"transferAndExecuteMsg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawnFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b50604051620041183803806200411883398101604081905262000034916200011d565b84848484848a6200004533620000b4565b6001600160a01b03908116608052600580546001600160a01b0319908116978316979097179055600680548716958216959095179094556007805486169385169390931790925560088054851691841691909117905560098054909316911617905550620001b1945050505050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811681146200011a57600080fd5b50565b60008060008060008060c087890312156200013757600080fd5b8651620001448162000104565b6020880151909650620001578162000104565b60408801519095506200016a8162000104565b60608801519094506200017d8162000104565b6080880151909350620001908162000104565b60a0880151909250620001a38162000104565b809150509295509295509295565b608051613f44620001d4600039600081816104eb01526108430152613f446000f3fe6080604052600436106101fe5760003560e01c806382980dc41161011d578063ccf2683b116100b0578063dfa2dbaf1161007f578063f2fde38b11610064578063f2fde38b146105bd578063f60bbe2a146105dd578063f83b0fb9146105f357600080fd5b8063dfa2dbaf1461057d578063e2c1ed251461059d57600080fd5b8063ccf2683b146104d9578063cd2abd661461050d578063d8257d171461054a578063db2c20c81461056a57600080fd5b806395e911a8116100ec57806395e911a8146104705780639b05a775146104865780639f3ce55a146104a6578063c66a9c5a146104b957600080fd5b806382980dc4146103da57806382efd502146104125780638da5cb5b1461043257806395b12c271461045057600080fd5b8063468a2d04116101955780635b3e5f50116101645780635b3e5f5014610367578063723d0a9d146103945780637b80ab20146103b45780637d7a101d146103c757600080fd5b8063468a2d04146102eb5780635335dca2146102fe578063584e45e114610331578063588be02b1461034757600080fd5b80633f395aff116101d15780633f395aff1461028557806340d0d026146102985780634289fbb3146102b85780634586f331146102cb57600080fd5b806303cbfe661461020357806306c28bd6146102255780632ff4c41114610245578063359ef75b14610265575b600080fd5b34801561020f57600080fd5b5061022361021e366004612f40565b610613565b005b34801561023157600080fd5b50610223610240366004612f5b565b61070c565b34801561025157600080fd5b50610223610260366004612fc0565b610798565b34801561027157600080fd5b50610223610280366004613074565b610a33565b61022361029336600461311b565b610a4f565b3480156102a457600080fd5b506102236102b3366004613224565b610d3f565b6102236102c6366004613290565b610d97565b3480156102d757600080fd5b506102236102e6366004612f5b565b610e7f565b6102236102f9366004613308565b610edb565b34801561030a57600080fd5b5061031e6103193660046133cd565b610f3b565b6040519081526020015b60405180910390f35b34801561033d57600080fd5b5061031e600a5481565b34801561035357600080fd5b50610223610362366004612f40565b610f61565b34801561037357600080fd5b5061031e610382366004612f40565b60036020526000908152604090205481565b3480156103a057600080fd5b506102236103af366004613224565b61104e565b6102236103c236600461311b565b61109c565b6102236103d536600461340f565b6112a9565b3480156103e657600080fd5b506005546103fa906001600160a01b031681565b6040516001600160a01b039091168152602001610328565b34801561041e57600080fd5b5061022361042d366004612f40565b611306565b34801561043e57600080fd5b506000546001600160a01b03166103fa565b34801561045c57600080fd5b506008546103fa906001600160a01b031681565b34801561047c57600080fd5b5061031e60015481565b34801561049257600080fd5b506102236104a1366004612f40565b6113f3565b6102236104b4366004613489565b6114e0565b3480156104c557600080fd5b506009546103fa906001600160a01b031681565b3480156104e557600080fd5b506103fa7f000000000000000000000000000000000000000000000000000000000000000081565b34801561051957600080fd5b5061053d610528366004612f5b565b60046020526000908152604090205460ff1681565b604051610328919061350d565b34801561055657600080fd5b506007546103fa906001600160a01b031681565b61022361057836600461351b565b61153a565b34801561058957600080fd5b506006546103fa906001600160a01b031681565b3480156105a957600080fd5b506102236105b8366004612f5b565b61158e565b3480156105c957600080fd5b506102236105d8366004612f40565b61161a565b3480156105e957600080fd5b5061031e60025481565b3480156105ff57600080fd5b5061022361060e366004612f40565b6116f9565b336106266000546001600160a01b031690565b6001600160a01b03161461066f5760405162461bcd60e51b81526020600482018190526024820152600080516020613ec383398151915260448201526064015b60405180910390fd5b6001600160a01b0381166106b75760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b6044820152606401610666565b600680546001600160a01b0319166001600160a01b0383169081179091556040519081527fd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d906020015b60405180910390a150565b3361071f6000546001600160a01b031690565b6001600160a01b0316146107635760405162461bcd60e51b81526020600482018190526024820152600080516020613ec38339815191526044820152606401610666565b60018190556040518181527f892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f290602001610701565b600046306040516020016107e992919091825260601b6001600160601b03191660208201527f77697468647261774665650000000000000000000000000000000000000000006034820152603f0190565b60408051808303601f19018152828252805160209182012090830181905260608c901b6001600160601b0319168383015260548084018c9052825180850390910181526074840192839052633416de1160e11b90925292507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163682dbc229161088a918b908b908b908b908b908b9060780161377f565b60006040518083038186803b1580156108a257600080fd5b505afa1580156108b6573d6000803e3d6000fd5b505050506001600160a01b0389166000908152600360205260408120546108dd908a6137f3565b90506000811161092f5760405162461bcd60e51b815260206004820152601960248201527f4e6f206e657720616d6f756e7420746f207769746864726177000000000000006044820152606401610666565b6001600160a01b038a166000818152600360205260408082208c90555190919061c35090849084818181858888f193505050503d806000811461098e576040519150601f19603f3d011682016040523d82523d6000602084013e610993565b606091505b50509050806109e45760405162461bcd60e51b815260206004820152601660248201527f6661696c656420746f20776974686472617720666565000000000000000000006044820152606401610666565b604080516001600160a01b038d168152602081018490527f78473f3f373f7673597f4f0fa5873cb4d375fea6d4339ad6b56dbd411513cb3f910160405180910390a15050505050505050505050565b610a3b6117e6565b610a48858585858561184a565b5050505050565b6000610a5a88611902565b90506000808281526004602081905260409091205460ff1690811115610a8257610a826134e3565b14610acf5760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c7265616479206578656375746564000000000000006044820152606401610666565b6000818152600460208181526040808420805460ff1916909317909255815146918101919091526001600160601b03193060601b16918101919091527f4d657373616765576974685472616e73666572000000000000000000000000006054820152606701604051602081830303815290604052805190602001209050600560009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e8e6101000135604051602001610b8e959493929190613806565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b8152600401610bc5979695949392919061377f565b60006040518083038186803b158015610bdd57600080fd5b505afa158015610bf1573d6000803e3d6000fd5b50505050600080610c038b8e8e61216b565b90506001816002811115610c1957610c196134e3565b03610c275760019150610cef565b6002816002811115610c3b57610c3b6134e3565b03610cbb576000848152600460205260408120805460ff19166001835b02179055507fe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c6000858d60c0016020810190610c949190613828565b8e6101000135604051610caa9493929190613862565b60405180910390a150505050610d34565b610cc68b8e8e6122a6565b90506001816002811115610cdc57610cdc6134e3565b03610cea5760039150610cef565b600291505b60008481526004602081905260409091208054849260ff19909116906001908490811115610d1f57610d1f6134e3565b0217905550610d2f84838d6122e1565b505050505b505050505050505050565b610d58610d526040830160208401613895565b83612353565b610d93610d6582806138b6565b60208401610d776101408601866138fd565b610d856101608801886138fd565b6103c26101808a018a6138fd565b5050565b468503610dd85760405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9259608a1b6044820152606401610666565b6000610de48383610f3b565b905080341015610e295760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b6044820152606401610666565b336001600160a01b03167f172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f6688888888888834604051610e6e9796959493929190613947565b60405180910390a250505050505050565b33610e926000546001600160a01b031690565b6001600160a01b031614610ed65760405162461bcd60e51b81526020600482018190526024820152600080516020613ec38339815191526044820152606401610666565b600a55565b6000610ee68861259d565b9050610f2f8a8a838a8a8a8a8a8a6040518060400160405280600781526020017f4d65737361676500000000000000000000000000000000000000000000000000815250612654565b50505050505050505050565b600254600090610f4b9083613994565b600154610f5891906139ab565b90505b92915050565b33610f746000546001600160a01b031690565b6001600160a01b031614610fb85760405162461bcd60e51b81526020600482018190526024820152600080516020613ec38339815191526044820152606401610666565b6001600160a01b0381166110005760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b6044820152606401610666565b600580546001600160a01b0319166001600160a01b0383169081179091556040519081527fbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae05890602001610701565b611061610d526040830160208401613895565b610d9361106e82806138b6565b602084016110806101408601866138fd565b61108e6101608801886138fd565b6102936101808a018a6138fd565b60006110a788611902565b90506000808281526004602081905260409091205460ff16908111156110cf576110cf6134e3565b1461111c5760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c7265616479206578656375746564000000000000006044820152606401610666565b6000818152600460208181526040808420805460ff1916909317909255815146918101919091526001600160601b03193060601b16918101919091527f4d657373616765576974685472616e73666572526566756e64000000000000006054820152606d01604051602081830303815290604052805190602001209050600560009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e8e61010001356040516020016111db959493929190613806565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b8152600401611212979695949392919061377f565b60006040518083038186803b15801561122a57600080fd5b505afa15801561123e573d6000803e3d6000fd5b505050506000806112508b8e8e6128a8565b90506001816002811115611266576112666134e3565b036112745760019150610cef565b6002816002811115611288576112886134e3565b03610cea576000848152600460205260408120805460ff1916600183610c58565b6112b48383836128ff565b336001600160a01b03167fe66fbe37d84ca73c589f782ac278844918ea6c56a4917f58707f715588080df28686868686346040516112f7969594939291906139be565b60405180910390a25050505050565b336113196000546001600160a01b031690565b6001600160a01b03161461135d5760405162461bcd60e51b81526020600482018190526024820152600080516020613ec38339815191526044820152606401610666565b6001600160a01b0381166113a55760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b6044820152606401610666565b600880546001600160a01b0319166001600160a01b0383169081179091556040519081527ffb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e890602001610701565b336114066000546001600160a01b031690565b6001600160a01b03161461144a5760405162461bcd60e51b81526020600482018190526024820152600080516020613ec38339815191526044820152606401610666565b6001600160a01b0381166114925760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b6044820152606401610666565b600780546001600160a01b0319166001600160a01b0383169081179091556040519081527fa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad90602001610701565b6114eb8383836128ff565b336001600160a01b03167fce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4858585853460405161152c9594939291906139ff565b60405180910390a250505050565b600061154588612997565b9050610f2f8a8a838a8a8a8a8a8a6040518060400160405280600881526020017f4d65737361676532000000000000000000000000000000000000000000000000815250612654565b336115a16000546001600160a01b031690565b6001600160a01b0316146115e55760405162461bcd60e51b81526020600482018190526024820152600080516020613ec38339815191526044820152606401610666565b60028190556040518181527f210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c90602001610701565b3361162d6000546001600160a01b031690565b6001600160a01b0316146116715760405162461bcd60e51b81526020600482018190526024820152600080516020613ec38339815191526044820152606401610666565b6001600160a01b0381166116ed5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610666565b6116f681612a2d565b50565b3361170c6000546001600160a01b031690565b6001600160a01b0316146117505760405162461bcd60e51b81526020600482018190526024820152600080516020613ec38339815191526044820152606401610666565b6001600160a01b0381166117985760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b6044820152606401610666565b600980546001600160a01b0319166001600160a01b0383169081179091556040519081527f918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb590602001610701565b6000546001600160a01b03161561183f5760405162461bcd60e51b815260206004820152601160248201527f6f776e657220616c7265616479207365740000000000000000000000000000006044820152606401610666565b61184833612a2d565b565b6005546001600160a01b0316156118a35760405162461bcd60e51b815260206004820152601b60248201527f6c697175696469747942726964676520616c72656164792073657400000000006044820152606401610666565b600580546001600160a01b03199081166001600160a01b03978816179091556006805482169587169590951790945560078054851693861693909317909255600880548416918516919091179055600980549092169216919091179055565b6000808060016119156020860186613895565b6006811115611926576119266134e3565b03611ab15761193b6040850160208601612f40565b61194b6060860160408701612f40565b61195b6080870160608801612f40565b608087013561197060e0890160c08a01613828565b6040516001600160601b0319606096871b8116602083015294861b851660348201529290941b9092166048820152605c8101919091526001600160c01b031960c092831b8116607c8301524690921b909116608482015260e0850135608c82015260ac0160408051808303601f19018152908290528051602090910120600554633c64f04b60e01b8352600483018290529093506001600160a01b031691508190633c64f04b90602401602060405180830381865afa158015611a37573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a5b9190613a3a565b1515600114611aac5760405162461bcd60e51b815260206004820152601660248201527f6272696467652072656c6179206e6f74206578697374000000000000000000006044820152606401610666565b612136565b6002611ac06020860186613895565b6006811115611ad157611ad16134e3565b03611c2e5746611ae760c0860160a08701613828565b611af76060870160408801612f40565b611b076080880160608901612f40565b6040516001600160c01b031960c095861b811660208301529390941b90921660288401526001600160601b0319606091821b8116603085015291901b1660448201526080850135605882015260780160408051808303601f19018152908290528051602090910120600554631c13568560e31b8352600483018290529093506001600160a01b03169150819063e09ab42890602401602060405180830381865afa158015611bb9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bdd9190613a3a565b1515600114611aac5760405162461bcd60e51b815260206004820152601960248201527f627269646765207769746864726177206e6f74206578697374000000000000006044820152606401610666565b6003611c3d6020860186613895565b6006811115611c4e57611c4e6134e3565b1480611c7757506004611c646020860186613895565b6006811115611c7557611c756134e3565b145b15611eda57611c8c6060850160408601612f40565b611c9c6080860160608701612f40565b6080860135611cb16040880160208901612f40565b611cc160e0890160c08a01613828565b604051606095861b6001600160601b0319908116602083015294861b851660348201526048810193909352931b909116606882015260c09190911b6001600160c01b031916607c82015260e0850135608482015260a40160408051601f19818403018152919052805160209091012091506003611d416020860186613895565b6006811115611d5257611d526134e3565b03611e1957506006546040516301e6472560e01b8152600481018390526001600160a01b039091169081906301e64725906024015b602060405180830381865afa158015611da4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dc89190613a3a565b1515600114611aac5760405162461bcd60e51b815260206004820152601560248201527f6d696e74207265636f7264206e6f7420657869737400000000000000000000006044820152606401610666565b506007546040516301e6472560e01b8152600481018390526001600160a01b039091169081906301e6472590602401602060405180830381865afa158015611e65573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e899190613a3a565b1515600114611aac5760405162461bcd60e51b815260206004820152601960248201527f7769746864726177207265636f7264206e6f74206578697374000000000000006044820152606401610666565b6005611ee96020860186613895565b6006811115611efa57611efa6134e3565b1480611f2357506006611f106020860186613895565b6006811115611f2157611f216134e3565b145b15612136576005611f376020860186613895565b6006811115611f4857611f486134e3565b03611f5f57506008546001600160a01b0316611f6d565b506009546001600160a01b03165b611f7d6060850160408601612f40565b611f8d6080860160608701612f40565b6080860135611fa26040880160208901612f40565b611fb260e0890160c08a01613828565b604051606095861b6001600160601b0319908116602083015294861b85166034820152604881019390935290841b8316606883015260c01b6001600160c01b031916607c82015260e087013560848201529183901b1660a482015260b80160408051601f198184030181529190528051602090910120915060056120396020860186613895565b600681111561204a5761204a6134e3565b0361207c576040516301e6472560e01b8152600481018390526001600160a01b038216906301e6472590602401611d87565b6040516301e6472560e01b8152600481018390526001600160a01b038216906301e6472590602401602060405180830381865afa1580156120c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120e59190613a3a565b15156001146121365760405162461bcd60e51b815260206004820152601960248201527f7769746864726177207265636f7264206e6f74206578697374000000000000006044820152606401610666565b6000818360405160200161214c93929190613a73565b6040516020818303038152906040528051906020012092505050919050565b6000805a90506000806121846060880160408901612f40565b6001600160a01b031634631f34afff60e21b6121a660408b0160208c01612f40565b6121b660808c0160608d01612f40565b60808c01356121cb60e08e0160c08f01613828565b8c8c336040516024016121e49796959493929190613a9f565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516122229190613af8565b60006040518083038185875af1925050503d806000811461225f576040519150601f19603f3d011682016040523d82523d6000602084013e612264565b606091505b5091509150811561228d57808060200190518101906122839190613b14565b935050505061229f565b6122978382612a7d565b600093505050505b9392505050565b6000805a90506000806122bf6060880160408901612f40565b6001600160a01b031634632d5bd7e360e11b6121a660408b0160208c01612f40565b6122f16060820160408301612f40565b6001600160a01b03167fa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d6000858561232f60e0870160c08801613828565b866101000135604051612346959493929190613b35565b60405180910390a2505050565b6001826006811115612367576123676134e3565b03612407576005546001600160a01b031663cdd1b25d61238783806138b6565b61239460208601866138fd565b6123a160408801886138fd565b6123ae60608a018a6138fd565b6040518963ffffffff1660e01b81526004016123d1989796959493929190613b73565b600060405180830381600087803b1580156123eb57600080fd5b505af11580156123ff573d6000803e3d6000fd5b505050505050565b600282600681111561241b5761241b6134e3565b0361243b576005546001600160a01b031663a21a928061238783806138b6565b600382600681111561244f5761244f6134e3565b0361246f576006546001600160a01b031663f873430261238783806138b6565b6005826006811115612483576124836134e3565b03612535576008546001600160a01b031663f87343026124a383806138b6565b6124b060208601866138fd565b6124bd60408801886138fd565b6124ca60608a018a6138fd565b6040518963ffffffff1660e01b81526004016124ed989796959493929190613b73565b6020604051808303816000875af115801561250c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125309190613bd3565b505050565b6004826006811115612549576125496134e3565b03612569576007546001600160a01b031663a21a928061238783806138b6565b600682600681111561257d5761257d6134e3565b03610d93576009546001600160a01b031663a21a92806124a383806138b6565b6040805160a081018252600080825260606020830181905292820181905291810182905260808101919091526040805160a08101909152806125e26020850185612f40565b6001600160a01b031681526020016040518060200160405280600081525081526020018360200160208101906126189190612f40565b6001600160a01b031681526020016126366060850160408601613828565b67ffffffffffffffff16815260200183606001358152509050919050565b6000612661898c8c612b08565b90506000808281526004602081905260409091205460ff1690811115612689576126896134e3565b146126d65760405162461bcd60e51b815260206004820152601860248201527f6d65737361676520616c726561647920657865637574656400000000000000006044820152606401610666565b6000818152600460208181526040808420805460ff191690931790925590516127059146913091879101613bec565b60408051601f1981840301815282825280516020918201206005549184018190528383018690528251808503840181526060850193849052633416de1160e11b90935293506001600160a01b03169163682dbc2291612772918d908d908d908d908d908d9060640161377f565b60006040518083038186803b15801561278a57600080fd5b505afa15801561279e573d6000803e3d6000fd5b505050506000806127b08c8f8f612b9c565b905060018160028111156127c6576127c66134e3565b036127d45760019150612858565b60028160028111156127e8576127e86134e3565b036128535760008481526004602052604090819020805460ff1916905560608d015160808e015191517fe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c9261284292600192899290613862565b60405180910390a150505050610f2f565b600291505b60008481526004602081905260409091208054849260ff19909116906001908490811115612888576128886134e3565b021790555061289884838e612d7d565b5050505050505050505050505050565b6000805a90506000806128c16060880160408901612f40565b6001600160a01b0316346305e5a4c160e11b6128e360808b0160608c01612f40565b8a608001358a8a336040516024016121e4959493929190613c2b565b4683036129405760405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9259608a1b6044820152606401610666565b600061294c8383610f3b565b9050803410156129915760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b6044820152606401610666565b50505050565b6040805160a081018252600080825260606020830181905292820181905291810182905260808101919091526040805160a0810190915260008152602081016129e084806138b6565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050908252506020908101906126189060408601908601612f40565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60005a90506000600a5445612a9291906137f3565b90508084108015612aad5750612aa9604085613c6a565b8211155b15612ab457fe5b6000612abf84612dcb565b9050612aca81612e2a565b7fffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f81604051612af99190613c8c565b60405180910390a15050505050565b60208301518051600091908203612b50578451604051612b3e919060200160609190911b6001600160601b031916815260140190565b60405160208183030381529060405290505b600181866040015187606001518860800151468989604051602001612b7c989796959493929190613c9f565b604051602081830303815290604052805190602001209150509392505050565b6000805a905060006060866020015151600003612c8c5786604001516001600160a01b0316346040518060600160405280602c8152602001613ee3602c91398051602090910120895160608b0151604051612c019291908c908c903390602401613d20565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051612c3f9190613af8565b60006040518083038185875af1925050503d8060008114612c7c576040519150601f19603f3d011682016040523d82523d6000602084013e612c81565b606091505b509092509050612d63565b86604001516001600160a01b0316346040518060600160405280602a8152602001613e99602a91398051906020012089602001518a606001518a8a33604051602401612cdc959493929190613d54565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051612d1a9190613af8565b60006040518083038185875af1925050503d8060008114612d57576040519150601f19603f3d011682016040523d82523d6000602084013e612d5c565b606091505b5090925090505b811561228d57808060200190518101906122839190613b14565b80604001516001600160a01b03167fa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d6001858585606001518660800151604051612346959493929190613b35565b6060604482511015612e1057505060408051808201909152601d81527f5472616e73616374696f6e2072657665727465642073696c656e746c79000000602082015290565b60048201915081806020019051810190610f5b9190613dbc565b60408051808201909152600b8082527f4d53473a3a41424f52543a000000000000000000000000000000000000000000602083015282518391116125305760005b8251811015612f0957828181518110612e8657612e86613e69565b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916828281518110612ec557612ec5613e69565b01602001517fff000000000000000000000000000000000000000000000000000000000000001614612ef75750505050565b80612f0181613e7f565b915050612e6b565b508260405162461bcd60e51b81526004016106669190613c8c565b80356001600160a01b0381168114612f3b57600080fd5b919050565b600060208284031215612f5257600080fd5b610f5882612f24565b600060208284031215612f6d57600080fd5b5035919050565b60008083601f840112612f8657600080fd5b50813567ffffffffffffffff811115612f9e57600080fd5b6020830191508360208260051b8501011115612fb957600080fd5b9250929050565b60008060008060008060008060a0898b031215612fdc57600080fd5b612fe589612f24565b975060208901359650604089013567ffffffffffffffff8082111561300957600080fd5b6130158c838d01612f74565b909850965060608b013591508082111561302e57600080fd5b61303a8c838d01612f74565b909650945060808b013591508082111561305357600080fd5b506130608b828c01612f74565b999c989b5096995094979396929594505050565b600080600080600060a0868803121561308c57600080fd5b61309586612f24565b94506130a360208701612f24565b93506130b160408701612f24565b92506130bf60608701612f24565b91506130cd60808701612f24565b90509295509295909350565b60008083601f8401126130eb57600080fd5b50813567ffffffffffffffff81111561310357600080fd5b602083019150836020828501011115612fb957600080fd5b6000806000806000806000806000898b036101a081121561313b57600080fd5b8a3567ffffffffffffffff8082111561315357600080fd5b61315f8e838f016130d9565b909c509a508a9150610120601f198401121561317a57600080fd5b60208d0199506101408d013592508083111561319557600080fd5b6131a18e848f01612f74565b90995097506101608d01359250889150808311156131be57600080fd5b6131ca8e848f01612f74565b90975095506101808d01359250869150808311156131e757600080fd5b50506131f58c828d01612f74565b915080935050809150509295985092959850929598565b60006080828403121561321e57600080fd5b50919050565b6000806040838503121561323757600080fd5b823567ffffffffffffffff8082111561324f57600080fd5b61325b8683870161320c565b9350602085013591508082111561327157600080fd5b5083016101a0818603121561328557600080fd5b809150509250929050565b60008060008060008060a087890312156132a957600080fd5b6132b287612f24565b9550602087013594506132c760408801612f24565b935060608701359250608087013567ffffffffffffffff8111156132ea57600080fd5b6132f689828a016130d9565b979a9699509497509295939492505050565b60008060008060008060008060006101008a8c03121561332757600080fd5b893567ffffffffffffffff8082111561333f57600080fd5b61334b8d838e016130d9565b909b5099508991506133608d60208e0161320c565b985060a08c013591508082111561337657600080fd5b6133828d838e01612f74565b909850965060c08c013591508082111561339b57600080fd5b6133a78d838e01612f74565b909650945060e08c01359150808211156133c057600080fd5b506131f58c828d01612f74565b600080602083850312156133e057600080fd5b823567ffffffffffffffff8111156133f757600080fd5b613403858286016130d9565b90969095509350505050565b60008060008060006060868803121561342757600080fd5b853567ffffffffffffffff8082111561343f57600080fd5b61344b89838a016130d9565b909750955060208801359450604088013591508082111561346b57600080fd5b50613478888289016130d9565b969995985093965092949392505050565b6000806000806060858703121561349f57600080fd5b6134a885612f24565b935060208501359250604085013567ffffffffffffffff8111156134cb57600080fd5b6134d7878288016130d9565b95989497509550505050565b634e487b7160e01b600052602160045260246000fd5b60058110613509576135096134e3565b9052565b60208101610f5b82846134f9565b600080600080600080600080600060a08a8c03121561353957600080fd5b893567ffffffffffffffff8082111561355157600080fd5b61355d8d838e016130d9565b909b50995060208c013591508082111561357657600080fd5b6135828d838e0161320c565b985060408c013591508082111561359857600080fd5b6135a48d838e01612f74565b909850965060608c01359150808211156135bd57600080fd5b6135c98d838e01612f74565b909650945060808c01359150808211156133c057600080fd5b60005b838110156135fd5781810151838201526020016135e5565b50506000910152565b6000815180845261361e8160208601602086016135e2565b601f01601f19169290920160200192915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b81835260006020808501808196508560051b810191508460005b878110156136e05782840389528135601e1988360301811261369657600080fd5b8701858101903567ffffffffffffffff8111156136b257600080fd5b8036038213156136c157600080fd5b6136cc868284613632565b9a87019a9550505090840190600101613675565b5091979650505050505050565b8183526000602080850194508260005b85811015613729576001600160a01b0361371683612f24565b16875295820195908201906001016136fd565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83111561376657600080fd5b8260051b80836020870137939093016020019392505050565b608081526000613792608083018a613606565b82810360208401526137a581898b61365b565b905082810360408401526137ba8187896136ed565b905082810360608401526137cf818587613734565b9a9950505050505050505050565b634e487b7160e01b600052601160045260246000fd5b81810381811115610f5b57610f5b6137dd565b8581528460208201528284604083013760409201918201526060019392505050565b60006020828403121561383a57600080fd5b813567ffffffffffffffff8116811461229f57600080fd5b60028110613509576135096134e3565b608081016138708287613852565b84602083015267ffffffffffffffff8416604083015282606083015295945050505050565b6000602082840312156138a757600080fd5b81356007811061229f57600080fd5b6000808335601e198436030181126138cd57600080fd5b83018035915067ffffffffffffffff8211156138e857600080fd5b602001915036819003821315612fb957600080fd5b6000808335601e1984360301811261391457600080fd5b83018035915067ffffffffffffffff82111561392f57600080fd5b6020019150600581901b3603821315612fb957600080fd5b60006001600160a01b03808a16835288602084015280881660408401525085606083015260c0608083015261398060c083018587613632565b90508260a083015298975050505050505050565b8082028115828204841417610f5b57610f5b6137dd565b80820180821115610f5b57610f5b6137dd565b6080815260006139d260808301888a613632565b86602084015282810360408401526139eb818688613632565b915050826060830152979650505050505050565b6001600160a01b0386168152846020820152608060408201526000613a28608083018587613632565b90508260608301529695505050505050565b600060208284031215613a4c57600080fd5b8151801515811461229f57600080fd5b60028110613a6c57613a6c6134e3565b60f81b9052565b613a7d8185613a5c565b60609290921b6001600160601b03191660018301526015820152603501919050565b60006001600160a01b03808a168352808916602084015287604084015267ffffffffffffffff8716606084015260c06080840152613ae160c084018688613632565b915080841660a08401525098975050505050505050565b60008251613b0a8184602087016135e2565b9190910192915050565b600060208284031215613b2657600080fd5b81516003811061229f57600080fd5b60a08101613b438288613852565b856020830152613b5660408301866134f9565b67ffffffffffffffff939093166060820152608001529392505050565b608081526000613b87608083018a8c613632565b8281036020840152613b9a81898b61365b565b90508281036040840152613baf8187896136ed565b90508281036060840152613bc4818587613734565b9b9a5050505050505050505050565b600060208284031215613be557600080fd5b5051919050565b8381526bffffffffffffffffffffffff198360601b16602082015260008251613c1c8160348501602087016135e2565b91909101603401949350505050565b60006001600160a01b03808816835286602084015260806040840152613c55608084018688613632565b91508084166060840152509695505050505050565b600082613c8757634e487b7160e01b600052601260045260246000fd5b500490565b602081526000610f586020830184613606565b613ca9818a613a5c565b60008851613cbe816001850160208d016135e2565b80830190506bffffffffffffffffffffffff198960601b1660018201526001600160c01b0319808960c01b16601583015287601d830152808760c01b16603d830152508385604583013760009301604501928352509098975050505050505050565b60006001600160a01b03808816835267ffffffffffffffff8716602084015260806040840152613c55608084018688613632565b608081526000613d676080830188613606565b67ffffffffffffffff871660208401528281036040840152613d8a818688613632565b9150506001600160a01b03831660608301529695505050505050565b634e487b7160e01b600052604160045260246000fd5b600060208284031215613dce57600080fd5b815167ffffffffffffffff80821115613de657600080fd5b818401915084601f830112613dfa57600080fd5b815181811115613e0c57613e0c613da6565b604051601f8201601f19908116603f01168101908382118183101715613e3457613e34613da6565b81604052828152876020848701011115613e4d57600080fd5b613e5e8360208301602088016135e2565b979650505050505050565b634e487b7160e01b600052603260045260246000fd5b600060018201613e9157613e916137dd565b506001019056fe657865637574654d6573736167652862797465732c75696e7436342c62797465732c61646472657373294f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572657865637574654d65737361676528616464726573732c75696e7436342c62797465732c6164647265737329a26469706673582212200ee606fc1151b84213c23dfd294e0e281128248c80b28c385c3b71ff0a42109564736f6c63430008110033",
}

// MessageBusABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageBusMetaData.ABI instead.
var MessageBusABI = MessageBusMetaData.ABI

// MessageBusBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageBusMetaData.Bin instead.
var MessageBusBin = MessageBusMetaData.Bin

// DeployMessageBus deploys a new Ethereum contract, binding an instance of MessageBus to it.
func DeployMessageBus(auth *bind.TransactOpts, backend bind.ContractBackend, _sigsVerifier common.Address, _liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address, _pegBridgeV2 common.Address, _pegVaultV2 common.Address) (common.Address, *types.Transaction, *MessageBus, error) {
	parsed, err := MessageBusMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBusBin), backend, _sigsVerifier, _liquidityBridge, _pegBridge, _pegVault, _pegBridgeV2, _pegVaultV2)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageBus{MessageBusCaller: MessageBusCaller{contract: contract}, MessageBusTransactor: MessageBusTransactor{contract: contract}, MessageBusFilterer: MessageBusFilterer{contract: contract}}, nil
}

// MessageBus is an auto generated Go binding around an Ethereum contract.
type MessageBus struct {
	MessageBusCaller     // Read-only binding to the contract
	MessageBusTransactor // Write-only binding to the contract
	MessageBusFilterer   // Log filterer for contract events
}

// MessageBusCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageBusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageBusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageBusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageBusSession struct {
	Contract     *MessageBus       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageBusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageBusCallerSession struct {
	Contract *MessageBusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MessageBusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageBusTransactorSession struct {
	Contract     *MessageBusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MessageBusRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageBusRaw struct {
	Contract *MessageBus // Generic contract binding to access the raw methods on
}

// MessageBusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageBusCallerRaw struct {
	Contract *MessageBusCaller // Generic read-only contract binding to access the raw methods on
}

// MessageBusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageBusTransactorRaw struct {
	Contract *MessageBusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageBus creates a new instance of MessageBus, bound to a specific deployed contract.
func NewMessageBus(address common.Address, backend bind.ContractBackend) (*MessageBus, error) {
	contract, err := bindMessageBus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageBus{MessageBusCaller: MessageBusCaller{contract: contract}, MessageBusTransactor: MessageBusTransactor{contract: contract}, MessageBusFilterer: MessageBusFilterer{contract: contract}}, nil
}

// NewMessageBusCaller creates a new read-only instance of MessageBus, bound to a specific deployed contract.
func NewMessageBusCaller(address common.Address, caller bind.ContractCaller) (*MessageBusCaller, error) {
	contract, err := bindMessageBus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusCaller{contract: contract}, nil
}

// NewMessageBusTransactor creates a new write-only instance of MessageBus, bound to a specific deployed contract.
func NewMessageBusTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageBusTransactor, error) {
	contract, err := bindMessageBus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusTransactor{contract: contract}, nil
}

// NewMessageBusFilterer creates a new log filterer instance of MessageBus, bound to a specific deployed contract.
func NewMessageBusFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageBusFilterer, error) {
	contract, err := bindMessageBus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageBusFilterer{contract: contract}, nil
}

// bindMessageBus binds a generic wrapper to an already deployed contract.
func bindMessageBus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageBusABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBus *MessageBusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBus.Contract.MessageBusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBus *MessageBusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBus.Contract.MessageBusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBus *MessageBusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBus.Contract.MessageBusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBus *MessageBusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBus *MessageBusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBus *MessageBusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBus.Contract.contract.Transact(opts, method, params...)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBus *MessageBusCaller) CalcFee(opts *bind.CallOpts, _message []byte) (*big.Int, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "calcFee", _message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBus *MessageBusSession) CalcFee(_message []byte) (*big.Int, error) {
	return _MessageBus.Contract.CalcFee(&_MessageBus.CallOpts, _message)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBus *MessageBusCallerSession) CalcFee(_message []byte) (*big.Int, error) {
	return _MessageBus.Contract.CalcFee(&_MessageBus.CallOpts, _message)
}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBus *MessageBusCaller) ExecutedMessages(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "executedMessages", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBus *MessageBusSession) ExecutedMessages(arg0 [32]byte) (uint8, error) {
	return _MessageBus.Contract.ExecutedMessages(&_MessageBus.CallOpts, arg0)
}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBus *MessageBusCallerSession) ExecutedMessages(arg0 [32]byte) (uint8, error) {
	return _MessageBus.Contract.ExecutedMessages(&_MessageBus.CallOpts, arg0)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBus *MessageBusCaller) FeeBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "feeBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBus *MessageBusSession) FeeBase() (*big.Int, error) {
	return _MessageBus.Contract.FeeBase(&_MessageBus.CallOpts)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBus *MessageBusCallerSession) FeeBase() (*big.Int, error) {
	return _MessageBus.Contract.FeeBase(&_MessageBus.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBus *MessageBusCaller) FeePerByte(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "feePerByte")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBus *MessageBusSession) FeePerByte() (*big.Int, error) {
	return _MessageBus.Contract.FeePerByte(&_MessageBus.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBus *MessageBusCallerSession) FeePerByte() (*big.Int, error) {
	return _MessageBus.Contract.FeePerByte(&_MessageBus.CallOpts)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBus *MessageBusCaller) LiquidityBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "liquidityBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBus *MessageBusSession) LiquidityBridge() (common.Address, error) {
	return _MessageBus.Contract.LiquidityBridge(&_MessageBus.CallOpts)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBus *MessageBusCallerSession) LiquidityBridge() (common.Address, error) {
	return _MessageBus.Contract.LiquidityBridge(&_MessageBus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBus *MessageBusCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBus *MessageBusSession) Owner() (common.Address, error) {
	return _MessageBus.Contract.Owner(&_MessageBus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBus *MessageBusCallerSession) Owner() (common.Address, error) {
	return _MessageBus.Contract.Owner(&_MessageBus.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBus *MessageBusCaller) PegBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "pegBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBus *MessageBusSession) PegBridge() (common.Address, error) {
	return _MessageBus.Contract.PegBridge(&_MessageBus.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBus *MessageBusCallerSession) PegBridge() (common.Address, error) {
	return _MessageBus.Contract.PegBridge(&_MessageBus.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBus *MessageBusCaller) PegBridgeV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "pegBridgeV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBus *MessageBusSession) PegBridgeV2() (common.Address, error) {
	return _MessageBus.Contract.PegBridgeV2(&_MessageBus.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBus *MessageBusCallerSession) PegBridgeV2() (common.Address, error) {
	return _MessageBus.Contract.PegBridgeV2(&_MessageBus.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBus *MessageBusCaller) PegVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "pegVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBus *MessageBusSession) PegVault() (common.Address, error) {
	return _MessageBus.Contract.PegVault(&_MessageBus.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBus *MessageBusCallerSession) PegVault() (common.Address, error) {
	return _MessageBus.Contract.PegVault(&_MessageBus.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBus *MessageBusCaller) PegVaultV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "pegVaultV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBus *MessageBusSession) PegVaultV2() (common.Address, error) {
	return _MessageBus.Contract.PegVaultV2(&_MessageBus.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBus *MessageBusCallerSession) PegVaultV2() (common.Address, error) {
	return _MessageBus.Contract.PegVaultV2(&_MessageBus.CallOpts)
}

// PreExecuteMessageGasUsage is a free data retrieval call binding the contract method 0x584e45e1.
//
// Solidity: function preExecuteMessageGasUsage() view returns(uint256)
func (_MessageBus *MessageBusCaller) PreExecuteMessageGasUsage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "preExecuteMessageGasUsage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreExecuteMessageGasUsage is a free data retrieval call binding the contract method 0x584e45e1.
//
// Solidity: function preExecuteMessageGasUsage() view returns(uint256)
func (_MessageBus *MessageBusSession) PreExecuteMessageGasUsage() (*big.Int, error) {
	return _MessageBus.Contract.PreExecuteMessageGasUsage(&_MessageBus.CallOpts)
}

// PreExecuteMessageGasUsage is a free data retrieval call binding the contract method 0x584e45e1.
//
// Solidity: function preExecuteMessageGasUsage() view returns(uint256)
func (_MessageBus *MessageBusCallerSession) PreExecuteMessageGasUsage() (*big.Int, error) {
	return _MessageBus.Contract.PreExecuteMessageGasUsage(&_MessageBus.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBus *MessageBusCaller) SigsVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "sigsVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBus *MessageBusSession) SigsVerifier() (common.Address, error) {
	return _MessageBus.Contract.SigsVerifier(&_MessageBus.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBus *MessageBusCallerSession) SigsVerifier() (common.Address, error) {
	return _MessageBus.Contract.SigsVerifier(&_MessageBus.CallOpts)
}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBus *MessageBusCaller) WithdrawnFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MessageBus.contract.Call(opts, &out, "withdrawnFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBus *MessageBusSession) WithdrawnFees(arg0 common.Address) (*big.Int, error) {
	return _MessageBus.Contract.WithdrawnFees(&_MessageBus.CallOpts, arg0)
}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBus *MessageBusCallerSession) WithdrawnFees(arg0 common.Address) (*big.Int, error) {
	return _MessageBus.Contract.WithdrawnFees(&_MessageBus.CallOpts, arg0)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactor) ExecuteMessage(opts *bind.TransactOpts, _message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "executeMessage", _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusSession) ExecuteMessage(_message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessage(&_MessageBus.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactorSession) ExecuteMessage(_message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessage(&_MessageBus.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0xdb2c20c8.
//
// Solidity: function executeMessage(bytes _message, (bytes,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactor) ExecuteMessage0(opts *bind.TransactOpts, _message []byte, _route MsgDataTypesRouteInfo2, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "executeMessage0", _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0xdb2c20c8.
//
// Solidity: function executeMessage(bytes _message, (bytes,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusSession) ExecuteMessage0(_message []byte, _route MsgDataTypesRouteInfo2, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessage0(&_MessageBus.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0xdb2c20c8.
//
// Solidity: function executeMessage(bytes _message, (bytes,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactorSession) ExecuteMessage0(_message []byte, _route MsgDataTypesRouteInfo2, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessage0(&_MessageBus.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "executeMessageWithTransfer", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusSession) ExecuteMessageWithTransfer(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessageWithTransfer(&_MessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactorSession) ExecuteMessageWithTransfer(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessageWithTransfer(&_MessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "executeMessageWithTransferRefund", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessageWithTransferRefund(&_MessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBus *MessageBusTransactorSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.ExecuteMessageWithTransferRefund(&_MessageBus.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(address _liquidityBridge, address _pegBridge, address _pegVault, address _pegBridgeV2, address _pegVaultV2) returns()
func (_MessageBus *MessageBusTransactor) Init(opts *bind.TransactOpts, _liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address, _pegBridgeV2 common.Address, _pegVaultV2 common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "init", _liquidityBridge, _pegBridge, _pegVault, _pegBridgeV2, _pegVaultV2)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(address _liquidityBridge, address _pegBridge, address _pegVault, address _pegBridgeV2, address _pegVaultV2) returns()
func (_MessageBus *MessageBusSession) Init(_liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address, _pegBridgeV2 common.Address, _pegVaultV2 common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.Init(&_MessageBus.TransactOpts, _liquidityBridge, _pegBridge, _pegVault, _pegBridgeV2, _pegVaultV2)
}

// Init is a paid mutator transaction binding the contract method 0x359ef75b.
//
// Solidity: function init(address _liquidityBridge, address _pegBridge, address _pegVault, address _pegBridgeV2, address _pegVaultV2) returns()
func (_MessageBus *MessageBusTransactorSession) Init(_liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address, _pegBridgeV2 common.Address, _pegVaultV2 common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.Init(&_MessageBus.TransactOpts, _liquidityBridge, _pegBridge, _pegVault, _pegBridgeV2, _pegVaultV2)
}

// RefundAndExecuteMsg is a paid mutator transaction binding the contract method 0x40d0d026.
//
// Solidity: function refundAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBus *MessageBusTransactor) RefundAndExecuteMsg(opts *bind.TransactOpts, _transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "refundAndExecuteMsg", _transferParams, _msgParams)
}

// RefundAndExecuteMsg is a paid mutator transaction binding the contract method 0x40d0d026.
//
// Solidity: function refundAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBus *MessageBusSession) RefundAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBus.Contract.RefundAndExecuteMsg(&_MessageBus.TransactOpts, _transferParams, _msgParams)
}

// RefundAndExecuteMsg is a paid mutator transaction binding the contract method 0x40d0d026.
//
// Solidity: function refundAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBus *MessageBusTransactorSession) RefundAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBus.Contract.RefundAndExecuteMsg(&_MessageBus.TransactOpts, _transferParams, _msgParams)
}

// SendMessage is a paid mutator transaction binding the contract method 0x7d7a101d.
//
// Solidity: function sendMessage(bytes _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBus *MessageBusTransactor) SendMessage(opts *bind.TransactOpts, _receiver []byte, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "sendMessage", _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x7d7a101d.
//
// Solidity: function sendMessage(bytes _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBus *MessageBusSession) SendMessage(_receiver []byte, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBus.Contract.SendMessage(&_MessageBus.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x7d7a101d.
//
// Solidity: function sendMessage(bytes _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBus *MessageBusTransactorSession) SendMessage(_receiver []byte, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBus.Contract.SendMessage(&_MessageBus.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBus *MessageBusTransactor) SendMessage0(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "sendMessage0", _receiver, _dstChainId, _message)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBus *MessageBusSession) SendMessage0(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBus.Contract.SendMessage0(&_MessageBus.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBus *MessageBusTransactorSession) SendMessage0(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBus.Contract.SendMessage0(&_MessageBus.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBus *MessageBusTransactor) SendMessageWithTransfer(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "sendMessageWithTransfer", _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBus *MessageBusSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBus.Contract.SendMessageWithTransfer(&_MessageBus.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBus *MessageBusTransactorSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBus.Contract.SendMessageWithTransfer(&_MessageBus.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBus *MessageBusTransactor) SetFeeBase(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setFeeBase", _fee)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBus *MessageBusSession) SetFeeBase(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetFeeBase(&_MessageBus.TransactOpts, _fee)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBus *MessageBusTransactorSession) SetFeeBase(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetFeeBase(&_MessageBus.TransactOpts, _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBus *MessageBusTransactor) SetFeePerByte(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setFeePerByte", _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBus *MessageBusSession) SetFeePerByte(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetFeePerByte(&_MessageBus.TransactOpts, _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBus *MessageBusTransactorSession) SetFeePerByte(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetFeePerByte(&_MessageBus.TransactOpts, _fee)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBus *MessageBusTransactor) SetLiquidityBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setLiquidityBridge", _addr)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBus *MessageBusSession) SetLiquidityBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetLiquidityBridge(&_MessageBus.TransactOpts, _addr)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBus *MessageBusTransactorSession) SetLiquidityBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetLiquidityBridge(&_MessageBus.TransactOpts, _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBus *MessageBusTransactor) SetPegBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setPegBridge", _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBus *MessageBusSession) SetPegBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegBridge(&_MessageBus.TransactOpts, _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBus *MessageBusTransactorSession) SetPegBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegBridge(&_MessageBus.TransactOpts, _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBus *MessageBusTransactor) SetPegBridgeV2(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setPegBridgeV2", _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBus *MessageBusSession) SetPegBridgeV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegBridgeV2(&_MessageBus.TransactOpts, _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBus *MessageBusTransactorSession) SetPegBridgeV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegBridgeV2(&_MessageBus.TransactOpts, _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBus *MessageBusTransactor) SetPegVault(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setPegVault", _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBus *MessageBusSession) SetPegVault(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegVault(&_MessageBus.TransactOpts, _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBus *MessageBusTransactorSession) SetPegVault(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegVault(&_MessageBus.TransactOpts, _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBus *MessageBusTransactor) SetPegVaultV2(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setPegVaultV2", _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBus *MessageBusSession) SetPegVaultV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegVaultV2(&_MessageBus.TransactOpts, _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBus *MessageBusTransactorSession) SetPegVaultV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPegVaultV2(&_MessageBus.TransactOpts, _addr)
}

// SetPreExecuteMessageGasUsage is a paid mutator transaction binding the contract method 0x4586f331.
//
// Solidity: function setPreExecuteMessageGasUsage(uint256 _usage) returns()
func (_MessageBus *MessageBusTransactor) SetPreExecuteMessageGasUsage(opts *bind.TransactOpts, _usage *big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "setPreExecuteMessageGasUsage", _usage)
}

// SetPreExecuteMessageGasUsage is a paid mutator transaction binding the contract method 0x4586f331.
//
// Solidity: function setPreExecuteMessageGasUsage(uint256 _usage) returns()
func (_MessageBus *MessageBusSession) SetPreExecuteMessageGasUsage(_usage *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPreExecuteMessageGasUsage(&_MessageBus.TransactOpts, _usage)
}

// SetPreExecuteMessageGasUsage is a paid mutator transaction binding the contract method 0x4586f331.
//
// Solidity: function setPreExecuteMessageGasUsage(uint256 _usage) returns()
func (_MessageBus *MessageBusTransactorSession) SetPreExecuteMessageGasUsage(_usage *big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.SetPreExecuteMessageGasUsage(&_MessageBus.TransactOpts, _usage)
}

// TransferAndExecuteMsg is a paid mutator transaction binding the contract method 0x723d0a9d.
//
// Solidity: function transferAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBus *MessageBusTransactor) TransferAndExecuteMsg(opts *bind.TransactOpts, _transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "transferAndExecuteMsg", _transferParams, _msgParams)
}

// TransferAndExecuteMsg is a paid mutator transaction binding the contract method 0x723d0a9d.
//
// Solidity: function transferAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBus *MessageBusSession) TransferAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBus.Contract.TransferAndExecuteMsg(&_MessageBus.TransactOpts, _transferParams, _msgParams)
}

// TransferAndExecuteMsg is a paid mutator transaction binding the contract method 0x723d0a9d.
//
// Solidity: function transferAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBus *MessageBusTransactorSession) TransferAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBus.Contract.TransferAndExecuteMsg(&_MessageBus.TransactOpts, _transferParams, _msgParams)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBus *MessageBusTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBus *MessageBusSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.TransferOwnership(&_MessageBus.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBus *MessageBusTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBus.Contract.TransferOwnership(&_MessageBus.TransactOpts, newOwner)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBus *MessageBusTransactor) WithdrawFee(opts *bind.TransactOpts, _account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.contract.Transact(opts, "withdrawFee", _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBus *MessageBusSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.WithdrawFee(&_MessageBus.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBus *MessageBusTransactorSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBus.Contract.WithdrawFee(&_MessageBus.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// MessageBusCallRevertedIterator is returned from FilterCallReverted and is used to iterate over the raw logs and unpacked data for CallReverted events raised by the MessageBus contract.
type MessageBusCallRevertedIterator struct {
	Event *MessageBusCallReverted // Event containing the contract specifics and raw log

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
func (it *MessageBusCallRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusCallReverted)
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
		it.Event = new(MessageBusCallReverted)
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
func (it *MessageBusCallRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusCallRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusCallReverted represents a CallReverted event raised by the MessageBus contract.
type MessageBusCallReverted struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallReverted is a free log retrieval operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBus *MessageBusFilterer) FilterCallReverted(opts *bind.FilterOpts) (*MessageBusCallRevertedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return &MessageBusCallRevertedIterator{contract: _MessageBus.contract, event: "CallReverted", logs: logs, sub: sub}, nil
}

// WatchCallReverted is a free log subscription operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBus *MessageBusFilterer) WatchCallReverted(opts *bind.WatchOpts, sink chan<- *MessageBusCallReverted) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusCallReverted)
				if err := _MessageBus.contract.UnpackLog(event, "CallReverted", log); err != nil {
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

// ParseCallReverted is a log parse operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBus *MessageBusFilterer) ParseCallReverted(log types.Log) (*MessageBusCallReverted, error) {
	event := new(MessageBusCallReverted)
	if err := _MessageBus.contract.UnpackLog(event, "CallReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the MessageBus contract.
type MessageBusExecutedIterator struct {
	Event *MessageBusExecuted // Event containing the contract specifics and raw log

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
func (it *MessageBusExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusExecuted)
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
		it.Event = new(MessageBusExecuted)
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
func (it *MessageBusExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusExecuted represents a Executed event raised by the MessageBus contract.
type MessageBusExecuted struct {
	MsgType    uint8
	MsgId      [32]byte
	Status     uint8
	Receiver   common.Address
	SrcChainId uint64
	SrcTxHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d.
//
// Solidity: event Executed(uint8 msgType, bytes32 msgId, uint8 status, address indexed receiver, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBus *MessageBusFilterer) FilterExecuted(opts *bind.FilterOpts, receiver []common.Address) (*MessageBusExecutedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "Executed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusExecutedIterator{contract: _MessageBus.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d.
//
// Solidity: event Executed(uint8 msgType, bytes32 msgId, uint8 status, address indexed receiver, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBus *MessageBusFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *MessageBusExecuted, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "Executed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusExecuted)
				if err := _MessageBus.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0xa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d.
//
// Solidity: event Executed(uint8 msgType, bytes32 msgId, uint8 status, address indexed receiver, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBus *MessageBusFilterer) ParseExecuted(log types.Log) (*MessageBusExecuted, error) {
	event := new(MessageBusExecuted)
	if err := _MessageBus.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusFeeBaseUpdatedIterator is returned from FilterFeeBaseUpdated and is used to iterate over the raw logs and unpacked data for FeeBaseUpdated events raised by the MessageBus contract.
type MessageBusFeeBaseUpdatedIterator struct {
	Event *MessageBusFeeBaseUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusFeeBaseUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusFeeBaseUpdated)
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
		it.Event = new(MessageBusFeeBaseUpdated)
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
func (it *MessageBusFeeBaseUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusFeeBaseUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusFeeBaseUpdated represents a FeeBaseUpdated event raised by the MessageBus contract.
type MessageBusFeeBaseUpdated struct {
	FeeBase *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeBaseUpdated is a free log retrieval operation binding the contract event 0x892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f2.
//
// Solidity: event FeeBaseUpdated(uint256 feeBase)
func (_MessageBus *MessageBusFilterer) FilterFeeBaseUpdated(opts *bind.FilterOpts) (*MessageBusFeeBaseUpdatedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "FeeBaseUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusFeeBaseUpdatedIterator{contract: _MessageBus.contract, event: "FeeBaseUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeBaseUpdated is a free log subscription operation binding the contract event 0x892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f2.
//
// Solidity: event FeeBaseUpdated(uint256 feeBase)
func (_MessageBus *MessageBusFilterer) WatchFeeBaseUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusFeeBaseUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "FeeBaseUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusFeeBaseUpdated)
				if err := _MessageBus.contract.UnpackLog(event, "FeeBaseUpdated", log); err != nil {
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

// ParseFeeBaseUpdated is a log parse operation binding the contract event 0x892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f2.
//
// Solidity: event FeeBaseUpdated(uint256 feeBase)
func (_MessageBus *MessageBusFilterer) ParseFeeBaseUpdated(log types.Log) (*MessageBusFeeBaseUpdated, error) {
	event := new(MessageBusFeeBaseUpdated)
	if err := _MessageBus.contract.UnpackLog(event, "FeeBaseUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusFeePerByteUpdatedIterator is returned from FilterFeePerByteUpdated and is used to iterate over the raw logs and unpacked data for FeePerByteUpdated events raised by the MessageBus contract.
type MessageBusFeePerByteUpdatedIterator struct {
	Event *MessageBusFeePerByteUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusFeePerByteUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusFeePerByteUpdated)
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
		it.Event = new(MessageBusFeePerByteUpdated)
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
func (it *MessageBusFeePerByteUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusFeePerByteUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusFeePerByteUpdated represents a FeePerByteUpdated event raised by the MessageBus contract.
type MessageBusFeePerByteUpdated struct {
	FeePerByte *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeePerByteUpdated is a free log retrieval operation binding the contract event 0x210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c.
//
// Solidity: event FeePerByteUpdated(uint256 feePerByte)
func (_MessageBus *MessageBusFilterer) FilterFeePerByteUpdated(opts *bind.FilterOpts) (*MessageBusFeePerByteUpdatedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "FeePerByteUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusFeePerByteUpdatedIterator{contract: _MessageBus.contract, event: "FeePerByteUpdated", logs: logs, sub: sub}, nil
}

// WatchFeePerByteUpdated is a free log subscription operation binding the contract event 0x210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c.
//
// Solidity: event FeePerByteUpdated(uint256 feePerByte)
func (_MessageBus *MessageBusFilterer) WatchFeePerByteUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusFeePerByteUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "FeePerByteUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusFeePerByteUpdated)
				if err := _MessageBus.contract.UnpackLog(event, "FeePerByteUpdated", log); err != nil {
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

// ParseFeePerByteUpdated is a log parse operation binding the contract event 0x210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c.
//
// Solidity: event FeePerByteUpdated(uint256 feePerByte)
func (_MessageBus *MessageBusFilterer) ParseFeePerByteUpdated(log types.Log) (*MessageBusFeePerByteUpdated, error) {
	event := new(MessageBusFeePerByteUpdated)
	if err := _MessageBus.contract.UnpackLog(event, "FeePerByteUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusFeeWithdrawnIterator is returned from FilterFeeWithdrawn and is used to iterate over the raw logs and unpacked data for FeeWithdrawn events raised by the MessageBus contract.
type MessageBusFeeWithdrawnIterator struct {
	Event *MessageBusFeeWithdrawn // Event containing the contract specifics and raw log

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
func (it *MessageBusFeeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusFeeWithdrawn)
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
		it.Event = new(MessageBusFeeWithdrawn)
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
func (it *MessageBusFeeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusFeeWithdrawn represents a FeeWithdrawn event raised by the MessageBus contract.
type MessageBusFeeWithdrawn struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeeWithdrawn is a free log retrieval operation binding the contract event 0x78473f3f373f7673597f4f0fa5873cb4d375fea6d4339ad6b56dbd411513cb3f.
//
// Solidity: event FeeWithdrawn(address receiver, uint256 amount)
func (_MessageBus *MessageBusFilterer) FilterFeeWithdrawn(opts *bind.FilterOpts) (*MessageBusFeeWithdrawnIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "FeeWithdrawn")
	if err != nil {
		return nil, err
	}
	return &MessageBusFeeWithdrawnIterator{contract: _MessageBus.contract, event: "FeeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFeeWithdrawn is a free log subscription operation binding the contract event 0x78473f3f373f7673597f4f0fa5873cb4d375fea6d4339ad6b56dbd411513cb3f.
//
// Solidity: event FeeWithdrawn(address receiver, uint256 amount)
func (_MessageBus *MessageBusFilterer) WatchFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *MessageBusFeeWithdrawn) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "FeeWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusFeeWithdrawn)
				if err := _MessageBus.contract.UnpackLog(event, "FeeWithdrawn", log); err != nil {
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

// ParseFeeWithdrawn is a log parse operation binding the contract event 0x78473f3f373f7673597f4f0fa5873cb4d375fea6d4339ad6b56dbd411513cb3f.
//
// Solidity: event FeeWithdrawn(address receiver, uint256 amount)
func (_MessageBus *MessageBusFilterer) ParseFeeWithdrawn(log types.Log) (*MessageBusFeeWithdrawn, error) {
	event := new(MessageBusFeeWithdrawn)
	if err := _MessageBus.contract.UnpackLog(event, "FeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusLiquidityBridgeUpdatedIterator is returned from FilterLiquidityBridgeUpdated and is used to iterate over the raw logs and unpacked data for LiquidityBridgeUpdated events raised by the MessageBus contract.
type MessageBusLiquidityBridgeUpdatedIterator struct {
	Event *MessageBusLiquidityBridgeUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusLiquidityBridgeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusLiquidityBridgeUpdated)
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
		it.Event = new(MessageBusLiquidityBridgeUpdated)
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
func (it *MessageBusLiquidityBridgeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusLiquidityBridgeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusLiquidityBridgeUpdated represents a LiquidityBridgeUpdated event raised by the MessageBus contract.
type MessageBusLiquidityBridgeUpdated struct {
	LiquidityBridge common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidityBridgeUpdated is a free log retrieval operation binding the contract event 0xbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae058.
//
// Solidity: event LiquidityBridgeUpdated(address liquidityBridge)
func (_MessageBus *MessageBusFilterer) FilterLiquidityBridgeUpdated(opts *bind.FilterOpts) (*MessageBusLiquidityBridgeUpdatedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "LiquidityBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusLiquidityBridgeUpdatedIterator{contract: _MessageBus.contract, event: "LiquidityBridgeUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidityBridgeUpdated is a free log subscription operation binding the contract event 0xbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae058.
//
// Solidity: event LiquidityBridgeUpdated(address liquidityBridge)
func (_MessageBus *MessageBusFilterer) WatchLiquidityBridgeUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusLiquidityBridgeUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "LiquidityBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusLiquidityBridgeUpdated)
				if err := _MessageBus.contract.UnpackLog(event, "LiquidityBridgeUpdated", log); err != nil {
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

// ParseLiquidityBridgeUpdated is a log parse operation binding the contract event 0xbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae058.
//
// Solidity: event LiquidityBridgeUpdated(address liquidityBridge)
func (_MessageBus *MessageBusFilterer) ParseLiquidityBridgeUpdated(log types.Log) (*MessageBusLiquidityBridgeUpdated, error) {
	event := new(MessageBusLiquidityBridgeUpdated)
	if err := _MessageBus.contract.UnpackLog(event, "LiquidityBridgeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusMessageIterator is returned from FilterMessage and is used to iterate over the raw logs and unpacked data for Message events raised by the MessageBus contract.
type MessageBusMessageIterator struct {
	Event *MessageBusMessage // Event containing the contract specifics and raw log

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
func (it *MessageBusMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusMessage)
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
		it.Event = new(MessageBusMessage)
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
func (it *MessageBusMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusMessage represents a Message event raised by the MessageBus contract.
type MessageBusMessage struct {
	Sender     common.Address
	Receiver   common.Address
	DstChainId *big.Int
	Message    []byte
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessage is a free log retrieval operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) FilterMessage(opts *bind.FilterOpts, sender []common.Address) (*MessageBusMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusMessageIterator{contract: _MessageBus.contract, event: "Message", logs: logs, sub: sub}, nil
}

// WatchMessage is a free log subscription operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) WatchMessage(opts *bind.WatchOpts, sink chan<- *MessageBusMessage, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusMessage)
				if err := _MessageBus.contract.UnpackLog(event, "Message", log); err != nil {
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

// ParseMessage is a log parse operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) ParseMessage(log types.Log) (*MessageBusMessage, error) {
	event := new(MessageBusMessage)
	if err := _MessageBus.contract.UnpackLog(event, "Message", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusMessage2Iterator is returned from FilterMessage2 and is used to iterate over the raw logs and unpacked data for Message2 events raised by the MessageBus contract.
type MessageBusMessage2Iterator struct {
	Event *MessageBusMessage2 // Event containing the contract specifics and raw log

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
func (it *MessageBusMessage2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusMessage2)
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
		it.Event = new(MessageBusMessage2)
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
func (it *MessageBusMessage2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusMessage2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusMessage2 represents a Message2 event raised by the MessageBus contract.
type MessageBusMessage2 struct {
	Sender     common.Address
	Receiver   []byte
	DstChainId *big.Int
	Message    []byte
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessage2 is a free log retrieval operation binding the contract event 0xe66fbe37d84ca73c589f782ac278844918ea6c56a4917f58707f715588080df2.
//
// Solidity: event Message2(address indexed sender, bytes receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) FilterMessage2(opts *bind.FilterOpts, sender []common.Address) (*MessageBusMessage2Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "Message2", senderRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusMessage2Iterator{contract: _MessageBus.contract, event: "Message2", logs: logs, sub: sub}, nil
}

// WatchMessage2 is a free log subscription operation binding the contract event 0xe66fbe37d84ca73c589f782ac278844918ea6c56a4917f58707f715588080df2.
//
// Solidity: event Message2(address indexed sender, bytes receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) WatchMessage2(opts *bind.WatchOpts, sink chan<- *MessageBusMessage2, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "Message2", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusMessage2)
				if err := _MessageBus.contract.UnpackLog(event, "Message2", log); err != nil {
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

// ParseMessage2 is a log parse operation binding the contract event 0xe66fbe37d84ca73c589f782ac278844918ea6c56a4917f58707f715588080df2.
//
// Solidity: event Message2(address indexed sender, bytes receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) ParseMessage2(log types.Log) (*MessageBusMessage2, error) {
	event := new(MessageBusMessage2)
	if err := _MessageBus.contract.UnpackLog(event, "Message2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusMessageWithTransferIterator is returned from FilterMessageWithTransfer and is used to iterate over the raw logs and unpacked data for MessageWithTransfer events raised by the MessageBus contract.
type MessageBusMessageWithTransferIterator struct {
	Event *MessageBusMessageWithTransfer // Event containing the contract specifics and raw log

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
func (it *MessageBusMessageWithTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusMessageWithTransfer)
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
		it.Event = new(MessageBusMessageWithTransfer)
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
func (it *MessageBusMessageWithTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusMessageWithTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusMessageWithTransfer represents a MessageWithTransfer event raised by the MessageBus contract.
type MessageBusMessageWithTransfer struct {
	Sender        common.Address
	Receiver      common.Address
	DstChainId    *big.Int
	Bridge        common.Address
	SrcTransferId [32]byte
	Message       []byte
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageWithTransfer is a free log retrieval operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) FilterMessageWithTransfer(opts *bind.FilterOpts, sender []common.Address) (*MessageBusMessageWithTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "MessageWithTransfer", senderRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusMessageWithTransferIterator{contract: _MessageBus.contract, event: "MessageWithTransfer", logs: logs, sub: sub}, nil
}

// WatchMessageWithTransfer is a free log subscription operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) WatchMessageWithTransfer(opts *bind.WatchOpts, sink chan<- *MessageBusMessageWithTransfer, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "MessageWithTransfer", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusMessageWithTransfer)
				if err := _MessageBus.contract.UnpackLog(event, "MessageWithTransfer", log); err != nil {
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

// ParseMessageWithTransfer is a log parse operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBus *MessageBusFilterer) ParseMessageWithTransfer(log types.Log) (*MessageBusMessageWithTransfer, error) {
	event := new(MessageBusMessageWithTransfer)
	if err := _MessageBus.contract.UnpackLog(event, "MessageWithTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusNeedRetryIterator is returned from FilterNeedRetry and is used to iterate over the raw logs and unpacked data for NeedRetry events raised by the MessageBus contract.
type MessageBusNeedRetryIterator struct {
	Event *MessageBusNeedRetry // Event containing the contract specifics and raw log

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
func (it *MessageBusNeedRetryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusNeedRetry)
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
		it.Event = new(MessageBusNeedRetry)
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
func (it *MessageBusNeedRetryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusNeedRetryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusNeedRetry represents a NeedRetry event raised by the MessageBus contract.
type MessageBusNeedRetry struct {
	MsgType    uint8
	MsgId      [32]byte
	SrcChainId uint64
	SrcTxHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNeedRetry is a free log retrieval operation binding the contract event 0xe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c.
//
// Solidity: event NeedRetry(uint8 msgType, bytes32 msgId, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBus *MessageBusFilterer) FilterNeedRetry(opts *bind.FilterOpts) (*MessageBusNeedRetryIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "NeedRetry")
	if err != nil {
		return nil, err
	}
	return &MessageBusNeedRetryIterator{contract: _MessageBus.contract, event: "NeedRetry", logs: logs, sub: sub}, nil
}

// WatchNeedRetry is a free log subscription operation binding the contract event 0xe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c.
//
// Solidity: event NeedRetry(uint8 msgType, bytes32 msgId, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBus *MessageBusFilterer) WatchNeedRetry(opts *bind.WatchOpts, sink chan<- *MessageBusNeedRetry) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "NeedRetry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusNeedRetry)
				if err := _MessageBus.contract.UnpackLog(event, "NeedRetry", log); err != nil {
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

// ParseNeedRetry is a log parse operation binding the contract event 0xe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c.
//
// Solidity: event NeedRetry(uint8 msgType, bytes32 msgId, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBus *MessageBusFilterer) ParseNeedRetry(log types.Log) (*MessageBusNeedRetry, error) {
	event := new(MessageBusNeedRetry)
	if err := _MessageBus.contract.UnpackLog(event, "NeedRetry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageBus contract.
type MessageBusOwnershipTransferredIterator struct {
	Event *MessageBusOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageBusOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusOwnershipTransferred)
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
		it.Event = new(MessageBusOwnershipTransferred)
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
func (it *MessageBusOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusOwnershipTransferred represents a OwnershipTransferred event raised by the MessageBus contract.
type MessageBusOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBus *MessageBusFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageBusOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusOwnershipTransferredIterator{contract: _MessageBus.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBus *MessageBusFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageBusOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusOwnershipTransferred)
				if err := _MessageBus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBus *MessageBusFilterer) ParseOwnershipTransferred(log types.Log) (*MessageBusOwnershipTransferred, error) {
	event := new(MessageBusOwnershipTransferred)
	if err := _MessageBus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusPegBridgeUpdatedIterator is returned from FilterPegBridgeUpdated and is used to iterate over the raw logs and unpacked data for PegBridgeUpdated events raised by the MessageBus contract.
type MessageBusPegBridgeUpdatedIterator struct {
	Event *MessageBusPegBridgeUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusPegBridgeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusPegBridgeUpdated)
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
		it.Event = new(MessageBusPegBridgeUpdated)
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
func (it *MessageBusPegBridgeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusPegBridgeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusPegBridgeUpdated represents a PegBridgeUpdated event raised by the MessageBus contract.
type MessageBusPegBridgeUpdated struct {
	PegBridge common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPegBridgeUpdated is a free log retrieval operation binding the contract event 0xd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d.
//
// Solidity: event PegBridgeUpdated(address pegBridge)
func (_MessageBus *MessageBusFilterer) FilterPegBridgeUpdated(opts *bind.FilterOpts) (*MessageBusPegBridgeUpdatedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "PegBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusPegBridgeUpdatedIterator{contract: _MessageBus.contract, event: "PegBridgeUpdated", logs: logs, sub: sub}, nil
}

// WatchPegBridgeUpdated is a free log subscription operation binding the contract event 0xd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d.
//
// Solidity: event PegBridgeUpdated(address pegBridge)
func (_MessageBus *MessageBusFilterer) WatchPegBridgeUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusPegBridgeUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "PegBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusPegBridgeUpdated)
				if err := _MessageBus.contract.UnpackLog(event, "PegBridgeUpdated", log); err != nil {
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

// ParsePegBridgeUpdated is a log parse operation binding the contract event 0xd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d.
//
// Solidity: event PegBridgeUpdated(address pegBridge)
func (_MessageBus *MessageBusFilterer) ParsePegBridgeUpdated(log types.Log) (*MessageBusPegBridgeUpdated, error) {
	event := new(MessageBusPegBridgeUpdated)
	if err := _MessageBus.contract.UnpackLog(event, "PegBridgeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusPegBridgeV2UpdatedIterator is returned from FilterPegBridgeV2Updated and is used to iterate over the raw logs and unpacked data for PegBridgeV2Updated events raised by the MessageBus contract.
type MessageBusPegBridgeV2UpdatedIterator struct {
	Event *MessageBusPegBridgeV2Updated // Event containing the contract specifics and raw log

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
func (it *MessageBusPegBridgeV2UpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusPegBridgeV2Updated)
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
		it.Event = new(MessageBusPegBridgeV2Updated)
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
func (it *MessageBusPegBridgeV2UpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusPegBridgeV2UpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusPegBridgeV2Updated represents a PegBridgeV2Updated event raised by the MessageBus contract.
type MessageBusPegBridgeV2Updated struct {
	PegBridgeV2 common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPegBridgeV2Updated is a free log retrieval operation binding the contract event 0xfb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e8.
//
// Solidity: event PegBridgeV2Updated(address pegBridgeV2)
func (_MessageBus *MessageBusFilterer) FilterPegBridgeV2Updated(opts *bind.FilterOpts) (*MessageBusPegBridgeV2UpdatedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "PegBridgeV2Updated")
	if err != nil {
		return nil, err
	}
	return &MessageBusPegBridgeV2UpdatedIterator{contract: _MessageBus.contract, event: "PegBridgeV2Updated", logs: logs, sub: sub}, nil
}

// WatchPegBridgeV2Updated is a free log subscription operation binding the contract event 0xfb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e8.
//
// Solidity: event PegBridgeV2Updated(address pegBridgeV2)
func (_MessageBus *MessageBusFilterer) WatchPegBridgeV2Updated(opts *bind.WatchOpts, sink chan<- *MessageBusPegBridgeV2Updated) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "PegBridgeV2Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusPegBridgeV2Updated)
				if err := _MessageBus.contract.UnpackLog(event, "PegBridgeV2Updated", log); err != nil {
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

// ParsePegBridgeV2Updated is a log parse operation binding the contract event 0xfb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e8.
//
// Solidity: event PegBridgeV2Updated(address pegBridgeV2)
func (_MessageBus *MessageBusFilterer) ParsePegBridgeV2Updated(log types.Log) (*MessageBusPegBridgeV2Updated, error) {
	event := new(MessageBusPegBridgeV2Updated)
	if err := _MessageBus.contract.UnpackLog(event, "PegBridgeV2Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusPegVaultUpdatedIterator is returned from FilterPegVaultUpdated and is used to iterate over the raw logs and unpacked data for PegVaultUpdated events raised by the MessageBus contract.
type MessageBusPegVaultUpdatedIterator struct {
	Event *MessageBusPegVaultUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusPegVaultUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusPegVaultUpdated)
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
		it.Event = new(MessageBusPegVaultUpdated)
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
func (it *MessageBusPegVaultUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusPegVaultUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusPegVaultUpdated represents a PegVaultUpdated event raised by the MessageBus contract.
type MessageBusPegVaultUpdated struct {
	PegVault common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPegVaultUpdated is a free log retrieval operation binding the contract event 0xa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad.
//
// Solidity: event PegVaultUpdated(address pegVault)
func (_MessageBus *MessageBusFilterer) FilterPegVaultUpdated(opts *bind.FilterOpts) (*MessageBusPegVaultUpdatedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "PegVaultUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusPegVaultUpdatedIterator{contract: _MessageBus.contract, event: "PegVaultUpdated", logs: logs, sub: sub}, nil
}

// WatchPegVaultUpdated is a free log subscription operation binding the contract event 0xa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad.
//
// Solidity: event PegVaultUpdated(address pegVault)
func (_MessageBus *MessageBusFilterer) WatchPegVaultUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusPegVaultUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "PegVaultUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusPegVaultUpdated)
				if err := _MessageBus.contract.UnpackLog(event, "PegVaultUpdated", log); err != nil {
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

// ParsePegVaultUpdated is a log parse operation binding the contract event 0xa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad.
//
// Solidity: event PegVaultUpdated(address pegVault)
func (_MessageBus *MessageBusFilterer) ParsePegVaultUpdated(log types.Log) (*MessageBusPegVaultUpdated, error) {
	event := new(MessageBusPegVaultUpdated)
	if err := _MessageBus.contract.UnpackLog(event, "PegVaultUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusPegVaultV2UpdatedIterator is returned from FilterPegVaultV2Updated and is used to iterate over the raw logs and unpacked data for PegVaultV2Updated events raised by the MessageBus contract.
type MessageBusPegVaultV2UpdatedIterator struct {
	Event *MessageBusPegVaultV2Updated // Event containing the contract specifics and raw log

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
func (it *MessageBusPegVaultV2UpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusPegVaultV2Updated)
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
		it.Event = new(MessageBusPegVaultV2Updated)
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
func (it *MessageBusPegVaultV2UpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusPegVaultV2UpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusPegVaultV2Updated represents a PegVaultV2Updated event raised by the MessageBus contract.
type MessageBusPegVaultV2Updated struct {
	PegVaultV2 common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPegVaultV2Updated is a free log retrieval operation binding the contract event 0x918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb5.
//
// Solidity: event PegVaultV2Updated(address pegVaultV2)
func (_MessageBus *MessageBusFilterer) FilterPegVaultV2Updated(opts *bind.FilterOpts) (*MessageBusPegVaultV2UpdatedIterator, error) {

	logs, sub, err := _MessageBus.contract.FilterLogs(opts, "PegVaultV2Updated")
	if err != nil {
		return nil, err
	}
	return &MessageBusPegVaultV2UpdatedIterator{contract: _MessageBus.contract, event: "PegVaultV2Updated", logs: logs, sub: sub}, nil
}

// WatchPegVaultV2Updated is a free log subscription operation binding the contract event 0x918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb5.
//
// Solidity: event PegVaultV2Updated(address pegVaultV2)
func (_MessageBus *MessageBusFilterer) WatchPegVaultV2Updated(opts *bind.WatchOpts, sink chan<- *MessageBusPegVaultV2Updated) (event.Subscription, error) {

	logs, sub, err := _MessageBus.contract.WatchLogs(opts, "PegVaultV2Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusPegVaultV2Updated)
				if err := _MessageBus.contract.UnpackLog(event, "PegVaultV2Updated", log); err != nil {
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

// ParsePegVaultV2Updated is a log parse operation binding the contract event 0x918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb5.
//
// Solidity: event PegVaultV2Updated(address pegVaultV2)
func (_MessageBus *MessageBusFilterer) ParsePegVaultV2Updated(log types.Log) (*MessageBusPegVaultV2Updated, error) {
	event := new(MessageBusPegVaultV2Updated)
	if err := _MessageBus.contract.UnpackLog(event, "PegVaultV2Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverMetaData contains all meta data concerning the MessageBusReceiver contract.
var MessageBusReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_liquidityBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegBridgeV2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pegVaultV2\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"CallReverted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMsgDataTypes.MsgType\",\"name\":\"msgType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumMsgDataTypes.TxStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"name\":\"Executed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidityBridge\",\"type\":\"address\"}],\"name\":\"LiquidityBridgeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMsgDataTypes.MsgType\",\"name\":\"msgType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"msgId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"name\":\"NeedRetry\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegBridge\",\"type\":\"address\"}],\"name\":\"PegBridgeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegBridgeV2\",\"type\":\"address\"}],\"name\":\"PegBridgeV2Updated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegVault\",\"type\":\"address\"}],\"name\":\"PegVaultUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pegVaultV2\",\"type\":\"address\"}],\"name\":\"PegVaultV2Updated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.RouteInfo\",\"name\":\"_route\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"sender\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.RouteInfo2\",\"name\":\"_route\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"_transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"executedMessages\",\"outputs\":[{\"internalType\":\"enumMsgDataTypes.TxStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidityBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegBridgeV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pegVaultV2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preExecuteMessageGasUsage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.BridgeTransferParams\",\"name\":\"_transferParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.MsgWithTransferExecutionParams\",\"name\":\"_msgParams\",\"type\":\"tuple\"}],\"name\":\"refundAndExecuteMsg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setLiquidityBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegBridgeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setPegVaultV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_usage\",\"type\":\"uint256\"}],\"name\":\"setPreExecuteMessageGasUsage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.BridgeTransferParams\",\"name\":\"_transferParams\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"enumMsgDataTypes.TransferType\",\"name\":\"t\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"wdseq\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"refId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"srcTxHash\",\"type\":\"bytes32\"}],\"internalType\":\"structMsgDataTypes.TransferInfo\",\"name\":\"transfer\",\"type\":\"tuple\"},{\"internalType\":\"bytes[]\",\"name\":\"sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"}],\"internalType\":\"structMsgDataTypes.MsgWithTransferExecutionParams\",\"name\":\"_msgParams\",\"type\":\"tuple\"}],\"name\":\"transferAndExecuteMsg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620034793803806200347983398101604081905262000034916200010f565b6200003f33620000a2565b600280546001600160a01b03199081166001600160a01b039788161790915560038054821695871695909517909455600480548516938616939093179092556005805484169185169190911790556006805490921692169190911790556200017f565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b03811681146200010a57600080fd5b919050565b600080600080600060a086880312156200012857600080fd5b6200013386620000f2565b94506200014360208701620000f2565b93506200015360408701620000f2565b92506200016360608701620000f2565b91506200017360808701620000f2565b90509295509295909350565b6132ea806200018f6000396000f3fe60806040526004361061016a5760003560e01c806382efd502116100cb578063cd2abd661161007f578063dfa2dbaf11610059578063dfa2dbaf146103b9578063f2fde38b146103d9578063f83b0fb9146103f957600080fd5b8063cd2abd6614610349578063d8257d1714610386578063db2c20c8146103a657600080fd5b806395b12c27116100b057806395b12c27146102e95780639b05a77514610309578063c66a9c5a1461032957600080fd5b806382efd502146102ab5780638da5cb5b146102cb57600080fd5b8063584e45e111610122578063723d0a9d11610107578063723d0a9d146102405780637b80ab201461026057806382980dc41461027357600080fd5b8063584e45e1146101f7578063588be02b1461022057600080fd5b806340d0d0261161015357806340d0d026146101a45780634586f331146101c4578063468a2d04146101e457600080fd5b806303cbfe661461016f5780633f395aff14610191575b600080fd5b34801561017b57600080fd5b5061018f61018a36600461269b565b610419565b005b61018f61019f366004612744565b610524565b3480156101b057600080fd5b5061018f6101bf36600461284d565b610816565b3480156101d057600080fd5b5061018f6101df3660046128b9565b61086e565b61018f6101f23660046128d2565b6108dc565b34801561020357600080fd5b5061020d60075481565b6040519081526020015b60405180910390f35b34801561022c57600080fd5b5061018f61023b36600461269b565b61093c565b34801561024c57600080fd5b5061018f61025b36600461284d565b610a3b565b61018f61026e366004612744565b610a89565b34801561027f57600080fd5b50600254610293906001600160a01b031681565b6040516001600160a01b039091168152602001610217565b3480156102b757600080fd5b5061018f6102c636600461269b565b610c9a565b3480156102d757600080fd5b506000546001600160a01b0316610293565b3480156102f557600080fd5b50600554610293906001600160a01b031681565b34801561031557600080fd5b5061018f61032436600461269b565b610d99565b34801561033557600080fd5b50600654610293906001600160a01b031681565b34801561035557600080fd5b506103796103643660046128b9565b60016020526000908152604090205460ff1681565b60405161021791906129c1565b34801561039257600080fd5b50600454610293906001600160a01b031681565b61018f6103b43660046129cf565b610e98565b3480156103c557600080fd5b50600354610293906001600160a01b031681565b3480156103e557600080fd5b5061018f6103f436600461269b565b610eec565b34801561040557600080fd5b5061018f61041436600461269b565b610fdd565b3361042c6000546001600160a01b031690565b6001600160a01b0316146104875760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b6001600160a01b0381166104cf5760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b604482015260640161047e565b600380546001600160a01b0319166001600160a01b0383169081179091556040519081527fd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d906020015b60405180910390a150565b600061052f886110dc565b90506000808281526001602052604090205460ff16600481111561055557610555612997565b146105a25760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c726561647920657865637574656400000000000000604482015260640161047e565b6000818152600160209081526040808320805460ff19166004179055805146928101929092526bffffffffffffffffffffffff193060601b16908201527f4d657373616765576974685472616e73666572000000000000000000000000006054820152606701604051602081830303815290604052805190602001209050600260009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e8e6101000135604051602001610662959493929190612a96565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b81526004016106999796959493929190612c55565b60006040518083038186803b1580156106b157600080fd5b505afa1580156106c5573d6000803e3d6000fd5b505050506000806106d78b8e8e611957565b905060018160028111156106ed576106ed612997565b036106fb57600191506107c8565b600281600281111561070f5761070f612997565b0361079457600084815260016020819052604082208054909160ff1990911690835b02179055507fe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c6000858d60c001602081019061076d9190612cb3565b8e61010001356040516107839493929190612ced565b60405180910390a15050505061080b565b61079f8b8e8e611a92565b905060018160028111156107b5576107b5612997565b036107c357600391506107c8565b600291505b60008481526001602081905260409091208054849260ff19909116908360048111156107f6576107f6612997565b021790555061080684838d611acd565b505050505b505050505050505050565b61082f6108296040830160208401612d20565b83611b3f565b61086a61083c8280612d41565b6020840161084e610140860186612d88565b61085c610160880188612d88565b61026e6101808a018a612d88565b5050565b336108816000546001600160a01b031690565b6001600160a01b0316146108d75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161047e565b600755565b60006108e788611d89565b90506109308a8a838a8a8a8a8a8a6040518060400160405280600781526020017f4d65737361676500000000000000000000000000000000000000000000000000815250611e40565b50505050505050505050565b3361094f6000546001600160a01b031690565b6001600160a01b0316146109a55760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161047e565b6001600160a01b0381166109ed5760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b604482015260640161047e565b600280546001600160a01b0319166001600160a01b0383169081179091556040519081527fbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae05890602001610519565b610a4e6108296040830160208401612d20565b61086a610a5b8280612d41565b60208401610a6d610140860186612d88565b610a7b610160880188612d88565b61019f6101808a018a612d88565b6000610a94886110dc565b90506000808281526001602052604090205460ff166004811115610aba57610aba612997565b14610b075760405162461bcd60e51b815260206004820152601960248201527f7472616e7366657220616c726561647920657865637574656400000000000000604482015260640161047e565b6000818152600160209081526040808320805460ff19166004179055805146928101929092526bffffffffffffffffffffffff193060601b16908201527f4d657373616765576974685472616e73666572526566756e64000000000000006054820152606d01604051602081830303815290604052805190602001209050600260009054906101000a90046001600160a01b03166001600160a01b031663682dbc2282848e8e8e6101000135604051602001610bc7959493929190612a96565b6040516020818303038152906040528a8a8a8a8a8a6040518863ffffffff1660e01b8152600401610bfe9796959493929190612c55565b60006040518083038186803b158015610c1657600080fd5b505afa158015610c2a573d6000803e3d6000fd5b50505050600080610c3c8b8e8e612090565b90506001816002811115610c5257610c52612997565b03610c6057600191506107c8565b6002816002811115610c7457610c74612997565b036107c357600084815260016020819052604082208054909160ff199091169083610731565b33610cad6000546001600160a01b031690565b6001600160a01b031614610d035760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161047e565b6001600160a01b038116610d4b5760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b604482015260640161047e565b600580546001600160a01b0319166001600160a01b0383169081179091556040519081527ffb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e890602001610519565b33610dac6000546001600160a01b031690565b6001600160a01b031614610e025760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161047e565b6001600160a01b038116610e4a5760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b604482015260640161047e565b600480546001600160a01b0319166001600160a01b0383169081179091556040519081527fa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad90602001610519565b6000610ea3886120e7565b90506109308a8a838a8a8a8a8a8a6040518060400160405280600881526020017f4d65737361676532000000000000000000000000000000000000000000000000815250611e40565b33610eff6000546001600160a01b031690565b6001600160a01b031614610f555760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161047e565b6001600160a01b038116610fd15760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161047e565b610fda8161217d565b50565b33610ff06000546001600160a01b031690565b6001600160a01b0316146110465760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161047e565b6001600160a01b03811661108e5760405162461bcd60e51b815260206004820152600f60248201526e696e76616c6964206164647265737360881b604482015260640161047e565b600680546001600160a01b0319166001600160a01b0383169081179091556040519081527f918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb590602001610519565b6000808060016110ef6020860186612d20565b600681111561110057611100612997565b0361129057611115604085016020860161269b565b611125606086016040870161269b565b611135608087016060880161269b565b608087013561114a60e0890160c08a01612cb3565b6040516bffffffffffffffffffffffff19606096871b8116602083015294861b851660348201529290941b9092166048820152605c8101919091526001600160c01b031960c092831b8116607c8301524690921b909116608482015260e0850135608c82015260ac0160408051808303601f19018152908290528051602090910120600254633c64f04b60e01b8352600483018290529093506001600160a01b031691508190633c64f04b90602401602060405180830381865afa158015611216573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061123a9190612dd2565b151560011461128b5760405162461bcd60e51b815260206004820152601660248201527f6272696467652072656c6179206e6f7420657869737400000000000000000000604482015260640161047e565b611922565b600261129f6020860186612d20565b60068111156112b0576112b0612997565b0361141257466112c660c0860160a08701612cb3565b6112d6606087016040880161269b565b6112e6608088016060890161269b565b6040516001600160c01b031960c095861b811660208301529390941b90921660288401526bffffffffffffffffffffffff19606091821b8116603085015291901b1660448201526080850135605882015260780160408051808303601f19018152908290528051602090910120600254631c13568560e31b8352600483018290529093506001600160a01b03169150819063e09ab42890602401602060405180830381865afa15801561139d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113c19190612dd2565b151560011461128b5760405162461bcd60e51b815260206004820152601960248201527f627269646765207769746864726177206e6f7420657869737400000000000000604482015260640161047e565b60036114216020860186612d20565b600681111561143257611432612997565b148061145b575060046114486020860186612d20565b600681111561145957611459612997565b145b156116c157611470606085016040860161269b565b611480608086016060870161269b565b6080860135611495604088016020890161269b565b6114a560e0890160c08a01612cb3565b604051606095861b6bffffffffffffffffffffffff19908116602083015294861b851660348201526048810193909352931b909116606882015260c09190911b6001600160c01b031916607c82015260e0850135608482015260a40160408051601f1981840301815291905280516020909101209150600361152a6020860186612d20565b600681111561153b5761153b612997565b0361160257506003546040516301e6472560e01b8152600481018390526001600160a01b039091169081906301e64725906024015b602060405180830381865afa15801561158d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115b19190612dd2565b151560011461128b5760405162461bcd60e51b815260206004820152601560248201527f6d696e74207265636f7264206e6f742065786973740000000000000000000000604482015260640161047e565b50600480546040516301e6472560e01b81529182018390526001600160a01b03169081906301e6472590602401602060405180830381865afa15801561164c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116709190612dd2565b151560011461128b5760405162461bcd60e51b815260206004820152601960248201527f7769746864726177207265636f7264206e6f7420657869737400000000000000604482015260640161047e565b60056116d06020860186612d20565b60068111156116e1576116e1612997565b148061170a575060066116f76020860186612d20565b600681111561170857611708612997565b145b1561192257600561171e6020860186612d20565b600681111561172f5761172f612997565b0361174657506005546001600160a01b0316611754565b506006546001600160a01b03165b611764606085016040860161269b565b611774608086016060870161269b565b6080860135611789604088016020890161269b565b61179960e0890160c08a01612cb3565b604051606095861b6bffffffffffffffffffffffff19908116602083015294861b85166034820152604881019390935290841b8316606883015260c01b6001600160c01b031916607c82015260e087013560848201529183901b1660a482015260b80160408051601f198184030181529190528051602090910120915060056118256020860186612d20565b600681111561183657611836612997565b03611868576040516301e6472560e01b8152600481018390526001600160a01b038216906301e6472590602401611570565b6040516301e6472560e01b8152600481018390526001600160a01b038216906301e6472590602401602060405180830381865afa1580156118ad573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118d19190612dd2565b15156001146119225760405162461bcd60e51b815260206004820152601960248201527f7769746864726177207265636f7264206e6f7420657869737400000000000000604482015260640161047e565b6000818360405160200161193893929190612e0b565b6040516020818303038152906040528051906020012092505050919050565b6000805a9050600080611970606088016040890161269b565b6001600160a01b031634631f34afff60e21b61199260408b0160208c0161269b565b6119a260808c0160608d0161269b565b60808c01356119b760e08e0160c08f01612cb3565b8c8c336040516024016119d09796959493929190612e3c565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b0319909416939093179092529051611a0e9190612e95565b60006040518083038185875af1925050503d8060008114611a4b576040519150601f19603f3d011682016040523d82523d6000602084013e611a50565b606091505b50915091508115611a795780806020019051810190611a6f9190612eb1565b9350505050611a8b565b611a8383826121cd565b600093505050505b9392505050565b6000805a9050600080611aab606088016040890161269b565b6001600160a01b031634632d5bd7e360e11b61199260408b0160208c0161269b565b611add606082016040830161269b565b6001600160a01b03167fa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d60008585611b1b60e0870160c08801612cb3565b866101000135604051611b32959493929190612ed2565b60405180910390a2505050565b6001826006811115611b5357611b53612997565b03611bf3576002546001600160a01b031663cdd1b25d611b738380612d41565b611b806020860186612d88565b611b8d6040880188612d88565b611b9a60608a018a612d88565b6040518963ffffffff1660e01b8152600401611bbd989796959493929190612f10565b600060405180830381600087803b158015611bd757600080fd5b505af1158015611beb573d6000803e3d6000fd5b505050505050565b6002826006811115611c0757611c07612997565b03611c27576002546001600160a01b031663a21a9280611b738380612d41565b6003826006811115611c3b57611c3b612997565b03611c5b576003546001600160a01b031663f8734302611b738380612d41565b6005826006811115611c6f57611c6f612997565b03611d21576005546001600160a01b031663f8734302611c8f8380612d41565b611c9c6020860186612d88565b611ca96040880188612d88565b611cb660608a018a612d88565b6040518963ffffffff1660e01b8152600401611cd9989796959493929190612f10565b6020604051808303816000875af1158015611cf8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d1c9190612f70565b505050565b6004826006811115611d3557611d35612997565b03611d55576004546001600160a01b031663a21a9280611b738380612d41565b6006826006811115611d6957611d69612997565b0361086a576006546001600160a01b031663a21a9280611c8f8380612d41565b6040805160a081018252600080825260606020830181905292820181905291810182905260808101919091526040805160a0810190915280611dce602085018561269b565b6001600160a01b03168152602001604051806020016040528060008152508152602001836020016020810190611e04919061269b565b6001600160a01b03168152602001611e226060850160408601612cb3565b67ffffffffffffffff16815260200183606001358152509050919050565b6000611e4d898c8c612258565b90506000808281526001602052604090205460ff166004811115611e7357611e73612997565b14611ec05760405162461bcd60e51b815260206004820152601860248201527f6d65737361676520616c72656164792065786563757465640000000000000000604482015260640161047e565b6000818152600160209081526040808320805460ff1916600417905551611eed9146913091879101612f89565b60408051601f1981840301815282825280516020918201206002549184018190528383018690528251808503840181526060850193849052633416de1160e11b90935293506001600160a01b03169163682dbc2291611f5a918d908d908d908d908d908d90606401612c55565b60006040518083038186803b158015611f7257600080fd5b505afa158015611f86573d6000803e3d6000fd5b50505050600080611f988c8f8f6122f1565b90506001816002811115611fae57611fae612997565b03611fbc5760019150612042565b6002816002811115611fd057611fd0612997565b0361203d57600084815260016020819052604091829020805460ff1916905560608e015160808f015192517fe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c9361202c93928992909190612ced565b60405180910390a150505050610930565b600291505b60008481526001602081905260409091208054849260ff199091169083600481111561207057612070612997565b021790555061208084838e6124d2565b5050505050505050505050505050565b6000805a90506000806120a9606088016040890161269b565b6001600160a01b0316346305e5a4c160e11b6120cb60808b0160608c0161269b565b8a608001358a8a336040516024016119d0959493929190612fc8565b6040805160a081018252600080825260606020830181905292820181905291810182905260808101919091526040805160a0810190915260008152602081016121308480612d41565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250602090810190611e04906040860190860161269b565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60005a90506000600754456121e2919061301d565b905080841080156121fd57506121f9604085613030565b8211155b1561220457fe5b600061220f84612520565b905061221a81612585565b7fffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f816040516122499190613052565b60405180910390a15050505050565b602083015180516000919082036122a5578451604051612293919060200160609190911b6bffffffffffffffffffffffff1916815260140190565b60405160208183030381529060405290505b6001818660400151876060015188608001514689896040516020016122d1989796959493929190613065565b604051602081830303815290604052805190602001209150509392505050565b6000805a9050600060608660200151516000036123e15786604001516001600160a01b0316346040518060600160405280602c8152602001613289602c91398051602090910120895160608b01516040516123569291908c908c9033906024016130e6565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b03199094169390931790925290516123949190612e95565b60006040518083038185875af1925050503d80600081146123d1576040519150601f19603f3d011682016040523d82523d6000602084013e6123d6565b606091505b5090925090506124b8565b86604001516001600160a01b0316346040518060600160405280602a815260200161325f602a91398051906020012089602001518a606001518a8a3360405160240161243195949392919061311a565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b031990941693909317909252905161246f9190612e95565b60006040518083038185875af1925050503d80600081146124ac576040519150601f19603f3d011682016040523d82523d6000602084013e6124b1565b606091505b5090925090505b8115611a795780806020019051810190611a6f9190612eb1565b80604001516001600160a01b03167fa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d6001858585606001518660800151604051611b32959493929190612ed2565b606060448251101561256557505060408051808201909152601d81527f5472616e73616374696f6e2072657665727465642073696c656e746c79000000602082015290565b6004820191508180602001905181019061257f9190613182565b92915050565b60408051808201909152600b8082527f4d53473a3a41424f52543a00000000000000000000000000000000000000000060208301528251839111611d1c5760005b8251811015612664578281815181106125e1576125e161322f565b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168282815181106126205761262061322f565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016146126525750505050565b8061265c81613245565b9150506125c6565b508260405162461bcd60e51b815260040161047e9190613052565b80356001600160a01b038116811461269657600080fd5b919050565b6000602082840312156126ad57600080fd5b611a8b8261267f565b60008083601f8401126126c857600080fd5b50813567ffffffffffffffff8111156126e057600080fd5b6020830191508360208285010111156126f857600080fd5b9250929050565b60008083601f84011261271157600080fd5b50813567ffffffffffffffff81111561272957600080fd5b6020830191508360208260051b85010111156126f857600080fd5b6000806000806000806000806000898b036101a081121561276457600080fd5b8a3567ffffffffffffffff8082111561277c57600080fd5b6127888e838f016126b6565b909c509a508a9150610120601f19840112156127a357600080fd5b60208d0199506101408d01359250808311156127be57600080fd5b6127ca8e848f016126ff565b90995097506101608d01359250889150808311156127e757600080fd5b6127f38e848f016126ff565b90975095506101808d013592508691508083111561281057600080fd5b505061281e8c828d016126ff565b915080935050809150509295985092959850929598565b60006080828403121561284757600080fd5b50919050565b6000806040838503121561286057600080fd5b823567ffffffffffffffff8082111561287857600080fd5b61288486838701612835565b9350602085013591508082111561289a57600080fd5b5083016101a081860312156128ae57600080fd5b809150509250929050565b6000602082840312156128cb57600080fd5b5035919050565b60008060008060008060008060006101008a8c0312156128f157600080fd5b893567ffffffffffffffff8082111561290957600080fd5b6129158d838e016126b6565b909b50995089915061292a8d60208e01612835565b985060a08c013591508082111561294057600080fd5b61294c8d838e016126ff565b909850965060c08c013591508082111561296557600080fd5b6129718d838e016126ff565b909650945060e08c013591508082111561298a57600080fd5b5061281e8c828d016126ff565b634e487b7160e01b600052602160045260246000fd5b600581106129bd576129bd612997565b9052565b6020810161257f82846129ad565b600080600080600080600080600060a08a8c0312156129ed57600080fd5b893567ffffffffffffffff80821115612a0557600080fd5b612a118d838e016126b6565b909b50995060208c0135915080821115612a2a57600080fd5b612a368d838e01612835565b985060408c0135915080821115612a4c57600080fd5b612a588d838e016126ff565b909850965060608c0135915080821115612a7157600080fd5b612a7d8d838e016126ff565b909650945060808c013591508082111561298a57600080fd5b8581528460208201528284604083013760409201918201526060019392505050565b60005b83811015612ad3578181015183820152602001612abb565b50506000910152565b60008151808452612af4816020860160208601612ab8565b601f01601f19169290920160200192915050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b81835260006020808501808196508560051b810191508460005b87811015612bb65782840389528135601e19883603018112612b6c57600080fd5b8701858101903567ffffffffffffffff811115612b8857600080fd5b803603821315612b9757600080fd5b612ba2868284612b08565b9a87019a9550505090840190600101612b4b565b5091979650505050505050565b8183526000602080850194508260005b85811015612bff576001600160a01b03612bec8361267f565b1687529582019590820190600101612bd3565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115612c3c57600080fd5b8260051b80836020870137939093016020019392505050565b608081526000612c68608083018a612adc565b8281036020840152612c7b81898b612b31565b90508281036040840152612c90818789612bc3565b90508281036060840152612ca5818587612c0a565b9a9950505050505050505050565b600060208284031215612cc557600080fd5b813567ffffffffffffffff81168114611a8b57600080fd5b600281106129bd576129bd612997565b60808101612cfb8287612cdd565b84602083015267ffffffffffffffff8416604083015282606083015295945050505050565b600060208284031215612d3257600080fd5b813560078110611a8b57600080fd5b6000808335601e19843603018112612d5857600080fd5b83018035915067ffffffffffffffff821115612d7357600080fd5b6020019150368190038213156126f857600080fd5b6000808335601e19843603018112612d9f57600080fd5b83018035915067ffffffffffffffff821115612dba57600080fd5b6020019150600581901b36038213156126f857600080fd5b600060208284031215612de457600080fd5b81518015158114611a8b57600080fd5b60028110612e0457612e04612997565b60f81b9052565b612e158185612df4565b60609290921b6bffffffffffffffffffffffff191660018301526015820152603501919050565b60006001600160a01b03808a168352808916602084015287604084015267ffffffffffffffff8716606084015260c06080840152612e7e60c084018688612b08565b915080841660a08401525098975050505050505050565b60008251612ea7818460208701612ab8565b9190910192915050565b600060208284031215612ec357600080fd5b815160038110611a8b57600080fd5b60a08101612ee08288612cdd565b856020830152612ef360408301866129ad565b67ffffffffffffffff939093166060820152608001529392505050565b608081526000612f24608083018a8c612b08565b8281036020840152612f3781898b612b31565b90508281036040840152612f4c818789612bc3565b90508281036060840152612f61818587612c0a565b9b9a5050505050505050505050565b600060208284031215612f8257600080fd5b5051919050565b8381526bffffffffffffffffffffffff198360601b16602082015260008251612fb9816034850160208701612ab8565b91909101603401949350505050565b60006001600160a01b03808816835286602084015260806040840152612ff2608084018688612b08565b91508084166060840152509695505050505050565b634e487b7160e01b600052601160045260246000fd5b8181038181111561257f5761257f613007565b60008261304d57634e487b7160e01b600052601260045260246000fd5b500490565b602081526000611a8b6020830184612adc565b61306f818a612df4565b60008851613084816001850160208d01612ab8565b80830190506bffffffffffffffffffffffff198960601b1660018201526001600160c01b0319808960c01b16601583015287601d830152808760c01b16603d830152508385604583013760009301604501928352509098975050505050505050565b60006001600160a01b03808816835267ffffffffffffffff8716602084015260806040840152612ff2608084018688612b08565b60808152600061312d6080830188612adc565b67ffffffffffffffff871660208401528281036040840152613150818688612b08565b9150506001600160a01b03831660608301529695505050505050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561319457600080fd5b815167ffffffffffffffff808211156131ac57600080fd5b818401915084601f8301126131c057600080fd5b8151818111156131d2576131d261316c565b604051601f8201601f19908116603f011681019083821181831017156131fa576131fa61316c565b8160405282815287602084870101111561321357600080fd5b613224836020830160208801612ab8565b979650505050505050565b634e487b7160e01b600052603260045260246000fd5b60006001820161325757613257613007565b506001019056fe657865637574654d6573736167652862797465732c75696e7436342c62797465732c6164647265737329657865637574654d65737361676528616464726573732c75696e7436342c62797465732c6164647265737329a26469706673582212203967148941ba46a878a1579e21647777b61193f7f5e63f1ae85998ea4d7c78b864736f6c63430008110033",
}

// MessageBusReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageBusReceiverMetaData.ABI instead.
var MessageBusReceiverABI = MessageBusReceiverMetaData.ABI

// MessageBusReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageBusReceiverMetaData.Bin instead.
var MessageBusReceiverBin = MessageBusReceiverMetaData.Bin

// DeployMessageBusReceiver deploys a new Ethereum contract, binding an instance of MessageBusReceiver to it.
func DeployMessageBusReceiver(auth *bind.TransactOpts, backend bind.ContractBackend, _liquidityBridge common.Address, _pegBridge common.Address, _pegVault common.Address, _pegBridgeV2 common.Address, _pegVaultV2 common.Address) (common.Address, *types.Transaction, *MessageBusReceiver, error) {
	parsed, err := MessageBusReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBusReceiverBin), backend, _liquidityBridge, _pegBridge, _pegVault, _pegBridgeV2, _pegVaultV2)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageBusReceiver{MessageBusReceiverCaller: MessageBusReceiverCaller{contract: contract}, MessageBusReceiverTransactor: MessageBusReceiverTransactor{contract: contract}, MessageBusReceiverFilterer: MessageBusReceiverFilterer{contract: contract}}, nil
}

// MessageBusReceiver is an auto generated Go binding around an Ethereum contract.
type MessageBusReceiver struct {
	MessageBusReceiverCaller     // Read-only binding to the contract
	MessageBusReceiverTransactor // Write-only binding to the contract
	MessageBusReceiverFilterer   // Log filterer for contract events
}

// MessageBusReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageBusReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageBusReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageBusReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageBusReceiverSession struct {
	Contract     *MessageBusReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MessageBusReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageBusReceiverCallerSession struct {
	Contract *MessageBusReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MessageBusReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageBusReceiverTransactorSession struct {
	Contract     *MessageBusReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MessageBusReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageBusReceiverRaw struct {
	Contract *MessageBusReceiver // Generic contract binding to access the raw methods on
}

// MessageBusReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageBusReceiverCallerRaw struct {
	Contract *MessageBusReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// MessageBusReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageBusReceiverTransactorRaw struct {
	Contract *MessageBusReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageBusReceiver creates a new instance of MessageBusReceiver, bound to a specific deployed contract.
func NewMessageBusReceiver(address common.Address, backend bind.ContractBackend) (*MessageBusReceiver, error) {
	contract, err := bindMessageBusReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiver{MessageBusReceiverCaller: MessageBusReceiverCaller{contract: contract}, MessageBusReceiverTransactor: MessageBusReceiverTransactor{contract: contract}, MessageBusReceiverFilterer: MessageBusReceiverFilterer{contract: contract}}, nil
}

// NewMessageBusReceiverCaller creates a new read-only instance of MessageBusReceiver, bound to a specific deployed contract.
func NewMessageBusReceiverCaller(address common.Address, caller bind.ContractCaller) (*MessageBusReceiverCaller, error) {
	contract, err := bindMessageBusReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverCaller{contract: contract}, nil
}

// NewMessageBusReceiverTransactor creates a new write-only instance of MessageBusReceiver, bound to a specific deployed contract.
func NewMessageBusReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageBusReceiverTransactor, error) {
	contract, err := bindMessageBusReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverTransactor{contract: contract}, nil
}

// NewMessageBusReceiverFilterer creates a new log filterer instance of MessageBusReceiver, bound to a specific deployed contract.
func NewMessageBusReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageBusReceiverFilterer, error) {
	contract, err := bindMessageBusReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverFilterer{contract: contract}, nil
}

// bindMessageBusReceiver binds a generic wrapper to an already deployed contract.
func bindMessageBusReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageBusReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusReceiver *MessageBusReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusReceiver.Contract.MessageBusReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusReceiver *MessageBusReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.MessageBusReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusReceiver *MessageBusReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.MessageBusReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusReceiver *MessageBusReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusReceiver *MessageBusReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusReceiver *MessageBusReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.contract.Transact(opts, method, params...)
}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBusReceiver *MessageBusReceiverCaller) ExecutedMessages(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "executedMessages", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBusReceiver *MessageBusReceiverSession) ExecutedMessages(arg0 [32]byte) (uint8, error) {
	return _MessageBusReceiver.Contract.ExecutedMessages(&_MessageBusReceiver.CallOpts, arg0)
}

// ExecutedMessages is a free data retrieval call binding the contract method 0xcd2abd66.
//
// Solidity: function executedMessages(bytes32 ) view returns(uint8)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) ExecutedMessages(arg0 [32]byte) (uint8, error) {
	return _MessageBusReceiver.Contract.ExecutedMessages(&_MessageBusReceiver.CallOpts, arg0)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) LiquidityBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "liquidityBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) LiquidityBridge() (common.Address, error) {
	return _MessageBusReceiver.Contract.LiquidityBridge(&_MessageBusReceiver.CallOpts)
}

// LiquidityBridge is a free data retrieval call binding the contract method 0x82980dc4.
//
// Solidity: function liquidityBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) LiquidityBridge() (common.Address, error) {
	return _MessageBusReceiver.Contract.LiquidityBridge(&_MessageBusReceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) Owner() (common.Address, error) {
	return _MessageBusReceiver.Contract.Owner(&_MessageBusReceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) Owner() (common.Address, error) {
	return _MessageBusReceiver.Contract.Owner(&_MessageBusReceiver.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) PegBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "pegBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) PegBridge() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegBridge(&_MessageBusReceiver.CallOpts)
}

// PegBridge is a free data retrieval call binding the contract method 0xdfa2dbaf.
//
// Solidity: function pegBridge() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) PegBridge() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegBridge(&_MessageBusReceiver.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) PegBridgeV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "pegBridgeV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) PegBridgeV2() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegBridgeV2(&_MessageBusReceiver.CallOpts)
}

// PegBridgeV2 is a free data retrieval call binding the contract method 0x95b12c27.
//
// Solidity: function pegBridgeV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) PegBridgeV2() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegBridgeV2(&_MessageBusReceiver.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) PegVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "pegVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) PegVault() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegVault(&_MessageBusReceiver.CallOpts)
}

// PegVault is a free data retrieval call binding the contract method 0xd8257d17.
//
// Solidity: function pegVault() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) PegVault() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegVault(&_MessageBusReceiver.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCaller) PegVaultV2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "pegVaultV2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverSession) PegVaultV2() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegVaultV2(&_MessageBusReceiver.CallOpts)
}

// PegVaultV2 is a free data retrieval call binding the contract method 0xc66a9c5a.
//
// Solidity: function pegVaultV2() view returns(address)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) PegVaultV2() (common.Address, error) {
	return _MessageBusReceiver.Contract.PegVaultV2(&_MessageBusReceiver.CallOpts)
}

// PreExecuteMessageGasUsage is a free data retrieval call binding the contract method 0x584e45e1.
//
// Solidity: function preExecuteMessageGasUsage() view returns(uint256)
func (_MessageBusReceiver *MessageBusReceiverCaller) PreExecuteMessageGasUsage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusReceiver.contract.Call(opts, &out, "preExecuteMessageGasUsage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreExecuteMessageGasUsage is a free data retrieval call binding the contract method 0x584e45e1.
//
// Solidity: function preExecuteMessageGasUsage() view returns(uint256)
func (_MessageBusReceiver *MessageBusReceiverSession) PreExecuteMessageGasUsage() (*big.Int, error) {
	return _MessageBusReceiver.Contract.PreExecuteMessageGasUsage(&_MessageBusReceiver.CallOpts)
}

// PreExecuteMessageGasUsage is a free data retrieval call binding the contract method 0x584e45e1.
//
// Solidity: function preExecuteMessageGasUsage() view returns(uint256)
func (_MessageBusReceiver *MessageBusReceiverCallerSession) PreExecuteMessageGasUsage() (*big.Int, error) {
	return _MessageBusReceiver.Contract.PreExecuteMessageGasUsage(&_MessageBusReceiver.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) ExecuteMessage(opts *bind.TransactOpts, _message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "executeMessage", _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverSession) ExecuteMessage(_message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessage(&_MessageBusReceiver.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x468a2d04.
//
// Solidity: function executeMessage(bytes _message, (address,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) ExecuteMessage(_message []byte, _route MsgDataTypesRouteInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessage(&_MessageBusReceiver.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0xdb2c20c8.
//
// Solidity: function executeMessage(bytes _message, (bytes,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) ExecuteMessage0(opts *bind.TransactOpts, _message []byte, _route MsgDataTypesRouteInfo2, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "executeMessage0", _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0xdb2c20c8.
//
// Solidity: function executeMessage(bytes _message, (bytes,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverSession) ExecuteMessage0(_message []byte, _route MsgDataTypesRouteInfo2, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessage0(&_MessageBusReceiver.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0xdb2c20c8.
//
// Solidity: function executeMessage(bytes _message, (bytes,address,uint64,bytes32) _route, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) ExecuteMessage0(_message []byte, _route MsgDataTypesRouteInfo2, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessage0(&_MessageBusReceiver.TransactOpts, _message, _route, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "executeMessageWithTransfer", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverSession) ExecuteMessageWithTransfer(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessageWithTransfer(&_MessageBusReceiver.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x3f395aff.
//
// Solidity: function executeMessageWithTransfer(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) ExecuteMessageWithTransfer(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessageWithTransfer(&_MessageBusReceiver.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "executeMessageWithTransferRefund", _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessageWithTransferRefund(&_MessageBusReceiver.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x7b80ab20.
//
// Solidity: function executeMessageWithTransferRefund(bytes _message, (uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32) _transfer, bytes[] _sigs, address[] _signers, uint256[] _powers) payable returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) ExecuteMessageWithTransferRefund(_message []byte, _transfer MsgDataTypesTransferInfo, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.ExecuteMessageWithTransferRefund(&_MessageBusReceiver.TransactOpts, _message, _transfer, _sigs, _signers, _powers)
}

// RefundAndExecuteMsg is a paid mutator transaction binding the contract method 0x40d0d026.
//
// Solidity: function refundAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) RefundAndExecuteMsg(opts *bind.TransactOpts, _transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "refundAndExecuteMsg", _transferParams, _msgParams)
}

// RefundAndExecuteMsg is a paid mutator transaction binding the contract method 0x40d0d026.
//
// Solidity: function refundAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) RefundAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.RefundAndExecuteMsg(&_MessageBusReceiver.TransactOpts, _transferParams, _msgParams)
}

// RefundAndExecuteMsg is a paid mutator transaction binding the contract method 0x40d0d026.
//
// Solidity: function refundAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) RefundAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.RefundAndExecuteMsg(&_MessageBusReceiver.TransactOpts, _transferParams, _msgParams)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetLiquidityBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setLiquidityBridge", _addr)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetLiquidityBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetLiquidityBridge(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetLiquidityBridge is a paid mutator transaction binding the contract method 0x588be02b.
//
// Solidity: function setLiquidityBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetLiquidityBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetLiquidityBridge(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetPegBridge(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setPegBridge", _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetPegBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegBridge(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegBridge is a paid mutator transaction binding the contract method 0x03cbfe66.
//
// Solidity: function setPegBridge(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetPegBridge(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegBridge(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetPegBridgeV2(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setPegBridgeV2", _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetPegBridgeV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegBridgeV2(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegBridgeV2 is a paid mutator transaction binding the contract method 0x82efd502.
//
// Solidity: function setPegBridgeV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetPegBridgeV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegBridgeV2(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetPegVault(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setPegVault", _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetPegVault(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegVault(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegVault is a paid mutator transaction binding the contract method 0x9b05a775.
//
// Solidity: function setPegVault(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetPegVault(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegVault(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetPegVaultV2(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setPegVaultV2", _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetPegVaultV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegVaultV2(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPegVaultV2 is a paid mutator transaction binding the contract method 0xf83b0fb9.
//
// Solidity: function setPegVaultV2(address _addr) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetPegVaultV2(_addr common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPegVaultV2(&_MessageBusReceiver.TransactOpts, _addr)
}

// SetPreExecuteMessageGasUsage is a paid mutator transaction binding the contract method 0x4586f331.
//
// Solidity: function setPreExecuteMessageGasUsage(uint256 _usage) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) SetPreExecuteMessageGasUsage(opts *bind.TransactOpts, _usage *big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "setPreExecuteMessageGasUsage", _usage)
}

// SetPreExecuteMessageGasUsage is a paid mutator transaction binding the contract method 0x4586f331.
//
// Solidity: function setPreExecuteMessageGasUsage(uint256 _usage) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) SetPreExecuteMessageGasUsage(_usage *big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPreExecuteMessageGasUsage(&_MessageBusReceiver.TransactOpts, _usage)
}

// SetPreExecuteMessageGasUsage is a paid mutator transaction binding the contract method 0x4586f331.
//
// Solidity: function setPreExecuteMessageGasUsage(uint256 _usage) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) SetPreExecuteMessageGasUsage(_usage *big.Int) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.SetPreExecuteMessageGasUsage(&_MessageBusReceiver.TransactOpts, _usage)
}

// TransferAndExecuteMsg is a paid mutator transaction binding the contract method 0x723d0a9d.
//
// Solidity: function transferAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) TransferAndExecuteMsg(opts *bind.TransactOpts, _transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "transferAndExecuteMsg", _transferParams, _msgParams)
}

// TransferAndExecuteMsg is a paid mutator transaction binding the contract method 0x723d0a9d.
//
// Solidity: function transferAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) TransferAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.TransferAndExecuteMsg(&_MessageBusReceiver.TransactOpts, _transferParams, _msgParams)
}

// TransferAndExecuteMsg is a paid mutator transaction binding the contract method 0x723d0a9d.
//
// Solidity: function transferAndExecuteMsg((bytes,bytes[],address[],uint256[]) _transferParams, (bytes,(uint8,address,address,address,uint256,uint64,uint64,bytes32,bytes32),bytes[],address[],uint256[]) _msgParams) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) TransferAndExecuteMsg(_transferParams MsgDataTypesBridgeTransferParams, _msgParams MsgDataTypesMsgWithTransferExecutionParams) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.TransferAndExecuteMsg(&_MessageBusReceiver.TransactOpts, _transferParams, _msgParams)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusReceiver *MessageBusReceiverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.TransferOwnership(&_MessageBusReceiver.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusReceiver *MessageBusReceiverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusReceiver.Contract.TransferOwnership(&_MessageBusReceiver.TransactOpts, newOwner)
}

// MessageBusReceiverCallRevertedIterator is returned from FilterCallReverted and is used to iterate over the raw logs and unpacked data for CallReverted events raised by the MessageBusReceiver contract.
type MessageBusReceiverCallRevertedIterator struct {
	Event *MessageBusReceiverCallReverted // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverCallRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverCallReverted)
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
		it.Event = new(MessageBusReceiverCallReverted)
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
func (it *MessageBusReceiverCallRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverCallRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverCallReverted represents a CallReverted event raised by the MessageBusReceiver contract.
type MessageBusReceiverCallReverted struct {
	Reason string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallReverted is a free log retrieval operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterCallReverted(opts *bind.FilterOpts) (*MessageBusReceiverCallRevertedIterator, error) {

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverCallRevertedIterator{contract: _MessageBusReceiver.contract, event: "CallReverted", logs: logs, sub: sub}, nil
}

// WatchCallReverted is a free log subscription operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchCallReverted(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverCallReverted) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "CallReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverCallReverted)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "CallReverted", log); err != nil {
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

// ParseCallReverted is a log parse operation binding the contract event 0xffdd6142bbb721f3400e3908b04b86f60649b2e4d191e3f4c50c32c3e6471d2f.
//
// Solidity: event CallReverted(string reason)
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParseCallReverted(log types.Log) (*MessageBusReceiverCallReverted, error) {
	event := new(MessageBusReceiverCallReverted)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "CallReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the MessageBusReceiver contract.
type MessageBusReceiverExecutedIterator struct {
	Event *MessageBusReceiverExecuted // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverExecuted)
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
		it.Event = new(MessageBusReceiverExecuted)
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
func (it *MessageBusReceiverExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverExecuted represents a Executed event raised by the MessageBusReceiver contract.
type MessageBusReceiverExecuted struct {
	MsgType    uint8
	MsgId      [32]byte
	Status     uint8
	Receiver   common.Address
	SrcChainId uint64
	SrcTxHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0xa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d.
//
// Solidity: event Executed(uint8 msgType, bytes32 msgId, uint8 status, address indexed receiver, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterExecuted(opts *bind.FilterOpts, receiver []common.Address) (*MessageBusReceiverExecutedIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "Executed", receiverRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverExecutedIterator{contract: _MessageBusReceiver.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0xa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d.
//
// Solidity: event Executed(uint8 msgType, bytes32 msgId, uint8 status, address indexed receiver, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverExecuted, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "Executed", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverExecuted)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0xa635eb05143f74743822bbd96428928de4c8ee8cc578299749be9425c17bb34d.
//
// Solidity: event Executed(uint8 msgType, bytes32 msgId, uint8 status, address indexed receiver, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParseExecuted(log types.Log) (*MessageBusReceiverExecuted, error) {
	event := new(MessageBusReceiverExecuted)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverLiquidityBridgeUpdatedIterator is returned from FilterLiquidityBridgeUpdated and is used to iterate over the raw logs and unpacked data for LiquidityBridgeUpdated events raised by the MessageBusReceiver contract.
type MessageBusReceiverLiquidityBridgeUpdatedIterator struct {
	Event *MessageBusReceiverLiquidityBridgeUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverLiquidityBridgeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverLiquidityBridgeUpdated)
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
		it.Event = new(MessageBusReceiverLiquidityBridgeUpdated)
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
func (it *MessageBusReceiverLiquidityBridgeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverLiquidityBridgeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverLiquidityBridgeUpdated represents a LiquidityBridgeUpdated event raised by the MessageBusReceiver contract.
type MessageBusReceiverLiquidityBridgeUpdated struct {
	LiquidityBridge common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidityBridgeUpdated is a free log retrieval operation binding the contract event 0xbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae058.
//
// Solidity: event LiquidityBridgeUpdated(address liquidityBridge)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterLiquidityBridgeUpdated(opts *bind.FilterOpts) (*MessageBusReceiverLiquidityBridgeUpdatedIterator, error) {

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "LiquidityBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverLiquidityBridgeUpdatedIterator{contract: _MessageBusReceiver.contract, event: "LiquidityBridgeUpdated", logs: logs, sub: sub}, nil
}

// WatchLiquidityBridgeUpdated is a free log subscription operation binding the contract event 0xbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae058.
//
// Solidity: event LiquidityBridgeUpdated(address liquidityBridge)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchLiquidityBridgeUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverLiquidityBridgeUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "LiquidityBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverLiquidityBridgeUpdated)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "LiquidityBridgeUpdated", log); err != nil {
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

// ParseLiquidityBridgeUpdated is a log parse operation binding the contract event 0xbf9977180dc6e6cff25598c8e59150cecd7f8e448e092633d38ab7ee223ae058.
//
// Solidity: event LiquidityBridgeUpdated(address liquidityBridge)
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParseLiquidityBridgeUpdated(log types.Log) (*MessageBusReceiverLiquidityBridgeUpdated, error) {
	event := new(MessageBusReceiverLiquidityBridgeUpdated)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "LiquidityBridgeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverNeedRetryIterator is returned from FilterNeedRetry and is used to iterate over the raw logs and unpacked data for NeedRetry events raised by the MessageBusReceiver contract.
type MessageBusReceiverNeedRetryIterator struct {
	Event *MessageBusReceiverNeedRetry // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverNeedRetryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverNeedRetry)
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
		it.Event = new(MessageBusReceiverNeedRetry)
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
func (it *MessageBusReceiverNeedRetryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverNeedRetryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverNeedRetry represents a NeedRetry event raised by the MessageBusReceiver contract.
type MessageBusReceiverNeedRetry struct {
	MsgType    uint8
	MsgId      [32]byte
	SrcChainId uint64
	SrcTxHash  [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNeedRetry is a free log retrieval operation binding the contract event 0xe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c.
//
// Solidity: event NeedRetry(uint8 msgType, bytes32 msgId, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterNeedRetry(opts *bind.FilterOpts) (*MessageBusReceiverNeedRetryIterator, error) {

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "NeedRetry")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverNeedRetryIterator{contract: _MessageBusReceiver.contract, event: "NeedRetry", logs: logs, sub: sub}, nil
}

// WatchNeedRetry is a free log subscription operation binding the contract event 0xe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c.
//
// Solidity: event NeedRetry(uint8 msgType, bytes32 msgId, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchNeedRetry(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverNeedRetry) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "NeedRetry")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverNeedRetry)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "NeedRetry", log); err != nil {
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

// ParseNeedRetry is a log parse operation binding the contract event 0xe49c2c954d381d1448cf824743aeff9da7a1d82078a7c9e5817269cc359bd26c.
//
// Solidity: event NeedRetry(uint8 msgType, bytes32 msgId, uint64 srcChainId, bytes32 srcTxHash)
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParseNeedRetry(log types.Log) (*MessageBusReceiverNeedRetry, error) {
	event := new(MessageBusReceiverNeedRetry)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "NeedRetry", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageBusReceiver contract.
type MessageBusReceiverOwnershipTransferredIterator struct {
	Event *MessageBusReceiverOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverOwnershipTransferred)
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
		it.Event = new(MessageBusReceiverOwnershipTransferred)
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
func (it *MessageBusReceiverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverOwnershipTransferred represents a OwnershipTransferred event raised by the MessageBusReceiver contract.
type MessageBusReceiverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageBusReceiverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverOwnershipTransferredIterator{contract: _MessageBusReceiver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverOwnershipTransferred)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParseOwnershipTransferred(log types.Log) (*MessageBusReceiverOwnershipTransferred, error) {
	event := new(MessageBusReceiverOwnershipTransferred)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverPegBridgeUpdatedIterator is returned from FilterPegBridgeUpdated and is used to iterate over the raw logs and unpacked data for PegBridgeUpdated events raised by the MessageBusReceiver contract.
type MessageBusReceiverPegBridgeUpdatedIterator struct {
	Event *MessageBusReceiverPegBridgeUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverPegBridgeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverPegBridgeUpdated)
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
		it.Event = new(MessageBusReceiverPegBridgeUpdated)
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
func (it *MessageBusReceiverPegBridgeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverPegBridgeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverPegBridgeUpdated represents a PegBridgeUpdated event raised by the MessageBusReceiver contract.
type MessageBusReceiverPegBridgeUpdated struct {
	PegBridge common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPegBridgeUpdated is a free log retrieval operation binding the contract event 0xd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d.
//
// Solidity: event PegBridgeUpdated(address pegBridge)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterPegBridgeUpdated(opts *bind.FilterOpts) (*MessageBusReceiverPegBridgeUpdatedIterator, error) {

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "PegBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverPegBridgeUpdatedIterator{contract: _MessageBusReceiver.contract, event: "PegBridgeUpdated", logs: logs, sub: sub}, nil
}

// WatchPegBridgeUpdated is a free log subscription operation binding the contract event 0xd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d.
//
// Solidity: event PegBridgeUpdated(address pegBridge)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchPegBridgeUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverPegBridgeUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "PegBridgeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverPegBridgeUpdated)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "PegBridgeUpdated", log); err != nil {
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

// ParsePegBridgeUpdated is a log parse operation binding the contract event 0xd60e9ceb4f54f1bfb1741a4b35fc9d806d7ed48200b523203b92248ea38fa17d.
//
// Solidity: event PegBridgeUpdated(address pegBridge)
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParsePegBridgeUpdated(log types.Log) (*MessageBusReceiverPegBridgeUpdated, error) {
	event := new(MessageBusReceiverPegBridgeUpdated)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "PegBridgeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverPegBridgeV2UpdatedIterator is returned from FilterPegBridgeV2Updated and is used to iterate over the raw logs and unpacked data for PegBridgeV2Updated events raised by the MessageBusReceiver contract.
type MessageBusReceiverPegBridgeV2UpdatedIterator struct {
	Event *MessageBusReceiverPegBridgeV2Updated // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverPegBridgeV2UpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverPegBridgeV2Updated)
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
		it.Event = new(MessageBusReceiverPegBridgeV2Updated)
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
func (it *MessageBusReceiverPegBridgeV2UpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverPegBridgeV2UpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverPegBridgeV2Updated represents a PegBridgeV2Updated event raised by the MessageBusReceiver contract.
type MessageBusReceiverPegBridgeV2Updated struct {
	PegBridgeV2 common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPegBridgeV2Updated is a free log retrieval operation binding the contract event 0xfb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e8.
//
// Solidity: event PegBridgeV2Updated(address pegBridgeV2)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterPegBridgeV2Updated(opts *bind.FilterOpts) (*MessageBusReceiverPegBridgeV2UpdatedIterator, error) {

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "PegBridgeV2Updated")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverPegBridgeV2UpdatedIterator{contract: _MessageBusReceiver.contract, event: "PegBridgeV2Updated", logs: logs, sub: sub}, nil
}

// WatchPegBridgeV2Updated is a free log subscription operation binding the contract event 0xfb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e8.
//
// Solidity: event PegBridgeV2Updated(address pegBridgeV2)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchPegBridgeV2Updated(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverPegBridgeV2Updated) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "PegBridgeV2Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverPegBridgeV2Updated)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "PegBridgeV2Updated", log); err != nil {
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

// ParsePegBridgeV2Updated is a log parse operation binding the contract event 0xfb337a6c76476534518d5816caeb86263972470fedccfd047a35eb1825eaa9e8.
//
// Solidity: event PegBridgeV2Updated(address pegBridgeV2)
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParsePegBridgeV2Updated(log types.Log) (*MessageBusReceiverPegBridgeV2Updated, error) {
	event := new(MessageBusReceiverPegBridgeV2Updated)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "PegBridgeV2Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverPegVaultUpdatedIterator is returned from FilterPegVaultUpdated and is used to iterate over the raw logs and unpacked data for PegVaultUpdated events raised by the MessageBusReceiver contract.
type MessageBusReceiverPegVaultUpdatedIterator struct {
	Event *MessageBusReceiverPegVaultUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverPegVaultUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverPegVaultUpdated)
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
		it.Event = new(MessageBusReceiverPegVaultUpdated)
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
func (it *MessageBusReceiverPegVaultUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverPegVaultUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverPegVaultUpdated represents a PegVaultUpdated event raised by the MessageBusReceiver contract.
type MessageBusReceiverPegVaultUpdated struct {
	PegVault common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPegVaultUpdated is a free log retrieval operation binding the contract event 0xa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad.
//
// Solidity: event PegVaultUpdated(address pegVault)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterPegVaultUpdated(opts *bind.FilterOpts) (*MessageBusReceiverPegVaultUpdatedIterator, error) {

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "PegVaultUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverPegVaultUpdatedIterator{contract: _MessageBusReceiver.contract, event: "PegVaultUpdated", logs: logs, sub: sub}, nil
}

// WatchPegVaultUpdated is a free log subscription operation binding the contract event 0xa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad.
//
// Solidity: event PegVaultUpdated(address pegVault)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchPegVaultUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverPegVaultUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "PegVaultUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverPegVaultUpdated)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "PegVaultUpdated", log); err != nil {
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

// ParsePegVaultUpdated is a log parse operation binding the contract event 0xa9db0c32d9c6c2f75f3b95047a9e67cc1c010eab792a4e6ca777ce918ad94aad.
//
// Solidity: event PegVaultUpdated(address pegVault)
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParsePegVaultUpdated(log types.Log) (*MessageBusReceiverPegVaultUpdated, error) {
	event := new(MessageBusReceiverPegVaultUpdated)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "PegVaultUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusReceiverPegVaultV2UpdatedIterator is returned from FilterPegVaultV2Updated and is used to iterate over the raw logs and unpacked data for PegVaultV2Updated events raised by the MessageBusReceiver contract.
type MessageBusReceiverPegVaultV2UpdatedIterator struct {
	Event *MessageBusReceiverPegVaultV2Updated // Event containing the contract specifics and raw log

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
func (it *MessageBusReceiverPegVaultV2UpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusReceiverPegVaultV2Updated)
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
		it.Event = new(MessageBusReceiverPegVaultV2Updated)
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
func (it *MessageBusReceiverPegVaultV2UpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusReceiverPegVaultV2UpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusReceiverPegVaultV2Updated represents a PegVaultV2Updated event raised by the MessageBusReceiver contract.
type MessageBusReceiverPegVaultV2Updated struct {
	PegVaultV2 common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPegVaultV2Updated is a free log retrieval operation binding the contract event 0x918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb5.
//
// Solidity: event PegVaultV2Updated(address pegVaultV2)
func (_MessageBusReceiver *MessageBusReceiverFilterer) FilterPegVaultV2Updated(opts *bind.FilterOpts) (*MessageBusReceiverPegVaultV2UpdatedIterator, error) {

	logs, sub, err := _MessageBusReceiver.contract.FilterLogs(opts, "PegVaultV2Updated")
	if err != nil {
		return nil, err
	}
	return &MessageBusReceiverPegVaultV2UpdatedIterator{contract: _MessageBusReceiver.contract, event: "PegVaultV2Updated", logs: logs, sub: sub}, nil
}

// WatchPegVaultV2Updated is a free log subscription operation binding the contract event 0x918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb5.
//
// Solidity: event PegVaultV2Updated(address pegVaultV2)
func (_MessageBusReceiver *MessageBusReceiverFilterer) WatchPegVaultV2Updated(opts *bind.WatchOpts, sink chan<- *MessageBusReceiverPegVaultV2Updated) (event.Subscription, error) {

	logs, sub, err := _MessageBusReceiver.contract.WatchLogs(opts, "PegVaultV2Updated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusReceiverPegVaultV2Updated)
				if err := _MessageBusReceiver.contract.UnpackLog(event, "PegVaultV2Updated", log); err != nil {
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

// ParsePegVaultV2Updated is a log parse operation binding the contract event 0x918a691a2a82482a10e11f43d7b627b2ba220dd08f251cb61933c42560f6fcb5.
//
// Solidity: event PegVaultV2Updated(address pegVaultV2)
func (_MessageBusReceiver *MessageBusReceiverFilterer) ParsePegVaultV2Updated(log types.Log) (*MessageBusReceiverPegVaultV2Updated, error) {
	event := new(MessageBusReceiverPegVaultV2Updated)
	if err := _MessageBusReceiver.contract.UnpackLog(event, "PegVaultV2Updated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderMetaData contains all meta data concerning the MessageBusSender contract.
var MessageBusSenderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"_sigsVerifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBase\",\"type\":\"uint256\"}],\"name\":\"FeeBaseUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePerByte\",\"type\":\"uint256\"}],\"name\":\"FeePerByteUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FeeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Message\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"receiver\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"Message2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bridge\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTransferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"MessageWithTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"calcFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePerByte\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_receiver\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_srcBridge\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_srcTransferId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithTransfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFeeBase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"setFeePerByte\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sigsVerifier\",\"outputs\":[{\"internalType\":\"contractISigsVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_cumulativeFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawnFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161116638038061116683398101604081905261002f91610099565b61003833610049565b6001600160a01b03166080526100c9565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100ab57600080fd5b81516001600160a01b03811681146100c257600080fd5b9392505050565b60805161107b6100eb6000396000818161020c01526103e3015261107b6000f3fe6080604052600436106100d25760003560e01c80638da5cb5b1161007f578063ccf2683b11610059578063ccf2683b146101fa578063e2c1ed251461022e578063f2fde38b1461024e578063f60bbe2a1461026e57600080fd5b80638da5cb5b1461019f57806395e911a8146101d15780639f3ce55a146101e757600080fd5b80635335dca2116100b05780635335dca21461012c5780635b3e5f501461015f5780637d7a101d1461018c57600080fd5b806306c28bd6146100d75780632ff4c411146100f95780634289fbb314610119575b600080fd5b3480156100e357600080fd5b506100f76100f2366004610a27565b610284565b005b34801561010557600080fd5b506100f7610114366004610aa8565b61032e565b6100f7610127366004610b9e565b6105d3565b34801561013857600080fd5b5061014c610147366004610c16565b6106bb565b6040519081526020015b60405180910390f35b34801561016b57600080fd5b5061014c61017a366004610c58565b60036020526000908152604090205481565b6100f761019a366004610c73565b6106e1565b3480156101ab57600080fd5b506000546001600160a01b03165b6040516001600160a01b039091168152602001610156565b3480156101dd57600080fd5b5061014c60015481565b6100f76101f5366004610ced565b61073e565b34801561020657600080fd5b506101b97f000000000000000000000000000000000000000000000000000000000000000081565b34801561023a57600080fd5b506100f7610249366004610a27565b610798565b34801561025a57600080fd5b506100f7610269366004610c58565b610836565b34801561027a57600080fd5b5061014c60025481565b336102976000546001600160a01b031690565b6001600160a01b0316146102f25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064015b60405180910390fd5b60018190556040518181527f892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f2906020015b60405180910390a150565b6000463060405160200161038492919091825260601b6bffffffffffffffffffffffff191660208201527f77697468647261774665650000000000000000000000000000000000000000006034820152603f0190565b60408051808303601f19018152828252805160209182012090830181905260608c901b6bffffffffffffffffffffffff19168383015260548084018c9052825180850390910181526074840192839052633416de1160e11b90925292507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03169163682dbc229161042a918b908b908b908b908b908b90607801610e94565b60006040518083038186803b15801561044257600080fd5b505afa158015610456573d6000803e3d6000fd5b505050506001600160a01b03891660009081526003602052604081205461047d908a610f3f565b9050600081116104cf5760405162461bcd60e51b815260206004820152601960248201527f4e6f206e657720616d6f756e7420746f2077697468647261770000000000000060448201526064016102e9565b6001600160a01b038a166000818152600360205260408082208c90555190919061c35090849084818181858888f193505050503d806000811461052e576040519150601f19603f3d011682016040523d82523d6000602084013e610533565b606091505b50509050806105845760405162461bcd60e51b815260206004820152601660248201527f6661696c656420746f207769746864726177206665650000000000000000000060448201526064016102e9565b604080516001600160a01b038d168152602081018490527f78473f3f373f7673597f4f0fa5873cb4d375fea6d4339ad6b56dbd411513cb3f910160405180910390a15050505050505050505050565b4685036106145760405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9259608a1b60448201526064016102e9565b600061062083836106bb565b9050803410156106655760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b60448201526064016102e9565b336001600160a01b03167f172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66888888888888346040516106aa9796959493929190610f52565b60405180910390a250505050505050565b6002546000906106cb9083610f9f565b6001546106d89190610fb6565b90505b92915050565b6106ec838383610927565b336001600160a01b03167fe66fbe37d84ca73c589f782ac278844918ea6c56a4917f58707f715588080df286868686863460405161072f96959493929190610fc9565b60405180910390a25050505050565b610749838383610927565b336001600160a01b03167fce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4858585853460405161078a95949392919061100a565b60405180910390a250505050565b336107ab6000546001600160a01b031690565b6001600160a01b0316146108015760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102e9565b60028190556040518181527f210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c90602001610323565b336108496000546001600160a01b031690565b6001600160a01b03161461089f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102e9565b6001600160a01b03811661091b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016102e9565b610924816109bf565b50565b4683036109685760405162461bcd60e51b815260206004820152600f60248201526e125b9d985b1a590818da185a5b9259608a1b60448201526064016102e9565b600061097483836106bb565b9050803410156109b95760405162461bcd60e51b815260206004820152601060248201526f496e73756666696369656e742066656560801b60448201526064016102e9565b50505050565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600060208284031215610a3957600080fd5b5035919050565b80356001600160a01b0381168114610a5757600080fd5b919050565b60008083601f840112610a6e57600080fd5b50813567ffffffffffffffff811115610a8657600080fd5b6020830191508360208260051b8501011115610aa157600080fd5b9250929050565b60008060008060008060008060a0898b031215610ac457600080fd5b610acd89610a40565b975060208901359650604089013567ffffffffffffffff80821115610af157600080fd5b610afd8c838d01610a5c565b909850965060608b0135915080821115610b1657600080fd5b610b228c838d01610a5c565b909650945060808b0135915080821115610b3b57600080fd5b50610b488b828c01610a5c565b999c989b5096995094979396929594505050565b60008083601f840112610b6e57600080fd5b50813567ffffffffffffffff811115610b8657600080fd5b602083019150836020828501011115610aa157600080fd5b60008060008060008060a08789031215610bb757600080fd5b610bc087610a40565b955060208701359450610bd560408801610a40565b935060608701359250608087013567ffffffffffffffff811115610bf857600080fd5b610c0489828a01610b5c565b979a9699509497509295939492505050565b60008060208385031215610c2957600080fd5b823567ffffffffffffffff811115610c4057600080fd5b610c4c85828601610b5c565b90969095509350505050565b600060208284031215610c6a57600080fd5b6106d882610a40565b600080600080600060608688031215610c8b57600080fd5b853567ffffffffffffffff80821115610ca357600080fd5b610caf89838a01610b5c565b9097509550602088013594506040880135915080821115610ccf57600080fd5b50610cdc88828901610b5c565b969995985093965092949392505050565b60008060008060608587031215610d0357600080fd5b610d0c85610a40565b935060208501359250604085013567ffffffffffffffff811115610d2f57600080fd5b610d3b87828801610b5c565b95989497509550505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b81835260006020808501808196508560051b810191508460005b87811015610df55782840389528135601e19883603018112610dab57600080fd5b8701858101903567ffffffffffffffff811115610dc757600080fd5b803603821315610dd657600080fd5b610de1868284610d47565b9a87019a9550505090840190600101610d8a565b5091979650505050505050565b8183526000602080850194508260005b85811015610e3e576001600160a01b03610e2b83610a40565b1687529582019590820190600101610e12565b509495945050505050565b81835260007f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115610e7b57600080fd5b8260051b80836020870137939093016020019392505050565b608081526000885180608084015260005b81811015610ec2576020818c0181015160a0868401015201610ea5565b50600060a08285010152601f19601f8201168301905060a0838203016020840152610ef160a08201898b610d70565b90508281036040840152610f06818789610e02565b90508281036060840152610f1b818587610e49565b9a9950505050505050505050565b634e487b7160e01b600052601160045260246000fd5b818103818111156106db576106db610f29565b60006001600160a01b03808a16835288602084015280881660408401525085606083015260c06080830152610f8b60c083018587610d47565b90508260a083015298975050505050505050565b80820281158282048414176106db576106db610f29565b808201808211156106db576106db610f29565b608081526000610fdd60808301888a610d47565b8660208401528281036040840152610ff6818688610d47565b915050826060830152979650505050505050565b6001600160a01b0386168152846020820152608060408201526000611033608083018587610d47565b9050826060830152969550505050505056fea2646970667358221220e28d377c50635a5b229422b202a7e96458ebf96e4cde136dfbc29ff1c2df1b1464736f6c63430008110033",
}

// MessageBusSenderABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageBusSenderMetaData.ABI instead.
var MessageBusSenderABI = MessageBusSenderMetaData.ABI

// MessageBusSenderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageBusSenderMetaData.Bin instead.
var MessageBusSenderBin = MessageBusSenderMetaData.Bin

// DeployMessageBusSender deploys a new Ethereum contract, binding an instance of MessageBusSender to it.
func DeployMessageBusSender(auth *bind.TransactOpts, backend bind.ContractBackend, _sigsVerifier common.Address) (common.Address, *types.Transaction, *MessageBusSender, error) {
	parsed, err := MessageBusSenderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBusSenderBin), backend, _sigsVerifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageBusSender{MessageBusSenderCaller: MessageBusSenderCaller{contract: contract}, MessageBusSenderTransactor: MessageBusSenderTransactor{contract: contract}, MessageBusSenderFilterer: MessageBusSenderFilterer{contract: contract}}, nil
}

// MessageBusSender is an auto generated Go binding around an Ethereum contract.
type MessageBusSender struct {
	MessageBusSenderCaller     // Read-only binding to the contract
	MessageBusSenderTransactor // Write-only binding to the contract
	MessageBusSenderFilterer   // Log filterer for contract events
}

// MessageBusSenderCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageBusSenderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusSenderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageBusSenderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusSenderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageBusSenderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageBusSenderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageBusSenderSession struct {
	Contract     *MessageBusSender // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageBusSenderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageBusSenderCallerSession struct {
	Contract *MessageBusSenderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MessageBusSenderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageBusSenderTransactorSession struct {
	Contract     *MessageBusSenderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MessageBusSenderRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageBusSenderRaw struct {
	Contract *MessageBusSender // Generic contract binding to access the raw methods on
}

// MessageBusSenderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageBusSenderCallerRaw struct {
	Contract *MessageBusSenderCaller // Generic read-only contract binding to access the raw methods on
}

// MessageBusSenderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageBusSenderTransactorRaw struct {
	Contract *MessageBusSenderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageBusSender creates a new instance of MessageBusSender, bound to a specific deployed contract.
func NewMessageBusSender(address common.Address, backend bind.ContractBackend) (*MessageBusSender, error) {
	contract, err := bindMessageBusSender(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageBusSender{MessageBusSenderCaller: MessageBusSenderCaller{contract: contract}, MessageBusSenderTransactor: MessageBusSenderTransactor{contract: contract}, MessageBusSenderFilterer: MessageBusSenderFilterer{contract: contract}}, nil
}

// NewMessageBusSenderCaller creates a new read-only instance of MessageBusSender, bound to a specific deployed contract.
func NewMessageBusSenderCaller(address common.Address, caller bind.ContractCaller) (*MessageBusSenderCaller, error) {
	contract, err := bindMessageBusSender(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderCaller{contract: contract}, nil
}

// NewMessageBusSenderTransactor creates a new write-only instance of MessageBusSender, bound to a specific deployed contract.
func NewMessageBusSenderTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageBusSenderTransactor, error) {
	contract, err := bindMessageBusSender(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderTransactor{contract: contract}, nil
}

// NewMessageBusSenderFilterer creates a new log filterer instance of MessageBusSender, bound to a specific deployed contract.
func NewMessageBusSenderFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageBusSenderFilterer, error) {
	contract, err := bindMessageBusSender(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderFilterer{contract: contract}, nil
}

// bindMessageBusSender binds a generic wrapper to an already deployed contract.
func bindMessageBusSender(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageBusSenderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusSender *MessageBusSenderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusSender.Contract.MessageBusSenderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusSender *MessageBusSenderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusSender.Contract.MessageBusSenderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusSender *MessageBusSenderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusSender.Contract.MessageBusSenderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageBusSender *MessageBusSenderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageBusSender.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageBusSender *MessageBusSenderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageBusSender.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageBusSender *MessageBusSenderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageBusSender.Contract.contract.Transact(opts, method, params...)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBusSender *MessageBusSenderCaller) CalcFee(opts *bind.CallOpts, _message []byte) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "calcFee", _message)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBusSender *MessageBusSenderSession) CalcFee(_message []byte) (*big.Int, error) {
	return _MessageBusSender.Contract.CalcFee(&_MessageBusSender.CallOpts, _message)
}

// CalcFee is a free data retrieval call binding the contract method 0x5335dca2.
//
// Solidity: function calcFee(bytes _message) view returns(uint256)
func (_MessageBusSender *MessageBusSenderCallerSession) CalcFee(_message []byte) (*big.Int, error) {
	return _MessageBusSender.Contract.CalcFee(&_MessageBusSender.CallOpts, _message)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBusSender *MessageBusSenderCaller) FeeBase(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "feeBase")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBusSender *MessageBusSenderSession) FeeBase() (*big.Int, error) {
	return _MessageBusSender.Contract.FeeBase(&_MessageBusSender.CallOpts)
}

// FeeBase is a free data retrieval call binding the contract method 0x95e911a8.
//
// Solidity: function feeBase() view returns(uint256)
func (_MessageBusSender *MessageBusSenderCallerSession) FeeBase() (*big.Int, error) {
	return _MessageBusSender.Contract.FeeBase(&_MessageBusSender.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBusSender *MessageBusSenderCaller) FeePerByte(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "feePerByte")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBusSender *MessageBusSenderSession) FeePerByte() (*big.Int, error) {
	return _MessageBusSender.Contract.FeePerByte(&_MessageBusSender.CallOpts)
}

// FeePerByte is a free data retrieval call binding the contract method 0xf60bbe2a.
//
// Solidity: function feePerByte() view returns(uint256)
func (_MessageBusSender *MessageBusSenderCallerSession) FeePerByte() (*big.Int, error) {
	return _MessageBusSender.Contract.FeePerByte(&_MessageBusSender.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusSender *MessageBusSenderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusSender *MessageBusSenderSession) Owner() (common.Address, error) {
	return _MessageBusSender.Contract.Owner(&_MessageBusSender.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageBusSender *MessageBusSenderCallerSession) Owner() (common.Address, error) {
	return _MessageBusSender.Contract.Owner(&_MessageBusSender.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBusSender *MessageBusSenderCaller) SigsVerifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "sigsVerifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBusSender *MessageBusSenderSession) SigsVerifier() (common.Address, error) {
	return _MessageBusSender.Contract.SigsVerifier(&_MessageBusSender.CallOpts)
}

// SigsVerifier is a free data retrieval call binding the contract method 0xccf2683b.
//
// Solidity: function sigsVerifier() view returns(address)
func (_MessageBusSender *MessageBusSenderCallerSession) SigsVerifier() (common.Address, error) {
	return _MessageBusSender.Contract.SigsVerifier(&_MessageBusSender.CallOpts)
}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBusSender *MessageBusSenderCaller) WithdrawnFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MessageBusSender.contract.Call(opts, &out, "withdrawnFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBusSender *MessageBusSenderSession) WithdrawnFees(arg0 common.Address) (*big.Int, error) {
	return _MessageBusSender.Contract.WithdrawnFees(&_MessageBusSender.CallOpts, arg0)
}

// WithdrawnFees is a free data retrieval call binding the contract method 0x5b3e5f50.
//
// Solidity: function withdrawnFees(address ) view returns(uint256)
func (_MessageBusSender *MessageBusSenderCallerSession) WithdrawnFees(arg0 common.Address) (*big.Int, error) {
	return _MessageBusSender.Contract.WithdrawnFees(&_MessageBusSender.CallOpts, arg0)
}

// SendMessage is a paid mutator transaction binding the contract method 0x7d7a101d.
//
// Solidity: function sendMessage(bytes _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderTransactor) SendMessage(opts *bind.TransactOpts, _receiver []byte, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "sendMessage", _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x7d7a101d.
//
// Solidity: function sendMessage(bytes _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderSession) SendMessage(_receiver []byte, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SendMessage(&_MessageBusSender.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessage is a paid mutator transaction binding the contract method 0x7d7a101d.
//
// Solidity: function sendMessage(bytes _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) SendMessage(_receiver []byte, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SendMessage(&_MessageBusSender.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderTransactor) SendMessage0(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "sendMessage0", _receiver, _dstChainId, _message)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderSession) SendMessage0(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SendMessage0(&_MessageBusSender.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessage0 is a paid mutator transaction binding the contract method 0x9f3ce55a.
//
// Solidity: function sendMessage(address _receiver, uint256 _dstChainId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) SendMessage0(_receiver common.Address, _dstChainId *big.Int, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SendMessage0(&_MessageBusSender.TransactOpts, _receiver, _dstChainId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderTransactor) SendMessageWithTransfer(opts *bind.TransactOpts, _receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "sendMessageWithTransfer", _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SendMessageWithTransfer(&_MessageBusSender.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SendMessageWithTransfer is a paid mutator transaction binding the contract method 0x4289fbb3.
//
// Solidity: function sendMessageWithTransfer(address _receiver, uint256 _dstChainId, address _srcBridge, bytes32 _srcTransferId, bytes _message) payable returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) SendMessageWithTransfer(_receiver common.Address, _dstChainId *big.Int, _srcBridge common.Address, _srcTransferId [32]byte, _message []byte) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SendMessageWithTransfer(&_MessageBusSender.TransactOpts, _receiver, _dstChainId, _srcBridge, _srcTransferId, _message)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderTransactor) SetFeeBase(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "setFeeBase", _fee)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderSession) SetFeeBase(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SetFeeBase(&_MessageBusSender.TransactOpts, _fee)
}

// SetFeeBase is a paid mutator transaction binding the contract method 0x06c28bd6.
//
// Solidity: function setFeeBase(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) SetFeeBase(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SetFeeBase(&_MessageBusSender.TransactOpts, _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderTransactor) SetFeePerByte(opts *bind.TransactOpts, _fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "setFeePerByte", _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderSession) SetFeePerByte(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SetFeePerByte(&_MessageBusSender.TransactOpts, _fee)
}

// SetFeePerByte is a paid mutator transaction binding the contract method 0xe2c1ed25.
//
// Solidity: function setFeePerByte(uint256 _fee) returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) SetFeePerByte(_fee *big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.SetFeePerByte(&_MessageBusSender.TransactOpts, _fee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusSender *MessageBusSenderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusSender *MessageBusSenderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusSender.Contract.TransferOwnership(&_MessageBusSender.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageBusSender.Contract.TransferOwnership(&_MessageBusSender.TransactOpts, newOwner)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBusSender *MessageBusSenderTransactor) WithdrawFee(opts *bind.TransactOpts, _account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusSender.contract.Transact(opts, "withdrawFee", _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBusSender *MessageBusSenderSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.WithdrawFee(&_MessageBusSender.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x2ff4c411.
//
// Solidity: function withdrawFee(address _account, uint256 _cumulativeFee, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_MessageBusSender *MessageBusSenderTransactorSession) WithdrawFee(_account common.Address, _cumulativeFee *big.Int, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _MessageBusSender.Contract.WithdrawFee(&_MessageBusSender.TransactOpts, _account, _cumulativeFee, _sigs, _signers, _powers)
}

// MessageBusSenderFeeBaseUpdatedIterator is returned from FilterFeeBaseUpdated and is used to iterate over the raw logs and unpacked data for FeeBaseUpdated events raised by the MessageBusSender contract.
type MessageBusSenderFeeBaseUpdatedIterator struct {
	Event *MessageBusSenderFeeBaseUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderFeeBaseUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderFeeBaseUpdated)
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
		it.Event = new(MessageBusSenderFeeBaseUpdated)
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
func (it *MessageBusSenderFeeBaseUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderFeeBaseUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderFeeBaseUpdated represents a FeeBaseUpdated event raised by the MessageBusSender contract.
type MessageBusSenderFeeBaseUpdated struct {
	FeeBase *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeBaseUpdated is a free log retrieval operation binding the contract event 0x892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f2.
//
// Solidity: event FeeBaseUpdated(uint256 feeBase)
func (_MessageBusSender *MessageBusSenderFilterer) FilterFeeBaseUpdated(opts *bind.FilterOpts) (*MessageBusSenderFeeBaseUpdatedIterator, error) {

	logs, sub, err := _MessageBusSender.contract.FilterLogs(opts, "FeeBaseUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderFeeBaseUpdatedIterator{contract: _MessageBusSender.contract, event: "FeeBaseUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeBaseUpdated is a free log subscription operation binding the contract event 0x892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f2.
//
// Solidity: event FeeBaseUpdated(uint256 feeBase)
func (_MessageBusSender *MessageBusSenderFilterer) WatchFeeBaseUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusSenderFeeBaseUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBusSender.contract.WatchLogs(opts, "FeeBaseUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderFeeBaseUpdated)
				if err := _MessageBusSender.contract.UnpackLog(event, "FeeBaseUpdated", log); err != nil {
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

// ParseFeeBaseUpdated is a log parse operation binding the contract event 0x892dfdc99ecd3bb4f2f2cb118dca02f0bd16640ff156d3c6459d4282e336a5f2.
//
// Solidity: event FeeBaseUpdated(uint256 feeBase)
func (_MessageBusSender *MessageBusSenderFilterer) ParseFeeBaseUpdated(log types.Log) (*MessageBusSenderFeeBaseUpdated, error) {
	event := new(MessageBusSenderFeeBaseUpdated)
	if err := _MessageBusSender.contract.UnpackLog(event, "FeeBaseUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderFeePerByteUpdatedIterator is returned from FilterFeePerByteUpdated and is used to iterate over the raw logs and unpacked data for FeePerByteUpdated events raised by the MessageBusSender contract.
type MessageBusSenderFeePerByteUpdatedIterator struct {
	Event *MessageBusSenderFeePerByteUpdated // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderFeePerByteUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderFeePerByteUpdated)
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
		it.Event = new(MessageBusSenderFeePerByteUpdated)
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
func (it *MessageBusSenderFeePerByteUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderFeePerByteUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderFeePerByteUpdated represents a FeePerByteUpdated event raised by the MessageBusSender contract.
type MessageBusSenderFeePerByteUpdated struct {
	FeePerByte *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeePerByteUpdated is a free log retrieval operation binding the contract event 0x210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c.
//
// Solidity: event FeePerByteUpdated(uint256 feePerByte)
func (_MessageBusSender *MessageBusSenderFilterer) FilterFeePerByteUpdated(opts *bind.FilterOpts) (*MessageBusSenderFeePerByteUpdatedIterator, error) {

	logs, sub, err := _MessageBusSender.contract.FilterLogs(opts, "FeePerByteUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderFeePerByteUpdatedIterator{contract: _MessageBusSender.contract, event: "FeePerByteUpdated", logs: logs, sub: sub}, nil
}

// WatchFeePerByteUpdated is a free log subscription operation binding the contract event 0x210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c.
//
// Solidity: event FeePerByteUpdated(uint256 feePerByte)
func (_MessageBusSender *MessageBusSenderFilterer) WatchFeePerByteUpdated(opts *bind.WatchOpts, sink chan<- *MessageBusSenderFeePerByteUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageBusSender.contract.WatchLogs(opts, "FeePerByteUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderFeePerByteUpdated)
				if err := _MessageBusSender.contract.UnpackLog(event, "FeePerByteUpdated", log); err != nil {
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

// ParseFeePerByteUpdated is a log parse operation binding the contract event 0x210d4d5d2d36d571207dac98e383e2441c684684c885fb2d7c54f8d24422074c.
//
// Solidity: event FeePerByteUpdated(uint256 feePerByte)
func (_MessageBusSender *MessageBusSenderFilterer) ParseFeePerByteUpdated(log types.Log) (*MessageBusSenderFeePerByteUpdated, error) {
	event := new(MessageBusSenderFeePerByteUpdated)
	if err := _MessageBusSender.contract.UnpackLog(event, "FeePerByteUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderFeeWithdrawnIterator is returned from FilterFeeWithdrawn and is used to iterate over the raw logs and unpacked data for FeeWithdrawn events raised by the MessageBusSender contract.
type MessageBusSenderFeeWithdrawnIterator struct {
	Event *MessageBusSenderFeeWithdrawn // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderFeeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderFeeWithdrawn)
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
		it.Event = new(MessageBusSenderFeeWithdrawn)
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
func (it *MessageBusSenderFeeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderFeeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderFeeWithdrawn represents a FeeWithdrawn event raised by the MessageBusSender contract.
type MessageBusSenderFeeWithdrawn struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeeWithdrawn is a free log retrieval operation binding the contract event 0x78473f3f373f7673597f4f0fa5873cb4d375fea6d4339ad6b56dbd411513cb3f.
//
// Solidity: event FeeWithdrawn(address receiver, uint256 amount)
func (_MessageBusSender *MessageBusSenderFilterer) FilterFeeWithdrawn(opts *bind.FilterOpts) (*MessageBusSenderFeeWithdrawnIterator, error) {

	logs, sub, err := _MessageBusSender.contract.FilterLogs(opts, "FeeWithdrawn")
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderFeeWithdrawnIterator{contract: _MessageBusSender.contract, event: "FeeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFeeWithdrawn is a free log subscription operation binding the contract event 0x78473f3f373f7673597f4f0fa5873cb4d375fea6d4339ad6b56dbd411513cb3f.
//
// Solidity: event FeeWithdrawn(address receiver, uint256 amount)
func (_MessageBusSender *MessageBusSenderFilterer) WatchFeeWithdrawn(opts *bind.WatchOpts, sink chan<- *MessageBusSenderFeeWithdrawn) (event.Subscription, error) {

	logs, sub, err := _MessageBusSender.contract.WatchLogs(opts, "FeeWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderFeeWithdrawn)
				if err := _MessageBusSender.contract.UnpackLog(event, "FeeWithdrawn", log); err != nil {
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

// ParseFeeWithdrawn is a log parse operation binding the contract event 0x78473f3f373f7673597f4f0fa5873cb4d375fea6d4339ad6b56dbd411513cb3f.
//
// Solidity: event FeeWithdrawn(address receiver, uint256 amount)
func (_MessageBusSender *MessageBusSenderFilterer) ParseFeeWithdrawn(log types.Log) (*MessageBusSenderFeeWithdrawn, error) {
	event := new(MessageBusSenderFeeWithdrawn)
	if err := _MessageBusSender.contract.UnpackLog(event, "FeeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderMessageIterator is returned from FilterMessage and is used to iterate over the raw logs and unpacked data for Message events raised by the MessageBusSender contract.
type MessageBusSenderMessageIterator struct {
	Event *MessageBusSenderMessage // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderMessageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderMessage)
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
		it.Event = new(MessageBusSenderMessage)
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
func (it *MessageBusSenderMessageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderMessageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderMessage represents a Message event raised by the MessageBusSender contract.
type MessageBusSenderMessage struct {
	Sender     common.Address
	Receiver   common.Address
	DstChainId *big.Int
	Message    []byte
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessage is a free log retrieval operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) FilterMessage(opts *bind.FilterOpts, sender []common.Address) (*MessageBusSenderMessageIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBusSender.contract.FilterLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderMessageIterator{contract: _MessageBusSender.contract, event: "Message", logs: logs, sub: sub}, nil
}

// WatchMessage is a free log subscription operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) WatchMessage(opts *bind.WatchOpts, sink chan<- *MessageBusSenderMessage, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBusSender.contract.WatchLogs(opts, "Message", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderMessage)
				if err := _MessageBusSender.contract.UnpackLog(event, "Message", log); err != nil {
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

// ParseMessage is a log parse operation binding the contract event 0xce3972bfffe49d317e1d128047a97a3d86b25c94f6f04409f988ef854d25e0e4.
//
// Solidity: event Message(address indexed sender, address receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) ParseMessage(log types.Log) (*MessageBusSenderMessage, error) {
	event := new(MessageBusSenderMessage)
	if err := _MessageBusSender.contract.UnpackLog(event, "Message", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderMessage2Iterator is returned from FilterMessage2 and is used to iterate over the raw logs and unpacked data for Message2 events raised by the MessageBusSender contract.
type MessageBusSenderMessage2Iterator struct {
	Event *MessageBusSenderMessage2 // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderMessage2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderMessage2)
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
		it.Event = new(MessageBusSenderMessage2)
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
func (it *MessageBusSenderMessage2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderMessage2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderMessage2 represents a Message2 event raised by the MessageBusSender contract.
type MessageBusSenderMessage2 struct {
	Sender     common.Address
	Receiver   []byte
	DstChainId *big.Int
	Message    []byte
	Fee        *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessage2 is a free log retrieval operation binding the contract event 0xe66fbe37d84ca73c589f782ac278844918ea6c56a4917f58707f715588080df2.
//
// Solidity: event Message2(address indexed sender, bytes receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) FilterMessage2(opts *bind.FilterOpts, sender []common.Address) (*MessageBusSenderMessage2Iterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBusSender.contract.FilterLogs(opts, "Message2", senderRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderMessage2Iterator{contract: _MessageBusSender.contract, event: "Message2", logs: logs, sub: sub}, nil
}

// WatchMessage2 is a free log subscription operation binding the contract event 0xe66fbe37d84ca73c589f782ac278844918ea6c56a4917f58707f715588080df2.
//
// Solidity: event Message2(address indexed sender, bytes receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) WatchMessage2(opts *bind.WatchOpts, sink chan<- *MessageBusSenderMessage2, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBusSender.contract.WatchLogs(opts, "Message2", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderMessage2)
				if err := _MessageBusSender.contract.UnpackLog(event, "Message2", log); err != nil {
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

// ParseMessage2 is a log parse operation binding the contract event 0xe66fbe37d84ca73c589f782ac278844918ea6c56a4917f58707f715588080df2.
//
// Solidity: event Message2(address indexed sender, bytes receiver, uint256 dstChainId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) ParseMessage2(log types.Log) (*MessageBusSenderMessage2, error) {
	event := new(MessageBusSenderMessage2)
	if err := _MessageBusSender.contract.UnpackLog(event, "Message2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderMessageWithTransferIterator is returned from FilterMessageWithTransfer and is used to iterate over the raw logs and unpacked data for MessageWithTransfer events raised by the MessageBusSender contract.
type MessageBusSenderMessageWithTransferIterator struct {
	Event *MessageBusSenderMessageWithTransfer // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderMessageWithTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderMessageWithTransfer)
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
		it.Event = new(MessageBusSenderMessageWithTransfer)
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
func (it *MessageBusSenderMessageWithTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderMessageWithTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderMessageWithTransfer represents a MessageWithTransfer event raised by the MessageBusSender contract.
type MessageBusSenderMessageWithTransfer struct {
	Sender        common.Address
	Receiver      common.Address
	DstChainId    *big.Int
	Bridge        common.Address
	SrcTransferId [32]byte
	Message       []byte
	Fee           *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMessageWithTransfer is a free log retrieval operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) FilterMessageWithTransfer(opts *bind.FilterOpts, sender []common.Address) (*MessageBusSenderMessageWithTransferIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBusSender.contract.FilterLogs(opts, "MessageWithTransfer", senderRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderMessageWithTransferIterator{contract: _MessageBusSender.contract, event: "MessageWithTransfer", logs: logs, sub: sub}, nil
}

// WatchMessageWithTransfer is a free log subscription operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) WatchMessageWithTransfer(opts *bind.WatchOpts, sink chan<- *MessageBusSenderMessageWithTransfer, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MessageBusSender.contract.WatchLogs(opts, "MessageWithTransfer", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderMessageWithTransfer)
				if err := _MessageBusSender.contract.UnpackLog(event, "MessageWithTransfer", log); err != nil {
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

// ParseMessageWithTransfer is a log parse operation binding the contract event 0x172762498a59a3bc4fed3f2b63f94f17ea0193cffdc304fe7d3eaf4d342d2f66.
//
// Solidity: event MessageWithTransfer(address indexed sender, address receiver, uint256 dstChainId, address bridge, bytes32 srcTransferId, bytes message, uint256 fee)
func (_MessageBusSender *MessageBusSenderFilterer) ParseMessageWithTransfer(log types.Log) (*MessageBusSenderMessageWithTransfer, error) {
	event := new(MessageBusSenderMessageWithTransfer)
	if err := _MessageBusSender.contract.UnpackLog(event, "MessageWithTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageBusSenderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageBusSender contract.
type MessageBusSenderOwnershipTransferredIterator struct {
	Event *MessageBusSenderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageBusSenderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageBusSenderOwnershipTransferred)
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
		it.Event = new(MessageBusSenderOwnershipTransferred)
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
func (it *MessageBusSenderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageBusSenderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageBusSenderOwnershipTransferred represents a OwnershipTransferred event raised by the MessageBusSender contract.
type MessageBusSenderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusSender *MessageBusSenderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageBusSenderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusSender.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageBusSenderOwnershipTransferredIterator{contract: _MessageBusSender.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusSender *MessageBusSenderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageBusSenderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageBusSender.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageBusSenderOwnershipTransferred)
				if err := _MessageBusSender.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageBusSender *MessageBusSenderFilterer) ParseOwnershipTransferred(log types.Log) (*MessageBusSenderOwnershipTransferred, error) {
	event := new(MessageBusSenderOwnershipTransferred)
	if err := _MessageBusSender.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MsgDataTypesMetaData contains all meta data concerning the MsgDataTypes contract.
var MsgDataTypesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220303fcbe9493c87e94d350e1ea288681823a2766ea6623c40286be19f266c7f2264736f6c63430008110033",
}

// MsgDataTypesABI is the input ABI used to generate the binding from.
// Deprecated: Use MsgDataTypesMetaData.ABI instead.
var MsgDataTypesABI = MsgDataTypesMetaData.ABI

// MsgDataTypesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MsgDataTypesMetaData.Bin instead.
var MsgDataTypesBin = MsgDataTypesMetaData.Bin

// DeployMsgDataTypes deploys a new Ethereum contract, binding an instance of MsgDataTypes to it.
func DeployMsgDataTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MsgDataTypes, error) {
	parsed, err := MsgDataTypesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MsgDataTypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MsgDataTypes{MsgDataTypesCaller: MsgDataTypesCaller{contract: contract}, MsgDataTypesTransactor: MsgDataTypesTransactor{contract: contract}, MsgDataTypesFilterer: MsgDataTypesFilterer{contract: contract}}, nil
}

// MsgDataTypes is an auto generated Go binding around an Ethereum contract.
type MsgDataTypes struct {
	MsgDataTypesCaller     // Read-only binding to the contract
	MsgDataTypesTransactor // Write-only binding to the contract
	MsgDataTypesFilterer   // Log filterer for contract events
}

// MsgDataTypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type MsgDataTypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgDataTypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MsgDataTypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgDataTypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MsgDataTypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MsgDataTypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MsgDataTypesSession struct {
	Contract     *MsgDataTypes     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MsgDataTypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MsgDataTypesCallerSession struct {
	Contract *MsgDataTypesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MsgDataTypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MsgDataTypesTransactorSession struct {
	Contract     *MsgDataTypesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MsgDataTypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type MsgDataTypesRaw struct {
	Contract *MsgDataTypes // Generic contract binding to access the raw methods on
}

// MsgDataTypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MsgDataTypesCallerRaw struct {
	Contract *MsgDataTypesCaller // Generic read-only contract binding to access the raw methods on
}

// MsgDataTypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MsgDataTypesTransactorRaw struct {
	Contract *MsgDataTypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMsgDataTypes creates a new instance of MsgDataTypes, bound to a specific deployed contract.
func NewMsgDataTypes(address common.Address, backend bind.ContractBackend) (*MsgDataTypes, error) {
	contract, err := bindMsgDataTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MsgDataTypes{MsgDataTypesCaller: MsgDataTypesCaller{contract: contract}, MsgDataTypesTransactor: MsgDataTypesTransactor{contract: contract}, MsgDataTypesFilterer: MsgDataTypesFilterer{contract: contract}}, nil
}

// NewMsgDataTypesCaller creates a new read-only instance of MsgDataTypes, bound to a specific deployed contract.
func NewMsgDataTypesCaller(address common.Address, caller bind.ContractCaller) (*MsgDataTypesCaller, error) {
	contract, err := bindMsgDataTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MsgDataTypesCaller{contract: contract}, nil
}

// NewMsgDataTypesTransactor creates a new write-only instance of MsgDataTypes, bound to a specific deployed contract.
func NewMsgDataTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*MsgDataTypesTransactor, error) {
	contract, err := bindMsgDataTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MsgDataTypesTransactor{contract: contract}, nil
}

// NewMsgDataTypesFilterer creates a new log filterer instance of MsgDataTypes, bound to a specific deployed contract.
func NewMsgDataTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*MsgDataTypesFilterer, error) {
	contract, err := bindMsgDataTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MsgDataTypesFilterer{contract: contract}, nil
}

// bindMsgDataTypes binds a generic wrapper to an already deployed contract.
func bindMsgDataTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MsgDataTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MsgDataTypes *MsgDataTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MsgDataTypes.Contract.MsgDataTypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MsgDataTypes *MsgDataTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MsgDataTypes.Contract.MsgDataTypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MsgDataTypes *MsgDataTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MsgDataTypes.Contract.MsgDataTypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MsgDataTypes *MsgDataTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MsgDataTypes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MsgDataTypes *MsgDataTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MsgDataTypes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MsgDataTypes *MsgDataTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MsgDataTypes.Contract.contract.Transact(opts, method, params...)
}

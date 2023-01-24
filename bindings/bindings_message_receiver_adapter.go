// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// MessageReceiverAdapterMetaData contains all meta data concerning the MessageReceiverAdapter contract.
var MessageReceiverAdapterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AllowedSenderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"DelayPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"DelayedMessageAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"DelayedMessageExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"messageBus\",\"type\":\"address\"}],\"name\":\"MessageBusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"srcContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dstContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"delayedMessages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_srcContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_dstContract\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callData\",\"type\":\"bytes\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"}],\"name\":\"executeDelayedMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_sender\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_srcContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransfer\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransferFallback\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessageWithTransferRefund\",\"outputs\":[{\"internalType\":\"enumIMessageReceiverApp.ExecutionStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dstContract\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"_srcChainId\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"_srcContracts\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_alloweds\",\"type\":\"bool[]\"}],\"name\":\"setAllowedSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"setDelayPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200246338038062002463833981016040819052620000349162000264565b8062000040336200008a565b600180546001600160a01b0319166001600160a01b03929092169190911790556200006b33620000da565b6005805460ff60201b191690556200008333620001a4565b5062000296565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b03811660009081526002602052604090205460ff1615620001495760405162461bcd60e51b815260206004820152601b60248201527f4163636f756e7420697320616c726561647920676f7665726e6f72000000000060448201526064015b60405180910390fd5b6001600160a01b038116600081815260026020908152604091829020805460ff1916600117905590519182527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b591015b60405180910390a150565b6001600160a01b03811660009081526006602052604090205460ff16156200020f5760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c72656164792070617573657200000000000000604482015260640162000140565b6001600160a01b038116600081815260066020908152604091829020805460ff1916600117905590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8910162000199565b6000602082840312156200027757600080fd5b81516001600160a01b03811681146200028f57600080fd5b9392505050565b6121bd80620002a66000396000f3fe6080604052600436106101c25760003560e01c806380f51c12116100f7578063affed0e011610095578063e3eece2611610064578063e3eece26146104fd578063e43581b81461052d578063eecdac8814610566578063f2fde38b1461058657600080fd5b8063affed0e01461048d578063b1c94d94146104bf578063e026049c146104d5578063e0bfeaad146104ea57600080fd5b80638456cb59116100d15780638456cb59146104135780638da5cb5b146104285780639c649fdf1461045a578063a1a227fa1461046d57600080fd5b806380f51c12146103a357806382dc1ec4146103d357806383960ac7146103f357600080fd5b806346fbf68e116101645780635c975abb1161013e5780635c975abb1461034e5780636b2c0f551461036e5780636ef8d66d1461038e5780637cd2bffc1461033b57600080fd5b806346fbf68e146102e2578063547cad121461031b5780635ab7afc61461033b57600080fd5b80631d8f2426116101a05780631d8f2426146102505780633c4a25d01461028b5780633d572107146102ad5780633f4ba83a146102cd57600080fd5b806306302c94146101c7578063063ce4e51461021d5780630bcb49821461023d575b600080fd5b3480156101d357600080fd5b506102086101e2366004611910565b600760209081526000938452604080852082529284528284209052825290205460ff1681565b60405190151581526020015b60405180910390f35b61023061022b3660046119a2565b6105a6565b6040516102149190611a38565b61023061024b366004611a60565b610612565b34801561025c57600080fd5b5061027d61026b366004611ad4565b60036020526000908152604090205481565b604051908152602001610214565b34801561029757600080fd5b506102ab6102a6366004611aed565b610678565b005b3480156102b957600080fd5b506102ab6102c8366004611ad4565b6106ed565b3480156102d957600080fd5b506102ab610788565b3480156102ee57600080fd5b506102086102fd366004611aed565b6001600160a01b031660009081526006602052604090205460ff1690565b34801561032757600080fd5b506102ab610336366004611aed565b6107f1565b610230610349366004611b11565b6108b5565b34801561035a57600080fd5b50600554640100000000900460ff16610208565b34801561037a57600080fd5b506102ab610389366004611aed565b61091d565b34801561039a57600080fd5b506102ab61098f565b3480156103af57600080fd5b506102086103be366004611aed565b60066020526000908152604090205460ff1681565b3480156103df57600080fd5b506102ab6103ee366004611aed565b610998565b3480156103ff57600080fd5b506102ab61040e366004611bed565b610a0a565b34801561041f57600080fd5b506102ab610c3b565b34801561043457600080fd5b506000546001600160a01b03165b6040516001600160a01b039091168152602001610214565b610230610468366004611c80565b610ca2565b34801561047957600080fd5b50600154610442906001600160a01b031681565b34801561049957600080fd5b506005546104aa9063ffffffff1681565b60405163ffffffff9091168152602001610214565b3480156104cb57600080fd5b5061027d60045481565b3480156104e157600080fd5b506102ab610e71565b6102ab6104f8366004611ccd565b610e7a565b34801561050957600080fd5b50610208610518366004611aed565b60026020526000908152604090205460ff1681565b34801561053957600080fd5b50610208610548366004611aed565b6001600160a01b031660009081526002602052604090205460ff1690565b34801561057257600080fd5b506102ab610581366004611aed565b610f94565b34801561059257600080fd5b506102ab6105a1366004611aed565b611006565b6001546000906001600160a01b031633146106085760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064015b60405180910390fd5b9695505050505050565b6001546000906001600160a01b0316331461066f5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064016105ff565b95945050505050565b3361068b6000546001600160a01b031690565b6001600160a01b0316146106e15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105ff565b6106ea816110f4565b50565b3360009081526002602052604090205460ff1661074c5760405162461bcd60e51b815260206004820152601660248201527f43616c6c6572206973206e6f7420676f7665726e6f720000000000000000000060448201526064016105ff565b60048190556040518181527fc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6906020015b60405180910390a150565b3360009081526006602052604090205460ff166107e75760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f742070617573657200000000000000000000000060448201526064016105ff565b6107ef6111b1565b565b336108046000546001600160a01b031690565b6001600160a01b03161461085a5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105ff565b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f3f8223bcd8b3b875473e9f9e14e1ad075451a2b5ffd31591655da9a01516bf5e9060200161077d565b6001546000906001600160a01b031633146109125760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064016105ff565b979650505050505050565b336109306000546001600160a01b031690565b6001600160a01b0316146109865760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105ff565b6106ea81611259565b6107ef33611259565b336109ab6000546001600160a01b031690565b6001600160a01b031614610a015760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105ff565b6106ea81611312565b33610a1d6000546001600160a01b031690565b6001600160a01b031614610a735760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105ff565b828114610ac25760405162461bcd60e51b815260206004820152600f60248201527f6c656e677468206d69736d61746368000000000000000000000000000000000060448201526064016105ff565b60005b83811015610c3257828282818110610adf57610adf611d54565b9050602002016020810190610af49190611d6a565b6001600160a01b038816600090815260076020908152604080832067ffffffffffffffff8b168452909152812090878785818110610b3457610b34611d54565b9050602002016020810190610b499190611aed565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790557fa88a752d56e7264efad752e7313e00c0e04a332b664a44e2eceaf477674ab7d78787878785818110610ba657610ba6611d54565b9050602002016020810190610bbb9190611aed565b868686818110610bcd57610bcd611d54565b9050602002016020810190610be29190611d6a565b604080516001600160a01b03958616815267ffffffffffffffff9490941660208501529190931690820152901515606082015260800160405180910390a180610c2a81611da2565b915050610ac5565b50505050505050565b3360009081526006602052604090205460ff16610c9a5760405162461bcd60e51b815260206004820152601460248201527f43616c6c6572206973206e6f742070617573657200000000000000000000000060448201526064016105ff565b6107ef6113cf565b6001546000906001600160a01b03163314610cff5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d657373616765206275730000000000000060448201526064016105ff565b600554640100000000900460ff1615610d4d5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016105ff565b600080610d5c85870187611e2a565b6001600160a01b03808316600090815260076020908152604080832067ffffffffffffffff8e1684528252808320938e1683529290522054919350915060ff16610de85760405162461bcd60e51b815260206004820152601260248201527f6e6f7420616c6c6f7765642073656e646572000000000000000000000000000060448201526064016105ff565b60008060045411610dfa576000610dfd565b60015b90508015610e1d57610e118989858561145a565b6001935050505061066f565b7f89e5332ec6ff132178f395753d055e19755ab6620d41de1c27e0a135a54f6d6789898585604051610e529493929190611f0d565b60405180910390a1610e648383611592565b9998505050505050505050565b6107ef3361162b565b600554640100000000900460ff1615610ec85760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016105ff565b610f0c86868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152508892506116e4915050565b610f4c8484848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061159292505050565b507f89e5332ec6ff132178f395753d055e19755ab6620d41de1c27e0a135a54f6d678686868686604051610f84959493929190611f49565b60405180910390a1505050505050565b33610fa76000546001600160a01b031690565b6001600160a01b031614610ffd5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105ff565b6106ea8161162b565b336110196000546001600160a01b031690565b6001600160a01b03161461106f5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016105ff565b6001600160a01b0381166110eb5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016105ff565b6106ea81611822565b6001600160a01b03811660009081526002602052604090205460ff161561115d5760405162461bcd60e51b815260206004820152601b60248201527f4163636f756e7420697320616c726561647920676f7665726e6f72000000000060448201526064016105ff565b6001600160a01b038116600081815260026020908152604091829020805460ff1916600117905590519182527fdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5910161077d565b600554640100000000900460ff1661120b5760405162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f742070617573656400000000000000000000000060448201526064016105ff565b6005805464ff00000000191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a1565b6001600160a01b03811660009081526006602052604090205460ff166112c15760405162461bcd60e51b815260206004820152601560248201527f4163636f756e74206973206e6f7420706175736572000000000000000000000060448201526064016105ff565b6001600160a01b038116600081815260066020908152604091829020805460ff1916905590519182527fcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e910161077d565b6001600160a01b03811660009081526006602052604090205460ff161561137b5760405162461bcd60e51b815260206004820152601960248201527f4163636f756e7420697320616c7265616479207061757365720000000000000060448201526064016105ff565b6001600160a01b038116600081815260066020908152604091829020805460ff1916600117905590519182527f6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8910161077d565b600554640100000000900460ff161561141d5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016105ff565b6005805464ff0000000019166401000000001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861123c3390565b600554604051600091611480918791879187914691889163ffffffff1690602001611fa7565b60408051601f19818403018152918152815160209283012060008181526003909352912054909150156114f55760405162461bcd60e51b815260206004820152601e60248201527f64656c61796564206d65737361676520616c726561647920657869737473000060448201526064016105ff565b6000818152600360205260409081902042905560055490517fb86a62c31cff12d722cd0a54f5964552dfe3fc694470b892aebf3511ea60d1be91611549918491899189918991899163ffffffff1690612058565b60405180910390a1600580546001919060009061156d90849063ffffffff166120b3565b92506101000a81548163ffffffff021916908363ffffffff1602179055505050505050565b6000806000846001600160a01b031634856040516115b091906120d7565b60006040518083038185875af1925050503d80600081146115ed576040519150601f19603f3d011682016040523d82523d6000602084013e6115f2565b606091505b50915091508161161e576116058161187f565b60405162461bcd60e51b81526004016105ff91906120f3565b6001925050505b92915050565b6001600160a01b03811660009081526002602052604090205460ff166116935760405162461bcd60e51b815260206004820152601760248201527f4163636f756e74206973206e6f7420676f7665726e6f7200000000000000000060448201526064016105ff565b6001600160a01b038116600081815260026020908152604091829020805460ff1916905590519182527f1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b910161077d565b600085858546868660405160200161170196959493929190611fa7565b60408051601f198184030181529181528151602092830120600081815260039093529120549091506117755760405162461bcd60e51b815260206004820152601960248201527f64656c61796564206d657373616765206e6f742065786973740000000000000060448201526064016105ff565b6004546000828152600360205260409020546117919190612106565b42116117df5760405162461bcd60e51b815260206004820152601c60248201527f64656c61796564206d657373616765207374696c6c206c6f636b65640000000060448201526064016105ff565b60008181526003602052604080822091909155517f3dd29d696ff6a4e0e53f61fecdcfe71ab28e10acd8ff8c59d55744d92b3fa24590610f849083815260200190565b600080546001600160a01b0383811673ffffffffffffffffffffffffffffffffffffffff19831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60606044825110156118c457505060408051808201909152601d81527f5472616e73616374696f6e2072657665727465642073696c656e746c79000000602082015290565b600482019150818060200190518101906116259190612119565b6001600160a01b03811681146106ea57600080fd5b803567ffffffffffffffff8116811461190b57600080fd5b919050565b60008060006060848603121561192557600080fd5b8335611930816118de565b925061193e602085016118f3565b9150604084013561194e816118de565b809150509250925092565b60008083601f84011261196b57600080fd5b50813567ffffffffffffffff81111561198357600080fd5b60208301915083602082850101111561199b57600080fd5b9250929050565b600080600080600080608087890312156119bb57600080fd5b863567ffffffffffffffff808211156119d357600080fd5b6119df8a838b01611959565b90985096508691506119f360208a016118f3565b95506040890135915080821115611a0957600080fd5b50611a1689828a01611959565b9094509250506060870135611a2a816118de565b809150509295509295509295565b6020810160038310611a5a57634e487b7160e01b600052602160045260246000fd5b91905290565b600080600080600060808688031215611a7857600080fd5b8535611a83816118de565b945060208601359350604086013567ffffffffffffffff811115611aa657600080fd5b611ab288828901611959565b9094509250506060860135611ac6816118de565b809150509295509295909350565b600060208284031215611ae657600080fd5b5035919050565b600060208284031215611aff57600080fd5b8135611b0a816118de565b9392505050565b600080600080600080600060c0888a031215611b2c57600080fd5b8735611b37816118de565b96506020880135611b47816118de565b955060408801359450611b5c606089016118f3565b9350608088013567ffffffffffffffff811115611b7857600080fd5b611b848a828b01611959565b90945092505060a0880135611b98816118de565b8091505092959891949750929550565b60008083601f840112611bba57600080fd5b50813567ffffffffffffffff811115611bd257600080fd5b6020830191508360208260051b850101111561199b57600080fd5b60008060008060008060808789031215611c0657600080fd5b8635611c11816118de565b9550611c1f602088016118f3565b9450604087013567ffffffffffffffff80821115611c3c57600080fd5b611c488a838b01611ba8565b90965094506060890135915080821115611c6157600080fd5b50611c6e89828a01611ba8565b979a9699509497509295939492505050565b600080600080600060808688031215611c9857600080fd5b8535611ca3816118de565b9450611cb1602087016118f3565b9350604086013567ffffffffffffffff811115611aa657600080fd5b60008060008060008060a08789031215611ce657600080fd5b8635611cf1816118de565b9550611cff602088016118f3565b94506040870135611d0f816118de565b9350606087013567ffffffffffffffff811115611d2b57600080fd5b611d3789828a01611959565b909450925050608087013563ffffffff81168114611a2a57600080fd5b634e487b7160e01b600052603260045260246000fd5b600060208284031215611d7c57600080fd5b81358015158114611b0a57600080fd5b634e487b7160e01b600052601160045260246000fd5b600060018201611db457611db4611d8c565b5060010190565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611dfa57611dfa611dbb565b604052919050565b600067ffffffffffffffff821115611e1c57611e1c611dbb565b50601f01601f191660200190565b60008060408385031215611e3d57600080fd5b8235611e48816118de565b9150602083013567ffffffffffffffff811115611e6457600080fd5b8301601f81018513611e7557600080fd5b8035611e88611e8382611e02565b611dd1565b818152866020838501011115611e9d57600080fd5b816020840160208301376000602083830101528093505050509250929050565b60005b83811015611ed8578181015183820152602001611ec0565b50506000910152565b60008151808452611ef9816020860160208601611ebd565b601f01601f19169290920160200192915050565b60006001600160a01b03808716835267ffffffffffffffff86166020840152808516604084015250608060608301526106086080830184611ee1565b60006001600160a01b03808816835267ffffffffffffffff8716602084015280861660408401525060806060830152826080830152828460a0840137600060a0848401015260a0601f19601f85011683010190509695505050505050565b60006bffffffffffffffffffffffff19808960601b1683527fffffffffffffffff000000000000000000000000000000000000000000000000808960c01b166014850152818860601b16601c850152808760c01b16603085015250508351612016816038850160208801611ebd565b60e09390931b7fffffffff000000000000000000000000000000000000000000000000000000001660389290930191820192909252603c019695505050505050565b86815260006001600160a01b03808816602084015267ffffffffffffffff8716604084015280861660608401525060c0608083015261209a60c0830185611ee1565b905063ffffffff831660a0830152979650505050505050565b63ffffffff8181168382160190808211156120d0576120d0611d8c565b5092915050565b600082516120e9818460208701611ebd565b9190910192915050565b602081526000611b0a6020830184611ee1565b8082018082111561162557611625611d8c565b60006020828403121561212b57600080fd5b815167ffffffffffffffff81111561214257600080fd5b8201601f8101841361215357600080fd5b8051612161611e8382611e02565b81815285602083850101111561217657600080fd5b61066f826020830160208601611ebd56fea26469706673582212201e6036b70b0a64d7c239f001f713c11fe5a2ded162bb3fc73e75c2a272125bb064736f6c63430008110033",
}

// MessageReceiverAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageReceiverAdapterMetaData.ABI instead.
var MessageReceiverAdapterABI = MessageReceiverAdapterMetaData.ABI

// MessageReceiverAdapterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageReceiverAdapterMetaData.Bin instead.
var MessageReceiverAdapterBin = MessageReceiverAdapterMetaData.Bin

// DeployMessageReceiverAdapter deploys a new Ethereum contract, binding an instance of MessageReceiverAdapter to it.
func DeployMessageReceiverAdapter(auth *bind.TransactOpts, backend bind.ContractBackend, _messageBus common.Address) (common.Address, *types.Transaction, *MessageReceiverAdapter, error) {
	parsed, err := MessageReceiverAdapterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageReceiverAdapterBin), backend, _messageBus)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageReceiverAdapter{MessageReceiverAdapterCaller: MessageReceiverAdapterCaller{contract: contract}, MessageReceiverAdapterTransactor: MessageReceiverAdapterTransactor{contract: contract}, MessageReceiverAdapterFilterer: MessageReceiverAdapterFilterer{contract: contract}}, nil
}

// MessageReceiverAdapter is an auto generated Go binding around an Ethereum contract.
type MessageReceiverAdapter struct {
	MessageReceiverAdapterCaller     // Read-only binding to the contract
	MessageReceiverAdapterTransactor // Write-only binding to the contract
	MessageReceiverAdapterFilterer   // Log filterer for contract events
}

// MessageReceiverAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageReceiverAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageReceiverAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageReceiverAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageReceiverAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageReceiverAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageReceiverAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageReceiverAdapterSession struct {
	Contract     *MessageReceiverAdapter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MessageReceiverAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageReceiverAdapterCallerSession struct {
	Contract *MessageReceiverAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MessageReceiverAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageReceiverAdapterTransactorSession struct {
	Contract     *MessageReceiverAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MessageReceiverAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageReceiverAdapterRaw struct {
	Contract *MessageReceiverAdapter // Generic contract binding to access the raw methods on
}

// MessageReceiverAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageReceiverAdapterCallerRaw struct {
	Contract *MessageReceiverAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// MessageReceiverAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageReceiverAdapterTransactorRaw struct {
	Contract *MessageReceiverAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageReceiverAdapter creates a new instance of MessageReceiverAdapter, bound to a specific deployed contract.
func NewMessageReceiverAdapter(address common.Address, backend bind.ContractBackend) (*MessageReceiverAdapter, error) {
	contract, err := bindMessageReceiverAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapter{MessageReceiverAdapterCaller: MessageReceiverAdapterCaller{contract: contract}, MessageReceiverAdapterTransactor: MessageReceiverAdapterTransactor{contract: contract}, MessageReceiverAdapterFilterer: MessageReceiverAdapterFilterer{contract: contract}}, nil
}

// NewMessageReceiverAdapterCaller creates a new read-only instance of MessageReceiverAdapter, bound to a specific deployed contract.
func NewMessageReceiverAdapterCaller(address common.Address, caller bind.ContractCaller) (*MessageReceiverAdapterCaller, error) {
	contract, err := bindMessageReceiverAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterCaller{contract: contract}, nil
}

// NewMessageReceiverAdapterTransactor creates a new write-only instance of MessageReceiverAdapter, bound to a specific deployed contract.
func NewMessageReceiverAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageReceiverAdapterTransactor, error) {
	contract, err := bindMessageReceiverAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterTransactor{contract: contract}, nil
}

// NewMessageReceiverAdapterFilterer creates a new log filterer instance of MessageReceiverAdapter, bound to a specific deployed contract.
func NewMessageReceiverAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageReceiverAdapterFilterer, error) {
	contract, err := bindMessageReceiverAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterFilterer{contract: contract}, nil
}

// bindMessageReceiverAdapter binds a generic wrapper to an already deployed contract.
func bindMessageReceiverAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageReceiverAdapterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageReceiverAdapter *MessageReceiverAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageReceiverAdapter.Contract.MessageReceiverAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageReceiverAdapter *MessageReceiverAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.MessageReceiverAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageReceiverAdapter *MessageReceiverAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.MessageReceiverAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageReceiverAdapter *MessageReceiverAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageReceiverAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.contract.Transact(opts, method, params...)
}

// AllowedSender is a free data retrieval call binding the contract method 0x06302c94.
//
// Solidity: function allowedSender(address , uint64 , address ) view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterCaller) AllowedSender(opts *bind.CallOpts, arg0 common.Address, arg1 uint64, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _MessageReceiverAdapter.contract.Call(opts, &out, "allowedSender", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedSender is a free data retrieval call binding the contract method 0x06302c94.
//
// Solidity: function allowedSender(address , uint64 , address ) view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) AllowedSender(arg0 common.Address, arg1 uint64, arg2 common.Address) (bool, error) {
	return _MessageReceiverAdapter.Contract.AllowedSender(&_MessageReceiverAdapter.CallOpts, arg0, arg1, arg2)
}

// AllowedSender is a free data retrieval call binding the contract method 0x06302c94.
//
// Solidity: function allowedSender(address , uint64 , address ) view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterCallerSession) AllowedSender(arg0 common.Address, arg1 uint64, arg2 common.Address) (bool, error) {
	return _MessageReceiverAdapter.Contract.AllowedSender(&_MessageReceiverAdapter.CallOpts, arg0, arg1, arg2)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_MessageReceiverAdapter *MessageReceiverAdapterCaller) DelayPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageReceiverAdapter.contract.Call(opts, &out, "delayPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) DelayPeriod() (*big.Int, error) {
	return _MessageReceiverAdapter.Contract.DelayPeriod(&_MessageReceiverAdapter.CallOpts)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_MessageReceiverAdapter *MessageReceiverAdapterCallerSession) DelayPeriod() (*big.Int, error) {
	return _MessageReceiverAdapter.Contract.DelayPeriod(&_MessageReceiverAdapter.CallOpts)
}

// DelayedMessages is a free data retrieval call binding the contract method 0x1d8f2426.
//
// Solidity: function delayedMessages(bytes32 ) view returns(uint256)
func (_MessageReceiverAdapter *MessageReceiverAdapterCaller) DelayedMessages(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MessageReceiverAdapter.contract.Call(opts, &out, "delayedMessages", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayedMessages is a free data retrieval call binding the contract method 0x1d8f2426.
//
// Solidity: function delayedMessages(bytes32 ) view returns(uint256)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) DelayedMessages(arg0 [32]byte) (*big.Int, error) {
	return _MessageReceiverAdapter.Contract.DelayedMessages(&_MessageReceiverAdapter.CallOpts, arg0)
}

// DelayedMessages is a free data retrieval call binding the contract method 0x1d8f2426.
//
// Solidity: function delayedMessages(bytes32 ) view returns(uint256)
func (_MessageReceiverAdapter *MessageReceiverAdapterCallerSession) DelayedMessages(arg0 [32]byte) (*big.Int, error) {
	return _MessageReceiverAdapter.Contract.DelayedMessages(&_MessageReceiverAdapter.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterCaller) Governors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MessageReceiverAdapter.contract.Call(opts, &out, "governors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) Governors(arg0 common.Address) (bool, error) {
	return _MessageReceiverAdapter.Contract.Governors(&_MessageReceiverAdapter.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterCallerSession) Governors(arg0 common.Address) (bool, error) {
	return _MessageReceiverAdapter.Contract.Governors(&_MessageReceiverAdapter.CallOpts, arg0)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterCaller) IsGovernor(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _MessageReceiverAdapter.contract.Call(opts, &out, "isGovernor", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) IsGovernor(_account common.Address) (bool, error) {
	return _MessageReceiverAdapter.Contract.IsGovernor(&_MessageReceiverAdapter.CallOpts, _account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterCallerSession) IsGovernor(_account common.Address) (bool, error) {
	return _MessageReceiverAdapter.Contract.IsGovernor(&_MessageReceiverAdapter.CallOpts, _account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _MessageReceiverAdapter.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) IsPauser(account common.Address) (bool, error) {
	return _MessageReceiverAdapter.Contract.IsPauser(&_MessageReceiverAdapter.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterCallerSession) IsPauser(account common.Address) (bool, error) {
	return _MessageReceiverAdapter.Contract.IsPauser(&_MessageReceiverAdapter.CallOpts, account)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_MessageReceiverAdapter *MessageReceiverAdapterCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageReceiverAdapter.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) MessageBus() (common.Address, error) {
	return _MessageReceiverAdapter.Contract.MessageBus(&_MessageReceiverAdapter.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_MessageReceiverAdapter *MessageReceiverAdapterCallerSession) MessageBus() (common.Address, error) {
	return _MessageReceiverAdapter.Contract.MessageBus(&_MessageReceiverAdapter.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_MessageReceiverAdapter *MessageReceiverAdapterCaller) Nonce(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MessageReceiverAdapter.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) Nonce() (uint32, error) {
	return _MessageReceiverAdapter.Contract.Nonce(&_MessageReceiverAdapter.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint32)
func (_MessageReceiverAdapter *MessageReceiverAdapterCallerSession) Nonce() (uint32, error) {
	return _MessageReceiverAdapter.Contract.Nonce(&_MessageReceiverAdapter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageReceiverAdapter *MessageReceiverAdapterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageReceiverAdapter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) Owner() (common.Address, error) {
	return _MessageReceiverAdapter.Contract.Owner(&_MessageReceiverAdapter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageReceiverAdapter *MessageReceiverAdapterCallerSession) Owner() (common.Address, error) {
	return _MessageReceiverAdapter.Contract.Owner(&_MessageReceiverAdapter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MessageReceiverAdapter.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) Paused() (bool, error) {
	return _MessageReceiverAdapter.Contract.Paused(&_MessageReceiverAdapter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterCallerSession) Paused() (bool, error) {
	return _MessageReceiverAdapter.Contract.Paused(&_MessageReceiverAdapter.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MessageReceiverAdapter.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) Pausers(arg0 common.Address) (bool, error) {
	return _MessageReceiverAdapter.Contract.Pausers(&_MessageReceiverAdapter.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_MessageReceiverAdapter *MessageReceiverAdapterCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _MessageReceiverAdapter.Contract.Pausers(&_MessageReceiverAdapter.CallOpts, arg0)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) AddGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "addGovernor", _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.AddGovernor(&_MessageReceiverAdapter.TransactOpts, _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.AddGovernor(&_MessageReceiverAdapter.TransactOpts, _account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.AddPauser(&_MessageReceiverAdapter.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.AddPauser(&_MessageReceiverAdapter.TransactOpts, account)
}

// ExecuteDelayedMessage is a paid mutator transaction binding the contract method 0xe0bfeaad.
//
// Solidity: function executeDelayedMessage(address _srcContract, uint64 _srcChainId, address _dstContract, bytes _callData, uint32 _nonce) payable returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) ExecuteDelayedMessage(opts *bind.TransactOpts, _srcContract common.Address, _srcChainId uint64, _dstContract common.Address, _callData []byte, _nonce uint32) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "executeDelayedMessage", _srcContract, _srcChainId, _dstContract, _callData, _nonce)
}

// ExecuteDelayedMessage is a paid mutator transaction binding the contract method 0xe0bfeaad.
//
// Solidity: function executeDelayedMessage(address _srcContract, uint64 _srcChainId, address _dstContract, bytes _callData, uint32 _nonce) payable returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) ExecuteDelayedMessage(_srcContract common.Address, _srcChainId uint64, _dstContract common.Address, _callData []byte, _nonce uint32) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.ExecuteDelayedMessage(&_MessageReceiverAdapter.TransactOpts, _srcContract, _srcChainId, _dstContract, _callData, _nonce)
}

// ExecuteDelayedMessage is a paid mutator transaction binding the contract method 0xe0bfeaad.
//
// Solidity: function executeDelayedMessage(address _srcContract, uint64 _srcChainId, address _dstContract, bytes _callData, uint32 _nonce) payable returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) ExecuteDelayedMessage(_srcContract common.Address, _srcChainId uint64, _dstContract common.Address, _callData []byte, _nonce uint32) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.ExecuteDelayedMessage(&_MessageReceiverAdapter.TransactOpts, _srcContract, _srcChainId, _dstContract, _callData, _nonce)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x063ce4e5.
//
// Solidity: function executeMessage(bytes _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) ExecuteMessage(opts *bind.TransactOpts, _sender []byte, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "executeMessage", _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x063ce4e5.
//
// Solidity: function executeMessage(bytes _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) ExecuteMessage(_sender []byte, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.ExecuteMessage(&_MessageReceiverAdapter.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x063ce4e5.
//
// Solidity: function executeMessage(bytes _sender, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) ExecuteMessage(_sender []byte, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.ExecuteMessage(&_MessageReceiverAdapter.TransactOpts, _sender, _srcChainId, _message, _executor)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _srcContract, uint64 _srcChainId, bytes _message, address ) payable returns(uint8)
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) ExecuteMessage0(opts *bind.TransactOpts, _srcContract common.Address, _srcChainId uint64, _message []byte, arg3 common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "executeMessage0", _srcContract, _srcChainId, _message, arg3)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _srcContract, uint64 _srcChainId, bytes _message, address ) payable returns(uint8)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) ExecuteMessage0(_srcContract common.Address, _srcChainId uint64, _message []byte, arg3 common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.ExecuteMessage0(&_MessageReceiverAdapter.TransactOpts, _srcContract, _srcChainId, _message, arg3)
}

// ExecuteMessage0 is a paid mutator transaction binding the contract method 0x9c649fdf.
//
// Solidity: function executeMessage(address _srcContract, uint64 _srcChainId, bytes _message, address ) payable returns(uint8)
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) ExecuteMessage0(_srcContract common.Address, _srcChainId uint64, _message []byte, arg3 common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.ExecuteMessage0(&_MessageReceiverAdapter.TransactOpts, _srcContract, _srcChainId, _message, arg3)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) ExecuteMessageWithTransfer(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "executeMessageWithTransfer", _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.ExecuteMessageWithTransfer(&_MessageReceiverAdapter.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransfer is a paid mutator transaction binding the contract method 0x7cd2bffc.
//
// Solidity: function executeMessageWithTransfer(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) ExecuteMessageWithTransfer(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.ExecuteMessageWithTransfer(&_MessageReceiverAdapter.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) ExecuteMessageWithTransferFallback(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "executeMessageWithTransferFallback", _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.ExecuteMessageWithTransferFallback(&_MessageReceiverAdapter.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferFallback is a paid mutator transaction binding the contract method 0x5ab7afc6.
//
// Solidity: function executeMessageWithTransferFallback(address _sender, address _token, uint256 _amount, uint64 _srcChainId, bytes _message, address _executor) payable returns(uint8)
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) ExecuteMessageWithTransferFallback(_sender common.Address, _token common.Address, _amount *big.Int, _srcChainId uint64, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.ExecuteMessageWithTransferFallback(&_MessageReceiverAdapter.TransactOpts, _sender, _token, _amount, _srcChainId, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) ExecuteMessageWithTransferRefund(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "executeMessageWithTransferRefund", _token, _amount, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.ExecuteMessageWithTransferRefund(&_MessageReceiverAdapter.TransactOpts, _token, _amount, _message, _executor)
}

// ExecuteMessageWithTransferRefund is a paid mutator transaction binding the contract method 0x0bcb4982.
//
// Solidity: function executeMessageWithTransferRefund(address _token, uint256 _amount, bytes _message, address _executor) payable returns(uint8)
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) ExecuteMessageWithTransferRefund(_token common.Address, _amount *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.ExecuteMessageWithTransferRefund(&_MessageReceiverAdapter.TransactOpts, _token, _amount, _message, _executor)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) Pause() (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.Pause(&_MessageReceiverAdapter.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) Pause() (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.Pause(&_MessageReceiverAdapter.TransactOpts)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) RemoveGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "removeGovernor", _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.RemoveGovernor(&_MessageReceiverAdapter.TransactOpts, _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.RemoveGovernor(&_MessageReceiverAdapter.TransactOpts, _account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.RemovePauser(&_MessageReceiverAdapter.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.RemovePauser(&_MessageReceiverAdapter.TransactOpts, account)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) RenounceGovernor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "renounceGovernor")
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) RenounceGovernor() (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.RenounceGovernor(&_MessageReceiverAdapter.TransactOpts)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) RenounceGovernor() (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.RenounceGovernor(&_MessageReceiverAdapter.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) RenouncePauser() (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.RenouncePauser(&_MessageReceiverAdapter.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.RenouncePauser(&_MessageReceiverAdapter.TransactOpts)
}

// SetAllowedSender is a paid mutator transaction binding the contract method 0x83960ac7.
//
// Solidity: function setAllowedSender(address _dstContract, uint64 _srcChainId, address[] _srcContracts, bool[] _alloweds) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) SetAllowedSender(opts *bind.TransactOpts, _dstContract common.Address, _srcChainId uint64, _srcContracts []common.Address, _alloweds []bool) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "setAllowedSender", _dstContract, _srcChainId, _srcContracts, _alloweds)
}

// SetAllowedSender is a paid mutator transaction binding the contract method 0x83960ac7.
//
// Solidity: function setAllowedSender(address _dstContract, uint64 _srcChainId, address[] _srcContracts, bool[] _alloweds) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) SetAllowedSender(_dstContract common.Address, _srcChainId uint64, _srcContracts []common.Address, _alloweds []bool) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.SetAllowedSender(&_MessageReceiverAdapter.TransactOpts, _dstContract, _srcChainId, _srcContracts, _alloweds)
}

// SetAllowedSender is a paid mutator transaction binding the contract method 0x83960ac7.
//
// Solidity: function setAllowedSender(address _dstContract, uint64 _srcChainId, address[] _srcContracts, bool[] _alloweds) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) SetAllowedSender(_dstContract common.Address, _srcChainId uint64, _srcContracts []common.Address, _alloweds []bool) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.SetAllowedSender(&_MessageReceiverAdapter.TransactOpts, _dstContract, _srcChainId, _srcContracts, _alloweds)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) SetDelayPeriod(opts *bind.TransactOpts, _period *big.Int) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "setDelayPeriod", _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.SetDelayPeriod(&_MessageReceiverAdapter.TransactOpts, _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.SetDelayPeriod(&_MessageReceiverAdapter.TransactOpts, _period)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.SetMessageBus(&_MessageReceiverAdapter.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.SetMessageBus(&_MessageReceiverAdapter.TransactOpts, _messageBus)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.TransferOwnership(&_MessageReceiverAdapter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.TransferOwnership(&_MessageReceiverAdapter.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageReceiverAdapter.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterSession) Unpause() (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.Unpause(&_MessageReceiverAdapter.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MessageReceiverAdapter *MessageReceiverAdapterTransactorSession) Unpause() (*types.Transaction, error) {
	return _MessageReceiverAdapter.Contract.Unpause(&_MessageReceiverAdapter.TransactOpts)
}

// MessageReceiverAdapterAllowedSenderUpdatedIterator is returned from FilterAllowedSenderUpdated and is used to iterate over the raw logs and unpacked data for AllowedSenderUpdated events raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterAllowedSenderUpdatedIterator struct {
	Event *MessageReceiverAdapterAllowedSenderUpdated // Event containing the contract specifics and raw log

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
func (it *MessageReceiverAdapterAllowedSenderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageReceiverAdapterAllowedSenderUpdated)
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
		it.Event = new(MessageReceiverAdapterAllowedSenderUpdated)
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
func (it *MessageReceiverAdapterAllowedSenderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageReceiverAdapterAllowedSenderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageReceiverAdapterAllowedSenderUpdated represents a AllowedSenderUpdated event raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterAllowedSenderUpdated struct {
	DstContract common.Address
	SrcChainId  uint64
	SrcContract common.Address
	Allowed     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAllowedSenderUpdated is a free log retrieval operation binding the contract event 0xa88a752d56e7264efad752e7313e00c0e04a332b664a44e2eceaf477674ab7d7.
//
// Solidity: event AllowedSenderUpdated(address dstContract, uint64 srcChainId, address srcContract, bool allowed)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) FilterAllowedSenderUpdated(opts *bind.FilterOpts) (*MessageReceiverAdapterAllowedSenderUpdatedIterator, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.FilterLogs(opts, "AllowedSenderUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterAllowedSenderUpdatedIterator{contract: _MessageReceiverAdapter.contract, event: "AllowedSenderUpdated", logs: logs, sub: sub}, nil
}

// WatchAllowedSenderUpdated is a free log subscription operation binding the contract event 0xa88a752d56e7264efad752e7313e00c0e04a332b664a44e2eceaf477674ab7d7.
//
// Solidity: event AllowedSenderUpdated(address dstContract, uint64 srcChainId, address srcContract, bool allowed)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) WatchAllowedSenderUpdated(opts *bind.WatchOpts, sink chan<- *MessageReceiverAdapterAllowedSenderUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.WatchLogs(opts, "AllowedSenderUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageReceiverAdapterAllowedSenderUpdated)
				if err := _MessageReceiverAdapter.contract.UnpackLog(event, "AllowedSenderUpdated", log); err != nil {
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

// ParseAllowedSenderUpdated is a log parse operation binding the contract event 0xa88a752d56e7264efad752e7313e00c0e04a332b664a44e2eceaf477674ab7d7.
//
// Solidity: event AllowedSenderUpdated(address dstContract, uint64 srcChainId, address srcContract, bool allowed)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) ParseAllowedSenderUpdated(log types.Log) (*MessageReceiverAdapterAllowedSenderUpdated, error) {
	event := new(MessageReceiverAdapterAllowedSenderUpdated)
	if err := _MessageReceiverAdapter.contract.UnpackLog(event, "AllowedSenderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageReceiverAdapterDelayPeriodUpdatedIterator is returned from FilterDelayPeriodUpdated and is used to iterate over the raw logs and unpacked data for DelayPeriodUpdated events raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterDelayPeriodUpdatedIterator struct {
	Event *MessageReceiverAdapterDelayPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *MessageReceiverAdapterDelayPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageReceiverAdapterDelayPeriodUpdated)
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
		it.Event = new(MessageReceiverAdapterDelayPeriodUpdated)
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
func (it *MessageReceiverAdapterDelayPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageReceiverAdapterDelayPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageReceiverAdapterDelayPeriodUpdated represents a DelayPeriodUpdated event raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterDelayPeriodUpdated struct {
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDelayPeriodUpdated is a free log retrieval operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) FilterDelayPeriodUpdated(opts *bind.FilterOpts) (*MessageReceiverAdapterDelayPeriodUpdatedIterator, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.FilterLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterDelayPeriodUpdatedIterator{contract: _MessageReceiverAdapter.contract, event: "DelayPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayPeriodUpdated is a free log subscription operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) WatchDelayPeriodUpdated(opts *bind.WatchOpts, sink chan<- *MessageReceiverAdapterDelayPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.WatchLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageReceiverAdapterDelayPeriodUpdated)
				if err := _MessageReceiverAdapter.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
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

// ParseDelayPeriodUpdated is a log parse operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) ParseDelayPeriodUpdated(log types.Log) (*MessageReceiverAdapterDelayPeriodUpdated, error) {
	event := new(MessageReceiverAdapterDelayPeriodUpdated)
	if err := _MessageReceiverAdapter.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageReceiverAdapterDelayedMessageAddedIterator is returned from FilterDelayedMessageAdded and is used to iterate over the raw logs and unpacked data for DelayedMessageAdded events raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterDelayedMessageAddedIterator struct {
	Event *MessageReceiverAdapterDelayedMessageAdded // Event containing the contract specifics and raw log

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
func (it *MessageReceiverAdapterDelayedMessageAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageReceiverAdapterDelayedMessageAdded)
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
		it.Event = new(MessageReceiverAdapterDelayedMessageAdded)
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
func (it *MessageReceiverAdapterDelayedMessageAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageReceiverAdapterDelayedMessageAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageReceiverAdapterDelayedMessageAdded represents a DelayedMessageAdded event raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterDelayedMessageAdded struct {
	Id          [32]byte
	SrcContract common.Address
	SrcChainId  uint64
	DstContract common.Address
	CallData    []byte
	Nonce       uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDelayedMessageAdded is a free log retrieval operation binding the contract event 0xb86a62c31cff12d722cd0a54f5964552dfe3fc694470b892aebf3511ea60d1be.
//
// Solidity: event DelayedMessageAdded(bytes32 id, address srcContract, uint64 srcChainId, address dstContract, bytes callData, uint32 nonce)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) FilterDelayedMessageAdded(opts *bind.FilterOpts) (*MessageReceiverAdapterDelayedMessageAddedIterator, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.FilterLogs(opts, "DelayedMessageAdded")
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterDelayedMessageAddedIterator{contract: _MessageReceiverAdapter.contract, event: "DelayedMessageAdded", logs: logs, sub: sub}, nil
}

// WatchDelayedMessageAdded is a free log subscription operation binding the contract event 0xb86a62c31cff12d722cd0a54f5964552dfe3fc694470b892aebf3511ea60d1be.
//
// Solidity: event DelayedMessageAdded(bytes32 id, address srcContract, uint64 srcChainId, address dstContract, bytes callData, uint32 nonce)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) WatchDelayedMessageAdded(opts *bind.WatchOpts, sink chan<- *MessageReceiverAdapterDelayedMessageAdded) (event.Subscription, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.WatchLogs(opts, "DelayedMessageAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageReceiverAdapterDelayedMessageAdded)
				if err := _MessageReceiverAdapter.contract.UnpackLog(event, "DelayedMessageAdded", log); err != nil {
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

// ParseDelayedMessageAdded is a log parse operation binding the contract event 0xb86a62c31cff12d722cd0a54f5964552dfe3fc694470b892aebf3511ea60d1be.
//
// Solidity: event DelayedMessageAdded(bytes32 id, address srcContract, uint64 srcChainId, address dstContract, bytes callData, uint32 nonce)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) ParseDelayedMessageAdded(log types.Log) (*MessageReceiverAdapterDelayedMessageAdded, error) {
	event := new(MessageReceiverAdapterDelayedMessageAdded)
	if err := _MessageReceiverAdapter.contract.UnpackLog(event, "DelayedMessageAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageReceiverAdapterDelayedMessageExecutedIterator is returned from FilterDelayedMessageExecuted and is used to iterate over the raw logs and unpacked data for DelayedMessageExecuted events raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterDelayedMessageExecutedIterator struct {
	Event *MessageReceiverAdapterDelayedMessageExecuted // Event containing the contract specifics and raw log

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
func (it *MessageReceiverAdapterDelayedMessageExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageReceiverAdapterDelayedMessageExecuted)
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
		it.Event = new(MessageReceiverAdapterDelayedMessageExecuted)
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
func (it *MessageReceiverAdapterDelayedMessageExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageReceiverAdapterDelayedMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageReceiverAdapterDelayedMessageExecuted represents a DelayedMessageExecuted event raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterDelayedMessageExecuted struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDelayedMessageExecuted is a free log retrieval operation binding the contract event 0x3dd29d696ff6a4e0e53f61fecdcfe71ab28e10acd8ff8c59d55744d92b3fa245.
//
// Solidity: event DelayedMessageExecuted(bytes32 id)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) FilterDelayedMessageExecuted(opts *bind.FilterOpts) (*MessageReceiverAdapterDelayedMessageExecutedIterator, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.FilterLogs(opts, "DelayedMessageExecuted")
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterDelayedMessageExecutedIterator{contract: _MessageReceiverAdapter.contract, event: "DelayedMessageExecuted", logs: logs, sub: sub}, nil
}

// WatchDelayedMessageExecuted is a free log subscription operation binding the contract event 0x3dd29d696ff6a4e0e53f61fecdcfe71ab28e10acd8ff8c59d55744d92b3fa245.
//
// Solidity: event DelayedMessageExecuted(bytes32 id)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) WatchDelayedMessageExecuted(opts *bind.WatchOpts, sink chan<- *MessageReceiverAdapterDelayedMessageExecuted) (event.Subscription, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.WatchLogs(opts, "DelayedMessageExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageReceiverAdapterDelayedMessageExecuted)
				if err := _MessageReceiverAdapter.contract.UnpackLog(event, "DelayedMessageExecuted", log); err != nil {
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

// ParseDelayedMessageExecuted is a log parse operation binding the contract event 0x3dd29d696ff6a4e0e53f61fecdcfe71ab28e10acd8ff8c59d55744d92b3fa245.
//
// Solidity: event DelayedMessageExecuted(bytes32 id)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) ParseDelayedMessageExecuted(log types.Log) (*MessageReceiverAdapterDelayedMessageExecuted, error) {
	event := new(MessageReceiverAdapterDelayedMessageExecuted)
	if err := _MessageReceiverAdapter.contract.UnpackLog(event, "DelayedMessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageReceiverAdapterGovernorAddedIterator is returned from FilterGovernorAdded and is used to iterate over the raw logs and unpacked data for GovernorAdded events raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterGovernorAddedIterator struct {
	Event *MessageReceiverAdapterGovernorAdded // Event containing the contract specifics and raw log

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
func (it *MessageReceiverAdapterGovernorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageReceiverAdapterGovernorAdded)
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
		it.Event = new(MessageReceiverAdapterGovernorAdded)
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
func (it *MessageReceiverAdapterGovernorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageReceiverAdapterGovernorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageReceiverAdapterGovernorAdded represents a GovernorAdded event raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterGovernorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorAdded is a free log retrieval operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) FilterGovernorAdded(opts *bind.FilterOpts) (*MessageReceiverAdapterGovernorAddedIterator, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.FilterLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterGovernorAddedIterator{contract: _MessageReceiverAdapter.contract, event: "GovernorAdded", logs: logs, sub: sub}, nil
}

// WatchGovernorAdded is a free log subscription operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) WatchGovernorAdded(opts *bind.WatchOpts, sink chan<- *MessageReceiverAdapterGovernorAdded) (event.Subscription, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.WatchLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageReceiverAdapterGovernorAdded)
				if err := _MessageReceiverAdapter.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
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

// ParseGovernorAdded is a log parse operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) ParseGovernorAdded(log types.Log) (*MessageReceiverAdapterGovernorAdded, error) {
	event := new(MessageReceiverAdapterGovernorAdded)
	if err := _MessageReceiverAdapter.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageReceiverAdapterGovernorRemovedIterator is returned from FilterGovernorRemoved and is used to iterate over the raw logs and unpacked data for GovernorRemoved events raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterGovernorRemovedIterator struct {
	Event *MessageReceiverAdapterGovernorRemoved // Event containing the contract specifics and raw log

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
func (it *MessageReceiverAdapterGovernorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageReceiverAdapterGovernorRemoved)
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
		it.Event = new(MessageReceiverAdapterGovernorRemoved)
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
func (it *MessageReceiverAdapterGovernorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageReceiverAdapterGovernorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageReceiverAdapterGovernorRemoved represents a GovernorRemoved event raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterGovernorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorRemoved is a free log retrieval operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) FilterGovernorRemoved(opts *bind.FilterOpts) (*MessageReceiverAdapterGovernorRemovedIterator, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.FilterLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterGovernorRemovedIterator{contract: _MessageReceiverAdapter.contract, event: "GovernorRemoved", logs: logs, sub: sub}, nil
}

// WatchGovernorRemoved is a free log subscription operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) WatchGovernorRemoved(opts *bind.WatchOpts, sink chan<- *MessageReceiverAdapterGovernorRemoved) (event.Subscription, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.WatchLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageReceiverAdapterGovernorRemoved)
				if err := _MessageReceiverAdapter.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
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

// ParseGovernorRemoved is a log parse operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) ParseGovernorRemoved(log types.Log) (*MessageReceiverAdapterGovernorRemoved, error) {
	event := new(MessageReceiverAdapterGovernorRemoved)
	if err := _MessageReceiverAdapter.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageReceiverAdapterMessageBusUpdatedIterator is returned from FilterMessageBusUpdated and is used to iterate over the raw logs and unpacked data for MessageBusUpdated events raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterMessageBusUpdatedIterator struct {
	Event *MessageReceiverAdapterMessageBusUpdated // Event containing the contract specifics and raw log

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
func (it *MessageReceiverAdapterMessageBusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageReceiverAdapterMessageBusUpdated)
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
		it.Event = new(MessageReceiverAdapterMessageBusUpdated)
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
func (it *MessageReceiverAdapterMessageBusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageReceiverAdapterMessageBusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageReceiverAdapterMessageBusUpdated represents a MessageBusUpdated event raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterMessageBusUpdated struct {
	MessageBus common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMessageBusUpdated is a free log retrieval operation binding the contract event 0x3f8223bcd8b3b875473e9f9e14e1ad075451a2b5ffd31591655da9a01516bf5e.
//
// Solidity: event MessageBusUpdated(address messageBus)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) FilterMessageBusUpdated(opts *bind.FilterOpts) (*MessageReceiverAdapterMessageBusUpdatedIterator, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.FilterLogs(opts, "MessageBusUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterMessageBusUpdatedIterator{contract: _MessageReceiverAdapter.contract, event: "MessageBusUpdated", logs: logs, sub: sub}, nil
}

// WatchMessageBusUpdated is a free log subscription operation binding the contract event 0x3f8223bcd8b3b875473e9f9e14e1ad075451a2b5ffd31591655da9a01516bf5e.
//
// Solidity: event MessageBusUpdated(address messageBus)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) WatchMessageBusUpdated(opts *bind.WatchOpts, sink chan<- *MessageReceiverAdapterMessageBusUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.WatchLogs(opts, "MessageBusUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageReceiverAdapterMessageBusUpdated)
				if err := _MessageReceiverAdapter.contract.UnpackLog(event, "MessageBusUpdated", log); err != nil {
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

// ParseMessageBusUpdated is a log parse operation binding the contract event 0x3f8223bcd8b3b875473e9f9e14e1ad075451a2b5ffd31591655da9a01516bf5e.
//
// Solidity: event MessageBusUpdated(address messageBus)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) ParseMessageBusUpdated(log types.Log) (*MessageReceiverAdapterMessageBusUpdated, error) {
	event := new(MessageReceiverAdapterMessageBusUpdated)
	if err := _MessageReceiverAdapter.contract.UnpackLog(event, "MessageBusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageReceiverAdapterMessageReceivedIterator is returned from FilterMessageReceived and is used to iterate over the raw logs and unpacked data for MessageReceived events raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterMessageReceivedIterator struct {
	Event *MessageReceiverAdapterMessageReceived // Event containing the contract specifics and raw log

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
func (it *MessageReceiverAdapterMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageReceiverAdapterMessageReceived)
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
		it.Event = new(MessageReceiverAdapterMessageReceived)
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
func (it *MessageReceiverAdapterMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageReceiverAdapterMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageReceiverAdapterMessageReceived represents a MessageReceived event raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterMessageReceived struct {
	SrcContract common.Address
	SrcChainId  uint64
	DstContract common.Address
	CallData    []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMessageReceived is a free log retrieval operation binding the contract event 0x89e5332ec6ff132178f395753d055e19755ab6620d41de1c27e0a135a54f6d67.
//
// Solidity: event MessageReceived(address srcContract, uint64 srcChainId, address dstContract, bytes callData)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) FilterMessageReceived(opts *bind.FilterOpts) (*MessageReceiverAdapterMessageReceivedIterator, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterMessageReceivedIterator{contract: _MessageReceiverAdapter.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

// WatchMessageReceived is a free log subscription operation binding the contract event 0x89e5332ec6ff132178f395753d055e19755ab6620d41de1c27e0a135a54f6d67.
//
// Solidity: event MessageReceived(address srcContract, uint64 srcChainId, address dstContract, bytes callData)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *MessageReceiverAdapterMessageReceived) (event.Subscription, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageReceiverAdapterMessageReceived)
				if err := _MessageReceiverAdapter.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

// ParseMessageReceived is a log parse operation binding the contract event 0x89e5332ec6ff132178f395753d055e19755ab6620d41de1c27e0a135a54f6d67.
//
// Solidity: event MessageReceived(address srcContract, uint64 srcChainId, address dstContract, bytes callData)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) ParseMessageReceived(log types.Log) (*MessageReceiverAdapterMessageReceived, error) {
	event := new(MessageReceiverAdapterMessageReceived)
	if err := _MessageReceiverAdapter.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageReceiverAdapterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterOwnershipTransferredIterator struct {
	Event *MessageReceiverAdapterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageReceiverAdapterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageReceiverAdapterOwnershipTransferred)
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
		it.Event = new(MessageReceiverAdapterOwnershipTransferred)
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
func (it *MessageReceiverAdapterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageReceiverAdapterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageReceiverAdapterOwnershipTransferred represents a OwnershipTransferred event raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageReceiverAdapterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageReceiverAdapter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterOwnershipTransferredIterator{contract: _MessageReceiverAdapter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageReceiverAdapterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageReceiverAdapter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageReceiverAdapterOwnershipTransferred)
				if err := _MessageReceiverAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) ParseOwnershipTransferred(log types.Log) (*MessageReceiverAdapterOwnershipTransferred, error) {
	event := new(MessageReceiverAdapterOwnershipTransferred)
	if err := _MessageReceiverAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageReceiverAdapterPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterPausedIterator struct {
	Event *MessageReceiverAdapterPaused // Event containing the contract specifics and raw log

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
func (it *MessageReceiverAdapterPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageReceiverAdapterPaused)
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
		it.Event = new(MessageReceiverAdapterPaused)
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
func (it *MessageReceiverAdapterPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageReceiverAdapterPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageReceiverAdapterPaused represents a Paused event raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) FilterPaused(opts *bind.FilterOpts) (*MessageReceiverAdapterPausedIterator, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterPausedIterator{contract: _MessageReceiverAdapter.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MessageReceiverAdapterPaused) (event.Subscription, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageReceiverAdapterPaused)
				if err := _MessageReceiverAdapter.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) ParsePaused(log types.Log) (*MessageReceiverAdapterPaused, error) {
	event := new(MessageReceiverAdapterPaused)
	if err := _MessageReceiverAdapter.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageReceiverAdapterPauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterPauserAddedIterator struct {
	Event *MessageReceiverAdapterPauserAdded // Event containing the contract specifics and raw log

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
func (it *MessageReceiverAdapterPauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageReceiverAdapterPauserAdded)
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
		it.Event = new(MessageReceiverAdapterPauserAdded)
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
func (it *MessageReceiverAdapterPauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageReceiverAdapterPauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageReceiverAdapterPauserAdded represents a PauserAdded event raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterPauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*MessageReceiverAdapterPauserAddedIterator, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterPauserAddedIterator{contract: _MessageReceiverAdapter.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *MessageReceiverAdapterPauserAdded) (event.Subscription, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageReceiverAdapterPauserAdded)
				if err := _MessageReceiverAdapter.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) ParsePauserAdded(log types.Log) (*MessageReceiverAdapterPauserAdded, error) {
	event := new(MessageReceiverAdapterPauserAdded)
	if err := _MessageReceiverAdapter.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageReceiverAdapterPauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterPauserRemovedIterator struct {
	Event *MessageReceiverAdapterPauserRemoved // Event containing the contract specifics and raw log

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
func (it *MessageReceiverAdapterPauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageReceiverAdapterPauserRemoved)
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
		it.Event = new(MessageReceiverAdapterPauserRemoved)
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
func (it *MessageReceiverAdapterPauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageReceiverAdapterPauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageReceiverAdapterPauserRemoved represents a PauserRemoved event raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterPauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*MessageReceiverAdapterPauserRemovedIterator, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterPauserRemovedIterator{contract: _MessageReceiverAdapter.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *MessageReceiverAdapterPauserRemoved) (event.Subscription, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageReceiverAdapterPauserRemoved)
				if err := _MessageReceiverAdapter.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) ParsePauserRemoved(log types.Log) (*MessageReceiverAdapterPauserRemoved, error) {
	event := new(MessageReceiverAdapterPauserRemoved)
	if err := _MessageReceiverAdapter.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageReceiverAdapterUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterUnpausedIterator struct {
	Event *MessageReceiverAdapterUnpaused // Event containing the contract specifics and raw log

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
func (it *MessageReceiverAdapterUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageReceiverAdapterUnpaused)
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
		it.Event = new(MessageReceiverAdapterUnpaused)
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
func (it *MessageReceiverAdapterUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageReceiverAdapterUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageReceiverAdapterUnpaused represents a Unpaused event raised by the MessageReceiverAdapter contract.
type MessageReceiverAdapterUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MessageReceiverAdapterUnpausedIterator, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MessageReceiverAdapterUnpausedIterator{contract: _MessageReceiverAdapter.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MessageReceiverAdapterUnpaused) (event.Subscription, error) {

	logs, sub, err := _MessageReceiverAdapter.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageReceiverAdapterUnpaused)
				if err := _MessageReceiverAdapter.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MessageReceiverAdapter *MessageReceiverAdapterFilterer) ParseUnpaused(log types.Log) (*MessageReceiverAdapterUnpaused, error) {
	event := new(MessageReceiverAdapterUnpaused)
	if err := _MessageReceiverAdapter.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

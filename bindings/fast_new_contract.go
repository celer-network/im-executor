package bindings

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var AdapterABI, _ = abi.JSON(strings.NewReader(MessageReceiverAdapterABI))

func FastNewMessageReceiverAdapter(address common.Address, backend bind.ContractBackend) *MessageReceiverAdapter {
	contract := bind.NewBoundContract(address, AdapterABI, backend, backend, backend)
	return &MessageReceiverAdapter{MessageReceiverAdapterCaller: MessageReceiverAdapterCaller{contract: contract}, MessageReceiverAdapterTransactor: MessageReceiverAdapterTransactor{contract: contract}, MessageReceiverAdapterFilterer: MessageReceiverAdapterFilterer{contract: contract}}
}

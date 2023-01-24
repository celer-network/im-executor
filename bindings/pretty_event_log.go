package bindings

import "fmt"

func (ev *MessageReceiverAdapterDelayedMessageAdded) PrettyLog(chainId uint64) string {
	return fmt.Sprintf("DelayedMessageAdded, delayId: %x, adapter: %x, delayTx: %x, srcContract: %x, srcChainId: %d, dstContract: %x, dstChainId:%d, nonce: %d",
		ev.Id, ev.Raw.Address, ev.Raw.TxHash, ev.SrcContract, ev.SrcChainId, ev.DstContract, chainId, ev.Nonce)
}

func (ev *MessageReceiverAdapterDelayedMessageExecuted) PrettyLog(chainId uint64) string {
	return fmt.Sprintf("DelayedMessageExecuted, delayId: %x, adapter: %x, executeTx: %x, dstChainId: %d", ev.Id, ev.Raw.Address, ev.Raw.TxHash, chainId)
}

func (ev *MessageReceiverAdapterDelayPeriodUpdated) PrettyLog(chainId uint64) string {
	return fmt.Sprintf("DelayPeriodUpdated, adapter: %x, chainId: %d, delayPeriod: %s", ev.Raw.Address, chainId, ev.Period)
}

func (ev *MessageReceiverAdapterDelayedMessageAdded) String() string {
	return fmt.Sprintf("delayId: %x, adapter: %x, delayTx: %x, srcContract: %x, srcChainId: %d, dstContract: %x, nonce: %d",
		ev.Id, ev.Raw.Address, ev.Raw.TxHash, ev.SrcContract, ev.SrcChainId, ev.DstContract, ev.Nonce)
}

func (ev *MessageReceiverAdapterDelayedMessageExecuted) String() string {
	return fmt.Sprintf("delayId: %x, adapter: %x, executeTx: %x", ev.Id, ev.Raw.Address, ev.Raw.TxHash)
}

func (ev *MessageReceiverAdapterDelayPeriodUpdated) String() string {
	return fmt.Sprintf("adapter: %x, delayPeriod: %s", ev.Raw.Address, ev.Period)
}

package executor

import (
	"math/big"

	ethtypes "github.com/celer-network/im-executor/sgn-v2/eth"
	msgtypes "github.com/celer-network/im-executor/sgn-v2/x/message/types"
	"github.com/celer-network/im-executor/types"
)

func newTransferInfo(execCtx *msgtypes.ExecutionContext) ethtypes.MsgDataTypesTransferInfo {
	message := execCtx.Message
	transfer := execCtx.GetTransfer()
	amount, _ := new(big.Int).SetString(transfer.Amount, 10)

	return ethtypes.MsgDataTypesTransferInfo{
		T:          uint8(message.GetTransferType()),
		Sender:     ethtypes.Hex2Addr(message.Sender),
		Receiver:   ethtypes.Hex2Addr(message.Receiver),
		Token:      ethtypes.Bytes2Addr(transfer.Token),
		Amount:     amount,
		Wdseq:      transfer.WdSeqNum,
		SrcChainId: message.SrcChainId,
		RefId:      ethtypes.Bytes2Hash(execCtx.Message.GetTransferRefId()),
		SrcTxHash:  ethtypes.Hex2Hash(message.SrcTxHash),
	}
}

func newRouteInfo(execCtx *msgtypes.ExecutionContext) types.MessageRoute {
	message := execCtx.Message

	return types.MessageRoute{
		Sender:      message.Sender,
		Receiver:    message.Receiver,
		SrcChainId:  message.SrcChainId,
		SrcTxHash:   message.SrcTxHash,
		IsBytesAddr: message.IsBytesAddr,
	}
}

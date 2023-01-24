package types

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	cbrtypes "github.com/celer-network/im-executor/sgn-v2/x/cbridge/types"
	pegbrtypes "github.com/celer-network/im-executor/sgn-v2/x/pegbridge/types"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func NewMsgXferExecutionContext(
	ev *eth.MessageBusMessageWithTransfer,
	chainId uint64,
	dstToken eth.Addr,
	dstAmt string,
	dstBridge eth.Addr,
	transferType TransferType) *ExecutionContext {

	message := Message{
		SrcChainId:    chainId,
		DstChainId:    ev.DstChainId.Uint64(),
		Sender:        eth.Addr2Hex(ev.Sender),
		Receiver:      eth.Addr2Hex(ev.Receiver),
		Data:          ev.Message,
		Fee:           ev.Fee.String(),
		TransferType:  transferType,
		TransferRefId: ev.SrcTransferId[:],
		SrcTxHash:     ev.Raw.TxHash.Hex(),
	}
	transfer := &Transfer{
		Token:  dstToken.Bytes(),
		Amount: dstAmt,
	}
	execCtx := &ExecutionContext{
		Message:  message,
		Transfer: transfer,
	}
	execCtx.MessageId = execCtx.ComputeMessageId(dstBridge)
	return execCtx
}

func NewMsgXferRefundExecutionContext(
	ev *eth.MessageBusMessageWithTransfer,
	wdOnchain *cbrtypes.WithdrawOnchain,
	nonce uint64,
	bridge eth.Addr) *ExecutionContext {

	message := Message{
		SrcChainId:    wdOnchain.Chainid,
		DstChainId:    wdOnchain.Chainid,
		Sender:        eth.ZeroAddrHex,
		Receiver:      eth.Addr2Hex(ev.Sender),
		Data:          ev.Message,
		TransferType:  TRANSFER_TYPE_LIQUIDITY_WITHDRAW,
		TransferRefId: ev.SrcTransferId[:],
		Fee:           ev.Fee.String(),
		SrcTxHash:     ev.Raw.TxHash.Hex(),
	}
	transfer := &Transfer{
		Amount:   new(big.Int).SetBytes(wdOnchain.Amount).String(),
		Token:    wdOnchain.Token,
		WdSeqNum: nonce,
	}
	execCtx := &ExecutionContext{
		Message:  message,
		Transfer: transfer,
	}
	execCtx.MessageId = execCtx.ComputeMessageId(bridge)
	return execCtx
}

func NewMsgPegDepositRefundExecutionContext(
	ev *eth.MessageBusMessageWithTransfer,
	wdOnChain pegbrtypes.WithdrawOnChain,
	pegVault eth.Addr, vaultVersion uint32) *ExecutionContext {

	transferType := TRANSFER_TYPE_PEG_WITHDRAW
	if vaultVersion == 2 {
		transferType = TRANSFER_TYPE_PEG_V2_WITHDRAW
	}

	message := Message{
		SrcChainId:    wdOnChain.RefChainId,
		DstChainId:    wdOnChain.RefChainId,
		Sender:        eth.ZeroAddrHex,
		Receiver:      eth.Bytes2AddrHex(wdOnChain.Receiver),
		Data:          ev.Message,
		TransferType:  transferType,
		TransferRefId: wdOnChain.RefId,
		Fee:           ev.Fee.String(),
		SrcTxHash:     ev.Raw.TxHash.Hex(),
	}
	transfer := &Transfer{
		Amount: new(big.Int).SetBytes(wdOnChain.Amount).String(),
		Token:  wdOnChain.Token,
	}
	execCtx := &ExecutionContext{
		Message:  message,
		Transfer: transfer,
	}
	execCtx.MessageId = execCtx.ComputeMessageId(pegVault)
	return execCtx
}

func NewMsgPegBurnRefundExecutionContext(
	ev *eth.MessageBusMessageWithTransfer,
	mintOnChain pegbrtypes.MintOnChain,
	pegBridge eth.Addr, bridgeVersion uint32) *ExecutionContext {

	transferType := TRANSFER_TYPE_PEG_MINT
	if bridgeVersion == 2 {
		transferType = TRANSFER_TYPE_PEG_V2_MINT
	}

	message := Message{
		SrcChainId:    mintOnChain.RefChainId,
		DstChainId:    mintOnChain.RefChainId,
		Sender:        eth.ZeroAddrHex,
		Receiver:      eth.Bytes2AddrHex(mintOnChain.Account),
		Data:          ev.Message,
		TransferType:  transferType,
		TransferRefId: mintOnChain.RefId,
		Fee:           ev.Fee.String(),
		SrcTxHash:     ev.Raw.TxHash.Hex(),
	}
	transfer := &Transfer{
		Amount: new(big.Int).SetBytes(mintOnChain.Amount).String(),
		Token:  mintOnChain.Token,
	}
	execCtx := &ExecutionContext{
		Message:  message,
		Transfer: transfer,
	}
	execCtx.MessageId = execCtx.ComputeMessageId(pegBridge)
	return execCtx
}

func (c *ExecutionContext) MustMarshal() []byte {
	data, err := c.Marshal()
	if err != nil {
		log.Panicf("failed to marshal execCtx %+v", c)
	}
	return data
}

func (c *ExecutionContext) ComputeMessageId(bridgeAddr eth.Addr) []byte {
	msg := c.Message
	if msg.TransferType == TRANSFER_TYPE_NULL {
		return msg.ComputeMessageIdNoTransfer()
	}
	return c.ComputeMessageIdWithTransfer(bridgeAddr)
}

func (c *ExecutionContext) ComputeMessageIdWithTransfer(dstBridgeAddr eth.Addr) []byte {
	dstTransferId := c.ComputeDstTransferId(dstBridgeAddr)
	return ComputeMessageIdFromDstTransfer(dstTransferId, dstBridgeAddr)
}

func (c *ExecutionContext) ComputeDstTransferId(dstBridgeAddr eth.Addr) []byte {
	var dstTransferId []byte
	m := c.Message
	t := c.Transfer
	switch m.TransferType {
	case TRANSFER_TYPE_NULL:
		return nil
	case TRANSFER_TYPE_LIQUIDITY_RELAY:
		log.Debugf("TransferType:%s, %s, %s, %x, %s, %d, %d, %x", m.TransferType, m.Sender, m.Receiver, t.Token, t.Amount, m.SrcChainId, m.DstChainId, m.TransferRefId)
		dstTransferId = solsha3.SoliditySHA3(
			[]string{"address", "address", "address", "uint256", "uint64", "uint64", "bytes32"},
			m.Sender, m.Receiver, t.Token, t.Amount, m.SrcChainId, m.DstChainId, m.TransferRefId,
		)
	case TRANSFER_TYPE_LIQUIDITY_WITHDRAW:
		log.Debugf("TransferType:%s, %s, %s, %x, %s, %d, %d, %x", m.TransferType, m.Sender, m.Receiver, t.Token, t.Amount, m.SrcChainId, m.DstChainId, m.TransferRefId)
		dstTransferId = solsha3.SoliditySHA3(
			[]string{"uint64", "uint64", "address", "address", "uint256"},
			m.DstChainId, t.WdSeqNum, m.Receiver, t.Token, t.Amount,
		)
	case TRANSFER_TYPE_PEG_MINT, TRANSFER_TYPE_PEG_WITHDRAW:
		log.Debugf("TransferType:%s, %s, %x, %s, %s, %d, %x", m.TransferType, m.Receiver, t.Token, t.Amount, m.Sender, m.SrcChainId, m.TransferRefId)
		dstTransferId = solsha3.SoliditySHA3(
			[]string{"address", "address", "uint256", "address", "uint64", "bytes32"},
			m.Receiver, t.Token, t.Amount, m.Sender, m.SrcChainId, m.TransferRefId,
		)
	case TRANSFER_TYPE_PEG_V2_MINT, TRANSFER_TYPE_PEG_V2_WITHDRAW:
		log.Debugf("TransferType:%s, %s, %x, %s, %s, %d, %x", m.TransferType, m.Receiver, t.Token, t.Amount, m.Sender, m.SrcChainId, m.TransferRefId)
		dstTransferId = solsha3.SoliditySHA3(
			[]string{"address", "address", "uint256", "address", "uint64", "bytes32", "address"},
			[]interface{}{m.Receiver, t.Token, t.Amount, m.Sender, m.SrcChainId, m.TransferRefId, dstBridgeAddr},
		)
	}
	return dstTransferId
}

func (c *ExecutionContext) PrettyPrint() {
	if c == nil {
		fmt.Println("nil execution context")
		return
	}
	fmt.Printf("message_id: %x\n\n", c.MessageId)
	c.Message.PrettyPrint()
	if c.Transfer != nil {
		fmt.Println()
		fmt.Printf("token: %x\n", c.Transfer.Token)
		fmt.Printf("amount: %s\n", c.Transfer.Amount)
		if c.Transfer.WdSeqNum != 0 {
			fmt.Printf("wd_seq_num: %d\n", c.Transfer.WdSeqNum)
		}
	}
	sigs := []string{}
	for _, sig := range c.Message.Signatures {
		sigs = append(sigs, eth.Bytes2Hex(sig.SigBytes))
	}
	fmt.Printf("signatures %s", strings.Join(sigs, ","))
}

func ComputeMessageIdFromDstTransfer(dstTransferId []byte, dstBridgeAddr eth.Addr) []byte {
	// Prepend bridge address and hash again
	msgId := solsha3.SoliditySHA3(
		[]string{"uint8", "address", "bytes32"},
		uint8(MsgType_MSG_TYPE_MESSAGE_WITH_TRANSFER), dstBridgeAddr, dstTransferId,
	)
	log.Debugf("ComputeMessageIdFromDstTransferId, dstTransferId %x, dstBridgeAddr %x, messageId %x",
		dstTransferId, dstBridgeAddr, msgId)
	return msgId
}

package eth

import (
	"fmt"
	"time"
)

func (l *BridgeLiquidityAdded) String() string {
	return fmt.Sprintf("seqNum %d, provider %x, token %x, amount %s", l.Seqnum, l.Provider, l.Token, l.Amount)
}

func (s *BridgeSend) String() string {
	return fmt.Sprintf("transferId %x, sender %x, receiver %x, token %x, amount %s, dstChainId %d, nonce %d, maxSlippage %d",
		s.TransferId, s.Sender, s.Receiver, s.Token, s.Amount, s.DstChainId, s.Nonce, s.MaxSlippage)
}

func (r *BridgeRelay) String() string {
	return fmt.Sprintf("transferId %x, sender %x, receiver %x, token %x, amount %s, srcChainId %d, srcTransferId %x",
		r.TransferId, r.Sender, r.Receiver, r.Token, r.Amount, r.SrcChainId, r.SrcTransferId)
}

func (w *BridgeWithdrawDone) String() string {
	return fmt.Sprintf("withdrawId %x, seqNum %d, receiver %x, token %x, amount %s",
		w.WithdrawId, w.Seqnum, w.Receiver, w.Token, w.Amount)
}

func (s *BridgeSignersUpdated) String() string {
	var out string
	for i, addr := range s.Signers {
		out += fmt.Sprintf("<addr %x power %s> ", addr, s.Powers[i])
	}
	return fmt.Sprintf("< %s>", out)
}

func (ev *OriginalTokenVaultDeposited) String() string {
	return fmt.Sprintf("depositId %x, account %x, token %x, amount %s, mintChainId %d,  mintAccount %x",
		ev.DepositId, ev.Depositor, ev.Token, ev.Amount, ev.MintChainId, ev.MintAccount)
}

func (ev *PeggedTokenBridgeMint) String() string {
	return fmt.Sprintf("mintId %x, account %x, token %x, amount %s, depositChainId %d, depositId %x, depositor %x",
		ev.MintId, ev.Account, ev.Token, ev.Amount, ev.RefChainId, ev.RefId, ev.Depositor)
}

func (ev *PeggedTokenBridgeBurn) String() string {
	return fmt.Sprintf("burnId %x, account %x, token %x, amount %s, withdrawAccount %x",
		ev.BurnId, ev.Account, ev.Token, ev.Amount, ev.WithdrawAccount)
}

func (ev *OriginalTokenVaultWithdrawn) String() string {
	return fmt.Sprintf("withdrawId %x, receiver %x, token %x, amount %s, burnChainId %d, burnId %x, burnAccount %x",
		ev.WithdrawId, ev.Receiver, ev.Token, ev.Amount, ev.RefChainId, ev.RefId, ev.BurnAccount)
}

func (ev *OriginalTokenVaultV2Deposited) String() string {
	return fmt.Sprintf("depositId %x, account %x, token %x, amount %s, mintChainId %d, mintAccount %x, nonce %d",
		ev.DepositId, ev.Depositor, ev.Token, ev.Amount, ev.MintChainId, ev.MintAccount, ev.Nonce)
}

func (ev *PeggedTokenBridgeV2Mint) String() string {
	return fmt.Sprintf("mintId %x, account %x, token %x, amount %s, depositChainId %d, depositId %x, depositor %x",
		ev.MintId, ev.Account, ev.Token, ev.Amount, ev.RefChainId, ev.RefId, ev.Depositor)
}

func (ev *PeggedTokenBridgeV2Burn) String() string {
	return fmt.Sprintf("burnId %x, account %x, token %x, amount %s, toChainId %d, toAccount %x, nonce %d",
		ev.BurnId, ev.Account, ev.Token, ev.Amount, ev.ToChainId, ev.ToAccount, ev.Nonce)
}

func (ev *OriginalTokenVaultV2Withdrawn) String() string {
	return fmt.Sprintf("withdrawId %x, receiver %x, token %x, amount %s, burnChainId %d, burnId %x, burnAccount %x",
		ev.WithdrawId, ev.Receiver, ev.Token, ev.Amount, ev.RefChainId, ev.RefId, ev.BurnAccount)
}

func (ev *WithdrawInboxWithdrawalRequest) String() string {
	return fmt.Sprintf(
		"sender: %x receiver: %x toChain: %d fromChains: %v tokens: %v ratios: %v slippages: %v deadline: %s",
		ev.Sender, ev.Receiver, ev.ToChain, ev.FromChains, ev.Tokens, ev.Ratios, ev.Slippages, time.Unix(ev.Deadline.Int64(), 0))
}

func (ev *MessageBusMessage) String() string {
	return fmt.Sprintf("sender: %x, receiver: %x, dstChainId: %s, Message: %x, Fee: %s",
		ev.Sender, ev.Receiver, ev.DstChainId, ev.Message, ev.Fee)
}

func (ev *MessageBusMessage2) String() string {
	return fmt.Sprintf("sender: %x, receiver: %x, dstChainId: %s, Message: %x, Fee: %s",
		ev.Sender, ev.Receiver, ev.DstChainId, ev.Message, ev.Fee)
}

func (ev *MessageBusMessageWithTransfer) String() string {
	return fmt.Sprintf("sender: %x, receiver: %x, dstChainId: %s, bridgeAddr: %x, srcTransferId: %x, Message: %x, Fee: %s",
		ev.Sender, ev.Receiver, ev.DstChainId, ev.Bridge, ev.SrcTransferId, ev.Message, ev.Fee)
}

func (ev *MessageBusExecuted) String() string {
	return fmt.Sprintf("msgType: %d, msgId: %x, status: %d, receiver: %x, srcChainId: %d, srcTxhash: %x",
		ev.MsgType, ev.MsgId, ev.Status, ev.Receiver, ev.SrcChainId, ev.SrcTxHash)
}

func (ev *VotingEscrowTokenDeposited) String() string {
	return fmt.Sprintf(
		"locker: %x value: %s lockTime: %s lockType: %s timestamp: %s",
		ev.Locker, ev.Value, ev.LockTime, ev.LockType, ev.Timestamp)
}

func (ev *VotingEscrowTokenWithdrawn) String() string {
	return fmt.Sprintf(
		"locker: %x value: %s timestamp: %s",
		ev.Locker, ev.Value, ev.Timestamp)
}

// return human friendly string for logging
func (ev *BridgeSend) PrettyLog(srcChid uint64) string {
	// max slippage uint is float * 1e6 so percentage needs to divide by 1e4
	return fmt.Sprintf("send-%x src: %d-%x dstchid: %d sender: %x receiver: %x amt: %s maxslip: %f%%", ev.TransferId, srcChid, ev.Token, ev.DstChainId, ev.Sender, ev.Receiver, ev.Amount, float64(ev.MaxSlippage)/10000)
}

// onchid is the chainid this event happen
func (ev *BridgeLiquidityAdded) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("cbr-liqadd-%d-%d token: %x lp: %x amt: %s", onchid, ev.Seqnum, ev.Token, ev.Provider, ev.Amount)
}

func (ev *BridgeWithdrawDone) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("cbr-withdraw-%d chid: %d token: %x receiver: %x amt: %s", ev.Seqnum, onchid, ev.Token, ev.Receiver, ev.Amount)
}

// relay-%x is src transfer id!!! so we can easily correlate with send log
func (ev *BridgeRelay) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("cbr-relay-%x srcchid: %d dst: %d-%x sender: %x receiver: %x amt: %s thisXferId: %x", ev.SrcTransferId, ev.SrcChainId, onchid, ev.Token, ev.Sender, ev.Receiver, ev.Amount, ev.TransferId)
}

func (ev *BridgeSignersUpdated) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("cbr-signersUpdated-%d: %s", onchid, ev)
}

// onchid is the chainid this event happen
func (ev *OriginalTokenVaultDeposited) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-deposited-%d depositId: %x account: %x token: %x amount: %s mintChainId: %d mintAccount: %x",
		onchid, ev.DepositId, ev.Depositor, ev.Token, ev.Amount, ev.MintChainId, ev.MintAccount)
}

func (ev *PeggedTokenBridgeMint) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-mint-%d mintId: %x token: %x account: %x amount: %s depositChainId: %d depositId: %x depositor: %x",
		onchid, ev.MintId, ev.Token, ev.Account, ev.Amount, ev.RefChainId, ev.RefId, ev.Depositor)
}

func (ev *PeggedTokenBridgeBurn) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-burn-%d burnId: %x token: %x account: %x amount: %s withdrawAccount: %x",
		onchid, ev.BurnId, ev.Token, ev.Account, ev.Amount, ev.WithdrawAccount)
}

func (ev *OriginalTokenVaultWithdrawn) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-withdrawn-%d withdrawId: %x receiver: %x token: %x amount: %s burnChainId: %d burnId: %x burnAccount: %x",
		onchid, ev.WithdrawId, ev.Receiver, ev.Token, ev.Amount, ev.RefChainId, ev.RefId, ev.BurnAccount)
}

func (ev *OriginalTokenVaultV2Deposited) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-deposited-%d depositId: %x account: %x token: %x amount: %s mintChainId: %d mintAccount: %x",
		onchid, ev.DepositId, ev.Depositor, ev.Token, ev.Amount, ev.MintChainId, ev.MintAccount)
}

func (ev *PeggedTokenBridgeV2Mint) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-mint-%d mintId: %x token: %x account: %x amount: %s depositChainId: %d depositId: %x depositor: %x",
		onchid, ev.MintId, ev.Token, ev.Account, ev.Amount, ev.RefChainId, ev.RefId, ev.Depositor)
}

func (ev *PeggedTokenBridgeV2Burn) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-burn-%d burnId: %x token: %x account: %x amount: %s toChainId %d toAccount: %x",
		onchid, ev.BurnId, ev.Token, ev.Account, ev.Amount, ev.ToChainId, ev.ToAccount)
}

func (ev *OriginalTokenVaultV2Withdrawn) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"peg-withdrawn-%d withdrawId: %x receiver: %x token: %x amount: %s mintChainId: %d mintId: %x burnAccount: %x",
		onchid, ev.WithdrawId, ev.Receiver, ev.Token, ev.Amount, ev.RefChainId, ev.RefId, ev.BurnAccount)
}

func (ev *WithdrawInboxWithdrawalRequest) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"wdi-withdrawalrequest-%d sender: %x receiver: %x toChain: %d fromChains: %v tokens: %v ratios: %v slippages: %v deadline: %s",
		onchid, ev.Sender, ev.Receiver, ev.ToChain, ev.FromChains, ev.Tokens, ev.Ratios, ev.Slippages, time.Unix(ev.Deadline.Int64(), 0))
}

func (ev *MessageBusMessage) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("Message-%d: sender %x, receiver %x, dstChainId %s, txHash %x",
		onchid, ev.Sender, ev.Receiver, ev.DstChainId, ev.Raw.TxHash)
}

func (ev *MessageBusMessage2) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("Message-%d: sender %x, receiver %x, dstChainId %s, txHash %x",
		onchid, ev.Sender, ev.Receiver, ev.DstChainId, ev.Raw.TxHash)
}

func (ev *MessageBusMessageWithTransfer) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("MessageWithTransfer-%d: sender %x, receiver %x, dstChainId %s, bridge %x, srcTransferId %x, txHash %x",
		onchid, ev.Sender, ev.Receiver, ev.DstChainId, ev.Bridge, ev.SrcTransferId, ev.Raw.TxHash)
}

func (ev *MessageBusExecuted) PrettyLog(onchid uint64) string {
	return fmt.Sprintf("MessageExecuted-%d: msgType %d, msgId %x, status %d, receiver %x, srcChainId %d, srcTxhash %x",
		onchid, ev.MsgType, ev.MsgId, ev.Status, ev.Receiver, ev.SrcChainId, ev.SrcTxHash)
}

func (ev *VotingEscrowTokenDeposited) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"ve-deposited-%d locker: %x value: %s lockTime: %s timestamp: %s txHash: %x",
		onchid, ev.Locker, ev.Value, ev.LockTime, ev.Timestamp, ev.Raw.TxHash)
}

func (ev *VotingEscrowTokenWithdrawn) PrettyLog(onchid uint64) string {
	return fmt.Sprintf(
		"ve-withdrawn-%d locker: %x value: %s timestamp: %s txHash: %x",
		onchid, ev.Locker, ev.Value, ev.Timestamp, ev.Raw.TxHash)
}

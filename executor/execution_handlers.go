package executor

import (
	"fmt"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	cbrtypes "github.com/celer-network/im-executor/sgn-v2/x/cbridge/types"
	msgtypes "github.com/celer-network/im-executor/sgn-v2/x/message/types"
	"github.com/celer-network/im-executor/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

func (e *Executor) executeMessage(x *Execution) {
	execCtx := x.Record.ExecutionContext
	id := execCtx.MessageId

	// 2. build transaction data
	route := newRouteInfo(execCtx)
	msg, sigs, signers, powers, err := e.getMsgSignInfo(execCtx)
	if err != nil {
		log.Errorf("failed to query chain signers with chainId %d: %s", execCtx.Message.DstChainId, err.Error())
		return
	}

	// 3. update record status to executing
	db := dal.GetDB()
	err = db.UpdateStatus(id, types.ExecutionStatus_Executing)
	if err != nil {
		log.Errorln("cannot executeMessage", err)
		return
	}

	// 4. send transaction
	log.Infof("executing %s", x.Record)
	tx, err := x.Transactor.Transact(
		newTransactionCallback(id, "execute message"),
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*gethtypes.Transaction, error) {
			setValue(opts, x.Receiver.PayableValue)
			method := func() (*gethtypes.Transaction, error) {
				return x.ExecuteMessage(opts, msg, route, sigs, signers, powers)
			}
			return wrapRevertCheck(opts, method)
		}, x.GetTxOptions()...)
	if err != nil {
		handleExecuteMessageError(err, id)
		return
	}

	// 5. handle transaction result
	log.Infof("executed msg (id %x): txhash %x", id, tx.Hash())
	db.UpdateTx(id, tx.Hash())
}

func (e *Executor) executeMessageWithTransfer(x *Execution) {
	execCtx := x.Record.ExecutionContext
	id := execCtx.MessageId

	// 1. check on-chain if transfer is ready
	bridge := x.Chain.GetMsgBridgeAddr(execCtx.Message.GetTransferType())
	transferId := eth.Bytes2Hash(execCtx.ComputeDstTransferId(bridge))
	ready, err := x.Chain.IsTransferReady(transferId, execCtx.Message.TransferType)
	if err != nil {
		log.Errorf("failed to query on-chain transfer for message (id %x, transferType %s, dstTransferId %x): %s",
			execCtx.MessageId, execCtx.Message.TransferType, transferId, err.Error())
		return
	}
	if !ready {
		log.Debugf("[skip execution] message with transfer (id %x) because transfer is not seen on dst chain yet", execCtx.MessageId)
		return
	}

	// 2. build transaction data
	msg, sigs, signers, powers, err := e.getMsgSignInfo(execCtx)
	if err != nil {
		log.Errorf("failed to query chain signers with chainId %d: %s", execCtx.Message.DstChainId, err.Error())
		return
	}
	xfer := newTransferInfo(execCtx)

	// 3. update record status to executing
	db := dal.GetDB()
	err = db.UpdateStatus(id, types.ExecutionStatus_Executing)
	if err != nil {
		log.Errorf("cannot execute message with transfer %s", err.Error())
		return
	}

	// 4. send transaction
	log.Infof("executing %s", x.Record)
	tx, err := x.Transactor.Transact(
		newTransactionCallback(id, "execute message with transfer"),
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*gethtypes.Transaction, error) {
			setValue(opts, x.Receiver.PayableValue)
			method := func() (*gethtypes.Transaction, error) {
				return x.Chain.MessageBus.ExecuteMessageWithTransfer(opts, msg, xfer, sigs, signers, powers)
			}
			return wrapRevertCheck(opts, method)
		}, x.GetTxOptions()...)
	if err != nil {
		handleExecuteMessageWithTransferError(err, id)
		return
	}

	// 5. handle transaction result
	log.Infof("executed xferMsg (id %x): txhash %x", id, tx.Hash())
	db.UpdateTx(id, tx.Hash())
}

func (e *Executor) executeMessageWithTransferRefund(x *Execution) {
	db := dal.GetDB()
	execCtx := x.Record.ExecutionContext
	id := execCtx.MessageId

	// check if the fund is ready
	bridge := x.Chain.GetMsgBridgeAddr(execCtx.Message.GetTransferType())
	transferId := eth.Bytes2Hash(execCtx.ComputeDstTransferId(bridge))
	ready, err := x.Chain.IsTransferReady(transferId, execCtx.Message.TransferType)
	if err != nil {
		log.Errorf("failed to query on-chain transfer for message (id %x, transferType %s, dstTransferId %x): %s",
			execCtx.MessageId, execCtx.Message.TransferType, transferId, err.Error())
		return
	}
	if !ready {
		log.Debugf("[skip execution] message with transfer refund (id %x) because transfer is not seen on dst chain yet", execCtx.MessageId)
		return
	}

	// re-query the execCtx from SGN since when refund messages are first queried they always don't have any sigs
	res, err := e.sgn.GetExecutionContext(eth.Bytes2Hex(id))
	if err != nil {
		log.Errorf("failed to re-query message with transfer refund (id %x) for sigs: %s", id, err.Error())
		return
	}
	if !res.HasEnoughSigs {
		log.Debugf("re-queried message with transfer refund (id %x) does not have enough sigs yet", id)
		return
	}
	err = db.UpdateExecCtx(id, res.ExecutionContext)
	if err != nil {
		log.Errorf("failed to save new execution context for message with transfer refund (id %x): %s", id, err.Error())
		return
	}
	x.Record.ExecutionContext = res.ExecutionContext
	execCtx = res.ExecutionContext

	// 1. build transaction data
	xfer := newTransferInfo(execCtx)

	// 2. update record status to executing
	err = db.UpdateStatus(id, types.ExecutionStatus_Executing)
	if err != nil {
		log.Errorf("cannot execute refund %s", err.Error())
		return
	}

	// 3. send transaction
	signers, powers := res.ChainSigners.GetAddrsPowers()
	log.Infof("executing %s", x.Record.String())
	tx, err := x.Transactor.Transact(
		newTransactionCallback(id, "execute refund"),
		func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*gethtypes.Transaction, error) {
			method := func() (*gethtypes.Transaction, error) {
				return x.Chain.MessageBus.ExecuteMessageWithTransferRefund(opts, execCtx.Message.Data, xfer, res.EnoughSigs, signers, powers)
			}
			return wrapRevertCheck(opts, method)
		}, x.GetTxOptions()...)
	if err != nil {
		handleExecuteMessageWithTransferRefundError(err, id)
		return
	}

	// 4. handle transaction result
	log.Infof("executed refund (id %x): txhash %x", id, tx.Hash())
	db.UpdateTx(id, tx.Hash())
}

func wrapRevertCheck(opts *bind.TransactOpts, method func() (*gethtypes.Transaction, error)) (*gethtypes.Transaction, error) {
	if opts.GasLimit > 0 {
		gasLimit := opts.GasLimit
		opts.GasLimit = 0 // have to set gas limit to 0, otherwise geth would skip estimateGas
		opts.NoSend = true
		_, err := method()
		if err != nil {
			return nil, fmt.Errorf("tx dry-run error: %s", err.Error())
		}
		opts.GasLimit = gasLimit
		opts.NoSend = false
	}
	return method()
}

func (e *Executor) executePegRefundMint(x *Execution, pegBridgeVersion uint32) error {
	execCtx := x.Record.ExecutionContext
	message := &execCtx.Message
	burnId := message.GetTransferRefId()

	var err error
	if e.testMode {
		err = e.sgn.InitPegRefund(burnId)
	} else {
		err = e.gateway.InitPegRefund(burnId)
	}
	if err != nil {
		return fmt.Errorf("failed to init claim refund: %s", err.Error())
	}
	var refundTxFunc types.RefundTxFunc
	switch pegBridgeVersion {
	case 0:
		refundTxFunc = x.executePegMint
	case 2:
		refundTxFunc = x.executePegV2Mint
	default:
		return fmt.Errorf("invalid bridge version %d", pegBridgeVersion)
	}
	return e.sgn.PollAndExecutePegRefundMint(burnId, message.DstChainId, refundTxFunc, execCtx.MessageId, x.executeRefund)
}

func (e *Executor) executePegRefundWithdraw(x *Execution, vaultVersion uint32) error {
	execCtx := x.Record.ExecutionContext
	message := &execCtx.Message
	depositId := message.GetTransferRefId()

	var err error
	if e.testMode {
		err = e.sgn.InitPegRefund(depositId)
	} else {
		err = e.gateway.InitPegRefund(depositId)
	}
	if err != nil {
		return fmt.Errorf("failed to init claim refund: %s", err.Error())
	}
	var refundTxFunc types.RefundTxFunc
	switch vaultVersion {
	case 0:
		refundTxFunc = x.executePegWithdraw
	case 2:
		refundTxFunc = x.executePegV2Withdraw
	default:
		return fmt.Errorf("invalid vault version %d", vaultVersion)
	}
	return e.sgn.PollAndExecutePegRefundWithdraw(depositId, message.DstChainId, refundTxFunc, execCtx.MessageId, x.executeRefund)
}

func (e *Executor) executeLiqRefundWithdraw(x *Execution) error {
	execCtx := x.Record.ExecutionContext
	receiver := execCtx.Message.Receiver
	nonce := execCtx.GetTransfer().GetWdSeqNum()
	chainId := execCtx.Message.DstChainId
	return e.sgn.PollAndExecuteWithdraw(receiver, nonce, chainId, x.executeLiqWithdraw, execCtx.MessageId, x.executeRefund)
}

func (e *Executor) getMsgSignInfo(execCtx *msgtypes.ExecutionContext) (msg []byte, sigs [][]byte, signers []eth.Addr, powers []*big.Int, err error) {
	msg = execCtx.Message.Data
	chainSigners, err := e.sgn.GetChainSigners(execCtx.Message.DstChainId)
	if err != nil {
		return
	}
	// skip the pass check since the quorum check is already done at the server side
	_, sigs = cbrtypes.ValidateSignatureQuorum(execCtx.Message.Signatures, chainSigners.GetSortedSigners())
	signers, powers = chainSigners.GetAddrsPowers()
	return
}

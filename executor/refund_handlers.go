package executor

import (
	"math/big"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	"github.com/celer-network/im-executor/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (x *Execution) executeLiqWithdraw(
	opts *bind.TransactOpts, wdOnchain []byte, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) (*gethtypes.Transaction, error) {
	return x.Chain.LiqBridge.Withdraw(opts, wdOnchain, sortedSigs, signers, powers)
}

func (x *Execution) executePegWithdraw(
	opts *bind.TransactOpts, wdOnchain []byte, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) (*gethtypes.Transaction, error) {
	return x.Chain.PegVault.OriginalTokenVault.Withdraw(opts, wdOnchain, sortedSigs, signers, powers)
}

func (x *Execution) executePegV2Withdraw(
	opts *bind.TransactOpts, wdOnchain []byte, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) (*gethtypes.Transaction, error) {
	return x.Chain.PegVaultV2.OriginalTokenVaultV2.Withdraw(opts, wdOnchain, sortedSigs, signers, powers)
}

func (x *Execution) executePegMint(
	opts *bind.TransactOpts, mintOnChain []byte, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) (*gethtypes.Transaction, error) {
	return x.Chain.PegBridge.PeggedTokenBridge.Mint(opts, mintOnChain, sortedSigs, signers, powers)
}

func (x *Execution) executePegV2Mint(
	opts *bind.TransactOpts, mintOnChain []byte, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) (*gethtypes.Transaction, error) {
	return x.Chain.PegBridgeV2.PeggedTokenBridgeV2.Mint(opts, mintOnChain, sortedSigs, signers, powers)
}

func (x *Execution) executeRefund(execute types.RefundTxFunc, messageId []byte, req []byte, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) error {
	db := dal.GetDB()
	err := db.UpdateStatus(messageId, types.ExecutionStatus_Init_Refund_Executing)
	if err != nil {
		return err
	}
	tx, err := x.Transactor.Transact(&ethutils.TransactionStateHandler{
		OnMined: func(receipt *gethtypes.Receipt) {
			status := types.ExecutionStatus_Init_Refund_Failed
			if receipt.Status == gethtypes.ReceiptStatusSuccessful {
				log.Infof("Refund init (messageId %x) mined and succeeded: tx %x", messageId, receipt.TxHash)
				status = types.ExecutionStatus_Init_Refund_Executed
				// reset retry count to zero
				db.UpdateRetryCount(messageId, 0)
			} else {
				log.Errorf("Refund init (messageId %x) mined but failed: tx %x", messageId, receipt.TxHash)
			}
			db.UpdateStatus(messageId, status)
		},
		OnError: func(tx *gethtypes.Transaction, err error) {
			log.Errorf("execute refund init error: txhash %s, err %v", tx.Hash(), err)
			db.UpdateStatus(messageId, types.ExecutionStatus_Init_Refund_Failed)
		},
	}, func(transactor bind.ContractTransactor, opts *bind.TransactOpts) (*gethtypes.Transaction, error) {
		log.Debugf("refund tx gasLimit %s gasFeeCap %s gasTipCap %s gasPrice %s value %s", opts.GasLimit, opts.GasFeeCap, opts.GasTipCap, opts.GasPrice, opts.Value)
		return execute(opts, req, sortedSigs, signers, powers)
	})
	if err != nil {
		handleRefundError(err, messageId)
	} else {
		log.Infof("executed refund init (messageId %x): txhash %x", messageId, tx.Hash())
	}
	return nil
}

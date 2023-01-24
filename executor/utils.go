package executor

import (
	"math/big"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

func newTransactionCallback(id []byte, logmsg string) *ethutils.TransactionStateHandler {
	db := dal.GetDB()
	return &ethutils.TransactionStateHandler{
		OnMined: func(receipt *gethtypes.Receipt) {
			if receipt.Status == gethtypes.ReceiptStatusSuccessful {
				log.Infof("%s (id %x) mined: tx %x", logmsg, id, receipt.TxHash)
			} else {
				log.Errorf("%s (id %x) mined but failed: tx %x", logmsg, id, receipt.TxHash)
				db.UpdateStatus(id, types.ExecutionStatus_Failed)
			}
		},
		OnError: func(tx *gethtypes.Transaction, err error) {
			log.Errorf("%s error: txhash %s, err %v", logmsg, tx.Hash(), err)
			db.UpdateStatus(id, types.ExecutionStatus_Failed)
		},
	}
}

func newDelayTransactionCallback(id common.Hash, logmsg string) *ethutils.TransactionStateHandler {
	db := dal.GetDB()
	return &ethutils.TransactionStateHandler{
		OnMined: func(receipt *gethtypes.Receipt) {
			if receipt.Status == gethtypes.ReceiptStatusSuccessful {
				log.Infof("%s (delayId %x) mined: tx %x", logmsg, id, receipt.TxHash)
			} else {
				log.Errorf("%s (delayId %x) mined but failed: tx %x", logmsg, id, receipt.TxHash)
				db.UpdateDelayStatus(id, types.ExecutionStatus_Delay_Execution_Failed)
			}
		},
		OnError: func(tx *gethtypes.Transaction, err error) {
			log.Errorf("%s (delayId %x) error: txhash %s, err %v", logmsg, id, tx.Hash(), err)
			db.UpdateDelayStatus(id, types.ExecutionStatus_Delay_Execution_Failed)
		},
	}
}

func setValue(opts *bind.TransactOpts, value string) {
	if len(value) > 0 {
		val, ok := new(big.Int).SetString(value, 10)
		if ok {
			opts.Value = val
		}
	}
}

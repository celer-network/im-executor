package executor

import (
	"math/big"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func setValue(opts *bind.TransactOpts, value string) {
	if len(value) > 0 {
		val, ok := new(big.Int).SetString(value, 10)
		if ok {
			opts.Value = val
		}
	}
}

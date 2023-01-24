package executor

import (
	"strings"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/types"
	"github.com/ethereum/go-ethereum/common"
)

const MSG_ABORT_PREFIX = "MSG::ABORT:"

func handleExecuteMessageError(err error, id []byte) {
	db := dal.GetDB()
	// case 1: message already executed
	if strings.Contains(err.Error(), "message already executed") {
		log.Infof("message (id %x), message already executed", id)
		db.UpdateStatus(id, types.ExecutionStatus_Succeeded)
		return
	}
	// case 2: app contract specifically instructs the MessageBus to revert the outter tx, retry without delay
	if strings.Contains(err.Error(), MSG_ABORT_PREFIX) {
		handleAbort(err, id)
		return
	}
	// uncaught failures
	log.Errorf("cannot execute message (id %x): %s", id, err.Error())
	db.IncrAttemptAndRevert(id, types.ExecutionStatus_Unexecuted)
}

func handleExecuteDelayedMessageError(err error, id common.Hash, deadline time.Time) {
	db := dal.GetDB()
	// case 1: delayed message not exist
	if strings.Contains(err.Error(), "delayed message not exist") {
		log.Infof("delayed message (delayId %x) not exist", id)
		db.UpdateDelayStatus(id, types.ExecutionStatus_Ignored)
		return
	}
	// case 2: delayed message still locked
	if strings.Contains(err.Error(), "delayed message still locked") {
		log.Infof("delayed message (delayId %x) still locked", id)
		db.UpdateDelayDeadlineAndRevert(id, deadline)
		return
	}
	// uncaught failures
	log.Errorf("cannot execute delayed message (delayId %x): %s", id, err.Error())
	db.IncrDelayAttemptAndRevert(id, types.ExecutionStatus_Delayed)
}

func handleExecuteMessageWithTransferError(err error, id []byte) {
	db := dal.GetDB()
	// case 1: message already executed
	if strings.Contains(err.Error(), "transfer already executed") {
		log.Infof("message with transfer (id %x), transfer already executed, marking message as executed", id)
		db.UpdateStatus(id, types.ExecutionStatus_Succeeded)
		return
	}
	// case 2: relay is not ready due to rpc provider nodes inconsistency, retry without delay
	if strings.Contains(err.Error(), "bridge relay not exist") {
		log.Infof("message (id %x) transfer is seen by executor but upon execution relay does not exist. this can be due to inconsistency within the rpc provider, will retry execution", id)
		err = db.RevertStatus(id, types.ExecutionStatus_Unexecuted)
		if err != nil {
			log.Errorf("cannot revert message (id %x) status: %s", id, err.Error())
		}
		return
	}
	// case 3: app contract specifically instructs the MessageBus to revert the outter tx, retry without delay
	if strings.Contains(err.Error(), MSG_ABORT_PREFIX) {
		handleAbort(err, id)
		return
	}
	// uncaught failures
	log.Errorf("cannot execute message with transfer (id %x): %s", id, err.Error())
	db.IncrAttemptAndRevert(id, types.ExecutionStatus_Unexecuted)
}

func handleExecuteMessageWithTransferRefundError(err error, id []byte) {
	db := dal.GetDB()
	// case 1: refund already executed
	if strings.Contains(err.Error(), "transfer already executed") {
		log.Infof("message (id %x), transfer already executed", id)
		db.UpdateStatus(id, types.ExecutionStatus_Succeeded)
		return
	}
	// case 2: relay is not ready due to rpc provider nodes inconsistency, retry without delay
	if strings.Contains(err.Error(), "bridge relay not exist") {
		log.Infof("message (id %x) transfer is seen by executor but upon execution relay does not exist. this can be due to inconsistency within the rpc provider, will retry execution", id)
		err = db.RevertStatus(id, types.ExecutionStatus_Init_Refund_Executed)
		if err != nil {
			log.Errorf("cannot revert message (id %x) status: %s", id, err.Error())
		}
		return
	}
	// case 3: app contract specifically instructs the MessageBus to revert the outter tx, retry without delay
	if strings.Contains(err.Error(), MSG_ABORT_PREFIX) {
		handleAbort(err, id)
		return
	}
	// uncaught failures
	log.Errorf("cannot execute message with transfer (id %x): %s", id, err.Error())
	db.IncrAttemptAndRevert(id, types.ExecutionStatus_Init_Refund_Executed)
}

func handleRefundError(err error, id []byte) {
	db := dal.GetDB()
	// case 1: refund already executed
	if strings.Contains(err.Error(), "transfer exists") || strings.Contains(err.Error(), "record exists") {
		log.Infof("refund already executed (id %x)", id)
		db.UpdateStatus(id, types.ExecutionStatus_Init_Refund_Executed)
		return
	}
	log.Errorf("cannot execute refund (id %x): %s", id, err.Error())
	db.IncrAttemptAndRevert(id, types.ExecutionStatus_Unexecuted)
}

func handleAbort(err error, id []byte) {
	db := dal.GetDB()
	// app contract specifically instructs the MessageBus to revert the outter tx, retry without delay
	log.Infof("message (id %x) execution aborted: %s", id, err.Error())
	err = db.RevertStatus(id, types.ExecutionStatus_Unexecuted)
	if err != nil {
		log.Errorf("cannot revert message (id %x) status: %s", id, err.Error())
	}
}

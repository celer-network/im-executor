package dal

import (
	"database/sql"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/im-executor/dal/models"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	"github.com/celer-network/im-executor/types"
	"github.com/ethereum/go-ethereum/common"
)

var DelayPeriods = make(map[uint64]map[common.Address]*atomic.Uint64)

type DelayedMessage struct {
	DelayId     common.Hash
	MsgId       common.Hash
	Status      types.ExecutionStatus
	Attempts    uint64
	SrcContract common.Address
	Adapter     common.Address
	DstContract common.Address
	SrcChainID  uint64
	DstChainID  uint64
	Calldata    []byte
	Nonce       uint32
	From        time.Time
	Until       time.Time
	DelayTx     common.Hash
	ExecuteTx   common.Hash
	CreateTime  time.Time
	UpdateTime  time.Time
}

func (d *DelayedMessage) Columns() models.Columns {
	return []string{"delay_id", "msg_id", "status", "retry_count", "src_contract", "src_chain_id", "adapter", "dst_contract", "dst_chain_id", "calldata", "nonce", "from_time", "until", "delay_tx", "execute_tx", "create_time", "update_time"}
}

func (d *DelayedMessage) Scan(row models.Scanner) error {
	var (
		delayId     sql.NullString
		msgId       sql.NullString
		status      sql.NullInt64
		attempts    sql.NullInt64
		srcContract sql.NullString
		adapter     sql.NullString
		dstContract sql.NullString
		dstChainId  sql.NullInt64
		srcChainId  sql.NullInt64
		calldata    []byte
		nonce       sql.NullInt64
		from        sql.NullTime
		until       sql.NullTime
		delayTx     sql.NullString
		executeTx   sql.NullString
		createTime  sql.NullTime
		updateTime  sql.NullTime
	)
	err := row.Scan(&delayId, &msgId, &status, &attempts, &srcContract, &srcChainId, &adapter, &dstContract, &dstChainId, &calldata, &nonce, &from, &until, &delayTx, &executeTx, &createTime, &updateTime)
	if err != nil {
		return err
	}
	d.DelayId = eth.Hex2Hash(delayId.String)
	d.MsgId = eth.Hex2Hash(msgId.String)
	d.Status = types.ExecutionStatus(status.Int64)
	d.Attempts = uint64(attempts.Int64)
	d.SrcContract = eth.Hex2Addr(srcContract.String)
	d.Adapter = eth.Hex2Addr(adapter.String)
	d.DstContract = eth.Hex2Addr(dstContract.String)
	d.DstChainID = uint64(dstChainId.Int64)
	d.SrcChainID = uint64(srcChainId.Int64)
	d.Calldata = calldata
	d.Nonce = uint32(nonce.Int64)
	d.From = from.Time
	d.Until = until.Time
	d.DelayTx = eth.Hex2Hash(delayTx.String)
	d.ExecuteTx = eth.Hex2Hash(executeTx.String)
	d.CreateTime = createTime.Time
	d.UpdateTime = updateTime.Time
	return nil
}

func (d *DelayedMessage) String() string {
	attempt := fmt.Sprintf("attempt [%d/%d]", d.Attempts+1, types.MaxExecuteRetry)
	return fmt.Sprintf("delayed message, from %d:%x to %d:%x, id %x, delay tx %x, status %s, %s", d.SrcChainID, d.SrcContract, d.DstChainID, d.DstContract, d.MsgId, d.DelayTx, d.Status, attempt)
}

/**************************************************************************************************
 * Read methods
 **************************************************************************************************/

func (db *DB) GetDelayedMessagesToExecute() []*DelayedMessage {
	q := fmt.Sprintf(`SELECT %s FROM delay_message WHERE status = $1 and until < now()`, (&DelayedMessage{}).Columns())

	rows, err := db.Db.Query(q, types.ExecutionStatus_Delayed)
	if err != nil {
		log.Errorf("failed to get delayed messages with status %d: %v", types.ExecutionStatus_Delayed, err)
	}
	delayedMessages := []*DelayedMessage{}
	for rows.Next() {
		message := &DelayedMessage{}
		err := message.Scan(rows)
		if err != nil {
			log.Error(err)
			continue
		}
		delayedMessages = append(delayedMessages, message)
	}
	return delayedMessages
}

/**************************************************************************************************
 * Write methods
 **************************************************************************************************/

func (db *DB) InsertDelayedMessage(delayId, msgId common.Hash, srcContract, adatper, dstContract common.Address, srcChainId, dstChainId uint64, calldata []byte, nonce uint32, until time.Time, delayTx common.Hash) error {
	q := `INSERT INTO delay_message (delay_id, msg_id, status, retry_count, src_contract, src_chain_id, adapter, dst_contract, dst_chain_id, calldata, nonce, from_time, until, delay_tx, create_time, update_time)
			VALUES ($1, $2, 10, 0, $3, $4, $5, $6, $7, $8, $9, now(), $10, $11, now(), now()) ON CONFLICT DO NOTHING`

	_, err := db.Db.Exec(
		q,
		delayId.String(),
		msgId.String(),
		srcContract.String(),
		srcChainId,
		adatper.String(),
		dstContract.String(),
		dstChainId,
		calldata,
		nonce,
		until,
		delayTx.String(),
	)
	return err
}

func (db *DB) UpdateDelayExecuteTx(id, tx common.Hash) error {
	q := `UPDATE delay_message set execute_tx = $2, update_time = now() where delay_id = $1`
	_, err := db.Db.Exec(q, id.String(), tx.String())
	if err != nil {
		log.Errorf("failed to update execute tx for delayed message (delayId %x): %v", id, err)
	}
	return err
}

func (db *DB) UpdateDelayDeadlineAndRevert(id common.Hash, deadline time.Time) error {
	q := `UPDATE delay_message set until = $2, update_time = now() where delay_id = $1`
	_, err := db.Db.Exec(q, id.String(), deadline)
	if err != nil {
		log.Errorf("failed to update delay deadline for delayed message (delayId %x): %v", id, err)
		// won't return here, let status be reverted
	}
	err = db.RevertDelayStatus(id, types.ExecutionStatus_Delayed)
	if err != nil {
		log.Errorf("cannot revert delayed message (delayId %x) status: %s", id, err.Error())
	}
	return err
}

// UpdateDelayStatus updates the delayed message's status.
// `id` is delayId
func (db *DB) UpdateDelayStatus(id common.Hash, status types.ExecutionStatus) error {
	q := `SELECT status FROM delay_message WHERE delay_id = $1`
	var oldStatus uint64
	err := db.Db.QueryRow(q, id.String()).Scan(&oldStatus)
	if err != nil {
		log.Errorf("cannot query delayed message (delayId %x): %s", id, err.Error())
		return err
	}
	if uint64(status) <= oldStatus {
		log.Debugf("Skipping updating delayed message (delayId %x) because old status %d >= new status %d", id, oldStatus, status)
		return nil
	}
	q = `UPSERT INTO delay_message (delay_id, status, update_time) VALUES ($1, $2, now())`
	_, err = db.Db.Exec(q, id.String(), status)
	if err != nil {
		log.Errorf("cannot update delayed message status (delayId %x): %s", id, err.Error())
		return err
	}
	log.Infof("delay_message (delayId %x) status changed from %s to %s", id, types.ExecutionStatus(oldStatus), status)
	return nil
}

func (db *DB) RevertDelayStatus(id common.Hash, status types.ExecutionStatus) error {
	if status != types.ExecutionStatus_Delayed {
		return fmt.Errorf("revert delay status to %d is forbidden", status)
	}
	q := `UPDATE delay_message SET status = $1, update_time = now() where delay_id = $2`
	res, err := db.Db.Exec(q, status, id.String())
	log.Infof("delayed message (delayId %x) status reverted to %s", id, status)
	return sqldb.ChkExec(res, err, 1, "RevertDelayStatus")
}

func (db *DB) IncreaseDelayRetryCount(id common.Hash) (newCount uint64) {
	q := `SELECT retry_count FROM delay_message WHERE delay_id = $1`
	var oldCount uint64
	err := db.Db.QueryRow(q, id.String()).Scan(&oldCount)
	if err != nil {
		log.Errorf("cannot increase delayed message (delayId %x) retry count: %v", id, err)
		return 0
	}
	newCount = oldCount + 1
	q = `UPDATE delay_message SET retry_count = $1, update_time = now() where delay_id = $2`
	res, err := db.Db.Exec(q, newCount, id.String())
	if e := sqldb.ChkExec(res, err, 1, "IncreaseDelayRetryCount"); e != nil {
		log.Errorf("cannot increase delayed message (delayId %x) retry count: %v", id, e)
		return 0
	}
	return newCount
}

func (db *DB) IncrDelayAttemptAndRevert(id common.Hash, revertStatus types.ExecutionStatus) error {
	increasedRetryCount := db.IncreaseDelayRetryCount(id)
	if increasedRetryCount > types.MaxExecuteRetry {
		log.Warnf("delayed message (delayId %x) hit max retry count %d. it is marked as failed", id, types.MaxExecuteRetry)
		err := db.UpdateDelayStatus(id, types.ExecutionStatus_Delay_Execution_Failed)
		if err != nil {
			log.Errorf("cannot update delayed message (delayId %x) status: %s", id, err.Error())
			return err
		}
		return nil
	}
	err := db.RevertDelayStatus(id, revertStatus)
	if err != nil {
		log.Errorf("cannot revert delayed message (delayId %x) status: %s", id, err.Error())
		return err
	}
	return nil
}

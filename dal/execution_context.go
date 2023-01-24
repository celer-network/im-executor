package dal

import (
	"fmt"
	"strings"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/goutils/sqldb"
	"github.com/celer-network/im-executor/dal/models"
	msgtypes "github.com/celer-network/im-executor/sgn-v2/x/message/types"
	"github.com/celer-network/im-executor/types"
	"github.com/ethereum/go-ethereum/common"
)

/**************************************************************************************************
 * Read methods
 **************************************************************************************************/

type ExecutionRecordQuery struct {
	ID    []byte
	SrcTx string
	DstTx string
}

func (db *DB) GetExecutionRecord(q *ExecutionRecordQuery) (*models.ExecutionRecord, error) {
	record := &models.ExecutionRecord{}
	where := &WhereBuilder{}
	if len(q.ID) > 0 {
		where.And("id = '\\x%x'", q.ID)
	}
	if q.SrcTx != "" {
		where.And("src_tx = '%s'", q.SrcTx)
	}
	if q.DstTx != "" {
		where.And("tx = '%s'", q.DstTx)
	}

	query := fmt.Sprintf(`SELECT %s FROM execution_context WHERE %s`, record.Columns(), where)

	row := db.Db.QueryRow(query)
	err := record.Scan(row)
	if err != nil {
		return nil, err
	}
	return record, nil
}

type ExecutionRecordsQuery struct {
	SrcTx    string
	Statuses []int
}

func (db *DB) GetExecutionRecords(q *ExecutionRecordsQuery) ([]*models.ExecutionRecord, error) {
	where := &WhereBuilder{}
	if len(q.SrcTx) > 0 {
		where.And("tx = '%s'", q.SrcTx)
	}
	if len(q.Statuses) > 0 {
		intStrs := []string{}
		for _, s := range q.Statuses {
			intStrs = append(intStrs, fmt.Sprintf("%d", s))
		}
		where.And("status in (%s)", strings.Join(intStrs, ","))
	}

	query := fmt.Sprintf(`SELECT %s FROM execution_context WHERE %s`, (&models.ExecutionRecord{}).Columns(), where)

	rows, err := db.Db.Query(query)
	if err != nil {
		return nil, err
	}
	records := []*models.ExecutionRecord{}
	for rows.Next() {
		record := &models.ExecutionRecord{}
		err := record.Scan(rows)
		if err != nil {
			log.Error(err)
			continue
		}
		records = append(records, record)
	}
	return records, nil
}

func (db *DB) GetExecutionRecordByID(id []byte) *models.ExecutionRecord {
	record := &models.ExecutionRecord{}
	q := fmt.Sprintf(`SELECT %s FROM execution_context WHERE id = $1`, record.Columns())
	row := db.Db.QueryRow(q, id)
	err := record.Scan(row)
	if err != nil {
		log.Error(err)
		return nil
	}
	return record
}

func (db *DB) GetExecutionRecordsToExecute() []*models.ExecutionRecord {
	q := fmt.Sprintf(`SELECT %s FROM execution_context WHERE status in ($1, $2)`, (&models.ExecutionRecord{}).Columns())

	rows, err := db.Db.Query(q, types.ExecutionStatus_Unexecuted, types.ExecutionStatus_Init_Refund_Executed)
	if err != nil {
		log.Errorf("failed to get execution context with status %d and %d: %v",
			types.ExecutionStatus_Unexecuted, types.ExecutionStatus_Init_Refund_Executed, err)
	}
	records := []*models.ExecutionRecord{}
	for rows.Next() {
		record := &models.ExecutionRecord{}
		err := record.Scan(rows)
		if err != nil {
			log.Error(err)
			continue
		}
		records = append(records, record)
	}
	return records
}

/**************************************************************************************************
 * Write methods
 **************************************************************************************************/

func (db *DB) SaveExecutionContexts(execCtxs []*msgtypes.ExecutionContext) {
	for i := range execCtxs {
		execCtx := execCtxs[i]
		err := db.SaveExecutionContext(execCtx)
		if err != nil {
			log.Errorln("could not save execution context: failed to exec", err)
			continue
		}
	}
}

func (db *DB) SaveExecutionContext(execCtx *msgtypes.ExecutionContext) error {
	execCtxBytes, err := execCtx.Marshal()
	if err != nil {
		return err
	}
	q := `INSERT INTO execution_context (id, exec_ctx, chain_id, src_chain_id, src_tx, create_time, update_time) 
			VALUES ($1, $2, $3, $4, $5, now(), now()) ON CONFLICT DO NOTHING`

	_, err = db.Db.Exec(
		q,
		execCtx.MessageId,
		execCtxBytes,
		execCtx.Message.DstChainId,
		execCtx.Message.SrcChainId,
		execCtx.Message.SrcTxHash,
	)
	return err
}

func (db *DB) UpdateTx(id []byte, tx common.Hash) error {
	q := `UPDATE execution_context set tx = $2, update_time = now() where id = $1`
	_, err := db.Db.Exec(q, id, tx.String())
	if err != nil {
		log.Errorf("failed to update tx for execution_context (id %x)", id)
	}
	return err
}

// UpdateStatus updates the execution status of a message.
// `id` is computed differently for each type of message:
// message associated with peg mint: id = mintId = hash(account, token, amount, depositor, refChainId, refId)
// message associated with peg withdraw: id = mintId, same as peg mint
// message associated with liquidity send: id = dstTransferId = hash(sender, receiver, token, amount, srcChainId, dstChainId, srcTransferId)
// message associated with liquidity withdraw: id = wdId = hash(chainid, seqnum, receiver, token, amount)
// all above ids are again hashed with bridge address id = hash(bridgeAddr, id)
// no transfer associated: id = messageId
func (db *DB) UpdateStatus(id []byte, status types.ExecutionStatus) error {
	q := `SELECT status FROM execution_context WHERE id = $1`
	var oldStatus uint64
	err := db.Db.QueryRow(q, id).Scan(&oldStatus)
	if err != nil {
		log.Errorf("cannot query message (id %x): %s", id, err.Error())
		return err
	}
	if uint64(status) <= oldStatus {
		log.Debugf("Skipping updating execution_context (id %x) because old status %d >= new status %d", id, oldStatus, status)
		return nil
	}
	q = `UPSERT INTO execution_context (id, status, update_time) VALUES ($1, $2, now())`
	_, err = db.Db.Exec(q, id, status)
	if err != nil {
		log.Errorf("cannot update message status (id %x): %s", id, err.Error())
		return err
	}
	log.Infof("execution_context (id %x) status changed from %s to %s", id, types.ExecutionStatus(oldStatus), status)
	return nil
}

func (db *DB) RevertStatus(id []byte, status types.ExecutionStatus) error {
	if status != types.ExecutionStatus_Unexecuted && status != types.ExecutionStatus_Init_Refund_Executed {
		return fmt.Errorf("revert status to %d is forbidden", status)
	}
	q := `UPDATE execution_context SET status = $1, update_time = now() where id = $2`
	res, err := db.Db.Exec(q, status, id)
	log.Infof("message (id %x) status reverted to %s", id, status)
	return sqldb.ChkExec(res, err, 1, "RevertStatus")
}

func (db *DB) IncreaseRetryCount(id []byte) (newCount uint64) {
	q := `SELECT retry_count FROM execution_context WHERE id = $1`
	var oldCount uint64
	err := db.Db.QueryRow(q, id).Scan(&oldCount)
	if err != nil {
		log.Errorf("cannot increase message (id %x) retry count: %v", id, err)
		return 0
	}
	newCount = oldCount + 1
	q = `UPDATE execution_context SET retry_count = $1, update_time = now() where id = $2`
	res, err := db.Db.Exec(q, newCount, id)
	if e := sqldb.ChkExec(res, err, 1, "IncreaseRetryCount"); e != nil {
		log.Errorf("cannot increase message (id %x) retry count: %v", id, e)
		return 0
	}
	return newCount
}

func (db *DB) UpdateRetryCount(id []byte, retryCount uint64) {
	q := `UPDATE execution_context SET retry_count = $1, update_time = now() where id = $2`
	res, err := db.Db.Exec(q, retryCount, id)
	if e := sqldb.ChkExec(res, err, 1, "UpdateRetryCount"); e != nil {
		log.Errorf("cannot update message (id %x) retry count: %v", id, e)
	}
}

func (db *DB) IncrAttemptAndRevert(id []byte, revertStatus types.ExecutionStatus) error {
	increasedRetryCount := db.IncreaseRetryCount(id)
	if increasedRetryCount > types.MaxExecuteRetry {
		log.Warnf("message (id %x) hit max retry count %d. it is marked as failed", id, types.MaxExecuteRetry)
		err := db.UpdateStatus(id, types.ExecutionStatus_Failed)
		if err != nil {
			log.Errorf("cannot update message (id %x) status: %s", id, err.Error())
			return err
		}
		return nil
	}
	err := db.RevertStatus(id, revertStatus)
	if err != nil {
		log.Errorf("cannot revert message (id %x) status: %s", id, err.Error())
		return err
	}
	return nil
}

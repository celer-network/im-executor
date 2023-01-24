package models

import (
	"database/sql"
	fmt "fmt"
	"time"

	"github.com/celer-network/im-executor/sgn-v2/eth"
	msgtypes "github.com/celer-network/im-executor/sgn-v2/x/message/types"
	"github.com/celer-network/im-executor/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogo/protobuf/proto"
)

type ExecutionRecord struct {
	ID               common.Hash
	ExecutionContext *msgtypes.ExecutionContext
	ExecutionStatus  types.ExecutionStatus
	Attempts         uint64
	ChainID          uint64
	SrcChainID       uint64
	TX               common.Hash
	SrcTx            common.Hash
	CreateTime       time.Time
	UpdateTime       time.Time
}

var _ Record = &ExecutionRecord{}

func (r *ExecutionRecord) Columns() Columns {
	return []string{"id", "exec_ctx", "status", "retry_count", "chain_id", "src_chain_id", "tx", "src_tx", "create_time", "update_time"}
}

func (r *ExecutionRecord) Scan(row Scanner) error {
	var (
		id           []byte
		execCtxBytes []byte
		status       sql.NullInt64
		attempts     sql.NullInt64
		chainId      sql.NullInt64
		srcChainId   sql.NullInt64
		dstTx        sql.NullString
		srcTx        sql.NullString
		createTime   sql.NullTime
		updateTime   sql.NullTime
	)
	err := row.Scan(&id, &execCtxBytes, &status, &attempts, &chainId, &srcChainId, &dstTx, &srcTx, &createTime, &updateTime)
	if err != nil {
		return err
	}
	execCtx := &msgtypes.ExecutionContext{}
	err = proto.Unmarshal(execCtxBytes, execCtx)
	if err != nil {
		return err
	}
	r.ID = eth.Bytes2Hash(id)
	r.ExecutionContext = execCtx
	r.ExecutionStatus = types.ExecutionStatus(status.Int64)
	r.Attempts = uint64(attempts.Int64)
	r.ChainID = uint64(chainId.Int64)
	r.SrcChainID = uint64(srcChainId.Int64)
	r.TX = eth.Hex2Hash(dstTx.String)
	r.SrcTx = eth.Hex2Hash(srcTx.String)
	r.CreateTime = createTime.Time
	r.UpdateTime = updateTime.Time
	return nil
}

func (e *ExecutionRecord) GetLastExecutableStatus() types.ExecutionStatus {
	if e.ExecutionContext.Message.SrcChainId == e.ExecutionContext.Message.DstChainId &&
		(e.ExecutionStatus == types.ExecutionStatus_Executing ||
			e.ExecutionStatus == types.ExecutionStatus_Failed) {
		return types.ExecutionStatus_Init_Refund_Executed
	}
	return types.ExecutionStatus_Unexecuted
}

func (e *ExecutionRecord) String() string {
	xfer := e.ExecutionContext.Transfer
	msg := e.ExecutionContext.Message

	execType := "message"
	route := ""
	amount := ""
	attempt := fmt.Sprintf("attempt [%d/%d]", e.Attempts+1, types.MaxExecuteRetry)

	if msg.SrcChainId == msg.DstChainId {
		execType = fmt.Sprintf("refund -> %d", msg.DstChainId)
	} else if len(e.ExecutionContext.GetTransfer().GetAmount()) > 0 {
		execType = "message with transfer"
		route = fmt.Sprintf(" %d->%d", msg.SrcChainId, msg.DstChainId)
		amount = fmt.Sprintf(" amount %s token %x", xfer.Amount, xfer.Token)
	}
	srcTxInfo := ""
	if e.SrcTx != eth.ZeroHash {
		srcTxInfo = fmt.Sprintf(" src tx %s", e.SrcTx)
	}
	dstTxInfo := ""
	if e.TX != eth.ZeroHash {
		dstTxInfo = fmt.Sprintf(" dst tx %s", e.TX)
	}
	return fmt.Sprintf("%s%s%s%s%s id %x status %s, %s",
		execType, route, amount, srcTxInfo, dstTxInfo, e.ExecutionContext.MessageId, e.ExecutionStatus, attempt)
}

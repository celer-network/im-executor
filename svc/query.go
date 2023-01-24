package svc

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/dal/models"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	msgtypes "github.com/celer-network/im-executor/sgn-v2/x/message/types"
)

func (s *ExecutorService) QueryMessage(messageId string) (record *models.ExecutionRecord, msg *msgtypes.Message, err error) {
	if messageId == "" {
		return nil, nil, fmt.Errorf("messageId is required")
	}
	record = s.db.GetExecutionRecordByID(eth.Hex2Bytes(messageId))
	msg, err = s.sgn.GetMessage(messageId)
	if err != nil {
		log.Warnf("cannot query message from SGN: %s", err.Error())
	}
	return
}

func (s *ExecutorService) QueryExecutionRecord(query *dal.ExecutionRecordQuery) (*models.ExecutionRecord, error) {
	return s.db.GetExecutionRecord(query)
}

func (s *ExecutorService) QueryExecutionRecords(query *dal.ExecutionRecordsQuery) ([]*models.ExecutionRecord, error) {
	return s.db.GetExecutionRecords(query)
}

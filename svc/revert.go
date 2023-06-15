package svc

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/dal"
)

func (s *ExecutorService) RevertToLastExecutableStatus(query *dal.ExecutionRecordQuery) error {
	r, err := s.db.GetExecutionRecord(query)
	if err != nil {
		return err
	}
	if r == nil {
		return fmt.Errorf("execution record not found with query %+v", query)
	}
	err = s.db.RevertStatus(r.ID.Bytes(), r.GetLastExecutableStatus())
	if err != nil {
		return fmt.Errorf("failed to revert status: %s", err)
	}
	return nil
}

func (s *ExecutorService) RevertToLastExecutableStatusBySrcTx(srcTx string) error {
	records, err := s.db.GetExecutionRecords(&dal.ExecutionRecordsQuery{SrcTx: srcTx})
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return fmt.Errorf("execution records not found with srcTx %s", srcTx)
	}
	for _, r := range records {
		err = s.db.RevertStatus(r.ID.Bytes(), r.GetLastExecutableStatus())
		if err != nil {
			log.Errorf("failed to revert status for record %s: %s", r.ID, err.Error())
		}
	}
	return nil
}

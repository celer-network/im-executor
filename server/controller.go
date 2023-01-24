package server

import (
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/server/models"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	"github.com/celer-network/im-executor/svc"
)

func (s *RestServer) handleGetMessage(r *models.GetMessageRequest) (*models.GetMessageResponse, error) {
	if r.ID == "" {
		return nil, fmt.Errorf("param messageId is required")
	}
	record, msg, _ := s.executorSvc.QueryMessage(r.ID)
	return &models.GetMessageResponse{
		ExecutionRecord: record,
		Message:         msg,
	}, nil
}

func (s *RestServer) handleUnstuckTx(r *models.UnstuckTxRequest) (*models.UnstuckTxResponse, error) {
	log.Infof("handle /unstuck, req: %+v", r)
	gasOpts := &svc.GasPriceOptions{}
	gasOpts.GasPrice, _ = new(big.Int).SetString(r.GasPrice, 10)
	gasOpts.GasTipCap, _ = new(big.Int).SetString(r.GasTipCap, 10)
	gasOpts.GasFeeCap, _ = new(big.Int).SetString(r.GasFeeCap, 10)

	if r.ID != "" || r.DstTx != "" {
		query := &dal.ExecutionRecordQuery{
			ID:    eth.Hex2Bytes(r.ID),
			DstTx: r.DstTx,
		}
		replacementTx, err := s.executorSvc.UnstuckTxAndRevertStatus(query, r.Account, gasOpts, false)
		return &models.UnstuckTxResponse{
			ReplacementTx: replacementTx.Hex(),
		}, err
	}
	if r.SrcTx != "" {
		replacementTx, err := s.executorSvc.UnstuckTxAndRevertStatusBySrcTx(r.SrcTx, r.Account, gasOpts, false)
		return &models.UnstuckTxResponse{
			ReplacementTx: replacementTx.Hex(),
		}, err
	}
	return nil, fmt.Errorf("one of id, srcTx, or dstTx is required")
}

func (s *RestServer) handleRevertExecutionStatus(r *models.RevertExecutionStatusRequest) (*models.RevertExecutionStatusResponse, error) {
	log.Infof("handle /revert-execution-status, req: %v", r)
	if r.ID != "" || r.DstTx != "" {
		query := &dal.ExecutionRecordQuery{
			ID:    eth.Hex2Bytes(r.ID),
			DstTx: r.DstTx,
		}
		err := s.executorSvc.RevertToLastExecutableStatus(query)
		return &models.RevertExecutionStatusResponse{}, err
	}
	if r.SrcTx != "" {
		err := s.executorSvc.RevertToLastExecutableStatusBySrcTx(r.SrcTx)
		return &models.RevertExecutionStatusResponse{}, err
	}
	return nil, fmt.Errorf("one of id, srcTx, or dstTx is required")
}

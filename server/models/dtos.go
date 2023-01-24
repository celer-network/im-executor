package models

import (
	"github.com/celer-network/im-executor/dal/models"
	"github.com/celer-network/im-executor/sgn-v2/x/message/types"
)

type ServerResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type GetMessageRequest struct {
	ID string `json:"id"`
}

type GetMessageResponse struct {
	ExecutionRecord *models.ExecutionRecord
	Message         *types.Message
}

type UnstuckTxRequest struct {
	ID      string `json:"id"`
	SrcTx   string `json:"srcTx"`
	DstTx   string `json:"dstTx"`
	Account string `json:"account"`
	// legacy
	GasPrice string `json:"gasPrice"`
	// eip1559
	GasTipCap string `json:"gasTipCap"`
	GasFeeCap string `json:"gasFeeCap"`
}

type UnstuckTxResponse struct {
	ReplacementTx string `json:"replacementTx"`
}

type RevertExecutionStatusRequest struct {
	ID    string `json:"id"`
	SrcTx string `json:"srcTx"`
	DstTx string `json:"dstTx"`
}

type RevertExecutionStatusResponse struct {
}

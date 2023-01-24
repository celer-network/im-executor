package executor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/accounts"
	"github.com/celer-network/im-executor/chains"
	"github.com/celer-network/im-executor/contracts"
	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/dal/models"
	"github.com/celer-network/im-executor/sgn"
	msgtypes "github.com/celer-network/im-executor/sgn-v2/x/message/types"
	"github.com/celer-network/im-executor/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/viper"
)

type Executor struct {
	retry       *types.RetryConfig
	sgn         *sgn.SgnClient
	gateway     *sgn.GatewayClient
	accounts    *accounts.Accounts
	wg          sync.WaitGroup
	wg2         sync.WaitGroup
	parallelism int
	testMode    bool
	autoRefund  bool
}

func NewExecutor(gatewayClient *sgn.GatewayClient, sgnClient *sgn.SgnClient, accs *accounts.Accounts, testMode bool) *Executor {
	autoRefundEnabled := viper.GetBool(types.KeyEnableAutoRefund)
	if autoRefundEnabled {
		log.Infoln("auto refund enabled")
	} else {
		log.Infoln("auto refund disabled")
	}
	return &Executor{
		retry:       loadRetryConfig(),
		sgn:         sgnClient,
		gateway:     gatewayClient,
		accounts:    accs,
		parallelism: 10,
		testMode:    testMode,
		autoRefund:  autoRefundEnabled,
	}
}

func loadRetryConfig() *types.RetryConfig {
	c := &types.RetryConfig{}
	err := viper.UnmarshalKey(types.KeyExecutorRetry, c)
	if err != nil {
		log.Fatalf("error marshalling retry config")
	}
	c.ApplyDefaults()
	c.PrettyLog()
	return c
}

func (e *Executor) Start() {
	done := make(chan bool)
	go e.startFetchingExecCtxsFromSgn()
	go e.startExecuting()
	go chains.StartMonitoring(e.accounts.ReceiverContracts(), e.accounts.Addresses())
	log.Info("executor started")
	<-done
}

func (e *Executor) startFetchingExecCtxsFromSgn() {
	db := dal.GetDB()
	log.Infoln("Start fetching execution contexts from SGN")
	for {
		time.Sleep(8 * time.Second)
		execCtxs, err := e.sgn.GetExecutionContexts(e.accounts.ReceiverContracts())
		if err != nil {
			log.Errorln("failed to get messages", err)
			continue
		}
		if len(execCtxs) == 0 {
			continue
		}
		log.Tracef("Got %d execution contexts", len(execCtxs))
		execCtxsToSave := []*msgtypes.ExecutionContext{}
		for i := range execCtxs {
			execCtxsToSave = append(execCtxsToSave, &execCtxs[i])
		}
		db.SaveExecutionContexts(execCtxsToSave)
	}
}

func (e *Executor) startExecuting() {
	log.Infoln("Start executing")
	db := dal.GetDB()
	for {
		time.Sleep(3 * time.Second)
		executions := make([]*Execution, 0)

		records := db.GetExecutionRecordsToExecute()
		executions = append(executions, e.RecordToExecution(records, 0)...)
		delayedMessages := db.GetDelayedMessagesToExecute()
		executions = append(executions, e.DelayedMessageToExecution(delayedMessages, 0)...)

		if len(executions) == 0 {
			continue
		}
		e.Execute(executions)
	}
}

//func (e *Executor) startProcessingExecCtxsFromDb() {
//	log.Infoln("Start processing execution contexts from DB")
//	db := dal.GetDB()
//	for {
//		time.Sleep(3 * time.Second)
//		records := db.GetExecutionRecordsToExecute()
//		if len(records) == 0 {
//			continue
//		}
//		e.Execute(records, 0)
//	}
//}

//func (e *Executor) startExecutingDelayedMessagesFromDb() {
//	log.Infoln("Start executing delayed messages from DB")
//	db := dal.GetDB()
//	for {
//		time.Sleep(10 * time.Second)
//		delayedMessages := db.GetDelayedMessagesToExecute()
//		if len(delayedMessages) == 0 {
//			continue
//		}
//		e.ExecuteDelayedMessages(delayedMessages, 0)
//	}
//}

func (e *Executor) Execute(xs []*Execution) {
	// X workers processing messages at once
	// each worker is responsible for a chunk of the msgs
	chunkSize := len(xs) / e.parallelism
	if chunkSize < 1 {
		chunkSize = 1
	}
	log.Debugf("Executing %d messages with parallelism %d, chunk size %d", len(xs), e.parallelism, chunkSize)
	workerNum := 0
	for i := 0; i < len(xs); i += chunkSize {
		end := i + chunkSize
		if end > len(xs) {
			end = len(xs)
		}
		chunk := xs[i:end]
		e.wg.Add(1)
		log.Debugf("Worker #%d executing messages [%d:%d]", workerNum, i, end)
		go func(xs []*Execution) {
			defer e.wg.Done()
			for _, x := range xs {
				e.routeExecution(x)
			}
		}(chunk)
		workerNum++
	}
	// block until the current round of msgs are all done executing
	e.wg.Wait()
}

func (e *Executor) routeExecution(x *Execution) {
	if x.DelayedMessage != nil {
		e.executeDelayedMessage(x)
		return
	}

	execCtx := x.Record.ExecutionContext
	status := x.Record.ExecutionStatus

	// same chain ids mean it's a refund
	if execCtx.Message.SrcChainId == execCtx.Message.DstChainId {
		if !e.autoRefund {
			log.Debugf("skip executing refund for message (id %x) because enable_auto_refund is off", execCtx.MessageId)
			return
		}
		if status == types.ExecutionStatus_Init_Refund_Executed {
			e.executeMessageWithTransferRefund(x)
		} else if status == types.ExecutionStatus_Unexecuted {
			e.routeInitRefund(x)
		}
		return
	}
	// handle normal execution
	switch execCtx.Message.GetTransferType() {
	case msgtypes.TRANSFER_TYPE_NULL:
		e.executeMessage(x)
	case msgtypes.TRANSFER_TYPE_LIQUIDITY_RELAY,
		msgtypes.TRANSFER_TYPE_PEG_MINT,
		msgtypes.TRANSFER_TYPE_PEG_WITHDRAW,
		msgtypes.TRANSFER_TYPE_PEG_V2_MINT,
		msgtypes.TRANSFER_TYPE_PEG_V2_WITHDRAW:
		e.executeMessageWithTransfer(x)
	default:
		log.Errorf("cannot execute message (id %x): unsupported transfer type '%s'",
			execCtx.MessageId, execCtx.Message.GetTransferType())
	}
}

func (e *Executor) routeInitRefund(x *Execution) {
	execCtx := x.Record.ExecutionContext
	var err error
	switch x.Record.ExecutionContext.Message.GetTransferType() {
	case msgtypes.TRANSFER_TYPE_LIQUIDITY_WITHDRAW:
		err = e.executeLiqRefundWithdraw(x)
	case msgtypes.TRANSFER_TYPE_PEG_MINT:
		err = e.executePegRefundMint(x, 0)
	case msgtypes.TRANSFER_TYPE_PEG_WITHDRAW:
		err = e.executePegRefundWithdraw(x, 0)
	case msgtypes.TRANSFER_TYPE_PEG_V2_MINT:
		err = e.executePegRefundMint(x, 2)
	case msgtypes.TRANSFER_TYPE_PEG_V2_WITHDRAW:
		err = e.executePegRefundWithdraw(x, 2)
	default:
		err = fmt.Errorf("cannot init refund for message (id %x): unsupported transfer type %s",
			execCtx.MessageId, execCtx.Message.GetTransferType())
	}
	if err != nil {
		log.Errorln("init refund failed", err)
	}
}

func (e *Executor) RecordToExecution(records []*models.ExecutionRecord, gasLimit uint64) []*Execution {
	xs := make([]*Execution, 0)
	for _, r := range records {
		if gasLimit == 0 && e.retry.ShouldBackoff(r.Attempts, r.UpdateTime) {
			log.Debugf("backoff for %s", r.ID)
			continue
		}
		receiver := &contracts.ContractAddress{ChainId: r.ChainID, Address: r.ExecutionContext.Message.Receiver}
		sender := &contracts.ContractAddress{ChainId: r.SrcChainID, Address: r.ExecutionContext.Message.Sender}
		chain, ok := chains.GetChain(r.ChainID)
		if !ok {
			log.Errorf("chain %d not found", r.ChainID)
			continue
		}
		x, err := e.newExecution(r.ID, sender, receiver, chain, r, nil, gasLimit)
		if err != nil {
			log.Errorf("failed to new Execution: %s", err.Error())
			continue
		}
		xs = append(xs, x)
	}
	return xs
}

func (e *Executor) DelayedMessageToExecution(delayedMessages []*dal.DelayedMessage, gasLimit uint64) []*Execution {
	xs := make([]*Execution, 0)
	for _, dm := range delayedMessages {
		if gasLimit == 0 && e.retry.ShouldBackoff(dm.Attempts, dm.UpdateTime) {
			log.Debugf("backoff for %s", dm.MsgId)
			continue
		}
		receiver := &contracts.ContractAddress{ChainId: dm.DstChainID, Address: dm.Adapter.Hex()}
		sender := &contracts.ContractAddress{ChainId: dm.SrcChainID, Address: dm.SrcContract.Hex()}
		chain, ok := chains.GetChain(dm.DstChainID)
		if !ok {
			log.Errorf("chain %d not found", dm.DstChainID)
			continue
		}
		x, err := e.newExecution(dm.DelayId, sender, receiver, chain, nil, dm, gasLimit)
		if err != nil {
			log.Errorf("failed to new Execution: %s", err.Error())
			continue
		}
		xs = append(xs, x)
	}
	return xs
}

func (e *Executor) newExecution(msgId common.Hash, sender, receiver *contracts.ContractAddress, chain *chains.Chain, record *models.ExecutionRecord, delayedMessage *dal.DelayedMessage, gasLimit uint64) (*Execution, error) {
	acc, ok := e.accounts.AccountByReceiver(receiver)
	if !ok {
		if record != nil {
			dal.GetDB().UpdateStatus(msgId.Bytes(), types.ExecutionStatus_Ignored)
		} else {
			dal.GetDB().UpdateDelayStatus(msgId, types.ExecutionStatus_Ignored)
		}
		return nil, fmt.Errorf("ignoring message/delayed-message with id/delayId %x: cannot find account by receiver %s", msgId, receiver)
	}
	allowed := acc.IsSenderAllowed(sender, receiver)
	if !allowed {
		if record != nil {
			dal.GetDB().UpdateStatus(msgId.Bytes(), types.ExecutionStatus_Ignored)
		} else {
			dal.GetDB().UpdateDelayStatus(msgId, types.ExecutionStatus_Ignored)
		}
		return nil, fmt.Errorf("ignoring message/delayed-message with id/delayId %x: sender %s not allowed", msgId, sender)
	}
	txr, ok := acc.Transactor(chain.ChainID)
	if !ok {
		return nil, fmt.Errorf("transactor not registered for chainId %d", chain.ChainID)
	}
	recvContract, ok := acc.ReceiverContract(receiver)
	if !ok {
		return nil, fmt.Errorf("receiver contract not registered for %s", receiver)
	}
	bal, err := chain.EthClient.PendingBalanceAt(context.Background(), acc.Address)
	if err != nil {
		log.Debugf("failed to query balance for account %s on chain %d", acc.ID, chain.ChainID)
	}
	log.Debugf("account %s address %s chain %d gas token balance %s", acc.ID, acc.Address, chain.ChainID, bal)
	return &Execution{
		Chain:          chain,
		Transactor:     txr,
		Receiver:       recvContract,
		Record:         record,
		DelayedMessage: delayedMessage,
		GasLimit:       gasLimit,
	}, nil
}

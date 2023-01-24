package types

import (
	"fmt"
)

type ExecutionStatus uint64

const (
	ExecutionStatus_Unknown ExecutionStatus = iota
	// initial default status
	ExecutionStatus_Unexecuted

	// status branch: if the msg is the "refund" kind of message
	// executor needs to do InitWithdraw (if liq bridge) or ClaimRefund (if peg bridge) first before executing the message
	ExecutionStatus_Init_Refund_Executing
	// executor only executes the "refund" message at msgbus if the message is in this status
	ExecutionStatus_Init_Refund_Executed
	ExecutionStatus_Init_Refund_Failed

	ExecutionStatus_Executing

	// statuses after execution at msgbus
	ExecutionStatus_Succeeded
	ExecutionStatus_Fallback
	ExecutionStatus_Failed

	// ignored due to failed sender verification
	ExecutionStatus_Ignored

	// statuses related to delay message
	ExecutionStatus_Delayed
	ExecutionStatus_Delay_Executing
	ExecutionStatus_Delay_Executed
	ExecutionStatus_Delay_Execution_Failed
)

var ExecutionStatusMap = map[ExecutionStatus]string{
	0:  "Unknown",
	1:  "Unexecuted",
	2:  "Init_Refund_Executing",
	3:  "Init_Refund_Executed",
	4:  "Init_Refund_Failed",
	5:  "Executing",
	6:  "Succeeded",
	7:  "Fallback",
	8:  "Failed",
	9:  "Ignored",
	10: "Delayed",
	11: "Delay_Executing",
	12: "Delay_Executed",
	13: "Delay_Execution_Failed",
}

// txStatus is MessageReceiver's enum TxStatus
func NewExecutionStatus(msgbusStatus uint8) (ExecutionStatus, error) {
	switch msgbusStatus {
	case 1:
		return ExecutionStatus_Succeeded, nil
	case 2:
		return ExecutionStatus_Failed, nil
	case 3:
		return ExecutionStatus_Fallback, nil
	default:
		return ExecutionStatus_Unexecuted, fmt.Errorf("cannot map TxStatus (%d) to ExecutionStatus", msgbusStatus)
	}
}

func (s ExecutionStatus) String() string {
	return ExecutionStatusMap[s]
}

func (s ExecutionStatus) IsFinal() bool {
	return s == ExecutionStatus_Succeeded ||
		s == ExecutionStatus_Fallback ||
		s == ExecutionStatus_Failed ||
		s == ExecutionStatus_Delay_Executed ||
		s == ExecutionStatus_Delay_Execution_Failed
}

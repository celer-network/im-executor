package types

import "time"

// config keys
const (
	KeyService          = "service"
	KeyEnableAutoRefund = "executor.enable_auto_refund"
	KeyExecutorRetry    = "executor.retry"

	KeyGatewayGrpcUrl = "sgnd.gateway_grpc"
	KeySgnGrpcUrl     = "sgnd.sgn_grpc"

	KeyDbUrl = "db.url"

	KeyAlert = "alert"

	KeyServerEnable = "server.enable"
	KeyServerPort   = "server.port"
)

// monitor event names
const (
	MessageBusEventExecuted = "Executed"
	LiqBridgeEventRelay     = "Relay"
	PegBridgeEventMint      = "Mint"
	PegVaultEventWithdrawn  = "Withdrawn"
)

const (
	MaxPollingRetries    = 10
	PollingInterval      = 6 * time.Second
	GatewayTimeout       = 5 * time.Second
	MaxExecuteRetry      = 15 // TODO use types.DEFAULT_MAX_ATTEMPTS instead of this
	BalanceCheckInterval = 15 * time.Minute
)

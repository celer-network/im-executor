package cmd

import (
	"github.com/celer-network/im-executor/accounts"
	"github.com/celer-network/im-executor/chains"
	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/sgn"
	"github.com/celer-network/im-executor/types"
	"github.com/spf13/viper"
)

type ExecutorContext struct {
	DB            *dal.DB
	Chains        *chains.Chains
	SgnClient     *sgn.SgnClient
	GatewayClient *sgn.GatewayClient
	Accounts      *accounts.Accounts
}

// cache for repeated ctx access
var ctx *ExecutorContext

func GetExecutorContext() *ExecutorContext {
	if ctx == nil {
		initExecutorContext()
	}
	return ctx
}

func initExecutorContext() {
	testMode := viper.GetBool(FlagTest)
	gatewayClient := sgn.NewGatewayClient(viper.GetString(types.KeyGatewayGrpcUrl), testMode)
	sgnClient := sgn.NewSgnClient(viper.GetString(types.KeySgnGrpcUrl), testMode)

	ctx = &ExecutorContext{
		DB:            dal.InitDB(),
		Chains:        chains.InitChains(),
		SgnClient:     sgnClient,
		GatewayClient: gatewayClient,
		Accounts:      accounts.NewAccounts(),
	}
}

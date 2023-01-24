package cmd

import (
	"github.com/celer-network/im-executor/alert"
	"github.com/celer-network/im-executor/executor"
	"github.com/celer-network/im-executor/server"
	"github.com/celer-network/im-executor/svc"
	"github.com/celer-network/im-executor/types"
	"github.com/celer-network/im-executor/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CmdStart get the start command
func CmdStart() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start executor service",
		Run: func(cmd *cobra.Command, args []string) {
			test, err := cmd.InheritedFlags().GetBool("test")
			utils.CheckErr(err)

			port := viper.GetInt(types.KeyServerPort)
			enableServer := viper.GetBool(types.KeyServerEnable)

			start(test, enableServer, port)
		},
	}
	return cmd
}

func start(testMode bool, enableServer bool, port int) {
	alert.InitAlertConfig()
	ctx := GetExecutorContext()

	if enableServer {
		executorSvc := svc.NewExecutorService(ctx.DB, ctx.SgnClient, ctx.Chains, ctx.Accounts)
		server := server.NewRestServer(port, ctx.SgnClient, executorSvc)
		go server.Start()
	}
	executor.NewExecutor(ctx.GatewayClient, ctx.SgnClient, ctx.Accounts, testMode).Start()
}

package cmd

import (
	"fmt"

	"github.com/celer-network/im-executor/svc"
	"github.com/celer-network/im-executor/utils"
	"github.com/spf13/cobra"
)

// queryCmd represents the query command
func CmdQuery() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Short:                      "Query message info from DB and SGN",
		SuggestionsMinimumDistance: 2,
	}
	cmd.AddCommand(CmdQueryMsg())
	return cmd
}

func CmdQueryMsg() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "message [message-id]",
		Short: "Query message status",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			messageId := args[0]
			ctx := GetExecutorContext()
			svc := svc.NewExecutorService(ctx.DB, ctx.SgnClient, ctx.Chains, ctx.Accounts)
			record, msg, err := svc.QueryMessage(messageId)
			utils.CheckErr(err)
			if record != nil {
				fmt.Printf("Message from DB: \n%s\n", record.String())
			}
			if msg != nil {
				fmt.Printf("Message from SGN: \n%s\n", msg.PrettyLog())
			}
		},
	}
	return cmd
}

package cmd

import (
	"errors"

	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	"github.com/celer-network/im-executor/svc"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(CmdDB())
}

func CmdDB() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "db",
		Short: "common db operations",
	}
	cmd.AddCommand(CmdRevertToLastExecutableStatus())
	return cmd
}

func CmdRevertToLastExecutableStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "revert",
		Short: "revert a message's status to the last executable one",
		Args: func(cmd *cobra.Command, args []string) error {
			tx, _ := cmd.Flags().GetString(FlagDstTx)
			id, _ := cmd.Flags().GetString(FlagMessageId)

			if len(tx) > 0 && len(id) > 0 {
				return errors.New("use either '--id' or '--tx'")
			}
			if len(tx) == 0 && len(id) == 0 {
				return errors.New("'--id' or '--tx' is required")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			dstTxHex, _ := cmd.Flags().GetString(FlagDstTx)
			idHex, _ := cmd.Flags().GetString(FlagMessageId)

			r := GetExecutorContext()
			s := svc.NewExecutorService(r.DB, r.SgnClient, r.Chains, r.Accounts)
			err := s.RevertToLastExecutableStatus(&dal.ExecutionRecordQuery{
				ID:    eth.Hex2Bytes(idHex),
				DstTx: dstTxHex,
			})
			return err
		},
	}
	cmd.Flags().String(FlagDstTx, "", "the tx to look up a message with")
	cmd.Flags().String(FlagMessageId, "", "the message ID to look up a message with")

	return cmd
}

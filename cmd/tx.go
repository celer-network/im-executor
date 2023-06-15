package cmd

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/dal/models"
	"github.com/celer-network/im-executor/executor"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	"github.com/celer-network/im-executor/svc"
	"github.com/spf13/cobra"
)

const flagGasLimit = "gas-limit"

func CmdTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tx",
		Short: "transactions",
		Args:  verifyArgs,
	}
	cmd.PersistentFlags().String(FlagSrcTx, "", "the src tx to look up a message with")
	cmd.PersistentFlags().String(FlagDstTx, "", "the dst tx to look up a message with")
	cmd.PersistentFlags().String(FlagMessageId, "", "the message ID to look up a message with")
	cmd.PersistentFlags().String(FlagAccount, "", "(optional) if you have multiple account configured, you need to specify a account by it's ID")

	cmd.AddCommand(
		CmdUnstuckTx(),
		CmdExecute(),
	)
	return cmd
}

func verifyArgs(cmd *cobra.Command, args []string) error {
	id, _ := cmd.Flags().GetString(FlagMessageId)
	dstTx, _ := cmd.Flags().GetString(FlagDstTx)
	srcTx, _ := cmd.Flags().GetString(FlagSrcTx)
	acc, _ := cmd.Flags().GetString(FlagAccount)

	if !(len(dstTx) != 0 || len(id) != 0 || len(srcTx) != 0) {
		return errors.New("exactly one of '--id' or '--src-tx' or '--dst-tx'")
	}
	if len(acc) == 0 {
		accs := ctx.Accounts.AllAccounts()
		if len(accs) != 1 {
			return fmt.Errorf("--account not set but there are multiple accounts configured")
		}
	}
	return nil
}

func CmdUnstuckTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "unstuck",
		Short: "increase gas price and re-execute a message",
		Long:  "sends an empty replacement tx at a nonce, then reverts the execution status of the corresponding message to Unexecuted for the next round of execution",
		Args:  verifyUnstuckTx,
		RunE:  unstuckTx,
	}
	cmd.Flags().Uint64(FlagGasPrice, 0, "[Legacy] gas price override for legacy tx")
	cmd.Flags().Uint64(FlagGasTipCap, 0, "[EIP-1559] gas tip cap override for EIP-1559 tx")
	cmd.Flags().Uint64(FlagGasFeeCap, 0, "[EIP-1559] gas price override for EIP-1559 tx")
	cmd.Flags().Uint64(FlagChainId, 0, "chain id override for the chain to send the unstuck tx to")
	return cmd
}

func verifyUnstuckTx(cmd *cobra.Command, args []string) error {
	gasPrice, _ := cmd.Flags().GetUint64(FlagGasPrice)
	gasTipCap, _ := cmd.Flags().GetUint64(FlagGasTipCap)
	gasFeeCap, _ := cmd.Flags().GetUint64(FlagGasFeeCap)
	if gasPrice > 0 && (gasTipCap > 0 || gasFeeCap > 0) {
		return errors.New("use either '--gas-price' for legacy tx or '--gas-tip-cap and --gas-fee-cap' for EIP-1559 tx")
	}
	return nil
}

func unstuckTx(cmd *cobra.Command, args []string) (err error) {
	gasPrice, _ := cmd.Flags().GetInt64(FlagGasPrice)
	gasTipCap, _ := cmd.Flags().GetInt64(FlagGasTipCap)
	gasFeeCap, _ := cmd.Flags().GetInt64(FlagGasFeeCap)

	srcTxHex, _ := cmd.Flags().GetString(FlagSrcTx)
	dstTxHex, _ := cmd.Flags().GetString(FlagDstTx)
	idHex, _ := cmd.Flags().GetString(FlagMessageId)
	acc, _ := cmd.Flags().GetString(FlagAccount)

	ctx := GetExecutorContext()
	s := svc.NewExecutorService(ctx.DB, ctx.SgnClient, ctx.Chains, ctx.Accounts)

	gasOpts := &svc.GasPriceOptions{
		GasPrice:  big.NewInt(gasPrice),
		GasTipCap: big.NewInt(gasTipCap),
		GasFeeCap: big.NewInt(gasFeeCap),
	}
	var replacementTx eth.Hash
	if idHex != "" || dstTxHex != "" {
		query := &dal.ExecutionRecordQuery{
			ID:    eth.Hex2Bytes(idHex),
			DstTx: dstTxHex,
		}
		replacementTx, err = s.UnstuckTxAndRevertStatus(query, acc, gasOpts, false)
	}
	if srcTxHex != "" {
		replacementTx, err = s.UnstuckTxAndRevertStatusBySrcTx(srcTxHex, acc, gasOpts, false)
	}
	log.Infof("replacementTx: %s", replacementTx)
	return
}

func CmdExecute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "execute",
		Short: "execute a message",
		RunE:  execute,
	}
	cmd.Flags().Uint64(flagGasLimit, 0, "gas limit override")
	return cmd
}

func execute(cmd *cobra.Command, args []string) (err error) {
	gasLimit, _ := cmd.Flags().GetUint64(flagGasLimit)
	srcTxHex, _ := cmd.Flags().GetString(FlagSrcTx)
	dstTxHex, _ := cmd.Flags().GetString(FlagDstTx)
	idHex, _ := cmd.Flags().GetString(FlagMessageId)

	ctx := GetExecutorContext()
	x := executor.NewExecutor(ctx.GatewayClient, ctx.SgnClient, ctx.Accounts, false)
	r, err := ctx.DB.GetExecutionRecord(&dal.ExecutionRecordQuery{
		ID:    eth.Hex2Bytes(idHex),
		SrcTx: srcTxHex,
		DstTx: dstTxHex,
	})
	if err != nil {
		log.Errorf("GetExecutionRecord error: %s", err.Error())
		return err
	}
	if r.ExecutionStatus != r.GetLastExecutableStatus() {
		err = ctx.DB.RevertStatus(r.ID.Bytes(), r.GetLastExecutableStatus())
		if err != nil {
			log.Warnf("cannot revert status: %s", err.Error())
		}
	}
	log.Infof("executing message with gas limit %d", gasLimit)
	x.Execute([]*models.ExecutionRecord{r}, gasLimit)
	return
}

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/sgn-v2/common"
	"github.com/celer-network/im-executor/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// to be replaced by ldflags="-X 'cmd.Version=xxx'
var Version = "dev"

func init() {
	rootCmd.PersistentFlags().String(FlagHome, os.ExpandEnv("$HOME/.executor"), "home path")
	rootCmd.PersistentFlags().Bool(FlagTest, false, "start in CI test mode (internal use only)")
	rootCmd.PersistentFlags().String(FlagLoglevel, "info", "log level")

	rootCmd.AddCommand(
		CmdVersion(),
		CmdQuery(),
		CmdStart(),
		CmdTx(),
	)
}

var rootCmd = &cobra.Command{
	Use: "executor",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		home, err := cmd.Flags().GetString(FlagHome)
		utils.CheckErr(err)
		viper.SetDefault(FlagHome, home)

		ll, err := cmd.Flags().GetString(FlagLoglevel)
		utils.CheckErr(err)
		log.SetLevelByName(ll)

		testMode, err := cmd.Flags().GetBool(FlagTest)
		utils.CheckErr(err)
		viper.SetDefault(FlagTest, testMode)

		if testMode {
			log.Infoln("test mode on")
		}
		setupConfig(home)
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func setupConfig(home string) {
	// sets account address prefix for transactors
	sdkConfig := sdk.GetConfig()
	sdkConfig.SetBech32PrefixForAccount(common.Bech32PrefixAccAddr, common.Bech32PrefixAccPub)
	sdkConfig.Seal()

	readConfig(home, "config/executor.toml")
	readConfig(home, "config/cbridge.toml")
}

func readConfig(home, relativePath string) {
	path := filepath.Join(home, relativePath)
	viper.SetConfigFile(path)
	if err := viper.MergeInConfig(); err != nil {
		log.Fatalln("failed to load", path, err)
	}
}

func CmdVersion() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Version of the binary",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(Version)
		},
	}
	return cmd
}

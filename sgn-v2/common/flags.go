package common

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

const (
	FlagEthGateway               = "eth.gateway"
	FlagEthContractCelr          = "eth.contract_address.celr"
	FlagEthContractStaking       = "eth.contract_address.staking"
	FlagEthContractSgn           = "eth.contract_address.sgn"
	FlagEthContractStakingReward = "eth.contract_address.staking_reward"
	FlagEthContractViewer        = "eth.contract_address.viewer"
	FlagEthContractGovern        = "eth.contract_address.govern"
	FlagEthSignerKeystore        = "eth.signer_keystore"
	FlagEthSignerPassphrase      = "eth.signer_passphrase"
	FlagEthValidatorAddress      = "eth.validator_address"
	FlagEthPollInterval          = "eth.poll_interval"
	FlagEthBlockDelay            = "eth.block_delay"
	FlagEthChainId               = "eth.chain_id"
	FlagEthMonitorStartBlock     = "eth.monitor_start_block"
	FlagEthMaxBlockDelta         = "eth.max_block_delta"

	// Legacy gas price flags
	FlagEthMaxGasPriceGwei   = "eth.max_gas_price_gwei"
	FlagEthMinGasPriceGwei   = "eth.min_gas_price_gwei"
	FlagEthAddGasPriceGwei   = "eth.add_gas_price_gwei"
	FlagEthForceGasPriceGwei = "eth.force_gas_price_gwei"

	// EIP-1559 gas price flags
	FlagEthMaxFeePerGasGwei         = "eth.max_fee_per_gas_gwei"
	FlagEthMaxPriorityFeePerGasGwei = "eth.max_priority_fee_per_gas_gwei"

	FlagSgnValidatorAccount           = "sgnd.validator_account"
	FlagSgnTransactors                = "sgnd.transactors"
	FlagSgnPassphrase                 = "sgnd.passphrase"
	FlagSgnChainId                    = "sgnd.chain_id"
	FlagSgnNodeURI                    = "sgnd.node_uri"
	FlagSgnNodeGrpcUrl                = "sgnd.node_grpc_url"
	FlagSgnBaseGasPrice               = "sgnd.base_gas_price"
	FlagSgnPriceUpdateUrl             = "sgnd.price_update_url"
	FlagSgnKeyringBackend             = "sgnd.keyring_backend"
	FlagSgnGasAdjustment              = "sgnd.gas_adjustment"
	FlagSgnLivenessReportEndpoint     = "sgnd.liveness_report_endpoint"
	FlagSgnConsensusLogReportEndpoint = "sgnd.consensus_log_report_endpoint"
	FlagSgnWitnessMode                = "sgnd.witness_mode"

	FlagSgnCheckIntervalSlash     = "sgnd.check_interval.slash"
	FlagSgnCheckIntervalCbridge   = "sgnd.check_interval.cbridge"
	FlagSgnCheckIntervalVerifier  = "sgnd.check_interval.verifier"
	FlagSgnCheckIntervalCbrPrice  = "sgnd.check_interval.cbr_price"
	FlagSgnCheckIntervalMsgRefund = "sgnd.check_interval.msg_refund"
	FlagSgnCheckIntervalClearMsg  = "sgnd.check_interval.clear_msg"
	FlagSgnCheckIntervalContracts = "sgnd.check_interval.contracts"
	FlagSgnCheckIntervalVeToken   = "sgnd.check_interval.ve_token"

	FlagSgnCheckIntervalSigWaitBlk           = "sgnd.check_interval.sig_wait_blk"
	FlagSgnCheckIntervalNewSyncerWaitBlk     = "sgnd.check_interval.new_syncer_wait_blk"
	FlagSgnCheckIntervalNewSyncerMaxDuration = "sgnd.check_interval.new_syncer_max_duration"

	FlagConsensusTimeoutCommit = "consensus.timeout_commit"
	FlagConsensusAVLCacheSize  = "consensus.avl_cache_size"

	FlagLogLevel = "log.level"
	FlagLogColor = "log.color"

	FlagToStartGateway                            = "gateway.start_gateway"
	FlagGatewayDbUrl                              = "gateway.db_url"
	FlagGatewayAwsS3Region                        = "gateway.aws.s3.region"
	FlagGatewayAwsS3Bucket                        = "gateway.aws.s3.bucket"
	FlagGatewayAwsKey                             = "gateway.aws.key"
	FlagGatewayAwsSecret                          = "gateway.aws.secret"
	FlagGatewayIncentiveRewardsKeystore           = "gateway.incentive_rewards.keystore"
	FlagGatewayIncentiveRewardsPassphrase         = "gateway.incentive_rewards.passphrase"
	FlagGatewayIncentiveRewardsBscChainId         = "gateway.incentive_rewards.bsc_chain_id"
	FlagGatewayIncentiveRewardsBscContractAddress = "gateway.incentive_rewards.bsc_contract_address"

	FlagMultiChain = "multichain" // array of toml tables, each table represents one chain, see common/multichain.go for details

	// because on flow chain, one account can have multiple pubkeys, we have to know which account and pubkey index to use for sending tx
	// validators should use separate tool to add pubkey to their own flow account and record the index
	FlagFlowAccount     = "nonevm.flow_account"    // hex of 8 bytes address, w/ 0x prefix, may not have leading 00, eg. 0x01 is valid
	FlagFlowPubkeyIndex = "nonevm.flow_pubkey_idx" // index of pubkey to use. this pubkey's weight must >= 1000

	FlagBlocklistEthAddrs = "blocklist.eth_addrs"

	FlagBridgeDefaultCheckInterval = "bridge.default_check_interval"

	FlagInboundLiquidityLimit = "inbound_liquidity_limit"
)

const (
	DefaultSgnGasAdjustment            = 1.5
	DefaultSgnGasLimit                 = 300000
	DefaultSgnGetAllActiveMessageLimit = 50
)

func PostCommands(cmds ...*cobra.Command) []*cobra.Command {
	for _, c := range cmds {
		c.SetErr(c.ErrOrStderr())
	}
	return cmds
}

func GetCommands(cmds ...*cobra.Command) []*cobra.Command {
	for _, c := range cmds {
		c.Flags().Int64(flags.FlagHeight, 0, "Use a specific height to query state at (this can error if the node is pruning state)")

		c.SetErr(c.ErrOrStderr())
	}
	return cmds
}

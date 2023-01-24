package chains

const FlagMultiChain = "multichain"

// represent one chain in config file, include necessary info like chainid, gateway, cbridge address etc
type OneChainConfig struct {
	ChainID                                             uint64
	Name, Gateway                                       string
	BlkInterval, BlkDelay, MaxBlkDelta, ForwardBlkDelay uint64
	GasLimit                                            uint64
	AddGasEstimateRatio                                 float64
	// Legacy gas price flags
	AddGasGwei   uint64
	MinGasGwei   uint64
	MaxGasGwei   uint64
	ForceGasGwei string
	// EIP-1559 gas price flags
	MaxFeePerGasGwei         uint64
	MaxPriorityFeePerGasGwei uint64
	// cbridge contract address
	CBridge string
	// OriginalTokenVault contract address
	OTVault string
	// PeggedTokenBridge contract address
	PTBridge string
	// OriginalTokenVaultV2 contract address
	OTVault2 string
	// PeggedTokenBridgeV2 contract address
	PTBridge2 string
	// WithdrawInbox contract address
	WdInbox string
	// MsgBus contract address
	MsgBus string
	// MsgRecvAdapters contract address
	MsgRecvAdapters []string
	// FarmingRewards contract address
	FarmingRewards string
	// VeToken contract address
	VeToken string
	// if ProxyPort > 0, a proxy with this port will be created to support some special chain such as harmony, celo.
	// chainID will be used to determined which type proxy to create, so make sure the chainID is supported in the "endpoint-proxy"
	// create a proxy to the Gateway, and eth-client will be created to "127.0.0.1:ProxyPort"
	// more detail, https://github.com/celer-network/endpoint-proxy
	ProxyPort int
}

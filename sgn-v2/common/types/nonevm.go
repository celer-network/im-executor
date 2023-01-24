package types

func IsFlowChain(chid uint64) bool {
	switch chid {
	case uint64(NonEvmChainID_FLOW_MAINNET), uint64(NonEvmChainID_FLOW_TEST), uint64(NonEvmChainID_FLOW_EMULATOR):
		return true
	default:
		return false
	}
}

func IsAptosChain(chid uint64) bool {
	switch chid {
	case uint64(NonEvmChainID_APTOS_MAINNET), uint64(NonEvmChainID_APTOS_TEST), uint64(NonEvmChainID_APTOS_LOCAL):
		return true
	default:
		return false
	}
}

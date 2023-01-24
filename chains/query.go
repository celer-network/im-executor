package chains

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	msgtypes "github.com/celer-network/im-executor/sgn-v2/x/message/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (c *Chain) IsTransferReady(transferId eth.Hash, transferType msgtypes.TransferType) (ready bool, err error) {
	switch transferType {
	case msgtypes.TRANSFER_TYPE_LIQUIDITY_RELAY:
		ready, err = c.LiqBridge.Transfers(&bind.CallOpts{}, transferId)
	case msgtypes.TRANSFER_TYPE_PEG_MINT:
		ready, err = c.PegBridge.Records(&bind.CallOpts{}, transferId)
	case msgtypes.TRANSFER_TYPE_PEG_V2_MINT:
		ready, err = c.PegBridgeV2.Records(&bind.CallOpts{}, transferId)
	case msgtypes.TRANSFER_TYPE_PEG_WITHDRAW:
		ready, err = c.PegVault.Records(&bind.CallOpts{}, transferId)
	case msgtypes.TRANSFER_TYPE_PEG_V2_WITHDRAW:
		ready, err = c.PegVaultV2.Records(&bind.CallOpts{}, transferId)
	default:
		log.Panicf("unsupported transfer type %s", transferType)
	}
	return
}

func (c *Chain) GetMsgBridgeAddr(t msgtypes.TransferType) common.Address {
	var bridgeAddr common.Address
	switch t {
	case msgtypes.TRANSFER_TYPE_NULL:
		bridgeAddr = eth.ZeroAddr
	case msgtypes.TRANSFER_TYPE_LIQUIDITY_RELAY, msgtypes.TRANSFER_TYPE_LIQUIDITY_WITHDRAW:
		bridgeAddr = c.LiqBridge.Address
	case msgtypes.TRANSFER_TYPE_PEG_MINT:
		bridgeAddr = c.PegBridge.Address
	case msgtypes.TRANSFER_TYPE_PEG_V2_MINT:
		bridgeAddr = c.PegBridgeV2.Address
	case msgtypes.TRANSFER_TYPE_PEG_WITHDRAW:
		bridgeAddr = c.PegVault.Address
	case msgtypes.TRANSFER_TYPE_PEG_V2_WITHDRAW:
		bridgeAddr = c.PegVaultV2.Address
	}
	return bridgeAddr
}

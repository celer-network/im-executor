package chains

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	msgtypes "github.com/celer-network/im-executor/sgn-v2/x/message/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type delayedTransfer interface {
	DelayedTransfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
		Receiver  common.Address
		Token     common.Address
		Amount    *big.Int
		Timestamp *big.Int
	}, error)
}

func (c *Chain) IsTransferReady(transferId eth.Hash, transferType msgtypes.TransferType) (ready bool, err error) {
	var d delayedTransfer
	switch transferType {
	case msgtypes.TRANSFER_TYPE_LIQUIDITY_RELAY:
		ready, err = c.LiqBridge.Transfers(nil, transferId)
		d = c.LiqBridge
	case msgtypes.TRANSFER_TYPE_PEG_MINT:
		ready, err = c.PegBridge.Records(nil, transferId)
		d = c.PegBridge
	case msgtypes.TRANSFER_TYPE_PEG_V2_MINT:
		ready, err = c.PegBridgeV2.Records(nil, transferId)
		d = c.PegBridgeV2
	case msgtypes.TRANSFER_TYPE_PEG_WITHDRAW:
		ready, err = c.PegVault.Records(nil, transferId)
		d = c.PegVault
	case msgtypes.TRANSFER_TYPE_PEG_V2_WITHDRAW:
		ready, err = c.PegVaultV2.Records(nil, transferId)
		d = c.PegVaultV2
	default:
		log.Panicf("unsupported transfer type %s", transferType)
	}
	if err != nil || !ready {
		return
	}
	// delayedTransfer is deleted in contract once it is executed. so no delayedTransfer
	// found means either it's not a delayed transfer or the delayed transfer is executed
	delayInfo, err := d.DelayedTransfers(nil, transferId)
	if err != nil {
		return false, err
	}
	delayed := delayInfo.Receiver != eth.ZeroAddr
	if delayed {
		log.Infof("transfer (id %s) is ready but delayed", transferId)
		return false, nil
	}
	return true, nil
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

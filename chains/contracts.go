package chains

import (
	"github.com/celer-network/im-executor/sgn-v2/eth"
	"github.com/celer-network/im-executor/utils"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Contracts struct {
	MessageBus  *eth.MsgBusContract
	LiqBridge   *eth.BridgeContract
	PegBridge   *eth.PegBridgeContract
	PegBridgeV2 *eth.PegBridgeV2Contract
	PegVault    *eth.PegVaultContract
	PegVaultV2  *eth.PegVaultV2Contract
}

func NewContracts(ec *ethclient.Client, config *OneChainConfig) *Contracts {
	contracts := &Contracts{}
	var err error
	contracts.LiqBridge, err = eth.NewBridgeContract(eth.Hex2Addr(config.CBridge), ec)
	utils.CheckErr(err)
	contracts.PegBridge, err = eth.NewPegBridgeContract(eth.Hex2Addr(config.PTBridge), ec)
	utils.CheckErr(err)
	contracts.PegVault, err = eth.NewPegVaultContract(eth.Hex2Addr(config.OTVault), ec)
	utils.CheckErr(err)
	contracts.PegBridgeV2, err = eth.NewPegBridgeV2Contract(eth.Hex2Addr(config.PTBridge2), ec)
	utils.CheckErr(err)
	contracts.PegVaultV2, err = eth.NewPegVaultV2Contract(eth.Hex2Addr(config.OTVault2), ec)
	utils.CheckErr(err)
	contracts.MessageBus, err = eth.NewMsgBusContract(eth.Hex2Addr(config.MsgBus), ec)
	utils.CheckErr(err)
	return contracts
}

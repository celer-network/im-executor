package chains

import (
	"context"
	"strings"
	"time"

	"github.com/celer-network/goutils/eth/mon2"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/alert"
	"github.com/celer-network/im-executor/contracts"
	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	"github.com/celer-network/im-executor/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func (c *Chain) startMonitoringEvents(filters contracts.ReceiverContracts) {
	c.monitorBusExecuted(filters)
}

func (c *Chain) monitorBusExecuted(filters contracts.ReceiverContracts) {
	addrConfig := mon2.PerAddrCfg{
		Addr:    c.MessageBus.Address,
		ChkIntv: time.Minute,
		AbiStr:  eth.MessageBusABI,
	}
	msgbusABI, err := abi.JSON(strings.NewReader(eth.MessageBusABI))
	if err != nil {
		log.Fatalf("failed to parse MessageBus ABI")
	}
	evExecuted, ok := msgbusABI.Events["Executed"]
	if !ok {
		log.Fatalf("failed to get Executed event MessageBus ABI")
	}
	filters = filters.GetByChain(c.ChainID)
	if len(filters) > 0 {
		// first topic would be a signature of the event
		// while we switch eventName later, so we don't filter event signature here.
		addrs := filters.AddressHashes(c.ChainID)
		addrConfig.Topics = [][]common.Hash{{evExecuted.ID}, addrs}
	} else {
		log.Infof("no address filter for chain %d, stop monitoring MessageBus Executed on this chain", c.ChainID)
		return
	}
	log.Infof("monitoring MessageBus Executed on chain %d with address filters %#s", c.ChainID, addrConfig.Topics[1])
	db := dal.GetDB()
	go c.monitor2.MonAddr(addrConfig, func(_ string, eLog ethtypes.Log) {
		e, err := c.MessageBus.ParseExecuted(eLog)
		if err != nil {
			log.Errorln("cannot parse event Executed", err)
			return
		}
		log.Infof("monitorBusExecuted: got event Executed %v", e)
		status, err := types.NewExecutionStatus(e.Status)
		if err != nil {
			log.Errorln("monitorBusExecuted: ", err)
			return
		}
		err = db.UpdateStatus(e.MsgId[:], status)
		if err != nil {
			log.Errorf("failed to update execution_context %x: %v", e.MsgId[:], err)
		}
	})
}

func (c *Chain) startMonitoringBalance(signers []eth.Addr) {
	alertConfig := alert.GetConfig()
	if alertConfig == nil || len(alertConfig.LowGasThresholds) == 0 {
		log.Infof("opt out of low gas monitoring because alert.low_gas_thresholds for chain %d is not configured", c.ChainID)
		return
	}
	for _, addr := range signers {
		log.Infof("monitoring native token balance on chain %d at %s", c.ChainID, addr)
		localAddr := addr
		go func() {
			t := time.NewTicker(types.BalanceCheckInterval)
			defer t.Stop()
			for {
				select {
				case <-c.balAlertCtl:
					log.Infof("Terminated monitor for native token balance on chain %d at %s", c.ChainID, localAddr)
					return
				case <-t.C:
					c.checkSignerBalanceThenAlert(alertConfig, localAddr)
				}
			}
		}()
	}
}

func (c *Chain) checkSignerBalanceThenAlert(config *alert.AlertConfig, addr eth.Addr) {
	threshold, found := config.GetLowGasThreshold(c.ChainID)
	if !found {
		return
	}
	balance, err := c.EthClient.PendingBalanceAt(context.Background(), addr)
	if err != nil {
		log.Errorf("failed to query pending balance at %s: %s", addr, err.Error())
		return
	}
	if balance.Cmp(threshold) < 0 {
		alert.LowGasAlert(c.ChainID, addr, balance, threshold)
	}
}

package contracts

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	comtypes "github.com/celer-network/im-executor/sgn-v2/common/types"
	"github.com/celer-network/im-executor/sgn-v2/eth"
)

const (
	DEFAULT_CONTRACT_GROUP = "default"
)

type ReceiverContract struct {
	ChainId uint64 `mapstructure:"chain_id"`
	Address string `mapstructure:"address"`
	// the payable value to add when calling executeMessage or executeMessageWithTransfer
	PayableValue string `mapstructure:"add_payable_value_for_execution"`
	// sender verification
	AllowSenderGroups []string `mapstructure:"allow_sender_groups"`
}

func (c *ReceiverContract) ContractAddress() *ContractAddress {
	return &ContractAddress{
		Address: c.Address,
		ChainId: c.ChainId,
	}
}

func (c *ReceiverContract) GetAddress() eth.Addr {
	return eth.Hex2Addr(c.Address)
}

func (c *ReceiverContract) ContractInfo() *comtypes.ContractInfo {
	return &comtypes.ContractInfo{
		ChainId: c.ChainId,
		Address: c.Address,
	}
}

type ReceiverContracts []*ReceiverContract

func (c ReceiverContracts) ContractInfoList() []*comtypes.ContractInfo {
	infos := []*comtypes.ContractInfo{}
	for _, filter := range c {
		infos = append(infos, filter.ContractInfo())
	}
	return infos
}

func (c ReceiverContracts) Get(chid uint64, addr string) (ReceiverContract, bool) {
	for _, filter := range c {
		if chid == filter.ChainId && eth.Hex2Addr(filter.Address) == eth.Hex2Addr(addr) {
			return *filter, true
		}
	}
	return ReceiverContract{}, false
}

func (c ReceiverContracts) GetByChain(chid uint64) ReceiverContracts {
	f := ReceiverContracts{}
	for _, filter := range c {
		if filter.ChainId == chid {
			f = append(f, filter)
		}
	}
	return f
}

func (c ReceiverContracts) Addresses(chid uint64) []eth.Addr {
	addrs := []eth.Addr{}
	for _, filter := range c {
		if filter.ChainId == chid {
			addrs = append(addrs, filter.GetAddress())
		}
	}
	return addrs
}

func (c ReceiverContracts) AddressHashes(chid uint64) []eth.Hash {
	addrs := []eth.Hash{}
	for _, filter := range c {
		if filter.ChainId == chid {
			addrs = append(addrs, filter.GetAddress().Hash())
		}
	}
	return addrs
}

func (c ReceiverContracts) String() string {
	str := ""
	for _, f := range c {
		s := fmt.Sprintf("- chain_id %d address %s payable_value %s allow_sender_groups %q\n",
			f.ChainId, f.Address, f.PayableValue, f.AllowSenderGroups)
		str += s
	}
	return str
}

func (c ReceiverContracts) Validate() error {
	if len(c) == 0 {
		log.Warnf("empty executor contract filter")
	}
	log.Infoln("executor will submit execution for these contracts:")
	for _, f := range c {
		if len(f.AllowSenderGroups) == 0 {
			log.Warnf("allow_sender_groups not configured for chain %d contract %s, default to submitting txs originated from all senders to this receiver contract", f.ChainId, f.Address)
		}
	}
	log.Infof("contract filters\n%s", c)
	return nil
}

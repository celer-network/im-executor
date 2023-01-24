package accounts

import (
	"fmt"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/contracts"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	"github.com/celer-network/im-executor/types"
	"github.com/spf13/viper"
)

type Accounts struct {
	accounts []*Account
	// a reverse mapping of receiversToAccounts' receiver contracts to receiversToAccounts
	receiversToAccounts map[string]*Account
}

const DEFAULT_ACCOUNT_ID = "default"

func NewAccounts() *Accounts {
	log.Infof("initializing accounts...")
	confs := loadConfigs()
	accs := &Accounts{
		receiversToAccounts: make(map[string]*Account),
	}
	for _, conf := range confs {
		acc := newAccount(conf)
		accs.registerAccount(acc)
	}
	return accs
}

func (a *Accounts) registerAccount(acc *Account) {
	a.accounts = append(a.accounts, acc)
	for _, contract := range acc.ReceiverContracts {
		a.receiversToAccounts[contract.ContractAddress().String()] = acc
	}
}

func (a *Accounts) Default() (*Account, error) {
	if len(a.accounts) != 0 {
		return nil, fmt.Errorf("cannot get default account if there is not exactly 1 account")
	}
	return a.accounts[0], nil
}

func (a *Accounts) AccountByID(id string) (*Account, bool) {
	if id == "" {
		id = DEFAULT_ACCOUNT_ID
	}
	for _, acc := range a.accounts {
		if acc.ID == id {
			return acc, true
		}
	}
	return nil, false
}

func (a *Accounts) AccountByReceiver(receiver *contracts.ContractAddress) (*Account, bool) {
	acc, ok := a.receiversToAccounts[receiver.String()]
	return acc, ok
}

func (a *Accounts) AllAccounts() []*Account {
	return a.accounts
}

// Addresses returns the address of all accounts
func (a *Accounts) Addresses() map[uint64][]eth.Addr {
	addrsMap := make(map[uint64][]eth.Addr)
	for _, acc := range a.accounts {
		for _, receiverContract := range acc.ReceiverContracts {
			if addrs, ok := addrsMap[receiverContract.ChainId]; ok {
				addrs = append(addrs, acc.Address)
				addrsMap[receiverContract.ChainId] = addrs
			} else {
				addrsMap[receiverContract.ChainId] = []eth.Addr{acc.Address}
			}
		}
	}
	return addrsMap
}

// ReceiverContracts returns all configured receiver contracts for all accounts
func (a *Accounts) ReceiverContracts() contracts.ReceiverContracts {
	receivers := contracts.ReceiverContracts{}
	for _, acc := range a.receiversToAccounts {
		receivers = append(receivers, acc.ReceiverContracts...)
	}
	return receivers
}

func loadConfigs() []*AccountConfig {
	var configs []*AccountConfig
	err := viper.UnmarshalKey(types.KeyService, &configs)
	if err != nil {
		log.Fatalf("failed to load service configs:%v", err)
		return nil
	}
	if len(configs) == 0 {
		log.Fatalf("empty [[service]] configs in executor.toml")
	}
	// check if all accounts are configured with an id
	if len(configs) > 1 {
		for _, conf := range configs {
			if len(conf.ID) == 0 {
				log.Fatalf("id field in [[service]] config in executor.toml is required if there are multiple services")
			}
		}
	} else if len(configs[0].ID) == 0 {
		// if exactly one account is configured, setup a fallback default ID if no ID is set
		configs[0].ID = DEFAULT_ACCOUNT_ID
	}
	return configs
}

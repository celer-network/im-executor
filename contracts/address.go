package contracts

import (
	"fmt"

	"github.com/celer-network/im-executor/sgn-v2/eth"
)

type ContractAddress struct {
	ChainId uint64 `mapstructure:"chain_id"`
	Address string
}

func (c *ContractAddress) GetAddress() eth.Addr {
	return eth.Hex2Addr(c.Address)
}

func (c *ContractAddress) Equal(o *ContractAddress) bool {
	return c.GetAddress() == o.GetAddress() && c.ChainId == o.ChainId
}

func (a *ContractAddress) String() string {
	return fmt.Sprintf("%d-%s", a.ChainId, eth.Hex2Addr(a.Address))
}

type ContractAddresses []*ContractAddress

func (a ContractAddresses) Has(sender *ContractAddress) bool {
	for _, s := range a {
		if s.Equal(sender) {
			return true
		}
	}
	return false
}

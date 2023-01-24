package accounts

import (
	"github.com/celer-network/im-executor/contracts"
)

type AccountConfig struct {
	ID                string                      `mapstructure:"id"`
	Keystore          string                      `mapstructure:"signer_keystore"`
	Passphrase        string                      `mapstructure:"signer_passphrase"`
	ReceiverContracts contracts.ReceiverContracts `mapstructure:"contracts"`
	SenderGroups      contracts.SenderGroups      `mapstructure:"contract_sender_groups"`
}

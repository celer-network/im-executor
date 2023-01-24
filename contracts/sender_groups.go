package contracts

import (
	"fmt"
)

type SenderGroup struct {
	Name  string
	Allow ContractAddresses
}

func (s SenderGroup) String() string {
	str := fmt.Sprintf("- sender group '%s'\n", s.Name)
	for _, a := range s.Allow {
		str += fmt.Sprintf("- allow chain %d address %s\n", a.ChainId, a.Address)
	}
	return str
}

type SenderGroups []*SenderGroup

// Validate senderGroups does not contain duplicated name.
func (s SenderGroups) Validate() error {
	groupAllows := make(map[string]ContractAddresses)
	for _, group := range s {
		if groupAllows[group.Name] != nil {
			return fmt.Errorf("duplicated sender group in executor.toml [service.contract_sender_groups]: %s", group.Name)
		}
		groupAllows[group.Name] = group.Allow
	}
	return nil
}

func (s SenderGroups) AllowedSenders(senderGroupName string) ContractAddresses {
	for _, group := range s {
		if group.Name == senderGroupName {
			return group.Allow
		}
	}
	return []*ContractAddress{}
}

func (s SenderGroups) String() string {
	str := ""
	for _, sg := range s {
		str += sg.String()
	}
	return str
}

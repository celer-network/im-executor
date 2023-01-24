package types

import (
	"fmt"
	"time"

	yaml "gopkg.in/yaml.v2"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// Parameter keys
var (
	ParamStoreKeyTriggerSignCooldown = []byte("triggersigncooldown")
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns default pegbridge parameters
func DefaultParams() Params {
	return Params{
		TriggerSignCooldown: time.Minute, // 1 minute
	}
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(
			ParamStoreKeyTriggerSignCooldown, &p.TriggerSignCooldown, validateTriggerSignCooldown),
	}
}

// Validate performs validation on farming parameters.
func (p Params) Validate() error {
	if err := validateTriggerSignCooldown(p.TriggerSignCooldown); err != nil {
		return err
	}
	return nil
}

func validateTriggerSignCooldown(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v <= 0 {
		return fmt.Errorf("trigger sign cooldown must be positive: %d", v)
	}
	return nil
}

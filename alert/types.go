package alert

import (
	"math/big"

	"github.com/celer-network/goutils/log"
)

type AlertConfig struct {
	Type             string
	Webhook          string
	LowGasThresholds []*ChainIdThreshold `mapstructure:"low_gas_thresholds"`
}

func (a *AlertConfig) GetLowGasThreshold(chid uint64) (*big.Int, bool) {
	for _, threshold := range a.LowGasThresholds {
		if threshold.ChainID == chid {
			t, ok := new(big.Int).SetString(threshold.Threshold, 10)
			if !ok {
				log.Errorf("malformed threshold %s", threshold.Threshold)
			}
			return t, true
		}
	}
	return nil, false
}

type ChainIdThreshold struct {
	ChainID   uint64 `mapstructure:"chain_id"`
	Threshold string
}

type SlackRequest struct {
	Text string `json:"text"`
}

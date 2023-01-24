package alert

import (
	"fmt"
	"math/big"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	"github.com/celer-network/im-executor/types"
	"github.com/spf13/viper"
	"gopkg.in/resty.v1"
)

var config *AlertConfig

func InitAlertConfig() {
	err := viper.UnmarshalKey(types.KeyAlert, &config)
	if err != nil {
		log.Errorf("failed to init alert config: %s", err.Error())
	}
}

func GetConfig() *AlertConfig {
	return config
}

func LowGasAlert(chid uint64, addr eth.Addr, balance, threshold *big.Int) {
	message := fmt.Sprintf("executor account %s\nbelow %s on chain %d\nnow is %s", addr, threshold, chid, balance)
	SendSlackMessage(message)
}

func SendSlackMessage(message string) {
	log.Infof("sending slack message: %s", message)
	_, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(&SlackRequest{message}).
		Post(config.Webhook)
	if err != nil {
		log.Errorf("failed to send slack message %s", err.Error())
	}
}

package common

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client"
)

const (
	retryInterval = time.Second
)

func TsSecToTime(ts uint64) time.Time {
	return time.Unix(int64(ts), 0)
}

func RobustQueryWithData(cliCtx client.Context, route string, data []byte) ([]byte, error) {
	res, _, err := cliCtx.QueryWithData(route, data)
	if err != nil {
		time.Sleep(retryInterval)
		res, _, err = cliCtx.QueryWithData(route, data)
		return res, err
	}

	return res, err
}

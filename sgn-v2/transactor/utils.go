package transactor

import (
	"fmt"
	"sync"

	"github.com/celer-network/im-executor/sgn-v2/common"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/gammazero/deque"
)

type MsgQueue struct {
	queue deque.Deque
	lock  sync.RWMutex
}

func (q *MsgQueue) PushBack(msg sdk.Msg) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue.PushBack(msg)
}

func (q *MsgQueue) PushFront(msg sdk.Msg) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue.PushFront(msg)
}

func (q *MsgQueue) PopFront() sdk.Msg {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.queue.PopFront().(sdk.Msg)
}

func (q *MsgQueue) Len() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.queue.Len()
}

func QueryAccount(cliCtx client.Context, sgnAddr string) (account *authtypes.BaseAccount, err error) {
	route := fmt.Sprintf("custom/%s/%s", authtypes.QuerierRoute, authtypes.QueryAccount)
	params := authtypes.QueryAccountRequest{
		Address: sgnAddr,
	}
	data, err := cliCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return
	}
	res, err := common.RobustQueryWithData(cliCtx, route, data)
	if err != nil {
		return
	}

	account = new(authtypes.BaseAccount)
	err = cliCtx.LegacyAmino.UnmarshalJSON(res, account)
	return
}

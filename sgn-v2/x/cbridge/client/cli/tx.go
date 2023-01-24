package cli

import (
	"github.com/celer-network/im-executor/sgn-v2/transactor"
	"github.com/celer-network/im-executor/sgn-v2/x/cbridge/types"
)

// if err not nil, should return immediately when estimate gas
func InitWithdraw(t *transactor.Transactor, req *types.MsgInitWithdraw) (resp *types.MsgInitWithdrawResp, err error) {
	req.Creator = t.Key.GetAddress().String() // make sure the msg creator is the transactor
	_, err = t.LockSendTx(req)
	return
}

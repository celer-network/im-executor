package cli

import (
	"github.com/celer-network/im-executor/sgn-v2/transactor"
	"github.com/celer-network/im-executor/sgn-v2/x/pegbridge/types"
)

func InitClaimRefund(t *transactor.Transactor, req *types.MsgClaimRefund) (resp *types.MsgClaimRefundResponse, err error) {
	req.Sender = t.Key.GetAddress().String() // make sure the msg sender is the transactor
	_, err = t.LockSendTx(req)
	return
}

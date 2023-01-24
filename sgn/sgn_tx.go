package sgn

import (
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	cbrcli "github.com/celer-network/im-executor/sgn-v2/x/cbridge/client/cli"
	cbrtypes "github.com/celer-network/im-executor/sgn-v2/x/cbridge/types"
	pegbrcli "github.com/celer-network/im-executor/sgn-v2/x/pegbridge/client/cli"
	pegbrtypes "github.com/celer-network/im-executor/sgn-v2/x/pegbridge/types"
)

func (c *SgnClient) InitWithdraw(srcXferId []byte, nonce uint64) error {
	txr := c.txrs.GetTransactor()
	wdReq := &cbrtypes.WithdrawReq{
		XferId:       eth.Bytes2Hex(srcXferId),
		ReqId:        nonce,
		WithdrawType: cbrtypes.RefundTransfer,
	}
	wdReqBytes, err := wdReq.Marshal()
	if err != nil {
		return err
	}
	msg := &cbrtypes.MsgInitWithdraw{
		WithdrawReq: wdReqBytes,
		Creator:     txr.Key.GetAddress().String(),
	}
	_, err = cbrcli.InitWithdraw(txr, msg)
	return err
}

func (c *SgnClient) InitPegRefund(refId []byte) error {
	txr := c.txrs.GetTransactor()
	msg := &pegbrtypes.MsgClaimRefund{
		RefId:  eth.Bytes2Hex(refId),
		Sender: txr.Key.GetAddress().String(),
	}
	log.Infof("init peg refund (refId %x)", refId)
	_, err := pegbrcli.InitClaimRefund(txr, msg)
	return err
}

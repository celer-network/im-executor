package sgn

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	cbrtypes "github.com/celer-network/im-executor/sgn-v2/x/cbridge/types"
	"github.com/celer-network/im-executor/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GatewayClient struct {
	conn *grpc.ClientConn
	cli  types.WebClient
}

func NewGatewayClient(gatewayUrl string, testMode bool) *GatewayClient {
	if testMode {
		return nil
	}
	log.Infof("Dialing gateway grpc: %s", gatewayUrl)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})), grpc.WithBlock()}
	//opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	conn, err := grpc.DialContext(context, gatewayUrl, opts...)
	defer cancel()
	if err != nil {
		log.Fatalln("failed to initialize gateway grpc connection", err)
	}
	return &GatewayClient{
		conn: conn,
		cli:  types.NewWebClient(conn),
	}
}

func (g *GatewayClient) InitWithdraw(srcXferId []byte, nonce uint64) error {
	wdReq := &cbrtypes.WithdrawReq{
		XferId:       eth.Bytes2Hex(srcXferId),
		ReqId:        nonce,
		WithdrawType: cbrtypes.RefundTransfer,
	}
	wdReqBytes, err := wdReq.Marshal()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), types.GatewayTimeout)
	defer cancel()
	req := &types.InitWithdrawRequest{WithdrawReq: wdReqBytes}
	res, err := g.cli.InitWithdraw(ctx, req)
	if err != nil {
		return err
	}
	if res.Err.GetCode() != types.ErrCode_ERROR_CODE_UNDEFINED {
		return fmt.Errorf("initWithdraw err: %s", res.Err.Msg)
	}
	return nil
}

func (g *GatewayClient) InitPegRefund(refId []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), types.GatewayTimeout)
	defer cancel()
	req := &types.InitPegRefundRequest{RefId: refId}
	res, err := g.cli.InitPegRefund(ctx, req)
	if err != nil {
		return err
	}
	if res.Err.GetCode() != types.ErrCode_ERROR_CODE_UNDEFINED {
		return fmt.Errorf("InitPegRefund err: %s", res.Err.Msg)
	}
	return nil
}

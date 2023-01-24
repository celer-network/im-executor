package sgn

import (
	"context"
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/sgn-v2/app"
	"github.com/celer-network/im-executor/sgn-v2/common"
	"github.com/celer-network/im-executor/sgn-v2/transactor"
	cbrtypes "github.com/celer-network/im-executor/sgn-v2/x/cbridge/types"
	pegbrtypes "github.com/celer-network/im-executor/sgn-v2/x/pegbridge/types"
	"github.com/celer-network/im-executor/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type SgnClient struct {
	txrs     *transactor.TransactorPool
	grpcConn *grpc.ClientConn
}

func NewSgnClient(sgnUrl string, testMode bool) *SgnClient {
	txrs := newSgnTransactors(testMode)
	log.Infof("Dialing sgn node grpc: %s", sgnUrl)
	var opts []grpc.DialOption
	if !testMode {
		// "an empty tls.config would use aws public ca cert as instead"
		opts = []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})), grpc.WithBlock()}
	} else {
		opts = []grpc.DialOption{grpc.WithInsecure(), grpc.WithBlock()}
	}
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	grpcConn, err := grpc.DialContext(context, sgnUrl, opts...)
	defer cancel()
	if err != nil {
		log.Fatalln("failed to initialize sgn node grpc connection", err)
	}
	return &SgnClient{txrs, grpcConn}
}

func newSgnTransactors(testMode bool) *transactor.TransactorPool {
	if !testMode {
		return nil
	}
	encoding := app.MakeEncodingConfig()
	txrAddrs := viper.GetStringSlice(common.FlagSgnTransactors)
	chainId := viper.GetString(common.FlagSgnChainId)
	nodeUri := viper.GetString(common.FlagSgnNodeURI)
	home := viper.GetString(flags.FlagHome)
	log.Infof("Initializing sgn transactors with args: home %s, chainId %s, nodeuri %s, addrs %v", home, chainId, nodeUri, txrAddrs)
	txrs := transactor.NewTransactorPool(home, chainId, encoding.Amino, encoding.Codec, encoding.InterfaceRegistry)
	err := txrs.AddTransactors(
		nodeUri,
		viper.GetString(common.FlagSgnPassphrase),
		txrAddrs)
	if err != nil {
		log.Fatalf("failed to add transactors: %s", err.Error())
	}
	log.Infof("Initialized %d sgn transactors", len(txrAddrs))
	return txrs
}

func (c *SgnClient) PollAndExecuteWithdraw(addr string, nonce uint64, chainId uint64, execute types.RefundTxFunc, messageId []byte, handler types.ExecuteRefund) error {
	for try := 1; try <= types.MaxPollingRetries; try++ {
		log.Debugf("polling withdraw status (try %d/%d): addr %s, nonce %d, chainId %d", try, types.MaxPollingRetries, addr, nonce, chainId)
		time.Sleep(types.PollingInterval)
		// poll withdraw status until its status reaches WD_WAITING_FOR_LP
		detail, status, err := c.GetXferWithdrawStatus(addr, nonce, chainId)
		if err != nil {
			return fmt.Errorf("failed to get withdraw status (addr %x, nonce %d, chainId %d): %s", addr, nonce, chainId, err.Error())
		}
		if status != cbrtypes.WithdrawStatus_WD_WAITING_FOR_LP {
			log.Debugf("withdraw status %v is not WAITING_FOR_LP yet", status)
			continue
		}
		if status == cbrtypes.WithdrawStatus_WD_COMPLETED ||
			status == cbrtypes.WithdrawStatus_WD_DELAYED ||
			status == cbrtypes.WithdrawStatus_WD_FAILED {
			log.Warnf("withdraw status %v not executable, skipping", status)
			return nil
		}
		// prepare withdraw req info
		res, err := c.GetChainSigners(chainId)
		if err != nil {
			return fmt.Errorf("failed to query chain signers with chainId %d", chainId)
		}
		signers, powers := res.GetAddrsPowers()
		wdOnchain, sortedSigs := detail.GetWdOnchain(), detail.GetSortedSigsBytes()

		// execute the withdraw req onchain
		err = handler(execute, messageId, wdOnchain, sortedSigs, signers, powers)
		if err != nil {
			return fmt.Errorf("failed to execute liq withdraw (usr %s, nonce %d, chainId %d)", addr, nonce, chainId)
		} else {
			return nil
		}
	}
	return fmt.Errorf("PollAndExecuteWithdraw max retry reached for user %s, nonce %d, chainId %d", addr, nonce, chainId)
}

func (c *SgnClient) PollAndExecutePegRefundMint(burnId []byte, chainId uint64, execute types.RefundTxFunc, messageId []byte, handler types.ExecuteRefund) error {
	for try := 1; try <= types.MaxPollingRetries; try++ {
		log.Debugf("polling ClaimRefund status (try %d/%d): burnId %x, chainId %d", try, types.MaxPollingRetries, burnId, chainId)
		time.Sleep(types.PollingInterval)
		res, err := c.GetChainSigners(chainId)
		if err != nil {
			return fmt.Errorf("failed to query chain signers: %s", err.Error())
		}
		mintId, err := c.GetPegRefundClaimId(burnId)
		if err != nil {
			return fmt.Errorf("failed to query refund claim info for deposit (id %x): %s", burnId, err.Error())
		}
		mintInfo, err := c.GetPegMintInfo(mintId)
		if err != nil {
			if strings.Contains(err.Error(), pegbrtypes.ErrNoInfoFound.Error()) {
				log.Infof("peg mint info for burnId %x not found yet", burnId)
				continue
			}
			return fmt.Errorf("failed to query peg mint info: %s", err.Error())
		}
		pass, sigBytes := cbrtypes.ValidateSignatureQuorum(mintInfo.GetSignatures(), res.GetSortedSigners())
		if !pass {
			log.Infof("skip peg burn refund (burnId %x): not enough sigs yet", burnId)
			continue
		}
		signers, powers := res.GetAddrsPowers()
		err = handler(execute, messageId, mintInfo.MintProtoBytes, sigBytes, signers, powers)
		if err != nil {
			return fmt.Errorf("failed to execute peg burn (burnId %x) refund mint: %s", burnId, err.Error())
		} else {
			return nil
		}
	}
	return fmt.Errorf("PollAndExecutePegRefundMint max retry reached for burnId %x chainId %d", burnId, chainId)
}

func (c *SgnClient) PollAndExecutePegRefundWithdraw(depositId []byte, chainId uint64, execute types.RefundTxFunc, messageId []byte, handler types.ExecuteRefund) error {
	for try := 1; try <= types.MaxPollingRetries; try++ {
		log.Debugf("polling ClaimRefund status (try %d/%d): depositId %x, chainId %d", try, types.MaxPollingRetries, depositId, chainId)
		time.Sleep(types.PollingInterval)
		res, err := c.GetChainSigners(chainId)
		if err != nil {
			return fmt.Errorf("failed to query chain signers %s", err.Error())
		}
		withdrawId, err := c.GetPegRefundClaimId(depositId)
		if err != nil {
			return fmt.Errorf("failed to query refund claim info for deposit (id %x): %s", depositId, err.Error())
		}
		withdrawInfo, err := c.GetPegWithdrawInfo(withdrawId)
		if err != nil {
			if strings.Contains(err.Error(), pegbrtypes.ErrNoInfoFound.Error()) {
				log.Infof("peg withdraw info for depositId %x not found yet", depositId)
				continue
			}
			return fmt.Errorf("failed to query withdraw info %s", err.Error())
		}
		pass, sigBytes := cbrtypes.ValidateSignatureQuorum(withdrawInfo.GetSignatures(), res.GetSortedSigners())
		if !pass {
			log.Infof("skip peg deposit refund (depositId %X): not enough sigs yet", depositId)
			continue
		}
		signers, powers := res.GetAddrsPowers()
		err = handler(execute, messageId, withdrawInfo.WithdrawProtoBytes, sigBytes, signers, powers)
		if err != nil {
			return fmt.Errorf("failed to execute peg deposit (depositId %x) refund withdraw: %s", depositId, err.Error())
		} else {
			return nil
		}
	}
	return fmt.Errorf("PollAndExecutePegRefundWithdraw max retry reached for depositId %x chainId %d", depositId, chainId)
}

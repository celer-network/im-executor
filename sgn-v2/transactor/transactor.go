package transactor

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/sgn-v2/common"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

const (
	maxTxRetry         = 15
	maxTxQueryRetry    = 30
	txRetryDelay       = 1 * time.Second
	maxSignRetry       = 10
	signRetryDelay     = 100 * time.Millisecond
	maxWaitMinedRetry  = 5
	maxRawMsgBytesInTx = 500000

	appName = "sgn"
)

var errGasCode = fmt.Errorf("code %d", sdkerrors.ErrOutOfGas.ABCICode())
var errSeqCode = fmt.Errorf("code %d", sdkerrors.ErrWrongSequence.ABCICode())

type Transactor struct {
	TxFactory  clienttx.Factory
	CliCtx     client.Context
	Key        keyring.Info
	passphrase string
	msgQueue   MsgQueue
	sendLock   sync.Mutex

	// fields used by the WaitDone mechanism
	done     chan bool
	waitDone bool // whether there is currently a WaitDone() call in progress

	GrpcConn *grpc.ClientConn
}

func NewTransactor(
	homeDir, chainID, nodeURI, accAddr, passphrase string,
	legacyAmino *codec.LegacyAmino,
	cdc codec.Codec,
	interfaceRegistry codectypes.InterfaceRegistry,
) (*Transactor, error) {
	return NewTransactorWithGrpc(homeDir, chainID, nodeURI, accAddr, passphrase, "",
		legacyAmino, cdc, interfaceRegistry)
}

func NewTransactorWithGrpc(
	homeDir, chainID, nodeURI, accAddr, passphrase, grpcUrl string,
	legacyAmino *codec.LegacyAmino,
	cdc codec.Codec,
	interfaceRegistry codectypes.InterfaceRegistry,
) (*Transactor, error) {
	reader := strings.NewReader(passphrase + "\n")
	kb, err := keyring.New(appName,
		viper.GetString(common.FlagSgnKeyringBackend), homeDir, reader)
	if err != nil {
		return nil, err
	}

	addr, err := sdk.AccAddressFromBech32(accAddr)
	if err != nil {
		return nil, err
	}

	// may run into "resource temporarily unavailable" error if directly run it
	// retry when get this issue to avoid failure.
	var key keyring.Info
	for try := 0; try < maxSignRetry; try++ {
		key, err = kb.KeyByAddress(addr)
		if err == nil {
			break
		}
		if !strings.Contains(err.Error(), "resource temporarily unavailable") {
			log.Errorln("kb.GetByAddress error:", err)
			return nil, err
		}
		if try != maxSignRetry-1 {
			log.Debugln("retry kb.GetByAddress due to error:", err)
			time.Sleep(signRetryDelay)
		}
	}

	gasAdjustment := viper.GetFloat64(common.FlagSgnGasAdjustment)
	if gasAdjustment == 0 {
		gasAdjustment = common.DefaultSgnGasAdjustment
	}

	txConfig := tx.NewTxConfig((cdc).(*codec.ProtoCodec), tx.DefaultSignModes)
	cli, err := client.NewClientFromNode(nodeURI)
	if err != nil {
		log.Errorln("client.NewClientFromNode error:", err)
		return nil, err
	}
	cliCtx := client.Context{}.
		WithCodec(cdc).
		WithFromAddress(key.GetAddress()).
		WithFromName(key.GetName()).
		WithNodeURI(nodeURI).
		WithKeyring(kb).
		WithChainID(chainID).
		WithBroadcastMode(flags.BroadcastSync).
		WithTxConfig(txConfig).
		WithLegacyAmino(legacyAmino).
		WithAccountRetriever(authtypes.AccountRetriever{}).
		WithInterfaceRegistry(interfaceRegistry).
		WithClient(cli).
		WithHomeDir(homeDir)

	f := clienttx.Factory{}.
		WithKeybase(cliCtx.Keyring).
		WithTxConfig(cliCtx.TxConfig).
		WithAccountRetriever(cliCtx.AccountRetriever).
		WithAccountNumber(viper.GetUint64(flags.FlagAccountNumber)).
		WithSequence(viper.GetUint64(flags.FlagSequence)).
		WithGas(common.DefaultSgnGasLimit).
		WithGasAdjustment(gasAdjustment).
		WithChainID(chainID).
		WithMemo(viper.GetString(flags.FlagNote)).
		WithFees(viper.GetString(flags.FlagFees)).
		WithGasPrices(viper.GetString(flags.FlagGasPrices)).
		WithSignMode(signing.SignMode_SIGN_MODE_DIRECT).
		WithSimulateAndExecute(true)

	transactor := &Transactor{
		TxFactory:  f,
		CliCtx:     cliCtx,
		Key:        key,
		passphrase: passphrase,
		waitDone:   false,
		done:       make(chan bool),
	}

	if grpcUrl != "" {
		grpcConn, err := grpc.Dial(
			grpcUrl, grpc.WithInsecure(),
		)

		if err != nil {
			log.Errorln("grpc.Dial error:", err)
			return nil, err
		}

		transactor.GrpcConn = grpcConn
	}

	return transactor, nil
}

func NewCliTransactor(homeDir string, legacyAmino *codec.LegacyAmino, cdc codec.Codec, interfaceRegistry codectypes.InterfaceRegistry) (*Transactor, error) {
	return NewTransactor(
		homeDir,
		viper.GetString(common.FlagSgnChainId),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetString(common.FlagSgnValidatorAccount),
		viper.GetString(common.FlagSgnPassphrase),
		legacyAmino,
		cdc,
		interfaceRegistry,
	)
}

func NewTransactorWithCliCtx(cliCtx client.Context) (*Transactor, error) {
	return NewTransactor(
		cliCtx.HomeDir,
		viper.GetString(common.FlagSgnChainId),
		viper.GetString(common.FlagSgnNodeURI),
		viper.GetString(common.FlagSgnValidatorAccount),
		viper.GetString(common.FlagSgnPassphrase),
		cliCtx.LegacyAmino,
		cliCtx.Codec,
		cliCtx.InterfaceRegistry,
	)
}

// AddTxMsg add msg into a queue before actual broadcast
func (t *Transactor) AddTxMsg(msg sdk.Msg) {
	t.checkSigner([]sdk.Msg{msg})
	t.msgQueue.PushBack(msg)
}

// blocks until the msg queue is emptied
func (t *Transactor) WaitDone() {
	t.waitDone = true
	<-t.done
}

func (t *Transactor) sendTxMsgs(msgs []sdk.Msg, gas uint64) (*sdk.TxResponse, error) {
	var txResponseErr error
	for try := 0; try < maxTxRetry; try++ {
		txBytes, err := t.buildAndSignTx(msgs, gas)
		if err != nil {
			return nil, fmt.Errorf("buildAndSignTx err: %w", err)
		}
		log.Debugln("tx msg bytes size:", len(txBytes))
		txResponse, err := t.CliCtx.BroadcastTx(txBytes)
		if err != nil {
			return nil, fmt.Errorf("BroadcastTx err: %w", err)
		}

		if txResponse.Code == sdkerrors.SuccessABCICode {
			return txResponse, nil
		}

		txResponseErr = fmt.Errorf("BroadcastTx failed with code: %d, rawLog: %s, acct: %s",
			txResponse.Code, txResponse.RawLog, t.Key.GetAddress())
		if txResponse.Code == sdkerrors.ErrWrongSequence.ABCICode() && try < maxTxRetry-1 {
			log.Debugln(txResponseErr.Error(), "will retry")
			time.Sleep(txRetryDelay)
		} else {
			return txResponse, txResponseErr
		}
	}
	return nil, txResponseErr
}

// send a single msg so one fail won't affect others. this is only intended
// for initwithdraw/signagain request, if x/cbr return err, we return it immediately (wrapped in sendTxMsgs fmt.Errorf)
// if nil err, caller should query later. No waitmine
// note due to inherent async of estimategas and actual include in block, it's
// possible even this returns nil err, x/cbr still fails
// lock to ensure req are serialized even gateway handle concurrent initwithdraw from clients
func (t *Transactor) LockSendTx(msg sdk.Msg) (*sdk.TxResponse, error) {
	t.checkSigner([]sdk.Msg{msg})
	t.sendLock.Lock()
	defer t.sendLock.Unlock()
	return t.sendTxMsgs([]sdk.Msg{msg}, 0) // 0 gas so estimate will be called
}

func (t *Transactor) LockSendTxs(msgs []sdk.Msg) (*sdk.TxResponse, error) {
	t.checkSigner(msgs)
	t.sendLock.Lock()
	defer t.sendLock.Unlock()
	return t.sendTxMsgs(msgs, 0) // 0 gas so estimate will be called
}

func (t *Transactor) buildAndSignTx(msgs []sdk.Msg, gas uint64) ([]byte, error) {
	txf := t.TxFactory
	txf, err := prepareFactory(t.CliCtx, txf)
	if err != nil {
		return nil, err
	}

	if gas != 0 {
		txf = txf.WithGas(gas)
	} else if txf.SimulateAndExecute() || t.CliCtx.Simulate {
		for try := 0; try < maxTxRetry; try++ {
			_, adjusted, err := clienttx.CalculateGas(t.CliCtx, txf, msgs...)
			if err != nil {
				if strings.Contains(err.Error(), "account sequence mismatch") && try < maxTxRetry-1 {
					log.Debugln(err, "increment seq and retry")
					txf = txf.WithSequence(txf.Sequence() + 1)
					continue
				}
				return nil, fmt.Errorf("CalculateGas err: %w", err)
			}
			txf = txf.WithGas(adjusted)
			break
		}
	}

	tx, err := clienttx.BuildUnsignedTx(txf, msgs...)
	if err != nil {
		return nil, err
	}

	tx.SetFeeGranter(t.CliCtx.GetFeeGranterAddress())

	err = clienttx.Sign(txf, t.CliCtx.GetFromName(), tx, true)
	if err != nil {
		return nil, err
	}

	return t.CliCtx.TxConfig.TxEncoder()(tx.GetTx())
}

func (t *Transactor) waitMined(txHash string) (*sdk.TxResponse, error) {
	var err error
	mined := false
	var txResponse *sdk.TxResponse
	for try := 0; try < maxTxQueryRetry; try++ {
		time.Sleep(txRetryDelay)
		if txResponse, err = tx.QueryTx(t.CliCtx, txHash); err == nil {
			mined = true
			break
		}
	}
	if !mined {
		return txResponse, fmt.Errorf("tx not mined, err: %w", err)
	} else if txResponse.Code != sdkerrors.SuccessABCICode {
		if txResponse.Code == sdkerrors.ErrOutOfGas.ABCICode() { // out of gas
			return txResponse, fmt.Errorf("tx failed with %w, %s", errGasCode, txResponse.RawLog)
		} else if txResponse.Code == sdkerrors.ErrWrongSequence.ABCICode() {
			return txResponse, fmt.Errorf("tx failed with %w, %s", errSeqCode, txResponse.RawLog)
		} else {
			return txResponse, fmt.Errorf("tx failed with code %d, %s", txResponse.Code, txResponse.RawLog)
		}
	}
	return txResponse, nil
}

// prepareFactory ensures the account defined by ctx.GetFromAddress() exists and
// if the account number and/or the account sequence number are zero (not set),
// they will be queried for and set on the provided Factory. A new Factory with
// the updated fields will be returned.
func prepareFactory(clientCtx client.Context, txf clienttx.Factory) (clienttx.Factory, error) {
	from := clientCtx.GetFromAddress()

	if err := txf.AccountRetriever().EnsureExists(clientCtx, from); err != nil {
		return txf, err
	}

	initNum, initSeq := txf.AccountNumber(), txf.Sequence()
	if initNum == 0 || initSeq == 0 {
		num, seq, err := txf.AccountRetriever().GetAccountNumberSequence(clientCtx, from)
		if err != nil {
			return txf, err
		}

		if initNum == 0 {
			txf = txf.WithAccountNumber(num)
		}

		if initSeq == 0 {
			txf = txf.WithSequence(seq)
		}
	}

	return txf, nil
}

func (t *Transactor) checkSigner(msgs []sdk.Msg) {
	for _, msg := range msgs {
		if t.Key.GetAddress().String() != msg.GetSigners()[0].String() {
			log.Fatal("tx msg signer is not the transactor")
		}
	}
}

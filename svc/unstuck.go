package svc

import (
	"context"
	"fmt"
	"math/big"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/goutils/log"
	"github.com/celer-network/im-executor/chains"
	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/dal/models"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	"github.com/celer-network/im-executor/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type GasPriceOptions struct {
	// legacy
	GasPrice *big.Int
	// eip1559
	GasTipCap *big.Int
	GasFeeCap *big.Int
}

func (s *ExecutorService) UnstuckTxAndRevertStatusBySrcTx(
	srcTxHash string, accountID string, gasOpts *GasPriceOptions, wait bool) (eth.Hash, error) {
	records, err := s.db.GetExecutionRecords(&dal.ExecutionRecordsQuery{SrcTx: srcTxHash})
	if err != nil {
		return eth.ZeroHash, err
	}
	if len(records) == 0 {
		return eth.ZeroHash, nil
	}
	replacementTx, err := s.unstuckTxAndRevertStatus(records[0], accountID, gasOpts, wait)
	if err != nil {
		return eth.ZeroHash, err
	}
	for i := 1; i < len(records); i++ {
		err := s.db.RevertStatus(records[i].ID.Bytes(), records[i].GetLastExecutableStatus())
		if err != nil {
			log.Errorf("failed to revert status: %s", err)
		}
	}
	return replacementTx, nil
}

func (s *ExecutorService) UnstuckTxAndRevertStatus(
	query *dal.ExecutionRecordQuery, accountID string, gasOpts *GasPriceOptions, wait bool) (eth.Hash, error) {
	r, err := s.QueryExecutionRecord(query)
	if err != nil {
		return eth.ZeroHash, err
	}
	return s.unstuckTxAndRevertStatus(r, accountID, gasOpts, wait)
}

func (s *ExecutorService) unstuckTxAndRevertStatus(
	r *models.ExecutionRecord, accountID string, gasOpts *GasPriceOptions, wait bool) (eth.Hash, error) {
	if r.ExecutionStatus != types.ExecutionStatus_Executing {
		return eth.ZeroHash, fmt.Errorf("cannot send unstuck tx on execution record status %s", r.ExecutionStatus)
	}
	chain, ok := s.chains.GetChain(r.ChainID)
	if !ok {
		return eth.ZeroHash, fmt.Errorf("failed to get chain %d", r.ChainID)
	}
	acc, ok := s.accounts.AccountByID(accountID)
	if !ok {
		return eth.ZeroHash, fmt.Errorf("account with id %s not found", accountID)
	}
	signer, ok := acc.Signer(chain.ChainID)
	if !ok {
		return eth.ZeroHash, fmt.Errorf("account with id %s does not have a signer", accountID)
	}
	nonce, err := chain.EthClient.NonceAt(context.Background(), acc.Address, nil)
	if err != nil {
		return eth.ZeroHash, fmt.Errorf("failed to get nonce at %s: %s", acc.Address, err)
	}
	tx, err := buildEmptyTxAtNonce(chain, acc.Address, nonce, gasOpts)
	if err != nil {
		return eth.ZeroHash, fmt.Errorf("failed to build tx: %s", err)
	}
	log.Infof("sending an empty tx at nonce %d...", nonce)
	txBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return eth.ZeroHash, fmt.Errorf("failed to encode tx: %s", err)
	}
	rawSignedTx, err := signer.SignEthTransaction(txBytes)
	if err != nil {
		return eth.ZeroHash, fmt.Errorf("failed to sign tx: %s", err)
	}
	err = rlp.DecodeBytes(rawSignedTx, tx)
	if err != nil {
		return eth.ZeroHash, fmt.Errorf("failed to unmarshal tx: %s", err)
	}
	err = chain.EthClient.SendTransaction(context.Background(), tx)
	if err != nil {
		return eth.ZeroHash, fmt.Errorf("failed to send tx: %s", err)
	}
	log.Infof("tx sent %s", tx.Hash())
	if wait {
		return tx.Hash(), s.waitThenRevertStatus(chain, tx, r)
	} else {
		go s.waitThenRevertStatus(chain, tx, r)
	}
	return tx.Hash(), nil
}

func (s *ExecutorService) waitThenRevertStatus(chain *chains.Chain, tx *ethtypes.Transaction, r *models.ExecutionRecord) error {
	blkDelay := chain.BlkDelay
	if blkDelay == 0 {
		blkDelay = 5
	}
	log.Infof("waiting for %d blocks before reverting record status...", blkDelay)
	ethutils.WaitMined(context.Background(), chain.EthClient, tx, ethutils.WithBlockDelay(blkDelay))
	log.Infof("done")
	err := s.db.RevertStatus(r.ID.Bytes(), r.GetLastExecutableStatus())
	if err != nil {
		return fmt.Errorf("failed to revert status: %s", err)
	}
	return nil
}

func buildEmptyTxAtNonce(chain *chains.Chain, to eth.Addr, nonce uint64, o *GasPriceOptions) (*ethtypes.Transaction, error) {
	ec := chain.EthClient
	bgCtx := context.Background()
	head, err := ec.HeaderByNumber(bgCtx, nil)
	if err != nil {
		return nil, fmt.Errorf("HeaderByNumber err %w", err)
	}
	use1559 := o.GasPrice == nil && head.BaseFee != nil
	suggestedPrice, err := ec.SuggestGasPrice(bgCtx)
	if err != nil {
		return nil, fmt.Errorf("SuggestGasPrice err %w", err)
	}
	if use1559 {
		if o.GasTipCap == nil {
			o.GasTipCap = new(big.Int).Sub(suggestedPrice, head.BaseFee)
			log.Infof("using suggested gas tip cap %s", o.GasTipCap)
		}
		if o.GasFeeCap == nil {
			o.GasFeeCap = new(big.Int).Add(suggestedPrice, head.BaseFee)
			log.Infof("using suggested gas fee cap %s", o.GasFeeCap)
		}
		return ethtypes.NewTx(&ethtypes.DynamicFeeTx{
			Nonce:     nonce,
			To:        &to,
			Gas:       21000,
			GasTipCap: o.GasFeeCap,
			GasFeeCap: o.GasFeeCap,
		}), nil
	} else {
		if o.GasPrice == nil {
			o.GasPrice = suggestedPrice
		}
		return ethtypes.NewTx(&ethtypes.LegacyTx{
			Nonce:    nonce,
			To:       &to,
			Gas:      21000,
			GasPrice: o.GasPrice,
		}), nil
	}
}

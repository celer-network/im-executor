package executor

import (
	"fmt"
	"math/big"

	ethutils "github.com/celer-network/goutils/eth"
	"github.com/celer-network/im-executor/accounts"
	"github.com/celer-network/im-executor/chains"
	"github.com/celer-network/im-executor/contracts"
	"github.com/celer-network/im-executor/dal"
	"github.com/celer-network/im-executor/dal/models"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	"github.com/celer-network/im-executor/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
)

// Execution is a context object that is assembled before the execution of a message so that later steps of the execution
// don't need to acquire the relavant dependencies over and over.
type Execution struct {
	Chain          *chains.Chain
	Transactor     *ethutils.Transactor
	Receiver       *contracts.ReceiverContract
	Record         *models.ExecutionRecord
	DelayedMessage *dal.DelayedMessage
	// Gas limit override. Useful for force sending out a tx and debugging in Tenderly
	GasLimit uint64
}

func newExecution(acc *accounts.Account, chain *chains.Chain, record *models.ExecutionRecord, gasLimit uint64) (*Execution, error) {
	txr, ok := acc.Transactor(chain.ChainID)
	if !ok {
		return nil, fmt.Errorf("cannot create Execution: transactor not registered for chainId %d", chain.ChainID)
	}
	receiver := &contracts.ContractAddress{
		ChainId: chain.ChainID,
		Address: record.ExecutionContext.Message.Receiver,
	}
	recvContract, ok := acc.ReceiverContract(receiver)
	if !ok {
		return nil, fmt.Errorf("cannot create Execution: receiver contract not registered for %s", receiver)
	}
	return &Execution{
		Chain:      chain,
		Transactor: txr,
		Receiver:   recvContract,
		Record:     record,
		GasLimit:   gasLimit,
	}, nil
}

func (e *Execution) GetTxOptions() []ethutils.TxOption {
	options := []ethutils.TxOption{}
	if e.GasLimit > 0 {
		options = append(options, ethutils.WithGasLimit(e.GasLimit))
	}
	return options
}

func (x *Execution) ExecuteMessage(opts *bind.TransactOpts, msg []byte, route types.MessageRoute, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) (*gethtypes.Transaction, error) {
	if route.IsBytesAddr {
		routeInfo := eth.MsgDataTypesRouteInfo2{
			Sender:     eth.Hex2Bytes(route.Sender),
			Receiver:   eth.Hex2Addr(route.Receiver),
			SrcChainId: route.SrcChainId,
			SrcTxHash:  eth.Hex2Hash(route.SrcTxHash),
		}
		return x.Chain.MessageBus.ExecuteMessage0(opts, msg, routeInfo, sortedSigs, signers, powers)
	}
	routeInfo := eth.MsgDataTypesRouteInfo{
		Sender:     eth.Hex2Addr(route.Sender),
		Receiver:   eth.Hex2Addr(route.Receiver),
		SrcChainId: route.SrcChainId,
		SrcTxHash:  eth.Hex2Hash(route.SrcTxHash),
	}
	return x.Chain.MessageBus.ExecuteMessage(opts, msg, routeInfo, sortedSigs, signers, powers)
}

func (x *Execution) ExecuteDelayedMessage(opts *bind.TransactOpts, adapterAddr, srcContract eth.Addr, srcChainId uint64, dstContract eth.Addr, calldata []byte, nonce uint32) (*gethtypes.Transaction, error) {
	adapter, ok := x.Chain.Contracts.MsgRecvAdapters[adapterAddr]
	if !ok {
		return nil, fmt.Errorf("%x not configurred in adapters of chain %d", adapterAddr, x.Chain.ChainID)
	}
	return adapter.ExecuteDelayedMessage(opts, srcContract, srcChainId, dstContract, calldata, nonce)
}

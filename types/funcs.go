package types

import (
	"math/big"

	"github.com/celer-network/im-executor/sgn-v2/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type MessageRoute struct {
	Sender      string
	Receiver    string
	SrcChainId  uint64
	SrcTxHash   string
	IsBytesAddr bool
}

type ExecuteRefund func(execute RefundTxFunc, messageId []byte, wdOnchain []byte, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) error

type RefundTxFunc func(
	opts *bind.TransactOpts,
	wdOnchain []byte,
	sortedSigs [][]byte,
	signers []eth.Addr,
	powers []*big.Int) (*ethtypes.Transaction, error)

type ExecuteMsg func(execute ExecuteMsgTxFunc, messageId []byte, msg []byte, route MessageRoute, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) error

type ExecuteMsgTxFunc func(opts *bind.TransactOpts, msg []byte, route MessageRoute, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) (*ethtypes.Transaction, error)

type ExecuteMsgWithXfer func(execute ExecuteMsgWithXferTxFunc, messageId []byte, msg []byte, xfer eth.MsgDataTypesTransferInfo, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) error

type ExecuteMsgWithXferTxFunc func(opts *bind.TransactOpts, msg []byte, xfer eth.MsgDataTypesTransferInfo, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) (*ethtypes.Transaction, error)

type ExecuteMsgWithXferRefund func(execute ExecuteMsgWithXferRefundTxFunc, messageId []byte, msg []byte, xfer eth.MsgDataTypesTransferInfo, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) error

type ExecuteMsgWithXferRefundTxFunc func(opts *bind.TransactOpts, msg []byte, xfer eth.MsgDataTypesTransferInfo, sortedSigs [][]byte, signers []eth.Addr, powers []*big.Int) (*ethtypes.Transaction, error)

type Allowed func(messageId []byte, srcChainId uint64, senderStr string)

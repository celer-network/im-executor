package types

import (
	"fmt"
	"math/big"

	"github.com/celer-network/im-executor/sgn-v2/eth"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func (m *XferRelay) GetSortedSigsBytes() [][]byte {
	if m != nil {
		sigs := make([][]byte, len(m.SortedSigs))
		for i := range m.SortedSigs {
			sigs[i] = m.SortedSigs[i].Sig
		}
		return sigs
	}
	return nil
}

func (m *XferRelay) SignersStr() string {
	var signers string
	for _, s := range m.SortedSigs {
		signers += fmt.Sprintf("%x ", s.Addr)
	}
	return fmt.Sprintf("signers: < %s>", signers)
}

func (m *WithdrawDetail) GetSortedSigsBytes() [][]byte {
	if m != nil {
		sigs := make([][]byte, len(m.SortedSigs))
		for i := range m.SortedSigs {
			sigs[i] = m.SortedSigs[i].Sig
		}
		return sigs
	}
	return nil
}

func (m *WithdrawDetail) SignersStr() string {
	var signers string
	for _, s := range m.SortedSigs {
		signers += fmt.Sprintf("%x ", s.Addr)
	}
	return fmt.Sprintf("signers: < %s>", signers)
}

func (m *ChainSigners) SignersStr() string {
	var signers string
	for _, s := range m.SortedSigs {
		signers += fmt.Sprintf("%x ", s.Addr)
	}
	return fmt.Sprintf("signers: < %s>", signers)
}

func (r *RelayOnChain) String() string {
	if r == nil {
		return ""
	}
	return fmt.Sprintf("sender %x, receiver %x, token %x, amount %s, src_chain_id %d, dst_chain_id %d, src_xfer_id %x",
		r.Sender, r.Receiver, r.Token, big.NewInt(0).SetBytes(r.Amount), r.SrcChainId, r.DstChainId, r.SrcTransferId)
}

func (w *WithdrawOnchain) String() string {
	if w == nil {
		return ""
	}
	return fmt.Sprintf("chainid %d, seqnum %d, receiver %x, token %x, amount %d, refid %x",
		w.Chainid, w.Seqnum, w.Receiver, w.Token, big.NewInt(0).SetBytes(w.Amount), w.Refid)
}

// todo for NON-EVM-CBR
func EncodeRelayOnChainToSign(chainId uint64, contractAddr eth.Addr, relayBytes []byte) []byte {
	domain := solsha3.SoliditySHA3(
		[]string{"uint256", "address", "string"},
		[]interface{}{new(big.Int).SetUint64(chainId), contractAddr, "Relay"},
	)
	return append(domain, relayBytes...)
}

// todo for NON-EVM-CBR
func EncodeWithdrawOnchainToSign(chainId uint64, contractAddr eth.Addr, withdrawBytes []byte) []byte {
	domain := solsha3.SoliditySHA3(
		[]string{"uint256", "address", "string"},
		[]interface{}{new(big.Int).SetUint64(chainId), contractAddr, "WithdrawMsg"},
	)
	return append(domain, withdrawBytes...)
}

package types

import (
	"fmt"
	"strings"

	commontypes "github.com/celer-network/im-executor/sgn-v2/common/types"
	"github.com/celer-network/im-executor/sgn-v2/eth"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "cbridge"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_cbridge"

	// DefaultParamspace default name for parameter store
	DefaultParamspace = ModuleName

	CBridgeStakeDenomPrefix = "CB-"
	CBridgeFeeDenomPrefix   = "CBF-"
)

var (
	ChainSignersKey  = []byte("signers-chain")
	LatestSignersKey = []byte("signers-latest")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetChainSignersKey(chid uint64) []byte {
	return append(ChainSignersKey, sdk.Uint64ToBigEndian(chid)...)
}

/* states owned by cbridge module
1. liquidity map, lm-chid-token-lp -> amount big.Int.Bytes
2. processed add liquidity event, evliqadd-chid-seq -> true, to avoid process same event again
3. send event, evsend-%x transferid, module has seen this event, value is enum status
4. no longer save evrelay # relay event, evrelay-%x relay transferid -> srcTransferid
5. xfer relay: xferRelay-%x, src transfer id, relay msg and sigs
6. no longer need withdrawSeq
7. withdraw detail, wdDetail-%x-%d user addr and reqid, value is onchain msg and sigs
8. xfer refund, xferRefund-%x src xfer id -> withdrawonchain, only for failed xfer. first set when apply send, but no reqid, later when user InitWithdraw, set reqid in it
9. lp fee, lpfee-chid-token-lp -> fee big.Int bytes on this (chain,token)
10. sgn fee, sgnfee-chid-token -> big.Int bytes
11. liquidity sum of liqsum-chid-token, always equal sum of all lm-chid-token-xxx, we keep sum to avoid iter over all lps
*/

// key for liquidity map, chainid-tokenaddr-lpaddr
// value is big.Int.Bytes()
func LiqMapKey(chid uint64, token, lp *commontypes.Address) []byte {
	return []byte(fmt.Sprintf("lm-%d-%s-%s", chid, token.String(), lp.String()))
}

// value is big.Int bytes of sum over all lm-chid-token-xxx
func LiqSumKey(chid uint64, token *commontypes.Address) []byte {
	return []byte(fmt.Sprintf("liqsum-%d-%s", chid, token.String()))
}

func LiqCapKey(chid uint64, token *commontypes.Address) []byte {
	return []byte(fmt.Sprintf("liqcap-%d-%s", chid, token.String()))
}

func GetLpAddrFromLiqMapKey(key []byte) (*commontypes.Address, error) {
	keystrs := strings.Split(string(key), "-")
	if len(keystrs) != 4 {
		return commontypes.ZeroAddr(), fmt.Errorf("invaid key")
	}
	return new(commontypes.Address).SetString(keystrs[3]), nil
}

// value is 0x01 to indicate has applied event
func EvLiqAddKey(chid, seq uint64) []byte {
	return []byte(fmt.Sprintf("evliqadd-%d-%d", chid, seq))
}

// tid is user's transfer if. value is enum xfer status
func EvSendKey(tid eth.Hash) []byte {
	return []byte(fmt.Sprintf("evsend-%x", tid))
}

func LPOriginKey(lp *commontypes.Address) []byte {
	return []byte(fmt.Sprintf("lporigin-%s", lp.String()))
}

// serialized relay msg and sigs, add sig when receive msg
func XferRelayKey(tid eth.Hash) []byte {
	return []byte(fmt.Sprintf("xferRelay-%x", tid))
}

func XferRefundKey(tid eth.Hash) []byte {
	return []byte(fmt.Sprintf("xferRefund-%x", tid))
}

func WdDetailKey(usraddr *commontypes.Address, reqid uint64) []byte {
	return []byte(fmt.Sprintf("wdDetail-%s-%d", usraddr.String(), reqid))
}

// for chid, token, how much fee this lp has earned
// NOTE: Cumulative amount, only increasing
func LpFeeKey(chid uint64, token, lp *commontypes.Address) []byte {
	return []byte(fmt.Sprintf("lpfee-%d-%s-%s", chid, token.String(), lp.String()))
}

// Tracks the **total** fee allocated for SGN delegators
// NOTE: Cumulative amount, only increasing
func SgnFeeKey(chid uint64, token *commontypes.Address) []byte {
	return []byte(fmt.Sprintf("sgnfee-%d-%s", chid, token.String()))
}

/* ================ config kv, all governable
1. fee percentage goes to cbridge lp, eg. 80 means 80% goes to lp
2. chid-tokenAddr -> asset symbol string eg. "USDT", all uppercase
3. symbol-chid -> ChainAsset, note proto has dup info symbol and chain_id
4. chid1-chid2 -> ChainPair. keys are sorted so chid1 < chid2
5. pick lp size, how many LPs on first select. value is big.Int bytes
6. chid -> gas price big.Int.Bytes.
7. chid -> GasTokenSymbol string.
8. symbol -> uint32(USD price * 1e(4+ExtraPower10) if ExtraPower10 isn't set, just 1e4
9. chid -> GasCostParam
10. chid -> GasCost
11. symbol-chid1-chid2 -> ChainPair. per (chainpair, token) info override
12. symbol -> uint32 ExtraPower10, only exist if asset USD float < $0.0001
*/

var (
	CfgKeyFeePerc        = []byte("cfg-feeperc")
	CfgKeyFeeSplitGlobal = []byte("cfg-feesplit")
	CfgKeyPickLpSize     = []byte("cfg-lpsize")
	CfgKeyMaxGainPerc    = []byte("cfg-maxgainperc")
)

func CfgKeyChain2Sym(chid uint64, addr *commontypes.Address) []byte {
	return []byte(fmt.Sprintf("cfg-ch2sym-%d-%s", chid, addr.String()))
}

func CfgKeySym2Info(sym string, chid uint64) []byte {
	return []byte(fmt.Sprintf("cfg-sym2info-%s-%d", sym, chid))
}

func CfgKeyChainPair(chid1, chid2 uint64) []byte {
	if chid1 > chid2 {
		panic(fmt.Sprintf("chid1 %d > chid2 %d", chid2, chid2))
	}
	return []byte(fmt.Sprintf("cfg-chpair-%d-%d", chid1, chid2))
}

func CfgKeyChain2GasPrice(chid uint64) []byte {
	return []byte(fmt.Sprintf("cfg-ch2gasprice-%d", chid))
}

func CfgKeyChain2GasTokenSymbol(chid uint64) []byte {
	return []byte(fmt.Sprintf("cfg-ch2gastokensymbol-%d", chid))
}

func CfgKeySymbol2UsdPrice(sym string) []byte {
	return []byte(fmt.Sprintf("cfg-symbol2usdprice-%s", sym))
}

func CfgKeySymbol2ExtraPower10(sym string) []byte {
	return []byte(fmt.Sprintf("cfg-symbol2extrapower10-%s", sym))
}

func GetSymbolFromStakeToken(token string) string {
	return strings.Replace(token, CBridgeStakeDenomPrefix, "", 1)
}

// store params used to calculate relay gas cost when genesis
func CfgKeyChain2RelayGasCostParam(chid uint64) []byte {
	return []byte(fmt.Sprintf("cfg-ch2relaygascostparam-%d", chid))
}

// store estimate relay gas cost. only updated when monitored CbrEventSignersUpdated
func CfgKeyChain2EstimateRelayGasCost(chid uint64) []byte {
	return []byte(fmt.Sprintf("cfg-ch2relaygascost-%d", chid))
}

func CfgKeyCbrContract(chid uint64) []byte {
	return []byte(fmt.Sprintf("cfg-cbrcontract-%d", chid))
}

// chid1 must be smaller than chid2
func CfgKeyChainPairAssetOverride(sym string, chid1, chid2 uint64) []byte {
	return []byte(fmt.Sprintf("cfg-override-%s-%d-%d", sym, chid1, chid2))
}

func CfgKeyFeeSplitPoolOverride(sym string, dstChid uint64) []byte {
	return []byte(fmt.Sprintf("cfg-feesplit-%s-%d", sym, dstChid))
}

func CfgKeyFeeSplitChainOverride(dstChid uint64) []byte {
	return []byte(fmt.Sprintf("cfg-feesplit-%d", dstChid))
}

func BalancerKey(balancer *commontypes.Address) []byte {
	return []byte(fmt.Sprintf("balancer-%s", balancer.String()))
}

func BlacklistKey(account *commontypes.Address) []byte {
	return []byte(fmt.Sprintf("blacklist-%s", account.String()))
}

func ConfiscationKey(chid uint64, token *commontypes.Address) []byte {
	return []byte(fmt.Sprintf("confiscation-%d-%s", chid, token.String()))
}

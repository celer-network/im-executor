package eth

import (
	"encoding/hex"
	"fmt"
	"math/big"

	ec "github.com/ethereum/go-ethereum/common"
)

var (
	// ZeroAddr is all 0s
	ZeroAddr Addr
	// ZeroAddrHex is string of 20 0s
	ZeroAddrHex = Addr2Hex(ZeroAddr)
	// ZeroBigInt is big.NewInt(0)
	ZeroBigInt = big.NewInt(0)
	// ZeroHash is all 0s
	ZeroHash Hash
)

// Hash is the type for ethereum hash type
type Hash = ec.Hash

// Addr is alias to geth common.Address
type Addr = ec.Address

// ========== Hex/Bytes ==========

// Hex2Bytes supports hex string with or without 0x prefix
// Calls hex.DecodeString directly and ignore err
// similar to ec.FromHex but better
func Hex2Bytes(s string) (b []byte) {
	if len(s) >= 2 && s[0] == '0' && (s[1] == 'x' || s[1] == 'X') {
		s = s[2:]
	}
	// hex.DecodeString expects an even-length string
	if len(s)%2 == 1 {
		s = "0" + s
	}
	b, _ = hex.DecodeString(s)
	return b
}

// Bytes2Hex returns hex string without 0x prefix
func Bytes2Hex(b []byte) string {
	return hex.EncodeToString(b)
}

// ========== Address ==========

// Hex2Addr accepts hex string with or without 0x prefix and return Addr
func Hex2Addr(s string) Addr {
	return ec.HexToAddress(s)
}

// Addr2Hex returns hex without 0x
func Addr2Hex(a Addr) string {
	return Bytes2Hex(a[:])
}

// Bytes2Addr returns Address from b
// Addr.Bytes() does the reverse
func Bytes2Addr(b []byte) Addr {
	return ec.BytesToAddress(b)
}

// Bytes2AddrHex returns hex without 0x
func Bytes2AddrHex(b []byte) string {
	return Addr2Hex(Bytes2Addr(b))
}

// FormatAddrHex formats a string into standard Addr string
func FormatAddrHex(s string) string {
	return Addr2Hex(Hex2Addr(s))
}

// ========== Hash ==========

// Hex2Hash accepts hex string with or without 0x prefix and return Hash
func Hex2Hash(s string) Hash {
	return ec.HexToHash(s)
}

// Bytes2Hash converts bytes to Hash
func Bytes2Hash(b []byte) Hash {
	return ec.BytesToHash(b)
}

// if in is 20 bytes, return directly. otherwise return a new []byte w/ len 20, left pad 0x00..[in]
// if len(in)>20, panic
func Pad20Bytes(in []byte) []byte {
	origLen := len(in)
	if origLen == 20 {
		return in
	}
	if origLen > 20 {
		panic(fmt.Sprintf("%x len %d > 20", in, origLen))
	}
	ret := make([]byte, 20)
	copy(ret[20-origLen:], in)
	return ret
}

func Pad32Bytes(in []byte) []byte {
	origLen := len(in)
	if origLen == 32 {
		return in
	}
	if origLen > 32 {
		panic(fmt.Sprintf("%x len %d > 32", in, origLen))
	}
	ret := make([]byte, 32)
	copy(ret[32-origLen:], in)
	return ret
}

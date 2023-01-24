package eth

import "math/big"

func SignerBytes(addrs []Addr, powers []*big.Int) []byte {
	var packed []byte
	for _, addr := range addrs {
		packed = append(packed, Pad32Bytes(addr.Bytes())...)
	}
	for _, power := range powers {
		packed = append(packed, Pad32Bytes(power.Bytes())...)
	}
	return packed
}

// ToPadBytes return big-endian/network order bytes, left padded to specific length
// if v is uint32: 4 bytes, int64/uint64: 8 bytes, *big.Int: 32 bytes or rlen bytes if set
// return nil if type not supported
func ToPadBytes(v interface{}, rlen ...int) []byte {
	var orig []byte
	var retlen int
	switch k := v.(type) {
	case uint8:
		retlen = 1
		orig = big.NewInt(int64(k)).Bytes()
	case uint32:
		retlen = 4
		orig = big.NewInt(int64(k)).Bytes()
	case int64:
		retlen = 8
		orig = big.NewInt(k).Bytes()
	case uint64:
		retlen = 8
		orig = new(big.Int).SetUint64(k).Bytes()
	case *big.Int:
		if len(rlen) == 1 {
			retlen = rlen[0]
		} else {
			retlen = 32
		}
		orig = k.Bytes()
	default:
		return nil
	}
	ret := make([]byte, retlen)
	copy(ret[retlen-len(orig):], orig)
	return ret
}

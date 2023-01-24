package types

import (
	"strings"

	"github.com/celer-network/im-executor/sgn-v2/eth"
)

const (
	Address_Length_20  = uint16(20)
	Address_Length_8   = uint16(8)
	Address_Max_Length = uint16(1 << 8)
	Address_Length_32  = uint16(32)
)

var AddressTypeLength = map[Address_Type]uint16{
	Address_Bytes20: Address_Length_20,
	Address_Bytes8:  Address_Length_8,
	Address_Hex:     Address_Max_Length,
	Address_String:  Address_Max_Length,
	Address_Bytes32: Address_Length_32,
}

func ZeroAddr() *Address {
	return &Address{
		Data: []byte{0x0},
		Type: Address_Hex,
	}
}

func (a *Address) IsValid() bool {
	// nil address
	if a == nil || a.Data == nil || len(a.Data) == 0 {
		return false
	}
	// zero address
	if a.Type == Address_Hex && len(a.Data) == 1 && a.Data[0] == byte(0x0) {
		return false
	}
	return true
}

// if _data is longer than Address_Max_Length, only the first Address_Max_Length bytes of _data are left.
func (a *Address) setDataAndType(_data []byte, _type Address_Type) *Address {
	if (_type == Address_Hex || _type == Address_String) && len(_data) > int(Address_Max_Length) {
		a.Data = make([]byte, Address_Max_Length)
	} else {
		a.Data = make([]byte, len(_data))
	}
	copy(a.Data, _data)
	a.Type = _type
	return a
}

// SetDefaultAddress sets data to hexadecimal Address
func (a *Address) SetDefaultAddress(data []byte) *Address {
	return a.setDataAndType(data, Address_Hex)
}

func (a *Address) SetBytes20Address(addr eth.Addr) *Address {
	return a.setDataAndType(addr.Bytes(), Address_Bytes20)
}

func (a *Address) SetBytes8Address(flowAddr [8]byte) *Address {
	return a.setDataAndType(flowAddr[:], Address_Bytes8)
}

func (a *Address) SetStrAddress(addrStr string) *Address {
	return a.setDataAndType([]byte(addrStr), Address_String)
}

func (a *Address) SetBytes32Address(aptosAddr [32]byte) *Address {
	return a.setDataAndType(aptosAddr[:], Address_Bytes32)
}

func (a *Address) SetString(addrStr string) *Address {
	if addrStr == "" {
		return ZeroAddr()
	} else if IsHexAddress(addrStr) {
		bs := eth.Hex2Bytes(addrStr)
		return a.SetBytes(bs)
	} else {
		// now, just for flow token. example, A.1654653399040a61.FlowToken.Vault
		// also for aptos. example 0x12345678::test_mint_coin::TestMintCoin
		return a.SetStrAddress(addrStr)
	}
}

// SetBytes return an Address defined by _bytes with auto length match. Address_String is not applicable.
// It is very dangerous to apply SetBytes with an unknown byte slice, especially when there are string-like addresses.
func (a Address) SetBytes(_bytes []byte) *Address {
	switch uint16(len(_bytes)) {
	case Address_Length_20:
		return a.setDataAndType(_bytes, Address_Bytes20)
	case Address_Length_8:
		return a.setDataAndType(_bytes, Address_Bytes8)
	case Address_Length_32:
		return a.setDataAndType(_bytes, Address_Bytes32)
	default:
		return a.setDataAndType(_bytes, Address_Hex)
	}
}

// SetBytesByChainId return an Address defined by _bytes and _chainId.
// Some chain has hex type account address and string type token address. While only one type is matched by a chainId,
// we recommend only use SetBytesByChainId to new an account address.
func (a *Address) SetBytesByChainId(_bytes []byte, _chainId uint64) *Address {
	return a.SetBytesByType(_bytes, GetAddrTypeByChainId(_chainId))
}

// SetStringByChainId is a helper func based on SetBytesByChainId.
func (a *Address) SetStringByChainId(_addrStr string, _chainId uint64) *Address {
	var _bytes []byte
	if IsHexAddress(_addrStr) {
		_bytes = eth.Hex2Bytes(_addrStr)
	} else {
		_bytes = []byte(_addrStr)
	}
	return a.SetBytesByChainId(_bytes, _chainId)
}

// SetBytesByType return an Address defined by _bytes and _type.
func (a *Address) SetBytesByType(_bytes []byte, _type Address_Type) *Address {
	if CheckValidity(_bytes, _type) {
		return a.setDataAndType(_bytes, _type)
	} else {
		return a.Copy(ZeroAddr())
	}
}

func (a *Address) Copy(expected *Address) *Address {
	return a.setDataAndType(expected.GetData(), expected.GetType())
}

func (a *Address) String() string {
	switch a.Type {
	case Address_Hex, Address_Bytes20, Address_Bytes8, Address_Bytes32:
		return eth.Bytes2Hex(a.GetData())
	case Address_String:
		return string(a.GetData())
	default:
		return ""
	}
}

func (a *Address) Hex() string {
	return a.String()
}

func (a *Address) Bytes() []byte {
	res := make([]byte, len(a.GetData()))
	copy(res, a.GetData())
	return res
}

func (a *Address) Length() uint16 {
	return uint16(len(a.GetData()))
}

// GetEthAddress would return the corresponding eth address of Address if its type is evm or (len=20)hexadecimal,
// else return the zero eth address.
func (a *Address) GetEthAddress() eth.Addr {
	if a.GetType() == Address_Bytes20 {
		return eth.Bytes2Addr(a.GetData())
	}
	if a.GetType() == Address_Hex && a.Length() == Address_Length_20 {
		return eth.Bytes2Addr(a.GetData())
	}
	return eth.ZeroAddr
}

// <1>
// same type        return a copy of A
// <2>
// diff type        if lengths match or toType is string&hex, reset address with proper type
// <3>
// sender
// NOTE: return a copy of A by default
//
//	EVM >> non-EVM, return a copy of A
//	non-EVM >> EVM, longer address will be replaced by full f evm address
//	                shorter address will be padded by 0 from left
//
// <4>
// receiver
// NOTE: return zero addr by default in order to produce error
//
//		EVM >> non-EVM, longer address should be corrected by additional event {Supplement} from TransferAgent contract, and
//	                     returned in above case<2>. If else, return zero addr.
//		                shorter address is gotten from trimming left-padding 0 of evm address
//		non-EVM >> EVM, should be returned in above case<1>. If else, return zero addr.
func (a *Address) convert(dstChainId uint64, isSender bool) *Address {
	if !a.IsValid() {
		return ZeroAddr()
	}
	toType := GetAddrTypeByChainId(dstChainId)
	fromType := a.GetType()
	if fromType == toType {
		return new(Address).Copy(a)
	}
	toLen := AddressTypeLength[toType]
	fromLen := a.Length()
	fromData := a.GetData()
	if fromLen == toLen || toType == Address_String || toType == Address_Hex {
		return new(Address).setDataAndType(fromData, toType)
	}
	if isSender {
		if toType == Address_Bytes20 {
			if fromLen > toLen {
				return new(Address).SetBytes20Address(eth.Hex2Addr(strings.Repeat(`ff`, int(Address_Length_20))))
			} else {
				return new(Address).SetBytes20Address(eth.Bytes2Addr(fromData))
			}
		} else {
			return new(Address).Copy(a)
		}
	} else {
		if fromType == Address_Bytes20 && fromLen > toLen {
			return new(Address).setDataAndType(fromData[fromLen-toLen:], toType)
		}
		return ZeroAddr()
	}
}

func (a *Address) ConvertSender(chid uint64) *Address {
	return a.convert(chid, true)
}

func (a *Address) ConvertReceiver(chid uint64) *Address {
	return a.convert(chid, false)
}

func CheckValidity(_bytes []byte, _type Address_Type) bool {
	// To make results from SetBytesByType be consistent with SetString and SetBytes, we don't block address longer than Address_Max_Length
	//if len(_bytes) > int(AddressTypeLength[_type]) {
	//	return false
	//}
	if _type != Address_String && _type != Address_Hex {
		if len(_bytes) != int(AddressTypeLength[_type]) {
			return false
		}
	}
	return true
}

func GetAddrTypeByChainId(chid uint64) Address_Type {
	if IsFlowChain(chid) {
		return Address_Bytes8
	} else if IsAptosChain(chid) {
		return Address_Bytes32
	} else if chid == 0 {
		return Address_Hex
	} else {
		return Address_Bytes20
	}
}

func GetZeroAddrByChainId(chid uint64) *Address {
	switch GetAddrTypeByChainId(chid) {
	case Address_Bytes8:
		return new(Address).SetBytes8Address([8]byte{})
	case Address_Bytes20:
		return new(Address).SetBytes20Address(eth.ZeroAddr)
	case Address_Bytes32:
		return new(Address).SetBytes32Address([32]byte{})
	default:
		return ZeroAddr()
	}
}

func GetConfiscateAddrByChainId(chid uint64) *Address {
	burnAddr := GetZeroAddrByChainId(chid)
	if burnAddr.Equal(ZeroAddr()) || burnAddr.Length() == 0 {
		// return zero address
		return ZeroAddr()
	}
	burnAddr.GetData()[0] = eth.Hex2Bytes("0xcf")[0]
	return burnAddr
}

// IsHexAddress verifies whether a string can represent a valid hex-encoded address or not.
func IsHexAddress(s string) bool {
	if has0xPrefix(s) {
		s = s[2:]
	}
	return isHex(s)
}

// has0xPrefix validates str begins with '0x' or '0X'.
func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}

// isHexCharacter returns bool of c being a valid hexadecimal.
func isHexCharacter(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'f') || ('A' <= c && c <= 'F')
}

// isHex validates whether each byte is valid hexadecimal string.
func isHex(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	for _, c := range []byte(str) {
		if !isHexCharacter(c) {
			return false
		}
	}
	return true
}

// GetFixedHexAddr For some other server such as gateway, sentinel, im-scan, addr in db need it.
func (a *Address) GetFixedHexAddr() string {
	if a.GetType() != Address_Hex {
		return ""
	}
	if a.Length() <= 20 {
		return eth.Bytes2Addr(a.GetData()).String()
	}
	return "0x" + eth.Bytes2Hex(a.GetData())
}

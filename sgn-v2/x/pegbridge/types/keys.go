package types

import (
	"encoding/binary"
	"fmt"

	commontypes "github.com/celer-network/im-executor/sgn-v2/common/types"
	"github.com/celer-network/im-executor/sgn-v2/eth"
)

const (
	// ModuleName defines the module name
	ModuleName = "pegbridge"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// Byte length of chain ID / nonce occupied
	Uint64ByteArrayLength = 8
	Uint32ByteArrayLength = 4

	PegBridgeFeeDenomPrefix = "PBF-"
)

var (
	OriginalTokenVaultPrefix = []byte{0x01}
	PeggedTokenBridgePrefix  = []byte{0x02}
	OrigPeggedPairPrefix     = []byte{0x03}
	PeggedOrigIndexPrefix    = []byte{0x04}
	DepositInfoPrefix        = []byte{0x05}
	WithdrawInfoPrefix       = []byte{0x06}
	MintInfoPrefix           = []byte{0x07}
	BurnInfoPrefix           = []byte{0x08}
	FeeClaimInfoPrefix       = []byte{0x09}
	TotalSupplyPrefix        = []byte{0x0a}
	RefundPrefix             = []byte{0x0b}
	RefundClaimInfoPrefix    = []byte{0x0c}

	VersionedVaultPrefix  = []byte{0x11}
	VersionedBridgePrefix = []byte{0x12}
	VaultVersionPrefix    = []byte{0x13}
	BridgeVersionPrefix   = []byte{0x14}

	VaultMigrationPrefix = []byte{0x21}

	ConfiscationKeyPrefix = []byte{0x31}

	CfgKeyFeeSplitGlobal = []byte("cfg-feesplit")
)

func GetOriginalTokenVaultKey(chainId uint64) []byte {
	chainIdBytes := make([]byte, Uint64ByteArrayLength)
	binary.LittleEndian.PutUint64(chainIdBytes, chainId)
	return append(OriginalTokenVaultPrefix, chainIdBytes...)
}

func GetPeggedTokenBridgeKey(chainId uint64) []byte {
	chainIdBytes := make([]byte, Uint64ByteArrayLength)
	binary.LittleEndian.PutUint64(chainIdBytes, chainId)
	return append(PeggedTokenBridgePrefix, chainIdBytes...)
}

func GetChainIdAddressBytes(chainId uint64, address *commontypes.Address) []byte {
	return []byte(fmt.Sprintf("%d-%s", chainId, address.String()))
}

func GetOrigPeggedPairKey(origChainId uint64, origAddress *commontypes.Address, peggedChainId uint64) []byte {
	origPeggedBytes := []byte(
		fmt.Sprintf("%s-%d", string(GetChainIdAddressBytes(origChainId, origAddress)), peggedChainId))
	return append(OrigPeggedPairPrefix, origPeggedBytes...)
}

func GetOrigPeggedByOrigPrefix(origChainId uint64, origAddress *commontypes.Address) []byte {
	return append(OrigPeggedPairPrefix, GetChainIdAddressBytes(origChainId, origAddress)...)
}

func GetOrigPeggedByOrigTokenAndPeggedChainIdPrefix(origChainId uint64, origAddress *commontypes.Address, peggedChainId uint64) []byte {
	origTokenAndPeggedChainIdBytes := []byte(
		fmt.Sprintf("%s-%s",
			string(GetChainIdAddressBytes(origChainId, origAddress)),
			fmt.Sprintf("%d-", peggedChainId)))
	return append(OrigPeggedPairPrefix, origTokenAndPeggedChainIdBytes...)
}

func GetPeggedOrigIndexKey(peggedChainId uint64, peggedAddress *commontypes.Address) []byte {
	return append(PeggedOrigIndexPrefix, GetChainIdAddressBytes(peggedChainId, peggedAddress)...)
}

func GetDepositInfoKey(depositId eth.Hash) []byte {
	return append(DepositInfoPrefix, depositId.Bytes()...)
}

func GetWithdrawInfoKey(withdrawId eth.Hash) []byte {
	return append(WithdrawInfoPrefix, withdrawId.Bytes()...)
}

func GetMintInfoKey(mintId eth.Hash) []byte {
	return append(MintInfoPrefix, mintId.Bytes()...)
}

func GetBurnInfoKey(burnId eth.Hash) []byte {
	return append(BurnInfoPrefix, burnId.Bytes()...)
}

func GetFeeClaimInfoKey(address eth.Addr, nonce uint64) []byte {
	nonceBytes := make([]byte, Uint64ByteArrayLength)
	binary.LittleEndian.PutUint64(nonceBytes, nonce)
	return append(FeeClaimInfoPrefix, append(address.Bytes(), nonceBytes...)...)
}

func GetTotalSupplyKey(peggedChainId uint64, peggedAddress *commontypes.Address) []byte {
	composalKeyBytes := GetChainIdAddressBytes(peggedChainId, peggedAddress)
	return append(TotalSupplyPrefix, composalKeyBytes...)
}

func GetRefundKey(depositId eth.Hash) []byte {
	return append(RefundPrefix, depositId.Bytes()...)
}

func GetRefundClaimInfoKey(depositId eth.Hash) []byte {
	return append(RefundClaimInfoPrefix, depositId.Bytes()...)
}

func GetConfiscationKey(chainId uint64, token *commontypes.Address) []byte {
	return append(ConfiscationKeyPrefix, GetChainIdAddressBytes(chainId, token)...)
}

// ---------- new versioned keys -------

func GetVersionedVaultKey(chainId uint64, version uint32) []byte {
	chainIdBytes := make([]byte, Uint64ByteArrayLength)
	binary.LittleEndian.PutUint64(chainIdBytes, chainId)
	versionBytes := make([]byte, Uint32ByteArrayLength)
	binary.LittleEndian.PutUint32(versionBytes, version)
	return append(VersionedVaultPrefix, append(chainIdBytes, versionBytes...)...)
}

func GetVersionedBridgeKey(chainId uint64, version uint32) []byte {
	chainIdBytes := make([]byte, Uint64ByteArrayLength)
	binary.LittleEndian.PutUint64(chainIdBytes, chainId)
	versionBytes := make([]byte, Uint32ByteArrayLength)
	binary.LittleEndian.PutUint32(versionBytes, version)
	return append(VersionedBridgePrefix, append(chainIdBytes, versionBytes...)...)
}

func GetVaultVersionKey(chainId uint64, vaultAddr *commontypes.Address) []byte {
	chainIdBytes := make([]byte, Uint64ByteArrayLength)
	binary.LittleEndian.PutUint64(chainIdBytes, chainId)
	return append(VaultVersionPrefix, append(chainIdBytes, vaultAddr.Bytes()...)...)
}

func GetBridgeVersionKey(chainId uint64, bridgeAddr *commontypes.Address) []byte {
	chainIdBytes := make([]byte, Uint64ByteArrayLength)
	binary.LittleEndian.PutUint64(chainIdBytes, chainId)
	return append(BridgeVersionPrefix, append(chainIdBytes, bridgeAddr.Bytes()...)...)
}

// ---------- vault migration ----------

func GetVaultMigrationKey(chainId uint64, token *commontypes.Address) []byte {
	return append(GetVaultMigrationByChainPrefix(chainId), token.Bytes()...)
}

func GetVaultMigrationByChainPrefix(chainId uint64) []byte {
	chainIdBytes := make([]byte, Uint64ByteArrayLength)
	binary.LittleEndian.PutUint64(chainIdBytes, chainId)
	return append(VaultMigrationPrefix, chainIdBytes...)
}

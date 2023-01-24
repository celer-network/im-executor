package types

import (
	"github.com/celer-network/im-executor/sgn-v2/eth"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/pegbridge module sentinel errors
var (
	ErrNoInfoFound               = sdkerrors.Register(ModuleName, 7100, "no info found")
	ErrNoPeggedTokenBridgeFound  = sdkerrors.Register(ModuleName, 7101, "no PeggedTokenBridge found")
	ErrNoOriginalTokenVaultFound = sdkerrors.Register(ModuleName, 7102, "no OriginalTokenVault found")
)

// WrapErrNoInfoFound returns an error if no info is found for an ID
func WrapErrNoInfoFound(id eth.Hash) error {
	return sdkerrors.Wrapf(ErrNoInfoFound, "%x", id)
}

// WrapErrNoPeggedTokenBridgeFound returns an error if no PeggedTokenBridge contract is found for a chainId
func WrapErrNoPeggedTokenBridgeFound(chainId uint64) error {
	return sdkerrors.Wrapf(ErrNoPeggedTokenBridgeFound, "chainId: %d", chainId)
}

// WrapErrNoOriginalTokenVaultFound returns an error if no OriginalTokenVault contract is found for a chainId
func WrapErrNoOriginalTokenVaultFound(chainId uint64) error {
	return sdkerrors.Wrapf(ErrNoOriginalTokenVaultFound, "chainId: %d", chainId)
}

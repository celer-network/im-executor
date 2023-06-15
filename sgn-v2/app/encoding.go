package app

import (
	"github.com/celer-network/im-executor/sgn-v2/app/params"
	"github.com/cosmos/cosmos-sdk/std"
)

func MakeEncodingConfig() params.EncodingConfig {
	encodingConfig := params.MakeEncodingConfig()
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

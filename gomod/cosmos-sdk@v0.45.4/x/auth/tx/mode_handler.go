package tx

import (
	"fmt"

	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
)


var DefaultSignModes = []signingtypes.SignMode{
	signingtypes.SignMode_SIGN_MODE_DIRECT,
	signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
}



func makeSignModeHandler(modes []signingtypes.SignMode) signing.SignModeHandler {
	if len(modes) < 1 {
		panic(fmt.Errorf("no sign modes enabled"))
	}

	handlers := make([]signing.SignModeHandler, len(modes))

	for i, mode := range modes {
		switch mode {
		case signingtypes.SignMode_SIGN_MODE_DIRECT:
			handlers[i] = signModeDirectHandler{}
		case signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON:
			handlers[i] = signModeLegacyAminoJSONHandler{}
		default:
			panic(fmt.Errorf("unsupported sign mode %+v", mode))
		}
	}

	return signing.NewSignModeHandlerMap(
		modes[0],
		handlers,
	)
}

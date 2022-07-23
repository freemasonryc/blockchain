package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrAddressBookSet = sdkerrors.Register(ModuleName, 201, "address set error")
	ErrGatewayNum     = sdkerrors.Register(ModuleName, 202, "Number segment overrun")
	ErrDelegationCoin = sdkerrors.Register(ModuleName, 203, "Invalid amount")
	ErrGatewayNumber  = sdkerrors.Register(ModuleName, 204, "Number Already registered")
)

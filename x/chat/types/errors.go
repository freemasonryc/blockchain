package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrUserHasExisted = sdkerrors.Register(ModuleName, 7, "user has existed")
	ErrRegister       = sdkerrors.Register(ModuleName, 8, "user register faild")
	ErrUserNotFound   = sdkerrors.Register(ModuleName, 9, "user not found")
	ErrAddressFormat  = sdkerrors.Register(ModuleName, 10, "address format error")
	ErrTransfer       = sdkerrors.Register(ModuleName, 11, "transfer error")
	ErrBurn           = sdkerrors.Register(ModuleName, 12, "burn error")
	ErrUserUpdate     = sdkerrors.Register(ModuleName, 13, "user info update error")
	ErrMortgageAmount = sdkerrors.Register(ModuleName, 14, "mortgage amount error")
	ErrAddressBookSet = sdkerrors.Register(ModuleName, 15, "address set error")
)

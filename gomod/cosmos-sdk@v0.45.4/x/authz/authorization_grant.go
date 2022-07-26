package authz

import (
	"time"

	proto "github.com/gogo/protobuf/proto"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


func NewGrant(  a Authorization, expiration time.Time) (Grant, error) {




	g := Grant{
		Expiration: expiration,
	}
	msg, ok := a.(proto.Message)
	if !ok {
		return Grant{}, sdkerrors.Wrapf(sdkerrors.ErrPackAny, "cannot proto marshal %T", a)
	}

	any, err := cdctypes.NewAnyWithValue(msg)
	if err != nil {
		return Grant{}, err
	}
	g.Authorization = any

	return g, nil
}

var (
	_ cdctypes.UnpackInterfacesMessage = &Grant{}
)


func (g Grant) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	var authorization Authorization
	return unpacker.UnpackAny(g.Authorization, &authorization)
}


func (g Grant) GetAuthorization() Authorization {
	if g.Authorization == nil {
		return nil
	}
	a, ok := g.Authorization.GetCachedValue().(Authorization)
	if !ok {
		return nil
	}
	return a
}

func (g Grant) ValidateBasic() error {
	av := g.Authorization.GetCachedValue()
	a, ok := av.(Authorization)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "expected %T, got %T", (Authorization)(nil), av)
	}
	return a.ValidateBasic()
}

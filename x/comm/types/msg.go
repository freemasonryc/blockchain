package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	_ sdk.Msg = &MsgGatewayRegister{}
	_ sdk.Msg = &MsgAddressBookSave{}
)

const (
	TypeMsgAddressSave     = "address_save"
	TypeMsgGatewayRegister = "gateway_register"
)


func NewMsgAddressSave(fromAddress string, address_book []string) *MsgAddressBookSave {
	return &MsgAddressBookSave{
		FromAddress: fromAddress,
		AddressBook: address_book,
	}
}

func (msg MsgAddressBookSave) Route() string { return RouterKey }
func (msg MsgAddressBookSave) Type() string  { return TypeMsgAddressSave }
func (msg MsgAddressBookSave) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}
func (msg *MsgAddressBookSave) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgAddressBookSave) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid send address")
	}

	return nil
}
func (m MsgAddressBookSave) XXX_MessageName() string {
	return TypeMsgAddressSave
}


func NewMsgGatewayRegister(address, gatewayName, delegation string, indexNumber []string, commission stakingtypes.CommissionRates) *MsgGatewayRegister {
	return &MsgGatewayRegister{
		Address:     address,
		GatewayName: gatewayName,
		Delegation:  delegation,
		IndexNumber: indexNumber,
		Commission:  commission,
	}
}

func (msg MsgGatewayRegister) Route() string { return RouterKey }
func (msg MsgGatewayRegister) Type() string  { return TypeMsgGatewayRegister }
func (msg MsgGatewayRegister) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}
func (msg *MsgGatewayRegister) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgGatewayRegister) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid send address")
	}

	_, err = sdk.ParseCoinNormalized(msg.Delegation)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid Delegation")
	}

	return nil
}
func (m MsgGatewayRegister) XXX_MessageName() string {
	return TypeMsgGatewayRegister
}

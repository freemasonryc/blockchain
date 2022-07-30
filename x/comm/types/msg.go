package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	_ sdk.Msg = &MsgGatewayRegister{}
	_ sdk.Msg = &MsgGatewayDelegate{}
	_ sdk.Msg = &MsgGatewayUndelegate{}
)

const (
	TypeMsgGatewayRegister     = "gateway_register"
	TypeMsgGatewayDelegation   = "gateway_delegation"
	TypeMsgGatewayUndelegation = "gateway_undelegation"
)


func NewMsgGatewayRegister(address, gatewayName, gatewayUrl, delegation, base64ValPubkey string, indexNumber []string, commission stakingtypes.CommissionRates) *MsgGatewayRegister {
	return &MsgGatewayRegister{
		Address:     address,
		GatewayName: gatewayName,
		GatewayUrl:  gatewayUrl,
		Delegation:  delegation,
		IndexNumber: indexNumber,
		Commission:  commission,
		PubKey:      base64ValPubkey,
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
	if err := msg.Commission.Validate(); err != nil {
		return err
	}
	for _, val := range msg.IndexNumber {
		if len(val) != 6 {
			return ErrGatewayNumLength
		}
	}

	return nil
}
func (m MsgGatewayRegister) XXX_MessageName() string {
	return TypeMsgGatewayRegister
}


func NewMsgGatewayDelegation(address, validatorAddress string, amount sdk.Coin, indexNumber []string) *MsgGatewayDelegate {
	return &MsgGatewayDelegate{
		DelegatorAddress: address,
		ValidatorAddress: validatorAddress,
		Amount:           amount,
		IndexNumber:      indexNumber,
	}
}

func (msg MsgGatewayDelegate) Route() string { return RouterKey }
func (msg MsgGatewayDelegate) Type() string  { return TypeMsgGatewayDelegation }
func (msg MsgGatewayDelegate) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}
func (msg *MsgGatewayDelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgGatewayDelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid send address")
	}
	if len(msg.IndexNumber) == 0 && msg.Amount.IsZero() {
		return sdkerrors.Wrap(err, "invalid message")
	}
	for _, val := range msg.IndexNumber {
		if len(val) != 6 {
			return ErrGatewayNumLength
		}
	}
	return nil
}
func (m MsgGatewayDelegate) XXX_MessageName() string {
	return TypeMsgGatewayDelegation
}


func NewMsgGatewayUndelegation(address, validatorAddress string, amount sdk.Coin, indexNumber []string) *MsgGatewayUndelegate {
	return &MsgGatewayUndelegate{
		DelegatorAddress: address,
		ValidatorAddress: validatorAddress,
		Amount:           amount,
		IndexNumber:      indexNumber,
	}
}

func (msg MsgGatewayUndelegate) Route() string { return RouterKey }
func (msg MsgGatewayUndelegate) Type() string  { return TypeMsgGatewayUndelegation }
func (msg MsgGatewayUndelegate) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}
func (msg *MsgGatewayUndelegate) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgGatewayUndelegate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid send address")
	}
	if len(msg.IndexNumber) == 0 && msg.Amount.IsZero() {
		return sdkerrors.Wrap(err, "invalid message")
	}
	for _, val := range msg.IndexNumber {
		if len(val) != 6 {
			return ErrGatewayNumLength
		}
	}

	return nil
}
func (m MsgGatewayUndelegate) XXX_MessageName() string {
	return TypeMsgGatewayUndelegation
}

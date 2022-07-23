package types

import (
	"freemasonry.cc/blockchain/cmd/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ethereum/go-ethereum/common"

	types "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgTest{}
	_ sdk.Msg = &MsgRegister{}
	_ sdk.Msg = &MsgMortgage{}
	_ sdk.Msg = &MsgSendGift{}
	_ sdk.Msg = &MsgSetChatFee{}
)

const (
	TypeMsgConvertCoin = "convert_coin"
	TypeMsgRegister    = "register"
	TypeMsgMortgage    = "mortgage"
	TypeMsgSetChatFee  = "set_chat_fee"
	TypeMsgSendGift    = "send_gift"
)


func NewMsgSendGift(fromAddress, toAddress string, giftId, giftAmount int64, giftValue types.Coin) *MsgSendGift {
	return &MsgSendGift{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		GiftId:      giftId,
		GiftAmount:  giftAmount,
		GiftValue:   giftValue,
	}
}

func (msg MsgSendGift) Route() string { return RouterKey }
func (msg MsgSendGift) Type() string  { return TypeMsgSendGift }
func (msg MsgSendGift) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}
func (msg *MsgSendGift) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgSendGift) ValidateBasic() error {

	if msg.GiftValue.Denom != config.BaseDenom {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "coin error")
	}

	if !msg.GiftValue.Amount.IsPositive() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "amount error")
	}
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}

	_, err = sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid to address")
	}

	_, err = sdk.AccAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid node address")
	}

	return nil
}
func (m MsgSendGift) XXX_MessageName() string {
	return TypeMsgSendGift
}


func NewMsgSetChatFee(fromAddress string, fee types.Coin) *MsgSetChatFee {
	return &MsgSetChatFee{
		FromAddress: fromAddress,
		Fee:         fee,
	}
}

func (msg MsgSetChatFee) Route() string { return RouterKey }
func (msg MsgSetChatFee) Type() string  { return TypeMsgSetChatFee }
func (msg MsgSetChatFee) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}
func (msg *MsgSetChatFee) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgSetChatFee) ValidateBasic() error {

	if msg.Fee.Denom != config.BaseDenom {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "coin error")
	}

	if !msg.Fee.Amount.IsPositive() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "amount error")
	}
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}
	return nil
}
func (m MsgSetChatFee) XXX_MessageName() string {
	return TypeMsgSetChatFee
}


func NewMsgMortgage(fromAddress, node_address string, mortgageAmount types.Coin) *MsgMortgage {
	return &MsgMortgage{
		FromAddress:    fromAddress,
		NodeAddress:    node_address,
		MortgageAmount: mortgageAmount,
	}
}

func (msg MsgMortgage) Route() string { return RouterKey }
func (msg MsgMortgage) Type() string  { return TypeMsgMortgage }
func (msg MsgMortgage) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}
func (msg *MsgMortgage) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgMortgage) ValidateBasic() error {

	if msg.MortgageAmount.Denom != config.BaseDenom {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "coin error")
	}

	if !msg.MortgageAmount.Amount.IsPositive() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "amount error")
	}
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}
	return nil
}
func (m MsgMortgage) XXX_MessageName() string {
	return TypeMsgMortgage
}


func NewMsgRegister(fromAddress, nodeAddress string, mortgageAmount types.Coin) *MsgRegister {
	return &MsgRegister{
		FromAddress:    fromAddress,
		NodeAddress:    nodeAddress,
		MortgageAmount: mortgageAmount,
	}
}
func (msg MsgRegister) Route() string { return RouterKey }
func (msg MsgRegister) Type() string  { return TypeMsgRegister }
func (msg MsgRegister) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}
func (msg *MsgRegister) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgRegister) ValidateBasic() error {

	if msg.MortgageAmount.Denom != config.BaseDenom {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "coin error")
	}

	if !msg.MortgageAmount.Amount.IsPositive() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "amount error")
	}
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}
	_, err = sdk.AccAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}
	return nil
}
func (m MsgRegister) XXX_MessageName() string {
	return TypeMsgRegister
}



func NewMsgTest(coin sdk.Coin, receiver common.Address, sender sdk.AccAddress) *MsgTestResponse { 
	return &MsgTestResponse{}
}


func (msg MsgTest) Route() string { return RouterKey }


func (msg MsgTest) Type() string { return TypeMsgConvertCoin }


func (msg MsgTest) ValidateBasic() error {
	if !msg.Coin.Amount.IsPositive() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "cannot mint a non-positive amount")
	}
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}
	if !common.IsHexAddress(msg.Receiver) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver hex address %s", msg.Receiver)
	}
	return nil
}


func (msg *MsgTest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}


func (msg MsgTest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil
	}

	return []sdk.AccAddress{addr}
}

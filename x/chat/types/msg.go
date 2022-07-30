package types

import (
	"freemasonry.cc/blockchain/cmd/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	types "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgRegister{}
	_ sdk.Msg = &MsgMortgage{}
	_ sdk.Msg = &MsgSendGift{}
	_ sdk.Msg = &MsgSetChatFee{}
	_ sdk.Msg = &MsgAddressBookSave{}
)

const (
	TypeMsgRegister        = "register"
	TypeMsgMortgage        = "mortgage"
	TypeMsgSetChatFee      = "set_chat_fee"
	TypeMsgSendGift        = "send_gift"
	TypeMsgAddressBookSave = "address_book_save"
	TypeMsgGetRewards      = "get_rewards"
	TypeMsgMobileTransfer  = "mobile_transfer"
)


func NewMsgMobileTransfer(fromAddress, toAddress, mobile string) *MsgMobileTransfer {
	return &MsgMobileTransfer{
		FromAddress: fromAddress,
		ToAddress:   toAddress,
		Mobile:      mobile,
	}
}

func (msg MsgMobileTransfer) Route() string { return RouterKey }
func (msg MsgMobileTransfer) Type() string  { return TypeMsgMobileTransfer }
func (msg MsgMobileTransfer) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}
func (msg *MsgMobileTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgMobileTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}

	_, err = sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}

	return nil
}
func (m MsgMobileTransfer) XXX_MessageName() string {
	return TypeMsgMobileTransfer
}


func NewMsgGetRewards(fromAddress string) *MsgGetRewards {
	return &MsgGetRewards{
		FromAddress: fromAddress,
	}
}

func (msg MsgGetRewards) Route() string { return RouterKey }
func (msg MsgGetRewards) Type() string  { return TypeMsgGetRewards }
func (msg MsgGetRewards) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil
	}
	return []sdk.AccAddress{addr}
}
func (msg *MsgGetRewards) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
func (msg MsgGetRewards) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}

	return nil
}
func (m MsgGetRewards) XXX_MessageName() string {
	return TypeMsgGetRewards
}


func NewMsgAddressBookSave(fromAddress string, AddressBook []string) *MsgAddressBookSave {
	return &MsgAddressBookSave{
		FromAddress: fromAddress,
		AddressBook: AddressBook,
	}
}

func (msg MsgAddressBookSave) Route() string { return RouterKey }
func (msg MsgAddressBookSave) Type() string  { return TypeMsgAddressBookSave }
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
		return sdkerrors.Wrap(err, "invalid sender address")
	}

	return nil
}
func (m MsgAddressBookSave) XXX_MessageName() string {
	return TypeMsgAddressBookSave
}


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

	_, err = sdk.ValAddressFromBech32(msg.NodeAddress)
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


func NewMsgRegister(fromAddress, nodeAddress, mobilePrefix string, mortgageAmount types.Coin) *MsgRegister {
	return &MsgRegister{
		FromAddress:    fromAddress,
		NodeAddress:    nodeAddress,
		MortgageAmount: mortgageAmount,
		MobilePrefix:   mobilePrefix,
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
	_, err = sdk.ValAddressFromBech32(msg.NodeAddress)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid sender address")
	}
	return nil
}
func (m MsgRegister) XXX_MessageName() string {
	return TypeMsgRegister
}



package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)



func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	
	cdc.RegisterConcrete(&MsgRegister{}, MsgTypeRegister, nil)
	cdc.RegisterConcrete(&MsgMortgage{}, MsgTypeMortgage, nil)
	cdc.RegisterConcrete(&MsgSetChatFee{}, MsgTypeSetChatFee, nil)
	cdc.RegisterConcrete(&MsgSendGift{}, MsgTypeSendGift, nil)
	cdc.RegisterConcrete(&MsgAddressBookSave{}, MsgTypeAddressBookSave, nil)
	cdc.RegisterConcrete(&MsgGetRewards{}, MsgTypeGetRewards, nil)
	cdc.RegisterConcrete(&MsgMobileTransfer{}, MsgTypeMobileTransfer, nil)
}


func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegister{},
		&MsgMortgage{},
		&MsgSetChatFee{},
		&MsgSendGift{},
		&MsgAddressBookSave{},
		&MsgGetRewards{},
		&MsgMobileTransfer{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

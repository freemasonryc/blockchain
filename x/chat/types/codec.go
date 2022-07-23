package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {

	cdc.RegisterConcrete(&MsgRegister{}, MsgTypeRegister, nil)
	cdc.RegisterConcrete(&MsgMortgage{}, MsgTypeMortgage, nil)
	cdc.RegisterConcrete(&MsgSetChatFee{}, MsgTypeSetChatFee, nil)
	cdc.RegisterConcrete(&MsgSendGift{}, MsgTypeSendGift, nil)
}


func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgTest{},
		&MsgRegister{},
		&MsgMortgage{},
		&MsgSetChatFee{},
		&MsgSendGift{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

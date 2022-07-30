package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {

	cdc.RegisterConcrete(&MsgGatewayRegister{}, MSG_GATEWAY_REGISTER, nil)
	cdc.RegisterConcrete(&MsgGatewayDelegate{}, MSG_GATEWAY_DELEGATION, nil)
	cdc.RegisterConcrete(&MsgGatewayUndelegate{}, MSG_GATEWAY_UNDELEGATION, nil)
}


func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgGatewayRegister{},
		&MsgGatewayDelegate{},
		&MsgGatewayUndelegate{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

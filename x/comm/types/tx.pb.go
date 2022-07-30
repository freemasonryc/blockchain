


package types

import (
	context "context"
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)


var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf





const _ = proto.GoGoProtoPackageIsVersion3 


type MsgGatewayRegister struct {
	
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	
	GatewayName string `protobuf:"bytes,2,opt,name=gateway_name,json=gatewayName,proto3" json:"gateway_name,omitempty"`
	
	GatewayUrl string `protobuf:"bytes,3,opt,name=gateway_url,json=gatewayUrl,proto3" json:"gateway_url,omitempty"`
	
	Delegation string `protobuf:"bytes,4,opt,name=delegation,proto3" json:"delegation,omitempty"`
	
	IndexNumber []string `protobuf:"bytes,5,rep,name=index_number,json=indexNumber,proto3" json:"index_number,omitempty"`
	
	PubKey               string                `protobuf:"bytes,6,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	Commission           types.CommissionRates `protobuf:"bytes,7,opt,name=commission,proto3" json:"commission"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MsgGatewayRegister) Reset()         { *m = MsgGatewayRegister{} }
func (m *MsgGatewayRegister) String() string { return proto.CompactTextString(m) }
func (*MsgGatewayRegister) ProtoMessage()    {}
func (*MsgGatewayRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{0}
}
func (m *MsgGatewayRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgGatewayRegister.Unmarshal(m, b)
}
func (m *MsgGatewayRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgGatewayRegister.Marshal(b, m, deterministic)
}
func (m *MsgGatewayRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGatewayRegister.Merge(m, src)
}
func (m *MsgGatewayRegister) XXX_Size() int {
	return xxx_messageInfo_MsgGatewayRegister.Size(m)
}
func (m *MsgGatewayRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGatewayRegister.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGatewayRegister proto.InternalMessageInfo

func (m *MsgGatewayRegister) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgGatewayRegister) GetGatewayName() string {
	if m != nil {
		return m.GatewayName
	}
	return ""
}

func (m *MsgGatewayRegister) GetGatewayUrl() string {
	if m != nil {
		return m.GatewayUrl
	}
	return ""
}

func (m *MsgGatewayRegister) GetDelegation() string {
	if m != nil {
		return m.Delegation
	}
	return ""
}

func (m *MsgGatewayRegister) GetIndexNumber() []string {
	if m != nil {
		return m.IndexNumber
	}
	return nil
}

func (m *MsgGatewayRegister) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func (m *MsgGatewayRegister) GetCommission() types.CommissionRates {
	if m != nil {
		return m.Commission
	}
	return types.CommissionRates{}
}


type MsgGatewayDelegate struct {
	DelegatorAddress string      `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty" yaml:"delegator_address"`
	ValidatorAddress string      `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`
	Amount           types1.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
	
	IndexNumber          []string `protobuf:"bytes,4,rep,name=index_number,json=indexNumber,proto3" json:"index_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgGatewayDelegate) Reset()         { *m = MsgGatewayDelegate{} }
func (m *MsgGatewayDelegate) String() string { return proto.CompactTextString(m) }
func (*MsgGatewayDelegate) ProtoMessage()    {}
func (*MsgGatewayDelegate) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{1}
}
func (m *MsgGatewayDelegate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgGatewayDelegate.Unmarshal(m, b)
}
func (m *MsgGatewayDelegate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgGatewayDelegate.Marshal(b, m, deterministic)
}
func (m *MsgGatewayDelegate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGatewayDelegate.Merge(m, src)
}
func (m *MsgGatewayDelegate) XXX_Size() int {
	return xxx_messageInfo_MsgGatewayDelegate.Size(m)
}
func (m *MsgGatewayDelegate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGatewayDelegate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGatewayDelegate proto.InternalMessageInfo


type MsgGatewayUndelegate struct {
	DelegatorAddress string      `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty" yaml:"delegator_address"`
	ValidatorAddress string      `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`
	Amount           types1.Coin `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount"`
	
	IndexNumber          []string `protobuf:"bytes,4,rep,name=index_number,json=indexNumber,proto3" json:"index_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgGatewayUndelegate) Reset()         { *m = MsgGatewayUndelegate{} }
func (m *MsgGatewayUndelegate) String() string { return proto.CompactTextString(m) }
func (*MsgGatewayUndelegate) ProtoMessage()    {}
func (*MsgGatewayUndelegate) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{2}
}
func (m *MsgGatewayUndelegate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgGatewayUndelegate.Unmarshal(m, b)
}
func (m *MsgGatewayUndelegate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgGatewayUndelegate.Marshal(b, m, deterministic)
}
func (m *MsgGatewayUndelegate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGatewayUndelegate.Merge(m, src)
}
func (m *MsgGatewayUndelegate) XXX_Size() int {
	return xxx_messageInfo_MsgGatewayUndelegate.Size(m)
}
func (m *MsgGatewayUndelegate) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGatewayUndelegate.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGatewayUndelegate proto.InternalMessageInfo


type MsgEmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgEmptyResponse) Reset()         { *m = MsgEmptyResponse{} }
func (m *MsgEmptyResponse) String() string { return proto.CompactTextString(m) }
func (*MsgEmptyResponse) ProtoMessage()    {}
func (*MsgEmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{3}
}
func (m *MsgEmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgEmptyResponse.Unmarshal(m, b)
}
func (m *MsgEmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgEmptyResponse.Marshal(b, m, deterministic)
}
func (m *MsgEmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgEmptyResponse.Merge(m, src)
}
func (m *MsgEmptyResponse) XXX_Size() int {
	return xxx_messageInfo_MsgEmptyResponse.Size(m)
}
func (m *MsgEmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgEmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgEmptyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgGatewayRegister)(nil), "freemasonry.comm.v1.MsgGatewayRegister")
	proto.RegisterType((*MsgGatewayDelegate)(nil), "freemasonry.comm.v1.MsgGatewayDelegate")
	proto.RegisterType((*MsgGatewayUndelegate)(nil), "freemasonry.comm.v1.MsgGatewayUndelegate")
	proto.RegisterType((*MsgEmptyResponse)(nil), "freemasonry.comm.v1.MsgEmptyResponse")
}

func init() { proto.RegisterFile("tx.proto", fileDescriptor_0fd2153dc07d3b5c) }

var fileDescriptor_0fd2153dc07d3b5c = []byte{
	
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xcd, 0xdf, 0x97, 0xf4, 0x9b, 0x20, 0x91, 0x0c, 0x95, 0x70, 0xa3, 0xaa, 0x09, 0x16, 0xd0,
	0xb0, 0xb1, 0x95, 0xb0, 0x40, 0xea, 0x8e, 0x00, 0x42, 0x15, 0x4a, 0x17, 0x96, 0xba, 0x61, 0x63,
	0x8d, 0xed, 0xcb, 0x60, 0xc5, 0x33, 0x63, 0x79, 0x26, 0x21, 0x7e, 0x03, 0x96, 0x88, 0x3d, 0x52,
	0x1f, 0x82, 0x87, 0xe0, 0x29, 0xba, 0x66, 0xcd, 0x13, 0x20, 0xdb, 0xe3, 0x24, 0x6e, 0x40, 0xed,
	0x03, 0xb0, 0xf3, 0x3d, 0xe7, 0xcc, 0x3d, 0x73, 0x8f, 0x67, 0x06, 0x1d, 0xa8, 0xb5, 0x15, 0x27,
	0x42, 0x09, 0xfc, 0xe0, 0x43, 0x02, 0xc0, 0x88, 0x14, 0x3c, 0x49, 0x2d, 0x5f, 0x30, 0x66, 0xad,
	0x26, 0x83, 0x63, 0x2a, 0x04, 0x8d, 0xc0, 0x26, 0x71, 0x68, 0x13, 0xce, 0x85, 0x22, 0x2a, 0x14,
	0x5c, 0x16, 0x4b, 0x06, 0x87, 0x54, 0x50, 0x91, 0x7f, 0xda, 0xd9, 0x97, 0x46, 0x8f, 0x7c, 0x21,
	0x99, 0x90, 0x6e, 0x41, 0x14, 0x85, 0xa6, 0x4e, 0x8a, 0xca, 0xf6, 0x88, 0x04, 0x7b, 0x35, 0xf1,
	0x40, 0x91, 0x89, 0xed, 0x8b, 0x90, 0x6b, 0xfe, 0xb1, 0xe6, 0xa5, 0x22, 0x8b, 0x90, 0xd3, 0x8d,
	0x44, 0xd7, 0x85, 0xca, 0xfc, 0xd6, 0x40, 0x78, 0x2e, 0xe9, 0x5b, 0xa2, 0xe0, 0x13, 0x49, 0x1d,
	0xa0, 0xa1, 0x54, 0x90, 0x60, 0x03, 0x75, 0x48, 0x10, 0x24, 0x20, 0xa5, 0x51, 0x1f, 0xd5, 0xc7,
	0xff, 0x3b, 0x65, 0x89, 0x1f, 0xa1, 0x7b, 0xb4, 0x10, 0xbb, 0x9c, 0x30, 0x30, 0x1a, 0x39, 0xdd,
	0xd5, 0xd8, 0x05, 0x61, 0x80, 0x87, 0xa8, 0x2c, 0xdd, 0x65, 0x12, 0x19, 0xcd, 0x5c, 0x81, 0x34,
	0x74, 0x99, 0x44, 0xf8, 0x04, 0xa1, 0x00, 0x22, 0xa0, 0x79, 0x00, 0x46, 0xab, 0xe0, 0xb7, 0x48,
	0xe6, 0x11, 0xf2, 0x00, 0xd6, 0x2e, 0x5f, 0x32, 0x0f, 0x12, 0xe3, 0xbf, 0x51, 0x33, 0xf3, 0xc8,
	0xb1, 0x8b, 0x1c, 0xc2, 0x0f, 0x51, 0x27, 0x5e, 0x7a, 0xee, 0x02, 0x52, 0xa3, 0x9d, 0xaf, 0x6f,
	0xc7, 0x4b, 0xef, 0x1d, 0xa4, 0x78, 0x8e, 0x50, 0x16, 0x78, 0x28, 0x65, 0xd6, 0xbb, 0x33, 0xaa,
	0x8f, 0xbb, 0xd3, 0x53, 0x4b, 0x27, 0x57, 0xce, 0xae, 0xb3, 0xb0, 0x5e, 0x6d, 0x94, 0x0e, 0x51,
	0x20, 0x67, 0xad, 0x1f, 0xd7, 0xc3, 0x9a, 0xb3, 0xd3, 0xc0, 0xfc, 0x52, 0xc9, 0xe7, 0x75, 0xb1,
	0x47, 0xc0, 0xe7, 0xa8, 0xaf, 0xf7, 0x2b, 0x12, 0xb7, 0x92, 0xd4, 0xec, 0xf8, 0xd7, 0xf5, 0xd0,
	0x48, 0x09, 0x8b, 0xce, 0xcc, 0x3d, 0x89, 0xe9, 0xf4, 0x36, 0xd8, 0x4b, 0x1d, 0xe8, 0x39, 0xea,
	0xaf, 0x48, 0x14, 0x06, 0x95, 0x56, 0x8d, 0x9b, 0xad, 0xf6, 0x24, 0xa6, 0xd3, 0xdb, 0x60, 0x65,
	0xab, 0x17, 0xa8, 0x4d, 0x98, 0x58, 0x72, 0x95, 0x67, 0xde, 0x9d, 0x1e, 0x95, 0x73, 0x67, 0x67,
	0x64, 0x67, 0xe8, 0x90, 0xeb, 0x49, 0xb5, 0x7c, 0x2f, 0xf0, 0xd6, 0x5e, 0xe0, 0x67, 0x07, 0x9f,
	0xaf, 0x86, 0xb5, 0x9f, 0x57, 0xc3, 0x9a, 0xf9, 0xb5, 0x81, 0x0e, 0xb7, 0x91, 0x5c, 0xf2, 0xe0,
	0x5f, 0x28, 0x35, 0x13, 0xa3, 0xde, 0x5c, 0xd2, 0x37, 0x2c, 0x56, 0xa9, 0x03, 0x32, 0x16, 0x5c,
	0xc2, 0xf4, 0x7b, 0x03, 0x35, 0xe7, 0x92, 0x62, 0x82, 0xee, 0xdf, 0xbc, 0x5f, 0xa7, 0xd6, 0x1f,
	0x5e, 0x08, 0x6b, 0xff, 0x22, 0x0e, 0x9e, 0xfc, 0x4d, 0x58, 0xb1, 0xc2, 0x3e, 0xea, 0x57, 0x8f,
	0x68, 0x76, 0x8d, 0x6e, 0x33, 0x29, 0x4f, 0xf3, 0x5d, 0x4d, 0x60, 0x63, 0xb2, 0xf3, 0xd3, 0x9f,
	0xdd, 0x62, 0xb2, 0x95, 0xde, 0xd1, 0x66, 0x36, 0x7e, 0xff, 0xb4, 0xa2, 0xf3, 0x6d, 0x2f, 0x12,
	0xfe, 0xc2, 0xff, 0x48, 0x42, 0x6e, 0xaf, 0xed, 0x6c, 0x9d, 0xad, 0xd2, 0x18, 0xa4, 0xd7, 0xce,
	0xdf, 0xb0, 0xe7, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x98, 0x75, 0x21, 0x79, 0x05, 0x00,
	0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type MsgClient interface {
	
	GatewayRegister(ctx context.Context, in *MsgGatewayRegister, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	
	GatewayDelegation(ctx context.Context, in *MsgGatewayDelegate, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	
	GatewayUndelegate(ctx context.Context, in *MsgGatewayUndelegate, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
}

type msgClient struct {
	cc *grpc.ClientConn
}

func NewMsgClient(cc *grpc.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) GatewayRegister(ctx context.Context, in *MsgGatewayRegister, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.comm.v1.Msg/GatewayRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GatewayDelegation(ctx context.Context, in *MsgGatewayDelegate, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.comm.v1.Msg/GatewayDelegation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GatewayUndelegate(ctx context.Context, in *MsgGatewayUndelegate, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.comm.v1.Msg/GatewayUndelegate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type MsgServer interface {
	
	GatewayRegister(context.Context, *MsgGatewayRegister) (*MsgEmptyResponse, error)
	
	GatewayDelegation(context.Context, *MsgGatewayDelegate) (*MsgEmptyResponse, error)
	
	GatewayUndelegate(context.Context, *MsgGatewayUndelegate) (*MsgEmptyResponse, error)
}


type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) GatewayRegister(ctx context.Context, req *MsgGatewayRegister) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GatewayRegister not implemented")
}
func (*UnimplementedMsgServer) GatewayDelegation(ctx context.Context, req *MsgGatewayDelegate) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GatewayDelegation not implemented")
}
func (*UnimplementedMsgServer) GatewayUndelegate(ctx context.Context, req *MsgGatewayUndelegate) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GatewayUndelegate not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_GatewayRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGatewayRegister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GatewayRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.comm.v1.Msg/GatewayRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GatewayRegister(ctx, req.(*MsgGatewayRegister))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GatewayDelegation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGatewayDelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GatewayDelegation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.comm.v1.Msg/GatewayDelegation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GatewayDelegation(ctx, req.(*MsgGatewayDelegate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GatewayUndelegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGatewayUndelegate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GatewayUndelegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.comm.v1.Msg/GatewayUndelegate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GatewayUndelegate(ctx, req.(*MsgGatewayUndelegate))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "freemasonry.comm.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GatewayRegister",
			Handler:    _Msg_GatewayRegister_Handler,
		},
		{
			MethodName: "GatewayDelegation",
			Handler:    _Msg_GatewayDelegation_Handler,
		},
		{
			MethodName: "GatewayUndelegate",
			Handler:    _Msg_GatewayUndelegate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tx.proto",
}

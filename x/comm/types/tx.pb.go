


package types

import (
	context "context"
	fmt "fmt"
	cdcTypes "github.com/cosmos/cosmos-sdk/codec/types"
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


type MsgAddressBookSave struct {
	FromAddress          string   `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	AddressBook          []string `protobuf:"bytes,2,rep,name=address_book,json=addressBook,proto3" json:"address_book,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgAddressBookSave) Reset()         { *m = MsgAddressBookSave{} }
func (m *MsgAddressBookSave) String() string { return proto.CompactTextString(m) }
func (*MsgAddressBookSave) ProtoMessage()    {}
func (*MsgAddressBookSave) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{0}
}
func (m *MsgAddressBookSave) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgAddressBookSave.Unmarshal(m, b)
}
func (m *MsgAddressBookSave) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgAddressBookSave.Marshal(b, m, deterministic)
}
func (m *MsgAddressBookSave) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddressBookSave.Merge(m, src)
}
func (m *MsgAddressBookSave) XXX_Size() int {
	return xxx_messageInfo_MsgAddressBookSave.Size(m)
}
func (m *MsgAddressBookSave) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddressBookSave.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddressBookSave proto.InternalMessageInfo

func (m *MsgAddressBookSave) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgAddressBookSave) GetAddressBook() []string {
	if m != nil {
		return m.AddressBook
	}
	return nil
}


type MsgGatewayRegister struct {
	
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	
	GatewayName string `protobuf:"bytes,2,opt,name=gateway_name,json=gatewayName,proto3" json:"gateway_name,omitempty"`
	
	Delegation string `protobuf:"bytes,3,opt,name=delegation,proto3" json:"delegation,omitempty"`
	
	IndexNumber []string `protobuf:"bytes,4,rep,name=index_number,json=indexNumber,proto3" json:"index_number,omitempty"`
	
	Pubkey               *cdcTypes.Any         `protobuf:"bytes,5,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Commission           types.CommissionRates `protobuf:"bytes,6,opt,name=commission,proto3" json:"commission"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MsgGatewayRegister) Reset()         { *m = MsgGatewayRegister{} }
func (m *MsgGatewayRegister) String() string { return proto.CompactTextString(m) }
func (*MsgGatewayRegister) ProtoMessage()    {}
func (*MsgGatewayRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{1}
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

func (m *MsgGatewayRegister) GetPubkey() *cdcTypes.Any {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

func (m *MsgGatewayRegister) GetCommission() types.CommissionRates {
	if m != nil {
		return m.Commission
	}
	return types.CommissionRates{}
}


type MsgEmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgEmptyResponse) Reset()         { *m = MsgEmptyResponse{} }
func (m *MsgEmptyResponse) String() string { return proto.CompactTextString(m) }
func (*MsgEmptyResponse) ProtoMessage()    {}
func (*MsgEmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{2}
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
	proto.RegisterType((*MsgAddressBookSave)(nil), "freemasonry.comm.v1.MsgAddressBookSave")
	proto.RegisterType((*MsgGatewayRegister)(nil), "freemasonry.comm.v1.MsgGatewayRegister")
	proto.RegisterType((*MsgEmptyResponse)(nil), "freemasonry.comm.v1.MsgEmptyResponse")
}

func init() { proto.RegisterFile("tx.proto", fileDescriptor_0fd2153dc07d3b5c) }

var fileDescriptor_0fd2153dc07d3b5c = []byte{
	
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x69, 0xbb, 0x14, 0xf0, 0x22, 0x81, 0x42, 0x0f, 0xd9, 0x0a, 0x41, 0xa9, 0x80, 0xed,
	0x05, 0x5b, 0x5d, 0x9e, 0x60, 0x8b, 0x80, 0x03, 0xea, 0x0a, 0x85, 0xdb, 0x5e, 0x22, 0x27, 0x9d,
	0x9a, 0x28, 0xb5, 0x27, 0x8a, 0xdd, 0x52, 0x3f, 0x13, 0x57, 0x1e, 0x02, 0xc1, 0x3b, 0xf0, 0x2c,
	0xc8, 0x7f, 0x0a, 0xdb, 0xc2, 0x4a, 0x7b, 0xcb, 0x7c, 0xf3, 0x9b, 0x6f, 0xac, 0xf9, 0x42, 0xee,
	0x9a, 0x2d, 0x6d, 0x5a, 0x34, 0x98, 0x3c, 0x5a, 0xb6, 0x00, 0x92, 0x6b, 0x54, 0xad, 0xa5, 0x25,
	0x4a, 0x49, 0x37, 0xd3, 0xe1, 0x89, 0x40, 0x14, 0x2b, 0x60, 0x1e, 0x29, 0xd6, 0x4b, 0xc6, 0x95,
	0x0d, 0xfc, 0xf0, 0x71, 0x6c, 0xf1, 0xa6, 0x62, 0x5c, 0x29, 0x34, 0xdc, 0x54, 0xa8, 0x74, 0xec,
	0x0e, 0x04, 0x0a, 0xf4, 0x9f, 0xcc, 0x7d, 0x45, 0xf5, 0xa4, 0x44, 0x2d, 0x51, 0xe7, 0xa1, 0x11,
	0x8a, 0xd8, 0x7a, 0x1e, 0x2a, 0xa6, 0x0d, 0xaf, 0x2b, 0x25, 0xd8, 0x66, 0x5a, 0x80, 0xe1, 0xd3,
	0x5d, 0x1d, 0xa8, 0xf1, 0x25, 0x49, 0xe6, 0x5a, 0x9c, 0x2f, 0x16, 0x2d, 0x68, 0x3d, 0x43, 0xac,
	0x3f, 0xf1, 0x0d, 0x24, 0xcf, 0xc8, 0xfd, 0x65, 0x8b, 0x32, 0xe7, 0x41, 0x4f, 0x3b, 0xa3, 0xce,
	0xe4, 0x5e, 0x76, 0xec, 0xb4, 0x88, 0x3a, 0x24, 0x76, 0xf3, 0x02, 0xb1, 0x4e, 0xbb, 0xa3, 0x9e,
	0x43, 0xf8, 0x5f, 0xa7, 0xf1, 0xd7, 0xae, 0x37, 0x7f, 0xcf, 0x0d, 0x7c, 0xe1, 0x36, 0x03, 0x51,
	0x69, 0x03, 0x6d, 0x92, 0x92, 0x3b, 0xfb, 0xbe, 0xbb, 0xd2, 0x79, 0x8a, 0x00, 0xe7, 0x8a, 0x4b,
	0x48, 0xbb, 0x61, 0x6d, 0xd4, 0x2e, 0xb8, 0x84, 0xe4, 0x09, 0x21, 0x0b, 0x58, 0x81, 0xf0, 0xb7,
	0x49, 0x7b, 0x1e, 0xb8, 0xa2, 0x38, 0x8b, 0x4a, 0x2d, 0x60, 0x9b, 0xab, 0xb5, 0x2c, 0xa0, 0x4d,
	0x8f, 0xc2, 0xb3, 0xbc, 0x76, 0xe1, 0xa5, 0xe4, 0x1d, 0xe9, 0x37, 0xeb, 0xa2, 0x06, 0x9b, 0xde,
	0x1e, 0x75, 0x26, 0xc7, 0x67, 0x03, 0x1a, 0x0e, 0x4f, 0x77, 0x99, 0xd0, 0x73, 0x65, 0x67, 0xe9,
	0x8f, 0x6f, 0xaf, 0x06, 0xf1, 0xa0, 0x65, 0x6b, 0x1b, 0x83, 0xf4, 0xe3, 0xba, 0xf8, 0x00, 0x36,
	0x8b, 0xd3, 0xc9, 0x9c, 0x10, 0x97, 0x6a, 0xa5, 0xb5, 0x7b, 0x4a, 0xdf, 0x7b, 0x9d, 0xd2, 0x38,
	0xb2, 0xbb, 0x72, 0xbc, 0x3a, 0x7d, 0xf3, 0x87, 0xcc, 0xb8, 0x01, 0x3d, 0x3b, 0xfa, 0xfe, 0xeb,
	0xe9, 0xad, 0xec, 0x8a, 0xc1, 0x38, 0x21, 0x0f, 0xe7, 0x5a, 0xbc, 0x95, 0x8d, 0xb1, 0x19, 0xe8,
	0x06, 0x95, 0x86, 0xb3, 0x9f, 0x1d, 0xd2, 0x9b, 0x6b, 0x91, 0x70, 0xf2, 0xe0, 0x30, 0xa2, 0x53,
	0xfa, 0x9f, 0xdf, 0x8b, 0xfe, 0x9b, 0xe5, 0xf0, 0xc5, 0x75, 0xe0, 0xde, 0x2a, 0xb7, 0xe2, 0x30,
	0xa8, 0x6b, 0x57, 0x1c, 0x80, 0x37, 0x5c, 0x31, 0x9b, 0x5c, 0xbe, 0xdc, 0xe3, 0x4a, 0x56, 0xac,
	0xb0, 0xac, 0xcb, 0xcf, 0xbc, 0x52, 0x6c, 0xcb, 0xdc, 0x1c, 0x33, 0xb6, 0x01, 0x5d, 0xf4, 0x7d,
	0x14, 0xaf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x59, 0x9d, 0xdb, 0x4d, 0x03, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type MsgClient interface {
	
	AddressBookSave(ctx context.Context, in *MsgAddressBookSave, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	
	GatewayRegister(ctx context.Context, in *MsgGatewayRegister, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
}

type msgClient struct {
	cc *grpc.ClientConn
}

func NewMsgClient(cc *grpc.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddressBookSave(ctx context.Context, in *MsgAddressBookSave, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.comm.v1.Msg/AddressBookSave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GatewayRegister(ctx context.Context, in *MsgGatewayRegister, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.comm.v1.Msg/GatewayRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type MsgServer interface {
	
	AddressBookSave(context.Context, *MsgAddressBookSave) (*MsgEmptyResponse, error)
	
	GatewayRegister(context.Context, *MsgGatewayRegister) (*MsgEmptyResponse, error)
}


type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddressBookSave(ctx context.Context, req *MsgAddressBookSave) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressBookSave not implemented")
}
func (*UnimplementedMsgServer) GatewayRegister(ctx context.Context, req *MsgGatewayRegister) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GatewayRegister not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddressBookSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddressBookSave)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddressBookSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.comm.v1.Msg/AddressBookSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddressBookSave(ctx, req.(*MsgAddressBookSave))
	}
	return interceptor(ctx, in, info, handler)
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

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "freemasonry.comm.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddressBookSave",
			Handler:    _Msg_AddressBookSave_Handler,
		},
		{
			MethodName: "GatewayRegister",
			Handler:    _Msg_GatewayRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tx.proto",
}

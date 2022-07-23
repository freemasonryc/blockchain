


package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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


type MsgTest struct {


	Coin types.Coin `protobuf:"bytes,1,opt,name=coin,proto3" json:"coin"`

	Receiver string `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`

	Sender               string   `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgTest) Reset()         { *m = MsgTest{} }
func (m *MsgTest) String() string { return proto.CompactTextString(m) }
func (*MsgTest) ProtoMessage()    {}
func (*MsgTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{0}
}
func (m *MsgTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgTest.Unmarshal(m, b)
}
func (m *MsgTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgTest.Marshal(b, m, deterministic)
}
func (m *MsgTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTest.Merge(m, src)
}
func (m *MsgTest) XXX_Size() int {
	return xxx_messageInfo_MsgTest.Size(m)
}
func (m *MsgTest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTest proto.InternalMessageInfo

func (m *MsgTest) GetCoin() types.Coin {
	if m != nil {
		return m.Coin
	}
	return types.Coin{}
}

func (m *MsgTest) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func (m *MsgTest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

type MsgRegister struct {
	FromAddress          string     `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	NodeAddress          string     `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
	MortgageAmount       types.Coin `protobuf:"bytes,3,opt,name=mortgage_amount,json=mortgageAmount,proto3" json:"mortgage_amount" yaml:"mortgage_amount"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MsgRegister) Reset()         { *m = MsgRegister{} }
func (m *MsgRegister) String() string { return proto.CompactTextString(m) }
func (*MsgRegister) ProtoMessage()    {}
func (*MsgRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{1}
}
func (m *MsgRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgRegister.Unmarshal(m, b)
}
func (m *MsgRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgRegister.Marshal(b, m, deterministic)
}
func (m *MsgRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegister.Merge(m, src)
}
func (m *MsgRegister) XXX_Size() int {
	return xxx_messageInfo_MsgRegister.Size(m)
}
func (m *MsgRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegister.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegister proto.InternalMessageInfo

func (m *MsgRegister) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgRegister) GetNodeAddress() string {
	if m != nil {
		return m.NodeAddress
	}
	return ""
}

func (m *MsgRegister) GetMortgageAmount() types.Coin {
	if m != nil {
		return m.MortgageAmount
	}
	return types.Coin{}
}

type MsgMortgage struct {
	FromAddress          string     `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	NodeAddress          string     `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
	MortgageAmount       types.Coin `protobuf:"bytes,3,opt,name=mortgage_amount,json=mortgageAmount,proto3" json:"mortgage_amount" yaml:"mortgage_amount"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MsgMortgage) Reset()         { *m = MsgMortgage{} }
func (m *MsgMortgage) String() string { return proto.CompactTextString(m) }
func (*MsgMortgage) ProtoMessage()    {}
func (*MsgMortgage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{2}
}
func (m *MsgMortgage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgMortgage.Unmarshal(m, b)
}
func (m *MsgMortgage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgMortgage.Marshal(b, m, deterministic)
}
func (m *MsgMortgage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMortgage.Merge(m, src)
}
func (m *MsgMortgage) XXX_Size() int {
	return xxx_messageInfo_MsgMortgage.Size(m)
}
func (m *MsgMortgage) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMortgage.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMortgage proto.InternalMessageInfo

func (m *MsgMortgage) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgMortgage) GetNodeAddress() string {
	if m != nil {
		return m.NodeAddress
	}
	return ""
}

func (m *MsgMortgage) GetMortgageAmount() types.Coin {
	if m != nil {
		return m.MortgageAmount
	}
	return types.Coin{}
}

type MsgSetChatFee struct {
	FromAddress          string     `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	Fee                  types.Coin `protobuf:"bytes,2,opt,name=fee,proto3" json:"fee" yaml:"fee"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MsgSetChatFee) Reset()         { *m = MsgSetChatFee{} }
func (m *MsgSetChatFee) String() string { return proto.CompactTextString(m) }
func (*MsgSetChatFee) ProtoMessage()    {}
func (*MsgSetChatFee) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{3}
}
func (m *MsgSetChatFee) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSetChatFee.Unmarshal(m, b)
}
func (m *MsgSetChatFee) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSetChatFee.Marshal(b, m, deterministic)
}
func (m *MsgSetChatFee) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetChatFee.Merge(m, src)
}
func (m *MsgSetChatFee) XXX_Size() int {
	return xxx_messageInfo_MsgSetChatFee.Size(m)
}
func (m *MsgSetChatFee) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetChatFee.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetChatFee proto.InternalMessageInfo

func (m *MsgSetChatFee) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgSetChatFee) GetFee() types.Coin {
	if m != nil {
		return m.Fee
	}
	return types.Coin{}
}

type MsgSendGift struct {
	FromAddress          string     `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	NodeAddress          string     `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
	ToAddress            string     `protobuf:"bytes,3,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty" yaml:"to_address"`
	GiftId               int64      `protobuf:"varint,4,opt,name=gift_id,json=giftId,proto3" json:"gift_id,omitempty" yaml:"gift_id"`
	GiftAmount           int64      `protobuf:"varint,5,opt,name=gift_amount,json=giftAmount,proto3" json:"gift_amount,omitempty" yaml:"gift_amount"`
	GiftValue            types.Coin `protobuf:"bytes,6,opt,name=gift_value,json=giftValue,proto3" json:"gift_value" yaml:"gift_value"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MsgSendGift) Reset()         { *m = MsgSendGift{} }
func (m *MsgSendGift) String() string { return proto.CompactTextString(m) }
func (*MsgSendGift) ProtoMessage()    {}
func (*MsgSendGift) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{4}
}
func (m *MsgSendGift) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSendGift.Unmarshal(m, b)
}
func (m *MsgSendGift) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSendGift.Marshal(b, m, deterministic)
}
func (m *MsgSendGift) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendGift.Merge(m, src)
}
func (m *MsgSendGift) XXX_Size() int {
	return xxx_messageInfo_MsgSendGift.Size(m)
}
func (m *MsgSendGift) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendGift.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendGift proto.InternalMessageInfo

func (m *MsgSendGift) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgSendGift) GetNodeAddress() string {
	if m != nil {
		return m.NodeAddress
	}
	return ""
}

func (m *MsgSendGift) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *MsgSendGift) GetGiftId() int64 {
	if m != nil {
		return m.GiftId
	}
	return 0
}

func (m *MsgSendGift) GetGiftAmount() int64 {
	if m != nil {
		return m.GiftAmount
	}
	return 0
}

func (m *MsgSendGift) GetGiftValue() types.Coin {
	if m != nil {
		return m.GiftValue
	}
	return types.Coin{}
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
	return fileDescriptor_0fd2153dc07d3b5c, []int{5}
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

type MsgTestResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgTestResponse) Reset()         { *m = MsgTestResponse{} }
func (m *MsgTestResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTestResponse) ProtoMessage()    {}
func (*MsgTestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{6}
}
func (m *MsgTestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgTestResponse.Unmarshal(m, b)
}
func (m *MsgTestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgTestResponse.Marshal(b, m, deterministic)
}
func (m *MsgTestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTestResponse.Merge(m, src)
}
func (m *MsgTestResponse) XXX_Size() int {
	return xxx_messageInfo_MsgTestResponse.Size(m)
}
func (m *MsgTestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTestResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgTest)(nil), "freemasonry.chat.v1.MsgTest")
	proto.RegisterType((*MsgRegister)(nil), "freemasonry.chat.v1.MsgRegister")
	proto.RegisterType((*MsgMortgage)(nil), "freemasonry.chat.v1.MsgMortgage")
	proto.RegisterType((*MsgSetChatFee)(nil), "freemasonry.chat.v1.MsgSetChatFee")
	proto.RegisterType((*MsgSendGift)(nil), "freemasonry.chat.v1.MsgSendGift")
	proto.RegisterType((*MsgEmptyResponse)(nil), "freemasonry.chat.v1.MsgEmptyResponse")
	proto.RegisterType((*MsgTestResponse)(nil), "freemasonry.chat.v1.MsgTestResponse")
}

func init() { proto.RegisterFile("tx.proto", fileDescriptor_0fd2153dc07d3b5c) }

var fileDescriptor_0fd2153dc07d3b5c = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x95, 0x3b, 0x6f, 0x13, 0x4b,
	0x14, 0xc7, 0xaf, 0x1f, 0xd7, 0x89, 0xc7, 0xf7, 0x26, 0x64, 0x03, 0xc1, 0xb1, 0xa2, 0xac, 0x19,
	0xf1, 0xb0, 0x84, 0xb4, 0x83, 0x13, 0x24, 0xa4, 0x34, 0x28, 0x8e, 0x20, 0xa2, 0x70, 0xb3, 0x46,
	0x20, 0xd1, 0x58, 0xe3, 0xdd, 0xe3, 0xc9, 0x0a, 0xef, 0x8c, 0x35, 0x33, 0x59, 0xc5, 0x2d, 0x35,
	0x1d, 0x5f, 0x8a, 0x1a, 0x7a, 0x57, 0x34, 0xb4, 0x2e, 0xa8, 0xd1, 0xcc, 0xee, 0x3a, 0x4e, 0x84,
	0x43, 0x10, 0x05, 0x12, 0xdd, 0x9c, 0xc7, 0xef, 0xbf, 0xe7, 0xcc, 0xe8, 0x9c, 0x45, 0xab, 0xfa,
	0xcc, 0x1b, 0x4b, 0xa1, 0x85, 0xb3, 0x39, 0x94, 0x00, 0x31, 0x55, 0x82, 0xcb, 0x89, 0x17, 0x9c,
	0x50, 0xed, 0x25, 0xed, 0xc6, 0x0e, 0x13, 0x82, 0x8d, 0x80, 0xd0, 0x71, 0x44, 0x28, 0xe7, 0x42,
	0x53, 0x1d, 0x09, 0xae, 0x52, 0xa4, 0x71, 0x93, 0x09, 0x26, 0xec, 0x91, 0x98, 0x53, 0xe6, 0xdd,
	0x0d, 0x84, 0x8a, 0x85, 0x22, 0x03, 0xaa, 0x80, 0x24, 0xed, 0x01, 0x68, 0xda, 0x26, 0x81, 0x88,
	0x78, 0x1a, 0xc7, 0x12, 0xad, 0x74, 0x15, 0x7b, 0x09, 0x4a, 0x3b, 0xfb, 0xa8, 0x6c, 0x02, 0xf5,
	0x42, 0xb3, 0xd0, 0xaa, 0xed, 0x6d, 0x7b, 0x29, 0xe9, 0x19, 0xd2, 0xcb, 0x48, 0xef, 0x48, 0x44,
	0xbc, 0x53, 0xfe, 0x38, 0x75, 0xff, 0xf1, 0x6d, 0xb2, 0xd3, 0x40, 0xab, 0x12, 0x02, 0x88, 0x12,
	0x90, 0xf5, 0x62, 0xb3, 0xd0, 0xaa, 0xfa, 0x73, 0xdb, 0xd9, 0x42, 0x15, 0x05, 0x3c, 0x04, 0x59,
	0x2f, 0xd9, 0x48, 0x66, 0xe1, 0xaf, 0x05, 0x54, 0xeb, 0x2a, 0xe6, 0x03, 0x8b, 0x94, 0x06, 0xe9,
	0x1c, 0xa0, 0xff, 0x86, 0x52, 0xc4, 0x7d, 0x1a, 0x86, 0x12, 0x94, 0xb2, 0x05, 0x54, 0x3b, 0xb7,
	0x67, 0x53, 0x77, 0x73, 0x42, 0xe3, 0xd1, 0x01, 0x5e, 0x8c, 0x62, 0xbf, 0x66, 0xcc, 0xc3, 0xd4,
	0x32, 0x2c, 0x17, 0x21, 0xcc, 0xd9, 0xe2, 0x65, 0x76, 0x31, 0x8a, 0xfd, 0x9a, 0x31, 0x73, 0x76,
	0x80, 0xd6, 0x63, 0x21, 0x35, 0xa3, 0x0c, 0xfa, 0x34, 0x16, 0xa7, 0x5c, 0xdb, 0x42, 0xaf, 0xec,
	0x7d, 0xd7, 0xf4, 0x3e, 0x9b, 0xba, 0x5b, 0xa9, 0xfa, 0x25, 0x1e, 0xfb, 0x6b, 0xb9, 0xe7, 0x30,
	0x75, 0x64, 0xbd, 0x76, 0x33, 0xef, 0x5f, 0xdd, 0xeb, 0xfb, 0x02, 0xfa, 0xbf, 0xab, 0x58, 0x0f,
	0xf4, 0xd1, 0x09, 0xd5, 0xcf, 0xe1, 0xf7, 0xba, 0x7d, 0x8a, 0x4a, 0x43, 0x00, 0xdb, 0xe4, 0x95,
	0x55, 0x3a, 0x59, 0x95, 0x28, 0x53, 0x04, 0xc0, 0xbe, 0x21, 0xf1, 0xb7, 0xa2, 0xbd, 0xfa, 0x1e,
	0xf0, 0xf0, 0x38, 0x1a, 0xea, 0x3f, 0x76, 0xf5, 0x8f, 0x11, 0xd2, 0x62, 0x4e, 0xda, 0x51, 0xe8,
	0xdc, 0x9a, 0x4d, 0xdd, 0x8d, 0x94, 0x3c, 0x8f, 0x61, 0xbf, 0xaa, 0x45, 0x4e, 0x3d, 0x44, 0x2b,
	0x2c, 0x1a, 0xea, 0x7e, 0x14, 0xd6, 0xcb, 0xcd, 0x42, 0xab, 0xd4, 0x71, 0x66, 0x53, 0x77, 0x2d,
	0x45, 0xb2, 0x00, 0xf6, 0x2b, 0xe6, 0xf4, 0x22, 0x74, 0x9e, 0xa0, 0x9a, 0xf5, 0x65, 0x2f, 0xfb,
	0xaf, 0x05, 0xb6, 0x66, 0x53, 0xd7, 0x59, 0x00, 0xf2, 0x67, 0x43, 0xc6, 0x4a, 0x9f, 0xcc, 0xe9,
	0x21, 0x6b, 0xf5, 0x13, 0x3a, 0x3a, 0x85, 0x7a, 0xe5, 0x67, 0x77, 0xbd, 0x9d, 0xdd, 0xf5, 0xc6,
	0x82, 0xac, 0x45, 0xb1, 0x5f, 0x35, 0xc6, 0x2b, 0x7b, 0x76, 0xd0, 0x8d, 0xae, 0x62, 0xcf, 0xe2,
	0xb1, 0x9e, 0xf8, 0xa0, 0xc6, 0x82, 0x2b, 0xc0, 0x1b, 0x68, 0x3d, 0xdb, 0x33, 0xb9, 0x6b, 0xef,
	0x53, 0x09, 0x95, 0xba, 0x8a, 0x39, 0x0a, 0x95, 0xed, 0xfe, 0xd9, 0xf1, 0x7e, 0xb0, 0xf4, 0xbc,
	0x8c, 0x6a, 0xdc, 0xbd, 0x2a, 0x3a, 0xff, 0xcc, 0x83, 0x77, 0x9f, 0xbf, 0x7c, 0x28, 0xde, 0x71,
	0x5c, 0x02, 0x89, 0x59, 0x7b, 0x20, 0x83, 0xbd, 0x47, 0x24, 0x69, 0x13, 0x7d, 0x46, 0x02, 0xc1,
	0x13, 0x90, 0xba, 0x6f, 0xf7, 0x56, 0x0f, 0xad, 0xce, 0xf7, 0x4f, 0x73, 0x99, 0x74, 0x9e, 0xd1,
	0xb8, 0xb7, 0x2c, 0xe3, 0x42, 0x93, 0x46, 0xd4, 0x0c, 0xfa, 0xb1, 0x19, 0xf4, 0xa5, 0xa2, 0xf9,
	0x2a, 0xb8, 0xae, 0xe8, 0x6b, 0x84, 0x16, 0x26, 0x0a, 0x2f, 0x83, 0xce, 0x73, 0x7e, 0xa1, 0xda,
	0xf9, 0x6c, 0x34, 0x97, 0xcb, 0xa6, 0x19, 0xd7, 0x14, 0xed, 0xb4, 0xde, 0xdc, 0xbf, 0x90, 0x17,
	0x90, 0xc1, 0x48, 0x04, 0x6f, 0x83, 0x13, 0x1a, 0x71, 0x72, 0x46, 0x0c, 0x47, 0xf4, 0x64, 0x0c,
	0x6a, 0x50, 0xb1, 0x3f, 0xa0, 0xfd, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x18, 0xef, 0x9c,
	0xf5, 0x06, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type MsgClient interface {


	Test(ctx context.Context, in *MsgTest, opts ...grpc.CallOption) (*MsgTestResponse, error)
	Register(ctx context.Context, in *MsgRegister, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	MortGage(ctx context.Context, in *MsgMortgage, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	SetChatFee(ctx context.Context, in *MsgSetChatFee, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	SendGift(ctx context.Context, in *MsgSendGift, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
}

type msgClient struct {
	cc *grpc.ClientConn
}

func NewMsgClient(cc *grpc.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Test(ctx context.Context, in *MsgTest, opts ...grpc.CallOption) (*MsgTestResponse, error) {
	out := new(MsgTestResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/Test", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Register(ctx context.Context, in *MsgRegister, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MortGage(ctx context.Context, in *MsgMortgage, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/MortGage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetChatFee(ctx context.Context, in *MsgSetChatFee, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/SetChatFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendGift(ctx context.Context, in *MsgSendGift, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/SendGift", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type MsgServer interface {


	Test(context.Context, *MsgTest) (*MsgTestResponse, error)
	Register(context.Context, *MsgRegister) (*MsgEmptyResponse, error)
	MortGage(context.Context, *MsgMortgage) (*MsgEmptyResponse, error)
	SetChatFee(context.Context, *MsgSetChatFee) (*MsgEmptyResponse, error)
	SendGift(context.Context, *MsgSendGift) (*MsgEmptyResponse, error)
}


type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Test(ctx context.Context, req *MsgTest) (*MsgTestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (*UnimplementedMsgServer) Register(ctx context.Context, req *MsgRegister) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedMsgServer) MortGage(ctx context.Context, req *MsgMortgage) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MortGage not implemented")
}
func (*UnimplementedMsgServer) SetChatFee(ctx context.Context, req *MsgSetChatFee) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetChatFee not implemented")
}
func (*UnimplementedMsgServer) SendGift(ctx context.Context, req *MsgSendGift) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendGift not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Msg/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Test(ctx, req.(*MsgTest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegister)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Msg/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Register(ctx, req.(*MsgRegister))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MortGage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMortgage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MortGage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Msg/MortGage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MortGage(ctx, req.(*MsgMortgage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetChatFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetChatFee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetChatFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Msg/SetChatFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetChatFee(ctx, req.(*MsgSetChatFee))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendGift_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendGift)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendGift(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Msg/SendGift",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendGift(ctx, req.(*MsgSendGift))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "freemasonry.chat.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _Msg_Test_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Msg_Register_Handler,
		},
		{
			MethodName: "MortGage",
			Handler:    _Msg_MortGage_Handler,
		},
		{
			MethodName: "SetChatFee",
			Handler:    _Msg_SetChatFee_Handler,
		},
		{
			MethodName: "SendGift",
			Handler:    _Msg_SendGift_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tx.proto",
}

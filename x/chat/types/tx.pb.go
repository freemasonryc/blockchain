


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

type MsgRegister struct {
	FromAddress          string     `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	NodeAddress          string     `protobuf:"bytes,2,opt,name=node_address,json=nodeAddress,proto3" json:"node_address,omitempty" yaml:"node_address"`
	MortgageAmount       types.Coin `protobuf:"bytes,3,opt,name=mortgage_amount,json=mortgageAmount,proto3" json:"mortgage_amount" yaml:"mortgage_amount"`
	MobilePrefix         string     `protobuf:"bytes,4,opt,name=mobile_prefix,json=mobilePrefix,proto3" json:"mobile_prefix,omitempty" yaml:"mobile_prefix"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MsgRegister) Reset()         { *m = MsgRegister{} }
func (m *MsgRegister) String() string { return proto.CompactTextString(m) }
func (*MsgRegister) ProtoMessage()    {}
func (*MsgRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{0}
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

func (m *MsgRegister) GetMobilePrefix() string {
	if m != nil {
		return m.MobilePrefix
	}
	return ""
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
	return fileDescriptor_0fd2153dc07d3b5c, []int{1}
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
	return fileDescriptor_0fd2153dc07d3b5c, []int{2}
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
	return fileDescriptor_0fd2153dc07d3b5c, []int{3}
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

type MsgAddressBookSave struct {
	FromAddress          string   `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	AddressBook          []string `protobuf:"bytes,2,rep,name=address_book,json=addressBook,proto3" json:"address_book,omitempty" yaml:"address_book"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgAddressBookSave) Reset()         { *m = MsgAddressBookSave{} }
func (m *MsgAddressBookSave) String() string { return proto.CompactTextString(m) }
func (*MsgAddressBookSave) ProtoMessage()    {}
func (*MsgAddressBookSave) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{4}
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

type MsgGetRewards struct {
	FromAddress          string   `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgGetRewards) Reset()         { *m = MsgGetRewards{} }
func (m *MsgGetRewards) String() string { return proto.CompactTextString(m) }
func (*MsgGetRewards) ProtoMessage()    {}
func (*MsgGetRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{5}
}
func (m *MsgGetRewards) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgGetRewards.Unmarshal(m, b)
}
func (m *MsgGetRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgGetRewards.Marshal(b, m, deterministic)
}
func (m *MsgGetRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGetRewards.Merge(m, src)
}
func (m *MsgGetRewards) XXX_Size() int {
	return xxx_messageInfo_MsgGetRewards.Size(m)
}
func (m *MsgGetRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGetRewards.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGetRewards proto.InternalMessageInfo

func (m *MsgGetRewards) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

type MsgMobileTransfer struct {
	FromAddress          string   `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	ToAddress            string   `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty" yaml:"to_address"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty" yaml:"mobile"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgMobileTransfer) Reset()         { *m = MsgMobileTransfer{} }
func (m *MsgMobileTransfer) String() string { return proto.CompactTextString(m) }
func (*MsgMobileTransfer) ProtoMessage()    {}
func (*MsgMobileTransfer) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{6}
}
func (m *MsgMobileTransfer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgMobileTransfer.Unmarshal(m, b)
}
func (m *MsgMobileTransfer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgMobileTransfer.Marshal(b, m, deterministic)
}
func (m *MsgMobileTransfer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMobileTransfer.Merge(m, src)
}
func (m *MsgMobileTransfer) XXX_Size() int {
	return xxx_messageInfo_MsgMobileTransfer.Size(m)
}
func (m *MsgMobileTransfer) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMobileTransfer.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMobileTransfer proto.InternalMessageInfo

func (m *MsgMobileTransfer) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgMobileTransfer) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *MsgMobileTransfer) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
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
	return fileDescriptor_0fd2153dc07d3b5c, []int{7}
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
	return fileDescriptor_0fd2153dc07d3b5c, []int{8}
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
	proto.RegisterType((*MsgRegister)(nil), "freemasonry.chat.v1.MsgRegister")
	proto.RegisterType((*MsgMortgage)(nil), "freemasonry.chat.v1.MsgMortgage")
	proto.RegisterType((*MsgSetChatFee)(nil), "freemasonry.chat.v1.MsgSetChatFee")
	proto.RegisterType((*MsgSendGift)(nil), "freemasonry.chat.v1.MsgSendGift")
	proto.RegisterType((*MsgAddressBookSave)(nil), "freemasonry.chat.v1.MsgAddressBookSave")
	proto.RegisterType((*MsgGetRewards)(nil), "freemasonry.chat.v1.MsgGetRewards")
	proto.RegisterType((*MsgMobileTransfer)(nil), "freemasonry.chat.v1.MsgMobileTransfer")
	proto.RegisterType((*MsgEmptyResponse)(nil), "freemasonry.chat.v1.MsgEmptyResponse")
	proto.RegisterType((*MsgTestResponse)(nil), "freemasonry.chat.v1.MsgTestResponse")
}

func init() { proto.RegisterFile("tx.proto", fileDescriptor_0fd2153dc07d3b5c) }

var fileDescriptor_0fd2153dc07d3b5c = []byte{

	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0x4f, 0x6f, 0xd3, 0x3e,
	0x18, 0xfe, 0xb5, 0xdd, 0xaf, 0x6c, 0xee, 0xfe, 0xd0, 0x6c, 0x8c, 0x6e, 0x42, 0x4b, 0x65, 0x89,
	0x31, 0x84, 0x94, 0x68, 0x03, 0x09, 0x69, 0x12, 0x42, 0xeb, 0x04, 0x13, 0x42, 0x91, 0x50, 0x3a,
	0x81, 0xc4, 0x25, 0x72, 0x1a, 0xc7, 0x8b, 0xd6, 0xc4, 0x55, 0xec, 0x95, 0xf5, 0x3b, 0xec, 0x83,
	0x70, 0xe4, 0x63, 0xf0, 0x29, 0x72, 0xe7, 0x9a, 0x03, 0x07, 0x4e, 0xc8, 0x8e, 0xd3, 0xa6, 0x15,
	0x81, 0x4e, 0x3d, 0x20, 0x71, 0xf3, 0xeb, 0xf7, 0x7d, 0x9e, 0xf8, 0x7d, 0xfd, 0xf8, 0x69, 0xc1,
	0x32, 0xbf, 0x36, 0x06, 0x31, 0xe5, 0x54, 0xdb, 0xf4, 0x63, 0x8c, 0x43, 0xc4, 0x68, 0x14, 0x8f,
	0x8c, 0xde, 0x05, 0xe2, 0xc6, 0xf0, 0x70, 0xf7, 0x01, 0xa1, 0x94, 0xf4, 0xb1, 0x89, 0x06, 0x81,
	0x89, 0xa2, 0x88, 0x72, 0xc4, 0x03, 0x1a, 0xb1, 0x0c, 0xb2, 0xbb, 0x45, 0x28, 0xa1, 0x72, 0x69,
	0x8a, 0x95, 0xda, 0xdd, 0xeb, 0x51, 0x16, 0x52, 0x66, 0xba, 0x88, 0x61, 0x73, 0x78, 0xe8, 0x62,
	0x8e, 0x0e, 0xcd, 0x1e, 0x0d, 0xa2, 0x2c, 0x0f, 0x3f, 0x57, 0x41, 0xc3, 0x62, 0xc4, 0xc6, 0x24,
	0x60, 0x1c, 0xc7, 0xda, 0x31, 0x58, 0xf5, 0x63, 0x1a, 0x3a, 0xc8, 0xf3, 0x62, 0xcc, 0x58, 0xab,
	0xd2, 0xae, 0x1c, 0xac, 0x74, 0xee, 0xa7, 0x89, 0xbe, 0x39, 0x42, 0x61, 0xff, 0x18, 0x16, 0xb3,
	0xd0, 0x6e, 0x88, 0xf0, 0x24, 0x8b, 0x04, 0x36, 0xa2, 0x1e, 0x1e, 0x63, 0xab, 0xb3, 0xd8, 0x62,
	0x16, 0xda, 0x0d, 0x11, 0xe6, 0x58, 0x17, 0x6c, 0x84, 0x34, 0xe6, 0x04, 0x11, 0xec, 0xa0, 0x90,
	0x5e, 0x45, 0xbc, 0x55, 0x6b, 0x57, 0x0e, 0x1a, 0x47, 0x3b, 0x46, 0xd6, 0x81, 0x21, 0x3a, 0x30,
	0x54, 0x07, 0xc6, 0x29, 0x0d, 0xa2, 0xce, 0xde, 0xd7, 0x44, 0xff, 0x2f, 0x4d, 0xf4, 0xed, 0x8c,
	0x7d, 0x06, 0x0f, 0xed, 0xf5, 0x7c, 0xe7, 0x44, 0x6e, 0x68, 0x2f, 0xc0, 0x5a, 0x48, 0xdd, 0xa0,
	0x8f, 0x9d, 0x41, 0x8c, 0xfd, 0xe0, 0xba, 0xb5, 0x24, 0x0f, 0xd8, 0x4a, 0x13, 0x7d, 0x2b, 0xa7,
	0x28, 0xa4, 0xa1, 0xbd, 0x9a, 0xc5, 0xef, 0xb2, 0xf0, 0x5b, 0x45, 0x8e, 0xca, 0x52, 0xa4, 0xff,
	0xf2, 0xa8, 0xe0, 0x4d, 0x05, 0xac, 0x59, 0x8c, 0x74, 0x31, 0x3f, 0xbd, 0x40, 0xfc, 0x35, 0x5e,
	0xac, 0xdb, 0x97, 0xa0, 0xe6, 0x63, 0x2c, 0x9b, 0xfc, 0xed, 0x29, 0x35, 0x75, 0x4a, 0xa0, 0x18,
	0x31, 0x86, 0xb6, 0x40, 0xc2, 0xef, 0x99, 0x4a, 0xbb, 0x38, 0xf2, 0xce, 0x02, 0x9f, 0xff, 0xb5,
	0xd1, 0x3f, 0x03, 0x80, 0xd3, 0x31, 0xb2, 0x26, 0x91, 0xf7, 0xd2, 0x44, 0x6f, 0x66, 0xc8, 0x49,
	0x0e, 0xda, 0x2b, 0x9c, 0xe6, 0xa8, 0x27, 0xe0, 0x0e, 0x09, 0x7c, 0xee, 0x04, 0x9e, 0x54, 0x5c,
	0xad, 0xa3, 0xa5, 0x89, 0xbe, 0x9e, 0x41, 0x54, 0x02, 0xda, 0x75, 0xb1, 0x7a, 0xe3, 0x69, 0xcf,
	0x41, 0x43, 0xee, 0xa9, 0x9b, 0xfd, 0x5f, 0x02, 0xb6, 0xd3, 0x44, 0xd7, 0x0a, 0x80, 0xfc, 0xda,
	0x80, 0x88, 0x94, 0xba, 0xbb, 0x40, 0x46, 0xce, 0x10, 0xf5, 0xaf, 0x70, 0xab, 0xfe, 0xa7, 0x59,
	0xef, 0xa8, 0x59, 0x37, 0x0b, 0xb4, 0x12, 0x0a, 0xed, 0x15, 0x11, 0xbc, 0x97, 0xeb, 0x9b, 0x0a,
	0xd0, 0x2c, 0x46, 0x54, 0x27, 0x1d, 0x4a, 0x2f, 0xbb, 0x68, 0xb8, 0xb0, 0xf4, 0x55, 0xc2, 0x71,
	0x29, 0xbd, 0x6c, 0x55, 0xdb, 0xb5, 0x69, 0x6c, 0x31, 0x0b, 0xed, 0x06, 0x9a, 0x7c, 0x1b, 0xbe,
	0x95, 0xaa, 0x3c, 0xc3, 0xdc, 0xc6, 0x9f, 0x50, 0xec, 0xb1, 0x45, 0x0e, 0x02, 0xbf, 0x54, 0x40,
	0x53, 0xbe, 0x67, 0xf1, 0xc6, 0xcf, 0x63, 0x14, 0x31, 0x7f, 0x41, 0x03, 0x9c, 0x96, 0x47, 0x75,
	0x4e, 0x79, 0x3c, 0x06, 0xf5, 0xcc, 0x67, 0x94, 0xa0, 0x9a, 0x69, 0xa2, 0xaf, 0x15, 0xfd, 0x08,
	0xda, 0xaa, 0x00, 0x6a, 0xe0, 0xae, 0xc5, 0xc8, 0xab, 0x70, 0xc0, 0x47, 0x36, 0x66, 0x03, 0x1a,
	0x31, 0x0c, 0x9b, 0x60, 0xc3, 0x62, 0xe4, 0x1c, 0x33, 0x9e, 0x6f, 0x1d, 0xfd, 0x58, 0x02, 0x35,
	0x8b, 0x11, 0xad, 0x0b, 0x96, 0xc7, 0xc6, 0xde, 0x36, 0x7e, 0xf1, 0x93, 0x62, 0x14, 0xac, 0x7f,
	0xf7, 0x61, 0x59, 0xc5, 0xd4, 0xf7, 0x04, 0xa9, 0xb0, 0xc0, 0x33, 0x61, 0x81, 0xa5, 0xa4, 0xb9,
	0x49, 0xce, 0x4b, 0xfa, 0x01, 0x80, 0x82, 0xd7, 0xc0, 0x32, 0xd0, 0xa4, 0xe6, 0x16, 0xa7, 0x1d,
	0xbb, 0x46, 0xbb, 0x9c, 0x36, 0xab, 0x98, 0x97, 0x14, 0x81, 0x8d, 0xd9, 0x17, 0xf1, 0xa8, 0x0c,
	0x39, 0x53, 0x78, 0x8b, 0x81, 0x14, 0x64, 0x5e, 0x3a, 0x90, 0x49, 0xcd, 0xbc, 0xc4, 0x0e, 0x58,
	0x9f, 0x51, 0xfc, 0x7e, 0xf9, 0x25, 0x16, 0xeb, 0xe6, 0xfc, 0x40, 0xe7, 0xe0, 0xe3, 0xfe, 0x54,
	0x5d, 0xcf, 0x74, 0xfb, 0xb4, 0x77, 0xd9, 0xbb, 0x40, 0x41, 0x64, 0x5e, 0x9b, 0x02, 0x67, 0xf2,
	0xd1, 0x00, 0x33, 0xb7, 0x2e, 0xff, 0x82, 0x3c, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0x20, 0x03,
	0x91, 0xa5, 0xf7, 0x08, 0x00, 0x00,
}


var _ context.Context
var _ grpc.ClientConn



const _ = grpc.SupportPackageIsVersion4


//

type MsgClient interface {
	Register(ctx context.Context, in *MsgRegister, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	MortGage(ctx context.Context, in *MsgMortgage, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	SetChatFee(ctx context.Context, in *MsgSetChatFee, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	SendGift(ctx context.Context, in *MsgSendGift, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	AddressBookSave(ctx context.Context, in *MsgAddressBookSave, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	GetRewards(ctx context.Context, in *MsgGetRewards, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
	MobileTransfer(ctx context.Context, in *MsgMobileTransfer, opts ...grpc.CallOption) (*MsgEmptyResponse, error)
}

type msgClient struct {
	cc *grpc.ClientConn
}

func NewMsgClient(cc *grpc.ClientConn) MsgClient {
	return &msgClient{cc}
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

func (c *msgClient) AddressBookSave(ctx context.Context, in *MsgAddressBookSave, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/AddressBookSave", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) GetRewards(ctx context.Context, in *MsgGetRewards, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/GetRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MobileTransfer(ctx context.Context, in *MsgMobileTransfer, opts ...grpc.CallOption) (*MsgEmptyResponse, error) {
	out := new(MsgEmptyResponse)
	err := c.cc.Invoke(ctx, "/freemasonry.chat.v1.Msg/MobileTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


type MsgServer interface {
	Register(context.Context, *MsgRegister) (*MsgEmptyResponse, error)
	MortGage(context.Context, *MsgMortgage) (*MsgEmptyResponse, error)
	SetChatFee(context.Context, *MsgSetChatFee) (*MsgEmptyResponse, error)
	SendGift(context.Context, *MsgSendGift) (*MsgEmptyResponse, error)
	AddressBookSave(context.Context, *MsgAddressBookSave) (*MsgEmptyResponse, error)
	GetRewards(context.Context, *MsgGetRewards) (*MsgEmptyResponse, error)
	MobileTransfer(context.Context, *MsgMobileTransfer) (*MsgEmptyResponse, error)
}


type UnimplementedMsgServer struct {
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
func (*UnimplementedMsgServer) AddressBookSave(ctx context.Context, req *MsgAddressBookSave) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddressBookSave not implemented")
}
func (*UnimplementedMsgServer) GetRewards(ctx context.Context, req *MsgGetRewards) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRewards not implemented")
}
func (*UnimplementedMsgServer) MobileTransfer(ctx context.Context, req *MsgMobileTransfer) (*MsgEmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MobileTransfer not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
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
		FullMethod: "/freemasonry.chat.v1.Msg/AddressBookSave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddressBookSave(ctx, req.(*MsgAddressBookSave))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_GetRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGetRewards)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GetRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Msg/GetRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GetRewards(ctx, req.(*MsgGetRewards))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MobileTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMobileTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MobileTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/freemasonry.chat.v1.Msg/MobileTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MobileTransfer(ctx, req.(*MsgMobileTransfer))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "freemasonry.chat.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
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
		{
			MethodName: "AddressBookSave",
			Handler:    _Msg_AddressBookSave_Handler,
		},
		{
			MethodName: "GetRewards",
			Handler:    _Msg_GetRewards_Handler,
		},
		{
			MethodName: "MobileTransfer",
			Handler:    _Msg_MobileTransfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tx.proto",
}

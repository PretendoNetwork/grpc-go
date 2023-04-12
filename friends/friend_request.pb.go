// Code generated by protoc-gen-go. DO NOT EDIT.
// source: friend_request.proto

package grpc_friends

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FriendRequest struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Sender               uint32   `protobuf:"varint,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Recipient            uint32   `protobuf:"varint,3,opt,name=recipient,proto3" json:"recipient,omitempty"`
	Sent                 uint64   `protobuf:"varint,4,opt,name=sent,proto3" json:"sent,omitempty"`
	Expires              uint64   `protobuf:"varint,5,opt,name=expires,proto3" json:"expires,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FriendRequest) Reset()         { *m = FriendRequest{} }
func (m *FriendRequest) String() string { return proto.CompactTextString(m) }
func (*FriendRequest) ProtoMessage()    {}
func (*FriendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7112050df3d0ca77, []int{0}
}

func (m *FriendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FriendRequest.Unmarshal(m, b)
}
func (m *FriendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FriendRequest.Marshal(b, m, deterministic)
}
func (m *FriendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FriendRequest.Merge(m, src)
}
func (m *FriendRequest) XXX_Size() int {
	return xxx_messageInfo_FriendRequest.Size(m)
}
func (m *FriendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FriendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FriendRequest proto.InternalMessageInfo

func (m *FriendRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *FriendRequest) GetSender() uint32 {
	if m != nil {
		return m.Sender
	}
	return 0
}

func (m *FriendRequest) GetRecipient() uint32 {
	if m != nil {
		return m.Recipient
	}
	return 0
}

func (m *FriendRequest) GetSent() uint64 {
	if m != nil {
		return m.Sent
	}
	return 0
}

func (m *FriendRequest) GetExpires() uint64 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func init() {
	proto.RegisterType((*FriendRequest)(nil), "friends.FriendRequest")
}

func init() {
	proto.RegisterFile("friend_request.proto", fileDescriptor_7112050df3d0ca77)
}

var fileDescriptor_7112050df3d0ca77 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2b, 0xca, 0x4c,
	0xcd, 0x4b, 0x89, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x87, 0x88, 0x16, 0x2b, 0x35, 0x33, 0x72, 0xf1, 0xba, 0x81, 0xd9, 0x41, 0x10, 0x05,
	0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x4c, 0x99, 0x29,
	0x42, 0x62, 0x5c, 0x6c, 0xc5, 0xa9, 0x79, 0x29, 0xa9, 0x45, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xbc,
	0x41, 0x50, 0x9e, 0x90, 0x0c, 0x17, 0x67, 0x51, 0x6a, 0x72, 0x66, 0x41, 0x66, 0x6a, 0x5e, 0x89,
	0x04, 0x33, 0x58, 0x0a, 0x21, 0x20, 0x24, 0xc4, 0xc5, 0x52, 0x0c, 0x92, 0x60, 0x01, 0x9b, 0x03,
	0x66, 0x0b, 0x49, 0x70, 0xb1, 0xa7, 0x56, 0x14, 0x64, 0x16, 0xa5, 0x16, 0x4b, 0xb0, 0x82, 0x85,
	0x61, 0x5c, 0x27, 0xc1, 0x28, 0x7e, 0x3d, 0x7d, 0xeb, 0xf4, 0xa2, 0x82, 0xe4, 0x78, 0xa8, 0xc3,
	0x92, 0xd8, 0xc0, 0x0e, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x90, 0xb5, 0x7c, 0xb7, 0xc0,
	0x00, 0x00, 0x00,
}
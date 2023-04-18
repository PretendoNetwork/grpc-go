// Code generated by protoc-gen-go. DO NOT EDIT.
// source: deny_friend_request_rpc.proto

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

type DenyFriendRequestRequest struct {
	FriendRequestId      uint64   `protobuf:"varint,1,opt,name=friend_request_id,json=friendRequestId,proto3" json:"friend_request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DenyFriendRequestRequest) Reset()         { *m = DenyFriendRequestRequest{} }
func (m *DenyFriendRequestRequest) String() string { return proto.CompactTextString(m) }
func (*DenyFriendRequestRequest) ProtoMessage()    {}
func (*DenyFriendRequestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_104dd7fb22e5209e, []int{0}
}

func (m *DenyFriendRequestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DenyFriendRequestRequest.Unmarshal(m, b)
}
func (m *DenyFriendRequestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DenyFriendRequestRequest.Marshal(b, m, deterministic)
}
func (m *DenyFriendRequestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenyFriendRequestRequest.Merge(m, src)
}
func (m *DenyFriendRequestRequest) XXX_Size() int {
	return xxx_messageInfo_DenyFriendRequestRequest.Size(m)
}
func (m *DenyFriendRequestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DenyFriendRequestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DenyFriendRequestRequest proto.InternalMessageInfo

func (m *DenyFriendRequestRequest) GetFriendRequestId() uint64 {
	if m != nil {
		return m.FriendRequestId
	}
	return 0
}

type DenyFriendRequestResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DenyFriendRequestResponse) Reset()         { *m = DenyFriendRequestResponse{} }
func (m *DenyFriendRequestResponse) String() string { return proto.CompactTextString(m) }
func (*DenyFriendRequestResponse) ProtoMessage()    {}
func (*DenyFriendRequestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_104dd7fb22e5209e, []int{1}
}

func (m *DenyFriendRequestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DenyFriendRequestResponse.Unmarshal(m, b)
}
func (m *DenyFriendRequestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DenyFriendRequestResponse.Marshal(b, m, deterministic)
}
func (m *DenyFriendRequestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenyFriendRequestResponse.Merge(m, src)
}
func (m *DenyFriendRequestResponse) XXX_Size() int {
	return xxx_messageInfo_DenyFriendRequestResponse.Size(m)
}
func (m *DenyFriendRequestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DenyFriendRequestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DenyFriendRequestResponse proto.InternalMessageInfo

func (m *DenyFriendRequestResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*DenyFriendRequestRequest)(nil), "friends.DenyFriendRequestRequest")
	proto.RegisterType((*DenyFriendRequestResponse)(nil), "friends.DenyFriendRequestResponse")
}

func init() {
	proto.RegisterFile("deny_friend_request_rpc.proto", fileDescriptor_104dd7fb22e5209e)
}

var fileDescriptor_104dd7fb22e5209e = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x49, 0xcd, 0xab,
	0x8c, 0x4f, 0x2b, 0xca, 0x4c, 0xcd, 0x4b, 0x89, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x89,
	0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0xc8, 0x14, 0x2b, 0xb9,
	0x71, 0x49, 0xb8, 0xa4, 0xe6, 0x55, 0xba, 0x81, 0xb9, 0x41, 0x10, 0x75, 0x50, 0x4a, 0x48, 0x8b,
	0x4b, 0x10, 0xcd, 0x80, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0xfe, 0x34, 0x64,
	0x0d, 0x9e, 0x29, 0x4a, 0xa6, 0x5c, 0x92, 0x58, 0xcc, 0x29, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15,
	0x92, 0xe0, 0x62, 0x2f, 0x2e, 0x4d, 0x4e, 0x4e, 0x2d, 0x2e, 0x06, 0x6b, 0xe7, 0x08, 0x82, 0x71,
	0x9d, 0x04, 0xa3, 0xf8, 0xf5, 0xf4, 0xad, 0xd3, 0x8b, 0x0a, 0x92, 0xa1, 0x6e, 0x2d, 0x4e, 0x62,
	0x03, 0xbb, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xf4, 0xb2, 0x00, 0xc2, 0x00, 0x00,
	0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: send_user_notification_wiiu_rpc.proto

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

type SendUserNotificationWiiURequest struct {
	Pid                  uint32   `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	NotificationData     []byte   `protobuf:"bytes,2,opt,name=notification_data,json=notificationData,proto3" json:"notification_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendUserNotificationWiiURequest) Reset()         { *m = SendUserNotificationWiiURequest{} }
func (m *SendUserNotificationWiiURequest) String() string { return proto.CompactTextString(m) }
func (*SendUserNotificationWiiURequest) ProtoMessage()    {}
func (*SendUserNotificationWiiURequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4bc08365ab26817c, []int{0}
}

func (m *SendUserNotificationWiiURequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendUserNotificationWiiURequest.Unmarshal(m, b)
}
func (m *SendUserNotificationWiiURequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendUserNotificationWiiURequest.Marshal(b, m, deterministic)
}
func (m *SendUserNotificationWiiURequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendUserNotificationWiiURequest.Merge(m, src)
}
func (m *SendUserNotificationWiiURequest) XXX_Size() int {
	return xxx_messageInfo_SendUserNotificationWiiURequest.Size(m)
}
func (m *SendUserNotificationWiiURequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendUserNotificationWiiURequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendUserNotificationWiiURequest proto.InternalMessageInfo

func (m *SendUserNotificationWiiURequest) GetPid() uint32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *SendUserNotificationWiiURequest) GetNotificationData() []byte {
	if m != nil {
		return m.NotificationData
	}
	return nil
}

func init() {
	proto.RegisterType((*SendUserNotificationWiiURequest)(nil), "friends.SendUserNotificationWiiURequest")
}

func init() {
	proto.RegisterFile("send_user_notification_wiiu_rpc.proto", fileDescriptor_4bc08365ab26817c)
}

var fileDescriptor_4bc08365ab26817c = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x4e, 0xcd, 0x4b,
	0x89, 0x2f, 0x2d, 0x4e, 0x2d, 0x8a, 0xcf, 0xcb, 0x2f, 0xc9, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0xc9,
	0xcc, 0xcf, 0x8b, 0x2f, 0xcf, 0xcc, 0x2c, 0x8d, 0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x4f, 0x2b, 0xca, 0x4c, 0xcd, 0x4b, 0x29, 0x56, 0x4a, 0xe0, 0x92, 0x0f, 0x4e,
	0xcd, 0x4b, 0x09, 0x2d, 0x4e, 0x2d, 0xf2, 0x43, 0x52, 0x1f, 0x9e, 0x99, 0x19, 0x1a, 0x94, 0x5a,
	0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc0, 0xc5, 0x5c, 0x90, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x1b, 0x04, 0x62, 0x0a, 0x69, 0x73, 0x09, 0xa2, 0x18, 0x9e, 0x92, 0x58, 0x92, 0x28, 0xc1,
	0xa4, 0xc0, 0xa8, 0xc1, 0x13, 0x24, 0x80, 0x2c, 0xe1, 0x92, 0x58, 0x92, 0xe8, 0x24, 0x18, 0xc5,
	0xaf, 0xa7, 0x6f, 0x9d, 0x5e, 0x54, 0x90, 0x1c, 0x0f, 0xb5, 0x34, 0x89, 0x0d, 0xec, 0x08, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0x25, 0xb6, 0x40, 0xad, 0x00, 0x00, 0x00,
}

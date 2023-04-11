// Code generated by protoc-gen-go. DO NOT EDIT.
// source: smm_request_post_id_rpc.proto

package grpc_miiverse

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

type SMMRequestPostIdRequest struct {
	CourseId             uint64   `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SMMRequestPostIdRequest) Reset()         { *m = SMMRequestPostIdRequest{} }
func (m *SMMRequestPostIdRequest) String() string { return proto.CompactTextString(m) }
func (*SMMRequestPostIdRequest) ProtoMessage()    {}
func (*SMMRequestPostIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_02d24320ffde94aa, []int{0}
}

func (m *SMMRequestPostIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SMMRequestPostIdRequest.Unmarshal(m, b)
}
func (m *SMMRequestPostIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SMMRequestPostIdRequest.Marshal(b, m, deterministic)
}
func (m *SMMRequestPostIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SMMRequestPostIdRequest.Merge(m, src)
}
func (m *SMMRequestPostIdRequest) XXX_Size() int {
	return xxx_messageInfo_SMMRequestPostIdRequest.Size(m)
}
func (m *SMMRequestPostIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SMMRequestPostIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SMMRequestPostIdRequest proto.InternalMessageInfo

func (m *SMMRequestPostIdRequest) GetCourseId() uint64 {
	if m != nil {
		return m.CourseId
	}
	return 0
}

type SMMRequestPostIdResponse struct {
	PostId               uint64   `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SMMRequestPostIdResponse) Reset()         { *m = SMMRequestPostIdResponse{} }
func (m *SMMRequestPostIdResponse) String() string { return proto.CompactTextString(m) }
func (*SMMRequestPostIdResponse) ProtoMessage()    {}
func (*SMMRequestPostIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_02d24320ffde94aa, []int{1}
}

func (m *SMMRequestPostIdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SMMRequestPostIdResponse.Unmarshal(m, b)
}
func (m *SMMRequestPostIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SMMRequestPostIdResponse.Marshal(b, m, deterministic)
}
func (m *SMMRequestPostIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SMMRequestPostIdResponse.Merge(m, src)
}
func (m *SMMRequestPostIdResponse) XXX_Size() int {
	return xxx_messageInfo_SMMRequestPostIdResponse.Size(m)
}
func (m *SMMRequestPostIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SMMRequestPostIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SMMRequestPostIdResponse proto.InternalMessageInfo

func (m *SMMRequestPostIdResponse) GetPostId() uint64 {
	if m != nil {
		return m.PostId
	}
	return 0
}

func init() {
	proto.RegisterType((*SMMRequestPostIdRequest)(nil), "miiverse.SMMRequestPostIdRequest")
	proto.RegisterType((*SMMRequestPostIdResponse)(nil), "miiverse.SMMRequestPostIdResponse")
}

func init() {
	proto.RegisterFile("smm_request_post_id_rpc.proto", fileDescriptor_02d24320ffde94aa)
}

var fileDescriptor_02d24320ffde94aa = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0xce, 0xcd, 0x8d,
	0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x89, 0x2f, 0xc8, 0x2f, 0x2e, 0x89, 0xcf, 0x4c, 0x89,
	0x2f, 0x2a, 0x48, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0xcd, 0xcc, 0x2c, 0x4b,
	0x2d, 0x2a, 0x4e, 0x55, 0x32, 0xe3, 0x12, 0x0f, 0xf6, 0xf5, 0x0d, 0x82, 0xa8, 0x0c, 0xc8, 0x2f,
	0x2e, 0xf1, 0x4c, 0x81, 0x72, 0x84, 0xa4, 0xb9, 0x38, 0x93, 0xf3, 0x4b, 0x8b, 0x8a, 0x53, 0xe3,
	0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0x38, 0x20, 0x02, 0x9e, 0x29, 0x4a, 0xc6,
	0x5c, 0x12, 0x98, 0xfa, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xc4, 0xb9, 0xd8, 0xa1, 0x56,
	0x42, 0xb5, 0xb1, 0x15, 0x80, 0x15, 0x38, 0x09, 0x45, 0x09, 0xe8, 0xe9, 0x5b, 0xa7, 0x17, 0x15,
	0x24, 0xc7, 0xc3, 0x1c, 0x90, 0xc4, 0x06, 0x76, 0x91, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa3,
	0x0c, 0xb0, 0x03, 0xb2, 0x00, 0x00, 0x00,
}

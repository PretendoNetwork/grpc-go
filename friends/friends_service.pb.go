// Code generated by protoc-gen-go. DO NOT EDIT.
// source: friends_service.proto

package grpc_friends

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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

func init() {
	proto.RegisterFile("friends_service.proto", fileDescriptor_76c8dd950d0de92e)
}

var fileDescriptor_76c8dd950d0de92e = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0x2b, 0x31,
	0x14, 0x85, 0xfb, 0xe0, 0xf1, 0x0a, 0xd9, 0x3c, 0x1a, 0xa9, 0x48, 0xa5, 0x8a, 0xa3, 0x15, 0x41,
	0x4c, 0x41, 0x97, 0xae, 0x94, 0xaa, 0x74, 0x23, 0xa2, 0x14, 0x41, 0x8a, 0xa1, 0xcd, 0xdc, 0x8e,
	0x01, 0x9b, 0xc4, 0x24, 0xa3, 0x74, 0xe7, 0xaf, 0x76, 0x2d, 0x6d, 0x92, 0xe9, 0x4c, 0xa7, 0xad,
	0x6e, 0x73, 0xbe, 0x9c, 0x73, 0x73, 0x6e, 0x50, 0x7d, 0xa4, 0x39, 0x88, 0xd8, 0x50, 0x03, 0xfa,
	0x9d, 0x33, 0x20, 0x4a, 0x4b, 0x2b, 0x71, 0xd5, 0x1f, 0x37, 0xb6, 0x13, 0x29, 0x93, 0x57, 0x68,
	0xcf, 0x8e, 0x87, 0xe9, 0xa8, 0x0d, 0x63, 0x65, 0x27, 0x8e, 0x6a, 0xb4, 0x0c, 0x88, 0x98, 0xa6,
	0x06, 0x34, 0x15, 0xd2, 0xf2, 0x11, 0x67, 0x03, 0xcb, 0xa5, 0xa0, 0x1f, 0x9c, 0xa7, 0x54, 0x2b,
	0xe6, 0xb1, 0x9d, 0x04, 0xac, 0xa3, 0x9c, 0x2b, 0x55, 0x3c, 0x36, 0x39, 0x3d, 0x9a, 0xdb, 0x78,
	0x40, 0xc3, 0x5b, 0x0a, 0xc6, 0xe6, 0x98, 0xe3, 0x45, 0x0f, 0x8f, 0x18, 0xca, 0x05, 0x93, 0x63,
	0x2e, 0x92, 0x1c, 0xbc, 0x3b, 0x60, 0x0c, 0x94, 0x5d, 0xed, 0xd6, 0x8c, 0x41, 0x4c, 0x56, 0xca,
	0xa7, 0x5f, 0x7f, 0x51, 0xf5, 0xda, 0x15, 0x80, 0x9f, 0xd1, 0xd6, 0x03, 0x88, 0xb8, 0x67, 0x40,
	0xdf, 0xe6, 0xde, 0xf8, 0xc8, 0x79, 0x0f, 0x1f, 0x11, 0x5f, 0x13, 0x59, 0x85, 0xdc, 0x3b, 0xef,
	0xc6, 0x26, 0x71, 0x3d, 0x92, 0xd0, 0x23, 0xb9, 0x9a, 0xf6, 0x18, 0x55, 0x70, 0x1f, 0xd5, 0x6e,
	0xc0, 0x4e, 0xef, 0xba, 0xc4, 0xbb, 0x6e, 0xc7, 0xe0, 0xbd, 0xcc, 0xb8, 0xa4, 0x05, 0xc7, 0x68,
	0x1d, 0x62, 0x94, 0x14, 0x06, 0xa2, 0x0a, 0x7e, 0x41, 0xf5, 0x30, 0x9a, 0xd3, 0xfd, 0x75, 0xdc,
	0x2a, 0x8d, 0x5e, 0xd0, 0x43, 0xca, 0xe1, 0x4f, 0x58, 0x96, 0xf4, 0xf9, 0x07, 0x35, 0x0b, 0x93,
	0x78, 0xc4, 0x74, 0xfd, 0x7e, 0xf0, 0xc9, 0xf2, 0x89, 0x17, 0xb9, 0x10, 0x4d, 0x7e, 0x8b, 0x67,
	0x23, 0x0c, 0xd1, 0xc6, 0xc5, 0x6c, 0xf1, 0xc5, 0xa7, 0xee, 0x67, 0x46, 0x4b, 0xd4, 0x90, 0x76,
	0xb0, 0x1e, 0xca, 0x32, 0xfa, 0xa8, 0xd6, 0x01, 0x31, 0x29, 0x26, 0xcc, 0xd7, 0x55, 0xd2, 0xca,
	0xeb, 0x5a, 0x82, 0x04, 0xf7, 0xcb, 0xda, 0xd3, 0x7f, 0xd2, 0x3e, 0x4f, 0xb4, 0x62, 0xfe, 0x73,
	0x9a, 0xe1, 0xbf, 0xd9, 0x8f, 0x39, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x53, 0x39, 0xf0, 0xd7,
	0xa9, 0x03, 0x00, 0x00,
}

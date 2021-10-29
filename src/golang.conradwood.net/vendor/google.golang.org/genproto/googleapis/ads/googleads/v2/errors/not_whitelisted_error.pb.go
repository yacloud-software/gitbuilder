// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/errors/not_whitelisted_error.proto

package errors

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Enum describing possible not whitelisted errors.
type NotWhitelistedErrorEnum_NotWhitelistedError int32

const (
	// Enum unspecified.
	NotWhitelistedErrorEnum_UNSPECIFIED NotWhitelistedErrorEnum_NotWhitelistedError = 0
	// The received error code is not known in this version.
	NotWhitelistedErrorEnum_UNKNOWN NotWhitelistedErrorEnum_NotWhitelistedError = 1
	// Customer is not whitelisted for accessing this feature.
	NotWhitelistedErrorEnum_CUSTOMER_NOT_WHITELISTED_FOR_THIS_FEATURE NotWhitelistedErrorEnum_NotWhitelistedError = 2
)

var NotWhitelistedErrorEnum_NotWhitelistedError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CUSTOMER_NOT_WHITELISTED_FOR_THIS_FEATURE",
}

var NotWhitelistedErrorEnum_NotWhitelistedError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"CUSTOMER_NOT_WHITELISTED_FOR_THIS_FEATURE": 2,
}

func (x NotWhitelistedErrorEnum_NotWhitelistedError) String() string {
	return proto.EnumName(NotWhitelistedErrorEnum_NotWhitelistedError_name, int32(x))
}

func (NotWhitelistedErrorEnum_NotWhitelistedError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3a2ce7a1e2c42fa4, []int{0, 0}
}

// Container for enum describing possible not whitelisted errors.
type NotWhitelistedErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotWhitelistedErrorEnum) Reset()         { *m = NotWhitelistedErrorEnum{} }
func (m *NotWhitelistedErrorEnum) String() string { return proto.CompactTextString(m) }
func (*NotWhitelistedErrorEnum) ProtoMessage()    {}
func (*NotWhitelistedErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a2ce7a1e2c42fa4, []int{0}
}

func (m *NotWhitelistedErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotWhitelistedErrorEnum.Unmarshal(m, b)
}
func (m *NotWhitelistedErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotWhitelistedErrorEnum.Marshal(b, m, deterministic)
}
func (m *NotWhitelistedErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotWhitelistedErrorEnum.Merge(m, src)
}
func (m *NotWhitelistedErrorEnum) XXX_Size() int {
	return xxx_messageInfo_NotWhitelistedErrorEnum.Size(m)
}
func (m *NotWhitelistedErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_NotWhitelistedErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_NotWhitelistedErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.errors.NotWhitelistedErrorEnum_NotWhitelistedError", NotWhitelistedErrorEnum_NotWhitelistedError_name, NotWhitelistedErrorEnum_NotWhitelistedError_value)
	proto.RegisterType((*NotWhitelistedErrorEnum)(nil), "google.ads.googleads.v2.errors.NotWhitelistedErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/errors/not_whitelisted_error.proto", fileDescriptor_3a2ce7a1e2c42fa4)
}

var fileDescriptor_3a2ce7a1e2c42fa4 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xff, 0xe6, 0x07, 0x85, 0xe9, 0xc2, 0x12, 0x17, 0x8a, 0x48, 0x17, 0xd9, 0xb9, 0x70,
	0x02, 0x71, 0x37, 0xae, 0xd2, 0x76, 0xda, 0x06, 0x35, 0x29, 0x4d, 0xd2, 0x82, 0x04, 0x86, 0xd4,
	0x09, 0x63, 0xa0, 0x9d, 0x5b, 0x32, 0x63, 0x5d, 0xf9, 0x32, 0x2e, 0x7d, 0x14, 0x1f, 0xc5, 0xad,
	0x2f, 0x20, 0xc9, 0xd8, 0xba, 0xa9, 0xae, 0x72, 0xb8, 0x39, 0xdf, 0x99, 0x73, 0x2f, 0x22, 0x02,
	0x40, 0x2c, 0x0b, 0x37, 0xe7, 0xca, 0x35, 0xb2, 0x56, 0x1b, 0xcf, 0x2d, 0xaa, 0x0a, 0x2a, 0xe5,
	0x4a, 0xd0, 0xec, 0xf9, 0xb1, 0xd4, 0xc5, 0xb2, 0x54, 0xba, 0xe0, 0xac, 0x19, 0xe3, 0x75, 0x05,
	0x1a, 0xec, 0xae, 0x01, 0x70, 0xce, 0x15, 0xde, 0xb1, 0x78, 0xe3, 0x61, 0xc3, 0x9e, 0x9d, 0x6f,
	0xb3, 0xd7, 0xa5, 0x9b, 0x4b, 0x09, 0x3a, 0xd7, 0x25, 0x48, 0x65, 0x68, 0xe7, 0x05, 0x9d, 0x84,
	0xa0, 0xe7, 0x3f, 0xd9, 0xb4, 0xa6, 0xa8, 0x7c, 0x5a, 0x39, 0x0b, 0x74, 0xbc, 0xe7, 0x97, 0x7d,
	0x84, 0xda, 0x69, 0x18, 0x4f, 0x68, 0x3f, 0x18, 0x06, 0x74, 0xd0, 0xf9, 0x67, 0xb7, 0xd1, 0x61,
	0x1a, 0xde, 0x84, 0xd1, 0x3c, 0xec, 0xb4, 0xec, 0x4b, 0x74, 0xd1, 0x4f, 0xe3, 0x24, 0xba, 0xa3,
	0x53, 0x16, 0x46, 0x09, 0x9b, 0x8f, 0x83, 0x84, 0xde, 0x06, 0x71, 0x42, 0x07, 0x6c, 0x18, 0x4d,
	0x59, 0x32, 0x0e, 0x62, 0x36, 0xa4, 0x7e, 0x92, 0x4e, 0x69, 0xc7, 0xea, 0x7d, 0xb6, 0x90, 0xf3,
	0x00, 0x2b, 0xfc, 0xf7, 0x0e, 0xbd, 0xd3, 0x3d, 0x45, 0x26, 0x75, 0xff, 0x49, 0xeb, 0x7e, 0xf0,
	0xcd, 0x0a, 0x58, 0xe6, 0x52, 0x60, 0xa8, 0x84, 0x2b, 0x0a, 0xd9, 0x6c, 0xb7, 0xbd, 0xe5, 0xba,
	0x54, 0xbf, 0x9d, 0xf6, 0xda, 0x7c, 0x5e, 0xad, 0xff, 0x23, 0xdf, 0x7f, 0xb3, 0xba, 0x23, 0x13,
	0xe6, 0x73, 0x85, 0x8d, 0xac, 0xd5, 0xcc, 0xc3, 0xcd, 0x93, 0xea, 0x7d, 0x6b, 0xc8, 0x7c, 0xae,
	0xb2, 0x9d, 0x21, 0x9b, 0x79, 0x99, 0x31, 0x7c, 0x58, 0x8e, 0x99, 0x12, 0xe2, 0x73, 0x45, 0xc8,
	0xce, 0x42, 0xc8, 0xcc, 0x23, 0xc4, 0x98, 0x16, 0x07, 0x4d, 0xbb, 0xab, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf6, 0xa7, 0x9d, 0xee, 0xf7, 0x01, 0x00, 0x00,
}

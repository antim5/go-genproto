// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/feed_item_status.proto

package enums

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

// Possible statuses of a feed item.
type FeedItemStatusEnum_FeedItemStatus int32

const (
	// Not specified.
	FeedItemStatusEnum_UNSPECIFIED FeedItemStatusEnum_FeedItemStatus = 0
	// Used for return value only. Represents value unknown in this version.
	FeedItemStatusEnum_UNKNOWN FeedItemStatusEnum_FeedItemStatus = 1
	// Feed item is enabled.
	FeedItemStatusEnum_ENABLED FeedItemStatusEnum_FeedItemStatus = 2
	// Feed item has been removed.
	FeedItemStatusEnum_REMOVED FeedItemStatusEnum_FeedItemStatus = 3
)

var FeedItemStatusEnum_FeedItemStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}

var FeedItemStatusEnum_FeedItemStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x FeedItemStatusEnum_FeedItemStatus) String() string {
	return proto.EnumName(FeedItemStatusEnum_FeedItemStatus_name, int32(x))
}

func (FeedItemStatusEnum_FeedItemStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_15b96e530cd06c22, []int{0, 0}
}

// Container for enum describing possible statuses of a feed item.
type FeedItemStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedItemStatusEnum) Reset()         { *m = FeedItemStatusEnum{} }
func (m *FeedItemStatusEnum) String() string { return proto.CompactTextString(m) }
func (*FeedItemStatusEnum) ProtoMessage()    {}
func (*FeedItemStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_15b96e530cd06c22, []int{0}
}

func (m *FeedItemStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedItemStatusEnum.Unmarshal(m, b)
}
func (m *FeedItemStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedItemStatusEnum.Marshal(b, m, deterministic)
}
func (m *FeedItemStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedItemStatusEnum.Merge(m, src)
}
func (m *FeedItemStatusEnum) XXX_Size() int {
	return xxx_messageInfo_FeedItemStatusEnum.Size(m)
}
func (m *FeedItemStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedItemStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedItemStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.FeedItemStatusEnum_FeedItemStatus", FeedItemStatusEnum_FeedItemStatus_name, FeedItemStatusEnum_FeedItemStatus_value)
	proto.RegisterType((*FeedItemStatusEnum)(nil), "google.ads.googleads.v3.enums.FeedItemStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/feed_item_status.proto", fileDescriptor_15b96e530cd06c22)
}

var fileDescriptor_15b96e530cd06c22 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcf, 0x4a, 0xc3, 0x30,
	0x18, 0x77, 0x1d, 0x28, 0x64, 0xa0, 0xa5, 0xde, 0xc4, 0x1d, 0xb6, 0x07, 0x48, 0x0e, 0xf1, 0x14,
	0x4f, 0xa9, 0xcb, 0xe6, 0x50, 0xbb, 0xe1, 0x58, 0x05, 0x29, 0x8e, 0x68, 0x62, 0x28, 0xac, 0xc9,
	0x58, 0xb2, 0x3d, 0x90, 0x47, 0x1f, 0xc5, 0x27, 0x11, 0x9f, 0x42, 0x92, 0x6e, 0x85, 0x1d, 0xf4,
	0x52, 0x7e, 0xdf, 0xf7, 0xfb, 0xd3, 0x5f, 0x3e, 0x70, 0xa5, 0x8c, 0x51, 0x4b, 0x89, 0xb8, 0xb0,
	0xa8, 0x86, 0x1e, 0x6d, 0x31, 0x92, 0x7a, 0x53, 0x59, 0xf4, 0x2e, 0xa5, 0x58, 0x94, 0x4e, 0x56,
	0x0b, 0xeb, 0xb8, 0xdb, 0x58, 0xb8, 0x5a, 0x1b, 0x67, 0x92, 0x6e, 0x2d, 0x85, 0x5c, 0x58, 0xd8,
	0xb8, 0xe0, 0x16, 0xc3, 0xe0, 0xba, 0xb8, 0xdc, 0x87, 0xae, 0x4a, 0xc4, 0xb5, 0x36, 0x8e, 0xbb,
	0xd2, 0xe8, 0x9d, 0xb9, 0xff, 0x02, 0x92, 0xa1, 0x94, 0x62, 0xec, 0x64, 0x35, 0x0b, 0xa1, 0x4c,
	0x6f, 0xaa, 0xfe, 0x2d, 0x38, 0x3d, 0xdc, 0x26, 0x67, 0xa0, 0x33, 0xcf, 0x66, 0x53, 0x76, 0x33,
	0x1e, 0x8e, 0xd9, 0x20, 0x3e, 0x4a, 0x3a, 0xe0, 0x64, 0x9e, 0xdd, 0x65, 0x93, 0xa7, 0x2c, 0x6e,
	0xf9, 0x81, 0x65, 0x34, 0xbd, 0x67, 0x83, 0x38, 0xf2, 0xc3, 0x23, 0x7b, 0x98, 0xe4, 0x6c, 0x10,
	0xb7, 0xd3, 0xef, 0x16, 0xe8, 0xbd, 0x99, 0x0a, 0xfe, 0xdb, 0x31, 0x3d, 0x3f, 0xfc, 0xdb, 0xd4,
	0x57, 0x9b, 0xb6, 0x9e, 0xd3, 0x9d, 0x4b, 0x99, 0x25, 0xd7, 0x0a, 0x9a, 0xb5, 0x42, 0x4a, 0xea,
	0x50, 0x7c, 0x7f, 0x9f, 0x55, 0x69, 0xff, 0x38, 0xd7, 0x75, 0xf8, 0x7e, 0x44, 0xed, 0x11, 0xa5,
	0x9f, 0x51, 0x77, 0x54, 0x47, 0x51, 0x61, 0x61, 0x0d, 0x3d, 0xca, 0x31, 0xf4, 0xef, 0xb5, 0x5f,
	0x7b, 0xbe, 0xa0, 0xc2, 0x16, 0x0d, 0x5f, 0xe4, 0xb8, 0x08, 0xfc, 0x4f, 0xd4, 0xab, 0x97, 0x84,
	0x50, 0x61, 0x09, 0x69, 0x14, 0x84, 0xe4, 0x98, 0x90, 0xa0, 0x79, 0x3d, 0x0e, 0xc5, 0xf0, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x09, 0x71, 0x41, 0xfa, 0xc6, 0x01, 0x00, 0x00,
}
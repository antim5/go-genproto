// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/enums/conversion_action_type.proto

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

// Possible types of a conversion action.
type ConversionActionTypeEnum_ConversionActionType int32

const (
	// Not specified.
	ConversionActionTypeEnum_UNSPECIFIED ConversionActionTypeEnum_ConversionActionType = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionActionTypeEnum_UNKNOWN ConversionActionTypeEnum_ConversionActionType = 1
	// Conversions that occur when a user clicks on an ad's call extension.
	ConversionActionTypeEnum_AD_CALL ConversionActionTypeEnum_ConversionActionType = 2
	// Conversions that occur when a user on a mobile device clicks a phone
	// number.
	ConversionActionTypeEnum_CLICK_TO_CALL ConversionActionTypeEnum_ConversionActionType = 3
	// Conversions that occur when a user downloads a mobile app from the Google
	// Play Store.
	ConversionActionTypeEnum_GOOGLE_PLAY_DOWNLOAD ConversionActionTypeEnum_ConversionActionType = 4
	// Conversions that occur when a user makes a purchase in an app through
	// Android billing.
	ConversionActionTypeEnum_GOOGLE_PLAY_IN_APP_PURCHASE ConversionActionTypeEnum_ConversionActionType = 5
	// Call conversions that are tracked by the advertiser and uploaded.
	ConversionActionTypeEnum_UPLOAD_CALLS ConversionActionTypeEnum_ConversionActionType = 6
	// Conversions that are tracked by the advertiser and uploaded with
	// attributed clicks.
	ConversionActionTypeEnum_UPLOAD_CLICKS ConversionActionTypeEnum_ConversionActionType = 7
	// Conversions that occur on a webpage.
	ConversionActionTypeEnum_WEBPAGE ConversionActionTypeEnum_ConversionActionType = 8
	// Conversions that occur when a user calls a dynamically-generated phone
	// number from an advertiser's website.
	ConversionActionTypeEnum_WEBSITE_CALL ConversionActionTypeEnum_ConversionActionType = 9
)

var ConversionActionTypeEnum_ConversionActionType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AD_CALL",
	3: "CLICK_TO_CALL",
	4: "GOOGLE_PLAY_DOWNLOAD",
	5: "GOOGLE_PLAY_IN_APP_PURCHASE",
	6: "UPLOAD_CALLS",
	7: "UPLOAD_CLICKS",
	8: "WEBPAGE",
	9: "WEBSITE_CALL",
}

var ConversionActionTypeEnum_ConversionActionType_value = map[string]int32{
	"UNSPECIFIED":                 0,
	"UNKNOWN":                     1,
	"AD_CALL":                     2,
	"CLICK_TO_CALL":               3,
	"GOOGLE_PLAY_DOWNLOAD":        4,
	"GOOGLE_PLAY_IN_APP_PURCHASE": 5,
	"UPLOAD_CALLS":                6,
	"UPLOAD_CLICKS":               7,
	"WEBPAGE":                     8,
	"WEBSITE_CALL":                9,
}

func (x ConversionActionTypeEnum_ConversionActionType) String() string {
	return proto.EnumName(ConversionActionTypeEnum_ConversionActionType_name, int32(x))
}

func (ConversionActionTypeEnum_ConversionActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3050c3c63416f7a0, []int{0, 0}
}

// Container for enum describing possible types of a conversion action.
type ConversionActionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionActionTypeEnum) Reset()         { *m = ConversionActionTypeEnum{} }
func (m *ConversionActionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionActionTypeEnum) ProtoMessage()    {}
func (*ConversionActionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_3050c3c63416f7a0, []int{0}
}

func (m *ConversionActionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionActionTypeEnum.Unmarshal(m, b)
}
func (m *ConversionActionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionActionTypeEnum.Marshal(b, m, deterministic)
}
func (m *ConversionActionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionActionTypeEnum.Merge(m, src)
}
func (m *ConversionActionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionActionTypeEnum.Size(m)
}
func (m *ConversionActionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionActionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionActionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.enums.ConversionActionTypeEnum_ConversionActionType", ConversionActionTypeEnum_ConversionActionType_name, ConversionActionTypeEnum_ConversionActionType_value)
	proto.RegisterType((*ConversionActionTypeEnum)(nil), "google.ads.googleads.v3.enums.ConversionActionTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/enums/conversion_action_type.proto", fileDescriptor_3050c3c63416f7a0)
}

var fileDescriptor_3050c3c63416f7a0 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xdf, 0x8e, 0x94, 0x30,
	0x14, 0xc6, 0x85, 0xd5, 0x5d, 0xed, 0x6a, 0xac, 0x64, 0x2f, 0xd6, 0x3f, 0x1b, 0xb3, 0xfb, 0x00,
	0xe5, 0x82, 0xbb, 0x7a, 0x55, 0xa0, 0x22, 0x59, 0x02, 0x8d, 0x0c, 0x43, 0x34, 0x24, 0x04, 0x07,
	0x42, 0x26, 0xd9, 0x69, 0xc9, 0x94, 0x99, 0x64, 0x5e, 0xc7, 0x4b, 0x1f, 0xc5, 0x5b, 0xdf, 0xc2,
	0x0b, 0xe3, 0x23, 0x98, 0xb6, 0x42, 0xbc, 0x18, 0xf7, 0x06, 0x4e, 0xcf, 0x77, 0xbe, 0xdf, 0x69,
	0xcf, 0x01, 0xb8, 0x17, 0xa2, 0xbf, 0xeb, 0xdc, 0xa6, 0x95, 0xae, 0x09, 0x55, 0xb4, 0xf7, 0xdc,
	0x8e, 0xef, 0x36, 0xd2, 0x5d, 0x09, 0xbe, 0xef, 0xb6, 0x72, 0x2d, 0x78, 0xdd, 0xac, 0x46, 0xf5,
	0x1b, 0x0f, 0x43, 0x87, 0x86, 0xad, 0x18, 0x85, 0x73, 0x65, 0x0c, 0xa8, 0x69, 0x25, 0x9a, 0xbd,
	0x68, 0xef, 0x21, 0xed, 0x7d, 0xf5, 0x66, 0x42, 0x0f, 0x6b, 0xb7, 0xe1, 0x5c, 0x8c, 0x8d, 0x02,
	0x48, 0x63, 0xbe, 0xf9, 0x6d, 0x81, 0xcb, 0x60, 0xa6, 0x13, 0x0d, 0x5f, 0x1c, 0x86, 0x8e, 0xf2,
	0xdd, 0xe6, 0xe6, 0x87, 0x05, 0x2e, 0x8e, 0x89, 0xce, 0x73, 0x70, 0x5e, 0xa4, 0x39, 0xa3, 0x41,
	0xfc, 0x3e, 0xa6, 0x21, 0x7c, 0xe0, 0x9c, 0x83, 0xb3, 0x22, 0xbd, 0x4d, 0xb3, 0x32, 0x85, 0x96,
	0x3a, 0x90, 0xb0, 0x0e, 0x48, 0x92, 0x40, 0xdb, 0x79, 0x01, 0x9e, 0x05, 0x49, 0x1c, 0xdc, 0xd6,
	0x8b, 0xcc, 0xa4, 0x4e, 0x9c, 0x4b, 0x70, 0x11, 0x65, 0x59, 0x94, 0xd0, 0x9a, 0x25, 0xe4, 0x53,
	0x1d, 0x66, 0x65, 0x9a, 0x64, 0x24, 0x84, 0x0f, 0x9d, 0xb7, 0xe0, 0xf5, 0xbf, 0x4a, 0x9c, 0xd6,
	0x84, 0xb1, 0x9a, 0x15, 0x1f, 0x83, 0x0f, 0x24, 0xa7, 0xf0, 0x91, 0x03, 0xc1, 0xd3, 0x82, 0xa9,
	0x62, 0xcd, 0xca, 0xe1, 0xa9, 0xe2, 0x4f, 0x19, 0xd5, 0x26, 0x87, 0x67, 0xaa, 0x7f, 0x49, 0x7d,
	0x46, 0x22, 0x0a, 0x1f, 0x2b, 0x47, 0x49, 0xfd, 0x3c, 0x5e, 0x50, 0xd3, 0xfe, 0x89, 0xff, 0xcb,
	0x02, 0xd7, 0x2b, 0xb1, 0x41, 0xf7, 0x8e, 0xcd, 0x7f, 0x79, 0xec, 0xe1, 0x4c, 0xcd, 0x8c, 0x59,
	0x9f, 0xfd, 0xbf, 0xde, 0x5e, 0xdc, 0x35, 0xbc, 0x47, 0x62, 0xdb, 0xbb, 0x7d, 0xc7, 0xf5, 0x44,
	0xa7, 0xf5, 0x0d, 0x6b, 0xf9, 0x9f, 0x6d, 0xbe, 0xd3, 0xdf, 0xaf, 0xf6, 0x49, 0x44, 0xc8, 0x37,
	0xfb, 0x2a, 0x32, 0x28, 0xd2, 0x4a, 0x64, 0x42, 0x15, 0x2d, 0x3d, 0xa4, 0x36, 0x20, 0xbf, 0x4f,
	0x7a, 0x45, 0x5a, 0x59, 0xcd, 0x7a, 0xb5, 0xf4, 0x2a, 0xad, 0xff, 0xb4, 0xaf, 0x4d, 0x12, 0x63,
	0xd2, 0x4a, 0x8c, 0xe7, 0x0a, 0x8c, 0x97, 0x1e, 0xc6, 0xba, 0xe6, 0xcb, 0xa9, 0xbe, 0x98, 0xf7,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x91, 0xd6, 0x3d, 0x8e, 0x65, 0x02, 0x00, 0x00,
}
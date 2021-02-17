// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: google/maps/routes/v1/vehicle_emission_type.proto

package routes

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A set of values describing the vehicle's emission type.
// Applies only to the DRIVE travel mode.
type VehicleEmissionType int32

const (
	// No emission type specified. Default to GASOLINE.
	VehicleEmissionType_VEHICLE_EMISSION_TYPE_UNSPECIFIED VehicleEmissionType = 0
	// Gasoline/petrol fueled vehicle.
	VehicleEmissionType_GASOLINE VehicleEmissionType = 1
	// Electricity powered vehicle.
	VehicleEmissionType_ELECTRIC VehicleEmissionType = 2
	// Hybrid fuel (such as gasoline + electric) vehicle.
	VehicleEmissionType_HYBRID VehicleEmissionType = 3
)

// Enum value maps for VehicleEmissionType.
var (
	VehicleEmissionType_name = map[int32]string{
		0: "VEHICLE_EMISSION_TYPE_UNSPECIFIED",
		1: "GASOLINE",
		2: "ELECTRIC",
		3: "HYBRID",
	}
	VehicleEmissionType_value = map[string]int32{
		"VEHICLE_EMISSION_TYPE_UNSPECIFIED": 0,
		"GASOLINE":                          1,
		"ELECTRIC":                          2,
		"HYBRID":                            3,
	}
)

func (x VehicleEmissionType) Enum() *VehicleEmissionType {
	p := new(VehicleEmissionType)
	*p = x
	return p
}

func (x VehicleEmissionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VehicleEmissionType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_routes_v1_vehicle_emission_type_proto_enumTypes[0].Descriptor()
}

func (VehicleEmissionType) Type() protoreflect.EnumType {
	return &file_google_maps_routes_v1_vehicle_emission_type_proto_enumTypes[0]
}

func (x VehicleEmissionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VehicleEmissionType.Descriptor instead.
func (VehicleEmissionType) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_routes_v1_vehicle_emission_type_proto_rawDescGZIP(), []int{0}
}

var File_google_maps_routes_v1_vehicle_emission_type_proto protoreflect.FileDescriptor

var file_google_maps_routes_v1_vehicle_emission_type_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f,
	0x65, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0x64, 0x0a, 0x13, 0x56, 0x65,
	0x68, 0x69, 0x63, 0x6c, 0x65, 0x45, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x25, 0x0a, 0x21, 0x56, 0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x45, 0x4d, 0x49,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x41, 0x53, 0x4f,
	0x4c, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x52,
	0x49, 0x43, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x59, 0x42, 0x52, 0x49, 0x44, 0x10, 0x03,
	0x42, 0xae, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x18,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x45, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x04, 0x47, 0x4d, 0x52,
	0x53, 0xaa, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x4d, 0x61, 0x70, 0x73, 0x2e,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x4d, 0x61, 0x70, 0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x5c, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_routes_v1_vehicle_emission_type_proto_rawDescOnce sync.Once
	file_google_maps_routes_v1_vehicle_emission_type_proto_rawDescData = file_google_maps_routes_v1_vehicle_emission_type_proto_rawDesc
)

func file_google_maps_routes_v1_vehicle_emission_type_proto_rawDescGZIP() []byte {
	file_google_maps_routes_v1_vehicle_emission_type_proto_rawDescOnce.Do(func() {
		file_google_maps_routes_v1_vehicle_emission_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_routes_v1_vehicle_emission_type_proto_rawDescData)
	})
	return file_google_maps_routes_v1_vehicle_emission_type_proto_rawDescData
}

var file_google_maps_routes_v1_vehicle_emission_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_maps_routes_v1_vehicle_emission_type_proto_goTypes = []interface{}{
	(VehicleEmissionType)(0), // 0: google.maps.routes.v1.VehicleEmissionType
}
var file_google_maps_routes_v1_vehicle_emission_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_maps_routes_v1_vehicle_emission_type_proto_init() }
func file_google_maps_routes_v1_vehicle_emission_type_proto_init() {
	if File_google_maps_routes_v1_vehicle_emission_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_maps_routes_v1_vehicle_emission_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_routes_v1_vehicle_emission_type_proto_goTypes,
		DependencyIndexes: file_google_maps_routes_v1_vehicle_emission_type_proto_depIdxs,
		EnumInfos:         file_google_maps_routes_v1_vehicle_emission_type_proto_enumTypes,
	}.Build()
	File_google_maps_routes_v1_vehicle_emission_type_proto = out.File
	file_google_maps_routes_v1_vehicle_emission_type_proto_rawDesc = nil
	file_google_maps_routes_v1_vehicle_emission_type_proto_goTypes = nil
	file_google_maps_routes_v1_vehicle_emission_type_proto_depIdxs = nil
}

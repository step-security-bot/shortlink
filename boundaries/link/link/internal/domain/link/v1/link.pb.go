// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: domain/link/v1/link.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Link event
type LinkEvent int32

const (
	// Unspecified
	LinkEvent_LINK_EVENT_UNSPECIFIED LinkEvent = 0
	// Add link
	LinkEvent_LINK_EVENT_ADD LinkEvent = 1
	// Get link
	LinkEvent_LINK_EVENT_GET LinkEvent = 2
	// List link
	LinkEvent_LINK_EVENT_LIST LinkEvent = 3
	// Update link
	LinkEvent_LINK_EVENT_UPDATE LinkEvent = 4
	// Delete link
	LinkEvent_LINK_EVENT_DELETE LinkEvent = 5
)

// Enum value maps for LinkEvent.
var (
	LinkEvent_name = map[int32]string{
		0: "LINK_EVENT_UNSPECIFIED",
		1: "LINK_EVENT_ADD",
		2: "LINK_EVENT_GET",
		3: "LINK_EVENT_LIST",
		4: "LINK_EVENT_UPDATE",
		5: "LINK_EVENT_DELETE",
	}
	LinkEvent_value = map[string]int32{
		"LINK_EVENT_UNSPECIFIED": 0,
		"LINK_EVENT_ADD":         1,
		"LINK_EVENT_GET":         2,
		"LINK_EVENT_LIST":        3,
		"LINK_EVENT_UPDATE":      4,
		"LINK_EVENT_DELETE":      5,
	}
)

func (x LinkEvent) Enum() *LinkEvent {
	p := new(LinkEvent)
	*p = x
	return p
}

func (x LinkEvent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LinkEvent) Descriptor() protoreflect.EnumDescriptor {
	return file_domain_link_v1_link_proto_enumTypes[0].Descriptor()
}

func (LinkEvent) Type() protoreflect.EnumType {
	return &file_domain_link_v1_link_proto_enumTypes[0]
}

func (x LinkEvent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LinkEvent.Descriptor instead.
func (LinkEvent) EnumDescriptor() ([]byte, []int) {
	return file_domain_link_v1_link_proto_rawDescGZIP(), []int{0}
}

var File_domain_link_v1_link_proto protoreflect.FileDescriptor

var file_domain_link_v1_link_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2a, 0x92, 0x01, 0x0a, 0x09,
	0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x16, 0x4c, 0x49, 0x4e,
	0x4b, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x41, 0x44, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x49, 0x4e,
	0x4b, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x47, 0x45, 0x54, 0x10, 0x02, 0x12, 0x13, 0x0a,
	0x0f, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x49, 0x53, 0x54,
	0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x49, 0x4e,
	0x4b, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x05,
	0x42, 0xca, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72,
	0x69, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6c, 0x69,
	0x6e, 0x6b, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x4c, 0x58, 0xaa, 0x02, 0x0e, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x4c, 0x69, 0x6e, 0x6b, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x4c, 0x69, 0x6e, 0x6b, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x4c, 0x69, 0x6e, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_link_v1_link_proto_rawDescOnce sync.Once
	file_domain_link_v1_link_proto_rawDescData = file_domain_link_v1_link_proto_rawDesc
)

func file_domain_link_v1_link_proto_rawDescGZIP() []byte {
	file_domain_link_v1_link_proto_rawDescOnce.Do(func() {
		file_domain_link_v1_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_link_v1_link_proto_rawDescData)
	})
	return file_domain_link_v1_link_proto_rawDescData
}

var file_domain_link_v1_link_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_domain_link_v1_link_proto_goTypes = []interface{}{
	(LinkEvent)(0), // 0: domain.link.v1.LinkEvent
}
var file_domain_link_v1_link_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_domain_link_v1_link_proto_init() }
func file_domain_link_v1_link_proto_init() {
	if File_domain_link_v1_link_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_domain_link_v1_link_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_link_v1_link_proto_goTypes,
		DependencyIndexes: file_domain_link_v1_link_proto_depIdxs,
		EnumInfos:         file_domain_link_v1_link_proto_enumTypes,
	}.Build()
	File_domain_link_v1_link_proto = out.File
	file_domain_link_v1_link_proto_rawDesc = nil
	file_domain_link_v1_link_proto_goTypes = nil
	file_domain_link_v1_link_proto_depIdxs = nil
}

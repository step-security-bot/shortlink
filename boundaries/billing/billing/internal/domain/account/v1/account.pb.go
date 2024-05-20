// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: domain/account/v1/account.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Event is a billing account event
type Event int32

const (
	// event unspecified
	Event_EVENT_UNSPECIFIED Event = 0
	// account new
	Event_EVENT_ACCOUNT_NEW Event = 1
	// account delete
	Event_EVENT_ACCOUNT_DELETE Event = 2
)

// Enum value maps for Event.
var (
	Event_name = map[int32]string{
		0: "EVENT_UNSPECIFIED",
		1: "EVENT_ACCOUNT_NEW",
		2: "EVENT_ACCOUNT_DELETE",
	}
	Event_value = map[string]int32{
		"EVENT_UNSPECIFIED":    0,
		"EVENT_ACCOUNT_NEW":    1,
		"EVENT_ACCOUNT_DELETE": 2,
	}
)

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}

func (x Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event) Descriptor() protoreflect.EnumDescriptor {
	return file_domain_account_v1_account_proto_enumTypes[0].Descriptor()
}

func (Event) Type() protoreflect.EnumType {
	return &file_domain_account_v1_account_proto_enumTypes[0]
}

func (x Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event.Descriptor instead.
func (Event) EnumDescriptor() ([]byte, []int) {
	return file_domain_account_v1_account_proto_rawDescGZIP(), []int{0}
}

var File_domain_account_v1_account_proto protoreflect.FileDescriptor

var file_domain_account_v1_account_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x19, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x4f,
	0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15,
	0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f,
	0x4e, 0x45, 0x57, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41,
	0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x42,
	0x8e, 0x02, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x42, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x69, 0x65, 0x73,
	0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x42,
	0x41, 0xaa, 0x02, 0x19, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x25, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_account_v1_account_proto_rawDescOnce sync.Once
	file_domain_account_v1_account_proto_rawDescData = file_domain_account_v1_account_proto_rawDesc
)

func file_domain_account_v1_account_proto_rawDescGZIP() []byte {
	file_domain_account_v1_account_proto_rawDescOnce.Do(func() {
		file_domain_account_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_account_v1_account_proto_rawDescData)
	})
	return file_domain_account_v1_account_proto_rawDescData
}

var file_domain_account_v1_account_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_domain_account_v1_account_proto_goTypes = []interface{}{
	(Event)(0), // 0: domain.billing.account.v1.Event
}
var file_domain_account_v1_account_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_domain_account_v1_account_proto_init() }
func file_domain_account_v1_account_proto_init() {
	if File_domain_account_v1_account_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_domain_account_v1_account_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_account_v1_account_proto_goTypes,
		DependencyIndexes: file_domain_account_v1_account_proto_depIdxs,
		EnumInfos:         file_domain_account_v1_account_proto_enumTypes,
	}.Build()
	File_domain_account_v1_account_proto = out.File
	file_domain_account_v1_account_proto_rawDesc = nil
	file_domain_account_v1_account_proto_goTypes = nil
	file_domain_account_v1_account_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: meta.proto

package domain

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

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	ImageURL    string `protobuf:"bytes,2,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Keywords    string `protobuf:"bytes,4,opt,name=Keywords,proto3" json:"Keywords,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_meta_proto_rawDescGZIP(), []int{0}
}

func (x *Meta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Meta) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *Meta) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Meta) GetKeywords() string {
	if x != nil {
		return x.Keywords
	}
	return ""
}

var File_meta_proto protoreflect.FileDescriptor

var file_meta_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x70, 0x0a,
	0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52,
	0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52,
	0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42,
	0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61,
	0x74, 0x61, 0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_meta_proto_rawDescOnce sync.Once
	file_meta_proto_rawDescData = file_meta_proto_rawDesc
)

func file_meta_proto_rawDescGZIP() []byte {
	file_meta_proto_rawDescOnce.Do(func() {
		file_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_meta_proto_rawDescData)
	})
	return file_meta_proto_rawDescData
}

var file_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_meta_proto_goTypes = []interface{}{
	(*Meta)(nil), // 0: metadata_domain.Meta
}
var file_meta_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_meta_proto_init() }
func file_meta_proto_init() {
	if File_meta_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Meta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_meta_proto_goTypes,
		DependencyIndexes: file_meta_proto_depIdxs,
		MessageInfos:      file_meta_proto_msgTypes,
	}.Build()
	File_meta_proto = out.File
	file_meta_proto_rawDesc = nil
	file_meta_proto_goTypes = nil
	file_meta_proto_depIdxs = nil
}

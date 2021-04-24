// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: link_rpc.proto

package link_rpc

import (
	link "github.com/batazor/shortlink/internal/services/link/domain/link"
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

type LinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link   *link.Link `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	Filter string     `protobuf:"bytes,2,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *LinkRequest) Reset() {
	*x = LinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_link_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkRequest) ProtoMessage() {}

func (x *LinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_link_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkRequest.ProtoReflect.Descriptor instead.
func (*LinkRequest) Descriptor() ([]byte, []int) {
	return file_link_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *LinkRequest) GetLink() *link.Link {
	if x != nil {
		return x.Link
	}
	return nil
}

func (x *LinkRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

var File_link_rpc_proto protoreflect.FileDescriptor

var file_link_rpc_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x72, 0x70, 0x63, 0x1a, 0x0a, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x32, 0xbe, 0x01,
	0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1f, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0a, 0x2e,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x1a, 0x0a, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x00, 0x12, 0x1f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0a,
	0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x1a, 0x0a, 0x2e, 0x6c, 0x69, 0x6e,
	0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x15, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x73, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x0a, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x1a, 0x0a, 0x2e, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x0a, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b,
	0x1a, 0x0a, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x00, 0x42, 0x44,
	0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74,
	0x61, 0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x6c, 0x69, 0x6e, 0x6b,
	0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_link_rpc_proto_rawDescOnce sync.Once
	file_link_rpc_proto_rawDescData = file_link_rpc_proto_rawDesc
)

func file_link_rpc_proto_rawDescGZIP() []byte {
	file_link_rpc_proto_rawDescOnce.Do(func() {
		file_link_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_link_rpc_proto_rawDescData)
	})
	return file_link_rpc_proto_rawDescData
}

var file_link_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_link_rpc_proto_goTypes = []interface{}{
	(*LinkRequest)(nil), // 0: link_rpc.LinkRequest
	(*link.Link)(nil),   // 1: link.Link
	(*link.Links)(nil),  // 2: link.Links
}
var file_link_rpc_proto_depIdxs = []int32{
	1, // 0: link_rpc.LinkRequest.link:type_name -> link.Link
	1, // 1: link_rpc.Link.Add:input_type -> link.Link
	1, // 2: link_rpc.Link.Get:input_type -> link.Link
	0, // 3: link_rpc.Link.List:input_type -> link_rpc.LinkRequest
	1, // 4: link_rpc.Link.Update:input_type -> link.Link
	1, // 5: link_rpc.Link.Delete:input_type -> link.Link
	1, // 6: link_rpc.Link.Add:output_type -> link.Link
	1, // 7: link_rpc.Link.Get:output_type -> link.Link
	2, // 8: link_rpc.Link.List:output_type -> link.Links
	1, // 9: link_rpc.Link.Update:output_type -> link.Link
	1, // 10: link_rpc.Link.Delete:output_type -> link.Link
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_link_rpc_proto_init() }
func file_link_rpc_proto_init() {
	if File_link_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_link_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkRequest); i {
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
			RawDescriptor: file_link_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_link_rpc_proto_goTypes,
		DependencyIndexes: file_link_rpc_proto_depIdxs,
		MessageInfos:      file_link_rpc_proto_msgTypes,
	}.Build()
	File_link_rpc_proto = out.File
	file_link_rpc_proto_rawDesc = nil
	file_link_rpc_proto_goTypes = nil
	file_link_rpc_proto_depIdxs = nil
}

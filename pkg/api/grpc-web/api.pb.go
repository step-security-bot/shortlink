// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: api.proto

package grpcweb

import (
	link "github.com/batazor/shortlink/internal/services/api/domain/link"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

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
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkRequest) ProtoMessage() {}

func (x *LinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
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
	return file_api_proto_rawDescGZIP(), []int{0}
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

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x72, 0x70,
	0x63, 0x77, 0x65, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6c, 0x69,
	0x6e, 0x6b, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a,
	0x0b, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6c, 0x69, 0x6e,
	0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x32, 0x84, 0x02, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x41, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x77, 0x65, 0x62, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0b, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0x12, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x12, 0x3b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x0a, 0x2e, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x1a, 0x0a, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x7b, 0x68, 0x61, 0x73, 0x68, 0x7d, 0x12, 0x35, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x0a, 0x2e, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x1a, 0x0a, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c,
	0x69, 0x6e, 0x6b, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x22, 0x04, 0x2f, 0x61, 0x70,
	0x69, 0x3a, 0x01, 0x2a, 0x12, 0x45, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x0a, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x0b,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x7b, 0x68, 0x61, 0x73, 0x68, 0x7d, 0x42, 0x09, 0x5a, 0x07, 0x67,
	0x72, 0x70, 0x63, 0x77, 0x65, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_proto_goTypes = []interface{}{
	(*LinkRequest)(nil), // 0: grpcweb.LinkRequest
	(*link.Link)(nil),   // 1: link.Link
	(*link.Links)(nil),  // 2: link.Links
	(*empty.Empty)(nil), // 3: google.protobuf.Empty
}
var file_api_proto_depIdxs = []int32{
	1, // 0: grpcweb.LinkRequest.link:type_name -> link.Link
	0, // 1: grpcweb.Link.GetLinks:input_type -> grpcweb.LinkRequest
	1, // 2: grpcweb.Link.GetLink:input_type -> link.Link
	1, // 3: grpcweb.Link.CreateLink:input_type -> link.Link
	1, // 4: grpcweb.Link.DeleteLink:input_type -> link.Link
	2, // 5: grpcweb.Link.GetLinks:output_type -> link.Links
	1, // 6: grpcweb.Link.GetLink:output_type -> link.Link
	1, // 7: grpcweb.Link.CreateLink:output_type -> link.Link
	3, // 8: grpcweb.Link.DeleteLink:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

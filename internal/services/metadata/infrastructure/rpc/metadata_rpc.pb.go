// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: metadata_rpc.proto

package metadata_rpc

import (
	domain "github.com/batazor/shortlink/internal/services/metadata/domain"
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

type GetMetaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetMetaRequest) Reset() {
	*x = GetMetaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetaRequest) ProtoMessage() {}

func (x *GetMetaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetaRequest.ProtoReflect.Descriptor instead.
func (*GetMetaRequest) Descriptor() ([]byte, []int) {
	return file_metadata_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *GetMetaRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMetaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *domain.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *GetMetaResponse) Reset() {
	*x = GetMetaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMetaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetaResponse) ProtoMessage() {}

func (x *GetMetaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetaResponse.ProtoReflect.Descriptor instead.
func (*GetMetaResponse) Descriptor() ([]byte, []int) {
	return file_metadata_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *GetMetaResponse) GetMeta() *domain.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type SetMetaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *SetMetaRequest) Reset() {
	*x = SetMetaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMetaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMetaRequest) ProtoMessage() {}

func (x *SetMetaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMetaRequest.ProtoReflect.Descriptor instead.
func (*SetMetaRequest) Descriptor() ([]byte, []int) {
	return file_metadata_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *SetMetaRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SetMetaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *domain.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *SetMetaResponse) Reset() {
	*x = SetMetaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metadata_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMetaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMetaResponse) ProtoMessage() {}

func (x *SetMetaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_metadata_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMetaResponse.ProtoReflect.Descriptor instead.
func (*SetMetaResponse) Descriptor() ([]byte, []int) {
	return file_metadata_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *SetMetaResponse) GetMeta() *domain.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_metadata_rpc_proto protoreflect.FileDescriptor

var file_metadata_rpc_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72,
	0x70, 0x63, 0x1a, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64,
	0x22, 0x3c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x20,
	0x0a, 0x0e, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64,
	0x22, 0x3c, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x32, 0x96,
	0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x44, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x44, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x61, 0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metadata_rpc_proto_rawDescOnce sync.Once
	file_metadata_rpc_proto_rawDescData = file_metadata_rpc_proto_rawDesc
)

func file_metadata_rpc_proto_rawDescGZIP() []byte {
	file_metadata_rpc_proto_rawDescOnce.Do(func() {
		file_metadata_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_metadata_rpc_proto_rawDescData)
	})
	return file_metadata_rpc_proto_rawDescData
}

var file_metadata_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_metadata_rpc_proto_goTypes = []interface{}{
	(*GetMetaRequest)(nil),  // 0: metadata_rpc.GetMetaRequest
	(*GetMetaResponse)(nil), // 1: metadata_rpc.GetMetaResponse
	(*SetMetaRequest)(nil),  // 2: metadata_rpc.SetMetaRequest
	(*SetMetaResponse)(nil), // 3: metadata_rpc.SetMetaResponse
	(*domain.Meta)(nil),     // 4: metadata_domain.Meta
}
var file_metadata_rpc_proto_depIdxs = []int32{
	4, // 0: metadata_rpc.GetMetaResponse.meta:type_name -> metadata_domain.Meta
	4, // 1: metadata_rpc.SetMetaResponse.meta:type_name -> metadata_domain.Meta
	0, // 2: metadata_rpc.Metadata.Get:input_type -> metadata_rpc.GetMetaRequest
	2, // 3: metadata_rpc.Metadata.Set:input_type -> metadata_rpc.SetMetaRequest
	1, // 4: metadata_rpc.Metadata.Get:output_type -> metadata_rpc.GetMetaResponse
	3, // 5: metadata_rpc.Metadata.Set:output_type -> metadata_rpc.SetMetaResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_metadata_rpc_proto_init() }
func file_metadata_rpc_proto_init() {
	if File_metadata_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metadata_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetaRequest); i {
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
		file_metadata_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMetaResponse); i {
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
		file_metadata_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMetaRequest); i {
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
		file_metadata_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMetaResponse); i {
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
			RawDescriptor: file_metadata_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_metadata_rpc_proto_goTypes,
		DependencyIndexes: file_metadata_rpc_proto_depIdxs,
		MessageInfos:      file_metadata_rpc_proto_msgTypes,
	}.Build()
	File_metadata_rpc_proto = out.File
	file_metadata_rpc_proto_rawDesc = nil
	file_metadata_rpc_proto_goTypes = nil
	file_metadata_rpc_proto_depIdxs = nil
}

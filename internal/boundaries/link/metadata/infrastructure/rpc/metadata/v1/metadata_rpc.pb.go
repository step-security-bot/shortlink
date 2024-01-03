// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: infrastructure/rpc/metadata/v1/metadata_rpc.proto

package v1

import (
	v1 "github.com/shortlink-org/shortlink/internal/boundaries/link/metadata/domain/metadata/v1"
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

// MetadataServiceGetRequest is the request message for the Get method.
type MetadataServiceGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL is the URL to get the metadata for.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *MetadataServiceGetRequest) Reset() {
	*x = MetadataServiceGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataServiceGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataServiceGetRequest) ProtoMessage() {}

func (x *MetadataServiceGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataServiceGetRequest.ProtoReflect.Descriptor instead.
func (*MetadataServiceGetRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *MetadataServiceGetRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// MetadataServiceGetResponse is the response message for the Get method.
type MetadataServiceGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Meta is the metadata for the given URL.
	Meta *v1.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *MetadataServiceGetResponse) Reset() {
	*x = MetadataServiceGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataServiceGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataServiceGetResponse) ProtoMessage() {}

func (x *MetadataServiceGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataServiceGetResponse.ProtoReflect.Descriptor instead.
func (*MetadataServiceGetResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *MetadataServiceGetResponse) GetMeta() *v1.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

// MetadataServiceSetRequest is the request message for the Set method.
type MetadataServiceSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL is the URL to set the metadata for.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *MetadataServiceSetRequest) Reset() {
	*x = MetadataServiceSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataServiceSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataServiceSetRequest) ProtoMessage() {}

func (x *MetadataServiceSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataServiceSetRequest.ProtoReflect.Descriptor instead.
func (*MetadataServiceSetRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *MetadataServiceSetRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// MetadataServiceSetResponse is the response message for the Set method.
type MetadataServiceSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Meta is the metadata for the given URL.
	Meta *v1.Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *MetadataServiceSetResponse) Reset() {
	*x = MetadataServiceSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataServiceSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataServiceSetResponse) ProtoMessage() {}

func (x *MetadataServiceSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataServiceSetResponse.ProtoReflect.Descriptor instead.
func (*MetadataServiceSetResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *MetadataServiceSetResponse) GetMeta() *v1.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_infrastructure_rpc_metadata_v1_metadata_rpc_proto protoreflect.FileDescriptor

var file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_rawDesc = []byte{
	0x0a, 0x31, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x19, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x4a, 0x0a, 0x1a, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x2d, 0x0a,
	0x19, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x4a, 0x0a, 0x1a,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6d, 0x65,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x32, 0x91, 0x02, 0x0a, 0x0f, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7e, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x39, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7e, 0x0a, 0x03,
	0x53, 0x65, 0x74, 0x12, 0x39, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xa6, 0x02, 0x0a,
	0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x42, 0x10, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x70, 0x63,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x6f, 0x72,
	0x67, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x49,
	0x52, 0x4d, 0xaa, 0x02, 0x1e, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x52, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x1e, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2a, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x21, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x3a, 0x3a, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_rawDescOnce sync.Once
	file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_rawDescData = file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_rawDesc
)

func file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_rawDescGZIP() []byte {
	file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_rawDescOnce.Do(func() {
		file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_rawDescData)
	})
	return file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_rawDescData
}

var file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_goTypes = []interface{}{
	(*MetadataServiceGetRequest)(nil),  // 0: infrastructure.rpc.metadata.v1.MetadataServiceGetRequest
	(*MetadataServiceGetResponse)(nil), // 1: infrastructure.rpc.metadata.v1.MetadataServiceGetResponse
	(*MetadataServiceSetRequest)(nil),  // 2: infrastructure.rpc.metadata.v1.MetadataServiceSetRequest
	(*MetadataServiceSetResponse)(nil), // 3: infrastructure.rpc.metadata.v1.MetadataServiceSetResponse
	(*v1.Meta)(nil),                    // 4: domain.metadata.v1.Meta
}
var file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_depIdxs = []int32{
	4, // 0: infrastructure.rpc.metadata.v1.MetadataServiceGetResponse.meta:type_name -> domain.metadata.v1.Meta
	4, // 1: infrastructure.rpc.metadata.v1.MetadataServiceSetResponse.meta:type_name -> domain.metadata.v1.Meta
	0, // 2: infrastructure.rpc.metadata.v1.MetadataService.Get:input_type -> infrastructure.rpc.metadata.v1.MetadataServiceGetRequest
	2, // 3: infrastructure.rpc.metadata.v1.MetadataService.Set:input_type -> infrastructure.rpc.metadata.v1.MetadataServiceSetRequest
	1, // 4: infrastructure.rpc.metadata.v1.MetadataService.Get:output_type -> infrastructure.rpc.metadata.v1.MetadataServiceGetResponse
	3, // 5: infrastructure.rpc.metadata.v1.MetadataService.Set:output_type -> infrastructure.rpc.metadata.v1.MetadataServiceSetResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_init() }
func file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_init() {
	if File_infrastructure_rpc_metadata_v1_metadata_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataServiceGetRequest); i {
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
		file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataServiceGetResponse); i {
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
		file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataServiceSetRequest); i {
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
		file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataServiceSetResponse); i {
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
			RawDescriptor: file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_goTypes,
		DependencyIndexes: file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_depIdxs,
		MessageInfos:      file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_msgTypes,
	}.Build()
	File_infrastructure_rpc_metadata_v1_metadata_rpc_proto = out.File
	file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_rawDesc = nil
	file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_goTypes = nil
	file_infrastructure_rpc_metadata_v1_metadata_rpc_proto_depIdxs = nil
}

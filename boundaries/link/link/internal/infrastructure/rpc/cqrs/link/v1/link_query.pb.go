// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: infrastructure/rpc/cqrs/link/v1/link_query.proto

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

// GetRequest is the request message for the Get method.
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hash is the hash of the Link to get.
	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

// GetResponse is the response message for the Get method.
type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Link is the LinkView for the given hash.
	Link *LinkView `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetLink() *LinkView {
	if x != nil {
		return x.Link
	}
	return nil
}

// ListRequest is the request message for the List method.
type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Filter is the filter to apply to the Links.
	Filter string `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescGZIP(), []int{2}
}

func (x *ListRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

// ListResponse is the response message for the List method.
type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Links is the LinksView for the given filter.
	Links *LinksView `protobuf:"bytes,1,opt,name=links,proto3" json:"links,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescGZIP(), []int{3}
}

func (x *ListResponse) GetLinks() *LinksView {
	if x != nil {
		return x.Links
	}
	return nil
}

var File_infrastructure_rpc_cqrs_link_v1_link_query_proto protoreflect.FileDescriptor

var file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDesc = []byte{
	0x0a, 0x30, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x71, 0x72, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x71, 0x72, 0x73, 0x2f, 0x6c, 0x69, 0x6e,
	0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x20, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x22, 0x4c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3d, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x56, 0x69, 0x65, 0x77, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x22,
	0x25, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x50, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e,
	0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x56, 0x69, 0x65,
	0x77, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x32, 0xdd, 0x01, 0x0a, 0x10, 0x4c, 0x69, 0x6e,
	0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x2b, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x65, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63,
	0x71, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x71, 0x72,
	0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xb3, 0x02, 0x0a, 0x23, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x63, 0x71, 0x72, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31,
	0x42, 0x0e, 0x4c, 0x69, 0x6e, 0x6b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x69, 0x65,
	0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0xa2,
	0x02, 0x04, 0x49, 0x52, 0x43, 0x4c, 0xaa, 0x02, 0x1f, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x52, 0x70, 0x63, 0x2e, 0x43, 0x71, 0x72, 0x73,
	0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1f, 0x49, 0x6e, 0x66, 0x72, 0x61,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x43, 0x71,
	0x72, 0x73, 0x5c, 0x4c, 0x69, 0x6e, 0x6b, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2b, 0x49, 0x6e, 0x66,
	0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x52, 0x70, 0x63, 0x5c,
	0x43, 0x71, 0x72, 0x73, 0x5c, 0x4c, 0x69, 0x6e, 0x6b, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x23, 0x49, 0x6e, 0x66, 0x72, 0x61,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x3a, 0x3a,
	0x43, 0x71, 0x72, 0x73, 0x3a, 0x3a, 0x4c, 0x69, 0x6e, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescOnce sync.Once
	file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescData = file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDesc
)

func file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescGZIP() []byte {
	file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescOnce.Do(func() {
		file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescData)
	})
	return file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDescData
}

var file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_infrastructure_rpc_cqrs_link_v1_link_query_proto_goTypes = []interface{}{
	(*GetRequest)(nil),   // 0: infrastructure.rpc.cqrs.link.v1.GetRequest
	(*GetResponse)(nil),  // 1: infrastructure.rpc.cqrs.link.v1.GetResponse
	(*ListRequest)(nil),  // 2: infrastructure.rpc.cqrs.link.v1.ListRequest
	(*ListResponse)(nil), // 3: infrastructure.rpc.cqrs.link.v1.ListResponse
	(*LinkView)(nil),     // 4: infrastructure.rpc.cqrs.link.v1.LinkView
	(*LinksView)(nil),    // 5: infrastructure.rpc.cqrs.link.v1.LinksView
}
var file_infrastructure_rpc_cqrs_link_v1_link_query_proto_depIdxs = []int32{
	4, // 0: infrastructure.rpc.cqrs.link.v1.GetResponse.link:type_name -> infrastructure.rpc.cqrs.link.v1.LinkView
	5, // 1: infrastructure.rpc.cqrs.link.v1.ListResponse.links:type_name -> infrastructure.rpc.cqrs.link.v1.LinksView
	0, // 2: infrastructure.rpc.cqrs.link.v1.LinkQueryService.Get:input_type -> infrastructure.rpc.cqrs.link.v1.GetRequest
	2, // 3: infrastructure.rpc.cqrs.link.v1.LinkQueryService.List:input_type -> infrastructure.rpc.cqrs.link.v1.ListRequest
	1, // 4: infrastructure.rpc.cqrs.link.v1.LinkQueryService.Get:output_type -> infrastructure.rpc.cqrs.link.v1.GetResponse
	3, // 5: infrastructure.rpc.cqrs.link.v1.LinkQueryService.List:output_type -> infrastructure.rpc.cqrs.link.v1.ListResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_infrastructure_rpc_cqrs_link_v1_link_query_proto_init() }
func file_infrastructure_rpc_cqrs_link_v1_link_query_proto_init() {
	if File_infrastructure_rpc_cqrs_link_v1_link_query_proto != nil {
		return
	}
	file_infrastructure_rpc_cqrs_link_v1_link_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infrastructure_rpc_cqrs_link_v1_link_query_proto_goTypes,
		DependencyIndexes: file_infrastructure_rpc_cqrs_link_v1_link_query_proto_depIdxs,
		MessageInfos:      file_infrastructure_rpc_cqrs_link_v1_link_query_proto_msgTypes,
	}.Build()
	File_infrastructure_rpc_cqrs_link_v1_link_query_proto = out.File
	file_infrastructure_rpc_cqrs_link_v1_link_query_proto_rawDesc = nil
	file_infrastructure_rpc_cqrs_link_v1_link_query_proto_goTypes = nil
	file_infrastructure_rpc_cqrs_link_v1_link_query_proto_depIdxs = nil
}

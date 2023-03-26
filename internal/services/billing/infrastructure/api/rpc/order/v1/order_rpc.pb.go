// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: infrastructure/api/rpc/order/v1/order_rpc.proto

package order_rpc

import (
	v1 "github.com/shortlink-org/shortlink/internal/services/billing/domain/billing/order/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderHistoryResponse struct {
	state         protoimpl.MessageState
	unknownFields protoimpl.UnknownFields
	List          []*v1.Order `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	sizeCache     protoimpl.SizeCache
}

func (x *OrderHistoryResponse) Reset() {
	*x = OrderHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderHistoryResponse) ProtoMessage() {}

func (x *OrderHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderHistoryResponse.ProtoReflect.Descriptor instead.
func (*OrderHistoryResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *OrderHistoryResponse) GetList() []*v1.Order {
	if x != nil {
		return x.List
	}
	return nil
}

type OrderCreateRequest struct {
	state         protoimpl.MessageState
	Order         *v1.Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderCreateRequest) Reset() {
	*x = OrderCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateRequest) ProtoMessage() {}

func (x *OrderCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateRequest.ProtoReflect.Descriptor instead.
func (*OrderCreateRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *OrderCreateRequest) GetOrder() *v1.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type OrderCreateResponse struct {
	state         protoimpl.MessageState
	Order         *v1.Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderCreateResponse) Reset() {
	*x = OrderCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreateResponse) ProtoMessage() {}

func (x *OrderCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreateResponse.ProtoReflect.Descriptor instead.
func (*OrderCreateResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *OrderCreateResponse) GetOrder() *v1.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type OrderUpdateRequest struct {
	state         protoimpl.MessageState
	Order         *v1.Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderUpdateRequest) Reset() {
	*x = OrderUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderUpdateRequest) ProtoMessage() {}

func (x *OrderUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderUpdateRequest.ProtoReflect.Descriptor instead.
func (*OrderUpdateRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *OrderUpdateRequest) GetOrder() *v1.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type OrderUpdateResponse struct {
	state         protoimpl.MessageState
	Order         *v1.Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderUpdateResponse) Reset() {
	*x = OrderUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderUpdateResponse) ProtoMessage() {}

func (x *OrderUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderUpdateResponse.ProtoReflect.Descriptor instead.
func (*OrderUpdateResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *OrderUpdateResponse) GetOrder() *v1.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type OrderCloseRequest struct {
	state         protoimpl.MessageState
	Order         *v1.Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderCloseRequest) Reset() {
	*x = OrderCloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCloseRequest) ProtoMessage() {}

func (x *OrderCloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCloseRequest.ProtoReflect.Descriptor instead.
func (*OrderCloseRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *OrderCloseRequest) GetOrder() *v1.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type OrderCloseResponse struct {
	state         protoimpl.MessageState
	Order         *v1.Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderCloseResponse) Reset() {
	*x = OrderCloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCloseResponse) ProtoMessage() {}

func (x *OrderCloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCloseResponse.ProtoReflect.Descriptor instead.
func (*OrderCloseResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *OrderCloseResponse) GetOrder() *v1.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type OrderApproveRequest struct {
	state         protoimpl.MessageState
	Order         *v1.Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderApproveRequest) Reset() {
	*x = OrderApproveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderApproveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderApproveRequest) ProtoMessage() {}

func (x *OrderApproveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderApproveRequest.ProtoReflect.Descriptor instead.
func (*OrderApproveRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *OrderApproveRequest) GetOrder() *v1.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type OrderApproveResponse struct {
	state         protoimpl.MessageState
	Order         *v1.Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OrderApproveResponse) Reset() {
	*x = OrderApproveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderApproveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderApproveResponse) ProtoMessage() {}

func (x *OrderApproveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderApproveResponse.ProtoReflect.Descriptor instead.
func (*OrderApproveResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescGZIP(), []int{8}
}

func (x *OrderApproveResponse) GetOrder() *v1.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

var File_infrastructure_api_rpc_order_v1_order_rpc_proto protoreflect.FileDescriptor

var file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x23, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x22, 0x4a, 0x0a, 0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x13,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x4a, 0x0a, 0x12, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x34, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x22, 0x49, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x4a, 0x0a,
	0x12, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x13, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x34, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x4c, 0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x32, 0xdf, 0x04, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5f, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x35, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7a, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x33, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x7a, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x33, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77,
	0x0a, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x32, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x33, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7d, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x34, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x61, 0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescOnce sync.Once
	file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescData = file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDesc
)

func file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescGZIP() []byte {
	file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescOnce.Do(func() {
		file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescData)
	})
	return file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDescData
}

var file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_infrastructure_api_rpc_order_v1_order_rpc_proto_goTypes = []interface{}{
	(*OrderHistoryResponse)(nil), // 0: infrastructure.api.rpc.order.v1.OrderHistoryResponse
	(*OrderCreateRequest)(nil),   // 1: infrastructure.api.rpc.order.v1.OrderCreateRequest
	(*OrderCreateResponse)(nil),  // 2: infrastructure.api.rpc.order.v1.OrderCreateResponse
	(*OrderUpdateRequest)(nil),   // 3: infrastructure.api.rpc.order.v1.OrderUpdateRequest
	(*OrderUpdateResponse)(nil),  // 4: infrastructure.api.rpc.order.v1.OrderUpdateResponse
	(*OrderCloseRequest)(nil),    // 5: infrastructure.api.rpc.order.v1.OrderCloseRequest
	(*OrderCloseResponse)(nil),   // 6: infrastructure.api.rpc.order.v1.OrderCloseResponse
	(*OrderApproveRequest)(nil),  // 7: infrastructure.api.rpc.order.v1.OrderApproveRequest
	(*OrderApproveResponse)(nil), // 8: infrastructure.api.rpc.order.v1.OrderApproveResponse
	(*v1.Order)(nil),             // 9: domain.billing.order.v1.Order
	(*emptypb.Empty)(nil),        // 10: google.protobuf.Empty
}
var file_infrastructure_api_rpc_order_v1_order_rpc_proto_depIdxs = []int32{
	9,  // 0: infrastructure.api.rpc.order.v1.OrderHistoryResponse.list:type_name -> domain.billing.order.v1.Order
	9,  // 1: infrastructure.api.rpc.order.v1.OrderCreateRequest.order:type_name -> domain.billing.order.v1.Order
	9,  // 2: infrastructure.api.rpc.order.v1.OrderCreateResponse.order:type_name -> domain.billing.order.v1.Order
	9,  // 3: infrastructure.api.rpc.order.v1.OrderUpdateRequest.order:type_name -> domain.billing.order.v1.Order
	9,  // 4: infrastructure.api.rpc.order.v1.OrderUpdateResponse.order:type_name -> domain.billing.order.v1.Order
	9,  // 5: infrastructure.api.rpc.order.v1.OrderCloseRequest.order:type_name -> domain.billing.order.v1.Order
	9,  // 6: infrastructure.api.rpc.order.v1.OrderCloseResponse.order:type_name -> domain.billing.order.v1.Order
	9,  // 7: infrastructure.api.rpc.order.v1.OrderApproveRequest.order:type_name -> domain.billing.order.v1.Order
	9,  // 8: infrastructure.api.rpc.order.v1.OrderApproveResponse.order:type_name -> domain.billing.order.v1.Order
	10, // 9: infrastructure.api.rpc.order.v1.OrderService.OrderHistory:input_type -> google.protobuf.Empty
	1,  // 10: infrastructure.api.rpc.order.v1.OrderService.OrderCreate:input_type -> infrastructure.api.rpc.order.v1.OrderCreateRequest
	3,  // 11: infrastructure.api.rpc.order.v1.OrderService.OrderUpdate:input_type -> infrastructure.api.rpc.order.v1.OrderUpdateRequest
	5,  // 12: infrastructure.api.rpc.order.v1.OrderService.OrderClose:input_type -> infrastructure.api.rpc.order.v1.OrderCloseRequest
	7,  // 13: infrastructure.api.rpc.order.v1.OrderService.OrderApprove:input_type -> infrastructure.api.rpc.order.v1.OrderApproveRequest
	0,  // 14: infrastructure.api.rpc.order.v1.OrderService.OrderHistory:output_type -> infrastructure.api.rpc.order.v1.OrderHistoryResponse
	2,  // 15: infrastructure.api.rpc.order.v1.OrderService.OrderCreate:output_type -> infrastructure.api.rpc.order.v1.OrderCreateResponse
	4,  // 16: infrastructure.api.rpc.order.v1.OrderService.OrderUpdate:output_type -> infrastructure.api.rpc.order.v1.OrderUpdateResponse
	6,  // 17: infrastructure.api.rpc.order.v1.OrderService.OrderClose:output_type -> infrastructure.api.rpc.order.v1.OrderCloseResponse
	8,  // 18: infrastructure.api.rpc.order.v1.OrderService.OrderApprove:output_type -> infrastructure.api.rpc.order.v1.OrderApproveResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_infrastructure_api_rpc_order_v1_order_rpc_proto_init() }
func file_infrastructure_api_rpc_order_v1_order_rpc_proto_init() {
	if File_infrastructure_api_rpc_order_v1_order_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderHistoryResponse); i {
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
		file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateRequest); i {
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
		file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreateResponse); i {
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
		file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderUpdateRequest); i {
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
		file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderUpdateResponse); i {
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
		file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCloseRequest); i {
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
		file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCloseResponse); i {
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
		file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderApproveRequest); i {
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
		file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderApproveResponse); i {
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
			RawDescriptor: file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infrastructure_api_rpc_order_v1_order_rpc_proto_goTypes,
		DependencyIndexes: file_infrastructure_api_rpc_order_v1_order_rpc_proto_depIdxs,
		MessageInfos:      file_infrastructure_api_rpc_order_v1_order_rpc_proto_msgTypes,
	}.Build()
	File_infrastructure_api_rpc_order_v1_order_rpc_proto = out.File
	file_infrastructure_api_rpc_order_v1_order_rpc_proto_rawDesc = nil
	file_infrastructure_api_rpc_order_v1_order_rpc_proto_goTypes = nil
	file_infrastructure_api_rpc_order_v1_order_rpc_proto_depIdxs = nil
}

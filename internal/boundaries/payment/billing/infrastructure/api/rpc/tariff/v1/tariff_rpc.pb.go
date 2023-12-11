// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: infrastructure/api/rpc/tariff/v1/tariff_rpc.proto

package tariff_rpc

import (
	v1 "github.com/shortlink-org/shortlink/internal/boundaries/payment/billing/domain/billing/tariff/v1"
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

// TariffRequest is the tariff request message.
type TariffRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tariff is the tariff.
	Tariff *v1.Tariff `protobuf:"bytes,1,opt,name=tariff,proto3" json:"tariff,omitempty"`
}

func (x *TariffRequest) Reset() {
	*x = TariffRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TariffRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TariffRequest) ProtoMessage() {}

func (x *TariffRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TariffRequest.ProtoReflect.Descriptor instead.
func (*TariffRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *TariffRequest) GetTariff() *v1.Tariff {
	if x != nil {
		return x.Tariff
	}
	return nil
}

// TariffResponse is the tariff response message.
type TariffResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tariff is the tariff.
	Tariff *v1.Tariff `protobuf:"bytes,1,opt,name=tariff,proto3" json:"tariff,omitempty"`
}

func (x *TariffResponse) Reset() {
	*x = TariffResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TariffResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TariffResponse) ProtoMessage() {}

func (x *TariffResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TariffResponse.ProtoReflect.Descriptor instead.
func (*TariffResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *TariffResponse) GetTariff() *v1.Tariff {
	if x != nil {
		return x.Tariff
	}
	return nil
}

// TariffsResponse is the tariffs response message.
type TariffsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tariff is the tariff.
	List []*v1.Tariff `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *TariffsResponse) Reset() {
	*x = TariffsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TariffsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TariffsResponse) ProtoMessage() {}

func (x *TariffsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TariffsResponse.ProtoReflect.Descriptor instead.
func (*TariffsResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *TariffsResponse) GetList() []*v1.Tariff {
	if x != nil {
		return x.List
	}
	return nil
}

// TariffCreateRequest is the tariff create request message.
type TariffCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tariff is the tariff.
	Tariff *v1.Tariff `protobuf:"bytes,1,opt,name=tariff,proto3" json:"tariff,omitempty"`
}

func (x *TariffCreateRequest) Reset() {
	*x = TariffCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TariffCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TariffCreateRequest) ProtoMessage() {}

func (x *TariffCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TariffCreateRequest.ProtoReflect.Descriptor instead.
func (*TariffCreateRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *TariffCreateRequest) GetTariff() *v1.Tariff {
	if x != nil {
		return x.Tariff
	}
	return nil
}

// TariffCreateResponse is the tariff create response message.
type TariffCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tariff is the tariff.
	Tariff *v1.Tariff `protobuf:"bytes,1,opt,name=tariff,proto3" json:"tariff,omitempty"`
}

func (x *TariffCreateResponse) Reset() {
	*x = TariffCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TariffCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TariffCreateResponse) ProtoMessage() {}

func (x *TariffCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TariffCreateResponse.ProtoReflect.Descriptor instead.
func (*TariffCreateResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *TariffCreateResponse) GetTariff() *v1.Tariff {
	if x != nil {
		return x.Tariff
	}
	return nil
}

// TariffUpdateRequest is the tariff update request message.
type TariffUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tariff is the tariff.
	Tariff *v1.Tariff `protobuf:"bytes,1,opt,name=tariff,proto3" json:"tariff,omitempty"`
}

func (x *TariffUpdateRequest) Reset() {
	*x = TariffUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TariffUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TariffUpdateRequest) ProtoMessage() {}

func (x *TariffUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TariffUpdateRequest.ProtoReflect.Descriptor instead.
func (*TariffUpdateRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *TariffUpdateRequest) GetTariff() *v1.Tariff {
	if x != nil {
		return x.Tariff
	}
	return nil
}

// TariffUpdateResponse is the tariff update response message.
type TariffUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tariff is the tariff.
	Tariff *v1.Tariff `protobuf:"bytes,1,opt,name=tariff,proto3" json:"tariff,omitempty"`
}

func (x *TariffUpdateResponse) Reset() {
	*x = TariffUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TariffUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TariffUpdateResponse) ProtoMessage() {}

func (x *TariffUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TariffUpdateResponse.ProtoReflect.Descriptor instead.
func (*TariffUpdateResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *TariffUpdateResponse) GetTariff() *v1.Tariff {
	if x != nil {
		return x.Tariff
	}
	return nil
}

// TariffCloseRequest is the tariff close request message.
type TariffCloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tariff is the tariff.
	Tariff *v1.Tariff `protobuf:"bytes,1,opt,name=tariff,proto3" json:"tariff,omitempty"`
}

func (x *TariffCloseRequest) Reset() {
	*x = TariffCloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TariffCloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TariffCloseRequest) ProtoMessage() {}

func (x *TariffCloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TariffCloseRequest.ProtoReflect.Descriptor instead.
func (*TariffCloseRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *TariffCloseRequest) GetTariff() *v1.Tariff {
	if x != nil {
		return x.Tariff
	}
	return nil
}

// TariffCloseResponse is the tariff close response message.
type TariffCloseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tariff is the tariff.
	Tariff *v1.Tariff `protobuf:"bytes,1,opt,name=tariff,proto3" json:"tariff,omitempty"`
}

func (x *TariffCloseResponse) Reset() {
	*x = TariffCloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TariffCloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TariffCloseResponse) ProtoMessage() {}

func (x *TariffCloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TariffCloseResponse.ProtoReflect.Descriptor instead.
func (*TariffCloseResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescGZIP(), []int{8}
}

func (x *TariffCloseResponse) GetTariff() *v1.Tariff {
	if x != nil {
		return x.Tariff
	}
	return nil
}

var File_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto protoreflect.FileDescriptor

var file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDesc = []byte{
	0x0a, 0x31, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x20, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x0d, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x74, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x06, 0x74, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x22, 0x4a, 0x0a, 0x0e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x22, 0x47, 0x0a, 0x0f, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x13, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x38, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x52, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x22, 0x50, 0x0a, 0x14, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x52, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x22, 0x4f, 0x0a, 0x13,
	0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x22, 0x50, 0x0a,
	0x14, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x22,
	0x4e, 0x0a, 0x12, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x22,
	0x4f, 0x0a, 0x13, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x06, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x32, 0xd6, 0x04, 0x0a, 0x0d, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x6d, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x12, 0x2f, 0x2e, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x56, 0x0a, 0x07, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x0c, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x35, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72,
	0x69, 0x66, 0x66, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x36, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7f, 0x0a, 0x0c, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x35, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x36, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7c, 0x0a, 0x0b, 0x54,
	0x61, 0x72, 0x69, 0x66, 0x66, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x34, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x72, 0x69, 0x66, 0x66, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x35, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xaa, 0x02, 0x0a, 0x24, 0x63, 0x6f,
	0x6d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e,
	0x76, 0x31, 0x42, 0x0e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x52, 0x70, 0x63, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x5f,
	0x72, 0x70, 0x63, 0xa2, 0x02, 0x04, 0x49, 0x41, 0x52, 0x54, 0xaa, 0x02, 0x20, 0x49, 0x6e, 0x66,
	0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x2e,
	0x52, 0x70, 0x63, 0x2e, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x20,
	0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x41,
	0x70, 0x69, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x2c, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x54, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x24, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x3a, 0x3a, 0x54, 0x61, 0x72, 0x69,
	0x66, 0x66, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescOnce sync.Once
	file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescData = file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDesc
)

func file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescGZIP() []byte {
	file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescOnce.Do(func() {
		file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescData)
	})
	return file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDescData
}

var file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_goTypes = []interface{}{
	(*TariffRequest)(nil),        // 0: infrastructure.api.rpc.tariff.v1.TariffRequest
	(*TariffResponse)(nil),       // 1: infrastructure.api.rpc.tariff.v1.TariffResponse
	(*TariffsResponse)(nil),      // 2: infrastructure.api.rpc.tariff.v1.TariffsResponse
	(*TariffCreateRequest)(nil),  // 3: infrastructure.api.rpc.tariff.v1.TariffCreateRequest
	(*TariffCreateResponse)(nil), // 4: infrastructure.api.rpc.tariff.v1.TariffCreateResponse
	(*TariffUpdateRequest)(nil),  // 5: infrastructure.api.rpc.tariff.v1.TariffUpdateRequest
	(*TariffUpdateResponse)(nil), // 6: infrastructure.api.rpc.tariff.v1.TariffUpdateResponse
	(*TariffCloseRequest)(nil),   // 7: infrastructure.api.rpc.tariff.v1.TariffCloseRequest
	(*TariffCloseResponse)(nil),  // 8: infrastructure.api.rpc.tariff.v1.TariffCloseResponse
	(*v1.Tariff)(nil),            // 9: domain.billing.tariff.v1.Tariff
	(*emptypb.Empty)(nil),        // 10: google.protobuf.Empty
}
var file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_depIdxs = []int32{
	9,  // 0: infrastructure.api.rpc.tariff.v1.TariffRequest.tariff:type_name -> domain.billing.tariff.v1.Tariff
	9,  // 1: infrastructure.api.rpc.tariff.v1.TariffResponse.tariff:type_name -> domain.billing.tariff.v1.Tariff
	9,  // 2: infrastructure.api.rpc.tariff.v1.TariffsResponse.list:type_name -> domain.billing.tariff.v1.Tariff
	9,  // 3: infrastructure.api.rpc.tariff.v1.TariffCreateRequest.tariff:type_name -> domain.billing.tariff.v1.Tariff
	9,  // 4: infrastructure.api.rpc.tariff.v1.TariffCreateResponse.tariff:type_name -> domain.billing.tariff.v1.Tariff
	9,  // 5: infrastructure.api.rpc.tariff.v1.TariffUpdateRequest.tariff:type_name -> domain.billing.tariff.v1.Tariff
	9,  // 6: infrastructure.api.rpc.tariff.v1.TariffUpdateResponse.tariff:type_name -> domain.billing.tariff.v1.Tariff
	9,  // 7: infrastructure.api.rpc.tariff.v1.TariffCloseRequest.tariff:type_name -> domain.billing.tariff.v1.Tariff
	9,  // 8: infrastructure.api.rpc.tariff.v1.TariffCloseResponse.tariff:type_name -> domain.billing.tariff.v1.Tariff
	0,  // 9: infrastructure.api.rpc.tariff.v1.TariffService.Tariff:input_type -> infrastructure.api.rpc.tariff.v1.TariffRequest
	10, // 10: infrastructure.api.rpc.tariff.v1.TariffService.Tariffs:input_type -> google.protobuf.Empty
	3,  // 11: infrastructure.api.rpc.tariff.v1.TariffService.TariffCreate:input_type -> infrastructure.api.rpc.tariff.v1.TariffCreateRequest
	5,  // 12: infrastructure.api.rpc.tariff.v1.TariffService.TariffUpdate:input_type -> infrastructure.api.rpc.tariff.v1.TariffUpdateRequest
	7,  // 13: infrastructure.api.rpc.tariff.v1.TariffService.TariffClose:input_type -> infrastructure.api.rpc.tariff.v1.TariffCloseRequest
	1,  // 14: infrastructure.api.rpc.tariff.v1.TariffService.Tariff:output_type -> infrastructure.api.rpc.tariff.v1.TariffResponse
	2,  // 15: infrastructure.api.rpc.tariff.v1.TariffService.Tariffs:output_type -> infrastructure.api.rpc.tariff.v1.TariffsResponse
	4,  // 16: infrastructure.api.rpc.tariff.v1.TariffService.TariffCreate:output_type -> infrastructure.api.rpc.tariff.v1.TariffCreateResponse
	6,  // 17: infrastructure.api.rpc.tariff.v1.TariffService.TariffUpdate:output_type -> infrastructure.api.rpc.tariff.v1.TariffUpdateResponse
	8,  // 18: infrastructure.api.rpc.tariff.v1.TariffService.TariffClose:output_type -> infrastructure.api.rpc.tariff.v1.TariffCloseResponse
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_init() }
func file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_init() {
	if File_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TariffRequest); i {
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
		file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TariffResponse); i {
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
		file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TariffsResponse); i {
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
		file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TariffCreateRequest); i {
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
		file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TariffCreateResponse); i {
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
		file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TariffUpdateRequest); i {
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
		file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TariffUpdateResponse); i {
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
		file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TariffCloseRequest); i {
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
		file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TariffCloseResponse); i {
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
			RawDescriptor: file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_goTypes,
		DependencyIndexes: file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_depIdxs,
		MessageInfos:      file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_msgTypes,
	}.Build()
	File_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto = out.File
	file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_rawDesc = nil
	file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_goTypes = nil
	file_infrastructure_api_rpc_tariff_v1_tariff_rpc_proto_depIdxs = nil
}

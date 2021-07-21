// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: infrastructure/api/rpc/balance/v1/balance_rpc.proto

package v1

import (
	v1 "github.com/batazor/shortlink/internal/services/billing/domain/billing/balance/v1"
	v11 "github.com/batazor/shortlink/internal/services/billing/domain/billing/payment/v1"
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

type BalanceUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance *v1.Balance `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *BalanceUpdateRequest) Reset() {
	*x = BalanceUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceUpdateRequest) ProtoMessage() {}

func (x *BalanceUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceUpdateRequest.ProtoReflect.Descriptor instead.
func (*BalanceUpdateRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *BalanceUpdateRequest) GetBalance() *v1.Balance {
	if x != nil {
		return x.Balance
	}
	return nil
}

type BalanceUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payment *v11.Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *BalanceUpdateResponse) Reset() {
	*x = BalanceUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceUpdateResponse) ProtoMessage() {}

func (x *BalanceUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceUpdateResponse.ProtoReflect.Descriptor instead.
func (*BalanceUpdateResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *BalanceUpdateResponse) GetPayment() *v11.Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

type BalanceHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*v1.Balance `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *BalanceHistoryResponse) Reset() {
	*x = BalanceHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceHistoryResponse) ProtoMessage() {}

func (x *BalanceHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceHistoryResponse.ProtoReflect.Descriptor instead.
func (*BalanceHistoryResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *BalanceHistoryResponse) GetList() []*v1.Balance {
	if x != nil {
		return x.List
	}
	return nil
}

var File_infrastructure_api_rpc_balance_v1_balance_rpc_proto protoreflect.FileDescriptor

var file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_rawDesc = []byte{
	0x0a, 0x33, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x14, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3c, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x55, 0x0a,
	0x15, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x50, 0x0a, 0x16, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x32, 0xfe, 0x01, 0x0a, 0x0e, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x0e, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x39, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x84, 0x01, 0x0a, 0x0d, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x37, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x5a, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x61, 0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_rawDescOnce sync.Once
	file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_rawDescData = file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_rawDesc
)

func file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_rawDescGZIP() []byte {
	file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_rawDescOnce.Do(func() {
		file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_rawDescData)
	})
	return file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_rawDescData
}

var file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_goTypes = []interface{}{
	(*BalanceUpdateRequest)(nil),   // 0: infrastructure.api.rpc.balance.v1.BalanceUpdateRequest
	(*BalanceUpdateResponse)(nil),  // 1: infrastructure.api.rpc.balance.v1.BalanceUpdateResponse
	(*BalanceHistoryResponse)(nil), // 2: infrastructure.api.rpc.balance.v1.BalanceHistoryResponse
	(*v1.Balance)(nil),             // 3: domain.billing.balance.v1.Balance
	(*v11.Payment)(nil),            // 4: domain.billing.payment.v1.Payment
	(*emptypb.Empty)(nil),          // 5: google.protobuf.Empty
}
var file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_depIdxs = []int32{
	3, // 0: infrastructure.api.rpc.balance.v1.BalanceUpdateRequest.balance:type_name -> domain.billing.balance.v1.Balance
	4, // 1: infrastructure.api.rpc.balance.v1.BalanceUpdateResponse.payment:type_name -> domain.billing.payment.v1.Payment
	3, // 2: infrastructure.api.rpc.balance.v1.BalanceHistoryResponse.list:type_name -> domain.billing.balance.v1.Balance
	5, // 3: infrastructure.api.rpc.balance.v1.BalanceService.BalanceHistory:input_type -> google.protobuf.Empty
	0, // 4: infrastructure.api.rpc.balance.v1.BalanceService.BalanceUpdate:input_type -> infrastructure.api.rpc.balance.v1.BalanceUpdateRequest
	2, // 5: infrastructure.api.rpc.balance.v1.BalanceService.BalanceHistory:output_type -> infrastructure.api.rpc.balance.v1.BalanceHistoryResponse
	1, // 6: infrastructure.api.rpc.balance.v1.BalanceService.BalanceUpdate:output_type -> infrastructure.api.rpc.balance.v1.BalanceUpdateResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_init() }
func file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_init() {
	if File_infrastructure_api_rpc_balance_v1_balance_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceUpdateRequest); i {
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
		file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceUpdateResponse); i {
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
		file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceHistoryResponse); i {
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
			RawDescriptor: file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_goTypes,
		DependencyIndexes: file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_depIdxs,
		MessageInfos:      file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_msgTypes,
	}.Build()
	File_infrastructure_api_rpc_balance_v1_balance_rpc_proto = out.File
	file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_rawDesc = nil
	file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_goTypes = nil
	file_infrastructure_api_rpc_balance_v1_balance_rpc_proto_depIdxs = nil
}

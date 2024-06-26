// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: infrastructure/api/rpc/payment/v1/billing_rpc.proto

package payment_rpc

import (
	v1 "github.com/shortlink-org/shortlink/boundaries/billing/billing/internal/domain/payment/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Payment - information about payment
type Payment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask
	FieldMask *fieldmaskpb.FieldMask `protobuf:"bytes,6,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// ID payment
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name payment
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Status payment
	Status v1.StatusPayment `protobuf:"varint,3,opt,name=status,proto3,enum=domain.payment.v1.StatusPayment" json:"status,omitempty"`
	// User ID
	UserId string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Amount payment
	Amount int64 `protobuf:"varint,5,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Payment) Reset() {
	*x = Payment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payment) ProtoMessage() {}

func (x *Payment) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payment.ProtoReflect.Descriptor instead.
func (*Payment) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *Payment) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

func (x *Payment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Payment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Payment) GetStatus() v1.StatusPayment {
	if x != nil {
		return x.Status
	}
	return v1.StatusPayment(0)
}

func (x *Payment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Payment) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

// Payments - list payments
type Payments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of payments
	List []*Payment `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Payments) Reset() {
	*x = Payments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payments) ProtoMessage() {}

func (x *Payments) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payments.ProtoReflect.Descriptor instead.
func (*Payments) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *Payments) GetList() []*Payment {
	if x != nil {
		return x.List
	}
	return nil
}

// PaymentRequest is the request message for PaymentService.Payment.
type PaymentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID is the id of the payment.
	Payment *Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *PaymentRequest) Reset() {
	*x = PaymentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentRequest) ProtoMessage() {}

func (x *PaymentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentRequest.ProtoReflect.Descriptor instead.
func (*PaymentRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *PaymentRequest) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

// PaymentResponse is the response message for PaymentService.Payment.
type PaymentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Payment is the payment.
	Payment *Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *PaymentResponse) Reset() {
	*x = PaymentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentResponse) ProtoMessage() {}

func (x *PaymentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentResponse.ProtoReflect.Descriptor instead.
func (*PaymentResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *PaymentResponse) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

// PaymentsResponse is the response message for PaymentService.Payments.
type PaymentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Payments is the list of payments.
	List []*Payment `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PaymentsResponse) Reset() {
	*x = PaymentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentsResponse) ProtoMessage() {}

func (x *PaymentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentsResponse.ProtoReflect.Descriptor instead.
func (*PaymentsResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{4}
}

func (x *PaymentsResponse) GetList() []*Payment {
	if x != nil {
		return x.List
	}
	return nil
}

// PaymentCreateRequest is the request message for PaymentService.PaymentCreate.
type PaymentCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Payment is the payment.
	Payment *Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *PaymentCreateRequest) Reset() {
	*x = PaymentCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentCreateRequest) ProtoMessage() {}

func (x *PaymentCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentCreateRequest.ProtoReflect.Descriptor instead.
func (*PaymentCreateRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{5}
}

func (x *PaymentCreateRequest) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

// PaymentCreateResponse is the response message for PaymentService.PaymentCreate.
type PaymentCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Payment is the payment.
	Payment *Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *PaymentCreateResponse) Reset() {
	*x = PaymentCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentCreateResponse) ProtoMessage() {}

func (x *PaymentCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentCreateResponse.ProtoReflect.Descriptor instead.
func (*PaymentCreateResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{6}
}

func (x *PaymentCreateResponse) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

// PaymentCloseRequest is the request message for PaymentService.PaymentClose.
type PaymentCloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Payment is the payment.
	Payment *Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *PaymentCloseRequest) Reset() {
	*x = PaymentCloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentCloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentCloseRequest) ProtoMessage() {}

func (x *PaymentCloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentCloseRequest.ProtoReflect.Descriptor instead.
func (*PaymentCloseRequest) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{7}
}

func (x *PaymentCloseRequest) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

// PaymentCloseResponse is the response message for PaymentService.PaymentClose.
type PaymentCloseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Payment is the payment.
	Payment *Payment `protobuf:"bytes,1,opt,name=payment,proto3" json:"payment,omitempty"`
}

func (x *PaymentCloseResponse) Reset() {
	*x = PaymentCloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentCloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentCloseResponse) ProtoMessage() {}

func (x *PaymentCloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentCloseResponse.ProtoReflect.Descriptor instead.
func (*PaymentCloseResponse) Descriptor() ([]byte, []int) {
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP(), []int{8}
}

func (x *PaymentCloseResponse) GetPayment() *Payment {
	if x != nil {
		return x.Payment
	}
	return nil
}

var File_infrastructure_api_rpc_payment_v1_billing_rpc_proto protoreflect.FileDescriptor

var file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDesc = []byte{
	0x0a, 0x33, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61,
	0x73, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4a,
	0x0a, 0x08, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x56, 0x0a, 0x0e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x57, 0x0a, 0x0f, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x52, 0x0a, 0x10, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22,
	0x5c, 0x0a, 0x14, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5d, 0x0a,
	0x15, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5b, 0x0a, 0x13,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x5c, 0x0a, 0x14, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x44, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0xea, 0x03, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x72, 0x0a, 0x07, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59,
	0x0a, 0x08, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x33, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x84, 0x01, 0x0a, 0x0d, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x37, 0x2e, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x81, 0x01, 0x0a, 0x0c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x12, 0x36, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0xba, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0f,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x70, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x70, 0x63, 0xa2, 0x02, 0x04, 0x49, 0x41, 0x52,
	0x50, 0xaa, 0x02, 0x21, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x70, 0x63, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x21, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x2d, 0x49, 0x6e, 0x66, 0x72,
	0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x52,
	0x70, 0x63, 0x5c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x49, 0x6e, 0x66, 0x72,
	0x61, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x52, 0x70, 0x63, 0x3a, 0x3a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescOnce sync.Once
	file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescData = file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDesc
)

func file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescGZIP() []byte {
	file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescOnce.Do(func() {
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescData)
	})
	return file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDescData
}

var file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_goTypes = []interface{}{
	(*Payment)(nil),               // 0: infrastructure.api.rpc.payment.v1.Payment
	(*Payments)(nil),              // 1: infrastructure.api.rpc.payment.v1.Payments
	(*PaymentRequest)(nil),        // 2: infrastructure.api.rpc.payment.v1.PaymentRequest
	(*PaymentResponse)(nil),       // 3: infrastructure.api.rpc.payment.v1.PaymentResponse
	(*PaymentsResponse)(nil),      // 4: infrastructure.api.rpc.payment.v1.PaymentsResponse
	(*PaymentCreateRequest)(nil),  // 5: infrastructure.api.rpc.payment.v1.PaymentCreateRequest
	(*PaymentCreateResponse)(nil), // 6: infrastructure.api.rpc.payment.v1.PaymentCreateResponse
	(*PaymentCloseRequest)(nil),   // 7: infrastructure.api.rpc.payment.v1.PaymentCloseRequest
	(*PaymentCloseResponse)(nil),  // 8: infrastructure.api.rpc.payment.v1.PaymentCloseResponse
	(*fieldmaskpb.FieldMask)(nil), // 9: google.protobuf.FieldMask
	(v1.StatusPayment)(0),         // 10: domain.payment.v1.StatusPayment
	(*emptypb.Empty)(nil),         // 11: google.protobuf.Empty
}
var file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_depIdxs = []int32{
	9,  // 0: infrastructure.api.rpc.payment.v1.Payment.field_mask:type_name -> google.protobuf.FieldMask
	10, // 1: infrastructure.api.rpc.payment.v1.Payment.status:type_name -> domain.payment.v1.StatusPayment
	0,  // 2: infrastructure.api.rpc.payment.v1.Payments.list:type_name -> infrastructure.api.rpc.payment.v1.Payment
	0,  // 3: infrastructure.api.rpc.payment.v1.PaymentRequest.payment:type_name -> infrastructure.api.rpc.payment.v1.Payment
	0,  // 4: infrastructure.api.rpc.payment.v1.PaymentResponse.payment:type_name -> infrastructure.api.rpc.payment.v1.Payment
	0,  // 5: infrastructure.api.rpc.payment.v1.PaymentsResponse.list:type_name -> infrastructure.api.rpc.payment.v1.Payment
	0,  // 6: infrastructure.api.rpc.payment.v1.PaymentCreateRequest.payment:type_name -> infrastructure.api.rpc.payment.v1.Payment
	0,  // 7: infrastructure.api.rpc.payment.v1.PaymentCreateResponse.payment:type_name -> infrastructure.api.rpc.payment.v1.Payment
	0,  // 8: infrastructure.api.rpc.payment.v1.PaymentCloseRequest.payment:type_name -> infrastructure.api.rpc.payment.v1.Payment
	0,  // 9: infrastructure.api.rpc.payment.v1.PaymentCloseResponse.payment:type_name -> infrastructure.api.rpc.payment.v1.Payment
	2,  // 10: infrastructure.api.rpc.payment.v1.PaymentService.Payment:input_type -> infrastructure.api.rpc.payment.v1.PaymentRequest
	11, // 11: infrastructure.api.rpc.payment.v1.PaymentService.Payments:input_type -> google.protobuf.Empty
	5,  // 12: infrastructure.api.rpc.payment.v1.PaymentService.PaymentCreate:input_type -> infrastructure.api.rpc.payment.v1.PaymentCreateRequest
	7,  // 13: infrastructure.api.rpc.payment.v1.PaymentService.PaymentClose:input_type -> infrastructure.api.rpc.payment.v1.PaymentCloseRequest
	3,  // 14: infrastructure.api.rpc.payment.v1.PaymentService.Payment:output_type -> infrastructure.api.rpc.payment.v1.PaymentResponse
	4,  // 15: infrastructure.api.rpc.payment.v1.PaymentService.Payments:output_type -> infrastructure.api.rpc.payment.v1.PaymentsResponse
	6,  // 16: infrastructure.api.rpc.payment.v1.PaymentService.PaymentCreate:output_type -> infrastructure.api.rpc.payment.v1.PaymentCreateResponse
	8,  // 17: infrastructure.api.rpc.payment.v1.PaymentService.PaymentClose:output_type -> infrastructure.api.rpc.payment.v1.PaymentCloseResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_init() }
func file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_init() {
	if File_infrastructure_api_rpc_payment_v1_billing_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payment); i {
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
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payments); i {
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
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentRequest); i {
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
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentResponse); i {
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
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentsResponse); i {
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
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentCreateRequest); i {
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
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentCreateResponse); i {
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
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentCloseRequest); i {
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
		file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentCloseResponse); i {
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
			RawDescriptor: file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_goTypes,
		DependencyIndexes: file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_depIdxs,
		MessageInfos:      file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_msgTypes,
	}.Build()
	File_infrastructure_api_rpc_payment_v1_billing_rpc_proto = out.File
	file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_rawDesc = nil
	file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_goTypes = nil
	file_infrastructure_api_rpc_payment_v1_billing_rpc_proto_depIdxs = nil
}

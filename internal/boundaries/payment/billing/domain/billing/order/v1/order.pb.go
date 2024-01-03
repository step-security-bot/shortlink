// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: domain/billing/order/v1/order.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Order service
type StatusOrder int32

const (
	// unspecified
	StatusOrder_STATUS_ORDER_UNSPECIFIED StatusOrder = 0
	// create order
	StatusOrder_STATUS_ORDER_CREATE StatusOrder = 1
	// pending order
	StatusOrder_STATUS_ORDER_PENDING StatusOrder = 2
	// approve order
	StatusOrder_STATUS_ORDER_APPROVE StatusOrder = 3
	// paid order
	StatusOrder_STATUS_ORDER_PAID StatusOrder = 4
	// close order
	StatusOrder_STATUS_ORDER_CLOSE StatusOrder = 5
	// reject order
	StatusOrder_STATUS_ORDER_REJECT StatusOrder = 6
)

// Enum value maps for StatusOrder.
var (
	StatusOrder_name = map[int32]string{
		0: "STATUS_ORDER_UNSPECIFIED",
		1: "STATUS_ORDER_CREATE",
		2: "STATUS_ORDER_PENDING",
		3: "STATUS_ORDER_APPROVE",
		4: "STATUS_ORDER_PAID",
		5: "STATUS_ORDER_CLOSE",
		6: "STATUS_ORDER_REJECT",
	}
	StatusOrder_value = map[string]int32{
		"STATUS_ORDER_UNSPECIFIED": 0,
		"STATUS_ORDER_CREATE":      1,
		"STATUS_ORDER_PENDING":     2,
		"STATUS_ORDER_APPROVE":     3,
		"STATUS_ORDER_PAID":        4,
		"STATUS_ORDER_CLOSE":       5,
		"STATUS_ORDER_REJECT":      6,
	}
)

func (x StatusOrder) Enum() *StatusOrder {
	p := new(StatusOrder)
	*p = x
	return p
}

func (x StatusOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_domain_billing_order_v1_order_proto_enumTypes[0].Descriptor()
}

func (StatusOrder) Type() protoreflect.EnumType {
	return &file_domain_billing_order_v1_order_proto_enumTypes[0]
}

func (x StatusOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusOrder.Descriptor instead.
func (StatusOrder) EnumDescriptor() ([]byte, []int) {
	return file_domain_billing_order_v1_order_proto_rawDescGZIP(), []int{0}
}

// Event of order
type Event int32

const (
	// unspecified
	Event_EVENT_UNSPECIFIED Event = 0
	// new order
	Event_EVENT_ORDER_NEW Event = 1
	// update order
	Event_EVENT_ORDER_UPDATE Event = 2
	// approve order
	Event_EVENT_ORDER_APPROVED Event = 3
	// paid order
	Event_EVENT_ORDER_PAID Event = 4
	// close order
	Event_EVENT_ORDER_CLOSE Event = 5
	// reject order
	Event_EVENT_ORDER_REJECT Event = 6
)

// Enum value maps for Event.
var (
	Event_name = map[int32]string{
		0: "EVENT_UNSPECIFIED",
		1: "EVENT_ORDER_NEW",
		2: "EVENT_ORDER_UPDATE",
		3: "EVENT_ORDER_APPROVED",
		4: "EVENT_ORDER_PAID",
		5: "EVENT_ORDER_CLOSE",
		6: "EVENT_ORDER_REJECT",
	}
	Event_value = map[string]int32{
		"EVENT_UNSPECIFIED":    0,
		"EVENT_ORDER_NEW":      1,
		"EVENT_ORDER_UPDATE":   2,
		"EVENT_ORDER_APPROVED": 3,
		"EVENT_ORDER_PAID":     4,
		"EVENT_ORDER_CLOSE":    5,
		"EVENT_ORDER_REJECT":   6,
	}
)

func (x Event) Enum() *Event {
	p := new(Event)
	*p = x
	return p
}

func (x Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event) Descriptor() protoreflect.EnumDescriptor {
	return file_domain_billing_order_v1_order_proto_enumTypes[1].Descriptor()
}

func (Event) Type() protoreflect.EnumType {
	return &file_domain_billing_order_v1_order_proto_enumTypes[1]
}

func (x Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event.Descriptor instead.
func (Event) EnumDescriptor() ([]byte, []int) {
	return file_domain_billing_order_v1_order_proto_rawDescGZIP(), []int{1}
}

// Order service
type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask
	FieldMask *fieldmaskpb.FieldMask `protobuf:"bytes,5,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// order id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// user id
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// tariff id
	TariffId string `protobuf:"bytes,3,opt,name=tariff_id,json=tariffId,proto3" json:"tariff_id,omitempty"`
	// status order
	Status StatusOrder `protobuf:"varint,4,opt,name=status,proto3,enum=domain.billing.order.v1.StatusOrder" json:"status,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_billing_order_v1_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_domain_billing_order_v1_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_domain_billing_order_v1_order_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetTariffId() string {
	if x != nil {
		return x.TariffId
	}
	return ""
}

func (x *Order) GetStatus() StatusOrder {
	if x != nil {
		return x.Status
	}
	return StatusOrder_STATUS_ORDER_UNSPECIFIED
}

// Order list
type Orders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// order list
	List []*Order `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Orders) Reset() {
	*x = Orders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_billing_order_v1_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Orders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Orders) ProtoMessage() {}

func (x *Orders) ProtoReflect() protoreflect.Message {
	mi := &file_domain_billing_order_v1_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Orders.ProtoReflect.Descriptor instead.
func (*Orders) Descriptor() ([]byte, []int) {
	return file_domain_billing_order_v1_order_proto_rawDescGZIP(), []int{1}
}

func (x *Orders) GetList() []*Order {
	if x != nil {
		return x.List
	}
	return nil
}

var File_domain_billing_order_v1_order_proto protoreflect.FileDescriptor

var file_domain_billing_order_v1_order_proto_rawDesc = []byte{
	0x0a, 0x23, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc6, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3c, 0x0a, 0x06, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x32, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x2a, 0xc0, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f,
	0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x18,
	0x0a, 0x14, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45,
	0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x52, 0x44,
	0x45, 0x52, 0x5f, 0x50, 0x41, 0x49, 0x44, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10,
	0x05, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x06, 0x2a, 0xaa, 0x01, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x45,
	0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x01,
	0x12, 0x16, 0x0a, 0x12, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x44,
	0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x50, 0x41, 0x49, 0x44, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e,
	0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x05, 0x12,
	0x16, 0x0a, 0x12, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x52,
	0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x06, 0x42, 0x88, 0x02, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x6f, 0x72, 0x67, 0x2f,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x42, 0x4f, 0xaa, 0x02, 0x17, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x42,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x23, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x5c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x3a,
	0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_billing_order_v1_order_proto_rawDescOnce sync.Once
	file_domain_billing_order_v1_order_proto_rawDescData = file_domain_billing_order_v1_order_proto_rawDesc
)

func file_domain_billing_order_v1_order_proto_rawDescGZIP() []byte {
	file_domain_billing_order_v1_order_proto_rawDescOnce.Do(func() {
		file_domain_billing_order_v1_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_billing_order_v1_order_proto_rawDescData)
	})
	return file_domain_billing_order_v1_order_proto_rawDescData
}

var file_domain_billing_order_v1_order_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_domain_billing_order_v1_order_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_domain_billing_order_v1_order_proto_goTypes = []interface{}{
	(StatusOrder)(0),              // 0: domain.billing.order.v1.StatusOrder
	(Event)(0),                    // 1: domain.billing.order.v1.Event
	(*Order)(nil),                 // 2: domain.billing.order.v1.Order
	(*Orders)(nil),                // 3: domain.billing.order.v1.Orders
	(*fieldmaskpb.FieldMask)(nil), // 4: google.protobuf.FieldMask
}
var file_domain_billing_order_v1_order_proto_depIdxs = []int32{
	4, // 0: domain.billing.order.v1.Order.field_mask:type_name -> google.protobuf.FieldMask
	0, // 1: domain.billing.order.v1.Order.status:type_name -> domain.billing.order.v1.StatusOrder
	2, // 2: domain.billing.order.v1.Orders.list:type_name -> domain.billing.order.v1.Order
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_domain_billing_order_v1_order_proto_init() }
func file_domain_billing_order_v1_order_proto_init() {
	if File_domain_billing_order_v1_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_billing_order_v1_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_domain_billing_order_v1_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Orders); i {
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
			RawDescriptor: file_domain_billing_order_v1_order_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_billing_order_v1_order_proto_goTypes,
		DependencyIndexes: file_domain_billing_order_v1_order_proto_depIdxs,
		EnumInfos:         file_domain_billing_order_v1_order_proto_enumTypes,
		MessageInfos:      file_domain_billing_order_v1_order_proto_msgTypes,
	}.Build()
	File_domain_billing_order_v1_order_proto = out.File
	file_domain_billing_order_v1_order_proto_rawDesc = nil
	file_domain_billing_order_v1_order_proto_goTypes = nil
	file_domain_billing_order_v1_order_proto_depIdxs = nil
}

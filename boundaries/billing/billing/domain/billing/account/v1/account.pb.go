// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: domain/billing/account/v1/account.proto

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

// Event is a billing account event
type Event int32

const (
	// event unspecified
	Event_EVENT_UNSPECIFIED Event = 0
	// account new
	Event_EVENT_ACCOUNT_NEW Event = 1
	// account delete
	Event_EVENT_ACCOUNT_DELETE Event = 2
)

// Enum value maps for Event.
var (
	Event_name = map[int32]string{
		0: "EVENT_UNSPECIFIED",
		1: "EVENT_ACCOUNT_NEW",
		2: "EVENT_ACCOUNT_DELETE",
	}
	Event_value = map[string]int32{
		"EVENT_UNSPECIFIED":    0,
		"EVENT_ACCOUNT_NEW":    1,
		"EVENT_ACCOUNT_DELETE": 2,
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
	return file_domain_billing_account_v1_account_proto_enumTypes[0].Descriptor()
}

func (Event) Type() protoreflect.EnumType {
	return &file_domain_billing_account_v1_account_proto_enumTypes[0]
}

func (x Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event.Descriptor instead.
func (Event) EnumDescriptor() ([]byte, []int) {
	return file_domain_billing_account_v1_account_proto_rawDescGZIP(), []int{0}
}

// Account is a billing account
type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// field_mask
	FieldMask *fieldmaskpb.FieldMask `protobuf:"bytes,4,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// account id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// user id
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// tariff id
	TariffId string `protobuf:"bytes,3,opt,name=tariff_id,json=tariffId,proto3" json:"tariff_id,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_billing_account_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_domain_billing_account_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_domain_billing_account_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Account) GetTariffId() string {
	if x != nil {
		return x.TariffId
	}
	return ""
}

// Accounts is a list of billing accounts
type Accounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// accounts
	List []*Account `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Accounts) Reset() {
	*x = Accounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_billing_account_v1_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accounts) ProtoMessage() {}

func (x *Accounts) ProtoReflect() protoreflect.Message {
	mi := &file_domain_billing_account_v1_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accounts.ProtoReflect.Descriptor instead.
func (*Accounts) Descriptor() ([]byte, []int) {
	return file_domain_billing_account_v1_account_proto_rawDescGZIP(), []int{1}
}

func (x *Accounts) GetList() []*Account {
	if x != nil {
		return x.List
	}
	return nil
}

var File_domain_billing_account_v1_account_proto protoreflect.FileDescriptor

var file_domain_billing_account_v1_account_proto_rawDesc = []byte{
	0x0a, 0x27, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x69, 0x66, 0x66,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x69, 0x66,
	0x66, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12,
	0x36, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x2a, 0x4f, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x01, 0x12, 0x18,
	0x0a, 0x14, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x42, 0x8d, 0x02, 0x0a, 0x1d, 0x63, 0x6f, 0x6d,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b,
	0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x62,
	0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x42, 0x41, 0xaa, 0x02, 0x19, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x19, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x42,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x25, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x42, 0x69, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x5c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1c, 0x44, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x3a, 0x3a, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_billing_account_v1_account_proto_rawDescOnce sync.Once
	file_domain_billing_account_v1_account_proto_rawDescData = file_domain_billing_account_v1_account_proto_rawDesc
)

func file_domain_billing_account_v1_account_proto_rawDescGZIP() []byte {
	file_domain_billing_account_v1_account_proto_rawDescOnce.Do(func() {
		file_domain_billing_account_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_billing_account_v1_account_proto_rawDescData)
	})
	return file_domain_billing_account_v1_account_proto_rawDescData
}

var file_domain_billing_account_v1_account_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_domain_billing_account_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_domain_billing_account_v1_account_proto_goTypes = []interface{}{
	(Event)(0),                    // 0: domain.billing.account.v1.Event
	(*Account)(nil),               // 1: domain.billing.account.v1.Account
	(*Accounts)(nil),              // 2: domain.billing.account.v1.Accounts
	(*fieldmaskpb.FieldMask)(nil), // 3: google.protobuf.FieldMask
}
var file_domain_billing_account_v1_account_proto_depIdxs = []int32{
	3, // 0: domain.billing.account.v1.Account.field_mask:type_name -> google.protobuf.FieldMask
	1, // 1: domain.billing.account.v1.Accounts.list:type_name -> domain.billing.account.v1.Account
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_domain_billing_account_v1_account_proto_init() }
func file_domain_billing_account_v1_account_proto_init() {
	if File_domain_billing_account_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_billing_account_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_domain_billing_account_v1_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Accounts); i {
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
			RawDescriptor: file_domain_billing_account_v1_account_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_billing_account_v1_account_proto_goTypes,
		DependencyIndexes: file_domain_billing_account_v1_account_proto_depIdxs,
		EnumInfos:         file_domain_billing_account_v1_account_proto_enumTypes,
		MessageInfos:      file_domain_billing_account_v1_account_proto_msgTypes,
	}.Build()
	File_domain_billing_account_v1_account_proto = out.File
	file_domain_billing_account_v1_account_proto_rawDesc = nil
	file_domain_billing_account_v1_account_proto_goTypes = nil
	file_domain_billing_account_v1_account_proto_depIdxs = nil
}

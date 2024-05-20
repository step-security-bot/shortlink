// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: domain/payment/v1/payment.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// StatusPayment status payment
type StatusPayment int32

const (
	// Unspecified
	StatusPayment_STATUS_PAYMENT_UNSPECIFIED StatusPayment = 0
	// New payment
	StatusPayment_STATUS_PAYMENT_NEW StatusPayment = 1
	// Pending payment
	StatusPayment_STATUS_PAYMENT_PENDING StatusPayment = 2
	// Approve payment
	StatusPayment_STATUS_PAYMENT_APPROVE StatusPayment = 3
	// Close payment
	StatusPayment_STATUS_PAYMENT_CLOSE StatusPayment = 4
	// Reject payment
	StatusPayment_STATUS_PAYMENT_REJECT StatusPayment = 5
)

// Enum value maps for StatusPayment.
var (
	StatusPayment_name = map[int32]string{
		0: "STATUS_PAYMENT_UNSPECIFIED",
		1: "STATUS_PAYMENT_NEW",
		2: "STATUS_PAYMENT_PENDING",
		3: "STATUS_PAYMENT_APPROVE",
		4: "STATUS_PAYMENT_CLOSE",
		5: "STATUS_PAYMENT_REJECT",
	}
	StatusPayment_value = map[string]int32{
		"STATUS_PAYMENT_UNSPECIFIED": 0,
		"STATUS_PAYMENT_NEW":         1,
		"STATUS_PAYMENT_PENDING":     2,
		"STATUS_PAYMENT_APPROVE":     3,
		"STATUS_PAYMENT_CLOSE":       4,
		"STATUS_PAYMENT_REJECT":      5,
	}
)

func (x StatusPayment) Enum() *StatusPayment {
	p := new(StatusPayment)
	*p = x
	return p
}

func (x StatusPayment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusPayment) Descriptor() protoreflect.EnumDescriptor {
	return file_domain_payment_v1_payment_proto_enumTypes[0].Descriptor()
}

func (StatusPayment) Type() protoreflect.EnumType {
	return &file_domain_payment_v1_payment_proto_enumTypes[0]
}

func (x StatusPayment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusPayment.Descriptor instead.
func (StatusPayment) EnumDescriptor() ([]byte, []int) {
	return file_domain_payment_v1_payment_proto_rawDescGZIP(), []int{0}
}

var File_domain_payment_v1_payment_proto protoreflect.FileDescriptor

var file_domain_payment_v1_payment_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb4, 0x01, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x01,
	0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45,
	0x4e, 0x54, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41,
	0x50, 0x50, 0x52, 0x4f, 0x56, 0x45, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45,
	0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x41, 0x59,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x05, 0x42, 0xe5, 0x01,
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x6f, 0x72,
	0x67, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x44, 0x50, 0x58, 0xaa, 0x02, 0x11, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1d, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x13, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_payment_v1_payment_proto_rawDescOnce sync.Once
	file_domain_payment_v1_payment_proto_rawDescData = file_domain_payment_v1_payment_proto_rawDesc
)

func file_domain_payment_v1_payment_proto_rawDescGZIP() []byte {
	file_domain_payment_v1_payment_proto_rawDescOnce.Do(func() {
		file_domain_payment_v1_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_payment_v1_payment_proto_rawDescData)
	})
	return file_domain_payment_v1_payment_proto_rawDescData
}

var file_domain_payment_v1_payment_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_domain_payment_v1_payment_proto_goTypes = []interface{}{
	(StatusPayment)(0), // 0: domain.payment.v1.StatusPayment
}
var file_domain_payment_v1_payment_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_domain_payment_v1_payment_proto_init() }
func file_domain_payment_v1_payment_proto_init() {
	if File_domain_payment_v1_payment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_domain_payment_v1_payment_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_payment_v1_payment_proto_goTypes,
		DependencyIndexes: file_domain_payment_v1_payment_proto_depIdxs,
		EnumInfos:         file_domain_payment_v1_payment_proto_enumTypes,
	}.Build()
	File_domain_payment_v1_payment_proto = out.File
	file_domain_payment_v1_payment_proto_rawDesc = nil
	file_domain_payment_v1_payment_proto_goTypes = nil
	file_domain_payment_v1_payment_proto_depIdxs = nil
}

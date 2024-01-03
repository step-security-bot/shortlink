// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: domain/eventsourcing/v1/aggregate.proto

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

// BaseAggregate represents the basic information that all aggregates should have.
type BaseAggregate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Field mask indicating which fields have been updated.
	FieldMask *fieldmaskpb.FieldMask `protobuf:"bytes,5,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// The unique identifier of the aggregate.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The type of the aggregate.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// The version of the aggregate.
	Version int32 `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	// The changes that have been made to the aggregate, represented as a list of events.
	Changes []*Event `protobuf:"bytes,4,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (x *BaseAggregate) Reset() {
	*x = BaseAggregate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_eventsourcing_v1_aggregate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseAggregate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseAggregate) ProtoMessage() {}

func (x *BaseAggregate) ProtoReflect() protoreflect.Message {
	mi := &file_domain_eventsourcing_v1_aggregate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseAggregate.ProtoReflect.Descriptor instead.
func (*BaseAggregate) Descriptor() ([]byte, []int) {
	return file_domain_eventsourcing_v1_aggregate_proto_rawDescGZIP(), []int{0}
}

func (x *BaseAggregate) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

func (x *BaseAggregate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BaseAggregate) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *BaseAggregate) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *BaseAggregate) GetChanges() []*Event {
	if x != nil {
		return x.Changes
	}
	return nil
}

var File_domain_eventsourcing_v1_aggregate_proto protoreflect.FileDescriptor

var file_domain_eventsourcing_v1_aggregate_proto_rawDesc = []byte{
	0x0a, 0x27, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x0d, 0x42, 0x61,
	0x73, 0x65, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x42, 0xed,
	0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0e,
	0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x45, 0x58, 0xaa, 0x02, 0x17, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x17, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x19, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_eventsourcing_v1_aggregate_proto_rawDescOnce sync.Once
	file_domain_eventsourcing_v1_aggregate_proto_rawDescData = file_domain_eventsourcing_v1_aggregate_proto_rawDesc
)

func file_domain_eventsourcing_v1_aggregate_proto_rawDescGZIP() []byte {
	file_domain_eventsourcing_v1_aggregate_proto_rawDescOnce.Do(func() {
		file_domain_eventsourcing_v1_aggregate_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_eventsourcing_v1_aggregate_proto_rawDescData)
	})
	return file_domain_eventsourcing_v1_aggregate_proto_rawDescData
}

var file_domain_eventsourcing_v1_aggregate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_domain_eventsourcing_v1_aggregate_proto_goTypes = []interface{}{
	(*BaseAggregate)(nil),         // 0: domain.eventsourcing.v1.BaseAggregate
	(*fieldmaskpb.FieldMask)(nil), // 1: google.protobuf.FieldMask
	(*Event)(nil),                 // 2: domain.eventsourcing.v1.Event
}
var file_domain_eventsourcing_v1_aggregate_proto_depIdxs = []int32{
	1, // 0: domain.eventsourcing.v1.BaseAggregate.field_mask:type_name -> google.protobuf.FieldMask
	2, // 1: domain.eventsourcing.v1.BaseAggregate.changes:type_name -> domain.eventsourcing.v1.Event
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_domain_eventsourcing_v1_aggregate_proto_init() }
func file_domain_eventsourcing_v1_aggregate_proto_init() {
	if File_domain_eventsourcing_v1_aggregate_proto != nil {
		return
	}
	file_domain_eventsourcing_v1_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_domain_eventsourcing_v1_aggregate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseAggregate); i {
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
			RawDescriptor: file_domain_eventsourcing_v1_aggregate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_eventsourcing_v1_aggregate_proto_goTypes,
		DependencyIndexes: file_domain_eventsourcing_v1_aggregate_proto_depIdxs,
		MessageInfos:      file_domain_eventsourcing_v1_aggregate_proto_msgTypes,
	}.Build()
	File_domain_eventsourcing_v1_aggregate_proto = out.File
	file_domain_eventsourcing_v1_aggregate_proto_rawDesc = nil
	file_domain_eventsourcing_v1_aggregate_proto_goTypes = nil
	file_domain_eventsourcing_v1_aggregate_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: shortdb/domain/table/v1/stats.proto

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

type TableStats struct {
	state         protoimpl.MessageState
	unknownFields protoimpl.UnknownFields
	RowsCount     int64 `protobuf:"varint,1,opt,name=rows_count,json=rowsCount,proto3" json:"rows_count,omitempty"`
	sizeCache     protoimpl.SizeCache
	PageCount     int32 `protobuf:"varint,2,opt,name=page_count,json=pageCount,proto3" json:"page_count,omitempty"`
}

func (x *TableStats) Reset() {
	*x = TableStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortdb_domain_table_v1_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableStats) ProtoMessage() {}

func (x *TableStats) ProtoReflect() protoreflect.Message {
	mi := &file_shortdb_domain_table_v1_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableStats.ProtoReflect.Descriptor instead.
func (*TableStats) Descriptor() ([]byte, []int) {
	return file_shortdb_domain_table_v1_stats_proto_rawDescGZIP(), []int{0}
}

func (x *TableStats) GetRowsCount() int64 {
	if x != nil {
		return x.RowsCount
	}
	return 0
}

func (x *TableStats) GetPageCount() int32 {
	if x != nil {
		return x.PageCount
	}
	return 0
}

var File_shortdb_domain_table_v1_stats_proto protoreflect.FileDescriptor

var file_shortdb_domain_table_v1_stats_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x4a,
	0x0a, 0x0a, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x6f, 0x77, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x72, 0x6f, 0x77, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x61, 0x7a, 0x6f, 0x72,
	0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortdb_domain_table_v1_stats_proto_rawDescOnce sync.Once
	file_shortdb_domain_table_v1_stats_proto_rawDescData = file_shortdb_domain_table_v1_stats_proto_rawDesc
)

func file_shortdb_domain_table_v1_stats_proto_rawDescGZIP() []byte {
	file_shortdb_domain_table_v1_stats_proto_rawDescOnce.Do(func() {
		file_shortdb_domain_table_v1_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortdb_domain_table_v1_stats_proto_rawDescData)
	})
	return file_shortdb_domain_table_v1_stats_proto_rawDescData
}

var file_shortdb_domain_table_v1_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_shortdb_domain_table_v1_stats_proto_goTypes = []interface{}{
	(*TableStats)(nil), // 0: shortdb.domain.table.v1.TableStats
}
var file_shortdb_domain_table_v1_stats_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shortdb_domain_table_v1_stats_proto_init() }
func file_shortdb_domain_table_v1_stats_proto_init() {
	if File_shortdb_domain_table_v1_stats_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shortdb_domain_table_v1_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableStats); i {
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
			RawDescriptor: file_shortdb_domain_table_v1_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shortdb_domain_table_v1_stats_proto_goTypes,
		DependencyIndexes: file_shortdb_domain_table_v1_stats_proto_depIdxs,
		MessageInfos:      file_shortdb_domain_table_v1_stats_proto_msgTypes,
	}.Build()
	File_shortdb_domain_table_v1_stats_proto = out.File
	file_shortdb_domain_table_v1_stats_proto_rawDesc = nil
	file_shortdb_domain_table_v1_stats_proto_goTypes = nil
	file_shortdb_domain_table_v1_stats_proto_depIdxs = nil
}

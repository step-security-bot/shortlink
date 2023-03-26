// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: shortdb/domain/database/v1/database.proto

package v1

import (
	v1 "github.com/shortlink-org/shortlink/pkg/shortdb/domain/table/v1"
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

type DataBase struct {
	state         protoimpl.MessageState
	Tables        map[string]*v1.Table `protobuf:"bytes,2,rep,name=tables,proto3" json:"tables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Name          string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DataBase) Reset() {
	*x = DataBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortdb_domain_database_v1_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataBase) ProtoMessage() {}

func (x *DataBase) ProtoReflect() protoreflect.Message {
	mi := &file_shortdb_domain_database_v1_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataBase.ProtoReflect.Descriptor instead.
func (*DataBase) Descriptor() ([]byte, []int) {
	return file_shortdb_domain_database_v1_database_proto_rawDescGZIP(), []int{0}
}

func (x *DataBase) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataBase) GetTables() map[string]*v1.Table {
	if x != nil {
		return x.Tables
	}
	return nil
}

var File_shortdb_domain_database_v1_database_proto protoreflect.FileDescriptor

var file_shortdb_domain_database_v1_database_proto_rawDesc = []byte{
	0x0a, 0x29, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x23, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62,
	0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a,
	0x08, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a,
	0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x42,
	0x61, 0x73, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x1a, 0x59, 0x0a, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64,
	0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x61, 0x74, 0x61, 0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69,
	0x6e, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortdb_domain_database_v1_database_proto_rawDescOnce sync.Once
	file_shortdb_domain_database_v1_database_proto_rawDescData = file_shortdb_domain_database_v1_database_proto_rawDesc
)

func file_shortdb_domain_database_v1_database_proto_rawDescGZIP() []byte {
	file_shortdb_domain_database_v1_database_proto_rawDescOnce.Do(func() {
		file_shortdb_domain_database_v1_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortdb_domain_database_v1_database_proto_rawDescData)
	})
	return file_shortdb_domain_database_v1_database_proto_rawDescData
}

var file_shortdb_domain_database_v1_database_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_shortdb_domain_database_v1_database_proto_goTypes = []interface{}{
	(*DataBase)(nil), // 0: shortdb.domain.database.v1.DataBase
	nil,              // 1: shortdb.domain.database.v1.DataBase.TablesEntry
	(*v1.Table)(nil), // 2: shortdb.domain.table.v1.Table
}
var file_shortdb_domain_database_v1_database_proto_depIdxs = []int32{
	1, // 0: shortdb.domain.database.v1.DataBase.tables:type_name -> shortdb.domain.database.v1.DataBase.TablesEntry
	2, // 1: shortdb.domain.database.v1.DataBase.TablesEntry.value:type_name -> shortdb.domain.table.v1.Table
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_shortdb_domain_database_v1_database_proto_init() }
func file_shortdb_domain_database_v1_database_proto_init() {
	if File_shortdb_domain_database_v1_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shortdb_domain_database_v1_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataBase); i {
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
			RawDescriptor: file_shortdb_domain_database_v1_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shortdb_domain_database_v1_database_proto_goTypes,
		DependencyIndexes: file_shortdb_domain_database_v1_database_proto_depIdxs,
		MessageInfos:      file_shortdb_domain_database_v1_database_proto_msgTypes,
	}.Build()
	File_shortdb_domain_database_v1_database_proto = out.File
	file_shortdb_domain_database_v1_database_proto_rawDesc = nil
	file_shortdb_domain_database_v1_database_proto_goTypes = nil
	file_shortdb_domain_database_v1_database_proto_depIdxs = nil
}

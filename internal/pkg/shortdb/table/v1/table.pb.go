// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: shortdb/table/v1/table.proto

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

type Type int32

const (
	Type_TYPE_UNSPECIFIED Type = 0
	Type_TYPE_INTEGER     Type = 1
	Type_TYPE_STRING      Type = 2
	Type_TYPE_BOOLEAN     Type = 3
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_INTEGER",
		2: "TYPE_STRING",
		3: "TYPE_BOOLEAN",
	}
	Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_INTEGER":     1,
		"TYPE_STRING":      2,
		"TYPE_BOOLEAN":     3,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_shortdb_table_v1_table_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_shortdb_table_v1_table_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_shortdb_table_v1_table_proto_rawDescGZIP(), []int{0}
}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value map[string][]byte `protobuf:"bytes,2,rep,name=value,proto3" json:"value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortdb_table_v1_table_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_shortdb_table_v1_table_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_shortdb_table_v1_table_proto_rawDescGZIP(), []int{0}
}

func (x *Row) GetValue() map[string][]byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Fields map[string]Type `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=shortdb.table.v1.Type"`
	Rows   []*Row          `protobuf:"bytes,3,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *Table) Reset() {
	*x = Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortdb_table_v1_table_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Table) ProtoMessage() {}

func (x *Table) ProtoReflect() protoreflect.Message {
	mi := &file_shortdb_table_v1_table_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Table.ProtoReflect.Descriptor instead.
func (*Table) Descriptor() ([]byte, []int) {
	return file_shortdb_table_v1_table_proto_rawDescGZIP(), []int{1}
}

func (x *Table) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Table) GetFields() map[string]Type {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Table) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

type DataBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tables map[string]*Table `protobuf:"bytes,1,rep,name=tables,proto3" json:"tables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *DataBase) Reset() {
	*x = DataBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortdb_table_v1_table_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataBase) ProtoMessage() {}

func (x *DataBase) ProtoReflect() protoreflect.Message {
	mi := &file_shortdb_table_v1_table_proto_msgTypes[2]
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
	return file_shortdb_table_v1_table_proto_rawDescGZIP(), []int{2}
}

func (x *DataBase) GetTables() map[string]*Table {
	if x != nil {
		return x.Tables
	}
	return nil
}

var File_shortdb_table_v1_table_proto protoreflect.FileDescriptor

var file_shortdb_table_v1_table_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x22, 0x77, 0x0a, 0x03, 0x52, 0x6f, 0x77, 0x12, 0x36, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62,
	0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x77, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a,
	0x38, 0x0a, 0x0a, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd6, 0x01, 0x0a, 0x05, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64,
	0x62, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x1a,
	0x51, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x16, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x9e, 0x01, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x73, 0x65, 0x12,
	0x3e, 0x0a, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x1a,
	0x52, 0x0a, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x2a, 0x51, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x47, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f,
	0x4c, 0x45, 0x41, 0x4e, 0x10, 0x03, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x61, 0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2f, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortdb_table_v1_table_proto_rawDescOnce sync.Once
	file_shortdb_table_v1_table_proto_rawDescData = file_shortdb_table_v1_table_proto_rawDesc
)

func file_shortdb_table_v1_table_proto_rawDescGZIP() []byte {
	file_shortdb_table_v1_table_proto_rawDescOnce.Do(func() {
		file_shortdb_table_v1_table_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortdb_table_v1_table_proto_rawDescData)
	})
	return file_shortdb_table_v1_table_proto_rawDescData
}

var file_shortdb_table_v1_table_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shortdb_table_v1_table_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_shortdb_table_v1_table_proto_goTypes = []interface{}{
	(Type)(0),        // 0: shortdb.table.v1.Type
	(*Row)(nil),      // 1: shortdb.table.v1.Row
	(*Table)(nil),    // 2: shortdb.table.v1.Table
	(*DataBase)(nil), // 3: shortdb.table.v1.DataBase
	nil,              // 4: shortdb.table.v1.Row.ValueEntry
	nil,              // 5: shortdb.table.v1.Table.FieldsEntry
	nil,              // 6: shortdb.table.v1.DataBase.TablesEntry
}
var file_shortdb_table_v1_table_proto_depIdxs = []int32{
	4, // 0: shortdb.table.v1.Row.value:type_name -> shortdb.table.v1.Row.ValueEntry
	5, // 1: shortdb.table.v1.Table.fields:type_name -> shortdb.table.v1.Table.FieldsEntry
	1, // 2: shortdb.table.v1.Table.rows:type_name -> shortdb.table.v1.Row
	6, // 3: shortdb.table.v1.DataBase.tables:type_name -> shortdb.table.v1.DataBase.TablesEntry
	0, // 4: shortdb.table.v1.Table.FieldsEntry.value:type_name -> shortdb.table.v1.Type
	2, // 5: shortdb.table.v1.DataBase.TablesEntry.value:type_name -> shortdb.table.v1.Table
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_shortdb_table_v1_table_proto_init() }
func file_shortdb_table_v1_table_proto_init() {
	if File_shortdb_table_v1_table_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shortdb_table_v1_table_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_shortdb_table_v1_table_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Table); i {
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
		file_shortdb_table_v1_table_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_shortdb_table_v1_table_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shortdb_table_v1_table_proto_goTypes,
		DependencyIndexes: file_shortdb_table_v1_table_proto_depIdxs,
		EnumInfos:         file_shortdb_table_v1_table_proto_enumTypes,
		MessageInfos:      file_shortdb_table_v1_table_proto_msgTypes,
	}.Build()
	File_shortdb_table_v1_table_proto = out.File
	file_shortdb_table_v1_table_proto_rawDesc = nil
	file_shortdb_table_v1_table_proto_goTypes = nil
	file_shortdb_table_v1_table_proto_depIdxs = nil
}

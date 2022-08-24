// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: shortdb/domain/query/v1/query.proto

package v1

import (
	v1 "github.com/batazor/shortlink/pkg/shortdb/domain/field/v1"
	v11 "github.com/batazor/shortlink/pkg/shortdb/domain/index/v1"
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

// Type is the type of SQL query, e.g. SELECT/UPDATE
type Type int32

const (
	// TYPE_UNSPECIFIED is the zero value for a Type
	Type_TYPE_UNSPECIFIED Type = 0
	// TYPE_SELECT represents a SELECT query
	Type_TYPE_SELECT Type = 1
	// TYPE_UPDATE represents an UPDATE query
	Type_TYPE_UPDATE Type = 2
	// TYPE_INSERT represents an INSERT query
	Type_TYPE_INSERT Type = 3
	// TYPE_DELETE represents a DELETE query
	Type_TYPE_DELETE Type = 4
	// TYPE_CREATE_TABLE represents a CREATE TABLE query
	Type_TYPE_CREATE_TABLE Type = 5
	// TYPE_DROP_TABLE represents a DROP TABLE query
	Type_TYPE_DROP_TABLE Type = 6
	// TYPE_CREATE_INDEX represents a CREATE INDEX query
	Type_TYPE_CREATE_INDEX Type = 7
	// TYPE_CREATE_INDEX represents a DELETE INDEX query
	Type_TYPE_DELETE_INDEX Type = 8
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_SELECT",
		2: "TYPE_UPDATE",
		3: "TYPE_INSERT",
		4: "TYPE_DELETE",
		5: "TYPE_CREATE_TABLE",
		6: "TYPE_DROP_TABLE",
		7: "TYPE_CREATE_INDEX",
		8: "TYPE_DELETE_INDEX",
	}
	Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":  0,
		"TYPE_SELECT":       1,
		"TYPE_UPDATE":       2,
		"TYPE_INSERT":       3,
		"TYPE_DELETE":       4,
		"TYPE_CREATE_TABLE": 5,
		"TYPE_DROP_TABLE":   6,
		"TYPE_CREATE_INDEX": 7,
		"TYPE_DELETE_INDEX": 8,
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
	return file_shortdb_domain_query_v1_query_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_shortdb_domain_query_v1_query_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_shortdb_domain_query_v1_query_proto_rawDescGZIP(), []int{0}
}

// Operator is between operands in a condition
type Operator int32

const (
	// OPERATOR_UNSPECIFIED is the zero value for an Operator
	Operator_OPERATOR_UNSPECIFIED Operator = 0
	Operator_OPERATOR_EQ          Operator = 1 // Eq -> "="
	Operator_OPERATOR_NE          Operator = 2 // Ne -> "!="
	Operator_OPERATOR_GT          Operator = 3 // Gt -> ">"
	Operator_OPERATOR_LT          Operator = 4 // Lt -> "<"
	Operator_OPERATOR_GTE         Operator = 5 // Gte -> ">="
	Operator_OPERATOR_LTE         Operator = 6 // Lte -> "<="
)

// Enum value maps for Operator.
var (
	Operator_name = map[int32]string{
		0: "OPERATOR_UNSPECIFIED",
		1: "OPERATOR_EQ",
		2: "OPERATOR_NE",
		3: "OPERATOR_GT",
		4: "OPERATOR_LT",
		5: "OPERATOR_GTE",
		6: "OPERATOR_LTE",
	}
	Operator_value = map[string]int32{
		"OPERATOR_UNSPECIFIED": 0,
		"OPERATOR_EQ":          1,
		"OPERATOR_NE":          2,
		"OPERATOR_GT":          3,
		"OPERATOR_LT":          4,
		"OPERATOR_GTE":         5,
		"OPERATOR_LTE":         6,
	}
)

func (x Operator) Enum() *Operator {
	p := new(Operator)
	*p = x
	return p
}

func (x Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_shortdb_domain_query_v1_query_proto_enumTypes[1].Descriptor()
}

func (Operator) Type() protoreflect.EnumType {
	return &file_shortdb_domain_query_v1_query_proto_enumTypes[1]
}

func (x Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operator.Descriptor instead.
func (Operator) EnumDescriptor() ([]byte, []int) {
	return file_shortdb_domain_query_v1_query_proto_rawDescGZIP(), []int{1}
}

// Condition is a single boolean condition in a WHERE clause
type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LValue is the left hand side operand
	LValue string `protobuf:"bytes,1,opt,name=l_value,json=lValue,proto3" json:"l_value,omitempty"`
	// LValueIsField determines if Operand1 is a literal or a field name
	LValueIsField bool `protobuf:"varint,2,opt,name=l_value_is_field,json=lValueIsField,proto3" json:"l_value_is_field,omitempty"`
	// Operator is e.g. "=", ">"
	Operator Operator `protobuf:"varint,3,opt,name=operator,proto3,enum=shortdb.domain.query.v1.Operator" json:"operator,omitempty"`
	// RValue is the right hand side operand
	RValue string `protobuf:"bytes,4,opt,name=r_value,json=rValue,proto3" json:"r_value,omitempty"`
	// RValueIsField determines if Operand2 is a literal or a field name
	RValueIsField bool `protobuf:"varint,5,opt,name=r_value_is_field,json=rValueIsField,proto3" json:"r_value_is_field,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortdb_domain_query_v1_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_shortdb_domain_query_v1_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_shortdb_domain_query_v1_query_proto_rawDescGZIP(), []int{0}
}

func (x *Condition) GetLValue() string {
	if x != nil {
		return x.LValue
	}
	return ""
}

func (x *Condition) GetLValueIsField() bool {
	if x != nil {
		return x.LValueIsField
	}
	return false
}

func (x *Condition) GetOperator() Operator {
	if x != nil {
		return x.Operator
	}
	return Operator_OPERATOR_UNSPECIFIED
}

func (x *Condition) GetRValue() string {
	if x != nil {
		return x.RValue
	}
	return ""
}

func (x *Condition) GetRValueIsField() bool {
	if x != nil {
		return x.RValueIsField
	}
	return false
}

// Condition is a single boolean condition in a WHERE clause
type JoinCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LHS table name
	LTable string `protobuf:"bytes,1,opt,name=l_table,json=lTable,proto3" json:"l_table,omitempty"`
	// Level operand is the left hand side operand
	LOperand string `protobuf:"bytes,2,opt,name=l_operand,json=lOperand,proto3" json:"l_operand,omitempty"`
	// Operator is e.g. "=", ">"
	Operator Operator `protobuf:"varint,3,opt,name=operator,proto3,enum=shortdb.domain.query.v1.Operator" json:"operator,omitempty"`
	// RHS table name
	RTable string `protobuf:"bytes,4,opt,name=r_table,json=rTable,proto3" json:"r_table,omitempty"`
	// Right operand1 is the right hand side operand
	ROperand string `protobuf:"bytes,5,opt,name=r_operand,json=rOperand,proto3" json:"r_operand,omitempty"`
}

func (x *JoinCondition) Reset() {
	*x = JoinCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortdb_domain_query_v1_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinCondition) ProtoMessage() {}

func (x *JoinCondition) ProtoReflect() protoreflect.Message {
	mi := &file_shortdb_domain_query_v1_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinCondition.ProtoReflect.Descriptor instead.
func (*JoinCondition) Descriptor() ([]byte, []int) {
	return file_shortdb_domain_query_v1_query_proto_rawDescGZIP(), []int{1}
}

func (x *JoinCondition) GetLTable() string {
	if x != nil {
		return x.LTable
	}
	return ""
}

func (x *JoinCondition) GetLOperand() string {
	if x != nil {
		return x.LOperand
	}
	return ""
}

func (x *JoinCondition) GetOperator() Operator {
	if x != nil {
		return x.Operator
	}
	return Operator_OPERATOR_UNSPECIFIED
}

func (x *JoinCondition) GetRTable() string {
	if x != nil {
		return x.RTable
	}
	return ""
}

func (x *JoinCondition) GetROperand() string {
	if x != nil {
		return x.ROperand
	}
	return ""
}

type Join struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       string           `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Table      string           `protobuf:"bytes,2,opt,name=table,proto3" json:"table,omitempty"`
	Conditions []*JoinCondition `protobuf:"bytes,3,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *Join) Reset() {
	*x = Join{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortdb_domain_query_v1_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Join) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Join) ProtoMessage() {}

func (x *Join) ProtoReflect() protoreflect.Message {
	mi := &file_shortdb_domain_query_v1_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Join.ProtoReflect.Descriptor instead.
func (*Join) Descriptor() ([]byte, []int) {
	return file_shortdb_domain_query_v1_query_proto_rawDescGZIP(), []int{2}
}

func (x *Join) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Join) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *Join) GetConditions() []*JoinCondition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

// Query represents a parsed query
type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        Type               `protobuf:"varint,1,opt,name=type,proto3,enum=shortdb.domain.query.v1.Type" json:"type,omitempty"`
	Database    string             `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	TableName   string             `protobuf:"bytes,3,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	TableFields map[string]v1.Type `protobuf:"bytes,4,rep,name=table_fields,json=tableFields,proto3" json:"table_fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=shortdb.domain.field.v1.Type"`
	Indexs      []*v11.Index       `protobuf:"bytes,5,rep,name=indexs,proto3" json:"indexs,omitempty"`
	Conditions  []*Condition       `protobuf:"bytes,6,rep,name=conditions,proto3" json:"conditions,omitempty"`
	Updates     map[string]string  `protobuf:"bytes,7,rep,name=updates,proto3" json:"updates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Inserts     []*Query_Array     `protobuf:"bytes,8,rep,name=inserts,proto3" json:"inserts,omitempty"`
	// Used for SELECT (i.e. SELECTed field names) and INSERT (INSERTEDed field names)
	Fields      []string          `protobuf:"bytes,9,rep,name=fields,proto3" json:"fields,omitempty"`
	Aliases     map[string]string `protobuf:"bytes,10,rep,name=aliases,proto3" json:"aliases,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	OrderFields []string          `protobuf:"bytes,11,rep,name=order_fields,json=orderFields,proto3" json:"order_fields,omitempty"`
	OrderDir    []string          `protobuf:"bytes,12,rep,name=order_dir,json=orderDir,proto3" json:"order_dir,omitempty"`
	Joins       []*Join           `protobuf:"bytes,13,rep,name=joins,proto3" json:"joins,omitempty"`
	MaxRows     int32             `protobuf:"varint,14,opt,name=max_rows,json=maxRows,proto3" json:"max_rows,omitempty"`
	Limit       int32             `protobuf:"varint,15,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortdb_domain_query_v1_query_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_shortdb_domain_query_v1_query_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_shortdb_domain_query_v1_query_proto_rawDescGZIP(), []int{3}
}

func (x *Query) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNSPECIFIED
}

func (x *Query) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *Query) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *Query) GetTableFields() map[string]v1.Type {
	if x != nil {
		return x.TableFields
	}
	return nil
}

func (x *Query) GetIndexs() []*v11.Index {
	if x != nil {
		return x.Indexs
	}
	return nil
}

func (x *Query) GetConditions() []*Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

func (x *Query) GetUpdates() map[string]string {
	if x != nil {
		return x.Updates
	}
	return nil
}

func (x *Query) GetInserts() []*Query_Array {
	if x != nil {
		return x.Inserts
	}
	return nil
}

func (x *Query) GetFields() []string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Query) GetAliases() map[string]string {
	if x != nil {
		return x.Aliases
	}
	return nil
}

func (x *Query) GetOrderFields() []string {
	if x != nil {
		return x.OrderFields
	}
	return nil
}

func (x *Query) GetOrderDir() []string {
	if x != nil {
		return x.OrderDir
	}
	return nil
}

func (x *Query) GetJoins() []*Join {
	if x != nil {
		return x.Joins
	}
	return nil
}

func (x *Query) GetMaxRows() int32 {
	if x != nil {
		return x.MaxRows
	}
	return 0
}

func (x *Query) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type Query_Array struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []string `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Query_Array) Reset() {
	*x = Query_Array{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shortdb_domain_query_v1_query_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query_Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query_Array) ProtoMessage() {}

func (x *Query_Array) ProtoReflect() protoreflect.Message {
	mi := &file_shortdb_domain_query_v1_query_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query_Array.ProtoReflect.Descriptor instead.
func (*Query_Array) Descriptor() ([]byte, []int) {
	return file_shortdb_domain_query_v1_query_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Query_Array) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_shortdb_domain_query_v1_query_proto protoreflect.FileDescriptor

var file_shortdb_domain_query_v1_query_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x23,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2f, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xce, 0x01, 0x0a, 0x09, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x27, 0x0a, 0x10, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x73, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x49, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x3d, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x27, 0x0a, 0x10, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x73, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x72, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x49, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0xba, 0x01, 0x0a, 0x0d, 0x4a, 0x6f,
	0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6c,
	0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e,
	0x64, 0x12, 0x3d, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x5f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x78, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0xc7, 0x07, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x52, 0x0a, 0x0c, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x36, 0x0a, 0x06,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x06, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x73, 0x12, 0x42, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x45, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x3e, 0x0a, 0x07, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x52, 0x07, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x45, 0x0a, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x0b,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x0c,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x69, 0x72, 0x12, 0x33,
	0x0a, 0x05, 0x6a, 0x6f, 0x69, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x05, 0x6a, 0x6f,
	0x69, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x6f, 0x77, 0x73, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x52, 0x6f, 0x77, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x1a, 0x1d, 0x0a, 0x05, 0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x1a, 0x5d, 0x0a, 0x10, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x64, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a,
	0x0a, 0x0c, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0xba, 0x01, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04, 0x12, 0x15, 0x0a,
	0x11, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x52, 0x4f,
	0x50, 0x5f, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x10, 0x07,
	0x12, 0x15, 0x0a, 0x11, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x5f,
	0x49, 0x4e, 0x44, 0x45, 0x58, 0x10, 0x08, 0x2a, 0x8c, 0x01, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f,
	0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x45, 0x51, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x4e, 0x45, 0x10, 0x02,
	0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x47, 0x54, 0x10,
	0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x4c, 0x54,
	0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x47,
	0x54, 0x45, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52,
	0x5f, 0x4c, 0x54, 0x45, 0x10, 0x06, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x61, 0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x64, 0x62, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortdb_domain_query_v1_query_proto_rawDescOnce sync.Once
	file_shortdb_domain_query_v1_query_proto_rawDescData = file_shortdb_domain_query_v1_query_proto_rawDesc
)

func file_shortdb_domain_query_v1_query_proto_rawDescGZIP() []byte {
	file_shortdb_domain_query_v1_query_proto_rawDescOnce.Do(func() {
		file_shortdb_domain_query_v1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortdb_domain_query_v1_query_proto_rawDescData)
	})
	return file_shortdb_domain_query_v1_query_proto_rawDescData
}

var file_shortdb_domain_query_v1_query_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_shortdb_domain_query_v1_query_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_shortdb_domain_query_v1_query_proto_goTypes = []interface{}{
	(Type)(0),             // 0: shortdb.domain.query.v1.Type
	(Operator)(0),         // 1: shortdb.domain.query.v1.Operator
	(*Condition)(nil),     // 2: shortdb.domain.query.v1.Condition
	(*JoinCondition)(nil), // 3: shortdb.domain.query.v1.JoinCondition
	(*Join)(nil),          // 4: shortdb.domain.query.v1.Join
	(*Query)(nil),         // 5: shortdb.domain.query.v1.Query
	(*Query_Array)(nil),   // 6: shortdb.domain.query.v1.Query.Array
	nil,                   // 7: shortdb.domain.query.v1.Query.TableFieldsEntry
	nil,                   // 8: shortdb.domain.query.v1.Query.UpdatesEntry
	nil,                   // 9: shortdb.domain.query.v1.Query.AliasesEntry
	(*v11.Index)(nil),     // 10: shortdb.domain.index.v1.Index
	(v1.Type)(0),          // 11: shortdb.domain.field.v1.Type
}
var file_shortdb_domain_query_v1_query_proto_depIdxs = []int32{
	1,  // 0: shortdb.domain.query.v1.Condition.operator:type_name -> shortdb.domain.query.v1.Operator
	1,  // 1: shortdb.domain.query.v1.JoinCondition.operator:type_name -> shortdb.domain.query.v1.Operator
	3,  // 2: shortdb.domain.query.v1.Join.conditions:type_name -> shortdb.domain.query.v1.JoinCondition
	0,  // 3: shortdb.domain.query.v1.Query.type:type_name -> shortdb.domain.query.v1.Type
	7,  // 4: shortdb.domain.query.v1.Query.table_fields:type_name -> shortdb.domain.query.v1.Query.TableFieldsEntry
	10, // 5: shortdb.domain.query.v1.Query.indexs:type_name -> shortdb.domain.index.v1.Index
	2,  // 6: shortdb.domain.query.v1.Query.conditions:type_name -> shortdb.domain.query.v1.Condition
	8,  // 7: shortdb.domain.query.v1.Query.updates:type_name -> shortdb.domain.query.v1.Query.UpdatesEntry
	6,  // 8: shortdb.domain.query.v1.Query.inserts:type_name -> shortdb.domain.query.v1.Query.Array
	9,  // 9: shortdb.domain.query.v1.Query.aliases:type_name -> shortdb.domain.query.v1.Query.AliasesEntry
	4,  // 10: shortdb.domain.query.v1.Query.joins:type_name -> shortdb.domain.query.v1.Join
	11, // 11: shortdb.domain.query.v1.Query.TableFieldsEntry.value:type_name -> shortdb.domain.field.v1.Type
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_shortdb_domain_query_v1_query_proto_init() }
func file_shortdb_domain_query_v1_query_proto_init() {
	if File_shortdb_domain_query_v1_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shortdb_domain_query_v1_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
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
		file_shortdb_domain_query_v1_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinCondition); i {
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
		file_shortdb_domain_query_v1_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Join); i {
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
		file_shortdb_domain_query_v1_query_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_shortdb_domain_query_v1_query_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query_Array); i {
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
			RawDescriptor: file_shortdb_domain_query_v1_query_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shortdb_domain_query_v1_query_proto_goTypes,
		DependencyIndexes: file_shortdb_domain_query_v1_query_proto_depIdxs,
		EnumInfos:         file_shortdb_domain_query_v1_query_proto_enumTypes,
		MessageInfos:      file_shortdb_domain_query_v1_query_proto_msgTypes,
	}.Build()
	File_shortdb_domain_query_v1_query_proto = out.File
	file_shortdb_domain_query_v1_query_proto_rawDesc = nil
	file_shortdb_domain_query_v1_query_proto_goTypes = nil
	file_shortdb_domain_query_v1_query_proto_depIdxs = nil
}

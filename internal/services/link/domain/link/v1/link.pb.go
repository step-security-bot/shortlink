// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: domain/link/v1/link.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LinkEvent int32

const (
	LinkEvent_LINK_EVENT_UNSPECIFIED LinkEvent = 0
	LinkEvent_LINK_EVENT_ADD         LinkEvent = 1
	LinkEvent_LINK_EVENT_GET         LinkEvent = 2
	LinkEvent_LINK_EVENT_LIST        LinkEvent = 3
	LinkEvent_LINK_EVENT_UPDATE      LinkEvent = 4
	LinkEvent_LINK_EVENT_DELETE      LinkEvent = 5
)

// Enum value maps for LinkEvent.
var (
	LinkEvent_name = map[int32]string{
		0: "LINK_EVENT_UNSPECIFIED",
		1: "LINK_EVENT_ADD",
		2: "LINK_EVENT_GET",
		3: "LINK_EVENT_LIST",
		4: "LINK_EVENT_UPDATE",
		5: "LINK_EVENT_DELETE",
	}
	LinkEvent_value = map[string]int32{
		"LINK_EVENT_UNSPECIFIED": 0,
		"LINK_EVENT_ADD":         1,
		"LINK_EVENT_GET":         2,
		"LINK_EVENT_LIST":        3,
		"LINK_EVENT_UPDATE":      4,
		"LINK_EVENT_DELETE":      5,
	}
)

func (x LinkEvent) Enum() *LinkEvent {
	p := new(LinkEvent)
	*p = x
	return p
}

func (x LinkEvent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LinkEvent) Descriptor() protoreflect.EnumDescriptor {
	return file_domain_link_v1_link_proto_enumTypes[0].Descriptor()
}

func (LinkEvent) Type() protoreflect.EnumType {
	return &file_domain_link_v1_link_proto_enumTypes[0]
}

func (x LinkEvent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LinkEvent.Descriptor instead.
func (LinkEvent) EnumDescriptor() ([]byte, []int) {
	return file_domain_link_v1_link_proto_rawDescGZIP(), []int{0}
}

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldMask *fieldmaskpb.FieldMask `protobuf:"bytes,6,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// URL
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Hash by URL + salt
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// Describe of link
	Describe string `protobuf:"bytes,3,opt,name=describe,proto3" json:"describe,omitempty"`
	// Create at
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Update at
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_link_v1_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_domain_link_v1_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_domain_link_v1_link_proto_rawDescGZIP(), []int{0}
}

func (x *Link) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

func (x *Link) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Link) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *Link) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *Link) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Link) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Links struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link []*Link `protobuf:"bytes,1,rep,name=link,proto3" json:"link,omitempty"`
}

func (x *Links) Reset() {
	*x = Links{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_link_v1_link_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Links) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Links) ProtoMessage() {}

func (x *Links) ProtoReflect() protoreflect.Message {
	mi := &file_domain_link_v1_link_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Links.ProtoReflect.Descriptor instead.
func (*Links) Descriptor() ([]byte, []int) {
	return file_domain_link_v1_link_proto_rawDescGZIP(), []int{1}
}

func (x *Links) GetLink() []*Link {
	if x != nil {
		return x.Link
	}
	return nil
}

var File_domain_link_v1_link_proto protoreflect.FileDescriptor

var file_domain_link_v1_link_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9,
	0x01, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61,
	0x73, 0x6b, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x31, 0x0a, 0x05, 0x4c, 0x69,
	0x6e, 0x6b, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x2a, 0x92, 0x01,
	0x0a, 0x09, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x16, 0x4c,
	0x49, 0x4e, 0x4b, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4c, 0x49, 0x4e, 0x4b, 0x5f,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x44, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4c,
	0x49, 0x4e, 0x4b, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x47, 0x45, 0x54, 0x10, 0x02, 0x12,
	0x13, 0x0a, 0x0f, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x4c, 0x49,
	0x53, 0x54, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x49, 0x4e, 0x4b, 0x5f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x4c,
	0x49, 0x4e, 0x4b, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x10, 0x05, 0x42, 0xc3, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x4c, 0x69, 0x6e, 0x6b, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x6f, 0x72, 0x67,
	0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x6e,
	0x6b, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x44, 0x4c, 0x58, 0xaa, 0x02, 0x0e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x5c, 0x4c, 0x69, 0x6e, 0x6b, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x5c, 0x4c, 0x69, 0x6e, 0x6b, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x3a, 0x3a,
	0x4c, 0x69, 0x6e, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_link_v1_link_proto_rawDescOnce sync.Once
	file_domain_link_v1_link_proto_rawDescData = file_domain_link_v1_link_proto_rawDesc
)

func file_domain_link_v1_link_proto_rawDescGZIP() []byte {
	file_domain_link_v1_link_proto_rawDescOnce.Do(func() {
		file_domain_link_v1_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_link_v1_link_proto_rawDescData)
	})
	return file_domain_link_v1_link_proto_rawDescData
}

var file_domain_link_v1_link_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_domain_link_v1_link_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_domain_link_v1_link_proto_goTypes = []interface{}{
	(LinkEvent)(0),                // 0: domain.link.v1.LinkEvent
	(*Link)(nil),                  // 1: domain.link.v1.Link
	(*Links)(nil),                 // 2: domain.link.v1.Links
	(*fieldmaskpb.FieldMask)(nil), // 3: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_domain_link_v1_link_proto_depIdxs = []int32{
	3, // 0: domain.link.v1.Link.field_mask:type_name -> google.protobuf.FieldMask
	4, // 1: domain.link.v1.Link.created_at:type_name -> google.protobuf.Timestamp
	4, // 2: domain.link.v1.Link.updated_at:type_name -> google.protobuf.Timestamp
	1, // 3: domain.link.v1.Links.link:type_name -> domain.link.v1.Link
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_domain_link_v1_link_proto_init() }
func file_domain_link_v1_link_proto_init() {
	if File_domain_link_v1_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_link_v1_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
		file_domain_link_v1_link_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Links); i {
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
			RawDescriptor: file_domain_link_v1_link_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_link_v1_link_proto_goTypes,
		DependencyIndexes: file_domain_link_v1_link_proto_depIdxs,
		EnumInfos:         file_domain_link_v1_link_proto_enumTypes,
		MessageInfos:      file_domain_link_v1_link_proto_msgTypes,
	}.Build()
	File_domain_link_v1_link_proto = out.File
	file_domain_link_v1_link_proto_rawDesc = nil
	file_domain_link_v1_link_proto_goTypes = nil
	file_domain_link_v1_link_proto_depIdxs = nil
}

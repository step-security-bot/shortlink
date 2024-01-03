// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: domain/link_cqrs/v1/link.proto

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

// Link
type LinkView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// FieldMask
	FieldMask *fieldmaskpb.FieldMask `protobuf:"bytes,9,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// URL
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Hash by URL + salt
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// Describe of link
	Describe string `protobuf:"bytes,3,opt,name=describe,proto3" json:"describe,omitempty"`
	// Metadata
	ImageUrl string `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	// Meta description
	MetaDescription string `protobuf:"bytes,5,opt,name=meta_description,json=metaDescription,proto3" json:"meta_description,omitempty"`
	// Meta keywords
	MetaKeywords string `protobuf:"bytes,6,opt,name=meta_keywords,json=metaKeywords,proto3" json:"meta_keywords,omitempty"`
	// Create at
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Update at
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *LinkView) Reset() {
	*x = LinkView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_link_cqrs_v1_link_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkView) ProtoMessage() {}

func (x *LinkView) ProtoReflect() protoreflect.Message {
	mi := &file_domain_link_cqrs_v1_link_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkView.ProtoReflect.Descriptor instead.
func (*LinkView) Descriptor() ([]byte, []int) {
	return file_domain_link_cqrs_v1_link_proto_rawDescGZIP(), []int{0}
}

func (x *LinkView) GetFieldMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.FieldMask
	}
	return nil
}

func (x *LinkView) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *LinkView) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *LinkView) GetDescribe() string {
	if x != nil {
		return x.Describe
	}
	return ""
}

func (x *LinkView) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *LinkView) GetMetaDescription() string {
	if x != nil {
		return x.MetaDescription
	}
	return ""
}

func (x *LinkView) GetMetaKeywords() string {
	if x != nil {
		return x.MetaKeywords
	}
	return ""
}

func (x *LinkView) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *LinkView) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Links
type LinksView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Links
	Links []*LinkView `protobuf:"bytes,1,rep,name=links,proto3" json:"links,omitempty"`
}

func (x *LinksView) Reset() {
	*x = LinksView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_link_cqrs_v1_link_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinksView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinksView) ProtoMessage() {}

func (x *LinksView) ProtoReflect() protoreflect.Message {
	mi := &file_domain_link_cqrs_v1_link_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinksView.ProtoReflect.Descriptor instead.
func (*LinksView) Descriptor() ([]byte, []int) {
	return file_domain_link_cqrs_v1_link_proto_rawDescGZIP(), []int{1}
}

func (x *LinksView) GetLinks() []*LinkView {
	if x != nil {
		return x.Links
	}
	return nil
}

var File_domain_link_cqrs_v1_link_proto protoreflect.FileDescriptor

var file_domain_link_cqrs_v1_link_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x63, 0x71,
	0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x63, 0x71,
	0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x02, 0x0a, 0x08, 0x4c, 0x69, 0x6e,
	0x6b, 0x56, 0x69, 0x65, 0x77, 0x12, 0x39, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d,
	0x61, 0x73, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12,
	0x29, 0x0a, 0x10, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x65, 0x74, 0x61, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65,
	0x74, 0x61, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x40, 0x0a, 0x09, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x56, 0x69,
	0x65, 0x77, 0x12, 0x33, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f,
	0x63, 0x71, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x56, 0x69, 0x65, 0x77,
	0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x42, 0xe4, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x63, 0x71, 0x72, 0x73,
	0x2e, 0x76, 0x31, 0x42, 0x09, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x6f, 0x72, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x6c, 0x69,
	0x6e, 0x6b, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x63,
	0x71, 0x72, 0x73, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x44, 0x4c, 0x58, 0xaa, 0x02, 0x12, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x71, 0x72, 0x73, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x12, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c, 0x4c, 0x69, 0x6e, 0x6b, 0x43,
	0x71, 0x72, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5c,
	0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x71, 0x72, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x3a, 0x3a, 0x4c, 0x69, 0x6e, 0x6b, 0x43, 0x71, 0x72, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_link_cqrs_v1_link_proto_rawDescOnce sync.Once
	file_domain_link_cqrs_v1_link_proto_rawDescData = file_domain_link_cqrs_v1_link_proto_rawDesc
)

func file_domain_link_cqrs_v1_link_proto_rawDescGZIP() []byte {
	file_domain_link_cqrs_v1_link_proto_rawDescOnce.Do(func() {
		file_domain_link_cqrs_v1_link_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_link_cqrs_v1_link_proto_rawDescData)
	})
	return file_domain_link_cqrs_v1_link_proto_rawDescData
}

var file_domain_link_cqrs_v1_link_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_domain_link_cqrs_v1_link_proto_goTypes = []interface{}{
	(*LinkView)(nil),              // 0: domain.link_cqrs.v1.LinkView
	(*LinksView)(nil),             // 1: domain.link_cqrs.v1.LinksView
	(*fieldmaskpb.FieldMask)(nil), // 2: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_domain_link_cqrs_v1_link_proto_depIdxs = []int32{
	2, // 0: domain.link_cqrs.v1.LinkView.field_mask:type_name -> google.protobuf.FieldMask
	3, // 1: domain.link_cqrs.v1.LinkView.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: domain.link_cqrs.v1.LinkView.updated_at:type_name -> google.protobuf.Timestamp
	0, // 3: domain.link_cqrs.v1.LinksView.links:type_name -> domain.link_cqrs.v1.LinkView
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_domain_link_cqrs_v1_link_proto_init() }
func file_domain_link_cqrs_v1_link_proto_init() {
	if File_domain_link_cqrs_v1_link_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_link_cqrs_v1_link_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkView); i {
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
		file_domain_link_cqrs_v1_link_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinksView); i {
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
			RawDescriptor: file_domain_link_cqrs_v1_link_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_link_cqrs_v1_link_proto_goTypes,
		DependencyIndexes: file_domain_link_cqrs_v1_link_proto_depIdxs,
		MessageInfos:      file_domain_link_cqrs_v1_link_proto_msgTypes,
	}.Build()
	File_domain_link_cqrs_v1_link_proto = out.File
	file_domain_link_cqrs_v1_link_proto_rawDesc = nil
	file_domain_link_cqrs_v1_link_proto_goTypes = nil
	file_domain_link_cqrs_v1_link_proto_depIdxs = nil
}

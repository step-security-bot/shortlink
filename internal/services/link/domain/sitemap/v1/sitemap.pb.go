// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: domain/sitemap/v1/sitemap.proto

package v1

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// URL is a structure of <url> in <sitemap>
type Url struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Loc        string  `protobuf:"bytes,1,opt,name=loc,proto3" json:"loc,omitempty" xml:"loc"`
	LastMod    string  `protobuf:"bytes,2,opt,name=last_mod,json=lastMod,proto3" json:"last_mod,omitempty" xml:"lastmod"`
	ChangeFreq string  `protobuf:"bytes,3,opt,name=change_freq,json=changeFreq,proto3" json:"change_freq,omitempty" xml:"changefreq"`
	Priority   float32 `protobuf:"fixed32,4,opt,name=priority,proto3" json:"priority,omitempty" xml:"priority"`
}

func (x *Url) Reset() {
	*x = Url{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_sitemap_v1_sitemap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Url) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Url) ProtoMessage() {}

func (x *Url) ProtoReflect() protoreflect.Message {
	mi := &file_domain_sitemap_v1_sitemap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Url.ProtoReflect.Descriptor instead.
func (*Url) Descriptor() ([]byte, []int) {
	return file_domain_sitemap_v1_sitemap_proto_rawDescGZIP(), []int{0}
}

func (x *Url) GetLoc() string {
	if x != nil {
		return x.Loc
	}
	return ""
}

func (x *Url) GetLastMod() string {
	if x != nil {
		return x.LastMod
	}
	return ""
}

func (x *Url) GetChangeFreq() string {
	if x != nil {
		return x.ChangeFreq
	}
	return ""
}

func (x *Url) GetPriority() float32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

// Sitemap is a structure of <sitemap>
type Sitemap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url []*Url `protobuf:"bytes,1,rep,name=url,proto3" json:"url,omitempty" xml:"url"`
}

func (x *Sitemap) Reset() {
	*x = Sitemap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_sitemap_v1_sitemap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sitemap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sitemap) ProtoMessage() {}

func (x *Sitemap) ProtoReflect() protoreflect.Message {
	mi := &file_domain_sitemap_v1_sitemap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sitemap.ProtoReflect.Descriptor instead.
func (*Sitemap) Descriptor() ([]byte, []int) {
	return file_domain_sitemap_v1_sitemap_proto_rawDescGZIP(), []int{1}
}

func (x *Sitemap) GetUrl() []*Url {
	if x != nil {
		return x.Url
	}
	return nil
}

var File_domain_sitemap_v1_sitemap_proto protoreflect.FileDescriptor

var file_domain_sitemap_v1_sitemap_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x70,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x73, 0x69, 0x74, 0x65, 0x6d, 0x61,
	0x70, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x03, 0x55,
	0x72, 0x6c, 0x12, 0x20, 0x0a, 0x03, 0x6c, 0x6f, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0e, 0x9a, 0x84, 0x9e, 0x03, 0x09, 0x78, 0x6d, 0x6c, 0x3a, 0x22, 0x6c, 0x6f, 0x63, 0x22, 0x52,
	0x03, 0x6c, 0x6f, 0x63, 0x12, 0x2d, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0x9a, 0x84, 0x9e, 0x03, 0x0d, 0x78, 0x6d, 0x6c,
	0x3a, 0x22, 0x6c, 0x61, 0x73, 0x74, 0x6d, 0x6f, 0x64, 0x22, 0x52, 0x07, 0x6c, 0x61, 0x73, 0x74,
	0x4d, 0x6f, 0x64, 0x12, 0x36, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x66, 0x72,
	0x65, 0x71, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0x9a, 0x84, 0x9e, 0x03, 0x10, 0x78,
	0x6d, 0x6c, 0x3a, 0x22, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x66, 0x72, 0x65, 0x71, 0x22, 0x52,
	0x0a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x46, 0x72, 0x65, 0x71, 0x12, 0x2f, 0x0a, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x42, 0x13, 0x9a,
	0x84, 0x9e, 0x03, 0x0e, 0x78, 0x6d, 0x6c, 0x3a, 0x22, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x22, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x43, 0x0a, 0x07,
	0x53, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x70, 0x12, 0x38, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x73, 0x69,
	0x74, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x72, 0x6c, 0x42, 0x0e, 0x9a, 0x84,
	0x9e, 0x03, 0x09, 0x78, 0x6d, 0x6c, 0x3a, 0x22, 0x75, 0x72, 0x6c, 0x22, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x61, 0x74, 0x61, 0x7a, 0x6f, 0x72, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x6c, 0x69, 0x6e,
	0x6b, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x73, 0x69, 0x74, 0x65, 0x6d, 0x61, 0x70, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_domain_sitemap_v1_sitemap_proto_rawDescOnce sync.Once
	file_domain_sitemap_v1_sitemap_proto_rawDescData = file_domain_sitemap_v1_sitemap_proto_rawDesc
)

func file_domain_sitemap_v1_sitemap_proto_rawDescGZIP() []byte {
	file_domain_sitemap_v1_sitemap_proto_rawDescOnce.Do(func() {
		file_domain_sitemap_v1_sitemap_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_sitemap_v1_sitemap_proto_rawDescData)
	})
	return file_domain_sitemap_v1_sitemap_proto_rawDescData
}

var file_domain_sitemap_v1_sitemap_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_domain_sitemap_v1_sitemap_proto_goTypes = []interface{}{
	(*Url)(nil),     // 0: domain.sitemap.v1.Url
	(*Sitemap)(nil), // 1: domain.sitemap.v1.Sitemap
}
var file_domain_sitemap_v1_sitemap_proto_depIdxs = []int32{
	0, // 0: domain.sitemap.v1.Sitemap.url:type_name -> domain.sitemap.v1.Url
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_domain_sitemap_v1_sitemap_proto_init() }
func file_domain_sitemap_v1_sitemap_proto_init() {
	if File_domain_sitemap_v1_sitemap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_sitemap_v1_sitemap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Url); i {
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
		file_domain_sitemap_v1_sitemap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sitemap); i {
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
			RawDescriptor: file_domain_sitemap_v1_sitemap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_sitemap_v1_sitemap_proto_goTypes,
		DependencyIndexes: file_domain_sitemap_v1_sitemap_proto_depIdxs,
		MessageInfos:      file_domain_sitemap_v1_sitemap_proto_msgTypes,
	}.Build()
	File_domain_sitemap_v1_sitemap_proto = out.File
	file_domain_sitemap_v1_sitemap_proto_rawDesc = nil
	file_domain_sitemap_v1_sitemap_proto_goTypes = nil
	file_domain_sitemap_v1_sitemap_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: analyzer.proto

package proto

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

type AzRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Body     string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Type     string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Encoding string `protobuf:"bytes,4,opt,name=encoding,proto3" json:"encoding,omitempty"`
}

func (x *AzRequest) Reset() {
	*x = AzRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyzer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AzRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzRequest) ProtoMessage() {}

func (x *AzRequest) ProtoReflect() protoreflect.Message {
	mi := &file_analyzer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzRequest.ProtoReflect.Descriptor instead.
func (*AzRequest) Descriptor() ([]byte, []int) {
	return file_analyzer_proto_rawDescGZIP(), []int{0}
}

func (x *AzRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AzRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *AzRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AzRequest) GetEncoding() string {
	if x != nil {
		return x.Encoding
	}
	return ""
}

type AzResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *AzResponse) Reset() {
	*x = AzResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyzer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AzResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzResponse) ProtoMessage() {}

func (x *AzResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analyzer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzResponse.ProtoReflect.Descriptor instead.
func (*AzResponse) Descriptor() ([]byte, []int) {
	return file_analyzer_proto_rawDescGZIP(), []int{1}
}

func (x *AzResponse) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type Lineaccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Root  string `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	Key   string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Lineaccess) Reset() {
	*x = Lineaccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyzer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lineaccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lineaccess) ProtoMessage() {}

func (x *Lineaccess) ProtoReflect() protoreflect.Message {
	mi := &file_analyzer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lineaccess.ProtoReflect.Descriptor instead.
func (*Lineaccess) Descriptor() ([]byte, []int) {
	return file_analyzer_proto_rawDescGZIP(), []int{2}
}

func (x *Lineaccess) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Lineaccess) GetRoot() string {
	if x != nil {
		return x.Root
	}
	return ""
}

func (x *Lineaccess) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Lineaccess) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type Fileaccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Lines []*Lineaccess `protobuf:"bytes,2,rep,name=lines,proto3" json:"lines,omitempty"`
}

func (x *Fileaccess) Reset() {
	*x = Fileaccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyzer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Fileaccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fileaccess) ProtoMessage() {}

func (x *Fileaccess) ProtoReflect() protoreflect.Message {
	mi := &file_analyzer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Fileaccess.ProtoReflect.Descriptor instead.
func (*Fileaccess) Descriptor() ([]byte, []int) {
	return file_analyzer_proto_rawDescGZIP(), []int{3}
}

func (x *Fileaccess) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Fileaccess) GetLines() []*Lineaccess {
	if x != nil {
		return x.Lines
	}
	return nil
}

type Pageaccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Files []*Fileaccess `protobuf:"bytes,2,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *Pageaccess) Reset() {
	*x = Pageaccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyzer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pageaccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pageaccess) ProtoMessage() {}

func (x *Pageaccess) ProtoReflect() protoreflect.Message {
	mi := &file_analyzer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pageaccess.ProtoReflect.Descriptor instead.
func (*Pageaccess) Descriptor() ([]byte, []int) {
	return file_analyzer_proto_rawDescGZIP(), []int{4}
}

func (x *Pageaccess) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pageaccess) GetFiles() []*Fileaccess {
	if x != nil {
		return x.Files
	}
	return nil
}

type StoresigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *StoresigResponse) Reset() {
	*x = StoresigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_analyzer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoresigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoresigResponse) ProtoMessage() {}

func (x *StoresigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analyzer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoresigResponse.ProtoReflect.Descriptor instead.
func (*StoresigResponse) Descriptor() ([]byte, []int) {
	return file_analyzer_proto_rawDescGZIP(), []int{5}
}

func (x *StoresigResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_analyzer_proto protoreflect.FileDescriptor

var file_analyzer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x09, 0x41, 0x7a, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x22, 0x20, 0x0a, 0x0a,
	0x41, 0x7a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x5c,
	0x0a, 0x0a, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x72, 0x6f, 0x6f, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x49, 0x0a, 0x0a,
	0x46, 0x69, 0x6c, 0x65, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27,
	0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x22, 0x49, 0x0a, 0x0a, 0x50, 0x61, 0x67, 0x65, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x05, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0x7c, 0x0a, 0x08, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a,
	0x65, 0x72, 0x12, 0x30, 0x0a, 0x07, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x12, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x7a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x7a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x77, 0x70, 0x72, 0x2f, 0x73, 0x72, 0x63, 0x2f,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x7a, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_analyzer_proto_rawDescOnce sync.Once
	file_analyzer_proto_rawDescData = file_analyzer_proto_rawDesc
)

func file_analyzer_proto_rawDescGZIP() []byte {
	file_analyzer_proto_rawDescOnce.Do(func() {
		file_analyzer_proto_rawDescData = protoimpl.X.CompressGZIP(file_analyzer_proto_rawDescData)
	})
	return file_analyzer_proto_rawDescData
}

var file_analyzer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_analyzer_proto_goTypes = []interface{}{
	(*AzRequest)(nil),        // 0: proto.AzRequest
	(*AzResponse)(nil),       // 1: proto.AzResponse
	(*Lineaccess)(nil),       // 2: proto.Lineaccess
	(*Fileaccess)(nil),       // 3: proto.Fileaccess
	(*Pageaccess)(nil),       // 4: proto.Pageaccess
	(*StoresigResponse)(nil), // 5: proto.StoresigResponse
}
var file_analyzer_proto_depIdxs = []int32{
	2, // 0: proto.Fileaccess.lines:type_name -> proto.Lineaccess
	3, // 1: proto.Pageaccess.files:type_name -> proto.Fileaccess
	0, // 2: proto.Analyzer.Analyze:input_type -> proto.AzRequest
	4, // 3: proto.Analyzer.Storesignature:input_type -> proto.Pageaccess
	1, // 4: proto.Analyzer.Analyze:output_type -> proto.AzResponse
	5, // 5: proto.Analyzer.Storesignature:output_type -> proto.StoresigResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_analyzer_proto_init() }
func file_analyzer_proto_init() {
	if File_analyzer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_analyzer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AzRequest); i {
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
		file_analyzer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AzResponse); i {
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
		file_analyzer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lineaccess); i {
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
		file_analyzer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Fileaccess); i {
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
		file_analyzer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pageaccess); i {
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
		file_analyzer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoresigResponse); i {
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
			RawDescriptor: file_analyzer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analyzer_proto_goTypes,
		DependencyIndexes: file_analyzer_proto_depIdxs,
		MessageInfos:      file_analyzer_proto_msgTypes,
	}.Build()
	File_analyzer_proto = out.File
	file_analyzer_proto_rawDesc = nil
	file_analyzer_proto_goTypes = nil
	file_analyzer_proto_depIdxs = nil
}
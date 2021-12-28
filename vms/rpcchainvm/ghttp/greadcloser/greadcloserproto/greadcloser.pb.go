// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: greadcloser.proto

package greadcloserproto

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

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length int32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greadcloser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greadcloser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_greadcloser_proto_rawDescGZIP(), []int{0}
}

func (x *ReadRequest) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Read    []byte `protobuf:"bytes,1,opt,name=read,proto3" json:"read,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Errored bool   `protobuf:"varint,3,opt,name=errored,proto3" json:"errored,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greadcloser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greadcloser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_greadcloser_proto_rawDescGZIP(), []int{1}
}

func (x *ReadResponse) GetRead() []byte {
	if x != nil {
		return x.Read
	}
	return nil
}

func (x *ReadResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ReadResponse) GetErrored() bool {
	if x != nil {
		return x.Errored
	}
	return false
}

type CloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CloseRequest) Reset() {
	*x = CloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greadcloser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseRequest) ProtoMessage() {}

func (x *CloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greadcloser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseRequest.ProtoReflect.Descriptor instead.
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return file_greadcloser_proto_rawDescGZIP(), []int{2}
}

type CloseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CloseResponse) Reset() {
	*x = CloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greadcloser_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseResponse) ProtoMessage() {}

func (x *CloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greadcloser_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseResponse.ProtoReflect.Descriptor instead.
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return file_greadcloser_proto_rawDescGZIP(), []int{3}
}

var File_greadcloser_proto protoreflect.FileDescriptor

var file_greadcloser_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x65, 0x61, 0x64, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67, 0x72, 0x65, 0x61, 0x64, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x72,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x52, 0x0a, 0x0c,
	0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x64,
	0x22, 0x0e, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x0f, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x99, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x04,
	0x52, 0x65, 0x61, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x65, 0x61, 0x64, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x72, 0x65, 0x61, 0x64, 0x63, 0x6c, 0x6f, 0x73, 0x65,
	0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x1e, 0x2e, 0x67,
	0x72, 0x65, 0x61, 0x64, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67,
	0x72, 0x65, 0x61, 0x64, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x51, 0x5a,
	0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x61, 0x72,
	0x65, 0x2d, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x6c, 0x61,
	0x72, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x76, 0x6d, 0x2f, 0x67, 0x68,
	0x74, 0x74, 0x70, 0x2f, 0x67, 0x72, 0x65, 0x61, 0x64, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x2f,
	0x67, 0x72, 0x65, 0x61, 0x64, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greadcloser_proto_rawDescOnce sync.Once
	file_greadcloser_proto_rawDescData = file_greadcloser_proto_rawDesc
)

func file_greadcloser_proto_rawDescGZIP() []byte {
	file_greadcloser_proto_rawDescOnce.Do(func() {
		file_greadcloser_proto_rawDescData = protoimpl.X.CompressGZIP(file_greadcloser_proto_rawDescData)
	})
	return file_greadcloser_proto_rawDescData
}

var file_greadcloser_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_greadcloser_proto_goTypes = []interface{}{
	(*ReadRequest)(nil),   // 0: greadcloserproto.ReadRequest
	(*ReadResponse)(nil),  // 1: greadcloserproto.ReadResponse
	(*CloseRequest)(nil),  // 2: greadcloserproto.CloseRequest
	(*CloseResponse)(nil), // 3: greadcloserproto.CloseResponse
}
var file_greadcloser_proto_depIdxs = []int32{
	0, // 0: greadcloserproto.Reader.Read:input_type -> greadcloserproto.ReadRequest
	2, // 1: greadcloserproto.Reader.Close:input_type -> greadcloserproto.CloseRequest
	1, // 2: greadcloserproto.Reader.Read:output_type -> greadcloserproto.ReadResponse
	3, // 3: greadcloserproto.Reader.Close:output_type -> greadcloserproto.CloseResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_greadcloser_proto_init() }
func file_greadcloser_proto_init() {
	if File_greadcloser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greadcloser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_greadcloser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
		file_greadcloser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseRequest); i {
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
		file_greadcloser_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseResponse); i {
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
			RawDescriptor: file_greadcloser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greadcloser_proto_goTypes,
		DependencyIndexes: file_greadcloser_proto_depIdxs,
		MessageInfos:      file_greadcloser_proto_msgTypes,
	}.Build()
	File_greadcloser_proto = out.File
	file_greadcloser_proto_rawDesc = nil
	file_greadcloser_proto_goTypes = nil
	file_greadcloser_proto_depIdxs = nil
}

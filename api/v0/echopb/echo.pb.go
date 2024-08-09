// kvdbserver gRPC API
// API version: 0.2.0
//
// This package contains Protobuf definitions related to echo RPCs.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: api/v0/echo/echo.proto

package echopb

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

type UnaryEchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message to be sent.
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UnaryEchoRequest) Reset() {
	*x = UnaryEchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v0_echo_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryEchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryEchoRequest) ProtoMessage() {}

func (x *UnaryEchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0_echo_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryEchoRequest.ProtoReflect.Descriptor instead.
func (*UnaryEchoRequest) Descriptor() ([]byte, []int) {
	return file_api_v0_echo_echo_proto_rawDescGZIP(), []int{0}
}

func (x *UnaryEchoRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UnaryEchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message to be received.
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UnaryEchoResponse) Reset() {
	*x = UnaryEchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v0_echo_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryEchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryEchoResponse) ProtoMessage() {}

func (x *UnaryEchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v0_echo_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryEchoResponse.ProtoReflect.Descriptor instead.
func (*UnaryEchoResponse) Descriptor() ([]byte, []int) {
	return file_api_v0_echo_echo_proto_rawDescGZIP(), []int{1}
}

func (x *UnaryEchoResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_api_v0_echo_echo_proto protoreflect.FileDescriptor

var file_api_v0_echo_echo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x65, 0x63,
	0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30,
	0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x22, 0x24, 0x0a, 0x10, 0x55, 0x6e, 0x61, 0x72, 0x79,
	0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x25, 0x0a,
	0x11, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x32, 0x5f, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x63, 0x68, 0x6f,
	0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70, 0x62,
	0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x30, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x70,
	0x62, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x30, 0x2f,
	0x65, 0x63, 0x68, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v0_echo_echo_proto_rawDescOnce sync.Once
	file_api_v0_echo_echo_proto_rawDescData = file_api_v0_echo_echo_proto_rawDesc
)

func file_api_v0_echo_echo_proto_rawDescGZIP() []byte {
	file_api_v0_echo_echo_proto_rawDescOnce.Do(func() {
		file_api_v0_echo_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v0_echo_echo_proto_rawDescData)
	})
	return file_api_v0_echo_echo_proto_rawDescData
}

var file_api_v0_echo_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_v0_echo_echo_proto_goTypes = []interface{}{
	(*UnaryEchoRequest)(nil),  // 0: api.v0.echopb.UnaryEchoRequest
	(*UnaryEchoResponse)(nil), // 1: api.v0.echopb.UnaryEchoResponse
}
var file_api_v0_echo_echo_proto_depIdxs = []int32{
	0, // 0: api.v0.echopb.EchoService.UnaryEcho:input_type -> api.v0.echopb.UnaryEchoRequest
	1, // 1: api.v0.echopb.EchoService.UnaryEcho:output_type -> api.v0.echopb.UnaryEchoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_v0_echo_echo_proto_init() }
func file_api_v0_echo_echo_proto_init() {
	if File_api_v0_echo_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v0_echo_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryEchoRequest); i {
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
		file_api_v0_echo_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryEchoResponse); i {
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
			RawDescriptor: file_api_v0_echo_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v0_echo_echo_proto_goTypes,
		DependencyIndexes: file_api_v0_echo_echo_proto_depIdxs,
		MessageInfos:      file_api_v0_echo_echo_proto_msgTypes,
	}.Build()
	File_api_v0_echo_echo_proto = out.File
	file_api_v0_echo_echo_proto_rawDesc = nil
	file_api_v0_echo_echo_proto_goTypes = nil
	file_api_v0_echo_echo_proto_depIdxs = nil
}

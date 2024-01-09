// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: proto/kvdbserver/storage.proto

package kvdbserver

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

type SetStringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key used to store the value.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Value to store.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetStringRequest) Reset() {
	*x = SetStringRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserver_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStringRequest) ProtoMessage() {}

func (x *SetStringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserver_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStringRequest.ProtoReflect.Descriptor instead.
func (*SetStringRequest) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserver_storage_proto_rawDescGZIP(), []int{0}
}

func (x *SetStringRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetStringRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SetStringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetStringResponse) Reset() {
	*x = SetStringResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserver_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetStringResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetStringResponse) ProtoMessage() {}

func (x *SetStringResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserver_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetStringResponse.ProtoReflect.Descriptor instead.
func (*SetStringResponse) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserver_storage_proto_rawDescGZIP(), []int{1}
}

type GetStringRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key used to get the value.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetStringRequest) Reset() {
	*x = GetStringRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserver_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStringRequest) ProtoMessage() {}

func (x *GetStringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserver_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStringRequest.ProtoReflect.Descriptor instead.
func (*GetStringRequest) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserver_storage_proto_rawDescGZIP(), []int{2}
}

func (x *GetStringRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetStringResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returned value.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// True if key exists. False otherwise.
	Found bool `protobuf:"varint,2,opt,name=found,proto3" json:"found,omitempty"`
}

func (x *GetStringResponse) Reset() {
	*x = GetStringResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserver_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStringResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStringResponse) ProtoMessage() {}

func (x *GetStringResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserver_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStringResponse.ProtoReflect.Descriptor instead.
func (*GetStringResponse) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserver_storage_proto_rawDescGZIP(), []int{3}
}

func (x *GetStringResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetStringResponse) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

type DeleteKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key to delete.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteKeyRequest) Reset() {
	*x = DeleteKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserver_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKeyRequest) ProtoMessage() {}

func (x *DeleteKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserver_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKeyRequest.ProtoReflect.Descriptor instead.
func (*DeleteKeyRequest) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserver_storage_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteKeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// True if key was deleted. False if it doesn't exist.
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteKeyResponse) Reset() {
	*x = DeleteKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserver_storage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteKeyResponse) ProtoMessage() {}

func (x *DeleteKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserver_storage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteKeyResponse.ProtoReflect.Descriptor instead.
func (*DeleteKeyResponse) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserver_storage_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteKeyResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_kvdbserver_storage_proto protoreflect.FileDescriptor

var file_proto_kvdbserver_storage_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x22,
	0x3a, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x53,
	0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x24, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x24, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x2d, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x86, 0x02, 0x0a,
	0x0e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x50, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x2e, 0x6b,
	0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x74,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65,
	0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x50, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1f,
	0x2e, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x12, 0x1f, 0x2e, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70,
	0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b,
	0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_kvdbserver_storage_proto_rawDescOnce sync.Once
	file_proto_kvdbserver_storage_proto_rawDescData = file_proto_kvdbserver_storage_proto_rawDesc
)

func file_proto_kvdbserver_storage_proto_rawDescGZIP() []byte {
	file_proto_kvdbserver_storage_proto_rawDescOnce.Do(func() {
		file_proto_kvdbserver_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_kvdbserver_storage_proto_rawDescData)
	})
	return file_proto_kvdbserver_storage_proto_rawDescData
}

var file_proto_kvdbserver_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_kvdbserver_storage_proto_goTypes = []interface{}{
	(*SetStringRequest)(nil),  // 0: kvdbserverapi.SetStringRequest
	(*SetStringResponse)(nil), // 1: kvdbserverapi.SetStringResponse
	(*GetStringRequest)(nil),  // 2: kvdbserverapi.GetStringRequest
	(*GetStringResponse)(nil), // 3: kvdbserverapi.GetStringResponse
	(*DeleteKeyRequest)(nil),  // 4: kvdbserverapi.DeleteKeyRequest
	(*DeleteKeyResponse)(nil), // 5: kvdbserverapi.DeleteKeyResponse
}
var file_proto_kvdbserver_storage_proto_depIdxs = []int32{
	0, // 0: kvdbserverapi.StorageService.SetString:input_type -> kvdbserverapi.SetStringRequest
	2, // 1: kvdbserverapi.StorageService.GetString:input_type -> kvdbserverapi.GetStringRequest
	4, // 2: kvdbserverapi.StorageService.DeleteKey:input_type -> kvdbserverapi.DeleteKeyRequest
	1, // 3: kvdbserverapi.StorageService.SetString:output_type -> kvdbserverapi.SetStringResponse
	3, // 4: kvdbserverapi.StorageService.GetString:output_type -> kvdbserverapi.GetStringResponse
	5, // 5: kvdbserverapi.StorageService.DeleteKey:output_type -> kvdbserverapi.DeleteKeyResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_kvdbserver_storage_proto_init() }
func file_proto_kvdbserver_storage_proto_init() {
	if File_proto_kvdbserver_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_kvdbserver_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetStringRequest); i {
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
		file_proto_kvdbserver_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetStringResponse); i {
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
		file_proto_kvdbserver_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStringRequest); i {
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
		file_proto_kvdbserver_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStringResponse); i {
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
		file_proto_kvdbserver_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKeyRequest); i {
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
		file_proto_kvdbserver_storage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteKeyResponse); i {
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
			RawDescriptor: file_proto_kvdbserver_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_kvdbserver_storage_proto_goTypes,
		DependencyIndexes: file_proto_kvdbserver_storage_proto_depIdxs,
		MessageInfos:      file_proto_kvdbserver_storage_proto_msgTypes,
	}.Build()
	File_proto_kvdbserver_storage_proto = out.File
	file_proto_kvdbserver_storage_proto_rawDesc = nil
	file_proto_kvdbserver_storage_proto_goTypes = nil
	file_proto_kvdbserver_storage_proto_depIdxs = nil
}

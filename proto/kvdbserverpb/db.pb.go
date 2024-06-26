// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: proto/kvdbserverpb/db.proto

package kvdbserverpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type CreateDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the database to create.
	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}

func (x *CreateDatabaseRequest) Reset() {
	*x = CreateDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserverpb_db_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatabaseRequest) ProtoMessage() {}

func (x *CreateDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserverpb_db_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatabaseRequest.ProtoReflect.Descriptor instead.
func (*CreateDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserverpb_db_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDatabaseRequest) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

type CreateDatabaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the created database.
	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}

func (x *CreateDatabaseResponse) Reset() {
	*x = CreateDatabaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserverpb_db_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDatabaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatabaseResponse) ProtoMessage() {}

func (x *CreateDatabaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserverpb_db_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatabaseResponse.ProtoReflect.Descriptor instead.
func (*CreateDatabaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserverpb_db_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDatabaseResponse) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

type GetAllDatabasesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllDatabasesRequest) Reset() {
	*x = GetAllDatabasesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserverpb_db_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDatabasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDatabasesRequest) ProtoMessage() {}

func (x *GetAllDatabasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserverpb_db_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDatabasesRequest.ProtoReflect.Descriptor instead.
func (*GetAllDatabasesRequest) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserverpb_db_proto_rawDescGZIP(), []int{2}
}

type GetAllDatabasesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of returned database names.
	DbNames []string `protobuf:"bytes,1,rep,name=db_names,json=dbNames,proto3" json:"db_names,omitempty"`
}

func (x *GetAllDatabasesResponse) Reset() {
	*x = GetAllDatabasesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserverpb_db_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDatabasesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDatabasesResponse) ProtoMessage() {}

func (x *GetAllDatabasesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserverpb_db_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDatabasesResponse.ProtoReflect.Descriptor instead.
func (*GetAllDatabasesResponse) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserverpb_db_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllDatabasesResponse) GetDbNames() []string {
	if x != nil {
		return x.DbNames
	}
	return nil
}

type GetDatabaseInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the database.
	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}

func (x *GetDatabaseInfoRequest) Reset() {
	*x = GetDatabaseInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserverpb_db_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDatabaseInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatabaseInfoRequest) ProtoMessage() {}

func (x *GetDatabaseInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserverpb_db_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatabaseInfoRequest.ProtoReflect.Descriptor instead.
func (*GetDatabaseInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserverpb_db_proto_rawDescGZIP(), []int{4}
}

func (x *GetDatabaseInfoRequest) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

type GetDatabaseInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Information about the database.
	Data *DatabaseInfo `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetDatabaseInfoResponse) Reset() {
	*x = GetDatabaseInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserverpb_db_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDatabaseInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatabaseInfoResponse) ProtoMessage() {}

func (x *GetDatabaseInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserverpb_db_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatabaseInfoResponse.ProtoReflect.Descriptor instead.
func (*GetDatabaseInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserverpb_db_proto_rawDescGZIP(), []int{5}
}

func (x *GetDatabaseInfoResponse) GetData() *DatabaseInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteDatabaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the database.
	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}

func (x *DeleteDatabaseRequest) Reset() {
	*x = DeleteDatabaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserverpb_db_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatabaseRequest) ProtoMessage() {}

func (x *DeleteDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserverpb_db_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatabaseRequest.ProtoReflect.Descriptor instead.
func (*DeleteDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserverpb_db_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteDatabaseRequest) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

type DeleteDatabaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the deleted database.
	DbName string `protobuf:"bytes,1,opt,name=db_name,json=dbName,proto3" json:"db_name,omitempty"`
}

func (x *DeleteDatabaseResponse) Reset() {
	*x = DeleteDatabaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserverpb_db_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteDatabaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatabaseResponse) ProtoMessage() {}

func (x *DeleteDatabaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserverpb_db_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatabaseResponse.ProtoReflect.Descriptor instead.
func (*DeleteDatabaseResponse) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserverpb_db_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteDatabaseResponse) GetDbName() string {
	if x != nil {
		return x.DbName
	}
	return ""
}

// DatabaseInfo represents information about a database.
type DatabaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the database.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// UTC timestamp when the database was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// UTC timestamp when the database was updated.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Size of the data in bytes.
	DataSize uint64 `protobuf:"varint,4,opt,name=data_size,json=dataSize,proto3" json:"data_size,omitempty"`
	// Number of keys in the database.
	KeyCount uint32 `protobuf:"varint,5,opt,name=key_count,json=keyCount,proto3" json:"key_count,omitempty"`
}

func (x *DatabaseInfo) Reset() {
	*x = DatabaseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_kvdbserverpb_db_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseInfo) ProtoMessage() {}

func (x *DatabaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_kvdbserverpb_db_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseInfo.ProtoReflect.Descriptor instead.
func (*DatabaseInfo) Descriptor() ([]byte, []int) {
	return file_proto_kvdbserverpb_db_proto_rawDescGZIP(), []int{8}
}

func (x *DatabaseInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DatabaseInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DatabaseInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *DatabaseInfo) GetDataSize() uint64 {
	if x != nil {
		return x.DataSize
	}
	return 0
}

func (x *DatabaseInfo) GetKeyCount() uint32 {
	if x != nil {
		return x.KeyCount
	}
	return 0
}

var File_proto_kvdbserverpb_db_proto protoreflect.FileDescriptor

var file_proto_kvdbserverpb_db_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x70, 0x62, 0x2f, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6b,
	0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0x31, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x62, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x62, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x22, 0x31, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x30, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x62,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x64, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xd2, 0x01, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6b, 0x65, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x9b, 0x03, 0x0a, 0x0f,
	0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5f, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x12, 0x24, 0x2e, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x62, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6b, 0x76, 0x64,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x2e, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x24, 0x2e, 0x6b, 0x76, 0x64,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6b, 0x76, 0x64, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_kvdbserverpb_db_proto_rawDescOnce sync.Once
	file_proto_kvdbserverpb_db_proto_rawDescData = file_proto_kvdbserverpb_db_proto_rawDesc
)

func file_proto_kvdbserverpb_db_proto_rawDescGZIP() []byte {
	file_proto_kvdbserverpb_db_proto_rawDescOnce.Do(func() {
		file_proto_kvdbserverpb_db_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_kvdbserverpb_db_proto_rawDescData)
	})
	return file_proto_kvdbserverpb_db_proto_rawDescData
}

var file_proto_kvdbserverpb_db_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_kvdbserverpb_db_proto_goTypes = []interface{}{
	(*CreateDatabaseRequest)(nil),   // 0: kvdbserverapi.CreateDatabaseRequest
	(*CreateDatabaseResponse)(nil),  // 1: kvdbserverapi.CreateDatabaseResponse
	(*GetAllDatabasesRequest)(nil),  // 2: kvdbserverapi.GetAllDatabasesRequest
	(*GetAllDatabasesResponse)(nil), // 3: kvdbserverapi.GetAllDatabasesResponse
	(*GetDatabaseInfoRequest)(nil),  // 4: kvdbserverapi.GetDatabaseInfoRequest
	(*GetDatabaseInfoResponse)(nil), // 5: kvdbserverapi.GetDatabaseInfoResponse
	(*DeleteDatabaseRequest)(nil),   // 6: kvdbserverapi.DeleteDatabaseRequest
	(*DeleteDatabaseResponse)(nil),  // 7: kvdbserverapi.DeleteDatabaseResponse
	(*DatabaseInfo)(nil),            // 8: kvdbserverapi.DatabaseInfo
	(*timestamppb.Timestamp)(nil),   // 9: google.protobuf.Timestamp
}
var file_proto_kvdbserverpb_db_proto_depIdxs = []int32{
	8, // 0: kvdbserverapi.GetDatabaseInfoResponse.data:type_name -> kvdbserverapi.DatabaseInfo
	9, // 1: kvdbserverapi.DatabaseInfo.created_at:type_name -> google.protobuf.Timestamp
	9, // 2: kvdbserverapi.DatabaseInfo.updated_at:type_name -> google.protobuf.Timestamp
	0, // 3: kvdbserverapi.DatabaseService.CreateDatabase:input_type -> kvdbserverapi.CreateDatabaseRequest
	2, // 4: kvdbserverapi.DatabaseService.GetAllDatabases:input_type -> kvdbserverapi.GetAllDatabasesRequest
	4, // 5: kvdbserverapi.DatabaseService.GetDatabaseInfo:input_type -> kvdbserverapi.GetDatabaseInfoRequest
	6, // 6: kvdbserverapi.DatabaseService.DeleteDatabase:input_type -> kvdbserverapi.DeleteDatabaseRequest
	1, // 7: kvdbserverapi.DatabaseService.CreateDatabase:output_type -> kvdbserverapi.CreateDatabaseResponse
	3, // 8: kvdbserverapi.DatabaseService.GetAllDatabases:output_type -> kvdbserverapi.GetAllDatabasesResponse
	5, // 9: kvdbserverapi.DatabaseService.GetDatabaseInfo:output_type -> kvdbserverapi.GetDatabaseInfoResponse
	7, // 10: kvdbserverapi.DatabaseService.DeleteDatabase:output_type -> kvdbserverapi.DeleteDatabaseResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_kvdbserverpb_db_proto_init() }
func file_proto_kvdbserverpb_db_proto_init() {
	if File_proto_kvdbserverpb_db_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_kvdbserverpb_db_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDatabaseRequest); i {
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
		file_proto_kvdbserverpb_db_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDatabaseResponse); i {
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
		file_proto_kvdbserverpb_db_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDatabasesRequest); i {
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
		file_proto_kvdbserverpb_db_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDatabasesResponse); i {
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
		file_proto_kvdbserverpb_db_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDatabaseInfoRequest); i {
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
		file_proto_kvdbserverpb_db_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDatabaseInfoResponse); i {
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
		file_proto_kvdbserverpb_db_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDatabaseRequest); i {
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
		file_proto_kvdbserverpb_db_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteDatabaseResponse); i {
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
		file_proto_kvdbserverpb_db_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseInfo); i {
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
			RawDescriptor: file_proto_kvdbserverpb_db_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_kvdbserverpb_db_proto_goTypes,
		DependencyIndexes: file_proto_kvdbserverpb_db_proto_depIdxs,
		MessageInfos:      file_proto_kvdbserverpb_db_proto_msgTypes,
	}.Build()
	File_proto_kvdbserverpb_db_proto = out.File
	file_proto_kvdbserverpb_db_proto_rawDesc = nil
	file_proto_kvdbserverpb_db_proto_goTypes = nil
	file_proto_kvdbserverpb_db_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: api/storage/v1/volume.proto

package storageV1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	v1 "github.com/realotz/mstore/api/core/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// 类型 0 全部 1 目录 2 字符串
	Type uint32 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ListFileReq) Reset() {
	*x = ListFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_storage_v1_volume_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFileReq) ProtoMessage() {}

func (x *ListFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_storage_v1_volume_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFileReq.ProtoReflect.Descriptor instead.
func (*ListFileReq) Descriptor() ([]byte, []int) {
	return file_api_storage_v1_volume_proto_rawDescGZIP(), []int{0}
}

func (x *ListFileReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListFileReq) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ListFileReq) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

// 存储卷列表
type ListFileReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*File `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int64   `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListFileReply) Reset() {
	*x = ListFileReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_storage_v1_volume_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFileReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFileReply) ProtoMessage() {}

func (x *ListFileReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_storage_v1_volume_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFileReply.ProtoReflect.Descriptor instead.
func (*ListFileReply) Descriptor() ([]byte, []int) {
	return file_api_storage_v1_volume_proto_rawDescGZIP(), []int{1}
}

func (x *ListFileReply) GetList() []*File {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ListFileReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size      int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Path      string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Ext       string `protobuf:"bytes,4,opt,name=ext,proto3" json:"ext,omitempty"`
	IsDir     bool   `protobuf:"varint,5,opt,name=is_dir,json=isDir,proto3" json:"is_dir,omitempty"`
	UpdatedAt int64  `protobuf:"varint,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_storage_v1_volume_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_api_storage_v1_volume_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_api_storage_v1_volume_proto_rawDescGZIP(), []int{2}
}

func (x *File) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *File) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *File) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

func (x *File) GetIsDir() bool {
	if x != nil {
		return x.IsDir
	}
	return false
}

func (x *File) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

// 存储卷列表请求
type ListVolumeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListVolumeReq) Reset() {
	*x = ListVolumeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_storage_v1_volume_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVolumeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVolumeReq) ProtoMessage() {}

func (x *ListVolumeReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_storage_v1_volume_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVolumeReq.ProtoReflect.Descriptor instead.
func (*ListVolumeReq) Descriptor() ([]byte, []int) {
	return file_api_storage_v1_volume_proto_rawDescGZIP(), []int{3}
}

// 存储卷列表
type ListVolumeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List  []*Volume `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Total int64     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListVolumeReply) Reset() {
	*x = ListVolumeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_storage_v1_volume_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVolumeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVolumeReply) ProtoMessage() {}

func (x *ListVolumeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_storage_v1_volume_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVolumeReply.ProtoReflect.Descriptor instead.
func (*ListVolumeReply) Descriptor() ([]byte, []int) {
	return file_api_storage_v1_volume_proto_rawDescGZIP(), []int{4}
}

func (x *ListVolumeReply) GetList() []*Volume {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *ListVolumeReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type DeleteVolumeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteVolumeReq) Reset() {
	*x = DeleteVolumeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_storage_v1_volume_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteVolumeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteVolumeReq) ProtoMessage() {}

func (x *DeleteVolumeReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_storage_v1_volume_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteVolumeReq.ProtoReflect.Descriptor instead.
func (*DeleteVolumeReq) Descriptor() ([]byte, []int) {
	return file_api_storage_v1_volume_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteVolumeReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateVolumeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 卷名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 供应类型 local 本地 s3 oss cos nfs
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	// 配置 json string
	ProviderConfig string `protobuf:"bytes,3,opt,name=provider_config,json=providerConfig,proto3" json:"provider_config,omitempty"`
}

func (x *CreateVolumeReq) Reset() {
	*x = CreateVolumeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_storage_v1_volume_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVolumeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVolumeReq) ProtoMessage() {}

func (x *CreateVolumeReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_storage_v1_volume_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVolumeReq.ProtoReflect.Descriptor instead.
func (*CreateVolumeReq) Descriptor() ([]byte, []int) {
	return file_api_storage_v1_volume_proto_rawDescGZIP(), []int{6}
}

func (x *CreateVolumeReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateVolumeReq) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *CreateVolumeReq) GetProviderConfig() string {
	if x != nil {
		return x.ProviderConfig
	}
	return ""
}

type Volume struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 卷名
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// 供应类型 local 本地 s3 oss cos nfs
	Provider string `protobuf:"bytes,2,opt,name=provider,proto3" json:"provider,omitempty"`
	// 配置 json string
	ProviderConfig string `protobuf:"bytes,3,opt,name=provider_config,json=providerConfig,proto3" json:"provider_config,omitempty"`
	// 卷id
	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	// 创建时间
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Volume) Reset() {
	*x = Volume{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_storage_v1_volume_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Volume) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Volume) ProtoMessage() {}

func (x *Volume) ProtoReflect() protoreflect.Message {
	mi := &file_api_storage_v1_volume_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Volume.ProtoReflect.Descriptor instead.
func (*Volume) Descriptor() ([]byte, []int) {
	return file_api_storage_v1_volume_proto_rawDescGZIP(), []int{7}
}

func (x *Volume) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Volume) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Volume) GetProviderConfig() string {
	if x != nil {
		return x.ProviderConfig
	}
	return ""
}

func (x *Volume) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Volume) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Volume) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_api_storage_v1_volume_proto protoreflect.FileDescriptor

var file_api_storage_v1_volume_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x22, 0x5e, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x37, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x8a, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73,
	0x5f, 0x64, 0x69, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x44, 0x69,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x0f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x22, 0x62, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x39, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x21, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6a, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0xe7, 0x01, 0x0a, 0x06, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x92,
	0x04, 0x0a, 0x0d, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x80, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x12, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x2e, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13,
	0x22, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x3a, 0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x12, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65,
	0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a,
	0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x82, 0x01, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x12, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x6f, 0x6c, 0x75, 0x6d,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x87, 0x01, 0x0a, 0x08, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x2c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x65, 0x61, 0x6c, 0x6f, 0x74, 0x7a, 0x2f, 0x6d, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_storage_v1_volume_proto_rawDescOnce sync.Once
	file_api_storage_v1_volume_proto_rawDescData = file_api_storage_v1_volume_proto_rawDesc
)

func file_api_storage_v1_volume_proto_rawDescGZIP() []byte {
	file_api_storage_v1_volume_proto_rawDescOnce.Do(func() {
		file_api_storage_v1_volume_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_storage_v1_volume_proto_rawDescData)
	})
	return file_api_storage_v1_volume_proto_rawDescData
}

var file_api_storage_v1_volume_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_storage_v1_volume_proto_goTypes = []interface{}{
	(*ListFileReq)(nil),           // 0: api.service.storage.v1.volume.ListFileReq
	(*ListFileReply)(nil),         // 1: api.service.storage.v1.volume.ListFileReply
	(*File)(nil),                  // 2: api.service.storage.v1.volume.File
	(*ListVolumeReq)(nil),         // 3: api.service.storage.v1.volume.ListVolumeReq
	(*ListVolumeReply)(nil),       // 4: api.service.storage.v1.volume.ListVolumeReply
	(*DeleteVolumeReq)(nil),       // 5: api.service.storage.v1.volume.DeleteVolumeReq
	(*CreateVolumeReq)(nil),       // 6: api.service.storage.v1.volume.CreateVolumeReq
	(*Volume)(nil),                // 7: api.service.storage.v1.volume.Volume
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*v1.Empty)(nil),              // 9: api.core.v1.Empty
}
var file_api_storage_v1_volume_proto_depIdxs = []int32{
	2, // 0: api.service.storage.v1.volume.ListFileReply.list:type_name -> api.service.storage.v1.volume.File
	7, // 1: api.service.storage.v1.volume.ListVolumeReply.list:type_name -> api.service.storage.v1.volume.Volume
	8, // 2: api.service.storage.v1.volume.Volume.created_at:type_name -> google.protobuf.Timestamp
	8, // 3: api.service.storage.v1.volume.Volume.updated_at:type_name -> google.protobuf.Timestamp
	6, // 4: api.service.storage.v1.volume.VolumeService.CreateVolume:input_type -> api.service.storage.v1.volume.CreateVolumeReq
	5, // 5: api.service.storage.v1.volume.VolumeService.DeleteVolume:input_type -> api.service.storage.v1.volume.DeleteVolumeReq
	3, // 6: api.service.storage.v1.volume.VolumeService.ListVolume:input_type -> api.service.storage.v1.volume.ListVolumeReq
	0, // 7: api.service.storage.v1.volume.VolumeService.ListFile:input_type -> api.service.storage.v1.volume.ListFileReq
	7, // 8: api.service.storage.v1.volume.VolumeService.CreateVolume:output_type -> api.service.storage.v1.volume.Volume
	9, // 9: api.service.storage.v1.volume.VolumeService.DeleteVolume:output_type -> api.core.v1.Empty
	4, // 10: api.service.storage.v1.volume.VolumeService.ListVolume:output_type -> api.service.storage.v1.volume.ListVolumeReply
	1, // 11: api.service.storage.v1.volume.VolumeService.ListFile:output_type -> api.service.storage.v1.volume.ListFileReply
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_storage_v1_volume_proto_init() }
func file_api_storage_v1_volume_proto_init() {
	if File_api_storage_v1_volume_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_storage_v1_volume_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFileReq); i {
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
		file_api_storage_v1_volume_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFileReply); i {
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
		file_api_storage_v1_volume_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
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
		file_api_storage_v1_volume_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVolumeReq); i {
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
		file_api_storage_v1_volume_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVolumeReply); i {
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
		file_api_storage_v1_volume_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteVolumeReq); i {
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
		file_api_storage_v1_volume_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVolumeReq); i {
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
		file_api_storage_v1_volume_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Volume); i {
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
			RawDescriptor: file_api_storage_v1_volume_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_storage_v1_volume_proto_goTypes,
		DependencyIndexes: file_api_storage_v1_volume_proto_depIdxs,
		MessageInfos:      file_api_storage_v1_volume_proto_msgTypes,
	}.Build()
	File_api_storage_v1_volume_proto = out.File
	file_api_storage_v1_volume_proto_rawDesc = nil
	file_api_storage_v1_volume_proto_goTypes = nil
	file_api_storage_v1_volume_proto_depIdxs = nil
}

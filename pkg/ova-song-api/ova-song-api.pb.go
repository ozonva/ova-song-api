// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: ova-song-api.proto

package ova_song_api

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

type CreateSongV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Author string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Year   int32  `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *CreateSongV1Request) Reset() {
	*x = CreateSongV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_song_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSongV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSongV1Request) ProtoMessage() {}

func (x *CreateSongV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ova_song_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSongV1Request.ProtoReflect.Descriptor instead.
func (*CreateSongV1Request) Descriptor() ([]byte, []int) {
	return file_ova_song_api_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSongV1Request) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSongV1Request) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *CreateSongV1Request) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

type CreateSongV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongId uint64 `protobuf:"varint,1,opt,name=song_id,json=songId,proto3" json:"song_id,omitempty"`
}

func (x *CreateSongV1Response) Reset() {
	*x = CreateSongV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_song_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSongV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSongV1Response) ProtoMessage() {}

func (x *CreateSongV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ova_song_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSongV1Response.ProtoReflect.Descriptor instead.
func (*CreateSongV1Response) Descriptor() ([]byte, []int) {
	return file_ova_song_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSongV1Response) GetSongId() uint64 {
	if x != nil {
		return x.SongId
	}
	return 0
}

type DescribeSongV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongId uint64 `protobuf:"varint,1,opt,name=song_id,json=songId,proto3" json:"song_id,omitempty"`
}

func (x *DescribeSongV1Request) Reset() {
	*x = DescribeSongV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_song_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSongV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSongV1Request) ProtoMessage() {}

func (x *DescribeSongV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ova_song_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSongV1Request.ProtoReflect.Descriptor instead.
func (*DescribeSongV1Request) Descriptor() ([]byte, []int) {
	return file_ova_song_api_proto_rawDescGZIP(), []int{2}
}

func (x *DescribeSongV1Request) GetSongId() uint64 {
	if x != nil {
		return x.SongId
	}
	return 0
}

type DescribeSongV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Song *Song `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
}

func (x *DescribeSongV1Response) Reset() {
	*x = DescribeSongV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_song_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeSongV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeSongV1Response) ProtoMessage() {}

func (x *DescribeSongV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ova_song_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeSongV1Response.ProtoReflect.Descriptor instead.
func (*DescribeSongV1Response) Descriptor() ([]byte, []int) {
	return file_ova_song_api_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeSongV1Response) GetSong() *Song {
	if x != nil {
		return x.Song
	}
	return nil
}

type ListSongsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset uint64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  uint64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListSongsV1Request) Reset() {
	*x = ListSongsV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_song_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSongsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSongsV1Request) ProtoMessage() {}

func (x *ListSongsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ova_song_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSongsV1Request.ProtoReflect.Descriptor instead.
func (*ListSongsV1Request) Descriptor() ([]byte, []int) {
	return file_ova_song_api_proto_rawDescGZIP(), []int{4}
}

func (x *ListSongsV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListSongsV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListSongsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Songs []*Song `protobuf:"bytes,1,rep,name=songs,proto3" json:"songs,omitempty"`
}

func (x *ListSongsV1Response) Reset() {
	*x = ListSongsV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_song_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSongsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSongsV1Response) ProtoMessage() {}

func (x *ListSongsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ova_song_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSongsV1Response.ProtoReflect.Descriptor instead.
func (*ListSongsV1Response) Descriptor() ([]byte, []int) {
	return file_ova_song_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListSongsV1Response) GetSongs() []*Song {
	if x != nil {
		return x.Songs
	}
	return nil
}

type RemoveSongV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SongId uint64 `protobuf:"varint,1,opt,name=song_id,json=songId,proto3" json:"song_id,omitempty"`
}

func (x *RemoveSongV1Request) Reset() {
	*x = RemoveSongV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_song_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveSongV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveSongV1Request) ProtoMessage() {}

func (x *RemoveSongV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_ova_song_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveSongV1Request.ProtoReflect.Descriptor instead.
func (*RemoveSongV1Request) Descriptor() ([]byte, []int) {
	return file_ova_song_api_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveSongV1Request) GetSongId() uint64 {
	if x != nil {
		return x.SongId
	}
	return 0
}

type RemoveSongV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Removed bool `protobuf:"varint,1,opt,name=removed,proto3" json:"removed,omitempty"`
}

func (x *RemoveSongV1Response) Reset() {
	*x = RemoveSongV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_song_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveSongV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveSongV1Response) ProtoMessage() {}

func (x *RemoveSongV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_ova_song_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveSongV1Response.ProtoReflect.Descriptor instead.
func (*RemoveSongV1Response) Descriptor() ([]byte, []int) {
	return file_ova_song_api_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveSongV1Response) GetRemoved() bool {
	if x != nil {
		return x.Removed
	}
	return false
}

type Song struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Year   int32  `protobuf:"varint,4,opt,name=year,proto3" json:"year,omitempty"`
}

func (x *Song) Reset() {
	*x = Song{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ova_song_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Song) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Song) ProtoMessage() {}

func (x *Song) ProtoReflect() protoreflect.Message {
	mi := &file_ova_song_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Song.ProtoReflect.Descriptor instead.
func (*Song) Descriptor() ([]byte, []int) {
	return file_ova_song_api_proto_rawDescGZIP(), []int{8}
}

func (x *Song) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Song) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Song) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Song) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

var File_ova_song_api_proto protoreflect.FileDescriptor

var file_ova_song_api_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6f, 0x76, 0x61, 0x2d, 0x73, 0x6f, 0x6e, 0x67, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6f, 0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61,
	0x70, 0x69, 0x22, 0x55, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22, 0x2f, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x73, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x15, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x6f, 0x6e, 0x67, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x16,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x56, 0x31, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x22, 0x42,
	0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x3f, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x56,
	0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x6f, 0x6e,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x05, 0x73, 0x6f,
	0x6e, 0x67, 0x73, 0x22, 0x2e, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x6f, 0x6e,
	0x67, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x6f,
	0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x6f, 0x6e,
	0x67, 0x49, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x6f, 0x6e,
	0x67, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x64, 0x22, 0x56, 0x0a, 0x04, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x32, 0xf3, 0x02,
	0x0a, 0x0a, 0x4f, 0x76, 0x61, 0x53, 0x6f, 0x6e, 0x67, 0x41, 0x70, 0x69, 0x12, 0x57, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x56, 0x31, 0x12, 0x21, 0x2e, 0x6f,
	0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x0e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62,
	0x65, 0x53, 0x6f, 0x6e, 0x67, 0x56, 0x31, 0x12, 0x23, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x53,
	0x6f, 0x6e, 0x67, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6f,
	0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x6e, 0x67,
	0x73, 0x56, 0x31, 0x12, 0x20, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x56, 0x31, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6f, 0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x6f, 0x6e, 0x67, 0x73, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x0c, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x56, 0x31, 0x12, 0x21, 0x2e, 0x6f, 0x76, 0x61,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x53, 0x6f, 0x6e, 0x67, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x6f, 0x76, 0x61, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x76, 0x61, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x73, 0x6f, 0x6e,
	0x67, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6f, 0x76, 0x61, 0x2d, 0x73, 0x6f,
	0x6e, 0x67, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x76, 0x61, 0x5f, 0x73, 0x6f, 0x6e, 0x67, 0x5f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ova_song_api_proto_rawDescOnce sync.Once
	file_ova_song_api_proto_rawDescData = file_ova_song_api_proto_rawDesc
)

func file_ova_song_api_proto_rawDescGZIP() []byte {
	file_ova_song_api_proto_rawDescOnce.Do(func() {
		file_ova_song_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_ova_song_api_proto_rawDescData)
	})
	return file_ova_song_api_proto_rawDescData
}

var file_ova_song_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ova_song_api_proto_goTypes = []interface{}{
	(*CreateSongV1Request)(nil),    // 0: ova.task.api.CreateSongV1Request
	(*CreateSongV1Response)(nil),   // 1: ova.task.api.CreateSongV1Response
	(*DescribeSongV1Request)(nil),  // 2: ova.task.api.DescribeSongV1Request
	(*DescribeSongV1Response)(nil), // 3: ova.task.api.DescribeSongV1Response
	(*ListSongsV1Request)(nil),     // 4: ova.task.api.ListSongsV1Request
	(*ListSongsV1Response)(nil),    // 5: ova.task.api.ListSongsV1Response
	(*RemoveSongV1Request)(nil),    // 6: ova.task.api.RemoveSongV1Request
	(*RemoveSongV1Response)(nil),   // 7: ova.task.api.RemoveSongV1Response
	(*Song)(nil),                   // 8: ova.task.api.Song
}
var file_ova_song_api_proto_depIdxs = []int32{
	8, // 0: ova.task.api.DescribeSongV1Response.song:type_name -> ova.task.api.Song
	8, // 1: ova.task.api.ListSongsV1Response.songs:type_name -> ova.task.api.Song
	0, // 2: ova.task.api.OvaSongApi.CreateSongV1:input_type -> ova.task.api.CreateSongV1Request
	2, // 3: ova.task.api.OvaSongApi.DescribeSongV1:input_type -> ova.task.api.DescribeSongV1Request
	4, // 4: ova.task.api.OvaSongApi.ListSongsV1:input_type -> ova.task.api.ListSongsV1Request
	6, // 5: ova.task.api.OvaSongApi.RemoveSongV1:input_type -> ova.task.api.RemoveSongV1Request
	1, // 6: ova.task.api.OvaSongApi.CreateSongV1:output_type -> ova.task.api.CreateSongV1Response
	3, // 7: ova.task.api.OvaSongApi.DescribeSongV1:output_type -> ova.task.api.DescribeSongV1Response
	5, // 8: ova.task.api.OvaSongApi.ListSongsV1:output_type -> ova.task.api.ListSongsV1Response
	7, // 9: ova.task.api.OvaSongApi.RemoveSongV1:output_type -> ova.task.api.RemoveSongV1Response
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ova_song_api_proto_init() }
func file_ova_song_api_proto_init() {
	if File_ova_song_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ova_song_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSongV1Request); i {
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
		file_ova_song_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSongV1Response); i {
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
		file_ova_song_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSongV1Request); i {
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
		file_ova_song_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeSongV1Response); i {
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
		file_ova_song_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSongsV1Request); i {
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
		file_ova_song_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSongsV1Response); i {
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
		file_ova_song_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveSongV1Request); i {
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
		file_ova_song_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveSongV1Response); i {
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
		file_ova_song_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Song); i {
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
			RawDescriptor: file_ova_song_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ova_song_api_proto_goTypes,
		DependencyIndexes: file_ova_song_api_proto_depIdxs,
		MessageInfos:      file_ova_song_api_proto_msgTypes,
	}.Build()
	File_ova_song_api_proto = out.File
	file_ova_song_api_proto_rawDesc = nil
	file_ova_song_api_proto_goTypes = nil
	file_ova_song_api_proto_depIdxs = nil
}
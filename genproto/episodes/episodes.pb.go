// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: episodes.proto

package episodes

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

type EpisodeCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodcastId   string   `protobuf:"bytes,1,opt,name=podcast_id,json=podcastId,proto3" json:"podcast_id,omitempty"`
	UserId      string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title       string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	FileAudio   []byte   `protobuf:"bytes,4,opt,name=file_audio,json=fileAudio,proto3" json:"file_audio,omitempty"`
	Description string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Duration    int64    `protobuf:"varint,6,opt,name=duration,proto3" json:"duration,omitempty"`
	Genre       string   `protobuf:"bytes,7,opt,name=genre,proto3" json:"genre,omitempty"`
	Tags        []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *EpisodeCreate) Reset() {
	*x = EpisodeCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_episodes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EpisodeCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpisodeCreate) ProtoMessage() {}

func (x *EpisodeCreate) ProtoReflect() protoreflect.Message {
	mi := &file_episodes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpisodeCreate.ProtoReflect.Descriptor instead.
func (*EpisodeCreate) Descriptor() ([]byte, []int) {
	return file_episodes_proto_rawDescGZIP(), []int{0}
}

func (x *EpisodeCreate) GetPodcastId() string {
	if x != nil {
		return x.PodcastId
	}
	return ""
}

func (x *EpisodeCreate) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *EpisodeCreate) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *EpisodeCreate) GetFileAudio() []byte {
	if x != nil {
		return x.FileAudio
	}
	return nil
}

func (x *EpisodeCreate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *EpisodeCreate) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *EpisodeCreate) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *EpisodeCreate) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_episodes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

func (x *Filter) ProtoReflect() protoreflect.Message {
	mi := &file_episodes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_episodes_proto_rawDescGZIP(), []int{1}
}

func (x *Filter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Filter) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Filter) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_episodes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_episodes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_episodes_proto_rawDescGZIP(), []int{2}
}

func (x *ID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Episode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PodcastId   string   `protobuf:"bytes,2,opt,name=podcast_id,json=podcastId,proto3" json:"podcast_id,omitempty"`
	UserId      string   `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title       string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	FileAudio   []byte   `protobuf:"bytes,5,opt,name=file_audio,json=fileAudio,proto3" json:"file_audio,omitempty"`
	Description string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Duration    int64    `protobuf:"varint,7,opt,name=duration,proto3" json:"duration,omitempty"`
	Genre       string   `protobuf:"bytes,8,opt,name=genre,proto3" json:"genre,omitempty"`
	Tags        []string `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	CreatedAt   string   `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string   `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Episode) Reset() {
	*x = Episode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_episodes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Episode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Episode) ProtoMessage() {}

func (x *Episode) ProtoReflect() protoreflect.Message {
	mi := &file_episodes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Episode.ProtoReflect.Descriptor instead.
func (*Episode) Descriptor() ([]byte, []int) {
	return file_episodes_proto_rawDescGZIP(), []int{3}
}

func (x *Episode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Episode) GetPodcastId() string {
	if x != nil {
		return x.PodcastId
	}
	return ""
}

func (x *Episode) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Episode) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Episode) GetFileAudio() []byte {
	if x != nil {
		return x.FileAudio
	}
	return nil
}

func (x *Episode) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Episode) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Episode) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *Episode) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Episode) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Episode) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type Episodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Episodes []*Episode `protobuf:"bytes,1,rep,name=episodes,proto3" json:"episodes,omitempty"`
}

func (x *Episodes) Reset() {
	*x = Episodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_episodes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Episodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Episodes) ProtoMessage() {}

func (x *Episodes) ProtoReflect() protoreflect.Message {
	mi := &file_episodes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Episodes.ProtoReflect.Descriptor instead.
func (*Episodes) Descriptor() ([]byte, []int) {
	return file_episodes_proto_rawDescGZIP(), []int{4}
}

func (x *Episodes) GetEpisodes() []*Episode {
	if x != nil {
		return x.Episodes
	}
	return nil
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_episodes_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_episodes_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_episodes_proto_rawDescGZIP(), []int{5}
}

type IDs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodcastId string         `protobuf:"bytes,1,opt,name=podcast_id,json=podcastId,proto3" json:"podcast_id,omitempty"`
	EpisodeId string         `protobuf:"bytes,2,opt,name=episode_id,json=episodeId,proto3" json:"episode_id,omitempty"`
	Episode   *EpisodeCreate `protobuf:"bytes,3,opt,name=episode,proto3" json:"episode,omitempty"`
}

func (x *IDs) Reset() {
	*x = IDs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_episodes_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDs) ProtoMessage() {}

func (x *IDs) ProtoReflect() protoreflect.Message {
	mi := &file_episodes_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDs.ProtoReflect.Descriptor instead.
func (*IDs) Descriptor() ([]byte, []int) {
	return file_episodes_proto_rawDescGZIP(), []int{6}
}

func (x *IDs) GetPodcastId() string {
	if x != nil {
		return x.PodcastId
	}
	return ""
}

func (x *IDs) GetEpisodeId() string {
	if x != nil {
		return x.EpisodeId
	}
	return ""
}

func (x *IDs) GetEpisode() *EpisodeCreate {
	if x != nil {
		return x.Episode
	}
	return nil
}

type IDsForDelete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PodcastId string `protobuf:"bytes,1,opt,name=podcast_id,json=podcastId,proto3" json:"podcast_id,omitempty"`
	EpisodeId string `protobuf:"bytes,2,opt,name=episode_id,json=episodeId,proto3" json:"episode_id,omitempty"`
}

func (x *IDsForDelete) Reset() {
	*x = IDsForDelete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_episodes_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDsForDelete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDsForDelete) ProtoMessage() {}

func (x *IDsForDelete) ProtoReflect() protoreflect.Message {
	mi := &file_episodes_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDsForDelete.ProtoReflect.Descriptor instead.
func (*IDsForDelete) Descriptor() ([]byte, []int) {
	return file_episodes_proto_rawDescGZIP(), []int{7}
}

func (x *IDsForDelete) GetPodcastId() string {
	if x != nil {
		return x.PodcastId
	}
	return ""
}

func (x *IDsForDelete) GetEpisodeId() string {
	if x != nil {
		return x.EpisodeId
	}
	return ""
}

type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_episodes_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_episodes_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_episodes_proto_rawDescGZIP(), []int{8}
}

func (x *Success) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_episodes_proto protoreflect.FileDescriptor

var file_episodes_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x22, 0xe4, 0x01, 0x0a, 0x0d, 0x45,
	0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x70, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67,
	0x73, 0x22, 0x46, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0xac, 0x02, 0x0a, 0x07, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x39,
	0x0a, 0x08, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x65, 0x70,
	0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65,
	0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x52,
	0x08, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69,
	0x64, 0x22, 0x76, 0x0a, 0x03, 0x49, 0x44, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6f,
	0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x70, 0x69, 0x73, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x70, 0x69,
	0x73, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x07, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64,
	0x65, 0x73, 0x2e, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x07, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x22, 0x4c, 0x0a, 0x0c, 0x49, 0x44, 0x73,
	0x46, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6f, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x70, 0x69, 0x73,
	0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x70,
	0x69, 0x73, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x22, 0x23, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xee, 0x02, 0x0a,
	0x0f, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3f, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x63, 0x61, 0x73,
	0x74, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x12, 0x17, 0x2e, 0x65, 0x70, 0x69, 0x73, 0x6f,
	0x64, 0x65, 0x73, 0x2e, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x1a, 0x0c, 0x2e, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x49, 0x44, 0x22,
	0x00, 0x12, 0x40, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73,
	0x42, 0x79, 0x50, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x10, 0x2e, 0x65, 0x70,
	0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a, 0x12, 0x2e,
	0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x70, 0x69,
	0x73, 0x6f, 0x64, 0x65, 0x12, 0x0d, 0x2e, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x2e,
	0x49, 0x44, 0x73, 0x1a, 0x0e, 0x2e, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x56,
	0x6f, 0x69, 0x64, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45,
	0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x2e, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65,
	0x73, 0x2e, 0x49, 0x44, 0x73, 0x46, 0x6f, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x0e,
	0x2e, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x00,
	0x12, 0x33, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x50, 0x6f, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x12, 0x0c, 0x2e, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x49, 0x44,
	0x1a, 0x11, 0x2e, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x11, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x64, 0x63, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12, 0x0c, 0x2e, 0x65, 0x70, 0x69,
	0x73, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x49, 0x44, 0x1a, 0x11, 0x2e, 0x65, 0x70, 0x69, 0x73, 0x6f,
	0x64, 0x65, 0x73, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x00, 0x42, 0x13, 0x5a,
	0x11, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_episodes_proto_rawDescOnce sync.Once
	file_episodes_proto_rawDescData = file_episodes_proto_rawDesc
)

func file_episodes_proto_rawDescGZIP() []byte {
	file_episodes_proto_rawDescOnce.Do(func() {
		file_episodes_proto_rawDescData = protoimpl.X.CompressGZIP(file_episodes_proto_rawDescData)
	})
	return file_episodes_proto_rawDescData
}

var file_episodes_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_episodes_proto_goTypes = []interface{}{
	(*EpisodeCreate)(nil), // 0: episodes.EpisodeCreate
	(*Filter)(nil),        // 1: episodes.Filter
	(*ID)(nil),            // 2: episodes.ID
	(*Episode)(nil),       // 3: episodes.Episode
	(*Episodes)(nil),      // 4: episodes.Episodes
	(*Void)(nil),          // 5: episodes.Void
	(*IDs)(nil),           // 6: episodes.IDs
	(*IDsForDelete)(nil),  // 7: episodes.IDsForDelete
	(*Success)(nil),       // 8: episodes.Success
}
var file_episodes_proto_depIdxs = []int32{
	3, // 0: episodes.Episodes.episodes:type_name -> episodes.Episode
	0, // 1: episodes.IDs.episode:type_name -> episodes.EpisodeCreate
	0, // 2: episodes.EpisodesService.CreatePodcastEpisode:input_type -> episodes.EpisodeCreate
	1, // 3: episodes.EpisodesService.GetEpisodesByPodcastId:input_type -> episodes.Filter
	6, // 4: episodes.EpisodesService.UpdateEpisode:input_type -> episodes.IDs
	7, // 5: episodes.EpisodesService.DeleteEpisode:input_type -> episodes.IDsForDelete
	2, // 6: episodes.EpisodesService.PublishPodcast:input_type -> episodes.ID
	2, // 7: episodes.EpisodesService.ValidatePodcastId:input_type -> episodes.ID
	2, // 8: episodes.EpisodesService.CreatePodcastEpisode:output_type -> episodes.ID
	4, // 9: episodes.EpisodesService.GetEpisodesByPodcastId:output_type -> episodes.Episodes
	5, // 10: episodes.EpisodesService.UpdateEpisode:output_type -> episodes.Void
	5, // 11: episodes.EpisodesService.DeleteEpisode:output_type -> episodes.Void
	8, // 12: episodes.EpisodesService.PublishPodcast:output_type -> episodes.Success
	8, // 13: episodes.EpisodesService.ValidatePodcastId:output_type -> episodes.Success
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_episodes_proto_init() }
func file_episodes_proto_init() {
	if File_episodes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_episodes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EpisodeCreate); i {
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
		file_episodes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
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
		file_episodes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_episodes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Episode); i {
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
		file_episodes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Episodes); i {
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
		file_episodes_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_episodes_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDs); i {
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
		file_episodes_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDsForDelete); i {
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
		file_episodes_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Success); i {
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
			RawDescriptor: file_episodes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_episodes_proto_goTypes,
		DependencyIndexes: file_episodes_proto_depIdxs,
		MessageInfos:      file_episodes_proto_msgTypes,
	}.Build()
	File_episodes_proto = out.File
	file_episodes_proto_rawDesc = nil
	file_episodes_proto_goTypes = nil
	file_episodes_proto_depIdxs = nil
}

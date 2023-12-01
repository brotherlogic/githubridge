// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: githubridge.proto

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

type CreateIssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Repo  string `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body  string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *CreateIssueRequest) Reset() {
	*x = CreateIssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_githubridge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIssueRequest) ProtoMessage() {}

func (x *CreateIssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_githubridge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIssueRequest.ProtoReflect.Descriptor instead.
func (*CreateIssueRequest) Descriptor() ([]byte, []int) {
	return file_githubridge_proto_rawDescGZIP(), []int{0}
}

func (x *CreateIssueRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *CreateIssueRequest) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *CreateIssueRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateIssueRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type CreateIssueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IssueId int64 `protobuf:"varint,1,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
}

func (x *CreateIssueResponse) Reset() {
	*x = CreateIssueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_githubridge_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIssueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIssueResponse) ProtoMessage() {}

func (x *CreateIssueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_githubridge_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIssueResponse.ProtoReflect.Descriptor instead.
func (*CreateIssueResponse) Descriptor() ([]byte, []int) {
	return file_githubridge_proto_rawDescGZIP(), []int{1}
}

func (x *CreateIssueResponse) GetIssueId() int64 {
	if x != nil {
		return x.IssueId
	}
	return 0
}

type GetIssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Repo string `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	Id   int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetIssueRequest) Reset() {
	*x = GetIssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_githubridge_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIssueRequest) ProtoMessage() {}

func (x *GetIssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_githubridge_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIssueRequest.ProtoReflect.Descriptor instead.
func (*GetIssueRequest) Descriptor() ([]byte, []int) {
	return file_githubridge_proto_rawDescGZIP(), []int{2}
}

func (x *GetIssueRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *GetIssueRequest) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *GetIssueRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetIssueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State    string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	Comments int32  `protobuf:"varint,2,opt,name=comments,proto3" json:"comments,omitempty"`
}

func (x *GetIssueResponse) Reset() {
	*x = GetIssueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_githubridge_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIssueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIssueResponse) ProtoMessage() {}

func (x *GetIssueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_githubridge_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIssueResponse.ProtoReflect.Descriptor instead.
func (*GetIssueResponse) Descriptor() ([]byte, []int) {
	return file_githubridge_proto_rawDescGZIP(), []int{3}
}

func (x *GetIssueResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *GetIssueResponse) GetComments() int32 {
	if x != nil {
		return x.Comments
	}
	return 0
}

type CloseIssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Repo string `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	Id   int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CloseIssueRequest) Reset() {
	*x = CloseIssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_githubridge_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseIssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseIssueRequest) ProtoMessage() {}

func (x *CloseIssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_githubridge_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseIssueRequest.ProtoReflect.Descriptor instead.
func (*CloseIssueRequest) Descriptor() ([]byte, []int) {
	return file_githubridge_proto_rawDescGZIP(), []int{4}
}

func (x *CloseIssueRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *CloseIssueRequest) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *CloseIssueRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CloseIssueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CloseIssueResponse) Reset() {
	*x = CloseIssueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_githubridge_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseIssueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseIssueResponse) ProtoMessage() {}

func (x *CloseIssueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_githubridge_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseIssueResponse.ProtoReflect.Descriptor instead.
func (*CloseIssueResponse) Descriptor() ([]byte, []int) {
	return file_githubridge_proto_rawDescGZIP(), []int{5}
}

type CommentOnIssueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Repo    string `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`
	Id      int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	Comment string `protobuf:"bytes,4,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CommentOnIssueRequest) Reset() {
	*x = CommentOnIssueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_githubridge_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentOnIssueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentOnIssueRequest) ProtoMessage() {}

func (x *CommentOnIssueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_githubridge_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentOnIssueRequest.ProtoReflect.Descriptor instead.
func (*CommentOnIssueRequest) Descriptor() ([]byte, []int) {
	return file_githubridge_proto_rawDescGZIP(), []int{6}
}

func (x *CommentOnIssueRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *CommentOnIssueRequest) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *CommentOnIssueRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentOnIssueRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type CommentOnIssueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommentOnIssueResponse) Reset() {
	*x = CommentOnIssueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_githubridge_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentOnIssueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentOnIssueResponse) ProtoMessage() {}

func (x *CommentOnIssueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_githubridge_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentOnIssueResponse.ProtoReflect.Descriptor instead.
func (*CommentOnIssueResponse) Descriptor() ([]byte, []int) {
	return file_githubridge_proto_rawDescGZIP(), []int{7}
}

var File_githubridge_proto protoreflect.FileDescriptor

var file_githubridge_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x22, 0x66, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65,
	0x70, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x30, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x44, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x4b, 0x0a, 0x11, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x69,
	0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x65, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xd9, 0x02, 0x0a, 0x12, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x72, 0x69,
	0x64, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49,
	0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x72,
	0x69, 0x64, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0a, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x72, 0x69, 0x64, 0x67,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x72, 0x69, 0x64, 0x67,
	0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4f,
	0x6e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x22, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x72,
	0x69, 0x64, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x6e, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x4f, 0x6e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_githubridge_proto_rawDescOnce sync.Once
	file_githubridge_proto_rawDescData = file_githubridge_proto_rawDesc
)

func file_githubridge_proto_rawDescGZIP() []byte {
	file_githubridge_proto_rawDescOnce.Do(func() {
		file_githubridge_proto_rawDescData = protoimpl.X.CompressGZIP(file_githubridge_proto_rawDescData)
	})
	return file_githubridge_proto_rawDescData
}

var file_githubridge_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_githubridge_proto_goTypes = []interface{}{
	(*CreateIssueRequest)(nil),     // 0: githubridge.CreateIssueRequest
	(*CreateIssueResponse)(nil),    // 1: githubridge.CreateIssueResponse
	(*GetIssueRequest)(nil),        // 2: githubridge.GetIssueRequest
	(*GetIssueResponse)(nil),       // 3: githubridge.GetIssueResponse
	(*CloseIssueRequest)(nil),      // 4: githubridge.CloseIssueRequest
	(*CloseIssueResponse)(nil),     // 5: githubridge.CloseIssueResponse
	(*CommentOnIssueRequest)(nil),  // 6: githubridge.CommentOnIssueRequest
	(*CommentOnIssueResponse)(nil), // 7: githubridge.CommentOnIssueResponse
}
var file_githubridge_proto_depIdxs = []int32{
	0, // 0: githubridge.GithubridgeService.CreateIssue:input_type -> githubridge.CreateIssueRequest
	2, // 1: githubridge.GithubridgeService.GetIssue:input_type -> githubridge.GetIssueRequest
	4, // 2: githubridge.GithubridgeService.CloseIssue:input_type -> githubridge.CloseIssueRequest
	6, // 3: githubridge.GithubridgeService.CommentOnIssue:input_type -> githubridge.CommentOnIssueRequest
	1, // 4: githubridge.GithubridgeService.CreateIssue:output_type -> githubridge.CreateIssueResponse
	3, // 5: githubridge.GithubridgeService.GetIssue:output_type -> githubridge.GetIssueResponse
	5, // 6: githubridge.GithubridgeService.CloseIssue:output_type -> githubridge.CloseIssueResponse
	7, // 7: githubridge.GithubridgeService.CommentOnIssue:output_type -> githubridge.CommentOnIssueResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_githubridge_proto_init() }
func file_githubridge_proto_init() {
	if File_githubridge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_githubridge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIssueRequest); i {
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
		file_githubridge_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIssueResponse); i {
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
		file_githubridge_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIssueRequest); i {
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
		file_githubridge_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIssueResponse); i {
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
		file_githubridge_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseIssueRequest); i {
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
		file_githubridge_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseIssueResponse); i {
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
		file_githubridge_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentOnIssueRequest); i {
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
		file_githubridge_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentOnIssueResponse); i {
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
			RawDescriptor: file_githubridge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_githubridge_proto_goTypes,
		DependencyIndexes: file_githubridge_proto_depIdxs,
		MessageInfos:      file_githubridge_proto_msgTypes,
	}.Build()
	File_githubridge_proto = out.File
	file_githubridge_proto_rawDesc = nil
	file_githubridge_proto_goTypes = nil
	file_githubridge_proto_depIdxs = nil
}

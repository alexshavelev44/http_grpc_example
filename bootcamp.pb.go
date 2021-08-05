// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: bootcamp.proto

package http_grpc_example

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootcamp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_bootcamp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_bootcamp_proto_rawDescGZIP(), []int{0}
}

func (x *Member) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Member) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Member) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type GetAllMembersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllMembersRequest) Reset() {
	*x = GetAllMembersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootcamp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMembersRequest) ProtoMessage() {}

func (x *GetAllMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bootcamp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMembersRequest.ProtoReflect.Descriptor instead.
func (*GetAllMembersRequest) Descriptor() ([]byte, []int) {
	return file_bootcamp_proto_rawDescGZIP(), []int{1}
}

type GetAllMembersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members []*Member `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *GetAllMembersResponse) Reset() {
	*x = GetAllMembersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootcamp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllMembersResponse) ProtoMessage() {}

func (x *GetAllMembersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bootcamp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllMembersResponse.ProtoReflect.Descriptor instead.
func (*GetAllMembersResponse) Descriptor() ([]byte, []int) {
	return file_bootcamp_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllMembersResponse) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

type GetMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMemberRequest) Reset() {
	*x = GetMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootcamp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemberRequest) ProtoMessage() {}

func (x *GetMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bootcamp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemberRequest.ProtoReflect.Descriptor instead.
func (*GetMemberRequest) Descriptor() ([]byte, []int) {
	return file_bootcamp_proto_rawDescGZIP(), []int{3}
}

func (x *GetMemberRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMemberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Member *Member `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *GetMemberResponse) Reset() {
	*x = GetMemberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootcamp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMemberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemberResponse) ProtoMessage() {}

func (x *GetMemberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bootcamp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemberResponse.ProtoReflect.Descriptor instead.
func (*GetMemberResponse) Descriptor() ([]byte, []int) {
	return file_bootcamp_proto_rawDescGZIP(), []int{4}
}

func (x *GetMemberResponse) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

type CreateMemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateMemberRequest) Reset() {
	*x = CreateMemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootcamp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMemberRequest) ProtoMessage() {}

func (x *CreateMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bootcamp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMemberRequest.ProtoReflect.Descriptor instead.
func (*CreateMemberRequest) Descriptor() ([]byte, []int) {
	return file_bootcamp_proto_rawDescGZIP(), []int{5}
}

func (x *CreateMemberRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateMemberRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreateMemberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateMemberResponse) Reset() {
	*x = CreateMemberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bootcamp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMemberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMemberResponse) ProtoMessage() {}

func (x *CreateMemberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bootcamp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMemberResponse.ProtoReflect.Descriptor instead.
func (*CreateMemberResponse) Descriptor() ([]byte, []int) {
	return file_bootcamp_proto_rawDescGZIP(), []int{6}
}

func (x *CreateMemberResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_bootcamp_proto protoreflect.FileDescriptor

var file_bootcamp_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x22, 0x42, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x4c, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x22, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x46, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3f, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x32, 0xa7, 0x02, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x74, 0x63, 0x61, 0x6d, 0x70, 0x12,
	0x62, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x12, 0x27, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x23, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x40, 0x5a, 0x3e,
	0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x65, 0x78,
	0x73, 0x68, 0x61, 0x76, 0x65, 0x6c, 0x65, 0x76, 0x34, 0x34, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2d,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x3b, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bootcamp_proto_rawDescOnce sync.Once
	file_bootcamp_proto_rawDescData = file_bootcamp_proto_rawDesc
)

func file_bootcamp_proto_rawDescGZIP() []byte {
	file_bootcamp_proto_rawDescOnce.Do(func() {
		file_bootcamp_proto_rawDescData = protoimpl.X.CompressGZIP(file_bootcamp_proto_rawDescData)
	})
	return file_bootcamp_proto_rawDescData
}

var file_bootcamp_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bootcamp_proto_goTypes = []interface{}{
	(*Member)(nil),                // 0: http_grpc_example.Member
	(*GetAllMembersRequest)(nil),  // 1: http_grpc_example.GetAllMembersRequest
	(*GetAllMembersResponse)(nil), // 2: http_grpc_example.GetAllMembersResponse
	(*GetMemberRequest)(nil),      // 3: http_grpc_example.GetMemberRequest
	(*GetMemberResponse)(nil),     // 4: http_grpc_example.GetMemberResponse
	(*CreateMemberRequest)(nil),   // 5: http_grpc_example.CreateMemberRequest
	(*CreateMemberResponse)(nil),  // 6: http_grpc_example.CreateMemberResponse
}
var file_bootcamp_proto_depIdxs = []int32{
	0, // 0: http_grpc_example.GetAllMembersResponse.members:type_name -> http_grpc_example.Member
	0, // 1: http_grpc_example.GetMemberResponse.member:type_name -> http_grpc_example.Member
	1, // 2: http_grpc_example.Bootcamp.GetAllMembers:input_type -> http_grpc_example.GetAllMembersRequest
	3, // 3: http_grpc_example.Bootcamp.GetMember:input_type -> http_grpc_example.GetMemberRequest
	5, // 4: http_grpc_example.Bootcamp.CreateMember:input_type -> http_grpc_example.CreateMemberRequest
	2, // 5: http_grpc_example.Bootcamp.GetAllMembers:output_type -> http_grpc_example.GetAllMembersResponse
	4, // 6: http_grpc_example.Bootcamp.GetMember:output_type -> http_grpc_example.GetMemberResponse
	6, // 7: http_grpc_example.Bootcamp.CreateMember:output_type -> http_grpc_example.CreateMemberResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_bootcamp_proto_init() }
func file_bootcamp_proto_init() {
	if File_bootcamp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bootcamp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_bootcamp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMembersRequest); i {
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
		file_bootcamp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllMembersResponse); i {
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
		file_bootcamp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMemberRequest); i {
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
		file_bootcamp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMemberResponse); i {
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
		file_bootcamp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMemberRequest); i {
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
		file_bootcamp_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMemberResponse); i {
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
			RawDescriptor: file_bootcamp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bootcamp_proto_goTypes,
		DependencyIndexes: file_bootcamp_proto_depIdxs,
		MessageInfos:      file_bootcamp_proto_msgTypes,
	}.Build()
	File_bootcamp_proto = out.File
	file_bootcamp_proto_rawDesc = nil
	file_bootcamp_proto_goTypes = nil
	file_bootcamp_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BootcampClient is the client API for Bootcamp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BootcampClient interface {
	// Unless you are very sure of what you're doing always follow the pattern
	// rpc {Name} ({Name}Request) returns ({Name}Response);
	// This allows us to extend each RPC with new request and response parameters.
	// Please note that any such extensions must be backward and forward compatible.
	GetAllMembers(ctx context.Context, in *GetAllMembersRequest, opts ...grpc.CallOption) (*GetAllMembersResponse, error)
	GetMember(ctx context.Context, in *GetMemberRequest, opts ...grpc.CallOption) (*GetMemberResponse, error)
	CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...grpc.CallOption) (*CreateMemberResponse, error)
}

type bootcampClient struct {
	cc grpc.ClientConnInterface
}

func NewBootcampClient(cc grpc.ClientConnInterface) BootcampClient {
	return &bootcampClient{cc}
}

func (c *bootcampClient) GetAllMembers(ctx context.Context, in *GetAllMembersRequest, opts ...grpc.CallOption) (*GetAllMembersResponse, error) {
	out := new(GetAllMembersResponse)
	err := c.cc.Invoke(ctx, "/http_grpc_example.Bootcamp/GetAllMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bootcampClient) GetMember(ctx context.Context, in *GetMemberRequest, opts ...grpc.CallOption) (*GetMemberResponse, error) {
	out := new(GetMemberResponse)
	err := c.cc.Invoke(ctx, "/http_grpc_example.Bootcamp/GetMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bootcampClient) CreateMember(ctx context.Context, in *CreateMemberRequest, opts ...grpc.CallOption) (*CreateMemberResponse, error) {
	out := new(CreateMemberResponse)
	err := c.cc.Invoke(ctx, "/http_grpc_example.Bootcamp/CreateMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BootcampServer is the server API for Bootcamp service.
type BootcampServer interface {
	// Unless you are very sure of what you're doing always follow the pattern
	// rpc {Name} ({Name}Request) returns ({Name}Response);
	// This allows us to extend each RPC with new request and response parameters.
	// Please note that any such extensions must be backward and forward compatible.
	GetAllMembers(context.Context, *GetAllMembersRequest) (*GetAllMembersResponse, error)
	GetMember(context.Context, *GetMemberRequest) (*GetMemberResponse, error)
	CreateMember(context.Context, *CreateMemberRequest) (*CreateMemberResponse, error)
}

// UnimplementedBootcampServer can be embedded to have forward compatible implementations.
type UnimplementedBootcampServer struct {
}

func (*UnimplementedBootcampServer) GetAllMembers(context.Context, *GetAllMembersRequest) (*GetAllMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMembers not implemented")
}
func (*UnimplementedBootcampServer) GetMember(context.Context, *GetMemberRequest) (*GetMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMember not implemented")
}
func (*UnimplementedBootcampServer) CreateMember(context.Context, *CreateMemberRequest) (*CreateMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMember not implemented")
}

func RegisterBootcampServer(s *grpc.Server, srv BootcampServer) {
	s.RegisterService(&_Bootcamp_serviceDesc, srv)
}

func _Bootcamp_GetAllMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootcampServer).GetAllMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/http_grpc_example.Bootcamp/GetAllMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootcampServer).GetAllMembers(ctx, req.(*GetAllMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bootcamp_GetMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootcampServer).GetMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/http_grpc_example.Bootcamp/GetMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootcampServer).GetMember(ctx, req.(*GetMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bootcamp_CreateMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootcampServer).CreateMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/http_grpc_example.Bootcamp/CreateMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootcampServer).CreateMember(ctx, req.(*CreateMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bootcamp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "http_grpc_example.Bootcamp",
	HandlerType: (*BootcampServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllMembers",
			Handler:    _Bootcamp_GetAllMembers_Handler,
		},
		{
			MethodName: "GetMember",
			Handler:    _Bootcamp_GetMember_Handler,
		},
		{
			MethodName: "CreateMember",
			Handler:    _Bootcamp_CreateMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bootcamp.proto",
}

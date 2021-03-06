// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.1
// source: cloudtrail.proto

package aws

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

type GetEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEventsRequest) Reset() {
	*x = GetEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudtrail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventsRequest) ProtoMessage() {}

func (x *GetEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudtrail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventsRequest.ProtoReflect.Descriptor instead.
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return file_cloudtrail_proto_rawDescGZIP(), []int{0}
}

type CloudTrailEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method    string                 `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TrailId   int32                  `protobuf:"varint,3,opt,name=trail_id,json=trailId,proto3" json:"trail_id,omitempty"`
	Sid       string                 `protobuf:"bytes,4,opt,name=sid,proto3" json:"sid,omitempty"`
	Region    string                 `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *CloudTrailEvent) Reset() {
	*x = CloudTrailEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudtrail_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudTrailEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudTrailEvent) ProtoMessage() {}

func (x *CloudTrailEvent) ProtoReflect() protoreflect.Message {
	mi := &file_cloudtrail_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudTrailEvent.ProtoReflect.Descriptor instead.
func (*CloudTrailEvent) Descriptor() ([]byte, []int) {
	return file_cloudtrail_proto_rawDescGZIP(), []int{1}
}

func (x *CloudTrailEvent) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *CloudTrailEvent) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *CloudTrailEvent) GetTrailId() int32 {
	if x != nil {
		return x.TrailId
	}
	return 0
}

func (x *CloudTrailEvent) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *CloudTrailEvent) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

type CloudTrailConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Write bool `protobuf:"varint,1,opt,name=write,proto3" json:"write,omitempty"`
	Read  bool `protobuf:"varint,2,opt,name=read,proto3" json:"read,omitempty"`
}

func (x *CloudTrailConfig) Reset() {
	*x = CloudTrailConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudtrail_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloudTrailConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloudTrailConfig) ProtoMessage() {}

func (x *CloudTrailConfig) ProtoReflect() protoreflect.Message {
	mi := &file_cloudtrail_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloudTrailConfig.ProtoReflect.Descriptor instead.
func (*CloudTrailConfig) Descriptor() ([]byte, []int) {
	return file_cloudtrail_proto_rawDescGZIP(), []int{2}
}

func (x *CloudTrailConfig) GetWrite() bool {
	if x != nil {
		return x.Write
	}
	return false
}

func (x *CloudTrailConfig) GetRead() bool {
	if x != nil {
		return x.Read
	}
	return false
}

type TrailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Config *CloudTrailConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *TrailRequest) Reset() {
	*x = TrailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudtrail_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrailRequest) ProtoMessage() {}

func (x *TrailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudtrail_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrailRequest.ProtoReflect.Descriptor instead.
func (*TrailRequest) Descriptor() ([]byte, []int) {
	return file_cloudtrail_proto_rawDescGZIP(), []int{3}
}

func (x *TrailRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TrailRequest) GetConfig() *CloudTrailConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type CreateTrailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arn string `protobuf:"bytes,1,opt,name=arn,proto3" json:"arn,omitempty"`
}

func (x *CreateTrailResponse) Reset() {
	*x = CreateTrailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudtrail_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTrailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTrailResponse) ProtoMessage() {}

func (x *CreateTrailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloudtrail_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTrailResponse.ProtoReflect.Descriptor instead.
func (*CreateTrailResponse) Descriptor() ([]byte, []int) {
	return file_cloudtrail_proto_rawDescGZIP(), []int{4}
}

func (x *CreateTrailResponse) GetArn() string {
	if x != nil {
		return x.Arn
	}
	return ""
}

type UpdateTrailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Arn string `protobuf:"bytes,1,opt,name=arn,proto3" json:"arn,omitempty"`
}

func (x *UpdateTrailResponse) Reset() {
	*x = UpdateTrailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudtrail_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTrailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTrailResponse) ProtoMessage() {}

func (x *UpdateTrailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloudtrail_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTrailResponse.ProtoReflect.Descriptor instead.
func (*UpdateTrailResponse) Descriptor() ([]byte, []int) {
	return file_cloudtrail_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTrailResponse) GetArn() string {
	if x != nil {
		return x.Arn
	}
	return ""
}

type DeleteTrailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Ok      bool   `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *DeleteTrailResponse) Reset() {
	*x = DeleteTrailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudtrail_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTrailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTrailResponse) ProtoMessage() {}

func (x *DeleteTrailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cloudtrail_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTrailResponse.ProtoReflect.Descriptor instead.
func (*DeleteTrailResponse) Descriptor() ([]byte, []int) {
	return file_cloudtrail_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTrailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *DeleteTrailResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type DeleteTrailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteTrailRequest) Reset() {
	*x = DeleteTrailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudtrail_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTrailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTrailRequest) ProtoMessage() {}

func (x *DeleteTrailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cloudtrail_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTrailRequest.ProtoReflect.Descriptor instead.
func (*DeleteTrailRequest) Descriptor() ([]byte, []int) {
	return file_cloudtrail_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTrailRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_cloudtrail_proto protoreflect.FileDescriptor

var file_cloudtrail_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0xa8, 0x01, 0x0a, 0x0f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x54, 0x72, 0x61,
	0x69, 0x6c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61,
	0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x72, 0x61,
	0x69, 0x6c, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x22, 0x3c,
	0x0a, 0x10, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x77, 0x72, 0x69, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x61, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x72, 0x65, 0x61, 0x64, 0x22, 0x58, 0x0a, 0x0c,
	0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x34, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x27, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x72, 0x6e, 0x22,
	0x27, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x72, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x72, 0x6e, 0x22, 0x3f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x28, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0xc9, 0x02, 0x0a, 0x11, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x54, 0x72, 0x61,
	0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72,
	0x61, 0x69, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x69,
	0x6c, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x69,
	0x6c, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c,
	0x12, 0x18, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x54, 0x72,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x61, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cloudtrail_proto_rawDescOnce sync.Once
	file_cloudtrail_proto_rawDescData = file_cloudtrail_proto_rawDesc
)

func file_cloudtrail_proto_rawDescGZIP() []byte {
	file_cloudtrail_proto_rawDescOnce.Do(func() {
		file_cloudtrail_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudtrail_proto_rawDescData)
	})
	return file_cloudtrail_proto_rawDescData
}

var file_cloudtrail_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cloudtrail_proto_goTypes = []interface{}{
	(*GetEventsRequest)(nil),      // 0: cloudtrail.GetEventsRequest
	(*CloudTrailEvent)(nil),       // 1: cloudtrail.CloudTrailEvent
	(*CloudTrailConfig)(nil),      // 2: cloudtrail.CloudTrailConfig
	(*TrailRequest)(nil),          // 3: cloudtrail.TrailRequest
	(*CreateTrailResponse)(nil),   // 4: cloudtrail.CreateTrailResponse
	(*UpdateTrailResponse)(nil),   // 5: cloudtrail.UpdateTrailResponse
	(*DeleteTrailResponse)(nil),   // 6: cloudtrail.DeleteTrailResponse
	(*DeleteTrailRequest)(nil),    // 7: cloudtrail.DeleteTrailRequest
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_cloudtrail_proto_depIdxs = []int32{
	8, // 0: cloudtrail.CloudTrailEvent.timestamp:type_name -> google.protobuf.Timestamp
	2, // 1: cloudtrail.TrailRequest.config:type_name -> cloudtrail.CloudTrailConfig
	0, // 2: cloudtrail.CloudTrailService.GetEvents:input_type -> cloudtrail.GetEventsRequest
	3, // 3: cloudtrail.CloudTrailService.CreateTrail:input_type -> cloudtrail.TrailRequest
	3, // 4: cloudtrail.CloudTrailService.UpdateTrail:input_type -> cloudtrail.TrailRequest
	7, // 5: cloudtrail.CloudTrailService.DeleteTrail:input_type -> cloudtrail.DeleteTrailRequest
	1, // 6: cloudtrail.CloudTrailService.GetEvents:output_type -> cloudtrail.CloudTrailEvent
	4, // 7: cloudtrail.CloudTrailService.CreateTrail:output_type -> cloudtrail.CreateTrailResponse
	5, // 8: cloudtrail.CloudTrailService.UpdateTrail:output_type -> cloudtrail.UpdateTrailResponse
	6, // 9: cloudtrail.CloudTrailService.DeleteTrail:output_type -> cloudtrail.DeleteTrailResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloudtrail_proto_init() }
func file_cloudtrail_proto_init() {
	if File_cloudtrail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudtrail_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventsRequest); i {
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
		file_cloudtrail_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudTrailEvent); i {
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
		file_cloudtrail_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloudTrailConfig); i {
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
		file_cloudtrail_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrailRequest); i {
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
		file_cloudtrail_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTrailResponse); i {
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
		file_cloudtrail_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTrailResponse); i {
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
		file_cloudtrail_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTrailResponse); i {
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
		file_cloudtrail_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTrailRequest); i {
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
			RawDescriptor: file_cloudtrail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cloudtrail_proto_goTypes,
		DependencyIndexes: file_cloudtrail_proto_depIdxs,
		MessageInfos:      file_cloudtrail_proto_msgTypes,
	}.Build()
	File_cloudtrail_proto = out.File
	file_cloudtrail_proto_rawDesc = nil
	file_cloudtrail_proto_goTypes = nil
	file_cloudtrail_proto_depIdxs = nil
}

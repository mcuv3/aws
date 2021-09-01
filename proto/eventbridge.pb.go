// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eventbridge.proto

package aws

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Services int32

const (
	Services_Lambda Services = 0
	Services_Sqs    Services = 1
	Services_Iam    Services = 2
)

var Services_name = map[int32]string{
	0: "Lambda",
	1: "Sqs",
	2: "Iam",
}

var Services_value = map[string]int32{
	"Lambda": 0,
	"Sqs":    1,
	"Iam":    2,
}

func (x Services) String() string {
	return proto.EnumName(Services_name, int32(x))
}

func (Services) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_60060df97c40c816, []int{0}
}

type TargetType int32

const (
	TargetType_LambdaFuntion TargetType = 0
	TargetType_SqsQueue      TargetType = 1
)

var TargetType_name = map[int32]string{
	0: "LambdaFuntion",
	1: "SqsQueue",
}

var TargetType_value = map[string]int32{
	"LambdaFuntion": 0,
	"SqsQueue":      1,
}

func (x TargetType) String() string {
	return proto.EnumName(TargetType_name, int32(x))
}

func (TargetType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_60060df97c40c816, []int{1}
}

type Event struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_60060df97c40c816, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EventPattern struct {
	Service              Services `protobuf:"varint,1,opt,name=service,proto3,enum=eventbridge.Services" json:"service,omitempty"`
	EventType            string   `protobuf:"bytes,2,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventPattern) Reset()         { *m = EventPattern{} }
func (m *EventPattern) String() string { return proto.CompactTextString(m) }
func (*EventPattern) ProtoMessage()    {}
func (*EventPattern) Descriptor() ([]byte, []int) {
	return fileDescriptor_60060df97c40c816, []int{1}
}

func (m *EventPattern) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventPattern.Unmarshal(m, b)
}
func (m *EventPattern) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventPattern.Marshal(b, m, deterministic)
}
func (m *EventPattern) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventPattern.Merge(m, src)
}
func (m *EventPattern) XXX_Size() int {
	return xxx_messageInfo_EventPattern.Size(m)
}
func (m *EventPattern) XXX_DiscardUnknown() {
	xxx_messageInfo_EventPattern.DiscardUnknown(m)
}

var xxx_messageInfo_EventPattern proto.InternalMessageInfo

func (m *EventPattern) GetService() Services {
	if m != nil {
		return m.Service
	}
	return Services_Lambda
}

func (m *EventPattern) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

type ScheduleEvent struct {
	Cron                 string   `protobuf:"bytes,1,opt,name=cron,proto3" json:"cron,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScheduleEvent) Reset()         { *m = ScheduleEvent{} }
func (m *ScheduleEvent) String() string { return proto.CompactTextString(m) }
func (*ScheduleEvent) ProtoMessage()    {}
func (*ScheduleEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_60060df97c40c816, []int{2}
}

func (m *ScheduleEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScheduleEvent.Unmarshal(m, b)
}
func (m *ScheduleEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScheduleEvent.Marshal(b, m, deterministic)
}
func (m *ScheduleEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleEvent.Merge(m, src)
}
func (m *ScheduleEvent) XXX_Size() int {
	return xxx_messageInfo_ScheduleEvent.Size(m)
}
func (m *ScheduleEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleEvent proto.InternalMessageInfo

func (m *ScheduleEvent) GetCron() string {
	if m != nil {
		return m.Cron
	}
	return ""
}

type Target struct {
	TargetType           TargetType `protobuf:"varint,1,opt,name=target_type,json=targetType,proto3,enum=eventbridge.TargetType" json:"target_type,omitempty"`
	TargetArn            string     `protobuf:"bytes,2,opt,name=target_arn,json=targetArn,proto3" json:"target_arn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Target) Reset()         { *m = Target{} }
func (m *Target) String() string { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()    {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_60060df97c40c816, []int{3}
}

func (m *Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Target.Unmarshal(m, b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Target.Marshal(b, m, deterministic)
}
func (m *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(m, src)
}
func (m *Target) XXX_Size() int {
	return xxx_messageInfo_Target.Size(m)
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

func (m *Target) GetTargetType() TargetType {
	if m != nil {
		return m.TargetType
	}
	return TargetType_LambdaFuntion
}

func (m *Target) GetTargetArn() string {
	if m != nil {
		return m.TargetArn
	}
	return ""
}

type CreateRuleRequest struct {
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*CreateRuleRequest_EventPattern
	//	*CreateRuleRequest_ScheduleEvent
	Event                isCreateRuleRequest_Event `protobuf_oneof:"event"`
	Targets              []*Target                 `protobuf:"bytes,5,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *CreateRuleRequest) Reset()         { *m = CreateRuleRequest{} }
func (m *CreateRuleRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRuleRequest) ProtoMessage()    {}
func (*CreateRuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60060df97c40c816, []int{4}
}

func (m *CreateRuleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRuleRequest.Unmarshal(m, b)
}
func (m *CreateRuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRuleRequest.Marshal(b, m, deterministic)
}
func (m *CreateRuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRuleRequest.Merge(m, src)
}
func (m *CreateRuleRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRuleRequest.Size(m)
}
func (m *CreateRuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRuleRequest proto.InternalMessageInfo

func (m *CreateRuleRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRuleRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type isCreateRuleRequest_Event interface {
	isCreateRuleRequest_Event()
}

type CreateRuleRequest_EventPattern struct {
	EventPattern *EventPattern `protobuf:"bytes,3,opt,name=event_pattern,json=eventPattern,proto3,oneof"`
}

type CreateRuleRequest_ScheduleEvent struct {
	ScheduleEvent *ScheduleEvent `protobuf:"bytes,4,opt,name=schedule_event,json=scheduleEvent,proto3,oneof"`
}

func (*CreateRuleRequest_EventPattern) isCreateRuleRequest_Event() {}

func (*CreateRuleRequest_ScheduleEvent) isCreateRuleRequest_Event() {}

func (m *CreateRuleRequest) GetEvent() isCreateRuleRequest_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *CreateRuleRequest) GetEventPattern() *EventPattern {
	if x, ok := m.GetEvent().(*CreateRuleRequest_EventPattern); ok {
		return x.EventPattern
	}
	return nil
}

func (m *CreateRuleRequest) GetScheduleEvent() *ScheduleEvent {
	if x, ok := m.GetEvent().(*CreateRuleRequest_ScheduleEvent); ok {
		return x.ScheduleEvent
	}
	return nil
}

func (m *CreateRuleRequest) GetTargets() []*Target {
	if m != nil {
		return m.Targets
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CreateRuleRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CreateRuleRequest_EventPattern)(nil),
		(*CreateRuleRequest_ScheduleEvent)(nil),
	}
}

type CreateRuleResponse struct {
	RuleArn              string   `protobuf:"bytes,1,opt,name=rule_arn,json=ruleArn,proto3" json:"rule_arn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRuleResponse) Reset()         { *m = CreateRuleResponse{} }
func (m *CreateRuleResponse) String() string { return proto.CompactTextString(m) }
func (*CreateRuleResponse) ProtoMessage()    {}
func (*CreateRuleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60060df97c40c816, []int{5}
}

func (m *CreateRuleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRuleResponse.Unmarshal(m, b)
}
func (m *CreateRuleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRuleResponse.Marshal(b, m, deterministic)
}
func (m *CreateRuleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRuleResponse.Merge(m, src)
}
func (m *CreateRuleResponse) XXX_Size() int {
	return xxx_messageInfo_CreateRuleResponse.Size(m)
}
func (m *CreateRuleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRuleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRuleResponse proto.InternalMessageInfo

func (m *CreateRuleResponse) GetRuleArn() string {
	if m != nil {
		return m.RuleArn
	}
	return ""
}

type EventBridgeResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status               string   `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventBridgeResponse) Reset()         { *m = EventBridgeResponse{} }
func (m *EventBridgeResponse) String() string { return proto.CompactTextString(m) }
func (*EventBridgeResponse) ProtoMessage()    {}
func (*EventBridgeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60060df97c40c816, []int{6}
}

func (m *EventBridgeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventBridgeResponse.Unmarshal(m, b)
}
func (m *EventBridgeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventBridgeResponse.Marshal(b, m, deterministic)
}
func (m *EventBridgeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBridgeResponse.Merge(m, src)
}
func (m *EventBridgeResponse) XXX_Size() int {
	return xxx_messageInfo_EventBridgeResponse.Size(m)
}
func (m *EventBridgeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBridgeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventBridgeResponse proto.InternalMessageInfo

func (m *EventBridgeResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *EventBridgeResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type DeleteRuleRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRuleRequest) Reset()         { *m = DeleteRuleRequest{} }
func (m *DeleteRuleRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRuleRequest) ProtoMessage()    {}
func (*DeleteRuleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60060df97c40c816, []int{7}
}

func (m *DeleteRuleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRuleRequest.Unmarshal(m, b)
}
func (m *DeleteRuleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRuleRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRuleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRuleRequest.Merge(m, src)
}
func (m *DeleteRuleRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRuleRequest.Size(m)
}
func (m *DeleteRuleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRuleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRuleRequest proto.InternalMessageInfo

func (m *DeleteRuleRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ChangeRuleStateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeRuleStateRequest) Reset()         { *m = ChangeRuleStateRequest{} }
func (m *ChangeRuleStateRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeRuleStateRequest) ProtoMessage()    {}
func (*ChangeRuleStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60060df97c40c816, []int{8}
}

func (m *ChangeRuleStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeRuleStateRequest.Unmarshal(m, b)
}
func (m *ChangeRuleStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeRuleStateRequest.Marshal(b, m, deterministic)
}
func (m *ChangeRuleStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeRuleStateRequest.Merge(m, src)
}
func (m *ChangeRuleStateRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeRuleStateRequest.Size(m)
}
func (m *ChangeRuleStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeRuleStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeRuleStateRequest proto.InternalMessageInfo

func (m *ChangeRuleStateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterEnum("eventbridge.Services", Services_name, Services_value)
	proto.RegisterEnum("eventbridge.TargetType", TargetType_name, TargetType_value)
	proto.RegisterType((*Event)(nil), "eventbridge.Event")
	proto.RegisterType((*EventPattern)(nil), "eventbridge.EventPattern")
	proto.RegisterType((*ScheduleEvent)(nil), "eventbridge.ScheduleEvent")
	proto.RegisterType((*Target)(nil), "eventbridge.Target")
	proto.RegisterType((*CreateRuleRequest)(nil), "eventbridge.CreateRuleRequest")
	proto.RegisterType((*CreateRuleResponse)(nil), "eventbridge.CreateRuleResponse")
	proto.RegisterType((*EventBridgeResponse)(nil), "eventbridge.EventBridgeResponse")
	proto.RegisterType((*DeleteRuleRequest)(nil), "eventbridge.DeleteRuleRequest")
	proto.RegisterType((*ChangeRuleStateRequest)(nil), "eventbridge.ChangeRuleStateRequest")
}

func init() {
	proto.RegisterFile("eventbridge.proto", fileDescriptor_60060df97c40c816)
}

var fileDescriptor_60060df97c40c816 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x6e, 0xda, 0x30,
	0x14, 0xc6, 0x21, 0x14, 0x42, 0x0f, 0x7f, 0x06, 0xa7, 0x5a, 0x9b, 0x56, 0xda, 0x86, 0xd2, 0x8b,
	0xa1, 0x6a, 0x05, 0x89, 0xdd, 0x4c, 0xda, 0xcd, 0x06, 0xfb, 0xd3, 0x49, 0x93, 0xd6, 0x85, 0xee,
	0xa6, 0x17, 0xab, 0x0c, 0x39, 0xa2, 0x91, 0x20, 0x04, 0xdb, 0xe9, 0xd4, 0x37, 0xd8, 0x8b, 0xee,
	0x3d, 0x26, 0xdb, 0x49, 0x49, 0x4a, 0xab, 0x72, 0x77, 0xec, 0xf3, 0xe5, 0xf3, 0xe7, 0xdf, 0xb1,
	0x02, 0x6d, 0xba, 0xa1, 0x50, 0x4e, 0x78, 0xe0, 0xcf, 0xa8, 0x17, 0xf1, 0xa5, 0x5c, 0x62, 0x2d,
	0xb3, 0xe5, 0x1e, 0x40, 0xf9, 0xb3, 0x5a, 0x62, 0x13, 0xac, 0xc0, 0x77, 0x8a, 0x9d, 0x62, 0x77,
	0xd7, 0xb3, 0x02, 0xdf, 0xfd, 0x0d, 0x75, 0xdd, 0x38, 0x67, 0x52, 0x12, 0x0f, 0xb1, 0x0f, 0xb6,
	0x20, 0x7e, 0x13, 0x4c, 0x49, 0x8b, 0x9a, 0x83, 0xe7, 0xbd, 0xac, 0xf5, 0xd8, 0xf4, 0x84, 0x97,
	0xaa, 0xf0, 0x05, 0x80, 0x16, 0x5c, 0xc9, 0xdb, 0x88, 0x1c, 0x4b, 0x1b, 0xef, 0xea, 0x9d, 0x8b,
	0xdb, 0x88, 0xdc, 0x63, 0x68, 0x8c, 0xa7, 0xd7, 0xe4, 0xc7, 0x73, 0x32, 0x01, 0x10, 0x76, 0xa6,
	0x7c, 0x19, 0x26, 0x11, 0x74, 0xed, 0x32, 0xa8, 0x5c, 0x30, 0x3e, 0x23, 0x89, 0xef, 0xa0, 0x26,
	0x75, 0x65, 0xec, 0x4c, 0x84, 0x83, 0x5c, 0x04, 0xa3, 0x54, 0xe6, 0x1e, 0xc8, 0xbb, 0x5a, 0xe5,
	0x48, 0xbe, 0x64, 0x3c, 0x4c, 0x73, 0x98, 0x9d, 0x8f, 0x3c, 0x74, 0xff, 0x5a, 0xd0, 0x1e, 0x71,
	0x62, 0x92, 0xbc, 0x78, 0x4e, 0x1e, 0xad, 0x62, 0x12, 0x3a, 0x4c, 0xc8, 0x16, 0x94, 0x86, 0x51,
	0x35, 0x76, 0xa0, 0xe6, 0x93, 0x98, 0xf2, 0x20, 0x92, 0xc1, 0x32, 0x75, 0xca, 0x6e, 0xe1, 0x07,
	0x68, 0x98, 0x2b, 0x47, 0x06, 0x9a, 0x53, 0xea, 0x14, 0xbb, 0xb5, 0xc1, 0x61, 0x2e, 0x66, 0x96,
	0xea, 0x59, 0xc1, 0xab, 0x53, 0x96, 0xf2, 0x08, 0x9a, 0x22, 0xa1, 0x72, 0xa5, 0x1b, 0xce, 0x8e,
	0xb6, 0x38, 0xca, 0xc3, 0xce, 0x82, 0x3b, 0x2b, 0x78, 0x0d, 0x91, 0x23, 0x79, 0x0a, 0xb6, 0xb9,
	0x9f, 0x70, 0xca, 0x9d, 0x52, 0xb7, 0x36, 0xd8, 0x7b, 0x80, 0x93, 0x97, 0x6a, 0x86, 0x36, 0x94,
	0x75, 0xdb, 0xed, 0x03, 0x66, 0x49, 0x88, 0x68, 0x19, 0x0a, 0xc2, 0x43, 0xa8, 0x72, 0x15, 0x47,
	0xd1, 0x33, 0x38, 0x6c, 0xb5, 0x56, 0xec, 0xbe, 0xc2, 0x9e, 0x3e, 0x71, 0xa8, 0x8d, 0xef, 0xbe,
	0x70, 0xc0, 0x5e, 0x90, 0x10, 0x6c, 0x96, 0xf2, 0x4b, 0x97, 0xb8, 0x0f, 0x15, 0x21, 0x99, 0x8c,
	0x45, 0x42, 0x2f, 0x59, 0xb9, 0xaf, 0xa1, 0xfd, 0x89, 0xe6, 0xf4, 0xe4, 0x0c, 0xdc, 0x37, 0xb0,
	0x3f, 0xba, 0x66, 0xe1, 0x4c, 0x0b, 0xc7, 0x52, 0x65, 0x7d, 0x5c, 0x7d, 0xd2, 0x85, 0x6a, 0xfa,
	0x2e, 0x11, 0xa0, 0xf2, 0x9d, 0x2d, 0x26, 0x3e, 0x6b, 0x15, 0xd0, 0x86, 0xd2, 0x78, 0x25, 0x5a,
	0x45, 0x55, 0x7c, 0x63, 0x8b, 0x96, 0x75, 0x72, 0x0a, 0xb0, 0x7e, 0x3e, 0xd8, 0x86, 0x86, 0xd1,
	0x7e, 0x89, 0x43, 0x35, 0xd8, 0x56, 0x01, 0xeb, 0x50, 0x1d, 0xaf, 0xc4, 0xcf, 0x98, 0x62, 0x6a,
	0x15, 0x07, 0xff, 0x2c, 0xc0, 0xcc, 0xcd, 0x93, 0x43, 0xf0, 0x07, 0xc0, 0x1a, 0x20, 0xbe, 0xcc,
	0x51, 0xdf, 0x78, 0x63, 0x47, 0xaf, 0x1e, 0xed, 0x1b, 0x8e, 0x6e, 0x01, 0xcf, 0x01, 0x7e, 0x45,
	0xfe, 0xb6, 0x86, 0x9d, 0xcd, 0x77, 0x96, 0x9f, 0x8c, 0x71, 0x5c, 0x93, 0xbe, 0xe7, 0xb8, 0x31,
	0x82, 0xad, 0x1c, 0x2f, 0xe1, 0xd9, 0xbd, 0x91, 0xe0, 0x71, 0x3e, 0xe8, 0x83, 0x03, 0xdb, 0xc6,
	0x7b, 0x58, 0xbd, 0xac, 0xf4, 0xfa, 0xef, 0xd9, 0x1f, 0x31, 0xa9, 0xe8, 0x7f, 0xd7, 0xdb, 0xff,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x28, 0x82, 0x1a, 0xd0, 0x04, 0x00, 0x00,
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package aws

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// EventBridgeServiceClient is the client API for EventBridgeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventBridgeServiceClient interface {
	CreateRule(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*CreateRuleResponse, error)
	UpdateRule(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*EventBridgeResponse, error)
	DeleteRule(ctx context.Context, in *DeleteRuleRequest, opts ...grpc.CallOption) (*EventBridgeResponse, error)
	ChangeRuleState(ctx context.Context, in *ChangeRuleStateRequest, opts ...grpc.CallOption) (*EventBridgeResponse, error)
}

type eventBridgeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventBridgeServiceClient(cc grpc.ClientConnInterface) EventBridgeServiceClient {
	return &eventBridgeServiceClient{cc}
}

func (c *eventBridgeServiceClient) CreateRule(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*CreateRuleResponse, error) {
	out := new(CreateRuleResponse)
	err := c.cc.Invoke(ctx, "/eventbridge.EventBridgeService/CreateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventBridgeServiceClient) UpdateRule(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*EventBridgeResponse, error) {
	out := new(EventBridgeResponse)
	err := c.cc.Invoke(ctx, "/eventbridge.EventBridgeService/UpdateRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventBridgeServiceClient) DeleteRule(ctx context.Context, in *DeleteRuleRequest, opts ...grpc.CallOption) (*EventBridgeResponse, error) {
	out := new(EventBridgeResponse)
	err := c.cc.Invoke(ctx, "/eventbridge.EventBridgeService/DeleteRule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventBridgeServiceClient) ChangeRuleState(ctx context.Context, in *ChangeRuleStateRequest, opts ...grpc.CallOption) (*EventBridgeResponse, error) {
	out := new(EventBridgeResponse)
	err := c.cc.Invoke(ctx, "/eventbridge.EventBridgeService/ChangeRuleState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventBridgeServiceServer is the server API for EventBridgeService service.
// All implementations should embed UnimplementedEventBridgeServiceServer
// for forward compatibility
type EventBridgeServiceServer interface {
	CreateRule(context.Context, *CreateRuleRequest) (*CreateRuleResponse, error)
	UpdateRule(context.Context, *CreateRuleRequest) (*EventBridgeResponse, error)
	DeleteRule(context.Context, *DeleteRuleRequest) (*EventBridgeResponse, error)
	ChangeRuleState(context.Context, *ChangeRuleStateRequest) (*EventBridgeResponse, error)
}

// UnimplementedEventBridgeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEventBridgeServiceServer struct {
}

func (UnimplementedEventBridgeServiceServer) CreateRule(context.Context, *CreateRuleRequest) (*CreateRuleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRule not implemented")
}
func (UnimplementedEventBridgeServiceServer) UpdateRule(context.Context, *CreateRuleRequest) (*EventBridgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRule not implemented")
}
func (UnimplementedEventBridgeServiceServer) DeleteRule(context.Context, *DeleteRuleRequest) (*EventBridgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRule not implemented")
}
func (UnimplementedEventBridgeServiceServer) ChangeRuleState(context.Context, *ChangeRuleStateRequest) (*EventBridgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeRuleState not implemented")
}

// UnsafeEventBridgeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventBridgeServiceServer will
// result in compilation errors.
type UnsafeEventBridgeServiceServer interface {
	mustEmbedUnimplementedEventBridgeServiceServer()
}

func RegisterEventBridgeServiceServer(s *grpc.Server, srv EventBridgeServiceServer) {
	s.RegisterService(&_EventBridgeService_serviceDesc, srv)
}

func _EventBridgeService_CreateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventBridgeServiceServer).CreateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventbridge.EventBridgeService/CreateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventBridgeServiceServer).CreateRule(ctx, req.(*CreateRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventBridgeService_UpdateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventBridgeServiceServer).UpdateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventbridge.EventBridgeService/UpdateRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventBridgeServiceServer).UpdateRule(ctx, req.(*CreateRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventBridgeService_DeleteRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventBridgeServiceServer).DeleteRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventbridge.EventBridgeService/DeleteRule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventBridgeServiceServer).DeleteRule(ctx, req.(*DeleteRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventBridgeService_ChangeRuleState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeRuleStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventBridgeServiceServer).ChangeRuleState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventbridge.EventBridgeService/ChangeRuleState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventBridgeServiceServer).ChangeRuleState(ctx, req.(*ChangeRuleStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventBridgeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventbridge.EventBridgeService",
	HandlerType: (*EventBridgeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRule",
			Handler:    _EventBridgeService_CreateRule_Handler,
		},
		{
			MethodName: "UpdateRule",
			Handler:    _EventBridgeService_UpdateRule_Handler,
		},
		{
			MethodName: "DeleteRule",
			Handler:    _EventBridgeService_DeleteRule_Handler,
		},
		{
			MethodName: "ChangeRuleState",
			Handler:    _EventBridgeService_ChangeRuleState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eventbridge.proto",
}
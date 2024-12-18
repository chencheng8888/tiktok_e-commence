// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.26.1
// source: auth/v1/auther.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuthService_DeliverTokenByRPC_FullMethodName = "/auth.v1.AuthService/DeliverTokenByRPC"
	AuthService_VerifyTokenByRPC_FullMethodName  = "/auth.v1.AuthService/VerifyTokenByRPC"
	AuthService_AssignRole_FullMethodName        = "/auth.v1.AuthService/AssignRole"
	AuthService_RemoveRole_FullMethodName        = "/auth.v1.AuthService/RemoveRole"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	// 分发令牌
	DeliverTokenByRPC(ctx context.Context, in *DeliverTokenReq, opts ...grpc.CallOption) (*DeliveryResp, error)
	// 鉴权(判断请求能否通过)
	VerifyTokenByRPC(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyResp, error)
	AssignRole(ctx context.Context, in *AssignRoleReq, opts ...grpc.CallOption) (*AssignResp, error)
	RemoveRole(ctx context.Context, in *RemoveRoleReq, opts ...grpc.CallOption) (*RemoveResp, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) DeliverTokenByRPC(ctx context.Context, in *DeliverTokenReq, opts ...grpc.CallOption) (*DeliveryResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeliveryResp)
	err := c.cc.Invoke(ctx, AuthService_DeliverTokenByRPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) VerifyTokenByRPC(ctx context.Context, in *VerifyTokenReq, opts ...grpc.CallOption) (*VerifyResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyResp)
	err := c.cc.Invoke(ctx, AuthService_VerifyTokenByRPC_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) AssignRole(ctx context.Context, in *AssignRoleReq, opts ...grpc.CallOption) (*AssignResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AssignResp)
	err := c.cc.Invoke(ctx, AuthService_AssignRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RemoveRole(ctx context.Context, in *RemoveRoleReq, opts ...grpc.CallOption) (*RemoveResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveResp)
	err := c.cc.Invoke(ctx, AuthService_RemoveRole_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility.
type AuthServiceServer interface {
	// 分发令牌
	DeliverTokenByRPC(context.Context, *DeliverTokenReq) (*DeliveryResp, error)
	// 鉴权(判断请求能否通过)
	VerifyTokenByRPC(context.Context, *VerifyTokenReq) (*VerifyResp, error)
	AssignRole(context.Context, *AssignRoleReq) (*AssignResp, error)
	RemoveRole(context.Context, *RemoveRoleReq) (*RemoveResp, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) DeliverTokenByRPC(context.Context, *DeliverTokenReq) (*DeliveryResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeliverTokenByRPC not implemented")
}
func (UnimplementedAuthServiceServer) VerifyTokenByRPC(context.Context, *VerifyTokenReq) (*VerifyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyTokenByRPC not implemented")
}
func (UnimplementedAuthServiceServer) AssignRole(context.Context, *AssignRoleReq) (*AssignResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignRole not implemented")
}
func (UnimplementedAuthServiceServer) RemoveRole(context.Context, *RemoveRoleReq) (*RemoveResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRole not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}
func (UnimplementedAuthServiceServer) testEmbeddedByValue()                     {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_DeliverTokenByRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeliverTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DeliverTokenByRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_DeliverTokenByRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DeliverTokenByRPC(ctx, req.(*DeliverTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_VerifyTokenByRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).VerifyTokenByRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_VerifyTokenByRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).VerifyTokenByRPC(ctx, req.(*VerifyTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_AssignRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).AssignRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_AssignRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).AssignRole(ctx, req.(*AssignRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RemoveRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RemoveRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_RemoveRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RemoveRole(ctx, req.(*RemoveRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.v1.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeliverTokenByRPC",
			Handler:    _AuthService_DeliverTokenByRPC_Handler,
		},
		{
			MethodName: "VerifyTokenByRPC",
			Handler:    _AuthService_VerifyTokenByRPC_Handler,
		},
		{
			MethodName: "AssignRole",
			Handler:    _AuthService_AssignRole_Handler,
		},
		{
			MethodName: "RemoveRole",
			Handler:    _AuthService_RemoveRole_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/v1/auther.proto",
}

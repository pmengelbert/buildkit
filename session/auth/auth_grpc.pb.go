// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.11.4
// source: auth.proto

package auth

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
	Auth_Credentials_FullMethodName          = "/moby.filesync.v1.Auth/Credentials"
	Auth_FetchToken_FullMethodName           = "/moby.filesync.v1.Auth/FetchToken"
	Auth_GetTokenAuthority_FullMethodName    = "/moby.filesync.v1.Auth/GetTokenAuthority"
	Auth_VerifyTokenAuthority_FullMethodName = "/moby.filesync.v1.Auth/VerifyTokenAuthority"
)

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthClient interface {
	Credentials(ctx context.Context, in *CredentialsRequest, opts ...grpc.CallOption) (*CredentialsResponse, error)
	FetchToken(ctx context.Context, in *FetchTokenRequest, opts ...grpc.CallOption) (*FetchTokenResponse, error)
	GetTokenAuthority(ctx context.Context, in *GetTokenAuthorityRequest, opts ...grpc.CallOption) (*GetTokenAuthorityResponse, error)
	VerifyTokenAuthority(ctx context.Context, in *VerifyTokenAuthorityRequest, opts ...grpc.CallOption) (*VerifyTokenAuthorityResponse, error)
}

type authClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Credentials(ctx context.Context, in *CredentialsRequest, opts ...grpc.CallOption) (*CredentialsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CredentialsResponse)
	err := c.cc.Invoke(ctx, Auth_Credentials_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) FetchToken(ctx context.Context, in *FetchTokenRequest, opts ...grpc.CallOption) (*FetchTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FetchTokenResponse)
	err := c.cc.Invoke(ctx, Auth_FetchToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetTokenAuthority(ctx context.Context, in *GetTokenAuthorityRequest, opts ...grpc.CallOption) (*GetTokenAuthorityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTokenAuthorityResponse)
	err := c.cc.Invoke(ctx, Auth_GetTokenAuthority_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) VerifyTokenAuthority(ctx context.Context, in *VerifyTokenAuthorityRequest, opts ...grpc.CallOption) (*VerifyTokenAuthorityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VerifyTokenAuthorityResponse)
	err := c.cc.Invoke(ctx, Auth_VerifyTokenAuthority_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
// All implementations should embed UnimplementedAuthServer
// for forward compatibility.
type AuthServer interface {
	Credentials(context.Context, *CredentialsRequest) (*CredentialsResponse, error)
	FetchToken(context.Context, *FetchTokenRequest) (*FetchTokenResponse, error)
	GetTokenAuthority(context.Context, *GetTokenAuthorityRequest) (*GetTokenAuthorityResponse, error)
	VerifyTokenAuthority(context.Context, *VerifyTokenAuthorityRequest) (*VerifyTokenAuthorityResponse, error)
}

// UnimplementedAuthServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServer struct{}

func (UnimplementedAuthServer) Credentials(context.Context, *CredentialsRequest) (*CredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Credentials not implemented")
}
func (UnimplementedAuthServer) FetchToken(context.Context, *FetchTokenRequest) (*FetchTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchToken not implemented")
}
func (UnimplementedAuthServer) GetTokenAuthority(context.Context, *GetTokenAuthorityRequest) (*GetTokenAuthorityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenAuthority not implemented")
}
func (UnimplementedAuthServer) VerifyTokenAuthority(context.Context, *VerifyTokenAuthorityRequest) (*VerifyTokenAuthorityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyTokenAuthority not implemented")
}
func (UnimplementedAuthServer) testEmbeddedByValue() {}

// UnsafeAuthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServer will
// result in compilation errors.
type UnsafeAuthServer interface {
	mustEmbedUnimplementedAuthServer()
}

func RegisterAuthServer(s grpc.ServiceRegistrar, srv AuthServer) {
	// If the following call pancis, it indicates UnimplementedAuthServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Auth_ServiceDesc, srv)
}

func _Auth_Credentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Credentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_Credentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Credentials(ctx, req.(*CredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_FetchToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).FetchToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_FetchToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).FetchToken(ctx, req.(*FetchTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetTokenAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetTokenAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_GetTokenAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetTokenAuthority(ctx, req.(*GetTokenAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_VerifyTokenAuthority_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenAuthorityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).VerifyTokenAuthority(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Auth_VerifyTokenAuthority_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).VerifyTokenAuthority(ctx, req.(*VerifyTokenAuthorityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Auth_ServiceDesc is the grpc.ServiceDesc for Auth service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auth_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moby.filesync.v1.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Credentials",
			Handler:    _Auth_Credentials_Handler,
		},
		{
			MethodName: "FetchToken",
			Handler:    _Auth_FetchToken_Handler,
		},
		{
			MethodName: "GetTokenAuthority",
			Handler:    _Auth_GetTokenAuthority_Handler,
		},
		{
			MethodName: "VerifyTokenAuthority",
			Handler:    _Auth_VerifyTokenAuthority_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

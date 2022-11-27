// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1/authentication.proto

package pkg

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationClient interface {
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	GetSessions(ctx context.Context, in *GetSessionsRequest, opts ...grpc.CallOption) (*GetSessionsResponse, error)
	RevokeSession(ctx context.Context, in *RevokeSessionRequest, opts ...grpc.CallOption) (*RevokeSessionResponse, error)
	ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (Authentication_RefreshTokenClient, error)
	Enable2FA(ctx context.Context, in *Enable2FARequest, opts ...grpc.CallOption) (*Enable2FAResponse, error)
	Disable2FA(ctx context.Context, in *Disable2FARequest, opts ...grpc.CallOption) (*Disable2FAResponse, error)
	VerifyAccount(ctx context.Context, in *VerifyAccountRequest, opts ...grpc.CallOption) (*VerifyAccountResponse, error)
	DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error)
}

type authenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationClient(cc grpc.ClientConnInterface) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.Authentication/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.Authentication/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) GetSessions(ctx context.Context, in *GetSessionsRequest, opts ...grpc.CallOption) (*GetSessionsResponse, error) {
	out := new(GetSessionsResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.Authentication/GetSessions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) RevokeSession(ctx context.Context, in *RevokeSessionRequest, opts ...grpc.CallOption) (*RevokeSessionResponse, error) {
	out := new(RevokeSessionResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.Authentication/RevokeSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest, opts ...grpc.CallOption) (*ResetPasswordResponse, error) {
	out := new(ResetPasswordResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.Authentication/ResetPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.Authentication/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (Authentication_RefreshTokenClient, error) {
	stream, err := c.cc.NewStream(ctx, &Authentication_ServiceDesc.Streams[0], "/authentication.v1.Authentication/RefreshToken", opts...)
	if err != nil {
		return nil, err
	}
	x := &authenticationRefreshTokenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Authentication_RefreshTokenClient interface {
	Recv() (*RefreshTokenResponse, error)
	grpc.ClientStream
}

type authenticationRefreshTokenClient struct {
	grpc.ClientStream
}

func (x *authenticationRefreshTokenClient) Recv() (*RefreshTokenResponse, error) {
	m := new(RefreshTokenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *authenticationClient) Enable2FA(ctx context.Context, in *Enable2FARequest, opts ...grpc.CallOption) (*Enable2FAResponse, error) {
	out := new(Enable2FAResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.Authentication/Enable2FA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Disable2FA(ctx context.Context, in *Disable2FARequest, opts ...grpc.CallOption) (*Disable2FAResponse, error) {
	out := new(Disable2FAResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.Authentication/Disable2FA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) VerifyAccount(ctx context.Context, in *VerifyAccountRequest, opts ...grpc.CallOption) (*VerifyAccountResponse, error) {
	out := new(VerifyAccountResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.Authentication/VerifyAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) DeleteAccount(ctx context.Context, in *DeleteAccountRequest, opts ...grpc.CallOption) (*DeleteAccountResponse, error) {
	out := new(DeleteAccountResponse)
	err := c.cc.Invoke(ctx, "/authentication.v1.Authentication/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
// All implementations must embed UnimplementedAuthenticationServer
// for forward compatibility
type AuthenticationServer interface {
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	GetSessions(context.Context, *GetSessionsRequest) (*GetSessionsResponse, error)
	RevokeSession(context.Context, *RevokeSessionRequest) (*RevokeSessionResponse, error)
	ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	RefreshToken(*RefreshTokenRequest, Authentication_RefreshTokenServer) error
	Enable2FA(context.Context, *Enable2FARequest) (*Enable2FAResponse, error)
	Disable2FA(context.Context, *Disable2FARequest) (*Disable2FAResponse, error)
	VerifyAccount(context.Context, *VerifyAccountRequest) (*VerifyAccountResponse, error)
	DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error)
	mustEmbedUnimplementedAuthenticationServer()
}

// UnimplementedAuthenticationServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServer struct {
}

func (UnimplementedAuthenticationServer) SignIn(context.Context, *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedAuthenticationServer) SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedAuthenticationServer) GetSessions(context.Context, *GetSessionsRequest) (*GetSessionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSessions not implemented")
}
func (UnimplementedAuthenticationServer) RevokeSession(context.Context, *RevokeSessionRequest) (*RevokeSessionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeSession not implemented")
}
func (UnimplementedAuthenticationServer) ResetPassword(context.Context, *ResetPasswordRequest) (*ResetPasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedAuthenticationServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedAuthenticationServer) RefreshToken(*RefreshTokenRequest, Authentication_RefreshTokenServer) error {
	return status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (UnimplementedAuthenticationServer) Enable2FA(context.Context, *Enable2FARequest) (*Enable2FAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable2FA not implemented")
}
func (UnimplementedAuthenticationServer) Disable2FA(context.Context, *Disable2FARequest) (*Disable2FAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable2FA not implemented")
}
func (UnimplementedAuthenticationServer) VerifyAccount(context.Context, *VerifyAccountRequest) (*VerifyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAccount not implemented")
}
func (UnimplementedAuthenticationServer) DeleteAccount(context.Context, *DeleteAccountRequest) (*DeleteAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAuthenticationServer) mustEmbedUnimplementedAuthenticationServer() {}

// UnsafeAuthenticationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServer will
// result in compilation errors.
type UnsafeAuthenticationServer interface {
	mustEmbedUnimplementedAuthenticationServer()
}

func RegisterAuthenticationServer(s grpc.ServiceRegistrar, srv AuthenticationServer) {
	s.RegisterService(&Authentication_ServiceDesc, srv)
}

func _Authentication_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.Authentication/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.Authentication/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_GetSessions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).GetSessions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.Authentication/GetSessions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).GetSessions(ctx, req.(*GetSessionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_RevokeSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).RevokeSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.Authentication/RevokeSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).RevokeSession(ctx, req.(*RevokeSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.Authentication/ResetPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).ResetPassword(ctx, req.(*ResetPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.Authentication/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_RefreshToken_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RefreshTokenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AuthenticationServer).RefreshToken(m, &authenticationRefreshTokenServer{stream})
}

type Authentication_RefreshTokenServer interface {
	Send(*RefreshTokenResponse) error
	grpc.ServerStream
}

type authenticationRefreshTokenServer struct {
	grpc.ServerStream
}

func (x *authenticationRefreshTokenServer) Send(m *RefreshTokenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Authentication_Enable2FA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Enable2FARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Enable2FA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.Authentication/Enable2FA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Enable2FA(ctx, req.(*Enable2FARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Disable2FA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Disable2FARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Disable2FA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.Authentication/Disable2FA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Disable2FA(ctx, req.(*Disable2FARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_VerifyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).VerifyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.Authentication/VerifyAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).VerifyAccount(ctx, req.(*VerifyAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentication.v1.Authentication/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).DeleteAccount(ctx, req.(*DeleteAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authentication_ServiceDesc is the grpc.ServiceDesc for Authentication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authentication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.v1.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _Authentication_SignIn_Handler,
		},
		{
			MethodName: "SignUp",
			Handler:    _Authentication_SignUp_Handler,
		},
		{
			MethodName: "GetSessions",
			Handler:    _Authentication_GetSessions_Handler,
		},
		{
			MethodName: "RevokeSession",
			Handler:    _Authentication_RevokeSession_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _Authentication_ResetPassword_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _Authentication_ChangePassword_Handler,
		},
		{
			MethodName: "Enable2FA",
			Handler:    _Authentication_Enable2FA_Handler,
		},
		{
			MethodName: "Disable2FA",
			Handler:    _Authentication_Disable2FA_Handler,
		},
		{
			MethodName: "VerifyAccount",
			Handler:    _Authentication_VerifyAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _Authentication_DeleteAccount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RefreshToken",
			Handler:       _Authentication_RefreshToken_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1/authentication.proto",
}

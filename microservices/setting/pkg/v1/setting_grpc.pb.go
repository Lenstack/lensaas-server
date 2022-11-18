// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1/setting.proto

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

// SettingClient is the client API for Setting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SettingClient interface {
	CreateSetting(ctx context.Context, in *CreateSettingRequest, opts ...grpc.CallOption) (*CreateSettingResponse, error)
}

type settingClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingClient(cc grpc.ClientConnInterface) SettingClient {
	return &settingClient{cc}
}

func (c *settingClient) CreateSetting(ctx context.Context, in *CreateSettingRequest, opts ...grpc.CallOption) (*CreateSettingResponse, error) {
	out := new(CreateSettingResponse)
	err := c.cc.Invoke(ctx, "/setting.v1.Setting/CreateSetting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingServer is the server API for Setting service.
// All implementations must embed UnimplementedSettingServer
// for forward compatibility
type SettingServer interface {
	CreateSetting(context.Context, *CreateSettingRequest) (*CreateSettingResponse, error)
	mustEmbedUnimplementedSettingServer()
}

// UnimplementedSettingServer must be embedded to have forward compatible implementations.
type UnimplementedSettingServer struct {
}

func (UnimplementedSettingServer) CreateSetting(context.Context, *CreateSettingRequest) (*CreateSettingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSetting not implemented")
}
func (UnimplementedSettingServer) mustEmbedUnimplementedSettingServer() {}

// UnsafeSettingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingServer will
// result in compilation errors.
type UnsafeSettingServer interface {
	mustEmbedUnimplementedSettingServer()
}

func RegisterSettingServer(s grpc.ServiceRegistrar, srv SettingServer) {
	s.RegisterService(&Setting_ServiceDesc, srv)
}

func _Setting_CreateSetting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSettingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingServer).CreateSetting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/setting.v1.Setting/CreateSetting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingServer).CreateSetting(ctx, req.(*CreateSettingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Setting_ServiceDesc is the grpc.ServiceDesc for Setting service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Setting_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "setting.v1.Setting",
	HandlerType: (*SettingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSetting",
			Handler:    _Setting_CreateSetting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/setting.proto",
}
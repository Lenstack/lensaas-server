package applications

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type IMiddlewareApplication interface {
	GrpcUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error)
	GrpcStreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error
}

type MiddlewareApplication struct {
	loggerManager *zap.Logger
}

func NewMiddlewareApplication(loggerManager *zap.Logger) *MiddlewareApplication {
	return &MiddlewareApplication{loggerManager: loggerManager}
}

func (m *MiddlewareApplication) GrpcUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	m.loggerManager.Sugar().Infof("gRPC unary interceptor: %s", info.FullMethod)
	return handler(ctx, req)
}

func (m *MiddlewareApplication) GrpcStreamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	m.loggerManager.Sugar().Infof("gRPC stream interceptor: %s", info.FullMethod)
	return handler(srv, stream)
}

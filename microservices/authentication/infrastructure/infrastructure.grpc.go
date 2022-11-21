package infrastructure

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/core/applications"
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/pkg/v1"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type IGrpcServer interface {
}

type GrpcServer struct {
	port string
}

func NewGrpcServer(port string, microservices applications.MicroserviceServer, loggerManager *zap.Logger) *GrpcServer {
	listenServer, err := net.Listen("tcp", ":"+port)
	if err != nil {
		loggerManager.Sugar().Fatalf("failed to listen: %v", err)
		return nil
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(loggerManager),
			grpc_recovery.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(loggerManager),
			grpc_recovery.StreamServerInterceptor(),
		)),
	)
	pkg.RegisterAuthenticationServer(grpcServer, &microservices)
	reflection.Register(grpcServer)

	loggerManager.Sugar().Infof("gRPC server listening on port %s", port)
	err = grpcServer.Serve(listenServer)
	if err != nil {
		loggerManager.Sugar().Errorf("failed to serve: %v", err)
		return nil
	}
	return &GrpcServer{}
}

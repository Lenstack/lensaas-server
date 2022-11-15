package infrastructure

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/core/applications"
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/pkg"
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
		grpc.UnaryInterceptor(microservices.MiddlewareApplication.GrpcUnaryInterceptor),
		grpc.StreamInterceptor(microservices.MiddlewareApplication.GrpcStreamInterceptor),
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

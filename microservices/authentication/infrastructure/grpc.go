package infrastructure

import (
	"github.com/Lenstack/Lensaas/microservices/authentication/core/applications"
	"github.com/Lenstack/Lensaas/microservices/authentication/pkg/v1"
	"github.com/Lenstack/Lensaas/microservices/authentication/util"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type GrpcServer struct {
	port string
}

func NewGrpcServer(port string, microservices applications.MicroserviceServer, loggerManager *zap.Logger) *GrpcServer {
	listenServer, err := net.Listen("tcp", ":"+port)
	if err != nil {
		loggerManager.Sugar().Fatalf("failed to listen: %v", err)
		return nil
	}

	tlsCredentials, err := util.LoadTLSCredentials()
	if err != nil {
		loggerManager.Sugar().Fatalf("failed to load tls credentials: %v", err)
		return nil
	}

	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	pkg.RegisterAuthenticationServer(grpcServer, &microservices)
	reflection.Register(grpcServer)

	loggerManager.Sugar().Infof("GRPC server listening on port %s", port)
	err = grpcServer.Serve(listenServer)
	if err != nil {
		loggerManager.Sugar().Errorf("failed to serve: %v", err)
		return nil
	}
	return &GrpcServer{}
}

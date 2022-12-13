package infrastructure

// openssl req -nodes -x509 -newkey rsa:4096 -keyout server-key.pem -out server-cert.pem -sha256 -days 365 -subj '/CN=localhost'
import (
	"crypto/tls"
	"github.com/Lenstack/Lensaas/microservices/authentication/core/applications"
	"github.com/Lenstack/Lensaas/microservices/authentication/pkg/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	cert, err := tls.LoadX509KeyPair("./cert/server-cert.pem", "./cert/server-key.pem")
	if err != nil {
		loggerManager.Sugar().Fatalf("failed to load key pair: %v", err)
		return nil
	}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}

	grpcServer := grpc.NewServer(opts...)
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

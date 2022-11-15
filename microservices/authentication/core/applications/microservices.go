package applications

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/core/services"
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/pkg"
)

type MicroserviceServer struct {
	pkg.UnimplementedAuthenticationServer
	MiddlewareApplication MiddlewareApplication
	AuthenticationService services.AuthenticationService
}

func NewMicroserviceServer(
	middlewareApplication MiddlewareApplication,
	authenticationService services.AuthenticationService,
) *MicroserviceServer {
	return &MicroserviceServer{
		MiddlewareApplication: middlewareApplication,
		AuthenticationService: authenticationService,
	}
}

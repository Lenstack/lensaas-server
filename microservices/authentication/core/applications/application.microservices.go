package applications

import (
	"github.com/Lenstack/Lensaas/microservices/authentication/core/services"
	"github.com/Lenstack/Lensaas/microservices/authentication/pkg/v1"
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

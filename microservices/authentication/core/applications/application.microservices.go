package applications

import (
	"github.com/Lenstack/Lensaas/microservices/authentication/core/services"
	"github.com/Lenstack/Lensaas/microservices/authentication/pkg/v1"
)

type MicroserviceServer struct {
	pkg.UnimplementedAuthenticationServer
	AuthenticationService services.AuthenticationService
}

func NewMicroserviceServer(
	authenticationService services.AuthenticationService,
) *MicroserviceServer {
	return &MicroserviceServer{
		AuthenticationService: authenticationService,
	}
}

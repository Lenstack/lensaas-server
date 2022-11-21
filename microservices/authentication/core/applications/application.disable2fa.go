package applications

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/pkg/v1"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) Disable2FA(ctx context.Context, req *pkg.Disable2FARequest) (*pkg.Disable2FAResponse, error) {
	return &pkg.Disable2FAResponse{Message: "2fa has been disabled"}, nil
}

package applications

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/pkg/v1"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) VerifyEmail(ctx context.Context, req *pkg.VerifyEmailRequest) (*pkg.VerifyEmailResponse, error) {
	return &pkg.VerifyEmailResponse{Message: "the account has been activated"}, nil
}

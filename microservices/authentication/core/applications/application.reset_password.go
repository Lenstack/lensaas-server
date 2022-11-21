package applications

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/pkg/v1"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) ResetPassword(ctx context.Context, req *pkg.ResetPasswordRequest) (*pkg.ResetPasswordResponse, error) {
	return &pkg.ResetPasswordResponse{Message: "the password has been updated"}, nil
}

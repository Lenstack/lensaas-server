package applications

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/pkg/v1"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) ChangePassword(ctx context.Context, req *pkg.ChangePasswordRequest) (*pkg.ChangePasswordResponse, error) {
	return &pkg.ChangePasswordResponse{Message: "the password has been changed"}, nil
}

package applications

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/pkg/v1"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) VerifyAccount(ctx context.Context, req *pkg.VerifyAccountRequest) (*pkg.VerifyAccountResponse, error) {
	return &pkg.VerifyAccountResponse{Message: "the account has been verified"}, nil
}

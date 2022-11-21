package applications

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/pkg/v1"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) DeleteAccount(ctx context.Context, req *pkg.DeleteAccountRequest) (*pkg.DeleteAccountResponse, error) {
	return &pkg.DeleteAccountResponse{Message: "the account has been deleted"}, nil
}

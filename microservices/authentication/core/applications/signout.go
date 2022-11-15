package applications

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/pkg"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) SignOut(ctx context.Context, req *pkg.SignOutRequest) (*pkg.SignOutResponse, error) {
	message, err := ms.AuthenticationService.SignOut(req.GetToken())
	if err != nil {
		return nil, err
	}
	return &pkg.SignOutResponse{Message: message}, nil
}

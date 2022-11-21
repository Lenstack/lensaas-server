package applications

import (
	"fmt"
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/pkg/v1"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) SignUp(ctx context.Context, req *pkg.SignUpRequest) (*pkg.SignUpResponse, error) {
	userId, err := ms.AuthenticationService.SignUp(req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return &pkg.SignUpResponse{Message: fmt.Sprintf("the user has been created successfuly: %s", userId)}, nil
}

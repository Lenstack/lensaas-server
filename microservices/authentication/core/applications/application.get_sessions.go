package applications

import (
	"github.com/Lenstack/clean-grpc-microservices-gateway-ui/tree/master/microservices/authentication/pkg/v1"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) GetSessions(ctx context.Context, req *pkg.GetSessionsRequest) (*pkg.GetSessionsResponse, error) {
	return &pkg.GetSessionsResponse{Sessions: []string{""}}, nil
}

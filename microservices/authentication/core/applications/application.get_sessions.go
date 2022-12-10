package applications

import (
	"github.com/Lenstack/Lensaas/microservices/authentication/pkg/v1"
	"golang.org/x/net/context"
)

func (ms *MicroserviceServer) GetSessions(ctx context.Context, req *pkg.GetSessionsRequest) (*pkg.GetSessionsResponse, error) {
	return &pkg.GetSessionsResponse{Sessions: []string{""}}, nil
}

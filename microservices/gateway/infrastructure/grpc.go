package infrastructure

import (
	"go.uber.org/zap"
)

type GrpcClient struct {
	target string
	logger *zap.Logger
}

func NewGrpcClient(target string, logger *zap.Logger) *GrpcClient {
	return &GrpcClient{}
}

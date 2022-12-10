package infrastructure

import (
	"go.uber.org/zap"
	"strings"
)

type LoggerManager struct {
	Logger *zap.Logger
}

func NewLoggerManager(environment string) *LoggerManager {
	if strings.ToLower(environment) == "production" {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		return &LoggerManager{Logger: logger}
	}

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	return &LoggerManager{Logger: logger}
}

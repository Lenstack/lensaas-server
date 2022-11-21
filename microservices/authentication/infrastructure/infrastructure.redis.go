package infrastructure

import (
	"github.com/go-redis/redis/v9"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

type IRedisManager interface {
}

type RedisManager struct {
	Client *redis.Client
}

func NewRedisManager(host string, port string, password string, loggerManager *zap.Logger) *RedisManager {
	ctx := context.Background()

	loggerManager.Sugar().Infof("Redis server listening on port %s", port)
	rdb := redis.NewClient(&redis.Options{Addr: host + ":" + port, Password: password, DB: 0})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		loggerManager.Sugar().Fatalf("failed to connect to Redis: %v", err)
	}
	return &RedisManager{Client: rdb}
}

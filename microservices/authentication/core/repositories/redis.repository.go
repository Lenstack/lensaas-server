package repositories

import (
	"github.com/go-redis/redis/v9"
	"golang.org/x/net/context"
	"time"
)

type IRedisRepository interface {
	SetStringValue(key string, value interface{}, expiration time.Duration) error
	SetHashValue(key string, field string, value interface{}, expiration time.Duration) error
	GetStringValue(key string) (interface{}, error)
	GetHashValue(key string, field string) (interface{}, error)
}

type RedisRepository struct {
	Client *redis.Client
}

func (rr *RedisRepository) SetStringValue(key string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	return rr.Client.Set(ctx, key, value, expiration).Err()
}

func (rr *RedisRepository) GetStringValue(key string) (interface{}, error) {
	ctx := context.Background()
	return rr.Client.Get(ctx, key).Result()
}

func (rr *RedisRepository) SetHashValue(key string, field string, value interface{}, expiration time.Duration) error {
	ctx := context.Background()
	err := rr.Client.HSet(ctx, key, field, value).Err()
	if err != nil {
		return err
	}
	rr.Client.Expire(ctx, key, expiration)
	return nil
}

func (rr *RedisRepository) GetHashValue(key string, field string) (interface{}, error) {
	ctx := context.Background()
	return rr.Client.HGet(ctx, key, field).Result()
}

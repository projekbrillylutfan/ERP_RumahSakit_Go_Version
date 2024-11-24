package impl_service

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	Client *redis.Client
}

func NewRedisService(client *redis.Client) *RedisService {
	return &RedisService{Client: client}
}

func (r *RedisService) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := r.Client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisService) Get(ctx context.Context, key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("key not found")
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func (r *RedisService) Del(ctx context.Context, key string) error {
	err := r.Client.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
package types

import (
	"context"
	redis "github.com/redis/go-redis/v9"
	"main/_lib/logger"
	"time"
)

type RedisClient struct {
	Client *redis.Client
	Ctx    context.Context
}

func (rc RedisClient) Get(key string) string {
	value, err := rc.Client.Get(rc.Ctx, key).Result()
	if err != nil {
		logger.Info("ERROR! Get redis value error.")
	}
	return value
}

func (rc RedisClient) SetEx(key string, value string, expireSec uint) {
	err := rc.Client.Set(rc.Ctx, key, value, time.Duration(expireSec*1e9)).Err()
	if err != nil {
		logger.Info("ERROR! Set redis value error.")
	}
}

func (rc RedisClient) Set(key string, value string) {
	rc.SetEx(key, value, 0)
}

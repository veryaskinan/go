package redis

import (
	"context"
	redis "github.com/redis/go-redis/v9"
	"main/_lib/env"
	"main/_lib/redis/types"
)

func NewClient() types.RedisClient {
	redisAddr := env.Get("REDIS_HOST") + ":" + env.Get("REDIS_PORT")
	redisPass := env.Get("REDIS_PASSWORD")
	redisDb := env.GetInt("REDIS_BD")
	return types.RedisClient{
		Client: redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: redisPass,
			DB:       redisDb,
		}),
		Ctx: context.Background(),
	}
}

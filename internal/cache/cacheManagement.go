// Package cache gerencia o Redis
package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func ConectarRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis_encurtador:6379",
		Password: "",
		DB:       0,
	})
	err := RedisClient.Ping(Ctx).Err()
	if err != nil {
		panic(fmt.Sprintf("Não foi possível conectar ao Redis: %v", err))
	}
	fmt.Println("Redis conectado.")
}

func AdicionarLinkRedis(codigo, url string) error {
	return RedisClient.Set(Ctx, codigo, url, 0).Err()
}

func BuscarLinkRedis(codigo string) (string, error) {
	return RedisClient.Get(Ctx, codigo).Result()
}

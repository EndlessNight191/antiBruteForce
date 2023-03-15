package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func InitCache() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_HOST") + ":" + strconv.Itoa(viper.GetInt("REDIS_PORT")),
		Password: viper.GetString("REDIS_PASSWORD"),
		DB:       viper.GetInt("REDIS_DB_INDEX"),
	})

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if status := RedisClient.WithContext(ctx).Ping(); status.Err() != nil {
		return fmt.Errorf("error init redis: %w", status.Err())
	}

	return nil
}

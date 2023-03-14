package cache

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func Init() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_ADRESS"),
		Password: viper.GetString("REDIS_PASSWORD"),
		DB:       viper.GetInt("REDIS_DB_INDEX"),
	})

	if _, err := RedisClient.Ping().Result(); err != nil {
		return fmt.Errorf("error init redis: %w", err)
	}
	return nil
}
package cache

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_ADRESS"),
		Password: viper.GetString("REDIS_PASSWORD"),
		DB:       viper.GetInt("REDIS_DB_INDEX"),
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}

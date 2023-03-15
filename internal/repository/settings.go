package repository

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func GetLimitSettingInt(keyRedis string, keyEnv string) (int, error) {
	result, err := getSettingFromRedis(string(keyRedis), string(keyEnv))
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(result)
	if err != nil {
		return 0, fmt.Errorf("atoi: %v", err)
	}

	return num, nil
}

func getSettingFromRedis(keyRedis string, keyEnv string) (string, error) {
	value, err := RedisClient.Get(keyRedis).Result()

	if err != nil {
		return "", fmt.Errorf("get key: %v", err)
	}

	if err == redis.Nil {
		settingEnv := viper.GetString(keyEnv)
		RedisClient.Append(keyRedis, settingEnv)
		return settingEnv, nil
	}

	return value, nil
}

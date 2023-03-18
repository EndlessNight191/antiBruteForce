package repository

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func (r *ClientRepository) GetLimitSettingInt(keyRedis string, keyEnv string) (int, error) {
	result, err := r.getSettingFromRedis(string(keyRedis), string(keyEnv))
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(result)
	if err != nil {
		return 0, fmt.Errorf("atoi: %v", err)
	}

	return num, nil
}

func (r *ClientRepository) getSettingFromRedis(keyRedis string, keyEnv string) (string, error) {
	value, err := r.redisClient.Get(keyRedis).Result()

	if err != nil {
		return "", fmt.Errorf("get key: %v", err)
	}

	if err == redis.Nil {
		settingEnv := viper.GetString(keyEnv)
		r.redisClient.Append(keyRedis, settingEnv)
		return settingEnv, nil
	}

	return value, nil
}

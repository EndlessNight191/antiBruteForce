package repository

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func (r *ClientRepository) GetLimitSettingInt(keyRedis string) (int, error) {
	result, err := r.GetSettingFromRedis(string(keyRedis))
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(result)
	if err != nil {
		return 0, fmt.Errorf("atoi: %v", err)
	}

	return num, nil
}

func (r *ClientRepository) GetSettingFromRedis(key string) (string, error) {
	value, err := r.redisClient.Get(key).Result()

	if err != nil {
		return "", fmt.Errorf("get key: %v", err)
	}

	if err == redis.Nil {
		settingEnv := viper.GetString(key)
		r.redisClient.Append(key, settingEnv)
		return settingEnv, nil
	}

	return value, nil
}


func (r *ClientRepository) UpdateSetting(keyRedis string, newValue int) error {
	err := r.redisClient.Set(string(keyRedis), newValue, 0).Err()
	if err != nil {
		return fmt.Errorf("update key: %v", err)
	}

	return nil
}
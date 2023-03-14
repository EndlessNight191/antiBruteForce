package repository

import (
	"fmt"
	"strconv"
	"test/internal/Infrastructure/cache"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func GetLimitCommon() (int, error) {
	result, err := getSettingFromRedis("maxLimitCommon", "MAX_LIMIT_COMMON")
	if err != nil {
		return 0, err //fmt.error
	}

	num, err := strconv.Atoi(result)
	if err != nil {
		return 0, fmt.Errorf("atoi: %v", err)
	}

	return num, nil
}

func GetLimitIp() (int, error) {
	result, err := getSettingFromRedis("maxLimitIp", "MAX_LIMIT_IP")
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func GetLimitLogin() (int, error) {
	result, err := getSettingFromRedis("maxLimitLogin", "MAX_LIMIT_LOGIN")
	if err != nil {
		return 0, err
	}
	num, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func GetLimitPassword(a string, b string) (int, error) {
	result, err := getSettingFromRedis(a, b)
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func getSettingFromRedis(keyRedis string, keyEnv string) (string, error) {
	value, err := cache.RedisClient.Get(keyRedis).Result()
	if err == redis.Nil {
		settingEnv := viper.GetString(keyEnv)
		cache.RedisClient.Set(keyRedis, settingEnv, 600000)
		cache.RedisClient.Append(keyRedis, settingEnv)
		return settingEnv, nil
	}

	// if err != nil {
	//     return "", err
	// }

	return value, nil
}

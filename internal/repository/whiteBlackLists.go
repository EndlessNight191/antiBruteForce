package repository

import (
	"test/internal/Infrastructure/cache"
)

func CheckWhiteList(ip string) (bool, error) {
    isInWhiteList, err := cache.RedisClient.SIsMember("whiteList", string(ip)).Result()
    if err != nil {
        return false, err
    }

    return isInWhiteList, nil
}

func CheckBlackList(ip string) (bool, error) {
    isInBlackList, err := cache.RedisClient.SIsMember("blackList", string(ip)).Result()
    if err != nil {
        return false, err
    }

    return isInBlackList, nil
}

func AddWhiteList(ip string) error {
    if err := cache.RedisClient.SAdd("whiteList", string(ip)).Err(); err != nil {
        return err
    }
    return nil
}

func AddBlackList(ip string) error {
    if err := cache.RedisClient.SAdd("blackList", string(ip)).Err(); err != nil {
        return err
    }
    return nil
}
package repository

import "test/internal/Infrastructure/cache"


func IncrementAuthAttemptsCommon(ip string, login string, password string) (int64, error) {
    key := ip + ":" + login + ":" + password
    result, err := cache.RedisClient.Incr(key).Result()
    if err != nil {
        return 0, err
    }

	if result == 1 {
		if err = cache.RedisClient.Expire(key, 60000).Err(); err != nil {
			return 0, err
		}  // потом создать глобальную переменную или настройку для експаира?
	}
    return result, nil
}

func IncrementIp(ip string) (int64, error) {
    key := "ip:" + ip
    result, err := cache.RedisClient.Incr(key).Result()
    if err != nil {
        return 0, err
    }

	if result == 1 {
		if err = cache.RedisClient.Expire(key, 60000).Err(); err != nil {
			return 0, err
		}  // потом создать глобальную переменную или настройку для експаира?
	}
    return result, nil
}

func IncrementLogin(login string) (int64, error) {
    key := "login:" + login
    result, err := cache.RedisClient.Incr(key).Result()
    if err != nil {
        return 0, err
    }

	if result == 1 {
		if err = cache.RedisClient.Expire(key, 60000).Err(); err != nil {
			return 0, err
		}  // потом создать глобальную переменную или настройку для експаира?
	}
    return result, nil
}

func IncrementPassword(password string) (int64, error) {
    key := "password:" + password
    result, err := cache.RedisClient.Incr(key).Result()
    if err != nil {
        return 0, err
    }

	if result == 1 {
		if err = cache.RedisClient.Expire(key, 60000).Err(); err != nil {
			return 0, err
		}  // потом создать глобальную переменную или настройку для експаира?
	}
    return result, nil
}
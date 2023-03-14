package repository

import (
	"test/internal/Infrastructure/cache"
	"test/internal/domain"
)

func IncrementAuthAttemptsCommon(request domain.IncomingRequest) (int64, error) {
	key := joinToFormatCommon(request.IP, request.Login, request.Password)

	result, err := cache.RedisClient.Incr(key).Result()
	if err != nil {
		return 0, err //fmt.errorf
	}

	if result == 1 {
		if err := cache.RedisClient.Expire(key, 60000).Err(); err != nil {
			return 0, err //fmt.errorf
		} // потом создать глобальную переменную или настройку для експаира?
	}

	return result, nil
}

func IncrementIp(ip string) (int64, error) {
	key := joinToFormatIp(ip)
	result, err := cache.RedisClient.Incr(key).Result()
	if err != nil {
		return 0, err
	}

	if result == 1 {
		if err = cache.RedisClient.Expire(key, 60000).Err(); err != nil {
			return 0, err
		} // потом создать глобальную переменную или настройку для експаира?
	}
	return result, nil
}

func IncrementLogin(login string) (int64, error) {
	key := joinToFormatLogin(login)
	result, err := cache.RedisClient.Incr(key).Result()
	if err != nil {
		return 0, err
	}

	if result == 1 {
		if err = cache.RedisClient.Expire(key, 60000).Err(); err != nil {
			return 0, err
		} // потом создать глобальную переменную или настройку для експаира?
	}
	return result, nil
}

func IncrementPassword(password string) (int64, error) {
	key := joinToFormatPassword(password)
	result, err := cache.RedisClient.Incr(key).Result()
	if err != nil {
		return 0, err
	}

	if result == 1 {
		if err = cache.RedisClient.Expire(key, 60000).Err(); err != nil {
			return 0, err
		} // потом создать глобальную переменную или настройку для експаира?
	}
	return result, nil
}

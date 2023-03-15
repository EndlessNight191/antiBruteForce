package repository

func IncrementByKey(key string) (int64, error) {
	result, err := RedisClient.Incr(key).Result()
	if err != nil {
		return 0, err
	}

	if result == 1 {
		if err = RedisClient.Expire(key, 60000).Err(); err != nil {
			return 0, err
		} // потом создать глобальную переменную или настройку для експаира?
	}
	return result, nil
}
package repository

func IncrementByKey(key string) (int64, error) {
	result, err := RedisClient.Incr(key).Result()
	if err != nil {
		return 0, err
	}

	if result == 1 {
        if err := addExpair(key); err != nil {
            return 0, err
        }
	}
	return result, nil
}
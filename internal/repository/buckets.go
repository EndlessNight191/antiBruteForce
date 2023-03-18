package repository

type bucket struct {
	cache CacheR
}

func (b bucket) IncrementByKey(key string) (int64, error) {
	b.cache.Incr()
	result, err := r.redisClient.Incr(key).Result()
	if err != nil {
		return 0, err
	}

	if result == 1 {
		if err := r.addExpair(key); err != nil {
			return 0, err
		}
	}

	return result, nil
}

package repository

func (r *ClientRepository) IncrementByKey(key string) (int64, error) {
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
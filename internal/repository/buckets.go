package repository

import "fmt"

func (r *ClientRepository) IncrementByKey(key string) (int64, error) {
	result, err := r.redisClient.Incr(key).Result()
	if err != nil {
		return 0, fmt.Errorf("error increment key redis: %w", err)
	}

	if result == 1 {
        if err := r.addExpair(key); err != nil {
            return 0, fmt.Errorf("error add expair  key redis: %w", err)
        }
	}
	return result, nil
}

func (r *ClientRepository) DeleteByKey(key string) error {
	if err := r.redisClient.Del(key).Err(); err != nil {
		return fmt.Errorf("error delete key redis: %w", err)
	}
	
	return nil
}
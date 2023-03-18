package repository

import "github.com/go-redis/redis"

type ClientRepository struct {
    redisClient *redis.Client
}

func NewRedis(redisClient *redis.Client) *ClientRepository {
    return &ClientRepository{redisClient}
}

type Repository interface {
	CheckWhiteList  	(ip string) 						(bool, error)
	CheckBlackList  	(ip string) 						(bool, error)
	AddWhiteList 		(ip string)  						(error)
	AddBlackList 		(ip string)  						(error)
	GetLimitSettingInt  (keyRedis string, keyEnv string) 	(int, error)
	getSettingFromRedis (keyRedis string, keyEnv string) 	(string, error)
	addExpair			(string) 							(error)
	IncrementByKey 		(string) 							(int64, error)
}

func NewRepository(redisClient *redis.Client) Repository {
    return &ClientRepository{redisClient}
}
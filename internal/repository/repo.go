package repository

import (
	"test/internal/domain"

	"github.com/go-redis/redis"
)

type ClientRepository struct {
    redisClient *redis.Client
	setting *domain.ConfigSetting
}

type Repository interface {
	CheckWhiteList  	(ip string) 						(bool, error)
	CheckBlackList  	(ip string) 						(bool, error)
	AddWhiteList 		(ip string)  						(error)
	AddBlackList 		(ip string)  						(error)
	GetLimitSettingInt  (keyRedis string) 					(int, error)
	GetSettingFromRedis (keyRedis string) 					(string, error)
	addExpair			(string) 							(error)
	DeleteByKey			(string)							(error)
	IncrementByKey 		(string) 							(int64, error)
	RemoveFromWhiteList (string)							(error)
	RemoveFromBlackList	(string)							(error)
	UpdateSetting		(keyRedis string, newValue int) 	(error)
}

func NewRepository(redisClient *redis.Client, setting *domain.ConfigSetting) Repository {
    return &ClientRepository{redisClient, setting}
}
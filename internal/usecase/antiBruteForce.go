package usecase

import (
	"test/internal/domain"
	"test/internal/repository"
)

type status int

const (
	StatusBlack = status(iota)
	StatusWhite
	StatusNone
)

const (
	maxLimitCommon = "maxLimitCommon"
	maxLimitIp = "maxLimitIp"
	maxLimitLogin = "maxLimitLogin"
	maxLimitPassword= "maxLimitPassword"
	maxLimitCommonEnv = "MAX_LIMIT_COMMON"
	maxLimitIpEnv = "MAX_LIMIT_IP"
	maxLimitLoginEnv = "MAX_LIMIT_LOGIN"
	maxLimitPasswordEnv= "MAX_LIMIT_PASSWORD"
)

type access struct {
	isBlack bool
	isWhite bool
}

func AllowAccess(request domain.IncomingRequest) (domain.ResponseIsAccess, error) {
	hashPassword, err := hashPassword(request.Password)
	if err != nil {
		return domain.ResponseIsAccess{}, err
	}
	request.Password = hashPassword

	isLists, err := checkIpInLists(request.IP)
	if err != nil {
		return domain.ResponseIsAccess{}, err
	}

	if isLists == StatusWhite {
		return domain.ResponseIsAccess{IsAccess: true}, nil // есть доступ по вайт листу
	}

	if isLists == StatusBlack {
		return domain.ResponseIsAccess{}, nil // нет доступа по вайт листу
	}

	isAccess, err := checkBackets(request)
	if err != nil || !isAccess {
		return domain.ResponseIsAccess{}, err
	}

	return domain.ResponseIsAccess{IsAccess: true}, nil
}

func checkIpInLists(ip string) (status, error) {
	isBlack, err := repository.CheckBlackList(ip)
	if err != nil || isBlack {
		return StatusBlack, err
	}

	isWhite, err := repository.CheckWhiteList(ip)
	if err != nil || isWhite {
		return StatusWhite, err
	}

	return StatusNone, nil
}

func checkBackets(request domain.IncomingRequest) (bool, error) {
	isAccessCommon, err := checkBacketCommon(request)
	if err != nil || !isAccessCommon {
		return false, err
	}

	isAccessIp, err := checkBacketIp(request.IP)
	if err != nil || !isAccessIp {
		return false, err
	}

	isAccessLogin, err := checkBacketLogin(request.Login)
	if err != nil || !isAccessLogin {
		return false, err
	}

	isAccessPassword, err := checkBacketPassword(request.Password)
	if err != nil || !isAccessPassword {
		return false, err
	}

	return true, nil
}

func checkBacketCommon(request domain.IncomingRequest) (bool, error) {
	key := joinToFormatCommon(request.IP, request.Login, request.Password)
	count, err := repository.IncrementByKey(key)
	if err != nil {
		return false, err
	}

	countLimit, err := repository.GetLimitSettingInt(maxLimitCommon, maxLimitCommonEnv)
	if err != nil {
		return false, err
	}

	if count > int64(countLimit) {
		return false, nil
	}

	return true, nil
}

func checkBacketIp(ip string) (bool, error) {
	key := joinToFormatIp(ip)
	count, err := repository.IncrementByKey(key)
	if err != nil {
		return false, err
	}

	countLimit, err := repository.GetLimitSettingInt(maxLimitIp, maxLimitIpEnv)
	if err != nil {
		return false, err
	}

	if count > int64(countLimit) {
		return false, nil
	}

	return true, nil
}

func checkBacketLogin(login string) (bool, error) {
	key := joinToFormatLogin(login)
	count, err := repository.IncrementByKey(key)
	if err != nil {
		return false, err
	}

	countLimit, err := repository.GetLimitSettingInt(maxLimitLogin, maxLimitLoginEnv)
	if err != nil {
		return false, err
	}

	if count > int64(countLimit) {
		return false, nil
	}

	return true, nil
}

func checkBacketPassword(password string) (bool, error) {
	key := joinToFormatPassword(password)
	count, err := repository.IncrementByKey(key)
	if err != nil {
		return false, err
	}

	countLimit, err := repository.GetLimitSettingInt(maxLimitPassword, maxLimitPasswordEnv)
	if err != nil {
		return false, err
	}

	if count > int64(countLimit) {
		return false, nil
	}

	return true, nil
}

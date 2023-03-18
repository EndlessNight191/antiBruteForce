package usecase

import (
	"test/internal/domain"
)

type status int

const (
	StatusBlack = status(iota)
	StatusWhite
	StatusNone
)

const (
	maxLimitCommon      = "maxLimitCommon"
	maxLimitIp          = "maxLimitIp"
	maxLimitLogin       = "maxLimitLogin"
	maxLimitPassword    = "maxLimitPassword"
	maxLimitCommonEnv   = "MAX_LIMIT_COMMON"
	maxLimitIpEnv       = "MAX_LIMIT_IP"
	maxLimitLoginEnv    = "MAX_LIMIT_LOGIN"
	maxLimitPasswordEnv = "MAX_LIMIT_PASSWORD"
)

func (uc UseCase) AllowAccess(request domain.IncomingRequest) (domain.ResponseIsAccess, error) {
	hashPassword, err := hashPassword(request.Password)
	if err != nil {
		return domain.ResponseIsAccess{}, err
	}
	request.Password = hashPassword

	isLists, err := uc.checkIpInLists(request.IP)
	if err != nil {
		return domain.ResponseIsAccess{}, err
	}

	if isLists == StatusWhite {
		return domain.ResponseIsAccess{IsAccess: true}, nil // есть доступ по вайт листу
	}

	if isLists == StatusBlack {
		return domain.ResponseIsAccess{}, nil // нет доступа по вайт листу
	}

	isAccess, err := uc.checkBackets(request)
	if err != nil || !isAccess {
		return domain.ResponseIsAccess{}, err
	}

	return domain.ResponseIsAccess{IsAccess: true}, nil
}

func (uc *UseCase) checkIpInLists(ip string) (status, error) {
	ip = deleteIpMask(ip)

	isBlack, err := uc.repo.CheckBlackList(ip)
	if err != nil || isBlack {
		return StatusBlack, err
	}

	isWhite, err := uc.repo.CheckWhiteList(ip)
	if err != nil || isWhite {
		return StatusWhite, err
	}

	return StatusNone, nil
}

func (uc *UseCase) checkBackets(request domain.IncomingRequest) (bool, error) {
	isAccessCommon, err := uc.checkBacketCommon(request)
	if err != nil || !isAccessCommon {
		return false, err
	}

	isAccessIp, err := uc.checkBacketIp(request.IP)
	if err != nil || !isAccessIp {
		return false, err
	}

	isAccessLogin, err := uc.checkBacketLogin(request.Login)
	if err != nil || !isAccessLogin {
		return false, err
	}

	isAccessPassword, err := uc.checkBacketPassword(request.Password)
	if err != nil || !isAccessPassword {
		return false, err
	}

	return true, nil
}

func (uc *UseCase) checkBacketCommon(request domain.IncomingRequest) (bool, error) {
	key := joinToFormatCommon(request.IP, request.Login, request.Password)
	count, err := uc.repo.IncrementByKey(key)
	if err != nil {
		return false, err
	}

	countLimit, err := uc.repo.GetLimitSettingInt(maxLimitCommon, maxLimitCommonEnv)
	if err != nil {
		return false, err
	}

	if count > int64(countLimit) {
		return false, nil
	}

	return true, nil
}

func (uc *UseCase) checkBacketIp(ip string) (bool, error) {
	key := joinToFormatIp(ip)
	count, err := uc.repo.IncrementByKey(key)
	if err != nil {
		return false, err
	}

	countLimit, err := uc.repo.GetLimitSettingInt(maxLimitIp, maxLimitIpEnv)
	if err != nil {
		return false, err
	}

	if count > int64(countLimit) {
		return false, nil
	}

	return true, nil
}

func (uc *UseCase) checkBacketLogin(login string) (bool, error) {
	key := joinToFormatLogin(login)
	count, err := uc.repo.IncrementByKey(key)
	if err != nil {
		return false, err
	}

	countLimit, err := uc.repo.GetLimitSettingInt(maxLimitLogin, maxLimitLoginEnv)
	if err != nil {
		return false, err
	}

	if count > int64(countLimit) {
		return false, nil
	}

	return true, nil
}

func (uc *UseCase) checkBacketPassword(password string) (bool, error) {
	key := joinToFormatPassword(password)
	count, err := uc.repo.IncrementByKey(key)
	if err != nil {
		return false, err
	}

	countLimit, err := uc.repo.GetLimitSettingInt(maxLimitPassword, maxLimitPasswordEnv)
	if err != nil {
		return false, err
	}

	if count > int64(countLimit) {
		return false, nil
	}

	return true, nil
}

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

	if isLists.isWhite {
		return domain.ResponseIsAccess{IsAccess: true}, nil // есть доступ по вайт листу
	}

	if isLists.isBlack {
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
	if err != nil {
		return StatusBlack, err
	}

	isWhite, err := repository.CheckWhiteList(ip)
	if err != nil {
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

// инкрементить редис, проверять, получать лимит бакета, проверять не превысился ли лимит
func checkBacketCommon(request domain.IncomingRequest) (bool, error) {
	count, err := repository.IncrementAuthAttemptsCommon(request)
	if err != nil {
		return false, err
	}

	countLimit, err := repository.GetLimitCommon()
	if err != nil {
		return false, err
	}

	if count > int64(countLimit) {
		return false, nil
	}

	return true, nil
}

func checkBacketIp(ip string) (bool, error) {
	count, err := repository.IncrementIp(ip)
	if err != nil {
		return false, err
	}

	countLimit, err := repository.GetLimitIp()
	if err != nil {
		return false, err
	}

	if count > int64(countLimit) {
		return false, nil
	}

	return true, nil
}

func checkBacketLogin(login string) (bool, error) {
	count, err := repository.IncrementLogin(login)
	if err != nil {
		return false, err
	}

	countLimit, err := repository.GetLimitLogin()
	if err != nil {
		return false, err
	}

	if count > int64(countLimit) {
		return false, nil
	}

	return true, nil
}

func checkBacketPassword(password string) (bool, error) {
	count, err := repository.IncrementPassword(password)
	if err != nil {
		return false, err
	}

	countLimit, err := repository.GetLimitPassword("max", "hui")
	if err != nil {
		return false, err
	}

	if count > int64(countLimit) {
		return false, nil
	}

	return true, nil
}

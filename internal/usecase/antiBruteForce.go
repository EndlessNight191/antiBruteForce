package usecase

import (
	"test/internal/domain"
	"test/internal/repository"
)

func AllowAccess(request domain.IncomingRequest) (bool, error) {
	
	isLists, err := checkIpInLists(request.IP)
	if err != nil {
		return false, err
	}

	if *isLists {
		return true, nil // есть доступ по вайт листу
	}

	if !*isLists {
		return false, nil // нет доступап по блек листу
	}

	// логика с бакетами, ipшника в листах нету


	return false, nil
}

func checkIpInLists(ip string) (*bool, error) {
	f := false
	t := true

	isNotAccess, err := repository.CheckBlackList(ip)
	if (err != nil) {
		return &f, err
	}
	if (isNotAccess) {
		return &f, nil
	}

	isAccess, err := repository.CheckWhiteList(ip)
	if (err != nil) {
		return &f, err
	}
	if (isAccess) {
		return &t, nil
	}

	return nil, nil
}
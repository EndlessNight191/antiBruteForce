package usecase

import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) (string, error) {
	passwordByte := []byte(string(password))

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }

	return string(hashedPassword), nil
}

func joinToFormatCommon(ip string, login string, password string) string {
	key := ip + ":" + login + ":" + password
	return key
}

func joinToFormatIp(ip string) string {
	key := "ip:" + ip
	return key
}

func joinToFormatLogin(login string) string {
	key := "login:" + login
	return key
}

func joinToFormatPassword(password string) string {
	key := "password:" + password
	return key
}
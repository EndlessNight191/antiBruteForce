package usecase

import (
	"strings"

	"golang.org/x/crypto/bcrypt"
)

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

func deleteIpMask(ip string) (string) {
	isMask := strings.Contains(ip, "/") 
	
	if !isMask {
		return ip
	}

	parts := strings.Split(ip, "/")
    return parts[0]
}

// func checkNetworkIp(ip string, ipMask string) (bool, error) {
// 	ipParse := net.ParseIP(ip)
//     _, subnet, err := net.ParseCIDR(ipMask)
// 	if err != nil {
// 		return false, fmt.Errorf("ParseCIDR error: %v", err)
// 	}

//     if subnet.Contains(ipParse) {
//         return true, nil // принадлежит под сети
//     } else {
//         return false, nil // не принадлежит под сети
//     }
// }
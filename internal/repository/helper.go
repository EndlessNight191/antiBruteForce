package repository

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
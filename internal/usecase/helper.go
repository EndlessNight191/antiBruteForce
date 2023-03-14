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
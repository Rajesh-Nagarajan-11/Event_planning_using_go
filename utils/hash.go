package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(res), err
}
func Checkpassword(password string, retrivepassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(retrivepassword), []byte(password))
	return err == nil
}

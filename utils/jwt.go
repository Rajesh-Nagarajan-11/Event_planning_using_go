package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "superkey"

func GenerateToken(email string, userid int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userid": userid,
		"exp":    time.Now().Add(2 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(secretkey))
}
func VerifyToken(token string) (float64, error) {
	parsedtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected sigin method")
		}
		return []byte(secretkey), nil
	})
	if err != nil {
		return 0, errors.New("couldn't parse token")
	}
	tokenisvalid := parsedtoken.Valid
	if !tokenisvalid {
		return 0, errors.New("invalid Token")
	}
	claims, ok := parsedtoken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid Token claims")
	}
	//email := claim["email"].(string)
	userid := claims["userid"].(float64)
	return userid, nil

}

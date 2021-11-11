package common

import (
	"rlab/model/user"
	"time"

	"github.com/golang-jwt/jwt"
)

const term = 6 * time.Hour

// TODO: 放到config中
var jwtKey = []byte("remote lab")

type claims struct {
	user.UserInfo
	jwt.StandardClaims
}

func ReleaseToken(user user.UserInfo) (string, error) {
	claim := &claims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(term).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "rlab",
			Subject:   "student",
		},
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString(jwtKey)
}

func ParseToken(tokenString string) (*claims, error) {

	// 将token解析结果保存到claims中
	claims := &claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return claims, err
}

package helper

import (
	"errors"
	"fmt"
	"github.com/shawu21/test/common"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaim struct {
	Student_number string
	jwt.RegisteredClaims
}

var myKey = []byte("luxuetaoxuelu")

func CreatToken(student_number string) (string, error) {
	claim := JwtClaim{
		Student_number: student_number,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().In(common.ChinaTime).Add(168 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().In(common.ChinaTime)),
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func VerifyToken(token string) (string, error) {
	tempToken, err := jwt.ParseWithClaims(token, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := tempToken.Claims.(*JwtClaim)
	fmt.Println(claims)
	if !ok {
		return "", errors.New("claims error")
	}
	if err := tempToken.Claims.Valid(); err != nil {
		return "", err
	}
	return claims.Student_number, nil
}

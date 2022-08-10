package helper

import (
	"errors"
	"log"
	"test/common"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtClaim struct {
	jwt.StandardClaims
	Student_number string `json:"student_number"`
}

func CreatrToken(student_number string) string {
	claim := JwtClaim{}
	claim.Student_number = student_number
	claim.ExpiresAt = time.Now().In(common.ChinaTime).Add(168 * time.Hour).Unix()
	claim.IssuedAt = time.Now().In(common.ChinaTime).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString([]byte("Matilda"))
	if err != nil {
		log.Println(err)
	}
	return signedToken
}

func VerifyToken(token string) (string, error) {
	TempToken, err := jwt.ParseWithClaims(token, &JwtClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("Matilda"), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := TempToken.Claims.(*JwtClaim)
	if !ok {
		return "", errors.New("发生错误")
	}
	if err := TempToken.Claims.Valid(); err != nil {
		return "", err
	}
	return claims.Student_number, nil
}

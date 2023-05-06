package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const MaxAge = 81600 * 7 //s
const mySecret = "1234567aa"

type MyClaims struct {
	Data interface{}
	jwt.StandardClaims
}

func SignDataToToken(signData interface{}) string {
	newWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS512, &MyClaims{
		Data: signData,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(MaxAge) * time.Second).Unix(),
		},
	})
	signedString, err := newWithClaims.SignedString([]byte(mySecret))
	if err != nil {
		return ""
	}
	return signedString
}

func DecodeSignStringToData(signedString string, password string) (data *MyClaims, err error) {
	res, err := jwt.ParseWithClaims(signedString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(mySecret), nil
	})
	claims := res.Claims.(*MyClaims)
	return claims, err
}

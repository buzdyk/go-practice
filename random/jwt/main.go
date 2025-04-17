package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims // https://datatracker.ietf.org/doc/html/rfc7519
	UserID               int
}

const TokenExp = time.Hour * 3
const SecretKey = "secret"

func main() {
	tokenString, err := BuildJWTString()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tokenString)
	fmt.Println(GetUserID(tokenString))
}

func BuildJWTString() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExp)),
		},
		UserID: 1,
	})

	if tokenString, err := token.SignedString([]byte(SecretKey)); err != nil {
		return "", err
	} else {
		return tokenString, nil
	}
}

func GetUserID(tokenString string) (int, error) {
	claims := &Claims{}

	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(SecretKey), nil
	}

	if token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc); err != nil {
		return 0, err
	} else if token.Valid == false {
		return 0, errors.New("token is not valid")
	}

	return claims.UserID, nil
}

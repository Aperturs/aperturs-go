package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtClaim struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

var jwtSecret = []byte(getJwtSecret())

func getJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return ""
	}
	return secret
}
func JwtGenerate(ctx context.Context, userID string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtClaim{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	token, err := t.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
func JwtValidate(ctx context.Context,token string) (*jwt.Token,error){
	return jwt.ParseWithClaims(
		token,
		&JwtClaim{},
		func(t *jwt.Token) (interface{}, error) {
			_,ok := t.Method.(*jwt.SigningMethodHMAC)
			if !ok{
				return nil,fmt.Errorf("There is a problem with the signing method")
			}
			return jwtSecret,nil
		})
}
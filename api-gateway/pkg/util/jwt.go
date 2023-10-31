package util

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type TokenClaims struct {
	jwt.RegisteredClaims
	UserId int64  `json:"userId"`
	Jti    string `json:"jti"`
}

func GetToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if len(authHeader) <= len("Bearer")+1 {
		return "", errors.New("Unauthorized")
	}

	tokenString := authHeader[len("Bearer")+1:]

	return tokenString, nil
}

func keyFunc(key string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New("Unauthorized")
		}
		return []byte(key), nil
	}
}

func ParseAccessToken(token string) (*TokenClaims, error) {
	claims := TokenClaims{}

	_, err := jwt.ParseWithClaims(token, &claims, keyFunc("key"))
	if err != nil {
		return nil, err
	}

	return &claims, nil
}

func OpaqueueToken(token string) string {
	return fmt.Sprintf("OPAQUEUE%v", token)
}

func DecodeOpaqueueTokenToJWTToken(opaqueueToken string) string {
	jwtToken := opaqueueToken[len("OPAQUEUE"):]
	return jwtToken
}

package db

import (
	"auth_svc/pkg/config"
	"auth_svc/pkg/models"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func keyFunc(key string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, errors.New(config.UNAUTHORIZED)
		}
		return []byte(key), nil
	}
}

func GetToken(ctx *gin.Context) (string, error) {
	authHeader := ctx.GetHeader("Authorization")
	if len(authHeader) <= len("Bearer")+1 {
		return "", errors.New(config.UNAUTHORIZED)
	}

	tokenString := authHeader[len("Bearer")+1:]

	return tokenString, nil
}

func ValidateAccessToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, keyFunc("key"))
}

func ParseAccessToken(token string) (*models.TokenClaims, error) {
	claims := models.TokenClaims{}

	_, err := jwt.ParseWithClaims(token, &claims, keyFunc("key"))
	if err != nil {
		return nil, err
	}

	return &claims, nil
}

func GenerateAccessToken(user *models.Users) (string, error) {
	// Generate random
	random, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	// Generate claims
	claims := models.TokenClaims{
		UserId: user.Id,
		Jti:    random.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign with secret
	signedToken, err := token.SignedString([]byte("key"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

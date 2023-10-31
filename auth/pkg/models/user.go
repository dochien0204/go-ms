package models

import "github.com/golang-jwt/jwt/v4"

type Users struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenClaims struct {
	jwt.RegisteredClaims
	UserId int64  `json:"userId"`
	Jti    string `json:"jti"`
}

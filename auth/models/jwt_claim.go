package models

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

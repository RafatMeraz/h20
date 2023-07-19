package models

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	UserId uint `json:"userId"`
	jwt.StandardClaims
}

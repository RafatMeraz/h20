package controllers

import (
	"github.com/RafatMeraz/h20/auth/models"
	"github.com/RafatMeraz/h20/config"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTTokenController struct{}

func (JWTTokenController) CreateJwtToken(userId uint) (string, int64, error) {
	expiryTime := time.Now().Add(config.AppConfiguration.TokenValidationTime) // expire time of token
	claims := models.JwtCustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(config.AppConfiguration.TokenSecret)
	if err != nil {
		return "", 0, err
	}

	return token, expiryTime.Unix(), nil
}

func (JWTTokenController) CheckTokenValidation(claims models.JwtCustomClaims) bool {
	err := claims.Valid()
	return err != nil
}

func (JWTTokenController) GetClaimsFromToken(strToken string) (models.JwtCustomClaims, error) {
	var claims models.JwtCustomClaims
	_, err := jwt.ParseWithClaims(strToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return config.AppConfiguration.TokenSecret, nil
	})
	if err != nil {
		return models.JwtCustomClaims{}, err
	}
	return claims, nil
}

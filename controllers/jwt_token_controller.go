package controllers

import (
	"github.com/RafatMeraz/h20/models"
	"github.com/golang-jwt/jwt"
	"os"
	"strconv"
	"time"
)

type JWTTokenController struct{}

func (JWTTokenController) CreateJwtToken(userId uint) (string, int64, error) {
	hours, _ := strconv.Atoi(os.Getenv("TOKEN_VALIDATION_TIME"))
	expiryTime := time.Now().Add(time.Hour * time.Duration(hours)) // expire time of token
	claims := models.JwtCustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return "", 0, err
	}

	return token, expiryTime.Unix(), nil
}

func (JWTTokenController) CheckTokenValidation(claims *models.JwtCustomClaims) bool {
	err := claims.Valid()
	if err != nil {
		return false
	}
	return true
}

func (JWTTokenController) GetClaimsFromToken(strToken string) (models.JwtCustomClaims, error) {
	var claims models.JwtCustomClaims
	_, err := jwt.ParseWithClaims(strToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return models.JwtCustomClaims{}, err
	}
	return claims, nil
}

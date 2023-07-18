package controllers

import (
	"github.com/RafatMeraz/h20/auth/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct{}

func (AuthController) Login(c echo.Context) error {
	var credentials models.Credentials
	if err := c.Bind(&credentials); err != nil {
		return err
	}
	if credentials.Username != "rafat" || credentials.Password != "pass" {
		return c.NoContent(http.StatusUnauthorized)
	}
	token, expiryAt, err := JWTTokenController{}.CreateJwtToken(1)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": token,
		"expiry_at":    expiryAt,
	})
}

package controllers

import (
	"github.com/RafatMeraz/h20/auth/models"
	"github.com/RafatMeraz/h20/auth/repositories"
	"github.com/RafatMeraz/h20/error_mapper"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	userRepository repositories.UserRepository
}

func (authController AuthController) Login(c echo.Context) error {
	var userDTO models.UserDTO

	if err := c.Bind(&userDTO); err != nil {
		return error_mapper.ErrorMapper{}.MapError(c, err)
	}

	userId, passwordCheckError := authController.userRepository.CheckPassword(userDTO)
	if passwordCheckError != nil {
		return error_mapper.ErrorMapper{}.MapError(c, passwordCheckError)
	}

	token, expiryAt, err := JWTTokenController{}.CreateJwtToken(userId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": token,
		"expiry_at":    expiryAt,
	})
}

package controllers

import (
	"github.com/RafatMeraz/h20/error_mapper"
	"github.com/RafatMeraz/h20/models"
	"github.com/RafatMeraz/h20/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthController struct {
	Repository repositories.UserRepository
}

func (authController AuthController) Login(c echo.Context) error {
	v := validator.New()
	var userReq models.UserRequest

	if err := c.Bind(&userReq); err != nil {
		return error_mapper.ErrorMapper{}.MapError(c, err)
	}

	if validationErr := v.Struct(userReq); validationErr != nil {
		return validationErr
	}

	userId, passwordCheckError := authController.Repository.CheckPassword(userReq)
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

func (authController AuthController) SignUp(c echo.Context) error {
	v := validator.New()
	var userReq models.UserRequest
	if err := c.Bind(&userReq); err != nil {
		return err
	}

	if validationErr := v.Struct(userReq); validationErr != nil {
		return validationErr
	}

	userExist, err := authController.Repository.CheckIfUserAlreadyExist(userReq)
	if err != nil {
		return err
	}
	if userExist {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "email already exists",
		})
	}

	if err := authController.Repository.CreateNewUser(userReq); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user created successfully",
	})
}

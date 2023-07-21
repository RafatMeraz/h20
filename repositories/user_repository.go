package repositories

import (
	"github.com/RafatMeraz/h20/database"
	"github.com/RafatMeraz/h20/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"strings"
)

type UserRepository struct{}

func (UserRepository) CreateNewUser(userRequest models.UserRequest) error {
	newUserModel := models.User{Name: userRequest.Name, Password: userRequest.Password, Email: userRequest.Email}
	result := database.Database.Instance().Create(&newUserModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (UserRepository) CheckIfUserAlreadyExist(request models.UserRequest) (bool, error) {
	var user models.User
	result := database.Database.Instance().Where("email = ?", strings.Trim(request.Email, " ")).Limit(1).Find(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, result.Error
	}
	return true, nil
}

func (UserRepository) CheckPassword(dto models.UserRequest) (uint, error) {
	newUserModel := models.User{
		Email:    dto.Email,
		Password: dto.Password,
	}

	result := database.Database.Instance().Where("email = ?", newUserModel.Email).Limit(1).Find(&newUserModel)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, gorm.ErrRecordNotFound
	}
	if newUserModel.Password != dto.Password {
		return 0, echo.ErrUnauthorized
	}
	return newUserModel.ID, nil
}

package repositories

import (
	"github.com/RafatMeraz/h20/auth/models"
	"github.com/RafatMeraz/h20/database"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (UserRepository) AddNewUser(dto models.UserDTO) error {
	newUserModel := models.User{Name: dto.Name, Password: dto.Password, Email: dto.Email}
	result := database.Database.Instance().Create(&newUserModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (UserRepository) CheckPassword(dto models.UserDTO) (uint, error) {
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

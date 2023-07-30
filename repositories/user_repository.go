package repositories

import "github.com/RafatMeraz/h20/models"

type UserRepository interface {
	CreateNewUser(userRequest models.UserRequest) error
	CheckIfUserAlreadyExist(request models.UserRequest) (bool, error)
	CheckPassword(dto models.UserRequest) (uint, error)
}

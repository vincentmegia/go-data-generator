package services

import (
	"github.com/vincentmegia/go-data-generator/internals/models"
	"github.com/vincentmegia/go-data-generator/internals/repositories"
)

type UserService struct {
	Users *[]models.User
}

var users *[]models.User = new([]models.User)

func (us *UserService) Init() {

}
func (us *UserService) AddUser(u *models.User) error {
	userRepository := repositories.UserRepository{}
	if error := userRepository.AddUser(u); error != nil {
		return error
	}
	return nil
}

// NOTE: no need to (check for existence, assumption is we create new data each time
func (us *UserService) BulkInsert(users *[]models.User) error {
	userRepository := repositories.UserRepository{}
	if error := userRepository.BulkInsert(users); error != nil {
		return error
	}
	return nil
}

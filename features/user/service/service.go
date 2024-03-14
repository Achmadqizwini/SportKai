package service

import (
	"errors"

	"fmt"
	"github.com/Achmadqizwini/SportKai/features/user"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository user.RepositoryInterface
	validate       *validator.Validate
}



func New(repo user.RepositoryInterface) user.ServiceInterface {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

// Create implements user.ServiceInterface
func (srv *userService) Create(input user.User) (err error) {
	input.PublicId = uuid.NewString()
	bytePass, errEncrypt := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if errEncrypt != nil {
		return errors.New("failed insert data, error query")
	}

	input.Password = string(bytePass)

	if errCreate := srv.userRepository.Create(input); errCreate != nil {
		return fmt.Errorf("failed to create new user: %v", err)
	}
	return nil
}

// Get implements user.ServiceInterface.
func (srv *userService) Get() ([]user.User, error) {
	userData, err := srv.userRepository.Get()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %v", err)
	}
	return userData, nil
}

// Update implements user.ServiceInterface.
func (srv *userService) Update(input user.User, id string) (user.User, error) {
	updatedUser, err := srv.userRepository.Update(input, id)
	if err != nil {
		return user.User{}, fmt.Errorf("failed to update user: %v", err)
	}
	return updatedUser, nil
}

// Delete implements user.ServiceInterface.
func (srv *userService) Delete(id string) error {
	err := srv.userRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}

// GetById implements user.ServiceInterface.
func (srv *userService) GetById(id string) (user.User, error) {
	res, err := srv.userRepository.GetById(id)
	if err != nil {
		return user.User{}, fmt.Errorf("failed to get user by id: %v", err)
	}
	return res, nil
}
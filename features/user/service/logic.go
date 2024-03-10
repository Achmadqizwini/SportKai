package service

import (
	"errors"

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
func (srv *userService) Create(input user.Core) (err error) {
	input.PublicId = uuid.NewString()
	bytePass, errEncrypt := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if errEncrypt != nil {
		return errors.New("failed insert data, error query")
	}

	input.Password = string(bytePass)

	if errCreate := srv.userRepository.Create(input); errCreate != nil {
		return errors.New("failed insert data, error query")
	}
	return nil
}

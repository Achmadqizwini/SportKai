package service

import (
	"github.com/Achmadqizwini/SportKai/features/user"
	"github.com/Achmadqizwini/SportKai/utils/logger"
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

var (
	logService = logger.NewLogger().Logger.With().Logger()
)

// Create implements user.ServiceInterface
func (svc *userService) Create(input user.User) (err error) {
	if errValidate := svc.validate.Struct(input); errValidate != nil {
		logService.Error().Err(errValidate).Msg("error validate input, please check your input")
		return errValidate
	}
	input.PublicId = uuid.NewString()
	bytePass, errEncrypt := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if errEncrypt != nil {
		return errEncrypt
	}

	input.Password = string(bytePass)

	if errCreate := svc.userRepository.Create(input); errCreate != nil {
		logService.Error().Err(err).Msg("failed to create new user")
		return errCreate
	}
	return nil
}

// Get implements user.ServiceInterface.
func (svc *userService) Get() ([]user.User, error) {
	userData, err := svc.userRepository.Get()
	if err != nil {
		logService.Error().Err(err).Msg("failed to retrieve users")
		return nil, err
	}
	return userData, nil
}

// Update implements user.ServiceInterface.
func (svc *userService) Update(input user.User, id string) (user.User, error) {
	updatedUser, err := svc.userRepository.Update(input, id)
	if err != nil {
		logService.Error().Err(err).Msg("failed to update user")
		return user.User{}, err
	}
	return updatedUser, nil
}

// Delete implements user.ServiceInterface.
func (svc *userService) Delete(id string) error {
	err := svc.userRepository.Delete(id)
	if err != nil {
		logService.Error().Err(err).Msg("failed to delete user")
		return err
	}
	return nil
}

// GetById implements user.ServiceInterface.
func (svc *userService) GetById(id string) (user.User, error) {
	res, err := svc.userRepository.GetById(id)
	if err != nil {
		logService.Error().Err(err).Msg("failed to get user by id")
		return user.User{}, err
	}
	return res, nil
}
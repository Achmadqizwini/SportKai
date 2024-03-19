package service

import (
	"github.com/Achmadqizwini/SportKai/features/user/model"
	repo "github.com/Achmadqizwini/SportKai/features/user/repository"
	"github.com/Achmadqizwini/SportKai/utils/logger"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ServiceInterface interface {
	Create(input model.User) error
	Get() ([]model.User, error)
	GetById(id string) (model.User, error)
	Update(input model.User, id string) (model.User, error)
	Delete(id string) error
}

type userService struct {
	userRepository repo.RepositoryInterface
	validate       *validator.Validate
}

func New(repo repo.RepositoryInterface) ServiceInterface {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

var (
	logService = logger.NewLogger().Logger.With().Logger()
)

// Create implements ServiceInterface
func (svc *userService) Create(input model.User) (err error) {
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

// Get implements ServiceInterface.
func (svc *userService) Get() ([]model.User, error) {
	userData, err := svc.userRepository.Get()
	if err != nil {
		logService.Error().Err(err).Msg("failed to retrieve users")
		return nil, err
	}
	return userData, nil
}

// Update implements ServiceInterface.
func (svc *userService) Update(input model.User, id string) (model.User, error) {
	updatedUser, err := svc.userRepository.Update(input, id)
	if err != nil {
		logService.Error().Err(err).Msg("failed to update user")
		return model.User{}, err
	}
	return updatedUser, nil
}

// Delete implements ServiceInterface.
func (svc *userService) Delete(id string) error {
	err := svc.userRepository.Delete(id)
	if err != nil {
		logService.Error().Err(err).Msg("failed to delete user")
		return err
	}
	return nil
}

// GetById implements ServiceInterface.
func (svc *userService) GetById(id string) (model.User, error) {
	res, err := svc.userRepository.GetById(id)
	if err != nil {
		logService.Error().Err(err).Msg("failed to get user by id")
		return model.User{}, err
	}
	return res, nil
}

package service

import (
	"github.com/Achmadqizwini/SportKai/features/auth/model"
	repo "github.com/Achmadqizwini/SportKai/features/auth/repository"
	"github.com/Achmadqizwini/SportKai/utils/logger"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type ServiceInterface interface {
	Login(input model.Auth) (model.Auth, error)
}

type authService struct {
	authRepository repo.RepositoryInterface
	validate       *validator.Validate
}

func New(repo repo.RepositoryInterface) ServiceInterface {
	return &authService{
		authRepository: repo,
		validate:       validator.New(),
	}
}

var logService = logger.NewLogger().Logger.With().Logger()

// Create implements auth.ServiceInterface.
func (c *authService) Login(input model.Auth) (model.Auth, error) {
	if errValidate := c.validate.Struct(input); errValidate != nil {
		logService.Error().Err(errValidate).Msg("error validate input, please check your input")
		return model.Auth{}, errValidate
	}
	input.PublicId = uuid.NewString()
	user, err := c.authRepository.FindUser(input)
	if err != nil {
		logService.Error().Err(err).Msg("failed to login, check you username or password")
		return model.Auth{}, err
	}
	errCheckPass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if errCheckPass != nil {
		return model.Auth{}, errCheckPass
	}

	return user, nil
}

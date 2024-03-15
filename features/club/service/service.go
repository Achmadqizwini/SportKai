package service

import (
	"github.com/Achmadqizwini/SportKai/features/club/model"
	repo "github.com/Achmadqizwini/SportKai/features/club/repository"
	"github.com/Achmadqizwini/SportKai/utils/logger"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	Create(input model.Club) error
	Get() ([]model.Club, error)
	GetById(id string) (model.Club, error)
	Update(input model.Club, id string) (model.Club, error)
	Delete(id string) error
}

type clubService struct {
	clubRepository repo.RepositoryInterface
	validate       *validator.Validate
}

func New(repo repo.RepositoryInterface) ServiceInterface {
	return &clubService{
		clubRepository: repo,
		validate:       validator.New(),
	}
}

var logService = logger.NewLogger().Logger.With().Logger()

// Create implements club.ServiceInterface.
func (c *clubService) Create(input model.Club) error {
	if errValidate := c.validate.Struct(input); errValidate != nil {
		logService.Error().Err(errValidate).Msg("error validate input, please check your input")
		return errValidate
	}
	input.PublicId = uuid.NewString()
	input.JoinedMember = 1
	if err := c.clubRepository.Create(input); err != nil {
		logService.Error().Err(err).Msg("failed to create new club")
		return err
	}
	return nil
}

// Delete implements club.ServiceInterface.
func (c *clubService) Delete(id string) error {
	panic("unimplemented")
}

// Get implements club.ServiceInterface.
func (c *clubService) Get() ([]model.Club, error) {
	panic("unimplemented")
}

// GetById implements club.ServiceInterface.
func (c *clubService) GetById(id string) (model.Club, error) {
	panic("unimplemented")
}

// Update implements club.ServiceInterface.
func (c *clubService) Update(input model.Club, id string) (model.Club, error) {
	panic("unimplemented")
}

package service

import (
	"github.com/Achmadqizwini/SportKai/features/clubMember/model"
	repo "github.com/Achmadqizwini/SportKai/features/clubMember/repository"
	"github.com/Achmadqizwini/SportKai/utils/logger"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	Create(input model.ClubMember) error
	Get() ([]model.ClubMember, error)
	GetById(id string) (model.ClubMember, error)
	Update(input model.ClubMember, id string) (model.ClubMember, error)
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

// Create implements ServiceInterface.
func (c *clubService) Create(input model.ClubMember) error {
	input.PublicId = uuid.NewString()
	if input.Status == "" {
		input.Status = "Admin"
	}
	return nil
}

// Delete implements ServiceInterface.
func (c *clubService) Delete(id string) error {
	panic("unimplemented")
}

// Get implements ServiceInterface.
func (c *clubService) Get() ([]model.ClubMember, error) {
	panic("unimplemented")
}

// GetById implements ServiceInterface.
func (c *clubService) GetById(id string) (model.ClubMember, error) {
	panic("unimplemented")
}

// Update implements ServiceInterface.
func (c *clubService) Update(input model.ClubMember, id string) (model.ClubMember, error) {
	panic("unimplemented")
}

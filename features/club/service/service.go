package service

import (
	"github.com/Achmadqizwini/SportKai/features/club/model"
	repo "github.com/Achmadqizwini/SportKai/features/club/repository"
	member "github.com/Achmadqizwini/SportKai/features/clubMember/model"
	memberRepo "github.com/Achmadqizwini/SportKai/features/clubMember/repository"
	"github.com/Achmadqizwini/SportKai/utils/logger"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	Create(input model.Club, user_id string) error
	Get() ([]model.Club, error)
	GetById(id string) (model.Club, error)
	Update(input model.Club, id string) (model.Club, error)
	Delete(id string) error
}

type clubService struct {
	clubRepository   repo.RepositoryInterface
	memberRepository memberRepo.RepositoryInterface
	validate         *validator.Validate
}

func New(repo repo.RepositoryInterface, mem memberRepo.RepositoryInterface) ServiceInterface {
	return &clubService{
		clubRepository:   repo,
		memberRepository: mem,
		validate:         validator.New(),
	}
}

var logService = logger.NewLogger().Logger.With().Logger()

// Create implements club.ServiceInterface.
func (c *clubService) Create(input model.Club, user_id string) error {
	if errValidate := c.validate.Struct(input); errValidate != nil {
		logService.Error().Err(errValidate).Msg("error validate input, please check your input")
		return errValidate
	}
	input.PublicId = uuid.NewString()
	lastId, err := c.clubRepository.Create(input)
	if err != nil {
		logService.Error().Err(err).Msg("failed to create new club")
		return err
	}
	memberInput := member.MemberPayload{
		PublicId: uuid.NewString(),
		ClubId:   lastId,
		UserId:   user_id,
		Status:   "Owner",
	}
	if err := c.memberRepository.Create(memberInput); err != nil {
		logService.Error().Err(err).Msg("failed to create new member")
		return err
	}
	return nil
}

// Delete implements club.ServiceInterface.
func (c *clubService) Delete(id string) error {
	err := c.clubRepository.Delete(id)
	if err != nil {
		logService.Error().Err(err).Msg("failed to delete club")
		return err
	}
	return nil
}

// Get implements club.ServiceInterface.
func (c *clubService) Get() ([]model.Club, error) {
	club, err := c.clubRepository.Get()
	if err != nil {
		logService.Error().Err(err).Msg("failed to retrieve club")
		return nil, err
	}
	return club, nil
}

// GetById implements club.ServiceInterface.
func (c *clubService) GetById(id string) (model.Club, error) {
	res, err := c.clubRepository.GetById(id)
	if err != nil {
		logService.Error().Err(err).Msg("failed to get club by id")
		return model.Club{}, err
	}
	return res, nil
}

// Update implements club.ServiceInterface.
func (c *clubService) Update(input model.Club, id string) (model.Club, error) {
	err := c.clubRepository.Update(input, id)
	if err != nil {
		logService.Error().Err(err).Msg("failed to update club club")
		return model.Club{}, err
	}
	updatedClub, err := c.clubRepository.GetById(id)
	if err != nil {
		logService.Error().Err(err).Msg("failed to retrieve updated club")
		return model.Club{}, err
	}
	return updatedClub, nil
}

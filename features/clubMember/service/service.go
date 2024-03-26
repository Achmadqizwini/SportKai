package service

import (
	"github.com/Achmadqizwini/SportKai/features/clubMember/model"
	repo "github.com/Achmadqizwini/SportKai/features/clubMember/repository"
	"github.com/Achmadqizwini/SportKai/utils/logger"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	Create(input model.MemberPayload) error
	Get() ([]model.MemberPayload, error)
	GetById(id string) (model.MemberPayload, error)
	Update(input model.ClubMember, id string) (model.MemberPayload, error)
	Delete(id string) error
}

type memberService struct {
	memberRepository repo.RepositoryInterface
	validate         *validator.Validate
}

func New(repo repo.RepositoryInterface) ServiceInterface {
	return &memberService{
		memberRepository: repo,
		validate:         validator.New(),
	}
}

var logService = logger.NewLogger().Logger.With().Logger()

// Create implements ServiceInterface.
func (c *memberService) Create(input model.MemberPayload) error {
	input.PublicId = uuid.NewString()
	if input.Status == "" {
		input.Status = "Requested"
	}
	err := c.memberRepository.Create(input)
	if err != nil {
		logService.Error().Err(err).Msg("failed to create new club member")
		return err
	}
	return nil
}

// Delete implements ServiceInterface.
func (c *memberService) Delete(id string) error {
	err := c.memberRepository.Delete(id)
	if err != nil {
		logService.Error().Err(err).Msg("failed to kick member")
		return err
	}
	return nil
}

// Get implements ServiceInterface.
func (c *memberService) Get() ([]model.MemberPayload, error) {
	member, err := c.memberRepository.Get()
	if err != nil {
		logService.Error().Err(err).Msg("failed to retrieve club member")
		return nil, err
	}
	return member, nil
}

// GetById implements ServiceInterface.
func (c *memberService) GetById(id string) (model.MemberPayload, error) {
	res, err := c.memberRepository.GetById(id)
	if err != nil {
		logService.Error().Err(err).Msg("failed to get member by id")
		return model.MemberPayload{}, err
	}
	return res, nil
}

// Update implements ServiceInterface.
func (c *memberService) Update(input model.ClubMember, id string) (model.MemberPayload, error) {
	err := c.memberRepository.Update(input, id)
	if err != nil {
		logService.Error().Err(err).Msg("failed to update club member")
		return model.MemberPayload{}, err
	}
	updatedMember, err := c.memberRepository.GetById(id)
	if err != nil {
		logService.Error().Err(err).Msg("failed to retrieve updated member")
		return model.MemberPayload{}, err
	}
	return updatedMember, nil
}

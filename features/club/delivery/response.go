package delivery

import (
	"github.com/Achmadqizwini/SportKai/features/club/model"
	"time"
)

type ClubResponse struct {
	PublicId     string    `json:"id"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	City         string    `json:"city"`
	Description  string    `json:"description"`
	JoinedMember int       `json:"joined_member"`
	MemberTotal  int       `json:"member_total"`
	Rules        string    `json:"rules"`
	Requirements string    `json:"requirements"`
	CreatedAt    time.Time `json:"created_at"`
}

func getClubResponse(res model.Club) ClubResponse {
	return ClubResponse{
		PublicId:     res.PublicId,
		Name:         res.Name,
		Address:      res.Address,
		City:         res.City,
		Description:  res.Description,
		JoinedMember: res.JoinedMember,
		MemberTotal:  res.MemberTotal,
		Rules:        res.Rules,
		Requirements: res.Requirements,
		CreatedAt:    res.CreatedAt,
	}
}

func getClubResponseList(dataCore []model.Club) []ClubResponse {
	var dataResponse []ClubResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, getClubResponse(v))
	}
	return dataResponse
}

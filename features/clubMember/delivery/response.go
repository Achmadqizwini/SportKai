package delivery

import (
	"github.com/Achmadqizwini/SportKai/features/clubMember/model"
	"time"
)

type MemberResponse struct {
	PublicId string    `json:"id"`
	User_id  uint      `json:"user_id"`
	Club_id  uint      `json:"club_id"`
	Status   string    `json:"status"`
	JoinedAt time.Time `json:"joined_at"`
	LeftAt   time.Time `json:"left_at"`
}

func getMemberResponse(res model.ClubMember) MemberResponse {
	return MemberResponse{
		PublicId: res.PublicId,
		User_id:  res.UserId,
		Club_id:  res.ClubId,
		Status:   res.Status,
		JoinedAt: res.JoinedAt,
		LeftAt:   res.LeftAt,
	}
}

func getMemberResponseList(dataCore []model.ClubMember) []MemberResponse {
	var dataResponse []MemberResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, getMemberResponse(v))
	}
	return dataResponse
}

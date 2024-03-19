package model

import (
	"time"
)

type ClubMember struct {
	ID       uint
	PublicId string
	UserId   int    `json:"user_id" form:"user_id"`
	ClubId   int    `json:"club_id" form:"club_id"`
	Status   string `json:"status" form:"status"`
	JoinedAt time.Time
	LeftAt   time.Time
}

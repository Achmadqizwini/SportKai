package model

import (
	"time"
)

type ClubMember struct {
	ID       uint
	PublicId string
	UserId   uint   `json:"user_id" form:"user_id"`
	ClubId   uint   `json:"club_id" form:"club_id"`
	Status   string `json:"status" form:"status"`
	JoinedAt *time.Time
	LeftAt   *time.Time
}

package model

import (
	"time"
)

type ClubMember struct {
	ID       uint
	PublicId string
	UserId   uint
	ClubId   uint
	Status   string
	JoinedAt *time.Time
	LeftAt   *time.Time
}

type MemberPayload struct {
	PublicId string
	UserId   string `json:"user_id" form:"user_id"`
	ClubId   string `json:"club_id" form:"club_id"`
	Status   string `json:"status" form:"status"`
}

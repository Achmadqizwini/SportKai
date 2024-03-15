package model

import (
	"time"
)

type ClubMember struct {
	ID       uint
	PublicId string
	UserId   string
	ClubId   string
	Status   string
	JoinedAt time.Time
	LeftAt   time.Time
}

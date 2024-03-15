package model

import (
	"time"
)

type Club struct {
	ID           uint
	PublicId     string
	Name         string `validate:"required" json:"name" form:"name"`
	Address      string `validate:"required" json:"address" form:"address"`
	City         string `validate:"required" json:"city" form:"city"`
	Description  string `json:"description" form:"description"`
	JoinedMember int    `json:"joined_member" form:"joined_member"`
	MemberTotal  int    `json:"member_total" form:"member_total"`
	Rules        string `validate:"required" json:"rules" form:"rules"`
	Requirements string `validate:"required" json:"requirements" form:"requirements"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

package user

import (
	"time"
)

type Core struct {
	ID              uint
	FullName        string `valiidate:"required"`
	Email           string `valiidate:"required,email"`
	Password        string `valiidate:"required"`
	Phone           string `valiidate:"required"`
	Gender          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ServiceInterface interface {
	Create(input Core) error
}

type RepositoryInterface interface {
	Create(input Core) error
}

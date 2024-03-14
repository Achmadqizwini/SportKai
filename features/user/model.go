package user

import (
	"time"
)

type User struct {
	ID              uint
	PublicId		string 
	FullName        string `validate:"required" json:"fullname" form:"fullname"`
	Email           string `validate:"required" json:"email" form:"email"`
	Password        string `validate:"required" json:"password" form:"password"`
	Phone           string `validate:"required" json:"phone" form:"phone"`
	Gender          string `validate:"required" json:"gender" form:"gender"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ServiceInterface interface {
	Create(input User) error
	Get() ([]User, error)
	GetById(id string) (User, error)
	Update(input User, id string) (User, error)
	Delete(id string) error 
}

type RepositoryInterface interface {
	Create(input User) error
	Get() ([]User, error)
	GetById(id string) (User, error)
	Update(input User, id string) (User, error)
	Delete(id string) error 
}

package model

type Auth struct {
	ID       uint
	PublicId string
	Email    string `validate:"required" json:"email" form:"email"`
	Password string `validate:"required" json:"password" form:"password"`
	FullName string
	Phone    string
	Gender   string
}

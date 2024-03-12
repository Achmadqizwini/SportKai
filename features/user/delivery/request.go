package delivery

import "github.com/Achmadqizwini/SportKai/features/user"

type InsertRequest struct {
	FullName        string `json:"full_name" form:"full_name"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	Phone           string `json:"phone" form:"phone"`
	Gender          string `json:"gender" form:"gender"`
}

type UpdateRequest struct {
	ID              uint   `json:"id" form:"id"`
	FullName        string `json:"full_name" form:"full_name"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	Phone           string `json:"phone" form:"phone"`
	Gender          string `json:"gender" form:"gender"`
}

func toCore(i interface{}) user.User {
	switch i.(type) {
	case InsertRequest:
		cnv := i.(InsertRequest)
		return user.User{
			FullName:        cnv.FullName,
			Email:           cnv.Email,
			Password:        cnv.Password,
			Phone:           cnv.Phone,
			Gender:          cnv.Gender,
		}

	case UpdateRequest:
		cnv := i.(UpdateRequest)
		return user.User{
			ID:              cnv.ID,
			FullName:        cnv.FullName,
			Email:           cnv.Email,
			Password:        cnv.Password,
			Phone:           cnv.Phone,
			Gender:          cnv.Gender,
		}
	}

	return user.User{}
}

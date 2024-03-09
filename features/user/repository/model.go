package repository

import (
	_user "github.com/Achmadqizwini/SportKai/features/user"

	"gorm.io/gorm"
)

// struct gorm model
type User struct {
	gorm.Model
	FullName        string `validate:"required"`
	Email           string `validate:"required,email,unique"`
	Password        string `validate:"required"`
	Phone           string `validate:"required"`
	Gender          string
}

func fromCore(dataCore _user.Core) User {
	userGorm := User{
		FullName:        dataCore.FullName,
		Email:           dataCore.Email,
		Password:        dataCore.Password,
		Phone:           dataCore.Phone,
		Gender:          dataCore.Gender,
	}
	return userGorm
}

func (dataModel *User) toCore() _user.Core {
	return _user.Core{
		ID:              dataModel.ID,
		FullName:        dataModel.FullName,
		Email:           dataModel.Email,
		Password:        dataModel.Password,
		Phone:           dataModel.Phone,
		Gender:          dataModel.Gender,
		CreatedAt:       dataModel.CreatedAt,
		UpdatedAt:       dataModel.UpdatedAt,
	}
}

func toCoreList(dataModel []User) []_user.Core {
	var dataCore []_user.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}

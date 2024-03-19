package delivery

import "github.com/Achmadqizwini/SportKai/features/user/model"

type UserResponse struct {
	PublicId string `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
}

func getUserResponse(dataCore model.User) UserResponse {
	return UserResponse{
		PublicId: dataCore.PublicId,
		FullName: dataCore.FullName,
		Email:    dataCore.Email,
		Phone:    dataCore.Phone,
		Gender:   dataCore.Gender,
	}
}

func getUserResponseList(dataCore []model.User) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, getUserResponse(v))
	}
	return dataResponse
}

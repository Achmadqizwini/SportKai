package delivery

import "github.com/Achmadqizwini/SportKai/features/user"

type UserResponse struct {
	ID              uint   `json:"id"`
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Gender          string `json:"gender"`
}

func fromCore(dataCore user.User) UserResponse {
	return UserResponse{
		ID:              dataCore.ID,
		FullName:        dataCore.FullName,
		Email:           dataCore.Email,
		Phone:           dataCore.Phone,
		Gender:          dataCore.Gender,
	}
}

func fromCoreList(dataCore []user.User) []UserResponse {
	var dataResponse []UserResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}

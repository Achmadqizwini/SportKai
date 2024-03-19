package delivery

import "github.com/Achmadqizwini/SportKai/features/auth/model"

type login struct {
	ID       string `json:"id"`
	FullName string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

func UserResponse(dataCore model.Auth, token string) login {
	return login{
		ID:       dataCore.PublicId,
		FullName: dataCore.FullName,
		Email:    dataCore.Email,
		Token:    token,
	}
}

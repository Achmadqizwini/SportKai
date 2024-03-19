package delivery

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Achmadqizwini/SportKai/features/auth/model"
	svc "github.com/Achmadqizwini/SportKai/features/auth/service"
	"github.com/Achmadqizwini/SportKai/utils/helper"
)

type AuthDelivery struct {
	authService svc.ServiceInterface
}

func New(service svc.ServiceInterface, r *http.ServeMux) {
	handler := &AuthDelivery{
		authService: service,
	}

	r.HandleFunc("POST /login", handler.Login)
	// r.HandleFunc("GET /auths", handler.Getauth)
	// r.HandleFunc("PUT /auths/{id}", handler.Updateauth)
	// r.HandleFunc("DELETE /auths/{id}", handler.Deleteauth)
	// r.HandleFunc("GET /auths/{id}", handler.GetauthById)

}

func (delivery *AuthDelivery) Login(w http.ResponseWriter, r *http.Request) {
	var authInput model.Auth
	var err error
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		err = json.NewDecoder(r.Body).Decode(&authInput)
	} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") || strings.HasPrefix(contentType, "multipart/form-data") {
		err = r.ParseForm()
		if err == nil {
			authInput.Email = r.Form.Get("email")
			authInput.Password = r.Form.Get("password")
		}
	} else {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(helper.FailedResponse("Unsupported content type"))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.FailedResponse("Error binding data " + err.Error()))
		return
	}

	user, token, err := delivery.authService.Login(authInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("Failed to login: " + err.Error()))
		return
	}
	res := UserResponse(user, token)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessWithDataResponse("Login Success", res))
}

package delivery

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Achmadqizwini/SportKai/features/user"
	"github.com/Achmadqizwini/SportKai/utils/helper"
)

type UserDelivery struct {
	userService user.ServiceInterface
}

func New(service user.ServiceInterface, r *http.ServeMux) {
	handler := &UserDelivery{
		userService: service,
	}

	r.HandleFunc("POST /users", handler.CreateUser)
	r.HandleFunc("GET /users", handler.GetUsers)
	r.HandleFunc("PUT /users", handler.UpdateUser)
	r.HandleFunc("DELETE /users", handler.DeleteUser)

}

func (delivery *UserDelivery) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userInput user.User
	var err error
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		err = json.NewDecoder(r.Body).Decode(&userInput)
	} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") || strings.HasPrefix(contentType, "multipart/form-data"){
		err = r.ParseForm()
		if err == nil {
			userInput.FullName = r.Form.Get("fullname")
			userInput.Email = r.Form.Get("email")
			userInput.Password = r.Form.Get("password")
			userInput.Phone = r.Form.Get("phone")
			userInput.Gender = r.Form.Get("gender")
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

	err = delivery.userService.Create(userInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("Failed to insert data " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessResponse("Success create new users"))
}

func (delivery *UserDelivery) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := delivery.userService.Get()
	data := getUserResponseList(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("Failed to retrieve data " + err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(helper.SuccessWithDataResponse("Success retrieve users", data))
	}
}

func (delivery *UserDelivery) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var userInput user.User
	var err error
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		err = json.NewDecoder(r.Body).Decode(&userInput)
	} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") || strings.HasPrefix(contentType, "multipart/form-data"){
		err = r.ParseForm()
		if err == nil {
			userInput.FullName = r.Form.Get("fullname")
			userInput.Email = r.Form.Get("email")
			userInput.Password = r.Form.Get("password")
			userInput.Phone = r.Form.Get("phone")
			userInput.Gender = r.Form.Get("gender")
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
	result, err := delivery.userService.Update(userInput, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("Failed to update data " + err.Error()))
		return
	}
	updatedUser := getUserResponse(result)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessWithDataResponse("Success update users", updatedUser))
}

func (delivery *UserDelivery) DeleteUser(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	err := delivery.userService.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("Failed to delete data " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessResponse("Success delete users"))
}
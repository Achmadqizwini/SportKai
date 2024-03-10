package delivery

import (
	"encoding/json"
	"net/http"

	"github.com/Achmadqizwini/SportKai/features/user"
	"github.com/Achmadqizwini/SportKai/utils/helper"
	"github.com/gorilla/mux"
)

type UserDelivery struct {
	userService user.ServiceInterface
}

func New(service user.ServiceInterface, r *mux.Router) {
	handler := &UserDelivery{
		userService: service,
	}

	r.HandleFunc("/users", handler.Create).Methods("POST")
}

func (delivery *UserDelivery) Create(w http.ResponseWriter, r *http.Request) {
	var userInput user.Core
	err := json.NewDecoder(r.Body).Decode(&userInput)
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

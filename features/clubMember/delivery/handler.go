package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Achmadqizwini/SportKai/features/clubMember/model"
	svc "github.com/Achmadqizwini/SportKai/features/clubMember/service"
	"github.com/Achmadqizwini/SportKai/utils/helper"
)

type ClubDelivery struct {
	clubService svc.ServiceInterface
}

func New(service svc.ServiceInterface, r *http.ServeMux) {
	handler := &ClubDelivery{
		clubService: service,
	}

	r.HandleFunc("POST /members", handler.CreateMember)
	// r.HandleFunc("GET /members", handler.GetClub)
	// r.HandleFunc("PUT /members/{id}", handler.UpdateClub)
	// r.HandleFunc("DELETE /members/{id}", handler.DeleteClub)
	// r.HandleFunc("GET /members/{id}", handler.GetClubById)

}

func (delivery *ClubDelivery) CreateMember(w http.ResponseWriter, r *http.Request) {
	var clubInput model.ClubMember
	var err error
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		err = json.NewDecoder(r.Body).Decode(&clubInput)
	} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") || strings.HasPrefix(contentType, "multipart/form-data") {
		err = r.ParseForm()
		if err == nil {
			clubInput.UserId, _ = strconv.Atoi(r.Form.Get("user_id"))
			clubInput.ClubId, _ = strconv.Atoi(r.Form.Get("club_id"))
			clubInput.Status = r.Form.Get("status")
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

	fmt.Println(clubInput)

	err = delivery.clubService.Create(clubInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("Failed to insert data: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessResponse("Success create new club member"))
}

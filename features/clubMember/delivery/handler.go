package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Achmadqizwini/SportKai/features/clubMember/model"
	svc "github.com/Achmadqizwini/SportKai/features/clubMember/service"
	"github.com/Achmadqizwini/SportKai/utils/helper"
)

type MemberDelivery struct {
	memberService svc.ServiceInterface
}

func New(service svc.ServiceInterface, r *http.ServeMux) {
	handler := &MemberDelivery{
		memberService: service,
	}

	r.HandleFunc("POST /members", handler.CreateMember)
	r.HandleFunc("GET /members", handler.GetMember)
	// r.HandleFunc("PUT /members/{id}", handler.UpdateClub)
	// r.HandleFunc("DELETE /members/{id}", handler.DeleteClub)
	// r.HandleFunc("GET /members/{id}", handler.GetClubById)

}

func (delivery *MemberDelivery) CreateMember(w http.ResponseWriter, r *http.Request) {
	var clubInput model.ClubMember
	var err error
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		err = json.NewDecoder(r.Body).Decode(&clubInput)
	} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") || strings.HasPrefix(contentType, "multipart/form-data") {
		err = r.ParseForm()
		if err == nil {
			id, _ := strconv.Atoi(r.Form.Get("user_id"))
			clubInput.UserId = uint(id)
			id, _ = strconv.Atoi(r.Form.Get("club_id"))
			clubInput.ClubId = uint(id)
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

	err = delivery.memberService.Create(clubInput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("Failed to insert data: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessResponse("Success create new club member"))
}

func (delivery *MemberDelivery) GetMember(w http.ResponseWriter, r *http.Request) {
	members, err := delivery.memberService.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("failed to retrieve club members"))
		return
	}
	response := getMemberResponseList(members)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessWithDataResponse("success retrieve club members", response))
}

package delivery

import (
	"encoding/json"
	"net/http"
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
	r.HandleFunc("PUT /members/{id}", handler.UpdateMember)
	r.HandleFunc("DELETE /members/{id}", handler.DeleteClub)
	r.HandleFunc("GET /members/{id}", handler.GetMemberById)

}

func (delivery *MemberDelivery) CreateMember(w http.ResponseWriter, r *http.Request) {
	var memberInput model.MemberPayload
	var err error
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		err = json.NewDecoder(r.Body).Decode(&memberInput)
	} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") || strings.HasPrefix(contentType, "multipart/form-data") {
		err = r.ParseForm()
		if err == nil {
			memberInput.UserId = r.Form.Get("user_id")
			memberInput.ClubId = r.Form.Get("club_id")
			memberInput.Status = r.Form.Get("status")
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

	err = delivery.memberService.Create(memberInput)
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

func (delivery *MemberDelivery) UpdateMember(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var memberInput model.ClubMember
	var err error
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		err = json.NewDecoder(r.Body).Decode(&memberInput)
	} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") || strings.HasPrefix(contentType, "multipart/form-data") {
		err = r.ParseForm()
		if err == nil {
			memberInput.Status = r.Form.Get("status")

		}
	} else {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode(helper.FailedResponse("Unsupported content type"))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.FailedResponse("Error binding data. " + err.Error()))
		return
	}
	result, err := delivery.memberService.Update(memberInput, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("Failed to update data: " + err.Error()))
		return
	}
	updatedUser := getMemberResponse(result)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessWithDataResponse("Success update users", updatedUser))
}

func (delivery *MemberDelivery) GetMemberById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	res, err := delivery.memberService.GetById(id)
	memberData := getMemberResponse(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("Failed to retrieve member: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessWithDataResponse("Success retrieve member", memberData))
}

func (delivery *MemberDelivery) DeleteClub(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := delivery.memberService.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("Failed to remove member: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessResponse("Success remove member"))
}

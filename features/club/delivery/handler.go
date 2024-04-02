package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Achmadqizwini/SportKai/features/club/model"
	svc "github.com/Achmadqizwini/SportKai/features/club/service"
	"github.com/Achmadqizwini/SportKai/middlewares"
	"github.com/Achmadqizwini/SportKai/utils/helper"
)

type ClubDelivery struct {
	clubService svc.ServiceInterface
}

func New(service svc.ServiceInterface, r *http.ServeMux) {
	handler := &ClubDelivery{
		clubService: service,
	}

	r.HandleFunc("POST /clubs", middlewares.JWTMiddleware(handler.CreateClub))
	r.HandleFunc("GET /clubs", handler.GetClub)
	r.HandleFunc("PUT /clubs/{id}", handler.UpdateClub)
	r.HandleFunc("DELETE /clubs/{id}", handler.DeleteClub)
	r.HandleFunc("GET /clubs/{id}", handler.GetClubById)

}

func (c *ClubDelivery) GetClub(w http.ResponseWriter, r *http.Request) {
	clubs, err := c.clubService.Get()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("failed to retrieve clubs"))
		return
	}
	response := getClubResponseList(clubs)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessWithDataResponse("success retrieve clubs", response))
}

func (delivery *ClubDelivery) CreateClub(w http.ResponseWriter, r *http.Request) {
	var clubInput model.Club
	var err error
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		err = json.NewDecoder(r.Body).Decode(&clubInput)
	} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") || strings.HasPrefix(contentType, "multipart/form-data") {
		err = r.ParseForm()
		if err == nil {
			clubInput.Name = r.Form.Get("name")
			clubInput.Address = r.Form.Get("address")
			clubInput.City = r.Form.Get("city")
			clubInput.Description = r.Form.Get("description")
			clubInput.MemberTotal, _ = strconv.Atoi(r.Form.Get("member_total"))
			clubInput.Rules = r.Form.Get("rules")
			clubInput.Requirements = r.Form.Get("requirements")
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
	ctx := r.Context()
	user_id := ctx.Value(middlewares.Val("user_id")).(string)
	err = delivery.clubService.Create(clubInput, user_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("Failed to insert data: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessResponse("Success create new clubs"))
}

func (delivery *ClubDelivery) GetClubById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	res, err := delivery.clubService.GetById(id)
	clubData := getClubResponse(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("Failed to retrieve club: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessWithDataResponse("Success retrieve club", clubData))
}

func (delivery *ClubDelivery) DeleteClub(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	err := delivery.clubService.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("Failed to delete club: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessResponse("Success delete club"))
}

func (delivery *ClubDelivery) UpdateClub(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var clubInput model.Club
	var err error
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		err = json.NewDecoder(r.Body).Decode(&clubInput)
	} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") || strings.HasPrefix(contentType, "multipart/form-data") {
		err = r.ParseForm()
		if err == nil {
			clubInput.Name = r.Form.Get("name")
			clubInput.Address = r.Form.Get("address")
			clubInput.City = r.Form.Get("city")
			clubInput.Description = r.Form.Get("description")
			clubInput.Rules = r.Form.Get("rules")
			clubInput.Requirements = r.Form.Get("requirements")

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
	result, err := delivery.clubService.Update(clubInput, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.FailedResponse("Failed to update data: " + err.Error()))
		return
	}
	updatedUser := getClubResponse(result)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.SuccessWithDataResponse("Success update club", updatedUser))
}

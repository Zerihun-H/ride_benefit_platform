package rest

import (
	"encoding/json"
	"net/http"
	"rideBenefit/internal/constant/model"
	partner "rideBenefit/internal/module/partner"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

// PartnerHandler contains the function of handler for domain Partner
type PartnerHandler interface {
	GetPartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddPartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdatePartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeletePartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type partnerHandler struct {
	PartnerCase partner.Usecase
}

// PartnerInit is to initialize the rest handler for domain Partner
func PartnerInit(PartnerCase partner.Usecase) PartnerHandler {
	return &partnerHandler{
		PartnerCase,
	}
}

func (dh *partnerHandler) GetPartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check if the diver ID param is valid
	partnerID := ps.ByName("partnerID")
	// Convert the partnerID string to uint64
	id, err := strconv.Atoi(partnerID)
	if err != nil {
		http.Error(w, model.ErrInvalidPartnerID.Error(), http.StatusBadRequest)
		return
	}

	partner, err := dh.PartnerCase.GetPartner(uint64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, gorm.ErrRecordNotFound.Error(), http.StatusNotFound)
			return
		} else {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(partner)
}

func (dh *partnerHandler) AddPartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse partner data
	partner := &model.Partner{}
	err := json.NewDecoder(r.Body).Decode(&partner)
	if err != nil {
		http.Error(w, model.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}
	part, err := dh.PartnerCase.AddPartner(partner)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(part)
}
func (dh *partnerHandler) UpdatePartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse partner data
	partner := &model.Partner{}
	err := json.NewDecoder(r.Body).Decode(&partner)
	if err != nil {
		http.Error(w, model.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}

	drv, err := dh.PartnerCase.UpdatePartner(partner)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drv)
}

func (dh *partnerHandler) DeletePartner(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check if the diver ID param is valid
	partnerID := ps.ByName("partnerID")
	// Convert the partnerID string to uint64
	id, err := strconv.Atoi(partnerID)
	if err != nil {
		http.Error(w, model.ErrInvalidPartnerID.Error(), http.StatusBadRequest)
	}
	err = dh.PartnerCase.DeletePartner(uint64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, model.ErrInvalidPartnerID.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

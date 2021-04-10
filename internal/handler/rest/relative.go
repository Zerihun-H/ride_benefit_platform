package rest

import (
	"encoding/json"
	"net/http"
	"rideBenefit/internal/constant"
	"rideBenefit/internal/constant/model"
	relative "rideBenefit/internal/module/relative"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

// RelativeHandler contains the function of handler for domain Relative
type RelativeHandler interface {
	GetRelative(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddRelative(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdateRelative(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteRelative(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type relativeHandler struct {
	RelativeCase relative.Usecase
}

// RelativeInit is to initialize the rest handler for domain Relative
func RelativeInit(RelativeCase relative.Usecase) RelativeHandler {
	return &relativeHandler{
		RelativeCase,
	}
}

func (dh *relativeHandler) GetRelative(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check if the diver ID param is valid
	relativeID := ps.ByName("relativeID")
	// Convert the relativeID string to uint64
	id, err := strconv.Atoi(relativeID)
	if err != nil {
		http.Error(w, constant.ErrInvalidRelativeID.Error(), http.StatusBadRequest)
		return
	}

	relative, err := dh.RelativeCase.GetRelative(uint64(id))
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
	json.NewEncoder(w).Encode(relative)
}

func (dh *relativeHandler) AddRelative(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse relative data
	relative := &model.Relative{}
	err := json.NewDecoder(r.Body).Decode(&relative)
	if err != nil {
		http.Error(w, constant.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}
	part, err := dh.RelativeCase.AddRelative(relative)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(part)
}
func (dh *relativeHandler) UpdateRelative(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse relative data
	relative := &model.Relative{}
	err := json.NewDecoder(r.Body).Decode(&relative)
	if err != nil {
		http.Error(w, constant.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}

	drv, err := dh.RelativeCase.UpdateRelative(relative)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drv)
}

func (dh *relativeHandler) DeleteRelative(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check if the diver ID param is valid
	relativeID := ps.ByName("relativeID")
	// Convert the relativeID string to uint64
	id, err := strconv.Atoi(relativeID)
	if err != nil {
		http.Error(w, constant.ErrInvalidRelativeID.Error(), http.StatusBadRequest)
	}
	err = dh.RelativeCase.DeleteRelative(uint64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, constant.ErrInvalidRelativeID.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

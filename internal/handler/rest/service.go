package rest

import (
	"encoding/json"
	"net/http"
	"rideBenefit/internal/constant"
	"rideBenefit/internal/constant/model"
	service "rideBenefit/internal/module/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

// ServiceHandler contains the function of handler for domain Service
type ServiceHandler interface {
	GetService(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddService(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdateService(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteService(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type serviceHandler struct {
	ServiceCase service.Usecase
}

// ServiceInit is to initialize the rest handler for domain Service
func ServiceInit(ServiceCase service.Usecase) ServiceHandler {
	return &serviceHandler{
		ServiceCase,
	}
}

func (dh *serviceHandler) GetService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check if the service ID param is valid
	serviceID := ps.ByName("serviceID")
	// Convert the serviceID string to uint64
	id, err := strconv.Atoi(serviceID)
	if err != nil {
		http.Error(w, constant.ErrInvalidServiceID.Error(), http.StatusBadRequest)
		return
	}

	service, err := dh.ServiceCase.GetService(uint64(id))
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
	json.NewEncoder(w).Encode(service)
}

func (dh *serviceHandler) AddService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse service data
	service := &model.Service{}
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		http.Error(w, constant.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}
	part, err := dh.ServiceCase.AddService(service)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(part)
}

func (dh *serviceHandler) UpdateService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse service data
	service := &model.Service{}
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		http.Error(w, constant.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}

	prt, err := dh.ServiceCase.UpdateService(service)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prt)
}

func (dh *serviceHandler) DeleteService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check if the service ID param is valid
	serviceID := ps.ByName("serviceID")
	// Convert the serviceID string to uint64
	id, err := strconv.Atoi(serviceID)
	if err != nil {
		http.Error(w, constant.ErrInvalidServiceID.Error(), http.StatusBadRequest)
	}
	err = dh.ServiceCase.DeleteService(uint64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, constant.ErrInvalidServiceID.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

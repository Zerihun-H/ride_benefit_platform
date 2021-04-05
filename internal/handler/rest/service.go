package rest

import (
	"net/http"
	service "rideBenefit/internal/module/service"

	"github.com/julienschmidt/httprouter"
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
}

func (dh *serviceHandler) AddService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (dh *serviceHandler) UpdateService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (dh *serviceHandler) DeleteService(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

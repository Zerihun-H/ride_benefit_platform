package rest

import (
	"net/http"
	Driver "rideBenefit/internal/module/driver"

	"github.com/julienschmidt/httprouter"
)

// DriverHandler contains the function of handler for domain Driver
type DriverHandler interface {
	GetDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdateDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type driverHandler struct {
	DriverCase Driver.Usecase
}

// DriverInit is to initialize the rest handler for domain Driver
func DriverInit(DriverCase Driver.Usecase) DriverHandler {
	return &driverHandler{
		DriverCase,
	}
}

func (dh *driverHandler) GetDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (dh *driverHandler) AddDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (dh *driverHandler) UpdateDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (dh *driverHandler) DeleteDriver(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

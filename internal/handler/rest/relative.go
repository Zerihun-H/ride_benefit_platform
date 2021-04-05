package rest

import (
	"net/http"
	relative "rideBenefit/internal/module/relative"

	"github.com/julienschmidt/httprouter"
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
}

func (dh *relativeHandler) AddRelative(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (dh *relativeHandler) UpdateRelative(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (dh *relativeHandler) DeleteRelative(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

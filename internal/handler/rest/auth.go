package rest

import (
	"net/http"
	auth "rideBenefit/internal/module/auth"

	"github.com/julienschmidt/httprouter"
)

// AuthHandler contains the function of handler for domain Auth
type AuthHandler interface {
	GetAuth(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddAuth(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdateAuth(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteAuth(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type authHandler struct {
	AuthCase auth.Usecase
}

// AuthInit is to initialize the rest handler for domain Auth
func AuthInit(AuthCase auth.Usecase) AuthHandler {
	return &authHandler{
		AuthCase,
	}
}

func (dh *authHandler) GetAuth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (dh *authHandler) AddAuth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (dh *authHandler) UpdateAuth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (dh *authHandler) DeleteAuth(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

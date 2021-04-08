package rest

import (
	"encoding/json"
	"net/http"
	"rideBenefit/internal/constant/model"
	"rideBenefit/internal/module/auth"

	"github.com/julienschmidt/httprouter"
)

// AuthHandler contains the function of handler for domain Auth
type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	RefreshAccessToken(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
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

func (ah *authHandler) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	login := &model.LoginModel{}

	err := json.NewDecoder(r.Body).Decode(login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validUser, accessToken, err := ah.AuthCase.Login(login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !validUser {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(accessToken))
	w.WriteHeader(http.StatusOK)
}

func (ah *authHandler) RefreshAccessToken(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.WriteHeader(http.StatusOK)
}

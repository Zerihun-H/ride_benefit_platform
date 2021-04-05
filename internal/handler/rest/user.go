package rest

import (
	"net/http"
	user "rideBenefit/internal/module/user"

	"github.com/julienschmidt/httprouter"
)

// UserHandler contains the function of handler for domain User
type UserHandler interface {
	GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type userHandler struct {
	UserCase user.Usecase
}

// UserInit is to initialize the rest handler for domain User
func UserInit(UserCase user.Usecase) UserHandler {
	return &userHandler{
		UserCase,
	}
}

func (dh *userHandler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (dh *userHandler) AddUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (dh *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (dh *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

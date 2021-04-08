package rest

import (
	"encoding/json"
	"net/http"
	"rideBenefit/internal/constant"
	"rideBenefit/internal/constant/model"
	user "rideBenefit/internal/module/user"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
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

func (uh *userHandler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check if the diver ID param is valid
	userID := ps.ByName("userID")
	// Convert the userID string to uint64
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, constant.ErrInvalidUserID.Error(), http.StatusBadRequest)
		return
	}

	user, err := uh.UserCase.GetUser(uint64(id))
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
	json.NewEncoder(w).Encode(user)
}

func (uh *userHandler) AddUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse user data
	user := &model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, constant.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}
	part, err := uh.UserCase.AddUser(user)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(part)
}
func (uh *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse user data
	user := &model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, constant.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}

	drv, err := uh.UserCase.UpdateUser(user)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drv)
}

func (uh *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check if the diver ID param is valid
	userID := ps.ByName("userID")
	// Convert the userID string to uint64
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, constant.ErrInvalidUserID.Error(), http.StatusBadRequest)
	}
	err = uh.UserCase.DeleteUser(uint64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, constant.ErrInvalidUserID.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

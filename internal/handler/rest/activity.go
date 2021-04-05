package rest

import (
	"net/http"
	activity "rideBenefit/internal/module/activity"

	"github.com/julienschmidt/httprouter"
)

// ActivityHandler contains the function of handler for domain Activity
type ActivityHandler interface {
	GetActivity(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddActivity(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdateActivity(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteActivity(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type activityHandler struct {
	ActivityCase activity.Usecase
}

// ActivityInit is to initialize the rest handler for domain Activity
func ActivityInit(ActivityCase activity.Usecase) ActivityHandler {
	return &activityHandler{
		ActivityCase,
	}
}

func (dh *activityHandler) GetActivity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (dh *activityHandler) AddActivity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (dh *activityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (dh *activityHandler) DeleteActivity(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

package rest

import (
	"net/http"

	"rideBenefit/internal/handler/rest"
	"rideBenefit/platform/httprouter"
)

// UserRouting returns the list of routers for domain user
func UserRouting(handler rest.UserHandler) []httprouter.Router {
	return []httprouter.Router{
		{
			Method:  http.MethodGet,
			Path:    "/user/:userID",
			Handler: handler.GetUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/user",
			Handler: handler.AddUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/user",
			Handler: handler.UpdateUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/user/:userID",
			Handler: handler.DeleteUser,
		},
	}
}

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
			Path:    "/users/:userID",
			Handler: handler.GetUser,
		},
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: handler.AddUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users",
			Handler: handler.UpdateUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/:userID",
			Handler: handler.DeleteUser,
		},
	}
}

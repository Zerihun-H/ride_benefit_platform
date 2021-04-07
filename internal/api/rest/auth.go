package rest

import (
	"net/http"

	"rideBenefit/internal/handler/rest"
	"rideBenefit/platform/httprouter"
)

// AuthRouting returns the list of routers for domain auth
func AuthRouting(handler rest.AuthHandler) []httprouter.Router {
	return []httprouter.Router{
		{ // login
			Method:  http.MethodPost,
			Path:    "/users/login",
			Handler: handler.Login,
		}, { // Update employee
			Method:  http.MethodPost,
			Path:    "auth/refreshAccessToken",
			Handler: handler.RefreshAccessToken,
		},
	}
}

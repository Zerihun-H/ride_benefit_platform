package rest

import (
	"net/http"

	"rideBenefit/internal/handler/rest"
	mw "rideBenefit/internal/handler/rest/middleware"
	"rideBenefit/platform/httprouter"

	"github.com/rileyr/middleware"
)

// AuthRouting returns the list of routers for domain auth
func AuthRouting(handler rest.AuthHandler) []httprouter.Router {
	return []httprouter.Router{
		{ // login
			Method:  http.MethodPost,
			Path:    "/users/login",
			Handler: handler.Login,
		},
		{ // refresh access token
			Method:      http.MethodPost,
			Path:        "/auth/refreshAccessToken",
			Handler:     handler.RefreshAccessToken,
			Middlewares: []middleware.Middleware{mw.ValidateAccessToken},
		}, { // refresh access token
			Method:  http.MethodGet,
			Path:    "/auth/role/:id/permissions",
			Handler: handler.RolePermissions,
			// Middlewares: []middleware.Middleware{mw.ValidateAccessToken},
		},
	}
}

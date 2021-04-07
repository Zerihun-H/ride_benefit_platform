package rest

import (
	"net/http"

	"rideBenefit/internal/handler/rest"
	"rideBenefit/platform/httprouter"
)

// ServiceRouting returns the list of routers for domain service
func ServiceRouting(handler rest.ServiceHandler) []httprouter.Router {
	return []httprouter.Router{
		{
			Method:  http.MethodGet,
			Path:    "/service/:serviceID",
			Handler: handler.GetService,
		},
		{
			Method:  http.MethodPost,
			Path:    "/service",
			Handler: handler.AddService,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/service",
			Handler: handler.UpdateService,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/service/:serviceID",
			Handler: handler.DeleteService,
		},
	}
}

package routing

import (
	"net/http"

	"rideBenefit/internal/handler/rest"
	"rideBenefit/platform/httprouter"
)

// RelativeRouting returns the list of routers for domain relative
func RelativeRouting(handler rest.RelativeHandler) []httprouter.Router {
	return []httprouter.Router{
		{
			Method:  http.MethodGet,
			Path:    "/relative/:relativeID",
			Handler: handler.GetRelative,
		},
		{
			Method:  http.MethodPost,
			Path:    "/relative",
			Handler: handler.AddRelative,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/relative",
			Handler: handler.UpdateRelative,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/relative/:relativeID",
			Handler: handler.DeleteRelative,
		},
	}
}

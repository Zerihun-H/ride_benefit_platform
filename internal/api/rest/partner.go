package routing

import (
	"net/http"

	"rideBenefit/internal/handler/rest"
	"rideBenefit/platform/httprouter"
)

// PartnerRouting returns the list of routers for domain partner
func PartnerRouting(handler rest.PartnerHandler) []httprouter.Router {
	return []httprouter.Router{
		{
			Method:  http.MethodGet,
			Path:    "/partner/:partnerID",
			Handler: handler.GetPartner,
		},
		{
			Method:  http.MethodPost,
			Path:    "/partner",
			Handler: handler.AddPartner,
		},
	}
}

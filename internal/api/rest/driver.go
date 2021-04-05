package routing

import (
	"net/http"

	"rideBenefit/internal/handler/rest"
	"rideBenefit/platform/httprouter"
)

// DriverRouting returns the list of routers for domain driver
func DriverRouting(handler rest.DriverHandler) []httprouter.Router {
	return []httprouter.Router{
		{
			Method:  http.MethodGet,
			Path:    "/driver/:driverID",
			Handler: handler.GetDriver,
		},
		{
			Method:  http.MethodPost,
			Path:    "/driver",
			Handler: handler.AddDriver,
		},
	}
}

package rest

import (
	"net/http"

	"rideBenefit/internal/handler/rest"
	"rideBenefit/platform/httprouter"
)

// ReportRouting returns the list of routers for domain report
func ReportRouting(handler rest.ReportHandler) []httprouter.Router {
	return []httprouter.Router{
		{
			Method:  http.MethodGet,
			Path:    "/report/:reportID",
			Handler: handler.GetReport,
		},
		{
			Method:  http.MethodPost,
			Path:    "/report",
			Handler: handler.AddReport,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/report",
			Handler: handler.UpdateReport,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/report/:reportID",
			Handler: handler.DeleteReport,
		},
	}
}

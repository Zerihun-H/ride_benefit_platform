package rest

import (
	"net/http"
	report "rideBenefit/internal/module/report"

	"github.com/julienschmidt/httprouter"
)

// ReportHandler contains the function of handler for domain Report
type ReportHandler interface {
	GetReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	AddReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	UpdateReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	DeleteReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type reportHandler struct {
	ReportCase report.Usecase
}

// ReportInit is to initialize the rest handler for domain Report
func ReportInit(ReportCase report.Usecase) ReportHandler {
	return &reportHandler{
		ReportCase,
	}
}

func (dh *reportHandler) GetReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (dh *reportHandler) AddReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (dh *reportHandler) UpdateReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
func (dh *reportHandler) DeleteReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

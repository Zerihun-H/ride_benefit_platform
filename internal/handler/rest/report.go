package rest

import (
	"encoding/json"
	"net/http"
	"rideBenefit/internal/constant/model"
	report "rideBenefit/internal/module/report"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
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
	// Check if the diver ID param is valid
	reportID := ps.ByName("reportID")
	// Convert the reportID string to uint64
	id, err := strconv.Atoi(reportID)
	if err != nil {
		http.Error(w, model.ErrInvalidReportID.Error(), http.StatusBadRequest)
		return
	}

	report, err := dh.ReportCase.GetReport(uint64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, gorm.ErrRecordNotFound.Error(), http.StatusNotFound)
			return
		} else {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func (dh *reportHandler) AddReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse report data
	report := &model.Report{}
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, model.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}
	part, err := dh.ReportCase.AddReport(report)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(part)
}
func (dh *reportHandler) UpdateReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Parse report data
	report := &model.Report{}
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, model.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}

	drv, err := dh.ReportCase.UpdateReport(report)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drv)
}

func (dh *reportHandler) DeleteReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Check if the diver ID param is valid
	reportID := ps.ByName("reportID")
	// Convert the reportID string to uint64
	id, err := strconv.Atoi(reportID)
	if err != nil {
		http.Error(w, model.ErrInvalidReportID.Error(), http.StatusBadRequest)
	}
	err = dh.ReportCase.DeleteReport(uint64(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, model.ErrInvalidReportID.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

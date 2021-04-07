package report

import (
	"rideBenefit/internal/constant/model"
	"rideBenefit/internal/repository"
	"rideBenefit/internal/storage/persistence"
)

// Usecase contains the function of business logic of domain report
type Usecase interface {
	GetReport(reportID uint64) (*model.Report, error)
	AddReport(report *model.Report) (*model.Report, error)
	UpdateReport(report *model.Report) (*model.Report, error)
	DeleteReport(reportID uint64) error
}

type service struct {
	reportRepo    repository.ReportRepository
	reportPersist persistence.ReportPersistence
}

// Initialize takes all necessary service for domain report to run the business logic of domain report
func Initialize(reportRepo repository.ReportRepository, reportPersist persistence.ReportPersistence) Usecase {
	return &service{
		reportRepo:    reportRepo,
		reportPersist: reportPersist,
	}
}

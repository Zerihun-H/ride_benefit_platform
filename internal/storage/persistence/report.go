package persistence

import (
	"rideBenefit/internal/constant/model"

	"gorm.io/gorm"
)

// ReportPersistence contains the list of functions for database table report
type ReportPersistence interface {
	GetReport(reportID uint64) (*model.Report, error)
	AddReport(report *model.Report) (*model.Report, error)
}

type reportPersistence struct {
	db *gorm.DB
}

// ReportInit is to init the report persistence that contains report data
func ReportInit(db *gorm.DB) ReportPersistence {
	return &reportPersistence{
		db,
	}
}

// GetReport using the report id fetchs the report from the report database
func (pp *reportPersistence) GetReport(reportID uint64) (*model.Report, error) {
	report := &model.Report{}
	if err := pp.db.Where("id = ?", reportID).First(report).Error; err != nil {

		return &model.Report{}, err
	}
	return report, nil
}

// AddReport is adds a report to the database given a valid report
func (pp *reportPersistence) AddReport(report *model.Report) (*model.Report, error) {
	if err := pp.db.Create(report).Error; err != nil {
		return nil, err
	}
	return report, nil
}

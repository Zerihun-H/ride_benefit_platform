package persistence

import (
	"rideBenefit/internal/constant/model"
	"rideBenefit/platform/cockroach"
)

// ReportPersistence contains the list of functions for database table reports
type ReportPersistence interface {
	GetReport(reportID uint64) (*model.Report, error)
	AddReport(report *model.Report) (*model.Report, error)
	UpdateReport(report *model.Report) (*model.Report, error)
	DeleteReport(reportID uint64) error
}

type reportPersistence struct {
	db cockroach.CockroachPlatform
}

// ReportInit is to init the report persistence that contains report data
func ReportInit(db cockroach.CockroachPlatform) ReportPersistence {
	return &reportPersistence{
		db,
	}
}

// Getreport using the report id fetchs the report from the report database
func (rp *reportPersistence) GetReport(reportID uint64) (*model.Report, error) {
	db, err := rp.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	report := &model.Report{}
	if err := db.Where("id = ?", reportID).First(report).Error; err != nil {

		return &model.Report{}, err
	}

	return report, nil
}

// AddReport is adds a report to the database given a valid diver
func (rp *reportPersistence) AddReport(report *model.Report) (*model.Report, error) {
	db, err := rp.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	if err := db.Create(report).Error; err != nil {
		return nil, err
	}

	return report, nil
}

// AddReport is adds a report to the database given a valid diver
func (rp *reportPersistence) UpdateReport(report *model.Report) (*model.Report, error) {
	db, err := rp.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	updatedReport := *report
	err = db.First(report).Error
	if err != nil {
		return nil, err
	}

	updatedReport.ID = report.ID
	err = db.Save(&updatedReport).Error
	if err != nil {
		return nil, err
	}
	return &updatedReport, nil
}

// AddReport is adds a report to the database given a valid diver
func (rp *reportPersistence) DeleteReport(reportID uint64) error {
	db, err := rp.db.Open()
	if err != nil {
		return err
	}
	dbc, err := db.DB()
	if err != nil {
		return err
	}
	defer dbc.Close()

	return db.Where("id = ?", reportID).Delete(&model.Report{}).Error
}

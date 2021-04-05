package report

import "rideBenefit/internal/constant/model"

func (s *service) GetReport(reportID uint64) (*model.Report, error) {
	// Some validation

	report, err := s.reportPersist.GetReport(reportID)
	if err != nil {
		return nil, err
	}

	return report, nil
}

func (s *service) AddReport(report *model.Report) (*model.Report, error) {
	// Check if the reported pruduct is valid and available
	// Check if the customer is valid
	// Check if the shop is valid
	drv, err := s.reportPersist.AddReport(report)
	if err != nil {
		return nil, err
	}

	return drv, nil
}

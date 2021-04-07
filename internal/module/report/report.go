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

	part, err := s.reportPersist.AddReport(report)
	if err != nil {
		return nil, err
	}

	return part, nil
}

func (s *service) UpdateReport(report *model.Report) (*model.Report, error) {

	part, err := s.reportPersist.UpdateReport(report)
	if err != nil {
		return nil, err
	}

	return part, nil
}

func (s *service) DeleteReport(reportID uint64) error {

	return s.reportPersist.DeleteReport(reportID)
}

package relative

import "rideBenefit/internal/constant/model"

func (s *service) GetRelative(relativeID uint64) (*model.Relative, error) {
	// Some validation

	relative, err := s.relativePersist.GetRelative(relativeID)
	if err != nil {
		return nil, err
	}

	return relative, nil
}

func (s *service) AddRelative(relative *model.Relative) (*model.Relative, error) {
	// Check if the relativeed pruduct is valid and available
	// Check if the customer is valid
	// Check if the shop is valid
	drv, err := s.relativePersist.AddRelative(relative)
	if err != nil {
		return nil, err
	}

	return drv, nil
}

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

	rel, err := s.relativePersist.AddRelative(relative)
	if err != nil {
		return nil, err
	}

	return rel, nil
}

func (s *service) UpdateRelative(relative *model.Relative) (*model.Relative, error) {

	rel, err := s.relativePersist.UpdateRelative(relative)
	if err != nil {
		return nil, err
	}

	return rel, nil
}

func (s *service) DeleteRelative(relativeID uint64) error {

	return s.relativePersist.DeleteRelative(relativeID)
}

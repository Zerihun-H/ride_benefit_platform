package partner

import "rideBenefit/internal/constant/model"

func (s *service) GetPartner(partnerID uint64) (*model.Partner, error) {
	// Some validation

	partner, err := s.partnerPersist.GetPartner(partnerID)
	if err != nil {
		return nil, err
	}

	return partner, nil
}

func (s *service) AddPartner(partner *model.Partner) (*model.Partner, error) {

	part, err := s.partnerPersist.AddPartner(partner)
	if err != nil {
		return nil, err
	}

	return part, nil
}

func (s *service) UpdatePartner(partner *model.Partner) (*model.Partner, error) {

	part, err := s.partnerPersist.UpdatePartner(partner)
	if err != nil {
		return nil, err
	}

	return part, nil
}

func (s *service) DeletePartner(partnerID uint64) error {

	return s.partnerPersist.DeletePartner(partnerID)
}

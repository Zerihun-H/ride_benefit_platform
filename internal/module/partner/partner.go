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
	// Check if the partnered pruduct is valid and available
	// Check if the customer is valid
	// Check if the shop is valid
	ordr, err := s.partnerPersist.AddPartner(partner)
	if err != nil {
		return nil, err
	}

	return ordr, nil
}

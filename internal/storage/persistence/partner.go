package persistence

import (
	"rideBenefit/internal/constant/model"
	"rideBenefit/platform/cockroach"
)

// PartnerPersistence contains the list of functions for database table partners
type PartnerPersistence interface {
	GetPartner(partnerID uint64) (*model.Partner, error)
	AddPartner(partner *model.Partner) (*model.Partner, error)
}

type partnerPersistence struct {
	db cockroach.CockroachPlatform
}

// PartnerInit is to init the partner persistence that contains partner data
func PartnerInit(db cockroach.CockroachPlatform) PartnerPersistence {
	return &partnerPersistence{
		db,
	}
}

// Getpartner using the partner id fetchs the partner from the partner database
func (pp *partnerPersistence) GetPartner(partnerID uint64) (*model.Partner, error) {
	partner := &model.Partner{}
	// if err := pp.db.Where("id = ?", partnerID).First(partner).Error; err != nil {

	// 	return &model.Partner{}, err
	// }
	return partner, nil
}

// AddPartner is adds a partner to the database given a valid diver
func (pp *partnerPersistence) AddPartner(partner *model.Partner) (*model.Partner, error) {
	// if err := pp.db.Create(partner).Error; err != nil {
	// 	return nil, err
	// }
	return partner, nil
}

package persistence

import (
	"rideBenefit/internal/constant/model"
	"rideBenefit/platform/cockroach"
)

// PartnerPersistence contains the list of functions for database table partners
type PartnerPersistence interface {
	GetPartner(partnerID uint64) (*model.Partner, error)
	AddPartner(partner *model.Partner) (*model.Partner, error)
	UpdatePartner(partner *model.Partner) (*model.Partner, error)
	DeletePartner(partnerID uint64) error
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
	db, err := pp.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	partner := &model.Partner{}
	if err := db.Where("id = ?", partnerID).First(partner).Error; err != nil {

		return &model.Partner{}, err
	}

	return partner, nil
}

// AddPartner is adds a partner to the database given a valid diver
func (pp *partnerPersistence) AddPartner(partner *model.Partner) (*model.Partner, error) {
	db, err := pp.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	if err := db.Create(partner).Error; err != nil {
		return nil, err
	}

	return partner, nil
}

// AddPartner is adds a partner to the database given a valid diver
func (pp *partnerPersistence) UpdatePartner(partner *model.Partner) (*model.Partner, error) {
	db, err := pp.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	updatedPartner := *partner
	err = db.First(partner).Error
	if err != nil {
		return nil, err
	}

	updatedPartner.ID = partner.ID
	err = db.Save(&updatedPartner).Error
	if err != nil {
		return nil, err
	}
	return &updatedPartner, nil
}

// AddPartner is adds a partner to the database given a valid diver
func (pp *partnerPersistence) DeletePartner(partnerID uint64) error {
	db, err := pp.db.Open()
	if err != nil {
		return err
	}
	dbc, err := db.DB()
	if err != nil {
		return err
	}
	defer dbc.Close()

	return db.Where("id = ?", partnerID).Delete(&model.Partner{}).Error
}

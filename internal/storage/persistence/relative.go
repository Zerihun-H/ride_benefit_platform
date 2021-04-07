package persistence

import (
	"rideBenefit/internal/constant/model"
	"rideBenefit/platform/cockroach"
)

// RelativePersistence contains the list of functions for database table relative
type RelativePersistence interface {
	GetRelative(relativeID uint64) (*model.Relative, error)
	AddRelative(relative *model.Relative) (*model.Relative, error)
	UpdateRelative(relative *model.Relative) (*model.Relative, error)
	DeleteRelative(relativeID uint64) error
}

type relativePersistence struct {
	db cockroach.CockroachPlatform
}

// RelativeInit is to init the relative persistence that contains relative data
func RelativeInit(db cockroach.CockroachPlatform) RelativePersistence {
	return &relativePersistence{
		db,
	}
}

// Getrelative using the relative id fetchs the relative from the relative database
func (pp *relativePersistence) GetRelative(relativeID uint64) (*model.Relative, error) {
	db, err := pp.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	relative := &model.Relative{}
	if err := db.Where("id = ?", relativeID).First(relative).Error; err != nil {

		return &model.Relative{}, err
	}

	return relative, nil
}

// AddRelative is adds a relative to the database given a valid diver
func (pp *relativePersistence) AddRelative(relative *model.Relative) (*model.Relative, error) {
	db, err := pp.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	if err := db.Create(relative).Error; err != nil {
		return nil, err
	}

	return relative, nil
}

// AddRelative is adds a relative to the database given a valid diver
func (pp *relativePersistence) UpdateRelative(relative *model.Relative) (*model.Relative, error) {
	db, err := pp.db.Open()
	if err != nil {
		return nil, err
	}
	dbc, err := db.DB()
	if err != nil {
		return nil, err
	}
	defer dbc.Close()
	updatedRelative := *relative
	err = db.First(relative).Error
	if err != nil {
		return nil, err
	}

	updatedRelative.ID = relative.ID
	err = db.Save(&updatedRelative).Error
	if err != nil {
		return nil, err
	}
	return &updatedRelative, nil
}

// AddRelative is adds a relative to the database given a valid diver
func (pp *relativePersistence) DeleteRelative(relativeID uint64) error {
	db, err := pp.db.Open()
	if err != nil {
		return err
	}
	dbc, err := db.DB()
	if err != nil {
		return err
	}
	defer dbc.Close()

	return db.Where("id = ?", relativeID).Delete(&model.Relative{}).Error
}

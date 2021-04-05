package persistence

import (
	"rideBenefit/internal/constant/model"

	"gorm.io/gorm"
)

// RelativePersistence contains the list of functions for database table relative
type RelativePersistence interface {
	GetRelative(relativeID uint64) (*model.Relative, error)
	AddRelative(relative *model.Relative) (*model.Relative, error)
}

type relativePersistence struct {
	db *gorm.DB
}

// RelativeInit is to init the relative persistence that contains relative data
func RelativeInit(db *gorm.DB) RelativePersistence {
	return &relativePersistence{
		db,
	}
}

// GetRelative using the relative id fetchs the relative from the relative database
func (pp *relativePersistence) GetRelative(relativeID uint64) (*model.Relative, error) {
	relative := &model.Relative{}
	if err := pp.db.Where("id = ?", relativeID).First(relative).Error; err != nil {

		return &model.Relative{}, err
	}
	return relative, nil
}

// AddRelative is adds a relative to the database given a valid relative
func (pp *relativePersistence) AddRelative(relative *model.Relative) (*model.Relative, error) {
	if err := pp.db.Create(relative).Error; err != nil {
		return nil, err
	}
	return relative, nil
}

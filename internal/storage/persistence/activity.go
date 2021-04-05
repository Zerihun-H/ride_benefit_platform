package persistence

import (
	"rideBenefit/internal/constant/model"

	"gorm.io/gorm"
)

// ActivityPersistence contains the list of functions for database table activitys
type ActivityPersistence interface {
	GetActivity(activityID uint64) (*model.Activity, error)
	AddActivity(activity *model.Activity) (*model.Activity, error)
}

type activityPersistence struct {
	db *gorm.DB
}

// ActivityInit is to init the activity persistence that contains activity data
func ActivityInit(db *gorm.DB) ActivityPersistence {
	return &activityPersistence{
		db,
	}
}

// Getactivity using the activity id fetchs the activity from the activity database
func (dp *activityPersistence) GetActivity(activityID uint64) (*model.Activity, error) {
	activity := &model.Activity{}
	if err := dp.db.Where("id = ?", activityID).First(activity).Error; err != nil {

		return &model.Activity{}, err
	}
	return activity, nil
}

// AddActivity is adds a activity to the database given a valid diver
func (dp *activityPersistence) AddActivity(activity *model.Activity) (*model.Activity, error) {
	if err := dp.db.Create(activity).Error; err != nil {
		return nil, err
	}
	return activity, nil
}

package activity

import "rideBenefit/internal/constant/model"

func (s *service) GetActivity(activityID uint64) (*model.Activity, error) {
	// Some validation

	activity, err := s.activityPersist.GetActivity(activityID)
	if err != nil {
		return nil, err
	}

	return activity, nil
}

func (s *service) AddActivity(activity *model.Activity) (*model.Activity, error) {
	// Check if the activityed pruduct is valid and available
	// Check if the customer is valid
	// Check if the shop is valid
	drv, err := s.activityPersist.AddActivity(activity)
	if err != nil {
		return nil, err
	}

	return drv, nil
}

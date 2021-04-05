package repository

// ActivityRepository contains the functions of data logic for domain activity
type ActivityRepository interface {
}

type activityRepository struct {
}

// ActivityInit initializes the data logic / repository for domain activity
func ActivityInit() ActivityRepository {
	return activityRepository{}
}

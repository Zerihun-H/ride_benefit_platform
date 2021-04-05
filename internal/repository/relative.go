package repository

// RelativeRepository contains the functions of data logic for domain relative
type RelativeRepository interface {
}

type relativeRepository struct {
}

// RelativeInit initializes the data logic / repository for domain relative
func RelativeInit() RelativeRepository {
	return relativeRepository{}
}

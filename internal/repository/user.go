package repository

// UserRepository contains the functions of data logic for domain user
type UserRepository interface {
}

type userRepository struct {
}

// PartnetInit initializes the data logic / repository for domain user
func UserInit() UserRepository {
	return userRepository{}
}

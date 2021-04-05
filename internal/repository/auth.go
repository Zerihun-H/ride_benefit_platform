package repository

// AuthRepository contains the functions of data logic for domain auth
type AuthRepository interface {
}

type authRepository struct {
}

// AuthInit initializes the data logic / repository for domain auth
func AuthInit() AuthRepository {
	return authRepository{}
}

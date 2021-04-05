package repository

// ServiceRepository contains the functions of data logic for domain service
type ServiceRepository interface {
}

type serviceRepository struct {
}

// ServiceInit initializes the data logic / repository for domain service
func Service() ServiceRepository {
	return serviceRepository{}
}

package repository

// DriverRepository contains the functions of data logic for domain driver
type DriverRepository interface {
}

type driverRepository struct {
}

// DriverInit initializes the data logic / repository for domain driver
func DriverInit() DriverRepository {
	return driverRepository{}
}

package repository

// PartnerRepository contains the functions of data logic for domain partner
type PartnerRepository interface {
}

type partnerRepository struct {
}

// PartnetInit initializes the data logic / repository for domain partner
func PartnerInit() PartnerRepository {
	return partnerRepository{}
}

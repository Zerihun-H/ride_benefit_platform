package model

import (
	"time"

	"gorm.io/gorm"
)

// Employee ...
type Employee struct {
	ID               uint64     `json:"id" gorm:"primarykey"`
	FirstName        string     `json:"firstName"`
	LastName         string     `json:"lastName"`
	Email            string     `json:"email"`
	PhoneNumber      string     `json:"phoneNumber"`
	Age              uint32     `json:"age"`
	Gender           string     `json:"gender"`
	BirthDate        *time.Time `json:"birthDate"`
	PhotoURL         string     `json:"photoURL"`
	IDPhotoURL       string     `json:"idPhotoURL"`
	Surname          string     `json:"surname"`
	EmergencyContact string     `json:"emergencyContact"`
	EmergencyNumber  string     `json:"emergencyNumber"`
	Suspended        bool       `json:"suspended"`
	Type             string     `json:"type"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

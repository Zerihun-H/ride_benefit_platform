package model

import (
	"time"

	"gorm.io/gorm"
)

type Relative struct {
	ID          uint64     `json:"id" gorm:"primarykey"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	Relation    string     `json:"relation"`
	PhoneNumber string     `json:"phoneNumber"`
	Age         uint32     `json:"age"`
	Gender      string     `json:"gender"`
	BirthDate   *time.Time `json:"birthDate"`
	PhotoURL    string     `json:"photoURL"`
	IDURL       string     `json:"idURL"`
	EmployeeID  uint64     `json:"employeeID"`
	Employee    Employee   `gorm:"foreignKey:EmployeeID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

package model

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	ID          uint    `json:"id" gorm:"primarykey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	PartnerID   uint64  `json:"partnerID"`
	EmployeeID  uint64  `json:"employeeID"`
	Amount      float64 `json:"amount"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

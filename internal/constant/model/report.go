package model

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	ID          uint   `json:"id" gorm:"primarykey"`
	Description string `json:"description"`
	Name        string `json:"name"`
	PartnerID   uint64 `json:"partnerID"`
	EmployeeID  uint64 `json:"employeeID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

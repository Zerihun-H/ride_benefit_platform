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
	Driver_id   uint64 `json:"driverID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

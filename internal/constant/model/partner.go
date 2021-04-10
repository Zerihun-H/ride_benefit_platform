package model

import (
	"time"

	"gorm.io/gorm"
)

type Partner struct {
	ID          uint64 `gorm:"primarykey"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	PhoneNumber string `json:"phoneNumber"`
	SupportMail string `json:"supportMail"`
	PhotoURL    string `json:"photoURL"`
	ServiceID   uint64 `json:"serviceID"`
	Description string `json:"description"`
	// Users       []User `json:"users" gorm:"foreignKey:PartnerID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint64 `gorm:"primarykey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	PhotoURL  string `json:"photoURL"`
	Username  string `json:"username"`
	PartnerID uint64 `json:"partnerID"`
	RoleID    uint64 `json:"roleID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

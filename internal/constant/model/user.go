package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                     uint64    `gorm:"primarykey"`
	Name                   string    `json:"name"`
	Email                  string    `json:"email"`
	EmailVerifiedAt        time.Time `json:"emailVerifiedAt"`
	Password               string    `json:"password"`
	TwoFactorSecret        string    `json:"twoFactorSecret"`
	TwoFactorRecoveryCodes string    `json:"twoFactorRecoveryCodes"`
	RememberToken          string    `json:"rememberToken"`
	PhotoURL               string    `json:"photoURL"`
	Username               string    `json:"username"`
	PartnerID              uint64    `json:"partnerID"`
	RoleID                 uint64    `json:"roleID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

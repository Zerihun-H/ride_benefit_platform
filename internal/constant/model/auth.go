package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          uint64           `json:"id" gorm:"primarykey"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Permissions []RolePermission `json:"permissions" gorm:"foreignKey:RoleID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Permission struct {
	ID          uint64 `json:"id" gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ResourceID  uint64 `json:"resourceID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RolePermission struct {
	ID           uint64 `json:"id" gorm:"primarykey"`
	RoleID       uint64 `json:"roleID"`
	PermissionID uint64 `json:"permissionID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Resource struct {
	ID   uint   `gorm:"primarykey"`
	Name string `json:"name"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}



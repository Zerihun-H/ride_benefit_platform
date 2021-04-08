package model

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
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

type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AccessTokenClaims ...
type AccessTokenClaims struct {
	UserID uint64 `json:"uid"`
	RoleID string `json:"rle"`
	jwt.StandardClaims
}

// RefreshTokenClaims ...
type RefreshTokenClaims struct {
	UserID uint64 `json:"uid"`
	RoleID string `json:"rle"`
	jwt.StandardClaims
}

package model

import (
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"gorm.io/gorm"
)

type Role struct {
	ID          uint64       `json:"id" gorm:"primarykey"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions;"`

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
	RoleID       uint64     `json:"roleID"`
	PermissionID uint64     `json:"permissionID"`
	Role         Role       `json:"Role" gorm:"foreignKey:RoleID"`
	Permission   Permission `json:"Permission" gorm:"foreignKey:RoleID"`
}

// For implementing an in memory role permission cache
type RolesPermissions map[string][]string

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

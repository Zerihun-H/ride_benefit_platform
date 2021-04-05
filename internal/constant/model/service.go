package model

import (
	"time"

	"gorm.io/gorm"
)

type Service struct {
	ID          uint64    `gorm:"primarykey"`
	Name        string    `json:"name"`
	Code        string    `json:"code"`
	Description string    `json:"description"`
	Partners    []Partner `json:"partners"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

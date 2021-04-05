package model

import (
	"time"

	"gorm.io/gorm"
)

type Activity struct {
	ID          uint      `gorm:"primarykey"`
	UserID      uint64    `json:"userID"`
	Operation   string    `json:"operation"`
	ResourcedID uint64    `json:"resourceID"`
	TimeStamp   time.Time `json:"timeStamp"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

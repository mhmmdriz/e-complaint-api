package entities

import (
	"time"

	"gorm.io/gorm"
)

type ComplaintFile struct {
	ID          int            `gorm:"primaryKey"`
	ComplaintID string         `gorm:"not null;type:varchar;size:15"`
	Path        string         `gorm:"not null"`
	CreateAt    time.Time      `gorm:"autoCreateTime"`
	UpdateAt    time.Time      `gorm:"autoUpdateTime"`
	DeleteAt    gorm.DeletedAt `gorm:"index"`
}
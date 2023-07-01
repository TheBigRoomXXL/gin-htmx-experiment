package db

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"serializer:unixtime;type:time"`
	UpdatedAt time.Time      `gorm:"serializer:unixtime;type:time"`
	DeletedAt gorm.DeletedAt `gorm:"index" gorm:"serializer:unixtime;type:time"`
}

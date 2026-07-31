package entities

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel contiene los campos comunes para todas las entidades
type BaseModel struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

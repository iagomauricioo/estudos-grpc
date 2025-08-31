package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint    `gorm:"primaryKey;autoIncrement"`
	Name        string  `gorm:"not null;size:255"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"not null;type:decimal(10,2)"`
	Quantity    int32   `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

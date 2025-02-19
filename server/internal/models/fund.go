package models

import (
	"time"

	"github.com/google/uuid"
)

// Fund represents a fund created by a fund manager
type Fund struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string    `gorm:"type:varchar(100);not null"`
	ManagerID    uuid.UUID `gorm:"type:uuid;not null"`
	TargetAmount float64   `gorm:"type:decimal(12,2);not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

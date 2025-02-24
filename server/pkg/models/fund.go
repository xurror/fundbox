package models

import (
	"time"

	"github.com/google/uuid"
)

// Fund represents a fund created and managed by a user
type Fund struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string    `gorm:"type:varchar(100);not null"`
	ManagerID    uuid.UUID `gorm:"type:uuid;not null"`
	TargetAmount float64   `gorm:"type:decimal(12,2);not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`

	// Relationships
	Manager       User           `gorm:"foreignKey:ManagerID"` // The user managing the fund
	Contributions []Contribution `gorm:"foreignKey:FundID"`    // Contributions to this fund
}

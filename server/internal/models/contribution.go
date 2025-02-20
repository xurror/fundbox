package models

import (
	"time"

	"github.com/google/uuid"
)

// Contribution represents a financial contribution to a fund
type Contribution struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	FundID        uuid.UUID  `gorm:"type:uuid;not null"`
	ContributorID *uuid.UUID `gorm:"type:uuid"` // Nullable for anonymous contributions
	Anonymous     bool       `gorm:"default:false"`
	Amount        float64    `gorm:"type:decimal(12,2);not null;check:amount > 0"`
	CreatedAt     time.Time  `gorm:"autoCreateTime"`
	UpdatedAt     time.Time  `gorm:"autoUpdateTime"`

	// Relationships
	Fund        Fund `gorm:"foreignKey:FundID"`        // Fund receiving the contribution
	Contributor User `gorm:"foreignKey:ContributorID"` // The user who made the contribution
}

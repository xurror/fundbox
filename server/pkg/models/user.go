package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents a system user who can be a contributor or a fund manager
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Auth0ID   string    `gorm:"type:varchar(255);uniqueIndex"` // Maps to Auth0 user
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relationships
	FundsManaged  []Fund         `gorm:"foreignKey:ManagerID"`     // Funds the user manages
	Contributions []Contribution `gorm:"foreignKey:ContributorID"` // Contributions made
}

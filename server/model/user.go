package model

import (
	"github.com/lib/pq"
)

// User represents a user in the system
type User struct {
	Auditable
	FirstName     string         `json:"first_name" gorm:"not null"`
	LastName      string         `json:"last_name" gorm:"not null"`
	Email         string         `json:"email" gorm:"unique;not null"`
	Password      string         `json:"password,omitempty" gorm:"column:password_hash;not null"`
	Roles         pq.StringArray `json:"roles" gorm:"not null;type:text[]"`
	Contributions []Contribution `json:"-" gorm:"-"`
}

func (user *User) HasRoles(roles []Role) bool {
	for _, role := range roles {
		for _, userRole := range user.Roles {
			if string(role) == userRole {
				return true
			}
		}
	}
	return false
}

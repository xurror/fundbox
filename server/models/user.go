package models

import (
	"github.com/google/uuid"
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

// CreateUser creates a new user in the database
func CreateUser(user *User) (*User, error) {
	result := db.Create(&user)
	return user, HandleError(result.Error)
}

// GetUsers retrieves a list of users
func GetUsers(limit, offset int) ([]*User, error) {
	var users []*User
	result := db.Limit(limit).Find(&users)
	return users, HandleError(result.Error)
}

// GetUser retrieves a user by ID
func GetUser(id uuid.UUID) (*User, error) {
	user := &User{}
	result := db.First(&user, id)
	return user, HandleError(result.Error)
}

// GetUserByEmail retrieves a user by email
func GetUserByEmail(email string) (*User, error) {
	user := &User{}
	result := db.First(&user, "email = ?", email)
	return user, HandleError(result.Error)
}

func GetUserContributions(userID uuid.UUID, limit, offset int) ([]*Contribution, error) {
	var contributions []*Contribution
	result := db.Limit(limit).Find(&contributions, "contributor_id = ?", userID)
	return contributions, HandleError(result.Error)
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

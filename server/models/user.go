package models

// User represents a user in the system
type User struct {
	Auditable
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password,omitempty" gorm:"column:password_hash;not null"`
}

// CreateUser creates a new user in the database
func CreateUser(user *User) (*User, error) {
	result := db.Create(&user)
	return user, HandleError(result.Error)
}

// GetUsers retrieves a list of users
func GetUsers(limit, offset int) ([]*User, error) {
	users := []*User{}
	result := db.Limit(limit).Find(&users)
	return users, HandleError(result.Error)
}

// GetUser retrieves a user by ID
func GetUser(id string) (*User, error) {
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

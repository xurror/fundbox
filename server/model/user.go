package model

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

var ErrNotFound = errors.New("not found")

// const (
// 	UsersTableName = "Fundbox"
// )

type UsersTable struct {
	TableName string
}

type User struct {
	Persistable
	Auth0Id string `json:"email" gorm:"unique;not null" dynamodbav:"auth0_id"`
}

func NewUser(id uuid.UUID, auth0Id string) *User {
	return &User{
		Persistable: Persistable{
			Id: id,
		},
		Auth0Id: auth0Id,
	}
}

// UserService provides user-related services
type UserService interface {
	CreateUser(ctx context.Context, auth0Id string) (*User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (*User, error)
	GetUserByAuth0Id(ctx context.Context, auth0Id string) (*User, error)
	// GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUsers(ctx context.Context, limit int, offset interface{}) ([]*User, interface{}, error)

	// GetUserContributions(ctx context.Context, userID uuid.UUID, limit, offset int) ([]*model.Contribution, error)
}

// GetKey returns the composite primary key of the movie in a format that can be
// sent to DynamoDB.
func (user User) GetKey() map[string]types.AttributeValue {
	id, err := attributevalue.Marshal(user.Id)
	if err != nil {
		panic(err)
	}
	auth0Id, err := attributevalue.Marshal(user.Auth0Id)
	if err != nil {
		panic(err)
	}
	return map[string]types.AttributeValue{"id": id, "auth0_id": auth0Id}
}

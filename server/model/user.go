package model

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

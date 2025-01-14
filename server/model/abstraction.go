package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Persistable struct {
	Id uuid.UUID `json:"id" gorm:"primary_key;type:uuid" dynamodbav:"id"`
}

type Auditable struct {
	Persistable
	CreatedAt time.Time      `json:"created_at" dynamodbav:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" dynamodbav:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index" dynamodbav:"deleted_at"`
}

package model

import (
	"github.com/google/uuid"
)

type Currency struct {
	Persistable
	Name string `json:"name" gorm:"unique;not null"`
}

type Amount struct {
	Persistable
	Value      float64   `json:"value" gorm:"type:decimal;default:0.0"`
	CurrencyID uuid.UUID `json:"-" gorm:"not null"`
	Currency   Currency  `json:"currency"`
}

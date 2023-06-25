package model

import (
	"github.com/google/uuid"
)

type Contribution struct {
	Auditable
	AmountID      uuid.UUID `json:"-" gorm:"not null;type:uuid"`
	Amount        Amount    `json:"amount"`
	FundID        uuid.UUID `json:"-" gorm:"not null;type:uuid"`
	Fund          Fund      `json:"fund"`
	ContributorID uuid.UUID `json:"-" gorm:"not null;type:uuid"`
	Contributor   User      `json:"contributor" gorm:"foreignKey:ContributorID;references:ID"`
}

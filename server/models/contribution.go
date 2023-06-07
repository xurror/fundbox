package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
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

func CreateContribution(contribution *Contribution) (*Contribution, error) {
	result := db.Preload(clause.Associations).Preload("Amount.Currency").Create(&contribution)
	return contribution, HandleError(result.Error)
}

func GetContribution(id string) (*Contribution, error) {
	contribution := &Contribution{}
	result := db.Preload(clause.Associations).Preload("Amount.Currency").First(&contribution, "id = ?", id)
	return contribution, HandleError(result.Error)
}

func GetContributions(limit, offset int) ([]*Contribution, error) {
	contributions := []*Contribution{}
	result := db.Preload(clause.Associations).Preload("Amount.Currency").Limit(limit).Find(&contributions)
	return contributions, HandleError(result.Error)
}

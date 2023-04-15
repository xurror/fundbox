package models

import "gorm.io/gorm/clause"

type Contribution struct {
	Auditable
	AmountID      string      `json:"-" gorm:"not null"`
	Amount        Amount      `json:"amount"`
	FundID        string      `json:"-" gorm:"not null"`
	Fund          Fund        `json:"fund"`
	ContributorID string      `json:"-" gorm:"not null"`
	Contributor   Contributor `json:"contributor"`
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

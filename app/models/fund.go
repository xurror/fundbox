package models

type Fund struct {
	Auditable
	Reason        string         `json:"reason" binding:"required" gorm:"not null"`
	Description   string         `json:"description"`
	Contributions []Contribution `json:"-"`
}

func CreateFund(fund *Fund) (*Fund, error) {
	result := db.Create(&fund)
	return fund, HandleError(result.Error)
}

func GetFund(id string) (*Fund, error) {
	fund := &Fund{}
	result := db.First(&fund, "id = ?", id)
	return fund, HandleError(result.Error)
}

func GetFunds(limit, offset int) ([]*Fund, error) {
	funds := []*Fund{}
	result := db.Limit(limit).Find(&funds)
	return funds, HandleError(result.Error)
}

func GetFundContributions(fundID string, limit, offset int) ([]*Contribution, error) {
	contributions := []*Contribution{}
	result := db.Limit(limit).Find(&contributions, "fund_id = ?", fundID)
	return contributions, HandleError(result.Error)
}

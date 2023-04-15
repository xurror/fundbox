package models

type Contributor struct {
	Auditable
	Name          string         `json:"name" gorm:"not null"`
	Contributions []Contribution `json:"-"`
}

func CreateContributor(contributor *Contributor) (*Contributor, error) {
	result := db.Create(&contributor)
	return contributor, HandleError(result.Error)
}

func GetContributor(id string) (*Contributor, error) {
	contributor := &Contributor{}
	result := db.First(&contributor, "id = ?", id)
	return contributor, HandleError(result.Error)
}

func GetContributors(limit, offset int) ([]*Contributor, error) {
	contributors := []*Contributor{}
	result := db.Limit(limit).Find(&contributors)
	return contributors, HandleError(result.Error)
}

func GetContributorContributions(contributorID string, limit, offset int) ([]*Contribution, error) {
	contributions := []*Contribution{}
	result := db.Limit(limit).Find(&contributions, "contributor_id = ?", contributorID)
	return contributions, HandleError(result.Error)
}

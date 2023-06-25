package model

type Fund struct {
	Auditable
	Reason        string         `json:"reason" binding:"required" gorm:"not null"`
	Description   string         `json:"description"`
	Contributions []Contribution `json:"-"`
}

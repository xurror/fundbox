package service

import (
	"getting-to-go/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ContributionService struct {
	db *gorm.DB
}

func NewContributionService(db *gorm.DB) *ContributionService {
	return &ContributionService{
		db: db,
	}
}

func (s *ContributionService) CreateContribution(fundID, contributorID uuid.UUID, amount float64, currencyID uuid.UUID) (*model.Contribution, error) {
	contribution := &model.Contribution{
		FundID:        fundID,
		ContributorID: contributorID,
	}

	contribution.Amount = model.Amount{Value: amount, CurrencyID: currencyID}
	result := s.db.Preload(clause.Associations).Preload("Amount.Currency").Create(&contribution)
	return contribution, model.HandleError(result.Error)
}

func (s *ContributionService) GetContribution(id string) (*model.Contribution, error) {
	contribution := &model.Contribution{}
	result := s.db.Preload(clause.Associations).Preload("Amount.Currency").First(&contribution, "id = ?", id)
	return contribution, model.HandleError(result.Error)
}

func (s *ContributionService) GetContributions(limit, offset int) ([]*model.Contribution, error) {
	contributions := []*model.Contribution{}
	result := s.db.Preload(clause.Associations).Preload("Amount.Currency").Limit(limit).Find(&contributions)
	return contributions, model.HandleError(result.Error)
}

func (s *ContributionService) GetContributionsByUserID(userID uuid.UUID, limit, offset int) ([]*model.Contribution, error) {
	var contributions []*model.Contribution
	result := s.db.Limit(limit).Find(&contributions, "contributor_id = ?", userID)
	return contributions, model.HandleError(result.Error)
}

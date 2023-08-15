package service

import (
	"getting-to-go/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CurrencyService struct {
	db *gorm.DB
}

func NewCurrencyService(db *gorm.DB) *CurrencyService {
	return &CurrencyService{
		db: db,
	}
}

func (s *CurrencyService) GetCurrencies() ([]*model.Currency, error) {
	var currencies []*model.Currency
	result := s.db.Preload(clause.Associations).Find(&currencies)
	return currencies, model.HandleError(result.Error)
}

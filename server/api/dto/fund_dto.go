package dto

import (
	"time"

	"community-funds/pkg/models"

	"github.com/google/uuid"
)

type FundDTO struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	TargetAmount float64   `json:"targetAmount"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func MapFundToDTO(fund models.Fund) FundDTO {
	return FundDTO{
		ID:           fund.ID,
		Name:         fund.Name,
		TargetAmount: fund.TargetAmount,
		CreatedAt:    fund.CreatedAt,
		UpdatedAt:    fund.UpdatedAt,
	}
}

func MapFundsToDTOs(funds []models.Fund) []FundDTO {
	dtos := make([]FundDTO, len(funds))
	for i, fund := range funds {
		dtos[i] = MapFundToDTO(fund)
	}
	return dtos
}

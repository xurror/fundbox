package dto

import (
	"time"

	"community-funds/pkg/models"

	"github.com/google/uuid"
)

// FundDTO represents the data returned to clients for a fund
type FundDTO struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	TargetAmount float64   `json:"targetAmount"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// MapFundToDTO converts a Fund model into a FundDTO
func MapFundToDTO(fund models.Fund) FundDTO {
	return FundDTO{
		ID:           fund.ID,
		Name:         fund.Name,
		TargetAmount: fund.TargetAmount,
		CreatedAt:    fund.CreatedAt,
		UpdatedAt:    fund.UpdatedAt,
	}
}

// MapFundsToDTOs converts a slice of Fund models into a slice of FundDTOs
func MapFundsToDTOs(funds []models.Fund) []FundDTO {
	dtos := make([]FundDTO, len(funds))
	for i, fund := range funds {
		dtos[i] = MapFundToDTO(fund)
	}
	return dtos
}

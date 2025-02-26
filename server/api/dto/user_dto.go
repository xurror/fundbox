package dto

import (
	"community-funds/pkg/models"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID uuid.UUID `json:"id"`
}

func MapUserToDTO(fund models.Fund) FundDTO {
	return FundDTO{
		ID:           fund.ID,
		Name:         fund.Name,
		TargetAmount: fund.TargetAmount,
		CreatedAt:    fund.CreatedAt,
		UpdatedAt:    fund.UpdatedAt,
	}
}

func MapUsersToDTOs(funds []models.Fund) []FundDTO {
	dtos := make([]FundDTO, len(funds))
	for i, fund := range funds {
		dtos[i] = MapFundToDTO(fund)
	}
	return dtos
}

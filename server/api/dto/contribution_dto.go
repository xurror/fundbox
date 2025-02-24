package dto

import (
	"time"

	"community-funds/pkg/models"

	"github.com/google/uuid"
)

// ContributionDTO represents a structured response for contributions
type ContributionDTO struct {
	ID              uuid.UUID  `json:"id"`
	FundID          uuid.UUID  `json:"fundId"`
	FundName        string     `json:"fundName"`
	ContributorID   *uuid.UUID `json:"contributorId,omitempty"`   // Null if anonymous
	ContributorName *string    `json:"contributorName,omitempty"` // Null if anonymous
	Amount          float64    `json:"amount"`
	Anonymous       bool       `json:"anonymous"`
	CreatedAt       time.Time  `json:"createdAt"`
}

// MapContributionToDTO converts a Contribution model into a ContributionDTO
func MapContributionToDTO(contribution models.Contribution) ContributionDTO {
	return ContributionDTO{
		ID:              contribution.ID,
		FundID:          contribution.FundID,
		FundName:        contribution.Fund.Name,
		ContributorID:   contribution.ContributorID,
		ContributorName: &contribution.Contributor.Name,
		Amount:          contribution.Amount,
		Anonymous:       contribution.Anonymous,
		CreatedAt:       contribution.CreatedAt,
	}
}

// MapContributionsToDTOs converts a slice of Contribution models into DTOs
func MapContributionsToDTOs(contributions []models.Contribution) []ContributionDTO {
	dtos := make([]ContributionDTO, len(contributions))
	for i, c := range contributions {
		dtos[i] = MapContributionToDTO(c)
	}
	return dtos
}

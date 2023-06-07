package resolvers

import "getting-to-go/models"

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users         []*models.User
	funds         []*models.Fund
	contributions []*models.Contribution
}

package resolver

import (
	"getting-to-go/service"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService         *service.UserService
	FundService         *service.FundService
	CurrencyService     *service.CurrencyService
	ContributionService *service.ContributionService
	AuthService         *service.AuthService
}

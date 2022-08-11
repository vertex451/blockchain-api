package graphql

import (
	"silverspase/blockchain-api/golang/internal/service/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	usecase usecase.Provider
}

// NewResolver creates new instance of Resolver.
func NewResolver(usecase usecase.Provider) *Resolver {
	return &Resolver{
		usecase: usecase,
	}
}

package graph

import "github.com/marceloaguero/go-graphql-products/providers/model/provider"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	usecase provider.Usecase
}

func NewResolver(uc provider.Usecase) *Resolver {
	return &Resolver{
		usecase: uc,
	}
}

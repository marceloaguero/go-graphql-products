package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/marceloaguero/go-graphql-products/products/model/product"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	usecase product.Usecase
}

func NewResolver(uc product.Usecase) *Resolver {
	return &Resolver{
		usecase: uc,
	}
}

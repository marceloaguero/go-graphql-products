package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/marceloaguero/go-graphql-products/products/graph/model"
	"github.com/marceloaguero/go-graphql-products/products/model/product"
)

// CreateProduct is the resolver for the createProduct field.
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.ProductInput) (*product.Product, error) {
	newProduct := &product.Product{
		Name:        input.Name,
		Description: *input.Description,
		Unit:        input.Unit,
		Price:       input.Price,
		Stock:       *input.Stock,
		IsActive:    *input.IsActive,
	}
	product, err := r.usecase.Create(newProduct)
	return product, err
}

// UpdateProduct is the resolver for the updateProduct field.
func (r *mutationResolver) UpdateProduct(ctx context.Context, id string, input model.ProductInput) (*product.Product, error) {
	panic(fmt.Errorf("not implemented: UpdateProduct - updateProduct"))
}

// DeleteProduct is the resolver for the deleteProduct field.
func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (*product.Product, error) {
	panic(fmt.Errorf("not implemented: DeleteProduct - deleteProduct"))
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*product.Product, error) {
	product, err := r.usecase.GetByID(id)
	return product, err
}

// ProductByName is the resolver for the productByName field.
func (r *queryResolver) ProductByName(ctx context.Context, name string) (*product.Product, error) {
	product, err := r.usecase.GetByName(name)
	return product, err
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*product.Product, error) {
	products, err := r.usecase.GetAll()
	return products, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

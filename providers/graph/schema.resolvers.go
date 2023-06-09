package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/marceloaguero/go-graphql-products/providers/graph/model"
	"github.com/marceloaguero/go-graphql-products/providers/model/provider"
)

// CreateProvider is the resolver for the createProvider field.
func (r *mutationResolver) CreateProvider(ctx context.Context, input model.ProviderInput) (*provider.Provider, error) {
	newProvider := &provider.Provider{
		Name:     input.Name,
		IsActive: *input.IsActive,
	}
	provider, err := r.usecase.Create(newProvider)
	return provider, err
}

// UpdateProvider is the resolver for the updateProvider field.
func (r *mutationResolver) UpdateProvider(ctx context.Context, id string, input model.ProviderInput) (*provider.Provider, error) {
	changedProvider := &provider.Provider{
		ID:       id,
		Name:     input.Name,
		IsActive: *input.IsActive,
	}
	provider, err := r.usecase.Update(changedProvider)
	return provider, err
}

// DeleteProvider is the resolver for the deleteProvider field.
func (r *mutationResolver) DeleteProvider(ctx context.Context, id string) (*provider.Provider, error) {
	panic(fmt.Errorf("not implemented: DeleteProvider - deleteProvider"))
}

// Provider is the resolver for the provider field.
func (r *queryResolver) Provider(ctx context.Context, id string) (*provider.Provider, error) {
	provider, err := r.usecase.GetByID(id)
	return provider, err
}

// ProviderByName is the resolver for the providerByName field.
func (r *queryResolver) ProviderByName(ctx context.Context, name string) (*provider.Provider, error) {
	provider, err := r.usecase.GetByName(name)
	return provider, err
}

// Providers is the resolver for the providers field.
func (r *queryResolver) Providers(ctx context.Context) ([]*provider.Provider, error) {
	providers, err := r.usecase.GetAll()
	return providers, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

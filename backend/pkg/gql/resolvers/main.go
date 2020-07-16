package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	gql "testapp/pkg/gql"
	models "testapp/pkg/gql/models"
)

// Resolver <>
type Resolver struct{}

func (r *mutationResolver) Createproduct(ctx context.Context, input models.ProductInput) (*models.Product, error) {
	panic("not implemented")
}

func (r *mutationResolver) Updateproduct(ctx context.Context, input models.ProductInput, productID string) (*models.Product, error) {
	panic("not implemented")
}

func (r *mutationResolver) Deleteproduct(ctx context.Context, productID string) (*bool, error) {
	panic("not implemented")
}

func (r *queryResolver) Products(ctx context.Context) ([]*models.Product, error) {
	panic("not implemented")
}

func (r *queryResolver) Productwithid(ctx context.Context, productID string) (*models.Product, error) {
	panic("not implemented")
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	gql "testapp/pkg/gql"
	models "testapp/pkg/gql/models"
)

// Resolver <>
type Resolver struct{}

// TODO: this mongodb application cause error and stuck here if db not exist
// okay we find the problem
// TODO: docker-compose not show any !(okay find the probelm is the same as mongodb problem)
func (r *mutationResolver) Createproduct(ctx context.Context, input models.ProductInput) (*models.Product, error) {
	panic("not implemented")
	// if !*db.Connecting {
	// 	panic("db not connect")
	// }
	// iData := db.Product{
	// 	Title:    *input.Title,
	// 	Price:    *input.Price,
	// 	Discount: *input.Discount,
	// 	Code:     *input.Code,
	// 	Content:  *input.Content,
	// }

	// result, err := db.CreateProduct(&iData)
	// if err != nil {
	// 	return nil, err
	// }
	// id := result.ID.Hex()
	// oData := models.Product{
	// 	ID:       &id,
	// 	Title:    &result.Title,
	// 	Price:    &result.Price,
	// 	Discount: &result.Discount,
	// 	Code:     &result.Code,
	// 	Content:  &result.Content,
	// }
	// return &oData, nil
}

func (r *mutationResolver) Updateproduct(ctx context.Context, input models.ProductInput, productID string) (*models.Product, error) {
	panic("not implemented")
}

func (r *mutationResolver) Deleteproduct(ctx context.Context, productID string) (*bool, error) {
	panic("not implemented")
}

func (r *queryResolver) Products(ctx context.Context) ([]*models.Product, error) {
	panic("not implemented")
	// if !*db.Connecting {
	// 	panic("db not connect")
	// }
	// result, err := db.ReadproductAll()
	// if err != nil {
	// 	return nil, err
	// }
	// oData := []*models.Product{}

	// for _, v := range result {
	// 	id := v.ID.Hex()
	// 	oObj := models.Product{
	// 		ID:       &id,
	// 		Title:    &v.Title,
	// 		Price:    &v.Price,
	// 		Code:     &v.Code,
	// 		Discount: &v.Discount,
	// 		Content:  &v.Content,
	// 	}
	// 	oData = append(oData, &oObj)
	// }

	// return oData, nil
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

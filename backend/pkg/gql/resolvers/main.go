package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"log"
	"sync"
	db "testapp/pkg/db"
	gql "testapp/pkg/gql"
	models "testapp/pkg/gql/models"

	"github.com/globalsign/mgo"
)

// Resolver <>
type Resolver struct {
}

func (r *mutationResolver) Createproduct(ctx context.Context, input models.ProductInput) (*models.Product, error) {
	// mgoInfo := &mgo.DialInfo{
	// 	Addrs:    []string{"localhost:27017"},
	// 	Database: "shop",
	// 	Username: "michael",
	// 	Password: "lneequal1",
	// }
	// session, err := mgo.DialWithInfo(mgoInfo)
	// defer session.Close()
	// if err != nil {
	// 	log.Fatalf("this is error : %v", err)
	// 	return nil, err
	// }

	// iData := db.Product{
	// 	Title:    *input.Title,
	// 	Price:    *input.Price,
	// 	Discount: *input.Discount,
	// 	Code:     *input.Code,
	// 	Content:  *input.Content,
	// }
	// database := session.DB("shop")
	// C := database.C("product")
	// err = C.Insert(&iData)
	// if err != nil {
	// 	log.Fatalf("this is error : %v", err)
	// 	return nil, err
	// }

	// dboData := db.Product{}
	// err = C.Find(bson.M{"title": input.Title}).One(&dboData)
	// if err != nil {
	// 	log.Fatalf("this is error : %v", err)
	// 	return nil, err
	// }
	// id := dboData.ID.Hex()
	// oData := models.Product{
	// 	ID:       &id,
	// 	Title:    &dboData.Title,
	// 	Price:    &dboData.Price,
	// 	Discount: &dboData.Discount,
	// 	Code:     &dboData.Code,
	// 	Content:  &dboData.Content,
	// 	Avator:   &dboData.Avator,
	// }
	var oData models.Product
	wg := sync.WaitGroup{}
	wg.Add(1)
	// okay here to read input chan
	go func(wg *sync.WaitGroup) {
		insData := db.Product{
			Title:    *input.Title,
			Price:    *input.Price,
			Discount: *input.Discount,
			Code:     *input.Code,
			Content:  *input.Content,
		}
		db.SetinsertInChan(insData)
		for {
			select {
			case result := <-db.GetinsertOutChan():
				id := result.ID.Hex()
				oData = models.Product{
					ID:       &id,
					Title:    &result.Title,
					Price:    &result.Price,
					Discount: &result.Discount,
					Code:     &result.Code,
					Content:  &result.Content,
					Avator:   &result.Avator,
				}
				wg.Done()
			}
		}
	}(&wg)
	wg.Wait()
	return &oData, nil
}

func (r *mutationResolver) Updateproduct(ctx context.Context, input models.ProductInput, productID string) (*models.Product, error) {
	panic("not implemented")
}

func (r *mutationResolver) Deleteproduct(ctx context.Context, productID string) (*bool, error) {
	panic("not implemented")
}

func (r *queryResolver) Products(ctx context.Context) ([]*models.Product, error) {
	mgoInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Database: "shop",
		Username: "michael",
		Password: "lneequal1",
	}
	session, err := mgo.DialWithInfo(mgoInfo)
	defer session.Close()
	if err != nil {
		log.Fatalf("this is error : %v", err)
		return nil, err
	}
	database := session.DB("shop")
	C := database.C("product")
	entry := []db.Product{}
	err = C.Find(nil).All(&entry)
	if err != nil {
		log.Fatalf("this is error : %v", err)
		return nil, err
	}
	oData := []*models.Product{}
	for _, v := range entry {
		id := v.ID.Hex()
		tmp := models.Product{
			ID:       &id,
			Title:    &v.Title,
			Price:    &v.Price,
			Discount: &v.Discount,
			Code:     &v.Code,
			Content:  &v.Content,
			Avator:   &v.Avator,
		}
		oData = append(oData, &tmp)
	}
	return oData, nil
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

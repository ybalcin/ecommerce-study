package adapters

import (
	"context"
	"fmt"
	"github.com/ybalcin/ecommerce-study/pkg/mgo"
	"testing"
)

//func TestNewProductRepository(t *testing.T) {
//	store := mgo.NewStore(context.Background(), "mongodb+srv://ecommerce-user:B9VeLojwHUidkeHP@cluster0.l1pmb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", "ecommerce")
//
//	repo := NewProductMongoRepository(store)
//	if repo == nil {
//		t.Fail()
//	}
//}
//
//func TestProductRepository_AddProduct(t *testing.T) {
//	ctx := context.Background()
//
//	store := mgo.NewStore(ctx, "mongodb+srv://ecommerce-user:B9VeLojwHUidkeHP@cluster0.l1pmb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", "ecommerce")
//
//	repo := NewProductMongoRepository(store)
//
//	product, err := product.NewProduct(product.ProductId(primitive.NewObjectID().Hex()), "p1", 123.5, 15)
//	if err != nil {
//		t.Fail()
//	}
//
//	if err := repo.AddProduct(ctx, product); err != nil {
//		t.Fail()
//	}
//}

func TestProductMongoRepository_GetProduct(t *testing.T) {
	ctx := context.Background()

	store := mgo.NewStore(ctx, "mongodb+srv://ecommerce-user:B9VeLojwHUidkeHP@cluster0.l1pmb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority", "ecommerce")

	repo, _ := NewProductMongoRepository(store)

	product, err := repo.GetProduct(ctx, "p1asd")
	if err != nil {
		t.Fail()
	}

	if product == nil {
		t.Fail()
	}

	fmt.Printf("%#v", product)
}

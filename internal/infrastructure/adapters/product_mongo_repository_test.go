package adapters

//
//import (
//	"context"
//	"fmt"
//	"github.com/ybalcin/ecommerce-study/internal/domain/models"
//	"github.com/ybalcin/ecommerce-study/pkg/mgo"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"testing"
//)
//
//func TestNewProductRepository(t *testing.T) {
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
//
//	repo := NewProductMongoRepository(store)
//
//	product, err := models.NewProduct(models.ProductId(primitive.NewObjectID().Hex()), "p1", 123.5, 15)
//	if err != nil {
//		t.Fail()
//	}
//
//	if err := repo.AddProduct(ctx, product); err != nil {
//		t.Fail()
//	}
//}
//
//func TestProductMongoRepository_GetProduct(t *testing.T) {
//	ctx := context.Background()
//
//
//	repo := NewProductMongoRepository(store)
//
//	product, err := repo.GetProduct(ctx, models.ProductId("61a75a6b9836030435bce659"))
//	if err != nil {
//		t.Fail()
//	}
//
//	if product == nil {
//		t.Fail()
//	}
//
//	fmt.Printf("%#v", product)
//}

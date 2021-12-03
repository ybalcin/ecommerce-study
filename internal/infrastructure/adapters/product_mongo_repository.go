package adapters

import (
	"context"
	"errors"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/pkg/mgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// productBson is collection store model
type productBson struct {
	Id    primitive.ObjectID `bson:"_id"`
	Code  string             `bson:"code"`
	Price int                `bson:"price"`
	Stock int                `bson:"stock"`
}

// productMongoRepository implements repositories.ProductRepository
type productMongoRepository struct {
	products *mgo.Collection
}

const productCollection string = "products"

func (p *productBson) hasValue() bool {
	if p == nil {
		return false
	}

	return p.Id != primitive.NilObjectID
}

// NewProductMongoRepository initializes new product repository
func NewProductMongoRepository(store *mgo.Store) (*productMongoRepository, error) {
	if store == nil {
		return nil, errors.New("adapters: store is nil")
	}

	return &productMongoRepository{products: store.Collection(productCollection)}, nil
}

// AddProduct adds product to collection
func (r *productMongoRepository) AddProduct(ctx context.Context, product *domain.Product) error {

	productBson := productBson{
		Id:    primitive.NewObjectID(),
		Code:  product.Code(),
		Price: product.Price(),
		Stock: product.Stock(),
	}

	if _, err := r.products.InsertOne(ctx, productBson); err != nil {
		return err
	}

	return nil
}

// GetProduct gets product from collection
func (r *productMongoRepository) GetProduct(ctx context.Context, productCode string) (*domain.Product, error) {
	productBson := new(productBson)

	if err := r.products.FindOne(ctx, bson.M{"code": productCode}, productBson); err != nil {
		return nil, err
	}

	if !productBson.hasValue() {
		return nil, nil
	}

	productModel, err := domain.NewProduct(productBson.Id.Hex(), productBson.Code, productBson.Price, productBson.Stock)
	if err != nil {
		return nil, err
	}

	return productModel, nil
}

// UpdateProduct updates product
func (r *productMongoRepository) UpdateProduct(ctx context.Context, product *domain.Product) error {
	find := bson.M{
		"_id": primitive.ObjectIDFromHex(product.Id()),
	}

	if err := r.products.UpdateOne(ctx, find, product); err != nil {
		return err
	}

	return nil
}

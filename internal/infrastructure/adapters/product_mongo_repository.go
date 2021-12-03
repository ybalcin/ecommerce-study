package adapters

import (
	"context"
	"errors"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/pkg/mgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"_id", -1}})
	findOptions.SetLimit(1)

	filter := bson.D{
		{"code", productCode},
	}

	var sProductBson []productBson

	if err := r.products.Find(ctx, filter, &sProductBson, findOptions); err != nil {
		return nil, err
	}

	if len(sProductBson) <= 0 {
		return nil, nil
	}

	productBson := sProductBson[0]

	productModel, err := domain.NewProduct(productBson.Id.Hex(), productBson.Code, productBson.Price, productBson.Stock)
	if err != nil {
		return nil, err
	}

	return productModel, nil
}

// UpdateProductStock updates product
func (r *productMongoRepository) UpdateProductStock(ctx context.Context, product *domain.Product) error {
	id, err := primitive.ObjectIDFromHex(product.Id())
	if err != nil {
		return err
	}

	find := bson.M{
		"_id": id,
	}

	updates := bson.D{
		{"$set", bson.D{{"stock", product.Stock()}}},
	}

	if err := r.products.UpdateOne(ctx, find, updates); err != nil {
		return err
	}

	return nil
}

// DropProducts deletes all products
func (r *productMongoRepository) DropProducts(ctx context.Context) error {
	if err := r.products.DeleteMany(ctx, bson.D{}); err != nil {
		return err
	}

	return nil
}

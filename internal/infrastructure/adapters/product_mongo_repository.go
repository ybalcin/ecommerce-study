package adapters

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/domain/models"
	"github.com/ybalcin/ecommerce-study/pkg/mgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// productBson is collection store model
type productBson struct {
	Id          primitive.ObjectID `bson:"_id"`
	ProductCode string             `bson:"product_code"`
	Price       float64            `bson:"price"`
	Stock       int                `bson:"stock"`
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
func NewProductMongoRepository(store *mgo.Store) *productMongoRepository {
	if store == nil {
		panic("adapters: mgo store is nil")
	}

	return &productMongoRepository{products: store.Collection(productCollection)}
}

// AddProduct adds product to collection
func (r *productMongoRepository) AddProduct(ctx context.Context, product *models.Product) error {
	objId, err := mgo.ToObjectID(product.Id().String())
	if err != nil {
		return err
	}

	productBson := productBson{
		Id:          objId,
		ProductCode: product.ProductCode(),
		Price:       product.Price(),
		Stock:       product.Stock(),
	}

	if _, err := r.products.InsertOne(ctx, productBson); err != nil {
		return err
	}

	return nil
}

// GetProduct gets product from collection
func (r *productMongoRepository) GetProduct(ctx context.Context, id models.ProductId) (*models.Product, error) {
	productBson := new(productBson)

	objId, err := mgo.ToObjectID(id.String())
	if err != nil {
		return nil, err
	}

	if err := r.products.FindOne(ctx, bson.M{"_id": objId}, productBson); err != nil {
		return nil, err
	}

	if !productBson.hasValue() {
		return nil, nil
	}

	productModel, err := models.NewProduct(models.ProductId(productBson.Id.Hex()), productBson.ProductCode, productBson.Price, productBson.Stock)
	if err != nil {
		return nil, err
	}

	return productModel, nil
}

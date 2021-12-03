package adapters

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/pkg/mgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type orderBson struct {
	Id            primitive.ObjectID `bson:"_id"`
	ProductCode   string             `bson:"product_code"`
	Quantity      int                `bson:"quantity"`
	UnitSalePrice int                `bson:"unit_sale_price"`
}

// OrderRepository implements repositories.OrderRepository
type OrderRepository struct {
	orders *mgo.Collection
}

const orderCollection string = "orders"

// NewOrderRepository initializes new order repository
func NewOrderRepository(store *mgo.Store) (*OrderRepository, error) {
	if store == nil {
		return nil, errors.New("adapters: store is nil")
	}

	return &OrderRepository{
		orders: store.Collection(orderCollection),
	}, nil
}

// GetOrders gets orders by productCode
func (r *OrderRepository) GetOrders(ctx context.Context, productCode string) ([]domain.Order, error) {
	var ordersBson []orderBson

	filter := bson.M{
		"product_code": productCode,
	}

	if err := r.orders.Find(ctx, filter, &ordersBson); err != nil {
		return nil, err
	}

	var orders []domain.Order
	for _, o := range ordersBson {
		order, err := o.mapToOrder()
		if err != nil {
			continue
		}

		orders = append(orders, *order)
	}

	return orders, nil
}

// AddOrder adds order to collection
func (r *OrderRepository) AddOrder(ctx context.Context, order *domain.Order) error {
	orderBson := &orderBson{
		Id:            primitive.NewObjectID(),
		ProductCode:   order.ProductCode(),
		Quantity:      order.Quantity(),
		UnitSalePrice: order.UnitSalePrice(),
	}

	if _, err := r.orders.InsertOne(ctx, orderBson); err != nil {
		return err
	}

	return nil
}

// DropOrders deletes all orders
func (r *OrderRepository) DropOrders(ctx context.Context) error {
	if err := r.orders.DeleteMany(ctx, bson.D{}); err != nil {
		return err
	}

	return nil
}

func (o *orderBson) mapToOrder() (*domain.Order, error) {
	order, err := domain.NewOrder(o.ProductCode, o.Quantity, o.UnitSalePrice)
	if err != nil {
		return nil, err
	}

	return order, nil
}

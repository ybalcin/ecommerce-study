package repositories

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/domain"
)

// OrderRepository interface wraps order operations
type OrderRepository interface {
	// GetOrders gets orders by product
	GetOrders(ctx context.Context, productCode string) ([]domain.Order, error)

	// AddOrder adds order to collection
	AddOrder(ctx context.Context, order *domain.Order) error

	// DropOrders deletes all orders
	DropOrders(ctx context.Context) error
}

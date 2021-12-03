package repositories

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/domain"
)

// ProductRepository interface wraps product operations
type ProductRepository interface {
	// AddProduct adds product to collection
	AddProduct(ctx context.Context, product *domain.Product) error

	// GetProduct gets product from collection
	GetProduct(ctx context.Context, productCode string) (*domain.Product, error)

	// UpdateProductStock updates product
	UpdateProductStock(ctx context.Context, product *domain.Product) error

	// DropProducts deletes all products
	DropProducts(ctx context.Context) error
}

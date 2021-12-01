package repositories

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/domain/models"
)

// ProductRepository interface wraps product operations
type ProductRepository interface {
	// AddProduct adds product to collection
	AddProduct(ctx context.Context, product *models.Product) error

	// GetProduct gets product from collection
	GetProduct(ctx context.Context, id models.ProductId) (*models.Campaign, error)
}

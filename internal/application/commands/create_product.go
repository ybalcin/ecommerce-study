package commands

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
)

type CreateProductCommand struct {
	ProductCode string
	Price       int
	Stock       int
}

type CreateProductCommandHandler struct {
	productRepository repositories.ProductRepository
}

// NewAddProductCommandHandler initializes new AddProductCommandHandler
func NewAddProductCommandHandler(productRepository repositories.ProductRepository) *CreateProductCommandHandler {
	return &CreateProductCommandHandler{
		productRepository: productRepository,
	}
}

func (h *CreateProductCommandHandler) Handle(ctx context.Context, c *CreateProductCommand) error {
	product, err := domain.NewProduct("", c.ProductCode, c.Price, c.Stock)
	if err != nil {
		return err
	}

	if err := h.productRepository.AddProduct(ctx, product); err != nil {
		return err
	}

	return nil
}

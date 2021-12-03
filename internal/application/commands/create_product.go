package commands

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/application"
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

// NewCreateProductCommandHandler initializes new AddProductCommandHandler
func NewCreateProductCommandHandler(productRepository repositories.ProductRepository) *CreateProductCommandHandler {
	return &CreateProductCommandHandler{
		productRepository: productRepository,
	}
}

func (h *CreateProductCommandHandler) Handle(ctx context.Context, c *CreateProductCommand) error {
	if h == nil {
		return application.ThrowCreateProductCommandHandlerCannotBeNilError()
	}

	if err := h.validate(); err != nil {
		return err
	}

	product, err := domain.NewProduct("", c.ProductCode, c.Price, c.Stock)
	if err != nil {
		return err
	}

	if err := h.productRepository.AddProduct(ctx, product); err != nil {
		return err
	}

	return nil
}

func (h *CreateProductCommandHandler) validate() error {
	if h.productRepository == nil {
		return application.ThrowProductRepositoryCannotBeNil()
	}

	return nil
}

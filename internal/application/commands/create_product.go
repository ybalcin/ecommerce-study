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

func (h *CreateProductCommandHandler) Handle(ctx context.Context, c *CreateProductCommand) (*createProductResponse, error) {
	if h == nil {
		return nil, application.ThrowCreateProductCommandHandlerCannotBeNilError()
	}

	if err := h.validate(); err != nil {
		return nil, err
	}

	if c == nil {
		return nil, application.ThrowCreateProductCommandCannotBeNilError()
	}

	product, err := domain.NewProduct("", c.ProductCode, c.Price, c.Stock)
	if err != nil {
		return nil, err
	}

	if err := h.productRepository.AddProduct(ctx, product); err != nil {
		return nil, err
	}

	return NewCreateProductResponse(product.Code(), product.Price(), product.Stock()), nil
}

func (h *CreateProductCommandHandler) validate() error {
	if h.productRepository == nil {
		return application.ThrowProductRepositoryCannotBeNil()
	}

	return nil
}

package createproduct

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
)

type Command struct {
	ProductCode string
	Price       int
	Stock       int
}

type Handler struct {
	productRepository repositories.ProductRepository
}

// NewHandler initializes new AddProductCommandHandler
func NewHandler(productRepository repositories.ProductRepository) *Handler {
	return &Handler{
		productRepository: productRepository,
	}
}

func (h *Handler) Handle(ctx context.Context, c *Command) (*response, error) {
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

	return NewResponse(product.Code(), product.Price(), product.Stock()), nil
}

func (h *Handler) validate() error {
	if h.productRepository == nil {
		return application.ThrowProductRepositoryCannotBeNil()
	}

	return nil
}

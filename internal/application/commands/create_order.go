package commands

import (
	"context"
	"fmt"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
	"github.com/ybalcin/ecommerce-study/internal/domain/services"
)

type CreateOrderCommand struct {
	ProductCode string
	Quantity    int
}

type CreateOrderCommandHandler struct {
	orderRepository    repositories.OrderRepository
	productRepository  repositories.ProductRepository
	campaignRepository repositories.CampaignRepository
	systemTime         *application.SystemTime
}

// NewCreateOrderCommandHandler initializes new CreateOrderCommandHandler
func NewCreateOrderCommandHandler(
	orderRepository repositories.OrderRepository,
	productRepository repositories.ProductRepository,
	campaignRepository repositories.CampaignRepository,
	systemTime *application.SystemTime) *CreateOrderCommandHandler {

	return &CreateOrderCommandHandler{
		orderRepository:    orderRepository,
		productRepository:  productRepository,
		campaignRepository: campaignRepository,
		systemTime:         systemTime,
	}
}

// Handle handles CreateOrderCommand
func (h *CreateOrderCommandHandler) Handle(ctx context.Context, c *CreateOrderCommand) (*createOrderResponse, error) {
	if h == nil {
		return nil, application.ThrowCreateOrderCommandHandlerCannotBeNilError()
	}

	if err := h.validate(); err != nil {
		return nil, err
	}

	if c == nil {
		return nil, application.ThrowCreateOrderCommandCannotBeNilError()
	}

	product, err := h.productRepository.GetProduct(ctx, c.ProductCode)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, application.ThrowProductNotFoundError(c.ProductCode)
	}
	if !product.InStock() {
		return nil, application.ThrowProductOutOfStockError(c.ProductCode)
	}

	if c.Quantity > product.Stock() {
		return nil, application.ThrowProductOutOfStockError(c.ProductCode)
	}

	campaign, err := h.campaignRepository.GetLatestCampaign(ctx, product.Code())
	if err != nil {
		return nil, err
	}
	if campaign == nil {
		return nil, application.ThrowCampaignNotFoundError(fmt.Sprintf("for %s product", c.ProductCode))
	}

	campaignService := services.NewCampaignService(campaign)

	err = campaignService.ApplyCampaignAndUpdateFields(product, c.Quantity, c.Quantity*product.Price(), h.systemTime.Time())
	if err != nil {
		return nil, err
	}

	order, err := domain.NewOrder(c.ProductCode, c.Quantity, product.Price())
	if err != nil {
		return nil, err
	}

	if err = h.orderRepository.AddOrder(ctx, order); err != nil {
		return nil, err
	}

	product.ReduceStock(order.Quantity())
	err = h.productRepository.UpdateProductStock(ctx, product)
	if err != nil {
		return nil, err
	}

	err = h.campaignRepository.UpdateCampaignTurnOverSales(ctx, campaign)
	if err != nil {
		return nil, err
	}

	return NewCreateOrderResponse(order.ProductCode(), order.Quantity()), nil
}

func (h *CreateOrderCommandHandler) validate() error {
	if h.campaignRepository == nil {
		return application.ThrowCreateOrderCommandHandlerCannotBeNilError()
	}
	if h.systemTime == nil {
		return application.ThrowSystemTimeCannotBeNilError()
	}
	if h.productRepository == nil {
		return application.ThrowProductRepositoryCannotBeNil()
	}
	if h.orderRepository == nil {
		return application.ThrowOrderRepositoryCannotBeNilError()
	}

	return nil
}

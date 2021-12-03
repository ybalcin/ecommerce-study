package getproductinfo

import (
	"context"
	"fmt"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
	"github.com/ybalcin/ecommerce-study/internal/domain/services"
)

type Query struct {
	Code string
}

type Handler struct {
	productRepository  repositories.ProductRepository
	campaignRepository repositories.CampaignRepository
	orderRepository    repositories.OrderRepository
	systemTime         *application.SystemTime
}

// NewHandler initializes NewGetProductInfoHandler
func NewHandler(
	productRepository repositories.ProductRepository,
	campaignRepository repositories.CampaignRepository,
	orderRepository repositories.OrderRepository,
	systemTime *application.SystemTime) *Handler {

	return &Handler{
		productRepository:  productRepository,
		campaignRepository: campaignRepository,
		orderRepository:    orderRepository,
		systemTime:         systemTime,
	}
}

// Handle handles Query
func (h *Handler) Handle(
	ctx context.Context,
	q *Query) (*response, error) {

	if h == nil {
		return nil, application.ThrowGetProductInfoQueryHandlerCannotNilError()
	}

	if err := h.validate(); err != nil {
		return nil, err
	}

	if q == nil {
		return nil, application.ThrowGetProductInfoQueryNilError()
	}

	product, err := h.productRepository.GetProduct(ctx, q.Code)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, application.ThrowProductNotFoundError(q.Code)
	}

	campaign, err := h.campaignRepository.GetLatestCampaign(ctx, product.Code())
	if err != nil {
		return nil, err
	}
	if campaign == nil {
		return nil, application.ThrowCampaignNotFoundError(fmt.Sprintf("for %s product", q.Code))
	}

	campaignService := services.NewCampaignService(campaign)

	if _, err := campaignService.ApplyCampaign(product, h.systemTime.Time()); err != nil {
		return nil, err
	}

	return NewResponse(product.Code(), product.Price(), product.Stock()), nil
}

func (h *Handler) validate() error {
	if h.productRepository == nil {
		return application.ThrowProductRepositoryCannotBeNil()
	}
	if h.campaignRepository == nil {
		return application.ThrowCampaignRepositoryCannotBeNilError()
	}
	if h.orderRepository == nil {
		return application.ThrowOrderRepositoryCannotBeNilError()
	}
	if h.systemTime == nil {
		return application.ThrowSystemTimeCannotBeNilError()
	}

	return nil
}

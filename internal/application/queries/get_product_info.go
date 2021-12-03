package queries

import (
	"context"
	"fmt"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
	"github.com/ybalcin/ecommerce-study/internal/domain/services"
)

type GetProductInfoQuery struct {
	Code string
}

type GetProductInfoQueryHandler struct {
	productRepository  repositories.ProductRepository
	campaignRepository repositories.CampaignRepository
	orderRepository    repositories.OrderRepository
	systemTime         *application.SystemTime
}

// NewGetProductInfoQueryHandler initializes NewGetProductInfoHandler
func NewGetProductInfoQueryHandler(
	productRepository repositories.ProductRepository,
	campaignRepository repositories.CampaignRepository,
	orderRepository repositories.OrderRepository,
	systemTime *application.SystemTime) *GetProductInfoQueryHandler {

	return &GetProductInfoQueryHandler{
		productRepository:  productRepository,
		campaignRepository: campaignRepository,
		orderRepository:    orderRepository,
		systemTime:         systemTime,
	}
}

// Handle handles GetProductInfoQuery
func (h *GetProductInfoQueryHandler) Handle(
	ctx context.Context,
	q *GetProductInfoQuery) (*getProductInfoQueryResponse, error) {

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

	if err := campaignService.ApplyCampaign(product, h.systemTime.Time()); err != nil {
		return nil, err
	}

	return NewGetProductInfoQueryResponse(product.Code(), product.Price(), product.Stock()), nil
}

func (h *GetProductInfoQueryHandler) validate() error {
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

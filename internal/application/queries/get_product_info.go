package queries

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
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

// NewGetProductInfoHandler initializes NewGetProductInfoHandler
func NewGetProductInfoHandler(
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

func (h *GetProductInfoQueryHandler) Handle(
	ctx context.Context,
	q *GetProductInfoQuery) (*getProductInfoQueryResponse, error) {

	product, err := h.productRepository.GetProduct(ctx, q.Code)
	if err != nil {
		return nil, err
	}

	campaign, err := h.campaignRepository.GetLatestCampaign(ctx, product.Code())
	if err != nil {
		return nil, err
	}

	if campaign != nil {
		if err := campaign.ApplyCampaign(product, h.systemTime.Time()); err != nil {
			return nil, err
		}
	}

	return NewGetProductInfoQueryResponse(product.Code(), product.Price(), product.Stock()), nil
}

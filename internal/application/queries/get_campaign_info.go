package queries

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
)

type GetCampaignInfoQuery struct {
	Name string
}

type GetCampaignInfoQueryHandler struct {
	campaignRepository repositories.CampaignRepository
	orderRepository    repositories.OrderRepository
	productRepository  repositories.ProductRepository
	systemTime         *application.SystemTime
}

// NewGetCampaignInfoQueryHandler initializes NewGetCampaignInfoQueryHandler
func NewGetCampaignInfoQueryHandler(
	campaignRepository repositories.CampaignRepository,
	orderRepository repositories.OrderRepository,
	productRepository repositories.ProductRepository,
	systemTime *application.SystemTime) *GetCampaignInfoQueryHandler {

	return &GetCampaignInfoQueryHandler{
		campaignRepository: campaignRepository,
		orderRepository:    orderRepository,
		productRepository:  productRepository,
		systemTime:         systemTime,
	}
}

// Handle handles GetCampaignInfoQuery
func (h *GetCampaignInfoQueryHandler) Handle(
	ctx context.Context,
	q *GetCampaignInfoQuery) (*getCampaignInfoResponse, error) {

	campaign, err := h.campaignRepository.GetCampaign(ctx, q.Name)
	if err != nil {
		return nil, err
	}

	orders, err := h.orderRepository.GetOrders(ctx, campaign.ProductCode())
	if err != nil {
		return nil, err
	}

	return NewGetCampaignInfoResponse(
		campaign.Name(),
		campaign.TargetSalesCount(),
		campaign.Status(h.systemTime.Time()),
		campaign.SalesCount(),
		campaign.TurnOver(),
		campaign.AverageSalePrice(orders)), nil
}

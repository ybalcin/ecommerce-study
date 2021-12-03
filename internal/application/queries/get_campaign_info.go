package queries

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
	"github.com/ybalcin/ecommerce-study/internal/domain/services"
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

	if h == nil {
		return nil, application.ThrowGetCampaignInfoQueryHandlerCannotBeNilError()
	}

	if err := h.validate(); err != nil {
		return nil, err
	}

	if q == nil {
		return nil, application.ThrowGetCampaignInfoQueryNilError()
	}

	campaign, err := h.campaignRepository.GetCampaign(ctx, q.Name)
	if err != nil {
		return nil, err
	}
	if campaign == nil {
		return nil, application.ThrowCampaignNotFoundError(q.Name)
	}

	orders, err := h.orderRepository.GetOrders(ctx, campaign.ProductCode())
	if err != nil {
		return nil, err
	}

	campaignService := services.NewCampaignService(campaign)

	return NewGetCampaignInfoResponse(
		campaign.Name(),
		campaign.TargetSalesCount(),
		campaign.Status(h.systemTime.Time()),
		campaign.SalesCount(),
		campaign.TurnOver(),
		campaignService.CalculateAverageSalePrice(orders)), nil
}

func (h *GetCampaignInfoQueryHandler) validate() error {
	if h.productRepository == nil {
		return application.ThrowProductRepositoryCannotBeNil()
	}
	if h.systemTime == nil {
		return application.ThrowSystemTimeCannotBeNilError()
	}
	if h.campaignRepository == nil {
		return application.ThrowCampaignRepositoryCannotBeNilError()
	}
	if h.orderRepository == nil {
		return application.ThrowOrderRepositoryCannotBeNilError()
	}

	return nil
}

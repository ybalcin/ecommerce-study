package getcampaigninfo

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
	"github.com/ybalcin/ecommerce-study/internal/domain/services"
)

type Query struct {
	Name string
}

type Handler struct {
	campaignRepository repositories.CampaignRepository
	orderRepository    repositories.OrderRepository
	systemTime         *application.SystemTime
}

// NewHandler initializes NewHandler
func NewHandler(
	campaignRepository repositories.CampaignRepository,
	orderRepository repositories.OrderRepository,
	systemTime *application.SystemTime) *Handler {

	return &Handler{
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

	return NewResponse(
		campaign.Name(),
		campaign.TargetSalesCount(),
		campaign.Status(h.systemTime.Time()),
		campaign.SalesCount(),
		campaign.TurnOver(),
		campaignService.CalculateAverageSalePrice(orders)), nil
}

func (h *Handler) validate() error {
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

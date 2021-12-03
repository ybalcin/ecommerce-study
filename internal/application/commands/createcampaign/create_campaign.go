package createcampaign

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
)

type Command struct {
	Name                   string
	ProductCode            string
	Duration               int
	PriceManipulationLimit int
	TargetSalesCount       int
}

type Handler struct {
	campaignRepository repositories.CampaignRepository
	systemTime         *application.SystemTime
}

// NewHandler initializes NewAddCampaignCommandHandler
func NewHandler(
	campaignRepository repositories.CampaignRepository,
	systemTime *application.SystemTime) *Handler {

	return &Handler{
		campaignRepository: campaignRepository,
		systemTime:         systemTime,
	}
}

// Handle handles CreateCampaignCommand
func (h *Handler) Handle(ctx context.Context, c *Command) (*response, error) {
	if h == nil {
		return nil, application.ThrowCreateCampaignCommandHandlerCannotBeNilError()
	}

	if c == nil {
		return nil, application.ThrowCreateCampaignCommandCannotNilError()
	}

	if err := h.validate(); err != nil {
		return nil, err
	}

	campaign, err := domain.NewCampaign(
		"",
		c.Name,
		c.ProductCode,
		c.Duration,
		c.PriceManipulationLimit,
		c.TargetSalesCount,
		0,
		0,
		h.systemTime.Time())

	if err != nil {
		return nil, err
	}

	if err := h.campaignRepository.AddCampaign(ctx, campaign); err != nil {
		return nil, err
	}

	return NewResponse(
		campaign.Name(),
		campaign.ProductCode(),
		campaign.Duration(),
		campaign.PriceManipulationLimit(),
		campaign.TargetSalesCount()), nil
}

func (h *Handler) validate() error {
	if h.campaignRepository == nil {
		return application.ThrowCampaignRepositoryCannotBeNilError()
	}
	if h.systemTime == nil {
		return application.ThrowSystemTimeCannotBeNilError()
	}

	return nil
}

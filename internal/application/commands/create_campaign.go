package commands

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
)

type CreateCampaignCommand struct {
	Name                   string
	ProductCode            string
	Duration               int
	PriceManipulationLimit int
	TargetSalesCount       int
}

type CreateCampaignCommandHandler struct {
	campaignRepository repositories.CampaignRepository
	systemTime         *application.SystemTime
}

// NewCreateCampaignCommandHandler initializes NewAddCampaignCommandHandler
func NewCreateCampaignCommandHandler(
	campaignRepository repositories.CampaignRepository,
	systemTime *application.SystemTime) *CreateCampaignCommandHandler {

	return &CreateCampaignCommandHandler{
		campaignRepository: campaignRepository,
		systemTime:         systemTime,
	}
}

// Handle handles CreateCampaignCommand
func (h *CreateCampaignCommandHandler) Handle(ctx context.Context, c *CreateCampaignCommand) (*createCampaignResponse, error) {
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

	return NewCreateCampaignResponse(
		campaign.Name(),
		campaign.ProductCode(),
		campaign.Duration(),
		campaign.PriceManipulationLimit(),
		campaign.TargetSalesCount()), nil
}

func (h *CreateCampaignCommandHandler) validate() error {
	if h.campaignRepository == nil {
		return application.ThrowCampaignRepositoryCannotBeNilError()
	}
	if h.systemTime == nil {
		return application.ThrowSystemTimeCannotBeNilError()
	}

	return nil
}

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

// NewAddCampaignCommandHandler initializes NewAddCampaignCommandHandler
func NewAddCampaignCommandHandler(
	campaignRepository repositories.CampaignRepository,
	systemTime *application.SystemTime) *CreateCampaignCommandHandler {

	return &CreateCampaignCommandHandler{
		campaignRepository: campaignRepository,
		systemTime:         systemTime,
	}
}

// Handle handles CreateCampaignCommand
func (h *CreateCampaignCommandHandler) Handle(ctx context.Context, c *CreateCampaignCommand) error {
	if h == nil {
		return application.ThrowCreateCampaignCommandHandlerCannotBeNilError()
	}

	if c == nil {
		return application.ThrowCreateCampaignCommandCannotNilError()
	}

	if err := h.validate(); err != nil {
		return err
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
		return err
	}

	if err := h.campaignRepository.AddCampaign(ctx, campaign); err != nil {
		return err
	}

	return nil
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

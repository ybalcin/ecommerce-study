package commands

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
	"time"
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
}

// NewAddCampaignCommandHandler initializes NewAddCampaignCommandHandler
func NewAddCampaignCommandHandler(campaignRepository repositories.CampaignRepository) *CreateCampaignCommandHandler {
	return &CreateCampaignCommandHandler{
		campaignRepository: campaignRepository,
	}
}

// Handle handles CreateCampaignCommand
func (h *CreateCampaignCommandHandler) Handle(ctx context.Context, c *CreateCampaignCommand) error {
	campaign, err := domain.NewCampaign(
		"",
		c.Name,
		c.ProductCode,
		c.Duration,
		c.PriceManipulationLimit,
		c.TargetSalesCount,
		0,
		0,
		time.Now().UTC())

	if err != nil {
		return err
	}

	if err := h.campaignRepository.AddCampaign(ctx, campaign); err != nil {
		return err
	}

	return nil
}

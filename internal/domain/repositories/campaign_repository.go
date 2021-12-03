package repositories

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/domain"
)

// CampaignRepository interface wraps campaign operations
type CampaignRepository interface {
	// AddCampaign adds campaign to collection
	AddCampaign(ctx context.Context, campaign *domain.Campaign) error

	// GetCampaign gets campaign info by name
	GetCampaign(ctx context.Context, name string) (*domain.Campaign, error)

	// GetLatestCampaign gets the latest campaign of product
	GetLatestCampaign(ctx context.Context, productCode string) (*domain.Campaign, error)

	// UpdateCampaignTurnOverSales updates campaign turn over and sales count
	UpdateCampaignTurnOverSales(ctx context.Context, campaign *domain.Campaign) error

	// DropCampaigns deletes all campaigns
	DropCampaigns(ctx context.Context) error
}

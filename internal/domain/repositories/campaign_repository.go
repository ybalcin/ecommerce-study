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

	// UpdateCampaign updates campaign
	UpdateCampaign(ctx context.Context, campaign *domain.Campaign) error
}

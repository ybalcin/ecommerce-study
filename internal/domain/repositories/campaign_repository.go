package repositories

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/domain/models"
)

// CampaignRepository interface wraps campaign operations
type CampaignRepository interface {
	// AddCampaign adds campaign to collection
	AddCampaign(ctx context.Context, campaign *models.Campaign) error
}

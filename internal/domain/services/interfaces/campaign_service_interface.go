package interfaces

import (
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"time"
)

type CampaignService interface {
	ApplyCampaignAndUpdateFields(product *domain.Product, orderQuantity, orderTotalPrice int, systemTime time.Time) error
	ApplyCampaign(product *domain.Product, systemTime time.Time) (bool, error)
	CalculateAverageSalePrice(orders []domain.Order) int
}

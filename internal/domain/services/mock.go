package services

import (
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"time"
)

type MockCampaignService struct {
	ApplyCampaignAndUpdateFieldsFn      func(product *domain.Product, orderQuantity, orderTotalPrice int, systemTime time.Time) error
	ApplyCampaignAndUpdateFieldsInvoked bool

	ApplyCampaignFn      func(product *domain.Product, systemTime time.Time) (bool, error)
	ApplyCampaignInvoked bool

	CalculateAverageSalePriceFn      func(orders []domain.Order) int
	CalculateAverageSalePriceInvoked bool
}

func (s MockCampaignService) ApplyCampaignAndUpdateFields(product *domain.Product, orderQuantity, orderTotalPrice int, systemTime time.Time) error {
	s.ApplyCampaignAndUpdateFieldsInvoked = true
	return s.ApplyCampaignAndUpdateFieldsFn(product, orderQuantity, orderTotalPrice, systemTime)
}

func (s MockCampaignService) ApplyCampaign(product *domain.Product, systemTime time.Time) (bool, error) {
	s.ApplyCampaignInvoked = true
	return s.ApplyCampaignFn(product, systemTime)
}

func (s MockCampaignService) CalculateAverageSalePrice(orders []domain.Order) int {
	s.CalculateAverageSalePriceInvoked = true
	return s.CalculateAverageSalePriceFn(orders)
}

package services

import (
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/internal/domain/errors"
	"time"
)

// CampaignService is campaign domain service
type CampaignService struct {
	campaign *domain.Campaign
}

// NewCampaignService initializes new campaign service
func NewCampaignService(campaign *domain.Campaign) *CampaignService {
	return &CampaignService{
		campaign: campaign,
	}
}

// ApplyCampaignAndUpdateFields applies campaign and updates sales count, turnover
func (s *CampaignService) ApplyCampaignAndUpdateFields(
	product *domain.Product,
	orderQuantity,
	orderTotalPrice int,
	systemTime time.Time) error {

	if err := s.ApplyCampaign(product, systemTime); err != nil {
		return err
	}

	s.campaign.UpdateSalesCount(orderQuantity)
	s.campaign.UpdateTurnOver(orderTotalPrice)

	return nil
}

// ApplyCampaign applies campaign to product
func (s *CampaignService) ApplyCampaign(product *domain.Product, systemTime time.Time) error {
	if !s.campaign.IsActive(systemTime) {
		return nil
	}

	if s.campaign.ProductCode() != product.Code() {
		return errors.ThrowCampaignApplyProductCodesNotEqualError()
	}

	discountRate := s.campaign.CalculateDiscountRate(systemTime)
	applyCampaign(product, discountRate)

	return nil
}

// CalculateAverageSalePrice calculates average sale price of campaign
func (s *CampaignService) CalculateAverageSalePrice(orders []domain.Order) int {
	if orders == nil {
		return 0
	}

	totalOrder := len(orders)
	var unitPriceSum int

	for _, o := range orders {
		unitPriceSum += o.UnitSalePrice()
	}

	return unitPriceSum / totalOrder
}

func applyCampaign(product *domain.Product, discountRate int) {
	newPrice := product.Price() - product.Price()*discountRate/100
	product.ApplyPrice(newPrice)
}

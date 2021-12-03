package commands

import "fmt"

type createCampaignResponse struct {
	Name                   string
	ProductCode            string
	Duration               int
	PriceManipulationLimit int
	TargetSalesCount       int
}

func NewCreateCampaignResponse(name, productCode string, duration, limit, targetSalesCount int) *createCampaignResponse {
	return &createCampaignResponse{
		Name:                   name,
		ProductCode:            productCode,
		Duration:               duration,
		PriceManipulationLimit: limit,
		TargetSalesCount:       targetSalesCount,
	}
}

func (r *createCampaignResponse) String() string {
	return fmt.Sprintf("Campaign created; name %s, product %s, duration %d, limit %d, target sales count %d",
		r.Name,
		r.ProductCode,
		r.Duration,
		r.PriceManipulationLimit,
		r.TargetSalesCount)
}

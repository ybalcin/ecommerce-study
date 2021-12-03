package createcampaign

import "fmt"

type response struct {
	Name                   string
	ProductCode            string
	Duration               int
	PriceManipulationLimit int
	TargetSalesCount       int
}

func NewResponse(name, productCode string, duration, limit, targetSalesCount int) *response {
	return &response{
		Name:                   name,
		ProductCode:            productCode,
		Duration:               duration,
		PriceManipulationLimit: limit,
		TargetSalesCount:       targetSalesCount,
	}
}

func (r *response) String() string {
	return fmt.Sprintf("Campaign created; name %s, product %s, duration %d, limit %d, target sales count %d",
		r.Name,
		r.ProductCode,
		r.Duration,
		r.PriceManipulationLimit,
		r.TargetSalesCount)
}

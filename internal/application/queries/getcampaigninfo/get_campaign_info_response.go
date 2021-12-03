package getcampaigninfo

import "fmt"

type response struct {
	Name             string
	TargetSalesCount int
	Status           string
	TotalSales       int
	TurnOver         int
	AverageItemPrice int
}

// NewResponse initializes new NewResponse
func NewResponse(
	name string,
	targetSalesCount int,
	status string,
	totalSales int,
	turnOver int,
	averageItemPrice int) *response {

	return &response{
		Name:             name,
		TargetSalesCount: targetSalesCount,
		Status:           status,
		TotalSales:       totalSales,
		TurnOver:         turnOver,
		AverageItemPrice: averageItemPrice,
	}
}

func (r *response) String() string {
	return fmt.Sprintf("Campaign %s info; Status %s, Target Sales %d, Total Sales %d, Turnover %d, Average Item Price %d",
		r.Name, r.Status, r.TargetSalesCount, r.TotalSales, r.TurnOver, r.AverageItemPrice)
}

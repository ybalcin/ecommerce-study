package queries

import "fmt"

type getCampaignInfoResponse struct {
	Name             string
	TargetSalesCount int
	Status           string
	TotalSales       int
	TurnOver         int
	AverageItemPrice int
}

// NewGetCampaignInfoResponse initializes new NewGetCampaignInfoResponse
func NewGetCampaignInfoResponse(
	name string,
	targetSalesCount int,
	status string,
	totalSales int,
	turnOver int,
	averageItemPrice int) *getCampaignInfoResponse {

	return &getCampaignInfoResponse{
		Name:             name,
		TargetSalesCount: targetSalesCount,
		Status:           status,
		TotalSales:       totalSales,
		TurnOver:         turnOver,
		AverageItemPrice: averageItemPrice,
	}
}

func (r *getCampaignInfoResponse) String() string {
	return fmt.Sprintf("Campaign %s info; Status %s, Target Sales %d, Total Sales %d, Turnover %d, Average Item Price %d",
		r.Name, r.Status, r.TargetSalesCount, r.TotalSales, r.TurnOver, r.AverageItemPrice)
}

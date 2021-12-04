package services_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/internal/domain/services"
	"testing"
	"time"
)

const dummy = "dummy"

func newSysTime() time.Time {
	var now = time.Now().UTC()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}

func TestCampaignService_ApplyCampaign(t *testing.T) {
	testCases := []struct {
		productCode            string
		price                  int
		priceManipulationLimit int
		duration               int
		sysTime                time.Time
		expected               int
		ok                     bool
	}{
		{"p1", 100, 20, 5, newSysTime(), 95, true},
		{"p1", 100, 20, 5, newSysTime().Add(time.Hour), 90, true},
		{"p2", 100, 20, 5, newSysTime().Add(time.Hour), 100, false},
		{"p1", 100, 20, 5, newSysTime().Add(time.Hour * time.Duration(5)), 100, false},
	}

	for _, c := range testCases {
		product, err := domain.NewProduct("", "p1", 100, 1)
		assert.Nil(t, err)

		campaign, err := domain.NewCampaign(
			"",
			dummy,
			c.productCode,
			c.duration,
			c.priceManipulationLimit,
			2,
			1,
			1,
			time.Now())

		campaignService := services.NewCampaignService(campaign)

		ok, err := campaignService.ApplyCampaign(product, c.sysTime)

		assert.Equal(t, c.expected, product.Price())
		assert.Equal(t, c.ok, ok)
	}
}

func TestCampaignService_CalculateAverageSalePrice(t *testing.T) {
	testCases := []struct {
		orderLength   int
		unitSalePrice int
		expected      int
	}{
		{2, 10, 10},
		{0, 10, 0},
		{1, 10, 10},
	}

	for _, c := range testCases {
		var orders []domain.Order
		for i := 0; i < c.orderLength; i++ {
			order, err := domain.NewOrder("p1", 1, c.unitSalePrice)
			assert.Nil(t, err)

			orders = append(orders, *order)
		}

		campaign, err := domain.NewCampaign(
			"",
			dummy,
			"asd",
			2,
			1,
			2,
			1,
			1,
			time.Now())
		assert.Nil(t, err)

		campaignService := services.NewCampaignService(campaign)

		actual := campaignService.CalculateAverageSalePrice(orders)
		assert.Equal(t, c.expected, actual)
	}
}

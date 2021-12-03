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
	}{
		{"p1", 100, 20, 5, newSysTime(), 95},
		{"p1", 100, 20, 5, newSysTime().Add(time.Hour), 90},
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

		err = campaignService.ApplyCampaign(product, c.sysTime)
		assert.Nil(t, err)
		assert.Equal(t, c.expected, product.Price())
	}
}

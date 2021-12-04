package createcampaign_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/createcampaign"
	"testing"
)

func TestBuild(t *testing.T) {
	testCases := []struct {
		t        string
		cmd      string
		expected *createcampaign.Command
	}{
		{
			"cmd valid",
			"create_campaign C11 P11 10 20 100",
			&createcampaign.Command{
				Name:                   "C11",
				ProductCode:            "P11",
				Duration:               10,
				PriceManipulationLimit: 20,
				TargetSalesCount:       100,
			},
		},
		{
			"cmd invalid",
			"create_campaign",
			&createcampaign.Command{
				Name:                   "",
				ProductCode:            "",
				Duration:               0,
				PriceManipulationLimit: 0,
				TargetSalesCount:       0,
			},
		},
		{
			"entry not equal",
			"dummy",
			nil,
		},
	}

	for _, c := range testCases {
		t.Run(c.t, func(t *testing.T) {
			actual := createcampaign.Build(c.cmd)
			assert.Equal(t, c.expected, actual)
		})
	}
}

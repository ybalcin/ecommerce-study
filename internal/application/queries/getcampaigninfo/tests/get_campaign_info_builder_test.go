package getcampaigninfo_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/application/queries/getcampaigninfo"
	"testing"
)

func TestBuild(t *testing.T) {
	testCases := []struct {
		t        string
		q        string
		expected *getcampaigninfo.Query
	}{
		{
			"cmd valid",
			"get_campaign_info C11",
			&getcampaigninfo.Query{Name: "C11"},
		},
		{
			"cmd invalid",
			"get_campaign_info",
			&getcampaigninfo.Query{Name: ""},
		},
		{
			"entry not equal",
			"dummy",
			nil,
		},
		{
			"query empty",
			"",
			nil,
		},
	}

	for _, c := range testCases {
		t.Run(c.t, func(t *testing.T) {
			actual := getcampaigninfo.Build(c.q)
			assert.Equal(t, c.expected, actual)
		})
	}
}

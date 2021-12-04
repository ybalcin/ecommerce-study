package getproductinfo_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/application/queries/getproductinfo"
	"testing"
)

func TestBuild(t *testing.T) {
	testCases := []struct {
		t        string
		q        string
		expected *getproductinfo.Query
	}{
		{
			"cmd valid",
			"get_product_info P11",
			&getproductinfo.Query{Code: "P11"},
		},
		{
			"cmd invalid",
			"get_product_info",
			&getproductinfo.Query{Code: ""},
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
			actual := getproductinfo.Build(c.q)
			assert.Equal(t, c.expected, actual)
		})
	}
}

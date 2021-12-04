package createproduct_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/createproduct"
	"testing"
)

func TestBuild(t *testing.T) {
	testCases := []struct {
		t        string
		cmd      string
		expected *createproduct.Command
	}{
		{
			"cmd valid",
			"create_product P11 100 1000",
			&createproduct.Command{
				ProductCode: "P11",
				Price:       100,
				Stock:       1000,
			},
		},
		{
			"cmd invalid",
			"create_product",
			&createproduct.Command{
				ProductCode: "",
				Price:       0,
				Stock:       0,
			},
		},
		{
			"entry not equal",
			"dummy",
			nil,
		},
		{
			"cmd empty",
			"",
			nil,
		},
	}

	for _, c := range testCases {
		t.Run(c.t, func(t *testing.T) {
			actual := createproduct.Build(c.cmd)
			assert.Equal(t, c.expected, actual)
		})
	}
}

package createorder_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/createorder"
	"testing"
)

func TestBuild(t *testing.T) {
	testCases := []struct {
		t        string
		cmd      string
		expected *createorder.Command
	}{
		{
			"cmd valid",
			"create_order P11 10",
			&createorder.Command{
				ProductCode: "P11",
				Quantity:    10,
			},
		},
		{
			"cmd invalid",
			"create_order",
			&createorder.Command{
				ProductCode: "",
				Quantity:    0,
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
			actual := createorder.Build(c.cmd)
			assert.Equal(t, c.expected, actual)
		})
	}
}

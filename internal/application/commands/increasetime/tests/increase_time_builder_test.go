package increasetime

import (
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/increasetime"
	"testing"
)

func TestBuild(t *testing.T) {
	testCases := []struct {
		t        string
		cmd      string
		expected *increasetime.Command
	}{
		{
			"cmd valid",
			"increase_time 1",
			&increasetime.Command{Hours: 1},
		},
		{
			"cmd invalid",
			"increase_time",
			&increasetime.Command{Hours: 0},
		},
		{
			"entry not equal",
			"dummy",
			nil,
		},
	}

	for _, c := range testCases {
		t.Run(c.t, func(t *testing.T) {
			actual := increasetime.Build(c.cmd)
			assert.Equal(t, c.expected, actual)
		})
	}
}

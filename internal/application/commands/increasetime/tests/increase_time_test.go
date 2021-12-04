package increasetime_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/increasetime"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	testCases := []struct {
		t    string
		h    *increasetime.Handler
		c    *increasetime.Command
		fail bool
	}{
		{
			"handler nil",
			nil,
			new(increasetime.Command),
			true,
		},
		{
			"sysTime nil",
			increasetime.NewHandler(nil),
			new(increasetime.Command),
			true,
		},
		{
			"command nil",
			increasetime.NewHandler(new(application.SystemTime)),
			nil,
			true,
		},
	}

	for _, c := range testCases {
		t.Run(c.t, func(t *testing.T) {
			resp, err := c.h.Handle(c.c)
			if c.fail {
				assert.Nil(t, resp)
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, resp)
			}
		})
	}
}

package domain_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"testing"
	"time"
)

const dummy = "dummy"

func newSysTime() time.Time {
	var now = time.Now().UTC()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
}

func TestCampaign_IsActive(t *testing.T) {
	testCases := []struct {
		duration  int
		sysTime   time.Time
		createdAt time.Time
		expected  bool
	}{
		{1, newSysTime(), newSysTime(), true},
		{1, newSysTime().Add(time.Hour), newSysTime(), false},
		{2, newSysTime().Add(time.Hour), newSysTime(), true},
	}

	for _, c := range testCases {
		campaign, err := domain.NewCampaign(
			dummy,
			dummy,
			dummy,
			c.duration,
			1,
			1,
			0,
			0,
			c.createdAt)
		assert.Nil(t, err)
		actual := campaign.IsActive(c.sysTime)
		assert.Equal(t, c.expected, actual)
	}
}

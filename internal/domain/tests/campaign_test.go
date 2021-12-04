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

func TestCampaign_CalculateDiscountRate(t *testing.T) {
	testCases := []struct {
		priceManipulationLimit int
		duration               int
		systemTime             time.Time
		expected               int
	}{
		{20, 5, newSysTime(), 5},
		{20, 5, newSysTime().Add(time.Hour), 10},
		{20, 5, newSysTime().Add(time.Hour * time.Duration(2)), 15},
		{20, 5, newSysTime().Add(time.Hour * time.Duration(3)), 20},
		{20, 5, newSysTime().Add(time.Hour * time.Duration(4)), 20},
		{20, 5, newSysTime().Add(time.Hour * time.Duration(5)), 20},
	}

	for _, c := range testCases {
		campaign, err := domain.NewCampaign(
			dummy,
			dummy,
			dummy,
			c.duration,
			c.priceManipulationLimit,
			1,
			0,
			0,
			time.Now())
		assert.Nil(t, err)

		actual := campaign.CalculateDiscountRate(c.systemTime)
		assert.Equal(t, c.expected, actual)
	}
}

func TestNewCampaign(t *testing.T) {
	testCases := []struct {
		id                     string
		name                   string
		productCode            string
		duration               int
		priceManipulationLimit int
		targetSalesCount       int
		salesCount             int
		turnOver               int
		createdAt              time.Time
		fail                   bool
	}{
		{dummy, "", dummy, 1, 1, 1, 1, 1, time.Now(), true},
		{dummy, dummy, "", 1, 1, 1, 1, 1, time.Now(), true},
		{dummy, dummy, dummy, 0, 1, 1, 1, 1, time.Now(), true},
		{dummy, dummy, dummy, 1, 0, 1, 1, 1, time.Now(), true},
		{dummy, dummy, dummy, 1, 1, 0, 1, 1, time.Now(), true},
		{dummy, dummy, dummy, 1, 1, 1, 0, 1, time.Now(), false},
		{dummy, dummy, dummy, 1, 1, 1, 1, 0, time.Now(), false},
		{dummy, dummy, dummy, 1, 1, 1, 1, 1, time.Now(), false},
		{dummy, dummy, dummy, 1, 100, 1, 1, 1, time.Now(), false},
		{dummy, dummy, dummy, 1, 101, 1, 1, 1, time.Now(), true},
	}

	for _, c := range testCases {
		actual, err := domain.NewCampaign(c.id, c.name, c.productCode, c.duration, c.priceManipulationLimit, c.targetSalesCount,
			c.salesCount, c.turnOver, c.createdAt)

		if c.fail {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.NotNil(t, actual)
		}
	}
}

func TestCampaign_Status(t *testing.T) {
	testCases := []struct {
		sysTime  time.Time
		expected string
	}{
		{newSysTime(), domain.CampaignStatusActive},
		{newSysTime().Add(time.Hour), domain.CampaignStatusEnded},
	}

	for _, c := range testCases {
		campaign, err := domain.NewCampaign(dummy, dummy, dummy, 1, 1, 1,
			0, 1, time.Now())

		assert.Nil(t, err)

		actual := campaign.Status(c.sysTime)
		assert.Equal(t, c.expected, actual)
	}
}

func TestCampaign_TargetFulfilled(t *testing.T) {
	testCases := []struct {
		targetCount int
		saleCount   int
		expected    bool
	}{
		{10, 10, true},
		{10, 15, true},
		{10, 1, false},
	}

	for _, c := range testCases {
		campaign, err := domain.NewCampaign(dummy, dummy, dummy, 1, 1, c.targetCount,
			c.saleCount, 1, time.Now())
		assert.Nil(t, err)
		actual := campaign.TargetFulfilled()
		assert.Equal(t, actual, c.expected)
	}
}

func TestCampaign_UpdateSalesCount(t *testing.T) {
	testCases := []struct {
		val       int
		saleCount int
	}{
		{1, 2},
		{0, 1},
	}

	for _, c := range testCases {
		campaign, err := domain.NewCampaign(dummy, dummy, dummy, 1, 1, 100,
			c.saleCount, 1, time.Now())
		assert.Nil(t, err)
		campaign.UpdateSalesCount(c.val)
		assert.Equal(t, campaign.SalesCount(), c.val+c.saleCount)
	}
}

func TestCampaign_UpdateTurnOver(t *testing.T) {
	testCases := []struct {
		val      int
		turnOver int
	}{
		{1, 2},
		{0, 1},
	}

	for _, c := range testCases {
		campaign, err := domain.NewCampaign(dummy, dummy, dummy, 1, 1, 100,
			5, c.turnOver, time.Now())
		assert.Nil(t, err)
		campaign.UpdateTurnOver(c.val)
		assert.Equal(t, campaign.TurnOver(), c.val+c.turnOver)
	}
}

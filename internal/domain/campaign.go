package domain

import (
	"github.com/ybalcin/ecommerce-study/internal/domain/errors"
	"time"
)

// Campaign model
// duration as hours
// priceManipulationLimit as percentage
type Campaign struct {
	id                     string
	name                   string
	productCode            string
	duration               int
	priceManipulationLimit int
	targetSalesCount       int
	salesCount             int
	turnOver               int
	createdAt              time.Time
}

const (
	campaignStatusActive string = "Active"
	campaignStatusEnded  string = "Ended"
)

// validate validates model
func (c *Campaign) validate() error {
	if c.name == "" {
		return errors.ThrowCampaignNameShouldNotBeEmptyError()
	}
	if c.productCode == "" {
		return errors.ThrowProductCodeShouldNotBeEmptyError()
	}
	if c.duration <= 0 {
		return errors.ThrowCampaignDurationIsInvalidError()
	}
	if c.priceManipulationLimit <= 0 || c.priceManipulationLimit > 100 {
		return errors.ThrowCampaignPriceManipulationLimitIsInvalidError()
	}

	return nil
}

// NewCampaign initializes new campaign
func NewCampaign(
	id,
	name,
	productCode string,
	duration,
	priceManipulationLimit,
	targetSalesCount,
	salesCount,
	turnOver int,
	createdAt time.Time) (*Campaign, error) {

	campaign := &Campaign{
		id:                     id,
		name:                   name,
		productCode:            productCode,
		duration:               duration,
		priceManipulationLimit: priceManipulationLimit,
		targetSalesCount:       targetSalesCount,
		salesCount:             salesCount,
		turnOver:               turnOver,
		createdAt:              createdAt,
	}

	if err := campaign.validate(); err != nil {
		return nil, err
	}

	return campaign, nil
}

// Id returns id of campaign
func (c *Campaign) Id() string {
	return c.id
}

// Name returns name of campaign
func (c *Campaign) Name() string {
	return c.name
}

// ProductCode returns product code of campaign
func (c *Campaign) ProductCode() string {
	return c.productCode
}

// Duration returns duration of campaign
func (c *Campaign) Duration() int {
	return c.duration
}

// PriceManipulationLimit returns price manipulation limit of campaign
func (c *Campaign) PriceManipulationLimit() int {
	return c.priceManipulationLimit
}

// TargetSalesCount returns target sales count of campaign
func (c *Campaign) TargetSalesCount() int {
	return c.targetSalesCount
}

// SalesCount returns sales count of campaign
func (c *Campaign) SalesCount() int {
	return c.salesCount
}

// CreatedAt returns created time of campaign
func (c *Campaign) CreatedAt() time.Time {
	return c.createdAt
}

// IsActive returns campaign current status
func (c *Campaign) IsActive(systemTime time.Time) bool {
	if c.TargetFulfilled() {
		return false
	}

	if c.duration > systemTime.Hour() {
		return true
	}

	return false
}

// Status returns current status of campaign
func (c *Campaign) Status(systemTime time.Time) string {
	if c.IsActive(systemTime) {
		return campaignStatusActive
	}

	return campaignStatusEnded
}

// TurnOver returns turnover price of campaign
func (c *Campaign) TurnOver() int {
	return c.turnOver
}

// TargetFulfilled returns whether the target sales limit has been fulfilled
func (c *Campaign) TargetFulfilled() bool {
	if c.salesCount >= c.targetSalesCount {
		return true
	}

	return false
}

// CalculateDiscountRate calculates discount rate
func (c *Campaign) CalculateDiscountRate(systemTime time.Time) int {
	discountTimePart := c.priceManipulationLimit / c.duration
	calculatedDiscountRate := (c.priceManipulationLimit / discountTimePart) * (systemTime.Hour() + 1)

	if calculatedDiscountRate > c.priceManipulationLimit {
		return c.priceManipulationLimit
	}

	return calculatedDiscountRate
}

// UpdateSalesCount updates sales count
func (c *Campaign) UpdateSalesCount(val int) {
	c.salesCount += val
}

// UpdateTurnOver updates turnover
func (c *Campaign) UpdateTurnOver(val int) {
	c.turnOver += val
}

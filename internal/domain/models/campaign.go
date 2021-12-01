package models

import "fmt"

type (
	CampaignId string

	// Campaign entity
	// duration as hours
	// priceManipulationLimit as percentage
	Campaign struct {
		id                     CampaignId `bson:"_id"`
		name                   string     `bson:"name"`
		productCode            string     `bson:"product_code"`
		duration               int        `bson:"duration"`
		priceManipulationLimit int        `bson:"price_manipulation_limit"`
		targetSalesCount       int        `bson:"target_sales_count"`
	}
)

// NewCampaign initializes new campaign
func NewCampaign(id CampaignId, name, productCode string, duration, priceManipulationLimit, targetSalesCount int) (*Campaign, error) {
	campaign := &Campaign{
		id:                     id,
		name:                   name,
		productCode:            productCode,
		duration:               duration,
		priceManipulationLimit: priceManipulationLimit,
		targetSalesCount:       targetSalesCount,
	}

	if err := campaign.validate(); err != nil {
		return nil, err
	}

	return campaign, nil
}

// Id returns id of campaign
func (c *Campaign) Id() CampaignId {
	return c.id
}

// String converts id to string
func (id CampaignId) String() string {
	return string(id)
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

// validate validates model
func (c *Campaign) validate() error {
	invalidErr := func(key string) error {
		return fmt.Errorf("models: campaign %s is invalid", key)
	}

	if c.name == "" {
		return invalidErr("name")
	}
	if c.id == "" {
		return invalidErr("id")
	}
	if c.productCode == "" {
		return invalidErr("productCode")
	}
	if c.duration <= 0 {
		return invalidErr("duration")
	}
	if c.priceManipulationLimit <= 0 || c.priceManipulationLimit > 100 {
		return invalidErr("priceManipulationLimit")
	}

	return nil
}

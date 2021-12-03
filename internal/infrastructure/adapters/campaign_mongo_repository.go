package adapters

import (
	"context"
	"errors"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/pkg/mgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type campaignBson struct {
	Id                     primitive.ObjectID `bson:"_id"`
	Name                   string             `bson:"name"`
	ProductCode            string             `bson:"product_code"`
	Duration               int                `bson:"duration"`
	PriceManipulationLimit int                `bson:"price_manipulation_limit"`
	TargetSalesCount       int                `bson:"target_sales_count"`
	SalesCount             int                `bson:"sales_count"`
	TurnOver               int                `bson:"turn_over"`
	CreatedAt              time.Time          `bson:"created_at"`
}

// campaignRepository implements repositories.CampaignRepository
type campaignRepository struct {
	campaigns *mgo.Collection
}

const campaignCollection string = "campaigns"

func (c *campaignBson) hasValue() bool {
	if c == nil {
		return false
	}

	return c.Id != primitive.NilObjectID
}

// NewCampaignRepository initializes new campaign repository
func NewCampaignRepository(store *mgo.Store) (*campaignRepository, error) {
	if store == nil {
		return nil, errors.New("adapters: mgo store is nil")
	}

	return &campaignRepository{
		campaigns: store.Collection(campaignCollection),
	}, nil
}

// AddCampaign adds campaign to collection
func (r *campaignRepository) AddCampaign(ctx context.Context, campaign *domain.Campaign) error {
	campaignBson := campaignBson{
		Id:                     primitive.NewObjectID(),
		Name:                   campaign.Name(),
		ProductCode:            campaign.ProductCode(),
		Duration:               campaign.Duration(),
		PriceManipulationLimit: campaign.PriceManipulationLimit(),
		TargetSalesCount:       campaign.TargetSalesCount(),
		SalesCount:             campaign.SalesCount(),
		TurnOver:               campaign.TurnOver(),
		CreatedAt:              campaign.CreatedAt(),
	}

	if _, err := r.campaigns.InsertOne(ctx, campaignBson); err != nil {
		return err
	}

	return nil
}

// GetCampaign gets campaign info
func (r *campaignRepository) GetCampaign(ctx context.Context, name string) (*domain.Campaign, error) {
	filter := bson.M{
		"name": name,
	}

	campaignBson := new(campaignBson)

	if err := r.campaigns.FindOne(ctx, filter, campaignBson); err != nil {
		return nil, err
	}

	if !campaignBson.hasValue() {
		return nil, nil
	}

	campaign, err := campaignBson.mapToCampaign()
	if err != nil {
		return nil, err
	}

	return campaign, nil
}

// GetLatestCampaign gets the latest campaign of product
func (r *campaignRepository) GetLatestCampaign(ctx context.Context, productCode string) (*domain.Campaign, error) {
	filterStage := bson.D{
		{"product_code", productCode},
	}

	sortStage := bson.D{
		{"$sort", bson.M{
			"created_at": -1,
		}},
	}

	campaignBson := new(campaignBson)

	if err := r.campaigns.FindOne(ctx, mongo.Pipeline{filterStage, sortStage}, campaignBson); err != nil {
		return nil, err
	}

	if !campaignBson.hasValue() {
		return nil, nil
	}

	campaign, err := campaignBson.mapToCampaign()
	if err != nil {
		return nil, err
	}

	return campaign, nil
}

// UpdateCampaign updates campaign
func (r *campaignRepository) UpdateCampaign(ctx context.Context, campaign *domain.Campaign) error {
	id, err := primitive.ObjectIDFromHex(campaign.Id())
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": id,
	}

	campaignBson, err := mapToCampaignBson(campaign)
	if err != nil {
		return err
	}

	if err = r.campaigns.UpdateOne(ctx, filter, campaignBson); err != nil {
		return err
	}

	return nil
}

func (c *campaignBson) mapToCampaign() (*domain.Campaign, error) {
	campaign, err := domain.NewCampaign(
		c.Id.Hex(),
		c.Name,
		c.ProductCode,
		c.Duration,
		c.PriceManipulationLimit,
		c.TargetSalesCount,
		c.SalesCount,
		c.TurnOver,
		c.CreatedAt)

	if err != nil {
		return nil, err
	}

	return campaign, nil
}

func mapToCampaignBson(campaign *domain.Campaign) (*campaignBson, error) {
	id, err := primitive.ObjectIDFromHex(campaign.Id())
	if err != nil {
		return nil, err
	}

	return &campaignBson{
		Id:                     id,
		Name:                   campaign.Name(),
		ProductCode:            campaign.ProductCode(),
		Duration:               campaign.Duration(),
		PriceManipulationLimit: campaign.PriceManipulationLimit(),
		TargetSalesCount:       campaign.TargetSalesCount(),
		SalesCount:             campaign.SalesCount(),
		TurnOver:               campaign.TurnOver(),
		CreatedAt:              campaign.CreatedAt(),
	}, nil
}

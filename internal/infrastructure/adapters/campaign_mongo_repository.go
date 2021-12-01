package adapters

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/domain/models"
	"github.com/ybalcin/ecommerce-study/pkg/mgo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type campaignBson struct {
	Id                     primitive.ObjectID `bson:"_id"`
	Name                   string             `bson:"name"`
	ProductCode            string             `bson:"product_code"`
	Duration               int                `bson:"duration"`
	PriceManipulationLimit int                `bson:"price_manipulation_limit"`
	TargetSalesCount       int                `bson:"target_sales_count"`
}

// campaignRepository implements repositories.CampaignRepository
type campaignRepository struct {
	campaigns *mgo.Collection
}

const campaignCollection string = "campaigns"

// NewCampaignRepository initializes new campaign repository
func NewCampaignRepository(store *mgo.Store) *campaignRepository {
	if store == nil {
		panic("adapters: mgo store is nil")
	}

	return &campaignRepository{campaigns: store.Collection(campaignCollection)}
}

// AddCampaign adds campaign to collection
func (r *campaignRepository) AddCampaign(ctx context.Context, campaign *models.Campaign) error {
	objId, err := mgo.ToObjectID(campaign.Id().String())
	if err != nil {
		return err
	}

	campaignBson := campaignBson{
		Id:                     objId,
		Name:                   campaign.Name(),
		ProductCode:            campaign.ProductCode(),
		Duration:               campaign.Duration(),
		PriceManipulationLimit: campaign.PriceManipulationLimit(),
		TargetSalesCount:       campaign.TargetSalesCount(),
	}

	if _, err := r.campaigns.InsertOne(ctx, campaignBson); err != nil {
		return err
	}

	return nil
}

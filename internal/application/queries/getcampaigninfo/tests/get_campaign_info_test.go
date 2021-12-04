package getcampaigninfo_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/application/queries/getcampaigninfo"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		t    string
		h    *getcampaigninfo.Handler
		q    *getcampaigninfo.Query
		fail bool
	}{
		{
			"handler nil",
			nil,
			new(getcampaigninfo.Query),
			true,
		},
		{
			"campaignRepository nil",
			getcampaigninfo.NewHandler(nil, new(repositories.MockOrderRepository), new(application.SystemTime)),
			new(getcampaigninfo.Query),
			true,
		},
		{
			"orderRepository nil",
			getcampaigninfo.NewHandler(new(repositories.MockCampaignRepository), nil, new(application.SystemTime)),
			new(getcampaigninfo.Query),
			true,
		},
		{
			"query nil",
			getcampaigninfo.NewHandler(new(repositories.MockCampaignRepository), new(repositories.MockOrderRepository), new(application.SystemTime)),
			nil,
			true,
		},
		{
			"campaignRepository.GetCampaign returns err",
			getcampaigninfo.NewHandler(repositories.MockCampaignRepository{
				GetCampaignFn: func(ctx context.Context, name string) (*domain.Campaign, error) {
					return nil, errors.New("")
				},
			}, new(repositories.MockOrderRepository), new(application.SystemTime)),
			new(getcampaigninfo.Query),
			true,
		},
		{
			"campaignRepository.GetCampaign returns nil campaign",
			getcampaigninfo.NewHandler(repositories.MockCampaignRepository{
				GetCampaignFn: func(ctx context.Context, name string) (*domain.Campaign, error) {
					return nil, nil
				},
			}, new(repositories.MockOrderRepository), new(application.SystemTime)),
			new(getcampaigninfo.Query),
			true,
		},
		{
			"orderRepository.GetOrders returns error",
			getcampaigninfo.NewHandler(repositories.MockCampaignRepository{
				GetCampaignFn: func(ctx context.Context, name string) (*domain.Campaign, error) {
					return new(domain.Campaign), nil
				},
			}, repositories.MockOrderRepository{
				GetOrdersFn: func(ctx context.Context, productCode string) ([]domain.Order, error) {
					return nil, errors.New("")
				},
			}, new(application.SystemTime)),
			new(getcampaigninfo.Query),
			true,
		},
		{
			"success",
			getcampaigninfo.NewHandler(repositories.MockCampaignRepository{
				GetCampaignFn: func(ctx context.Context, name string) (*domain.Campaign, error) {
					return new(domain.Campaign), nil
				},
			}, repositories.MockOrderRepository{
				GetOrdersFn: func(ctx context.Context, productCode string) ([]domain.Order, error) {
					return []domain.Order{}, nil
				},
			}, new(application.SystemTime)),
			new(getcampaigninfo.Query),
			false,
		},
	}

	for _, c := range testCases {
		t.Run(c.t, func(t *testing.T) {
			resp, err := c.h.Handle(ctx, c.q)
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

package getproductinfo_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/application/queries/getproductinfo"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		t    string
		h    *getproductinfo.Handler
		q    *getproductinfo.Query
		fail bool
	}{
		{
			"handler nil",
			nil,
			new(getproductinfo.Query),
			true,
		},
		{
			"product repo nil",
			getproductinfo.NewHandler(nil, new(repositories.MockCampaignRepository), new(application.SystemTime)),
			new(getproductinfo.Query),
			true,
		},
		{
			"campaign repo nil",
			getproductinfo.NewHandler(new(repositories.MockProductRepository), nil, new(application.SystemTime)),
			new(getproductinfo.Query),
			true,
		},
		{
			"sysTime nil",
			getproductinfo.NewHandler(new(repositories.MockProductRepository), new(repositories.MockCampaignRepository), nil),
			new(getproductinfo.Query),
			true,
		},
		{
			"query nil",
			getproductinfo.NewHandler(new(repositories.MockProductRepository), new(repositories.MockCampaignRepository), new(application.SystemTime)),
			nil,
			true,
		},
		{
			"productRepository.GetProduct returns error",
			getproductinfo.NewHandler(&repositories.MockProductRepository{
				GetProductFn: func(ctx context.Context, productCode string) (*domain.Product, error) {
					return nil, errors.New("")
				},
			}, new(repositories.MockCampaignRepository), new(application.SystemTime)),
			new(getproductinfo.Query),
			true,
		},
		{
			"productRepository.GetProduct returns nil product",
			getproductinfo.NewHandler(&repositories.MockProductRepository{
				GetProductFn: func(ctx context.Context, productCode string) (*domain.Product, error) {
					return nil, nil
				},
			}, new(repositories.MockCampaignRepository), new(application.SystemTime)),
			new(getproductinfo.Query),
			true,
		},
		{
			"campaignRepository.GetLatestCampaign returns error",
			getproductinfo.NewHandler(&repositories.MockProductRepository{
				GetProductFn: func(ctx context.Context, productCode string) (*domain.Product, error) {
					return new(domain.Product), nil
				},
			}, repositories.MockCampaignRepository{
				GetLatestCampaignFn: func(ctx context.Context, productCode string) (*domain.Campaign, error) {
					return nil, errors.New("")
				},
			}, new(application.SystemTime)),
			new(getproductinfo.Query),
			true,
		},
		{
			"campaignRepository.GetLatestCampaign returns nil campaign",
			getproductinfo.NewHandler(&repositories.MockProductRepository{
				GetProductFn: func(ctx context.Context, productCode string) (*domain.Product, error) {
					return new(domain.Product), nil
				},
			}, repositories.MockCampaignRepository{
				GetLatestCampaignFn: func(ctx context.Context, productCode string) (*domain.Campaign, error) {
					return nil, nil
				},
			}, new(application.SystemTime)),
			new(getproductinfo.Query),
			true,
		},
		{
			"campaignRepository.GetLatestCampaign returns nil campaign",
			getproductinfo.NewHandler(&repositories.MockProductRepository{
				GetProductFn: func(ctx context.Context, productCode string) (*domain.Product, error) {
					return new(domain.Product), nil
				},
			}, repositories.MockCampaignRepository{
				GetLatestCampaignFn: func(ctx context.Context, productCode string) (*domain.Campaign, error) {
					return nil, nil
				},
			}, new(application.SystemTime)),
			new(getproductinfo.Query),
			true,
		},
		{
			"success",
			getproductinfo.NewHandler(&repositories.MockProductRepository{
				GetProductFn: func(ctx context.Context, productCode string) (*domain.Product, error) {
					return new(domain.Product), nil
				},
			}, repositories.MockCampaignRepository{
				GetLatestCampaignFn: func(ctx context.Context, productCode string) (*domain.Campaign, error) {
					return new(domain.Campaign), nil
				},
			}, new(application.SystemTime)),
			new(getproductinfo.Query),
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

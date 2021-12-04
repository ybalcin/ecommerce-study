package createorder_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/createorder"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		t    string
		h    *createorder.Handler
		c    *createorder.Command
		fail bool
	}{
		{
			"nil handler",
			nil,
			new(createorder.Command),
			true,
		},
		{
			"order repo nil",
			createorder.NewHandler(nil, new(repositories.MockProductRepository), new(repositories.MockCampaignRepository), new(application.SystemTime)),
			new(createorder.Command),
			true,
		},
		{
			"product repo nil",
			createorder.NewHandler(new(repositories.MockOrderRepository), nil, new(repositories.MockCampaignRepository), new(application.SystemTime)),
			new(createorder.Command),
			true,
		},
		{
			"campaign repo nil",
			createorder.NewHandler(new(repositories.MockOrderRepository), new(repositories.MockProductRepository), nil, new(application.SystemTime)),
			new(createorder.Command),
			true,
		},
		{
			"sysTime nil",
			createorder.NewHandler(new(repositories.MockOrderRepository), new(repositories.MockProductRepository), new(repositories.MockCampaignRepository), nil),
			new(createorder.Command),
			true,
		},
		{
			"command nil",
			createorder.NewHandler(new(repositories.MockOrderRepository), new(repositories.MockProductRepository), new(repositories.MockCampaignRepository), new(application.SystemTime)),
			nil,
			true,
		},
		{
			"productRepository.GetProduct returns err",
			createorder.NewHandler(new(repositories.MockOrderRepository), &repositories.MockProductRepository{
				GetProductFn: func(ctx context.Context, productCode string) (*domain.Product, error) {
					return nil, errors.New("")
				},
			}, new(repositories.MockCampaignRepository), new(application.SystemTime)),
			new(createorder.Command),
			true,
		},
		{
			"product out of stock",
			createorder.NewHandler(new(repositories.MockOrderRepository), &repositories.MockProductRepository{
				GetProductFn: func(ctx context.Context, productCode string) (*domain.Product, error) {
					p, _ := domain.NewProduct("", "p1", 1, 1)
					p.ReduceStock(1)
					return p, nil
				},
			}, new(repositories.MockCampaignRepository), new(application.SystemTime)),
			new(createorder.Command),
			true,
		},
		{
			"product out of stock / order quantity greater than stock",
			createorder.NewHandler(new(repositories.MockOrderRepository), &repositories.MockProductRepository{
				GetProductFn: func(ctx context.Context, productCode string) (*domain.Product, error) {
					p, _ := domain.NewProduct("", "p1", 1, 1)
					return p, nil
				},
			}, new(repositories.MockCampaignRepository), new(application.SystemTime)),
			createorder.Build("create_order P11 10"),
			true,
		},
		{
			"campaignRepository.GetLatestCampaign returns err",
			createorder.NewHandler(new(repositories.MockOrderRepository), &repositories.MockProductRepository{
				GetProductFn: func(ctx context.Context, productCode string) (*domain.Product, error) {
					p, _ := domain.NewProduct("", "p1", 1, 100)
					return p, nil
				},
			}, repositories.MockCampaignRepository{
				GetLatestCampaignFn: func(ctx context.Context, productCode string) (*domain.Campaign, error) {
					return nil, errors.New("")
				},
			}, new(application.SystemTime)),
			createorder.Build("create_order P11 10"),
			true,
		},
		{
			"campaignRepository.GetLatestCampaign returns nil campaign",
			createorder.NewHandler(new(repositories.MockOrderRepository), &repositories.MockProductRepository{
				GetProductFn: func(ctx context.Context, productCode string) (*domain.Product, error) {
					p, _ := domain.NewProduct("", "p1", 1, 100)
					return p, nil
				},
			}, repositories.MockCampaignRepository{
				GetLatestCampaignFn: func(ctx context.Context, productCode string) (*domain.Campaign, error) {
					return nil, nil
				},
			}, new(application.SystemTime)),
			createorder.Build("create_order P11 10"),
			true,
		},
		{
			"orderRepository.AddOrder returns err",
			createorder.NewHandler(repositories.MockOrderRepository{
				AddOrderFn: func(ctx context.Context, order *domain.Order) error {
					return errors.New("")
				},
			}, &repositories.MockProductRepository{
				GetProductFn: func(ctx context.Context, productCode string) (*domain.Product, error) {
					p, _ := domain.NewProduct("", "p1", 1, 100)
					return p, nil
				},
			}, repositories.MockCampaignRepository{
				GetLatestCampaignFn: func(ctx context.Context, productCode string) (*domain.Campaign, error) {
					return new(domain.Campaign), nil
				},
			}, new(application.SystemTime)),
			createorder.Build("create_order P11 10"),
			true,
		},
		{
			"productRepository.UpdateProductStock returns err",
			createorder.NewHandler(repositories.MockOrderRepository{
				AddOrderFn: func(ctx context.Context, order *domain.Order) error {
					return nil
				},
			}, &repositories.MockProductRepository{
				GetProductFn: func(ctx context.Context, productCode string) (*domain.Product, error) {
					p, _ := domain.NewProduct("", "p1", 1, 100)
					return p, nil
				},
				UpdateProductStockFn: func(ctx context.Context, product *domain.Product) error {
					return errors.New("")
				},
			}, repositories.MockCampaignRepository{
				GetLatestCampaignFn: func(ctx context.Context, productCode string) (*domain.Campaign, error) {
					return new(domain.Campaign), nil
				},
			}, new(application.SystemTime)),
			createorder.Build("create_order P11 10"),
			true,
		},
		{
			"campaignRepository.UpdateCampaignTurnOverSales returns err",
			createorder.NewHandler(repositories.MockOrderRepository{
				AddOrderFn: func(ctx context.Context, order *domain.Order) error {
					return nil
				},
			}, &repositories.MockProductRepository{
				GetProductFn: func(ctx context.Context, productCode string) (*domain.Product, error) {
					p, _ := domain.NewProduct("", "p1", 1, 100)
					return p, nil
				},
				UpdateProductStockFn: func(ctx context.Context, product *domain.Product) error {
					return nil
				},
			}, repositories.MockCampaignRepository{
				GetLatestCampaignFn: func(ctx context.Context, productCode string) (*domain.Campaign, error) {
					return new(domain.Campaign), nil
				},
				UpdateCampaignTurnOverSalesFn: func(ctx context.Context, campaign *domain.Campaign) error {
					return errors.New("")
				},
			}, new(application.SystemTime)),
			createorder.Build("create_order P11 10"),
			true,
		},
		{
			"success",
			createorder.NewHandler(repositories.MockOrderRepository{
				AddOrderFn: func(ctx context.Context, order *domain.Order) error {
					return nil
				},
			}, &repositories.MockProductRepository{
				GetProductFn: func(ctx context.Context, productCode string) (*domain.Product, error) {
					p, _ := domain.NewProduct("", "p1", 1, 100)
					return p, nil
				},
				UpdateProductStockFn: func(ctx context.Context, product *domain.Product) error {
					return nil
				},
			}, repositories.MockCampaignRepository{
				GetLatestCampaignFn: func(ctx context.Context, productCode string) (*domain.Campaign, error) {
					return new(domain.Campaign), nil
				},
				UpdateCampaignTurnOverSalesFn: func(ctx context.Context, campaign *domain.Campaign) error {
					return nil
				},
			}, new(application.SystemTime)),
			createorder.Build("create_order P11 10"),
			false,
		},
	}

	for _, c := range testCases {
		t.Run(c.t, func(t *testing.T) {
			resp, err := c.h.Handle(ctx, c.c)
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

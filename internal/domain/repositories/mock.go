package repositories

import (
	"context"
	"github.com/ybalcin/ecommerce-study/internal/domain"
)

type MockCampaignRepository struct {
	AddCampaignFn      func(ctx context.Context, campaign *domain.Campaign) error
	AddCampaignInvoked bool

	GetCampaignFn      func(ctx context.Context, name string) (*domain.Campaign, error)
	GetCampaignInvoked bool

	GetLatestCampaignFn      func(ctx context.Context, productCode string) (*domain.Campaign, error)
	GetLatestCampaignInvoked bool

	UpdateCampaignTurnOverSalesFn      func(ctx context.Context, campaign *domain.Campaign) error
	UpdateCampaignTurnOverSalesInvoked bool

	DropCampaignsFn      func(ctx context.Context) error
	DropCampaignsInvoked bool
}

func (r MockCampaignRepository) AddCampaign(ctx context.Context, campaign *domain.Campaign) error {
	r.AddCampaignInvoked = true
	return r.AddCampaignFn(ctx, campaign)
}

func (r MockCampaignRepository) GetCampaign(ctx context.Context, name string) (*domain.Campaign, error) {
	r.GetCampaignInvoked = true
	return r.GetCampaignFn(ctx, name)
}

func (r MockCampaignRepository) GetLatestCampaign(ctx context.Context, productCode string) (*domain.Campaign, error) {
	r.GetLatestCampaignInvoked = true
	return r.GetLatestCampaignFn(ctx, productCode)
}

func (r MockCampaignRepository) UpdateCampaignTurnOverSales(ctx context.Context, campaign *domain.Campaign) error {
	r.UpdateCampaignTurnOverSalesInvoked = true
	return r.UpdateCampaignTurnOverSalesFn(ctx, campaign)
}

func (r MockCampaignRepository) DropCampaigns(ctx context.Context) error {
	r.DropCampaignsInvoked = true
	return r.DropCampaignsFn(ctx)
}

type MockOrderRepository struct {
	GetOrdersFn      func(ctx context.Context, productCode string) ([]domain.Order, error)
	GetOrdersInvoked bool

	AddOrderFn      func(ctx context.Context, order *domain.Order) error
	AddOrderInvoked bool

	DropOrdersFn      func(ctx context.Context) error
	DropOrdersInvoked bool
}

func (r MockOrderRepository) GetOrders(ctx context.Context, productCode string) ([]domain.Order, error) {
	r.GetOrdersInvoked = true
	return r.GetOrdersFn(ctx, productCode)
}

func (r MockOrderRepository) AddOrder(ctx context.Context, order *domain.Order) error {
	r.AddOrderInvoked = true
	return r.AddOrderFn(ctx, order)
}

func (r MockOrderRepository) DropOrders(ctx context.Context) error {
	r.DropOrdersInvoked = true
	return r.DropOrdersFn(ctx)
}

type MockProductRepository struct {
	AddProductFn      func(ctx context.Context, product *domain.Product) error
	AddProductInvoked bool

	GetProductFn      func(ctx context.Context, productCode string) (*domain.Product, error)
	GetProductInvoked bool

	UpdateProductStockFn      func(ctx context.Context, product *domain.Product) error
	UpdateProductStockInvoked bool

	DropProductsFn      func(ctx context.Context) error
	DropProductsInvoked bool
}

func (r MockProductRepository) AddProduct(ctx context.Context, product *domain.Product) error {
	r.AddProductInvoked = true
	return r.AddProductFn(ctx, product)
}

func (r MockProductRepository) GetProduct(ctx context.Context, productCode string) (*domain.Product, error) {
	r.GetProductInvoked = true
	return r.GetProductFn(ctx, productCode)
}

func (r MockProductRepository) UpdateProductStock(ctx context.Context, product *domain.Product) error {
	r.UpdateProductStockInvoked = true
	return r.UpdateProductStockFn(ctx, product)
}

func (r *MockProductRepository) DropProducts(ctx context.Context) error {
	r.DropProductsInvoked = true
	return r.DropProductsFn(ctx)
}

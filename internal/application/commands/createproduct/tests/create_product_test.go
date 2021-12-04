package createproduct_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/createproduct"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"github.com/ybalcin/ecommerce-study/internal/domain/repositories"
	"testing"
)

func TestHandler_Handle(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		t    string
		h    *createproduct.Handler
		c    *createproduct.Command
		fail bool
	}{
		{
			"handler nil",
			nil,
			new(createproduct.Command),
			true,
		},
		{
			"repo nil",
			createproduct.NewHandler(nil),
			new(createproduct.Command),
			true,
		},
		{
			"command nil",
			createproduct.NewHandler(new(repositories.MockProductRepository)),
			nil,
			true,
		},
		{
			"domain.NewProduct returns err",
			createproduct.NewHandler(new(repositories.MockProductRepository)),
			createproduct.Build("create_product"),
			true,
		},
		{
			"productRepository.AddProduct returns err",
			createproduct.NewHandler(&repositories.MockProductRepository{
				AddProductFn: func(ctx context.Context, product *domain.Product) error {
					return errors.New("")
				},
			}),
			createproduct.Build("create_product P11 100 1000"),
			true,
		},
		{
			"success",
			createproduct.NewHandler(&repositories.MockProductRepository{
				AddProductFn: func(ctx context.Context, product *domain.Product) error {
					return nil
				},
			}),
			createproduct.Build("create_product P11 100 1000"),
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

package domain_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"testing"
)

func TestNewOrder(t *testing.T) {
	testCases := []struct {
		productCode   string
		quantity      int
		unitSalePrice int
		fail          bool
	}{
		{"", 1, 1, true},
		{"dummy", 0, 1, true},
		{"dummy", 1, 0, true},
		{"dummy", 1, 1, false},
	}

	for _, c := range testCases {
		order, err := domain.NewOrder(c.productCode, c.quantity, c.unitSalePrice)
		if c.fail {
			assert.Nil(t, order)
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.NotNil(t, order)
		}
	}
}

func TestOrder_TotalPrice(t *testing.T) {
	testCases := []struct {
		quantity      int
		unitSalePrice int
	}{
		{2, 5},
		{1, 5},
	}

	for _, c := range testCases {
		order, err := domain.NewOrder("dummy", c.quantity, c.unitSalePrice)
		assert.Nil(t, err)
		assert.Equal(t, order.TotalPrice(), c.quantity*c.unitSalePrice)
	}
}

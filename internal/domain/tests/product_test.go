package domain_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/domain"
	"testing"
)

func TestNewProduct(t *testing.T) {
	testCases := []struct {
		id    string
		code  string
		price int
		stock int
		fail  bool
	}{
		{"", "p1", 1, 1, false},
		{"", "", 1, 1, true},
		{"", "p1", 0, 1, true},
		{"", "p1", 1, 0, true},
		{"asd", "p1", 1, 1, false},
	}

	for _, c := range testCases {
		product, err := domain.NewProduct(c.id, c.code, c.price, c.stock)
		if c.fail {
			assert.NotNil(t, err)
			assert.Nil(t, product)
		} else {
			assert.Nil(t, err)
			assert.NotNil(t, product)
		}
	}
}

func TestProduct_ReduceStock(t *testing.T) {
	testCases := []struct {
		handOnStock int
		val         int
		expected    int
	}{
		{10, 5, 5},
		{10, 10, 0},
		{10, 20, 10},
	}

	for _, c := range testCases {
		product, err := domain.NewProduct("", "dummy", 1, c.handOnStock)
		assert.Nil(t, err)
		product.ReduceStock(c.val)
		assert.Equal(t, product.Stock(), c.expected)
	}
}

func TestProduct_InStock(t *testing.T) {
	product, err := domain.NewProduct("", "dummy", 1, 1)
	assert.Nil(t, err)
	assert.Equal(t, true, product.InStock())
}

func TestProduct_ApplyPrice(t *testing.T) {
	product, err := domain.NewProduct("", "dummy", 1, 1)
	assert.Nil(t, err)
	product.ApplyPrice(10)
	assert.Equal(t, 10, product.Price())
}

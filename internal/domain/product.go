package domain

import (
	"github.com/ybalcin/ecommerce-study/internal/domain/errors"
)

// Product model
type Product struct {
	id    string
	code  string
	price int
	stock int
}

// NewProduct initializes new product
func NewProduct(id, productCode string, price int, stock int) (*Product, error) {
	product := &Product{
		id:    id,
		code:  productCode,
		price: price,
		stock: stock,
	}

	if err := product.validate(); err != nil {
		return nil, err
	}

	return product, nil
}

// Id returns id of product
func (p *Product) Id() string {
	return p.id
}

// Stock returns stock of product
func (p *Product) Stock() int {
	return p.stock
}

// Price returns price of product
func (p *Product) Price() int {
	return p.price
}

// Code returns product code of product
func (p *Product) Code() string {
	return p.code
}

// ReduceStock reduces stock
func (p *Product) ReduceStock(val int) {
	if val > p.stock {
		return
	}

	p.stock -= val
}

// InStock returns true if product in stock
func (p *Product) InStock() bool {
	return p.stock > 0
}

// ApplyPrice sets new price
func (p *Product) ApplyPrice(newPrice int) {
	p.price = newPrice
}

// validate validates model
func (p *Product) validate() error {

	if p.code == "" {
		return errors.ThrowProductCodeShouldNotBeEmptyError()
	}
	if p.price <= 0 {
		return errors.ThrowProductPriceValueIsInvalidError()
	}
	if p.stock <= 0 {
		return errors.ThrowProductStockValueIsInvalidError()
	}

	return nil
}

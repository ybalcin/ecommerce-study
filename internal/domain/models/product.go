package models

import "fmt"

type (
	ProductId string

	// Product model
	Product struct {
		id          ProductId
		productCode string
		price       float64
		stock       int
	}
)

// NewProduct initializes new product
func NewProduct(id ProductId, productCode string, price float64, stock int) (*Product, error) {
	product := &Product{
		id:          id,
		productCode: productCode,
		price:       price,
		stock:       stock,
	}

	if err := product.validate(); err != nil {
		return nil, err
	}

	return product, nil
}

// Stock returns stock of product
func (p *Product) Stock() int {
	return p.stock
}

// Price returns price of product
func (p *Product) Price() float64 {
	return p.price
}

// ProductCode returns product code of product
func (p *Product) ProductCode() string {
	return p.productCode
}

// Id returns id of product
func (p *Product) Id() ProductId {
	return p.id
}

// String converts product id to string
func (id ProductId) String() string {
	return string(id)
}

// validate validates model
func (p *Product) validate() error {
	invalidErr := func(key string) error {
		return fmt.Errorf("models: product %s is invalid", key)
	}

	if p.id == "" {
		return invalidErr("id")
	}
	if p.productCode == "" {
		return invalidErr("productCode")
	}
	if p.price <= 0 {
		return invalidErr("price")
	}
	if p.stock <= 0 {
		return invalidErr("stock")
	}

	return nil
}

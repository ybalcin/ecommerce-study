package domain

import "github.com/ybalcin/ecommerce-study/internal/domain/errors"

// Order entity
type Order struct {
	productCode   string
	quantity      int
	unitSalePrice int
}

// NewOrder initializes new order
func NewOrder(productCode string, quantity int, unitSalePrice int) (*Order, error) {
	order := &Order{
		productCode:   productCode,
		quantity:      quantity,
		unitSalePrice: unitSalePrice,
	}

	if err := order.validate(); err != nil {
		return nil, err
	}

	return order, nil
}

// ProductCode product code of order
func (o *Order) ProductCode() string {
	return o.productCode
}

// Quantity returns item quantity of order
func (o *Order) Quantity() int {
	return o.quantity
}

// UnitSalePrice returns unit sale price
func (o *Order) UnitSalePrice() int {
	return o.unitSalePrice
}

// TotalPrice returns total order price
func (o *Order) TotalPrice() int {
	return o.quantity * o.unitSalePrice
}

func (o *Order) validate() error {
	if o.productCode == "" {
		return errors.ThrowOrderProductCodeIsEmptyError()
	}
	if o.quantity <= 0 {
		return errors.ThrowOrderQuantityIsInvalidError()
	}
	if o.unitSalePrice <= 0 {
		return errors.ThrowOrderUnitSalePriceInvalidError()
	}

	return nil
}

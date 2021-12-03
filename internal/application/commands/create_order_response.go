package commands

import "fmt"

type createOrderResponse struct {
	ProductCode string
	Quantity    int
}

func NewCreateOrderResponse(productCode string, quantity int) *createOrderResponse {
	return &createOrderResponse{
		ProductCode: productCode,
		Quantity:    quantity,
	}
}

func (r *createOrderResponse) String() string {
	return fmt.Sprintf("Order created; product %s, quantity %d", r.ProductCode, r.Quantity)
}

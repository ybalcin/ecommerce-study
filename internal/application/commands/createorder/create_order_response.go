package createorder

import "fmt"

type response struct {
	ProductCode string
	Quantity    int
}

func NewResponse(productCode string, quantity int) *response {
	return &response{
		ProductCode: productCode,
		Quantity:    quantity,
	}
}

func (r *response) String() string {
	return fmt.Sprintf("Order created; product %s, quantity %d", r.ProductCode, r.Quantity)
}

package createproduct

import "fmt"

type response struct {
	ProductCode string
	Price       int
	Stock       int
}

func NewResponse(code string, price, stock int) *response {
	return &response{
		ProductCode: code,
		Price:       price,
		Stock:       stock,
	}
}

func (r *response) String() string {
	return fmt.Sprintf("Product created; code %s, price %d, stock %d", r.ProductCode, r.Price, r.Stock)
}

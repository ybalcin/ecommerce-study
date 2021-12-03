package getproductinfo

import "fmt"

type response struct {
	Code  string
	Price int
	Stock int
}

// NewResponse initializes response
func NewResponse(code string, price, stock int) *response {
	return &response{
		Code:  code,
		Price: price,
		Stock: stock,
	}
}

func (r *response) String() string {
	return fmt.Sprintf("Product %s info; price %d, stock %d", r.Code, r.Price, r.Stock)
}

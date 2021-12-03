package queries

import "fmt"

type getProductInfoQueryResponse struct {
	Code  string
	Price int
	Stock int
}

// NewGetProductInfoQueryResponse initializes getProductInfoQueryResponse
func NewGetProductInfoQueryResponse(code string, price, stock int) *getProductInfoQueryResponse {
	return &getProductInfoQueryResponse{
		Code:  code,
		Price: price,
		Stock: stock,
	}
}

func (r *getProductInfoQueryResponse) String() string {
	return fmt.Sprintf("Product %s info; price %d, stock %d", r.Code, r.Price, r.Stock)
}

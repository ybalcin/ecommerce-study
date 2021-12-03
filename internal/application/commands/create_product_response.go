package commands

import "fmt"

type createProductResponse struct {
	ProductCode string
	Price       int
	Stock       int
}

func NewCreateProductResponse(code string, price, stock int) *createProductResponse {
	return &createProductResponse{
		ProductCode: code,
		Price:       price,
		Stock:       stock,
	}
}

func (r *createProductResponse) String() string {
	return fmt.Sprintf("Product created; code %s, price %d, stock %d", r.ProductCode, r.Price, r.Stock)
}

package getproductinfo

import (
	"github.com/ybalcin/ecommerce-study/internal/common"
	"strings"
)

const entry string = "get_product_info"

// Build builds Query
func Build(query string) *Query {
	if query == "" {
		return nil
	}

	tokens := strings.Split(query, " ")

	if common.ValueOfSlice(0, tokens) != entry {
		return nil
	}

	code := common.ValueOfSlice(1, tokens)

	return &Query{
		Code: code,
	}
}

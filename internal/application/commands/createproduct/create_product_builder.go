package createproduct

import (
	"github.com/ybalcin/ecommerce-study/internal/common"
	"strings"
)

const entry string = "create_product"

// Build builds Command
func Build(cmd string) *Command {
	if cmd == "" {
		return nil
	}

	tokens := strings.Split(cmd, " ")

	if common.ValueOfSlice(0, tokens) != entry {
		return nil
	}

	code := common.ValueOfSlice(1, tokens)
	price := common.ValueOfSlice(2, tokens)
	stock := common.ValueOfSlice(3, tokens)

	return &Command{
		ProductCode: code,
		Price:       common.StringToInt(price),
		Stock:       common.StringToInt(stock),
	}
}

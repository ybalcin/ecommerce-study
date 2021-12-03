package createorder

import (
	"github.com/ybalcin/ecommerce-study/internal/common"
	"strings"
)

const entry string = "create_order"

// Build builds Command
func Build(cmd string) *Command {
	if cmd == "" {
		return nil
	}

	tokens := strings.Split(cmd, " ")

	if common.ValueOfSlice(0, tokens) != entry {
		return nil
	}

	productCode := common.ValueOfSlice(1, tokens)
	quantity := common.ValueOfSlice(2, tokens)

	return &Command{
		ProductCode: productCode,
		Quantity:    common.StringToInt(quantity),
	}
}

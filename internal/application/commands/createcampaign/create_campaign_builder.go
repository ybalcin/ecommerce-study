package createcampaign

import (
	"github.com/ybalcin/ecommerce-study/internal/common"
	"strings"
)

const entry string = "create_campaign"

// Build builds Command
func Build(cmd string) *Command {
	if cmd == "" {
		return nil
	}

	tokens := strings.Split(cmd, " ")

	if common.ValueOfSlice(0, tokens) != entry {
		return nil
	}

	name := common.ValueOfSlice(1, tokens)
	productCode := common.ValueOfSlice(2, tokens)
	duration := common.ValueOfSlice(3, tokens)
	pmLimit := common.ValueOfSlice(4, tokens)
	targetCount := common.ValueOfSlice(5, tokens)

	return &Command{
		Name:                   name,
		ProductCode:            productCode,
		Duration:               common.StringToInt(duration),
		PriceManipulationLimit: common.StringToInt(pmLimit),
		TargetSalesCount:       common.StringToInt(targetCount),
	}
}

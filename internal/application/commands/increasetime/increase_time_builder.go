package increasetime

import (
	"github.com/ybalcin/ecommerce-study/internal/common"
	"strings"
)

const entry string = "increase_time"

// Build builds Command
func Build(cmd string) *Command {
	if cmd == "" {
		return nil
	}

	tokens := strings.Split(cmd, " ")

	if common.ValueOfSlice(0, tokens) != entry {
		return nil
	}

	hours := common.ValueOfSlice(1, tokens)

	return &Command{
		Hours: common.StringToInt(hours),
	}
}

package commands

import (
	"github.com/ybalcin/ecommerce-study/internal/application"
)

type IncreaseHourCommand struct {
	Hours int
}

type IncreaseHourCommandHandler struct {
	systemTime *application.SystemTime
}

// NewIncreaseHourCommandHandler initializes new IncreaseHourCommandHandler
func NewIncreaseHourCommandHandler(systemTime *application.SystemTime) *IncreaseHourCommandHandler {
	return &IncreaseHourCommandHandler{
		systemTime: systemTime,
	}
}

// Handle handles IncreaseHourCommand
func (h *IncreaseHourCommandHandler) Handle(c *IncreaseHourCommand) {
	h.systemTime.Add(c.Hours)
}

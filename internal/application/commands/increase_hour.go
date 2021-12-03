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
func (h *IncreaseHourCommandHandler) Handle(c *IncreaseHourCommand) error {
	if h == nil {
		return application.ThrowIncreaseHourCommandHandlerNilError()
	}

	if err := h.validate(); err != nil {
		return err
	}

	if c == nil {
		return application.ThrowIncreaseHourCommandCannotBeNilError()
	}

	h.systemTime.Add(c.Hours)

	return nil
}

func (h *IncreaseHourCommandHandler) validate() error {
	if h.systemTime == nil {
		return application.ThrowSystemTimeCannotBeNilError()
	}

	return nil
}

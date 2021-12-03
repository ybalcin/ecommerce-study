package commands

import (
	"github.com/ybalcin/ecommerce-study/internal/application"
)

type IncreaseTimeCommand struct {
	Hours int
}

type IncreaseTimeCommandHandler struct {
	systemTime *application.SystemTime
}

// NewIncreaseTimeCommandHandler initializes new IncreaseTimeCommandHandler
func NewIncreaseTimeCommandHandler(systemTime *application.SystemTime) *IncreaseTimeCommandHandler {
	return &IncreaseTimeCommandHandler{
		systemTime: systemTime,
	}
}

// Handle handles IncreaseTimeCommand
func (h *IncreaseTimeCommandHandler) Handle(c *IncreaseTimeCommand) (*increaseTimeResponse, error) {
	if h == nil {
		return nil, application.ThrowIncreaseHourCommandHandlerNilError()
	}

	if err := h.validate(); err != nil {
		return nil, err
	}

	if c == nil {
		return nil, application.ThrowIncreaseTimeCommandCannotBeNilError()
	}

	h.systemTime.Add(c.Hours)

	return NewIncreaseTimeResponse(h.systemTime.String()), nil
}

func (h *IncreaseTimeCommandHandler) validate() error {
	if h.systemTime == nil {
		return application.ThrowSystemTimeCannotBeNilError()
	}

	return nil
}

package increasetime

import (
	"github.com/ybalcin/ecommerce-study/internal/application"
)

type Command struct {
	Hours int
}

type Handler struct {
	systemTime *application.SystemTime
}

// NewHandler initializes new Handler
func NewHandler(systemTime *application.SystemTime) *Handler {
	return &Handler{
		systemTime: systemTime,
	}
}

// Handle handles Command
func (h *Handler) Handle(c *Command) (*response, error) {
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

	return NewResponse(h.systemTime.String()), nil
}

func (h *Handler) validate() error {
	if h.systemTime == nil {
		return application.ThrowSystemTimeCannotBeNilError()
	}

	return nil
}

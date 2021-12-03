package commands_test

import (
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/application/commands"
	"testing"
)

func TestNewIncreaseHourCommandHandler(t *testing.T) {
	sysTime := application.NewSystemTime()

	sysTimeSnapshot := sysTime.Time()

	increaseHourCommandHandler := commands.NewIncreaseTimeCommandHandler(sysTime)

	increaseHourCommandHandler.Handle(&commands.IncreaseTimeCommand{Hours: 2})

	if !sysTime.Time().After(sysTimeSnapshot) {
		t.Fail()
	}
}

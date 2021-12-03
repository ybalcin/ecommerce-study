package increasetime_test

import (
	"github.com/ybalcin/ecommerce-study/internal/application"
	"github.com/ybalcin/ecommerce-study/internal/application/commands/increasetime"
	"testing"
)

func TestNewIncreaseHourCommandHandler(t *testing.T) {
	sysTime := application.NewSystemTime()

	sysTimeSnapshot := sysTime.Time()

	increaseHourCommandHandler := increasetime.NewHandler(sysTime)

	increaseHourCommandHandler.Handle(&increasetime.Command{Hours: 2})

	if !sysTime.Time().After(sysTimeSnapshot) {
		t.Fail()
	}
}

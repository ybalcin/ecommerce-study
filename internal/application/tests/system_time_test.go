package application_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/application"
	"testing"
	"time"
)

func TestSystemTime_Add(t *testing.T) {
	testCases := []int{
		1, 0, -1,
	}

	for _, c := range testCases {
		sysTime := application.NewSystemTime()
		sysTime.Add(c)
		assert.Equal(t, application.NewSystemTime().Time().Add(time.Hour*time.Duration(c)), sysTime.Time())
	}
}

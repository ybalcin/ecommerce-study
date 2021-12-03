package common_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ybalcin/ecommerce-study/internal/common"
	"testing"
)

func TestValueOfSlice(t *testing.T) {
	testSlice := []string{
		"test",
	}

	testCases := []struct {
		index    int
		expected string
	}{
		{0, "test"},
		{1, ""},
	}

	for _, c := range testCases {
		actual := common.ValueOfSlice(c.index, testSlice)
		assert.Equal(t, c.expected, actual)
	}
}

func TestStringToInt(t *testing.T) {
	testCases := []struct {
		val      string
		expected int
	}{
		{"0", 0},
		{"1", 1},
		{"", 0},
		{"asd", 0},
	}

	for _, c := range testCases {
		actual := common.StringToInt(c.val)
		assert.Equal(t, c.expected, actual)
	}
}

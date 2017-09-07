package reverseInteger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert := assert.New(t)
	var cases = []struct {
		input    int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{1000000003, 0},
	}

	for _, tt := range cases {
		assert.Equal(reverse(tt.input), tt.expected)
	}

}

func TestReverseWithGoroutine(t *testing.T) {
	assert := assert.New(t)
	var cases = []struct {
		input    int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{1000000003, 0},
	}

	for _, tt := range cases {
		assert.Equal(reverseWithGoroutine(tt.input), tt.expected)
	}
}

func TestReverseIn3msSolution(t *testing.T) {
	assert := assert.New(t)
	var cases = []struct {
		input    int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{1000000003, 0},
	}

	for _, tt := range cases {
		assert.Equal(reverseIn3msSolution(tt.input), tt.expected)
	}
}

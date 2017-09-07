package reverseInteger

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var addTwoTests = []struct {
		input    int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{1000000003, 0},
	}

	for _, tt := range addTwoTests {
		if result := reverse(tt.input); result != tt.expected {
			t.Fatalf("should be %v, but is:%v\n", tt.expected, result)
		}
	}

}

func TestReverseWithGoroutine(t *testing.T) {
	var addTwoTests = []struct {
		input    int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{1000000003, 0},
	}

	for _, tt := range addTwoTests {
		if result := reverseWithGoroutine(tt.input); result != tt.expected {
			t.Fatalf("should be %v, but is:%v\n", tt.expected, result)
		}
	}

}

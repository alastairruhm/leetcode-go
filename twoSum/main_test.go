package main

import (
	"reflect"
	"testing"
)

func TestAddTwo(t *testing.T) {
	var addTwoTests = []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
	}

	for _, tt := range addTwoTests {
		if result := twoSum(tt.nums, tt.target); !reflect.DeepEqual(result, tt.expected) {
			t.Fatalf("should be %v, but is:%v\n", tt.expected, result)
		}
	}

}

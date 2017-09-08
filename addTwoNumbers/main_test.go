package addTwoNumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	input1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	input2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	expected := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  8,
				Next: nil,
			},
		},
	}
	var cases = []struct {
		input1   *ListNode
		input2   *ListNode
		expected *ListNode
	}{
		{input1, input2, expected},
	}

	for _, tt := range cases {
		result := addTwoNumbers(tt.input1, tt.input2)
		assert.True(t, Compare(result, tt.expected))
	}
}

package longestSubstringWithoutRepeatingCharacters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var cases = []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"abcdef", 6},
	}

	for _, tt := range cases {
		result := lengthOfLongestSubstring(tt.input)
		assert.Equal(t, result, tt.expected)
	}
}

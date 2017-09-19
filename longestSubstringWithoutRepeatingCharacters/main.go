package longestSubstringWithoutRepeatingCharacters

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var target string
	for i := 0; i < len(s); i++ {
		t := string(s[i])

		for j := i + 1; j < len(s); j++ {
			if strings.Contains(t, string(s[j])) {
				break
			}
			t = t + string(s[j])
		}

		if len(t) > len(target) {
			target = t
		}
	}
	return len(target)
}

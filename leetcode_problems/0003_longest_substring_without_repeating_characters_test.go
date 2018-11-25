package leetcode_problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log("Test 0003_longest_substring_without_repeating_characters")
	list := []struct {
		String string
		Ans    int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
	}

	for _, item := range list {
		assert.Equal(t, LengthOfLongestSubstring(item.String), item.Ans)
	}
}

func LengthOfLongestSubstring(s string) int {
	char := map[rune]int{}
	pointer := 0
	length := 0
	longest := 0
	for i, v := range s {
		if _, ok := char[v]; !ok {
			char[v] = i
		} else {
			if char[v]+1 > pointer {
				pointer = char[v] + 1
			}
			char[v] = i
		}
		length = i - pointer + 1
		if length > longest {
			longest = length
		}
		// fmt.Printf("char: %s index: %d pointer: %d, length: %d, longest: %d\n", string(v), i, pointer, length, longest)
	}
	return longest
}

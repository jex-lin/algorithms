package leetcode_problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	t.Log("Test 0001_two_sum_test")
	list := []struct {
		Nums   []int
		Target int
		Ans    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for _, item := range list {
		// t.Log("Nums: ", item.Nums)
		assert.Equal(t, TwoSum(item.Nums, item.Target), item.Ans)
	}
}

// O(n)
func TwoSum(s []int, t int) []int {
	m := map[int]int{}
	for i, v := range s {
		if a, ok := m[t-v]; !ok {
			m[v] = i
		} else {
			return []int{a, i}
		}
	}
	return []int{-1, -1}
}

// O(n^2)
func BadTwoSum(s []int, t int) []int {
	for i, v := range s {
		for j := i + 1; j < len(s); j++ {
			if v+s[j] == t {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

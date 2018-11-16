package leetcode_problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	t.Log("Test 4_median_of_two_sorted_arrays_test")
	list := []struct {
		Num1 []int
		Num2 []int
		Ans  float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, item := range list {
		// t.Log("num1: ", item.Num1, " num2:", item.Num2)
		assert.Equal(t, FindMedianSortedArrays(item.Num1, item.Num2), item.Ans)
	}
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	return 2
}

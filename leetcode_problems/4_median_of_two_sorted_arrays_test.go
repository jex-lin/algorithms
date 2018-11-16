package leetcode_problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	t.Log("Test 4_median_of_two_sorted_arrays_test")
	list := []struct {
		Nums1 []int
		Nums2 []int
		Ans   float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, item := range list {
		// t.Log("num1: ", item.Num1, " num2:", item.Num2)
		assert.Equal(t, FindMedianSortedArrays(item.Nums1, item.Nums2), item.Ans)
	}
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := make([]int, 0)
	nums = append(nums, nums1...)
	nums = append(nums, nums2...)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / float64(2)
	}
	return float64(nums[len(nums)/2])
}

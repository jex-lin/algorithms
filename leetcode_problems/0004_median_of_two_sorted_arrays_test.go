package leetcode_problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	t.Log("Test 0004_median_of_two_sorted_arrays_test")
	list := []struct {
		Nums1 []int
		Nums2 []int
		Ans   float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{}, []int{1}, 1},
	}

	for _, item := range list {
		// t.Log("num1: ", item.Num1, " num2:", item.Num2)
		assert.Equal(t, item.Ans, FindMedianSortedArrays(item.Nums1, item.Nums2))
	}
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := make([]int, len(nums1)+len(nums2))
	var flag1, flag2 int
	var end1, end2 bool
	for i := 0; i < len(nums1)+len(nums2); i++ {
		// fmt.Println(flag1, nums1[flag1], flag2, nums2[flag2])
		if end1 || len(nums1) == 0 {
			nums[i] = nums2[flag2]
			flag2++
			continue
		}
		if end2 || len(nums2) == 0 {
			nums[i] = nums1[flag1]
			flag1++
			continue
		}
		if nums1[flag1] < nums2[flag2] {
			nums[i] = nums1[flag1]
			if flag1+1 < len(nums1) {
				flag1++
			} else {
				end1 = true
			}
		} else {
			nums[i] = nums2[flag2]
			if flag2+1 < len(nums2) {
				flag2++
			} else {
				end2 = true
			}
		}
	}

	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2-1]+nums[len(nums)/2]) / float64(2)
	}
	return float64(nums[len(nums)/2])
}

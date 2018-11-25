package leetcode_problems

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	t.Log("Test 0009_palindrome_number_test")
	list := []struct {
		Val int
		Ans bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{0, true},
		{11, true},
		{123321, true},
		{1234321, true},
	}

	for _, item := range list {
		t.Log(item.Val)
		assert.Equal(t, IsPalindrome(item.Val), item.Ans)
	}
}

// ref: https://leetcode.com/problems/palindrome-number/solution/
func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	r := 0
	for x > r {
		r = r*10 + x%10
		x /= 10
	}
	return x == r || x == r/10
}

// My answer, not good
func IsPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	s := strconv.Itoa(x)
	var flag bool
	for i := 0; i < len(s); i++ {
		match := s[i] == s[len(s)-1-i]
		if flag && !match {
			return false
		}
		if !flag && match {
			flag = true
		}
	}
	return flag
}

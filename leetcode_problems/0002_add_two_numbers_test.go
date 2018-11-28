package leetcode_problems

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestAddTwoNumbers(t *testing.T) {
	t.Log("Test 0002_add_two_numbers_test")
	list := []struct {
		L1  ListNode
		L2  ListNode
		Ans ListNode
	}{
		{
			ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
	}

	for _, item := range list {
		// fmt.Print("l1: ", item.L1)
		// showList(&item.L1)
		// fmt.Print("l2: ", item.L2)
		// showList(&item.L2)
		ans := AddTwoNumbers(&item.L1, &item.L2)
		assert.Equal(t, item.Ans.Val, ans.Val)
		assert.Equal(t, item.Ans.Next.Val, ans.Next.Val)
		assert.Equal(t, item.Ans.Next.Next.Val, ans.Next.Next.Val)
	}
}

// ref: https://leetcode.com/problems/add-two-numbers/discuss/185527/Golang-solutions
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	for prev, sum := head, 0; l1 != nil || l2 != nil || sum > 0; sum /= 10 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		prev.Next = &ListNode{
			Val: sum % 10,
		}
		prev = prev.Next
	}
	return head.Next
}

// My answer, not good
func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var res []ListNode
	var carryover int
	for {
		var v1, v2 int
		if l1 != nil {
			v1 = l1.Val
		}

		if l2 != nil {
			v2 = l2.Val
		}
		t := v1 + v2 + carryover

		carryover = 0
		val := 0
		if t >= 10 {
			val = t - 10
			carryover = 1
		} else {
			val = t
		}

		res = append(res, ListNode{Val: val})
		if l1 != nil && l1.Next != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}

		if l2 != nil && l2.Next != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}

		if l1 == nil && l2 == nil {
			if carryover == 1 {
				res = append(res, ListNode{Val: carryover})
				break
			} else {
				break
			}
		}
	}

	for i := 0; i < len(res); i++ {
		if i+1 < len(res) {
			res[i].Next = &res[i+1]
		}
	}
	return &res[0]
}

func showList(l *ListNode) {
	for {
		if l.Next != nil {
			fmt.Print(l.Val)
			l = l.Next
		} else {
			fmt.Print(l.Val)
			break
		}
	}
	fmt.Print("\n")
}

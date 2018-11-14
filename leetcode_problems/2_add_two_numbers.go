package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a2 := ListNode{3, nil}
	a1 := ListNode{4, &a2}
	a0 := ListNode{2, &a1}
	fmt.Print("l1: ")
	showList(&a0)

	b2 := ListNode{4, nil}
	b1 := ListNode{6, &b2}
	b0 := ListNode{5, &b1}
	fmt.Print("l2: ")
	showList(&b0)

	c := addTwoNumbers(&a0, &b0)
	fmt.Print("result: ")
	showList(c)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
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

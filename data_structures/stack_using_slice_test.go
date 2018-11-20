package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Stack struct {
	slice []int
}

func TestStackUsingSlice(t *testing.T) {
	t.Log("Test stack_using_slice")
	list := []struct {
		values []int
		ans    []int
	}{
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
	}

	for _, item := range list {
		var s Stack
		for _, v := range item.values {
			s.Push(v)
		}
		var tmp []int
		for range s.slice {
			tmp = append(tmp, s.Pop())
		}
		assert.Equal(t, reflect.DeepEqual(tmp, item.ans), true)
	}
}

func (s *Stack) Push(i int) {
	s.slice = append(s.slice, i)
}

// return the top item from the stack, but doesn't remove it from the stack
func (s *Stack) Peek() int {
	return s.slice[len(s.slice)-1]
}

func (s *Stack) Pop() int {
	ret := s.Peek()
	s.slice = s.slice[0 : len(s.slice)-1]
	return ret
}

func (s *Stack) String() string {
	return fmt.Sprint(s.slice)
}

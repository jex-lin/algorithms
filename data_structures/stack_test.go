package main

import (
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://flaviocopes.com/golang-data-structure-stack/

type Stack struct {
	items []int
	lock  sync.RWMutex
}

func TestStack(t *testing.T) {
	t.Log("Test stack_using_slice")
	list := []struct {
		stack []int
	}{
		{[]int{3, 9, 5, 4, 10, 8}},
	}

	for _, item := range list {
		var s Stack
		for _, i := range item.stack {
			s.Push(i)
		}
		assert.Equal(t, 8, s.Peek())
		assert.Equal(t, 8, s.Pop())
		assert.Equal(t, 5, s.Size())
		assert.Equal(t, false, s.IsEmpty())
		assert.Equal(t, true, reflect.DeepEqual(s.items, []int{3, 9, 5, 4, 10}))
	}
}

func (s *Stack) Push(i int) {
	s.lock.Lock()
	s.items = append(s.items, i)
	s.lock.Unlock()
}

func (s *Stack) Pop() int {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.lock.Unlock()
	return item
}

func (s *Stack) Peek() int {
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

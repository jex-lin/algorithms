package main

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ref: https://flaviocopes.com/golang-data-structure-linked-list/

type SinglyLinkedNode struct {
	value int
	next  *SinglyLinkedNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedNode
	size int
	lock sync.RWMutex
}

func TestSinglyLinkedList(t *testing.T) {
	t.Log("Test singly_linked_list")
	list := []struct {
		nodes   []int
		reverse []int
	}{
		{[]int{3, 9, 5, 4, 10, 8}, []int{8, 10, 4, 5, 9, 3}},
	}

	for _, item := range list {
		var q SinglyLinkedList
		assert.Equal(t, true, q.IsEmpty())
		for _, i := range item.nodes {
			q.Append(i)
		}
		assert.Equal(t, false, q.IsEmpty())
		assert.Equal(t, 6, q.Size())
		q.Insert(5, 7)
		q.RemoveAt(4)
		assert.Equal(t, 4, q.IndexOf(7))
	}

	// Reverse
	for _, item := range list {
		var q SinglyLinkedList
		assert.Equal(t, true, q.IsEmpty())
		for _, i := range item.nodes {
			q.Append(i)
		}
		q.Reverse()
		assert.Equal(t, true, reflect.DeepEqual(q.Slice(), item.reverse))
	}
}

func (l *SinglyLinkedList) Append(i int) {
	l.lock.Lock()
	node := SinglyLinkedNode{i, nil}
	if l.head == nil {
		l.head = &node
	} else {
		last := l.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &node
	}
	l.size++
	l.lock.Unlock()
}

// i stands for index
func (l *SinglyLinkedList) Insert(i int, val int) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	if i < 0 || i > l.size {
		return fmt.Errorf("Index out of bounds")
	}
	newNode := SinglyLinkedNode{val, nil}
	if i == 0 {
		newNode.next = l.head
		l.head = &newNode
		return nil
	}
	node := l.head
	j := 0
	for j < i-1 {
		j++
		node = node.next
	}
	newNode.next = node.next
	node.next = &newNode
	l.size++
	return nil
}

func (l *SinglyLinkedList) RemoveAt(i int) (int, error) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if i < 0 || i > l.size {
		return 0, fmt.Errorf("Index out of bounds")
	}
	node := l.head
	j := 0
	for j < i-1 {
		j++
		node = node.next
	}
	remove := node.next
	node.next = remove.next
	l.size--
	return remove.value, nil
}

func (l *SinglyLinkedList) IndexOf(val int) int {
	l.lock.Lock()
	defer l.lock.Unlock()
	node := l.head
	j := 0
	for {
		if node.value == val {
			break
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		j++
	}
	return j
}

func (l *SinglyLinkedList) IsEmpty() bool {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.head == nil {
		return true
	}
	return false
}

func (l *SinglyLinkedList) Size() int {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.head == nil {
		return 0
	}
	j := 1
	node := l.head
	for {
		if node.next == nil {
			break
		}
		j++
		node = node.next
	}
	return j
}

func (l *SinglyLinkedList) Reverse() {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.head == nil {
		return
	}

	var previous *SinglyLinkedNode
	current := l.head
	next := l.head.next
	current.next = nil
	for next != nil {
		previous, current, next = current, next, next.next
		current.next = previous
	}
	l.head = current
}

func (l *SinglyLinkedList) Slice() []int {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.head == nil {
		return []int{}
	}
	var s []int
	node := l.head
	for node != nil {
		s = append(s, node.value)
		node = node.next
	}
	return s
}

func (l *SinglyLinkedList) String() {
	l.lock.RLock()
	defer l.lock.RUnlock()
	node := l.head
	for {
		if node == nil {
			break
		}
		fmt.Print(node.value)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()
}

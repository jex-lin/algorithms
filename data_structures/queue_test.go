package main

import (
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://flaviocopes.com/golang-data-structure-queue/

type Queue struct {
	items []int
	lock  sync.RWMutex
}

func TestQueue(t *testing.T) {
	t.Log("Test queue")
	list := []struct {
		queue []int
	}{
		{[]int{3, 9, 5, 4, 10, 8}},
	}

	for _, item := range list {
		var q Queue
		for _, i := range item.queue {
			q.Enqueue(i)
		}
		assert.Equal(t, 3, q.Front())
		assert.Equal(t, 3, q.Dequeue())
		assert.Equal(t, 5, q.Size())
		assert.Equal(t, false, q.IsEmpty())
		assert.Equal(t, true, reflect.DeepEqual(q.items, []int{9, 5, 4, 10, 8}))
	}
}

func (q *Queue) Enqueue(i int) {
	q.lock.Lock()
	q.items = append(q.items, i)
	q.lock.Unlock()
}

func (q *Queue) Dequeue() int {
	q.lock.Lock()
	item := q.items[0]
	q.items = q.items[1:]
	q.lock.Unlock()
	return item
}

func (q *Queue) Front() int {
	q.lock.RLock()
	item := q.items[0]
	q.lock.RUnlock()
	return item
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

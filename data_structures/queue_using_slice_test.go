package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Queue struct {
	slice []int
}

func TestQueueUsingSlice(t *testing.T) {
	t.Log("Test queue_using_slice")
	list := []struct {
		values []int
		ans    []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
	}

	for _, item := range list {
		var q Queue
		for _, v := range item.values {
			q.Enqueue(v)
		}
		var tmp []int
		for range q.slice {
			tmp = append(tmp, q.Dequeue())
		}
		assert.Equal(t, reflect.DeepEqual(tmp, item.ans), true)
	}
}

func (q *Queue) Enqueue(i int) {
	q.slice = append(q.slice, i)
}

func (q *Queue) Dequeue() int {
	ret := q.slice[0]
	q.slice = q.slice[1:len(q.slice)]
	return ret
}

// Pretty print
func (q *Queue) String() string {
	return fmt.Sprint(q.slice)
}

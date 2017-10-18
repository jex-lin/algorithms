package main

import "fmt"

type Queue struct {
	slice []int
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

func main() {
	var q *Queue = new(Queue)
	q.Enqueue(1)
	fmt.Println("Queue: ", q)
	q.Enqueue(2)
	fmt.Println("Queue: ", q)
	q.Enqueue(3)
	fmt.Println("Queue: ", q)
	fmt.Println("Dequeue: ", q.Dequeue())
	fmt.Println("Queue: ", q)
	fmt.Println("Dequeue: ", q.Dequeue())
	fmt.Println("Queue: ", q)
	fmt.Println("Dequeue: ", q.Dequeue())

	/*
	   Queue:  [1]
	   Queue:  [1 2]
	   Queue:  [1 2 3]
	   Dequeue:  1
	   Queue:  [2 3]
	   Dequeue:  2
	   Queue:  [3]
	   Dequeue:  3
	*/
}

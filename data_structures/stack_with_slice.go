package main

import "fmt"

type Stack struct {
	slice []int
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

func main() {
	var s *Stack = new(Stack)
	s.Push(1)
	fmt.Println("Stack: ", s)
	s.Push(2)
	fmt.Println("Stack: ", s)
	s.Push(3)
	fmt.Println("Stack: ", s)
	s.Push(4)
	fmt.Println("Stack: ", s)
	fmt.Println("Pop: ", s.Pop())
	fmt.Println("Stack: ", s)
	fmt.Println("Pop: ", s.Pop())
	fmt.Println("Stack: ", s)
	fmt.Println("Pop: ", s.Pop())
	fmt.Println("Stack: ", s)
	fmt.Println("Pop: ", s.Pop())

	/*
	   Stack:  [1]
	   Stack:  [1 2]
	   Stack:  [1 2 3]
	   Stack:  [1 2 3 4]
	   Pop:  4
	   Stack:  [1 2 3]
	   Pop:  3
	   Stack:  [1 2]
	   Pop:  2
	   Stack:  [1]
	   Pop:  1
	*/
}

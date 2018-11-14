package main

import "fmt"

func main() {
	s := []int{2, 7, 11, 15}
	fmt.Printf("O(n): %v\n", twoSum(s, 9))
	fmt.Printf("O(n^2): %v\n", badTwoSum(s, 9))
}

// O(n)
func twoSum(s []int, t int) []int {
	m := map[int]int{}
	for i, v := range s {
		if a, ok := m[t-v]; !ok {
			m[v] = i
		} else {
			return []int{a, i}
		}
	}
	return []int{-1, -1}
}

// O(n^2)
func badTwoSum(s []int, t int) []int {
	for i, v := range s {
		for j := i + 1; j < len(s); j++ {
			if v+s[j] == t {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

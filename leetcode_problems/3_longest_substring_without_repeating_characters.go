package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	char := map[rune]int{}
	pointer := 0
	length := 0
	longest := 0
	for i, v := range s {
		if _, ok := char[v]; !ok {
			char[v] = i
		} else {
			if char[v]+1 > pointer {
				pointer = char[v] + 1
			}
			char[v] = i
		}
		length = i - pointer + 1
		if length > longest {
			longest = length
		}
		// fmt.Printf("char: %s index: %d pointer: %d, length: %d, longest: %d\n", string(v), i, pointer, length, longest)
	}
	return longest
}

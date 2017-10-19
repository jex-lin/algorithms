package main

import "fmt"

func main() {
	var numbers []int = []int{5, 4, 2, 3, 1, 0}
	fmt.Println("Our list of numbers is:", numbers)
	bubbleSort(numbers)
	fmt.Println("After sorting:", numbers)

	var numbers2 []int = []int{2, 1, 3, 4, 5, 6}
	fmt.Println("Our list of numbers is:", numbers2)
	bubbleSort(numbers2)
	fmt.Println("After sorting:", numbers2)
}

func bubbleSort(numbers []int) {
	var N int = len(numbers)
	for i := 0; i < N; i++ {
		fmt.Println("Doing a sweep:", numbers)
		if !sweep(numbers, i) {
			return // 如果都沒有交換, 代表排序已經完成, 直接提早結束
		}
	}
}

func sweep(numbers []int, prevPasses int) bool {
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1
	var didSwap bool = false // 判斷是否有交換

	for secondIndex < (N - prevPasses) { // 因為交換到最後一個一定是最大的, 所以每次比對到前一個位置就可以了
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]
		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true
		}

		firstIndex++
		secondIndex++
	}
	return didSwap
}

/*

	Our list of numbers is: [5 4 2 3 1 0]
	Doing a sweep: [5 4 2 3 1 0]
	Doing a sweep: [4 2 3 1 0 5]
	Doing a sweep: [2 3 1 0 4 5]
	Doing a sweep: [2 1 0 3 4 5]
	Doing a sweep: [1 0 2 3 4 5]
	Doing a sweep: [0 1 2 3 4 5]
	After sorting: [0 1 2 3 4 5]
	Our list of numbers is: [2 1 3 4 5 6]
	Doing a sweep: [2 1 3 4 5 6]
	Doing a sweep: [1 2 3 4 5 6]
	After sorting: [1 2 3 4 5 6]

*/

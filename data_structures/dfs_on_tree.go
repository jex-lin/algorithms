package main

import "fmt"

/*

	tree:
			4
		  2   6
		1  3

	index = 0 => [1, -1, -1]   第一個數代表 value, 第二個數代表左節點的 index 值, 第三個數代表右節點的 index 值
	index = 1 => [3, -1, -1]
	index = 2 => [2,  0,  1]
	index = 3 => [6, -1, -1]
	index = 4 => [4,  2,  3]

*/

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	nodes := read()
	for _, node := range nodes {
		printNode(&node)
	}
}

func read() []Node {
	var N int
	fmt.Scanf("%d", &N)
	var nodes []Node = make([]Node, N)
	for i := 0; i < N; i++ {
		var val, indexLeft, indexRight int
		fmt.Scanf("%d %d %d", &val, &indexLeft, &indexRight)
		// fmt.Printf("Value: %d IndexLeft: %d RightLeft: %d\n", val, indexLeft, indexRight)
		nodes[i].Value = val
		if indexLeft >= 0 {
			nodes[i].Left = &nodes[indexLeft]
		}
		if indexRight >= 0 {
			nodes[i].Right = &nodes[indexRight]
		}
	}
	return nodes
}

func printNode(n *Node) {
	fmt.Print("Value: ", n.Value)
	if n.Left != nil {
		fmt.Print(" Left: ", n.Left.Value)
	}
	if n.Right != nil {
		fmt.Print(" Right: ", n.Right.Value)
	}
	fmt.Println()
}

/*
	$ go run dfs_on_tree.go < dfs.in

	Value: 1
	Value: 3
	Value: 2 Left: 1 Right: 3
	Value: 6
	Value: 4 Left: 2 Right: 6
*/

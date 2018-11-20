package main

import (
	"sync"
	"testing"
)

// https://flaviocopes.com/golang-data-structure-binary-search-tree/

type Node struct {
	key   int
	value string
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
	lock sync.RWMutex
}

func TestBinarySearchTree(t *testing.T) {
	t.Log("Test binary_search_tree")
	list := []struct {
		tree []Node
		ans  string
	}{
		{
			tree: []Node{
				Node{key: 8, value: "8"},
				Node{key: 4, value: "4"},
				Node{key: 10, value: "10"},
				Node{key: 2, value: "2"},
				Node{key: 6, value: "6"},
				Node{key: 1, value: "1"},
				Node{key: 3, value: "3"},
				Node{key: 5, value: "5"},
				Node{key: 7, value: "7"},
				Node{key: 9, value: "9"},
			},
			ans: "cczxvzxcv",
		},
	}

	for _, item := range list {
		var tree BinarySearchTree
		for _, node := range item.tree {
			tree.Insert(node.key, node.value)
		}
		// TODO assert tree = ans
	}
}

// Insert new value
func (bst *BinarySearchTree) Insert(key int, value string) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := &Node{key, value, nil, nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

// internal function to find the correct place for a node in a tree
func insertNode(node, newNode *Node) {
	if newNode.key < node.key {

	} else {

	}
}

func (bst *BinarySearchTree) Remove(key int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	removeNode(bst.root, key)
}

func removeNode(node *Node, key int) *Node {
	return &Node{} // FIXME
}

func (bst *BinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
}

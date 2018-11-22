package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://flaviocopes.com/golang-data-structure-binary-search-tree/
// https://appliedgo.net/bintree/
// insert and delete https://www.youtube.com/watch?v=wcIRPqTR3Kc

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
			/*
							8
					   4		10
					2     6    9
				   1 3   5 7

			*/
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
		tree.String()
		assert.Equal(t, tree.IsExisted(6), true)
		t.Log("remove 6")
		tree.Remove(6)
		assert.Equal(t, tree.IsExisted(6), false)
		assert.Equal(t, tree.Min(), "1")
		assert.Equal(t, tree.Max(), "10")
		tree.String()
	}
}

// Insert new node
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

func insertNode(node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// Remove node
func (bst *BinarySearchTree) Remove(key int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	removeNode(bst.root, key)
}

func removeNode(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = removeNode(node.left, key)
		return node
	}
	if key > node.key {
		node.right = removeNode(node.right, key)
		return node
	}
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}
	leftmostrightside := node.right
	for {
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}
	node.key, node.value = leftmostrightside.key, leftmostrightside.value
	node.right = removeNode(node.right, node.key)
	return node
}

func (bst *BinarySearchTree) IsExisted(key int) bool {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return isExisted(bst.root, key)
}

func isExisted(n *Node, key int) bool {
	if n == nil {
		return false
	}
	if key < n.key {
		return isExisted(n.left, key)
	}
	if key > n.key {
		return isExisted(n.right, key)
	}
	return true
}

func (bst *BinarySearchTree) Max() string {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := bst.root
	if n == nil {
		return ""
	}
	for {
		if n.right == nil {
			return n.value
		}
		n = n.right
	}
}

func (bst *BinarySearchTree) Min() string {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := bst.root
	if n == nil {
		return ""
	}
	for {
		if n.left == nil {
			return n.value
		}
		n = n.left
	}
}

// prints a visual representation of the tree
func (bst *BinarySearchTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.right, level)
		fmt.Printf(format+"%d\n", n.key)
		stringify(n.left, level)
	}
}

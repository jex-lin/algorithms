package leetcode_problems

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestLevelOrder(t *testing.T) {
	t.Log("Test 102_binary_tree_level_order_traversal_test")
	list := []struct {
		Tree TreeNode
		Ans  [][]int
	}{
		{
			TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}},
			[][]int{[]int{3}, []int{9, 20}, []int{15, 7}},
		},
	}

	for _, item := range list {
		assert.Equal(t, reflect.DeepEqual(levelOrder(&item.Tree), item.Ans), true)
	}
}

// BFS
func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}

	var nodes []*TreeNode
	nodes = append(nodes, root)
	for len(nodes) > 0 {
		var tmp []int
		times := len(nodes)
		for i := 0; i < times; i++ {
			if nodes[i].Left != nil {
				nodes = append(nodes, nodes[i].Left)
			}
			if nodes[i].Right != nil {
				nodes = append(nodes, nodes[i].Right)
			}
			tmp = append(tmp, nodes[i].Val)
		}
		nodes = nodes[times:]
		ans = append(ans, tmp)
	}
	return ans
}

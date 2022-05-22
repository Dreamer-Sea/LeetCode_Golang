package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	node5 := &TreeNode{Val: 5, Left: nil, Right: nil}
	node2 := &TreeNode{Val: 2, Left: nil, Right: node5}
	node3 := &TreeNode{Val: 3, Left: nil, Right: nil}
	node1 := &TreeNode{Val: 1, Left: node2, Right: node3}
	binaryTreePaths(node1)
	fmt.Println(res3)
}

var res3 []string

func binaryTreePaths(root *TreeNode) []string {
	res3 = make([]string, 0)
	backtracking3(root, []int{})
	return res3
}

func backtracking3(node *TreeNode, temp []int) {
	if node == nil {
		return
	}
	temp = append(temp, node.Val)
	backtracking3(node.Left, temp)
	backtracking3(node.Right, temp)
	if node.Left == nil && node.Right == nil {
		tmp := ""
		for idx, e := range temp {
			tmp += strconv.Itoa(e)
			if idx != len(temp)-1 {
				tmp += "->"
			}
		}
		res3 = append(res3, tmp)
		return
	}
	temp = temp[:len(temp)-1]
}

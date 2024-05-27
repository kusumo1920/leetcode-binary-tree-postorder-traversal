package main

import "fmt"

func main() {
	input := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	output := postorderTraversal(input)
	fmt.Println(output)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := postorderTraversal(root.Left)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)
	return result
}

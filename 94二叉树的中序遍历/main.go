package main

import "fmt"

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	res := inorderTraversal(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// for循环实现中序遍历
func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return res
}

// 递归实现中序遍历
// func inorderTraversal(root *TreeNode) (res []int) {
// 	var inorder func(node *TreeNode)
// 	inorder = func(node *TreeNode) {
// 		if node == nil {
// 			return
// 		}
// 		inorder(node.Left)
// 		res = append(res, node.Val)
// 		inorder(node.Right)
// 	}
// 	inorder(root)
// 	return
// }

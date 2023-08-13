package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	res = append(res, []int{root.Val})
	temp := []int{}
	newStack := []*TreeNode{}
	for len(stack) > 0 {
		cur := stack[0]
		stack = stack[1:]
		if cur.Left != nil {
			temp = append(temp, cur.Left.Val)
			newStack = append(newStack, cur.Left)
		}
		if cur.Right != nil {
			temp = append(temp, cur.Right.Val)
			newStack = append(newStack, cur.Right)
		}
		if len(stack) == 0 && len(newStack) != 0 {
			stack = append(stack, newStack...)
			newStack = []*TreeNode{}
			res = append(res, temp)
			temp = []int{}
		}
	}
	return res
}

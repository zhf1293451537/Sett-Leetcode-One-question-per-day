package main

import "container/heap"

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 小根堆实现
func mergeKLists(lists []*ListNode) *ListNode {
	h := hp{}
	for _, head := range lists {
		if head != nil {
			h = append(h, head)
		}
	}
	heap.Init(&h)
	dummy := &ListNode{}
	cur := dummy
	for len(h) > 0 {
		node := heap.Pop(&h).(*ListNode)
		if node.Next != nil {
			heap.Push(&h, node.Next)
		}
		cur.Next = node
		cur = cur.Next
	}
	return dummy.Next
}

type hp []*ListNode

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(*ListNode)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

//暴力解决
// func mergeKLists(lists []*ListNode) *ListNode {
// 	src := []int{}
// 	for i := range lists {
// 		temp := lists[i]
// 		for temp != nil {
// 			src = append(src, temp.Val)
// 			temp = temp.Next
// 		}
// 	}
// 	if len(src) == 0 {
// 		return nil
// 	}
// 	sort.Slice(src, func(i, j int) bool { return src[i] < src[j] })
// 	res := &ListNode{}
// 	temp := res
// 	for i := range src {
// 		temp.Val = src[i]
// 		if i == len(src)-1 {
// 			break
// 		}
// 		temp.Next = &ListNode{Next: nil}
// 		temp = temp.Next
// 	}
// 	return res
// }

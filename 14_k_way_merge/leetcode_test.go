package k_way_merge

import (
	"container/heap"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Merge K Sorted Lists
// https://leetcode-cn.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	h := new(Heap)

	for index := range lists {
		node := lists[index]
		if node == nil {
			continue
		}
		heap.Push(h, node)
	}

	for h.Len() > 0 {
		e := heap.Pop(h).(*ListNode)
		if e.Next != nil {
			heap.Push(h, e.Next)
		}

		e.Next = nil
		cur.Next = e
		cur = cur.Next
	}

	return head.Next
}

func TestMergeKList(t *testing.T) {
	mergeKLists([]*ListNode{
		MakeLinkedList([]int{1, 2, 3}),
		MakeLinkedList([]int{4, 5, 6}),
	})
}

type Heap []*ListNode

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func MakeLinkedList(arr []int) *ListNode {
	var head *ListNode
	// make nodes
	var pre *ListNode
	for index, num := range arr {
		if index == 0 {
			head = &ListNode{
				Val: num,
			}
			pre = head
		} else {
			cur := &ListNode{
				Val: num,
			}
			pre.Next = cur
			pre = cur
		}
	}
	return head
}

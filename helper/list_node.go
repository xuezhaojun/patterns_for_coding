package helper

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
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

func PrintLinkedList(head *ListNode) {
	cur := head
	fmt.Println("list begin:")
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
	fmt.Println("list end")
	return
}

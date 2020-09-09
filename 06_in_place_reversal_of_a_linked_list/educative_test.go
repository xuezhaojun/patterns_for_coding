package in_place_reversal_of_a_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// Reverse a LinkedList
// https://leetcode-cn.com/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var pre *ListNode
	cur := head

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}

// Reverse a Sub-list
// https://leetcode-cn.com/problems/reverse-linked-list-ii/
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	var pre, t, h *ListNode

	preHead := &ListNode{
		Next: head,
	}

	// 找到启始反转位置
	cur := preHead // 通过preHead减少对反转设计头部的判断
	// 当然这里也可以在第一部分先找continue点，然后继续计算
	index := 0
	var con *ListNode
	for cur != nil && index < m {
		if index == m-1 {
			con = cur
		}
		index++
		cur = cur.Next
	}

	// 将中间部分反转，并记录首尾
	count := n - m + 1 // eg. 4-2 +1 = 3
	t = cur
	for count > 0 {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		count--
	}
	h = pre

	// 将首尾调转
	con.Next = h
	t.Next = cur

	return preHead.Next
}

// Reverse every K-element Sub-list
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
func reverseKGroup(head *ListNode, k int) *ListNode {
	preTail := &ListNode{}
	cur := head
	headMoved := false
	for cur != nil {
		count := k

		var oldTail, oldHead *ListNode
		var newTail, newHead *ListNode
		// 找到对应的newHead,和newTail
		oldHead = cur
		var pre *ListNode
		for cur != nil {
			pre = cur
			cur = cur.Next
			count--
		}
		oldTail = pre

		if count > 0 {
			// 进行从 oldHead 到 oldTail 的反转
			c := oldTail
			var p *ListNode
			for c != oldTail.Next {
				next := c.Next
				c.Next = p
				p = c
				c = next
			}
		}
		newHead = oldHead
		newTail = oldTail

		// 拼接
		preTail.Next = newHead
		preTail = newTail

		// headmove
		if !headMoved {
			head = newHead
			headMoved = true
		}
	}
	return head
}

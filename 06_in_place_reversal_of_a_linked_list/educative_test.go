package in_place_reversal_of_a_linked_list

import "testing"

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
// 如果本题进行考察的话，即属于考察人员的复杂问题的快速解决能力，目前这部分我还有欠缺，比如下列代码还不够简洁，同时解题速度也不够快
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
		for cur != nil && count > 0 {
			pre = cur
			cur = cur.Next
			count--
		}
		oldTail = pre

		if count == 0 {
			// 进行从 oldHead 到 oldTail 的反转
			c := oldHead
			var p *ListNode
			for count < k {
				next := c.Next
				c.Next = p
				p = c
				c = next
				count++
			}
			newHead = oldTail
			newTail = oldHead
		} else {
			newHead = oldHead
			newTail = oldTail
		}

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

func TestReverseKGroup(t *testing.T) {
	n5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	n4 := &ListNode{
		Val:  4,
		Next: n5,
	}
	n3 := &ListNode{
		Val:  3,
		Next: n4,
	}
	n2 := &ListNode{
		Val:  2,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}
	reverseKGroup(n1, 3)
}

// Reverse alternating K-element Sub-list
// 以下算法选取了 educative 上的算法流程; 整体的思路确实更加的清晰一些
func ReverseAlternateKElements(head *ListNode, k int) {
	// k = 1的时候相当于无反转
	if k <= 1 || head == nil {
		return
	}

	current := head
	var previous *ListNode

	for {
		lastNodeOfPreviousPart := previous // 这里的命名虽然长但是相比之下相当的清楚
		lastNodeOfSubList := current       // 当前节点的最后一个子节点（反转后）
		var next *ListNode

		// 进行k个节点的反转
		i := 0
		for current != nil && i < k {
			next = current.Next
			current.Next = previous
			previous = current
			current = next
			i++
		}

		// 将当前的sublist 和 preSubList connect
		if lastNodeOfPreviousPart != nil {
			lastNodeOfPreviousPart.Next = previous
		} else {
			head = previous
		}

		lastNodeOfSubList.Next = current

		i = 0
		for current != nil && i < k {
			previous = current
			current = current.Next
			i++
		}

		if current == nil {
			break
		}
	}
}

// Rotate List
// https://leetcode-cn.com/problems/rotate-list/submissions/
// 同样原地完成， 需要时间O(N)，共遍历两次
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	// 首先遍历一次链表获取长度
	cur := head
	lengthOfList := 0
	for cur != nil {
		cur = cur.Next
		lengthOfList++
	}

	// 对k进行取模，判断最小需要移动的量
	k = k % lengthOfList
	if k == 0 { // 这里稍微有一点点丑，因为下面的算法并不是兼容k=0的情况
		return head
	}

	// 进行移动
	oldHead := head
	cur = head
	var pre *ListNode
	count := 0
	for {
		if count == lengthOfList-k {
			break
		}
		pre = cur
		cur = cur.Next
		count++
	}

	// 设定新的首尾
	head = cur // k = 0 的情况下，cur is nil
	pre.Next = nil

	// 将后面这个部分接入
	for cur != nil {
		pre = cur
		cur = cur.Next
	}
	pre.Next = oldHead

	return head
}

// 快慢指针算法， 也常常被成为 Hare & Tortoise algorithm 龟兔算法，在处理环形的链表和数组的时候，往往很有效果
// 由于两个指针移动速度一快一慢，如果linkedlist是环形的，那么最终fast必然会catch到slow
package fast_and_slow_pointers

type ListNode struct {
	Val  int
	Next *ListNode
}

// 其中最简单的一个例子：
// 找到一个linkedList中是否存在环
// 对应leetcode中的题目： https://leetcode-cn.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head, head
	// 本题需要注意的此处的写法，比我自已处理的要简洁，在逻辑上推算一下
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 环路检测【重要/经典问题】
// leetcode题目： https://leetcode-cn.com/problems/linked-list-cycle-lcci/
// 设置一个end节点
// slow pointer 一个个的将node指向end
// fast pointer 持续前进
// 如果 fast 的下一个等于nil，或者下下一个等于nil，说明无环
// 如果 fast 的下一个或者下下一个等于 end，则当前节点为head
// 这个是我自己的算法，但是用到了额外的空间O(1),同时破坏了原来的数据结构
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow, fast := head, head
	end := &ListNode{}

	for fast.Next != nil || fast.Next.Next != nil {
		if fast.Next == end {
			return fast
		}
		if fast.Next.Next == end {
			return fast.Next
		}
		fast = fast.Next.Next
		preslow := slow
		slow = slow.Next
		preslow.Next = end
	}

	return nil
}

// 本题需要一定的数学推导，其中非常重要的一点就是，再两个pointer相遇之后，重新再从head出发，和slow以相同的速度进行移动
// 当此次两个pointer会在入口处相遇
// m = xn - y 的原因
// 通俗的解法来讲：【这个可能更好理解一些】
// 1. fast 和 slow 相遇
// 2. 计算得出cycle的长度
// 3. 让两个pointer回到起点，其中一个pointer先走一个cycle的长度
// 4. 然后两个pointer一起走，最终meet在入口处
func detectCycleStandard(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	slow, fast := head, head

	for {
		if !(fast.Next != nil && fast.Next.Next != nil) {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	if fast != slow {
		return nil
	}

	pointer1 := head

	for pointer1 != slow {
		pointer1 = pointer1.Next
		slow = slow.Next
	}

	return pointer1
}

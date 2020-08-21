// 快慢指针算法， 也常常被成为 Hare & Tortoise algorithm 龟兔算法，在处理环形的链表和数组的时候，往往很有效果
// 由于两个指针移动速度一快一慢，如果linkedlist是环形的，那么最终fast必然会catch到slow
// 龟兔算法的使用场景：
// * 判断是否成环（以时间换空间的方式，如果使用空间就需要使用map来缓存所有已经出现过的结果）
package fast_and_slow_pointers

import (
	"fmt"
	"testing"
)

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

// Happy number
// leetcode传送门 https://leetcode-cn.com/problems/happy-number/submissions/
// 直接使用缓存的方式，将所有的结果缓存下来，一旦出现重复就返回false
// 这种方法可以通过测试，但是存在内存占用很大的问题： 时间100% 空间7%
func isHappy(num int) bool {
	// resultSet 缓存了所有出现过的数字的结果
	resultSet := make(map[int]int)

	getNext := func(n int) (bool, int) {
		no := n
		if result, ok := resultSet[n]; ok {
			return true, result
		}
		result := 0
		for n != 0 {
			c := n % 10
			result += c * c
			n = n / 10
		}
		resultSet[no] = result
		return false, result
	}

	for {
		exist, next := getNext(num)
		if exist {
			return false
		}
		if next == 1 {
			return true
		}
		num = next
	}
}

// 此处可以学到的思想
// 这个关于时间的预估非常有意思，可以看leetcode的官方解法的说明
func isHappyLessSpace(num int) bool {
	// getNext算法的时间复杂度为O(log^n)
	// 开10的方
	getNext := func(n int) int {
		result := 0
		for n != 0 {
			c := n % 10
			result += c * c
			n = n / 10
		}
		return result
	}

	slow, fast := num, num

	for {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
		if fast == 1 {
			return true
		}
		if slow == fast {
			return false
		}
	}
}

func TestIsHappy(t *testing.T) {
	fmt.Println(isHappy(2))
}

// findMiddle
// 找到链表的中间位置
func findMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

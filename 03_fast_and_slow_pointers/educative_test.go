package fast_and_slow_pointers

import (
	"fmt"
	"patterns_for_coding_questions_in_golang/helper"
	"testing"
)

// 其中最简单的一个例子：
// 找到一个linkedList中是否存在环
// 对应leetcode中的题目： https://leetcode-cn.com/problems/linked-list-cycle/
func hasCycle(head *helper.ListNode) bool {
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
func detectCycle(head *helper.ListNode) *helper.ListNode {
	if head == nil {
		return head
	}

	slow, fast := head, head
	end := &helper.ListNode{}

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
func detectCycleStandard(head *helper.ListNode) *helper.ListNode {
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
func findMiddle(head *helper.ListNode) *helper.ListNode {
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

// isPalindrome
// leetcode : https://leetcode-cn.com/problems/palindrome-linked-list/
// 判断一个链表是否回文链表
// 这个题需要破坏原来的链表结构，本答案中最后并没有执行链表的恢复，但是实际中使用中，需要将这个链表保持原来的结构，则需要在比较中再次执行恢复
// 时间要求：O(n)
// 空间要求：O(1)
func isPalindrome(head *helper.ListNode) bool {
	if head == nil {
		return true
	}

	slow, fast := head, head
	// find middle
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse nodes after slow(include slow)
	last := slow
	cur := slow.Next
	for cur != nil {
		next := cur.Next
		cur.Next = last
		last = cur
		cur = next
	}

	p1, p2 := head, last

	for {
		if p1.Val != p2.Val {
			return false
		}

		// 这个在很多官方的解法中，使用的事p1/p2是否未nil
		// 如果未nil，则表示比完了, 但是对应我的reverse，这个p2==slow更好理解
		if p2 == slow {
			break
		}

		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

// Rerange a linkedList
func RearrangeALinkedlist(head *helper.ListNode) *helper.ListNode {
	if head == nil {
		return head
	}

	// find middle
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse
	last := slow
	cur := slow.Next
	slow.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = last
		last = cur
		cur = next
	}

	// migrate
	p1 := head
	p2 := last
	for p2.Next != nil {
		nextOfP1 := p1.Next
		nextOfP2 := p2.Next
		p1.Next = p2
		p1 = nextOfP1
		p2.Next = p1
		p2 = nextOfP2
	}

	return head
}

func TestRearrangeALinkedlist(t *testing.T) {
	helper.PrintLinkedList(RearrangeALinkedlist(helper.MakeLinkedList([]int{2, 4, 6, 8, 10, 12}))) // 2，12，4，10，6，8
	helper.PrintLinkedList(RearrangeALinkedlist(helper.MakeLinkedList([]int{2, 4, 6, 8, 10})))     // 2，10，4，8，6
}

// cycle in a circular array
// leetcode相同题目： https://leetcode-cn.com/problems/circular-array-loop/
// 本题在educative上的解释我觉得更加合理：
// 如果不记录遍历结果，那么就需要对每一个起点开始进行一次寻环算法，那么就需要O(n^2)的时间，需要O(1)的空间
// 如果通过map记录遍历结果，那么arr中的每一个元素只需要遍历一次，时间复杂度O(n)，但是空间复杂度O(1)
// 本题包含了空间时间互换的思维
func circularArrayLoop(nums []int) bool {
	// 这个算法中最重要的部分，就是这个数组中index的“左右横跳”
	next := func(curIndex, add int, positive bool) (success bool, nextIndex int) {
		// 在 direction 不同的情况下，直接返回无next，不成同向环
		if (add > 0 && !positive) || (add < 0 && positive) {
			return false, -1
		}
		// 此处应该是我想错了，不需要减1
		nextIndex = ((curIndex+add)%len(nums) + len(nums)) % len(nums) // wrap the negative situdation
		// 如果环长度小于1
		if curIndex == nextIndex {
			return false, -1
		}
		return true, nextIndex
	}
	// 以下算法中，缓存了每次的查询结果
	for index, num := range nums {
		slow, fast := index, index
		positive := num > 0
		for {
			// add 1 to slow
			var ok bool
			ok, slow = next(slow, nums[slow], positive)
			if !ok {
				break
			}
			// add 2 to fast
			ok, fast = next(fast, nums[fast], positive)
			if !ok {
				break
			}
			ok, fast = next(fast, nums[fast], positive)
			if !ok {
				break
			}
			// if fast == slow means there is a loop
			if fast == slow {
				// 成环
				return true
			}
		}
	}

	return false
}

func TestCircularArrayLoop(t *testing.T) {
	fmt.Println(circularArrayLoop([]int{1, 2, -1, 2, 2})) // True
	fmt.Println(circularArrayLoop([]int{2, 2, -2, 2}))    // False
}

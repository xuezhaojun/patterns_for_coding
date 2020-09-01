package fast_and_slow_pointers

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/
// 本题求一个链表中的倒数第k个数字
// 方案1：遍历两遍，第一遍获取到长度，第二遍移动length-k个单位，获取倒数k的值
// 方案2【快慢指针】：快指针比慢指针先走k步，快指针到tail的时候，慢指针到了倒数第k的位置
func kthToLast(head *ListNode, k int) int {
	fast, slow := head, head
	// fast move k
	for k > 0 {
		fast = fast.Next
		k--
	}
	// move fast and slow
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Val
}

// https://leetcode-cn.com/problems/find-the-duplicate-number/
func findDuplicate(nums []int) int {
	return 0
}

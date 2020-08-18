package monotone_stack

import (
	"container/list"
	"fmt"
	"testing"
)

// https://leetcode-cn.com/problems/sliding-window-maximum/
// 滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	result := []int{}
	queue := list.New()

	// init first windows
	for i := 0; i < k; i++ {
		for queue.Len() != 0 {
			back := queue.Back() // 队尾的下标
			if nums[back.Value.(int)] < nums[i] {
				// 如果队尾小于当前最新insert的元素，则将队尾去掉
				queue.Remove(back)
			} else {
				break
			}
		}
		queue.PushBack(i)
	}
	result = append(result, nums[queue.Front().Value.(int)])

	// process with rest of nums
	for i := k; i < len(nums); i++ {
		// insert
		for queue.Len() != 0 {
			back := queue.Back() // 队尾的下标
			if nums[back.Value.(int)] < nums[i] {
				// 如果队尾小于当前最新insert的元素，则将队尾去掉
				queue.Remove(back)
			} else {
				break
			}
		}
		queue.PushBack(i)
		// pop
		left := i - k // 最左边即将出队的数的下标
		if left == queue.Front().Value.(int) {
			queue.Remove(queue.Front())
		}
		// add new head to result
		result = append(result, nums[queue.Front().Value.(int)])
	}

	return result
}

func TestMaxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)) // [3,3,5,5,6,7]
	fmt.Println(maxSlidingWindow([]int{1}, 1))                        // [3,3,5,5,6,7]
}

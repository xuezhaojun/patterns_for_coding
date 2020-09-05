package cyclic_sort

import "testing"

// Cyclic sort
// 要求是in-place，所以并没有返回
// 时间要求 O(n)
// 空间要求 O(1)
// 本题利用的是给定数组的两个特性：
// 1. 在 1-n 的范围内
// 2. 无重复
// ps: 太喜欢找个算法了，干净有趣
func CyclicSort(nums []int) {
	for i := range nums {
		for nums[i] != i+1 {
			expectPos := nums[i] - 1
			nums[i], nums[expectPos] = nums[expectPos], nums[i]
		}
	}
	return
}

func TestCyclicSort(t *testing.T) {
	CyclicSort([]int{3, 1, 5, 4, 2})
	CyclicSort([]int{2, 6, 4, 3, 1, 5})
	CyclicSort([]int{1, 5, 6, 4, 3, 2})
}

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

// Find The Missing Number
// 关键信息：
// nums 无序，且对应缺失某一个序列中的num
// 本题的一个trick，因为nums的数量其实是缺1的，所以对应最后一个数，其实是可以不排的；
// 但是我这里不考虑空间使用问题，给原来的nums新加一个-1的位置，用这个-1位来表示缺失位
// 除了缺失位，其他所有的位置，都直到满足条件才会continue
func FindTheMissingNumber(nums []int) int {
	nums = append(nums, -1)
	for i := range nums {
		for nums[i] != i {
			expectPos := nums[i]
			if expectPos == -1 {
				break
			}
			nums[i], nums[expectPos] = nums[expectPos], nums[i]
		}
	}

	for i := range nums {
		if nums[i] == -1 {
			return i
		}
	}

	return -1
}

func TestFindTheMissingNumber(t *testing.T) {
	t.Log(FindTheMissingNumber([]int{4, 0, 3, 1}))             // 2
	t.Log(FindTheMissingNumber([]int{8, 3, 5, 2, 4, 6, 0, 1})) // 7
	t.Log(FindTheMissingNumber([]int{0, 4, 2, 1}))             // 3
}

// Find the Duplicate Number
func FindTheDuplicateNumber(nums []int) int {
	for i := range nums {
		for i != nums[i]-1 {
			expectPos := nums[i] - 1
			if nums[expectPos] == nums[i] {
				result := nums[i]
				return result
			} else {
				nums[i], nums[expectPos] = nums[expectPos], nums[i]
			}
		}
	}
	return -1
}

func TestFindTheDuplicateNumber(t *testing.T) {
	t.Log(FindTheDuplicateNumber([]int{1, 4, 4, 3, 2}))    // 4
	t.Log(FindTheDuplicateNumber([]int{2, 1, 3, 3, 5, 4})) // 3
	t.Log(FindTheDuplicateNumber([]int{2, 4, 1, 4, 4}))    // 4
}

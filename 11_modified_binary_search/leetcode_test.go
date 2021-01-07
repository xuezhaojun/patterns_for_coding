package modified_binary_search

import "testing"

// https://leetcode-cn.com/problems/binary-search/
func search(nums []int, target int) int {
	l, h := 0, len(nums)-1
	for l <= h {
		m := (l + h) / 2
		mNum := nums[m]
		switch {
		case mNum > target:
			h = m - 1
		case mNum < target:
			l = m + 1
		default:
			return m
		}
	}
	return -1
}

// https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target/
func nextGreatestLetter(letters []byte, target byte) byte {
	if target < letters[0] || target >= letters[len(letters)-1] {
		return letters[0]
	}
	l, h := 0, len(letters)-1
	var m int
	for l <= h {
		m = (l + h) / 2
		mNum := letters[m]
		if mNum > target {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	if letters[m] <= target {
		return letters[m+1]
	}
	return letters[m]
}

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	first := findFirstIndex(nums, target)
	if first == -1 {
		return []int{-1, -1}
	}
	last := findLastIndex(nums, target)
	return []int{first, last}
}

func findFirstIndex(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	var mi int
	for start < end {
		mi = (start + end) / 2
		m := nums[mi]
		switch {
		case m < target:
			start = mi + 1
		case m > target:
			end = mi - 1
		default:
			// case mid == target
			// 因为我们需要求first index，然后需要将end=mid-1
			end = mi
		}
	}
	// 最终跳出循环的时候，end = firstIndex
	if end >= 0 && nums[end] == target {
		return end
	}
	return -1
}

func findLastIndex(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	var mid int
	for start < end {
		mid = (start + end + 1) / 2
		m := nums[mid]
		switch {
		case m < target:
			start = mid + 1
		case m > target:
			end = mid - 1
		default:
			// case mid == target
			// 因为我们需要求first index，然后需要将end=mid-1
			start = mid
		}
	}
	// 最终跳出循环的时候，start = lastIndex
	return start
}

func TestSearchRange(t *testing.T) {
	// searchRange([]int{
	// 	5, 7, 7, 8, 8, 10,
	// }, 8)
	searchRange([]int{
		2, 2,
	}, 1)
}

// https://leetcode-cn.com/problems/find-peak-element/
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	start := 0
	end := len(nums) - 1
	var mid int
	for {
		if nums[start] >= nums[start+1] && nums[end-1] >= nums[end] {
			return start
		}
		if nums[start] <= nums[start+1] && nums[end-1] <= nums[end] {
			return end
		}
		mid = (start + end) / 2
		if nums[mid]-nums[mid+1] > 0 {
			end = mid
		} else {
			start = mid
		}
	}
}

func TestFindPeakElement(t *testing.T) {
	findPeakElement([]int{1, 2, 3, 1})
}

// Search in Rotated Array
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
func searchRotated(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		m := nums[mid]
		switch {
		case m < target:
			if m < nums[start] && target >= nums[start] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		case m > target:
			if m >= nums[start] && target < nums[start] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		default:
			return mid
		}
	}
	return -1
}

func TestSearchRotated(t *testing.T) {
	searchRotated([]int{4, 5, 6, 7, 0, 1, 2}, 0)
}

package two_pointers

import (
	"math"
	"sort"
	"testing"
)

// https://leetcode-cn.com/problems/3sum-closest/
// 最接近的三数之和
// 使用双指针算法的时间复杂度 O(n^2)
func threeSumClosest(nums []int, target int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	closestSum := math.MaxInt64
	minDiff := math.MaxInt64
	for i, num := range nums {
		begin := i + 1
		end := len(nums) - 1
		for begin < end {
			bv := nums[begin]
			ev := nums[end]
			sum := num + bv + ev
			sub := (sum - target)
			diff := sub * sub
			if minDiff > diff {
				minDiff = diff
				closestSum = sum
			}
			switch {
			case sum > target:
				end--
			case sum < target:
				begin++
			default:
				return sum
			}
		}
	}

	return closestSum
}

func TestThreeSumClosest(t *testing.T) {
	t.Log(threeSumClosest([]int{0, 2, 1, -3}, 1))
}

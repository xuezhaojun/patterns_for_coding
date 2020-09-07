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

// 长键按入
// https://leetcode-cn.com/problems/long-pressed-name/solution/
// 通过使用count的方式来解题，我也想到了，但是最终没有引入，反而导致思路越来越乱
// 有时候引入变量可以让code变复杂，但是思路变清晰，这时候要果断的引入变量【重要】
// 引入count的运算，让算法非常的简洁漂亮
func isLongPressedName(name string, typed string) bool {
	nlen := len(name)
	tlen := len(typed)
	i := 1
	j := 1
	counti := 0
	countj := 0
	for i < nlen || j < tlen {
		for i < nlen && name[i] == name[i-1] {
			i++
			counti++
		}
		for j < tlen && typed[j] == typed[j-1] {
			j++
			countj++
		}
		if typed[j-1] != name[i-1] || countj < counti {
			return false
		}
		i++
		j++
		counti, countj = 0, 0
	}
	return name[nlen-1] == typed[tlen-1]
}

func TestIsLongPressedName(t *testing.T) {
	t.Log(isLongPressedName("leelee", "lleeelee"))
	t.Log(isLongPressedName("kikcxmvzi", "kiikcxxmmvvzz"))
}

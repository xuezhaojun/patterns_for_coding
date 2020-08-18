package sliding_window

import (
	"fmt"
	"testing"
)

// https://leetcode-cn.com/problems/max-consecutive-ones-iii/
// 最大连续1的个数
// 数学上的证明：
// A 属于 B， f(A) = False 且 f(B) = True 的情况不存在
// 即，所有的替换k个的子串都会被滑到，所以算法成立
func longestOnes(A []int, K int) int {
	left := 0
	countZero := 0
	maxLength := 0
	for right := range A {
		// expand
		if A[right] == 0 {
			countZero += 1
		}

		if countZero <= K {
			if right-left+1 > maxLength {
				maxLength = right - left + 1
			}
			continue
		}

		// shrink
		for countZero > K {
			if A[left] == 0 {
				countZero--
			}
			left++
		}
	}
	return maxLength
}

func TestLongestOnes(t *testing.T) {
	fmt.Println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))                         // 6
	fmt.Println(longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3)) // 10
}

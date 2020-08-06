package sliding_window

import "testing"

// FindAveragesOfSubArrays
func FindAveragesOfSubArrays(K int, arr []int) []float64 {
	results := []float64{}
	windowsSum, windowsStart := 0, 0
	for windowsEnd := range arr {
		windowsSum += arr[windowsEnd]
		if windowsEnd >= K-1 {
			result := float64(windowsSum) / float64(K)
			results = append(results, result)
			windowsSum -= arr[windowsStart]
			windowsStart += 1
		}
	}
	return results
}

func TestFindAveragesOfSubArrays(t *testing.T) {
	// 时间复杂度为O(n), 因为每一个元素都只被加减1次
	results := FindAveragesOfSubArrays(5, []int{1, 3, 2, 6, -1, 4, 1, 8, 2})
	t.Log("Averages of subarrays of size K: ", results)
}

// FindMaxSubArrayOfSizeK
func FindMaxSubArrayOfSizeK(K int, arr []int) int {
	max := 0
	windowsSum, windowsStart := 0, 0
	for windowsEnd := range arr {
		windowsSum += arr[windowsEnd]
		if windowsEnd >= K-1 {
			if windowsSum > max {
				max = windowsSum
			}
			windowsSum -= arr[windowsStart]
			windowsStart += 1
		}
	}
	return max
}

func TestFindMaxSubArrayOfSizeK(t *testing.T) {
	t.Log(FindMaxSubArrayOfSizeK(2, []int{2, 3, 4, 1, 5})) // 预期结果为7
}

// SmallestSubarrayWithGivenSum
func SmallestSubarrayWithGivenSum(arr []int, s int) int {
	sumShortest := len(arr) + 1
	sumLen := 0
	windowsSum, windowsStart := 0, 0

	for windowsEnd := range arr {
		windowsSum += arr[windowsEnd]
		sumLen += 1
		for windowsSum >= s {
			if sumLen < sumShortest {
				sumShortest = sumLen
			}
			sumLen -= 1
			windowsSum -= arr[windowsStart]
			windowsStart += 1
		}
	}

	if sumShortest == len(arr)+1 {
		return 0
	}

	return sumShortest
}

func TestSmallestSubarrayWithGivenSum(t *testing.T) {
	t.Log(SmallestSubarrayWithGivenSum([]int{2, 1, 5, 2, 3, 2}, 7)) // 预期：2
	t.Log(SmallestSubarrayWithGivenSum([]int{2, 1, 5, 2, 8}, 7))    // 预期：1
	t.Log(SmallestSubarrayWithGivenSum([]int{3, 4, 1, 1, 6}, 8))    // 预期：3
}

// LongestSubstringKDistinct
func LongestSubstringKDistinct(s string, k int) int {
	longest := 0
	windowsStart := 0
	windows := make(map[byte]int)
	for windowsEnd := range s {
		windows[s[windowsEnd]] += 1
		if len(windows) <= k && windowsEnd-windowsStart+1 > longest {
			longest = windowsEnd - windowsStart + 1
		}
		for len(windows) > k && windowsStart <= windowsEnd {
			if v, ok := windows[s[windowsStart]]; ok && v == 1 {
				delete(windows, s[windowsStart])
			} else {
				windows[s[windowsStart]] -= 1
			}
			windowsStart += 1
		}
	}
	return longest
}

func TestLongestSubstringKDistinct(t *testing.T) {
	t.Log(LongestSubstringKDistinct("araaci", 2)) // 预期：4
	t.Log(LongestSubstringKDistinct("raa", 1))    // 预期：2
	t.Log(LongestSubstringKDistinct("cbbebi", 3)) // 预期：5
}

// Fruits into Baskets
func FruitsIntoBaskets(fruits []byte) int {
	longest := 0
	windowsMap := make(map[byte]int)
	windowsStart := 0
	for windowsEnd, f := range fruits {
		windowsMap[f] += 1
		if len(windowsMap) <= 2 && longest < windowsEnd-windowsStart+1 {
			longest = windowsEnd - windowsStart + 1
		}
		for len(windowsMap) > 2 {
			if v, ok := windowsMap[fruits[windowsStart]]; ok && v == 1 {
				delete(windowsMap, fruits[windowsStart])
			} else {
				windowsMap[fruits[windowsStart]] -= 1
			}
			windowsStart += 1
		}
	}
	return longest
}

func TestFruitsIntoBaskets(t *testing.T) {
	t.Log(FruitsIntoBaskets([]byte{'A', 'B', 'C', 'A', 'C'}))      // 预期：3
	t.Log(FruitsIntoBaskets([]byte{'A', 'B', 'C', 'B', 'B', 'C'})) // 预期：5
}

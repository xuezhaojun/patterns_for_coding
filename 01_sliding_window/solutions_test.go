package sliding_window

import "testing"

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

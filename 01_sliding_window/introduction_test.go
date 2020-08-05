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

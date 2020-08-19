// 双指针算法的使用场景：【重要】
// 用于找到一个【有序的】序列中，满足特定要求的几个元素(一对，三个，或者多个，甚至一个序列)
// 比如： 给定一个数字序列和一个target，找到序列中满足 a+b=target 的两个元素
// 心得：
// * 对于Triple的数组查询来说，双指针算法的优势就是将O(n^3)简化到O(n^2)
package two_pointers

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

// PairWithTargetSum
func PairWithTargetSum(arr []int, target int) []int {
	begin := 0
	end := len(arr) - 1
	for begin < end {
		if arr[begin]+arr[end] > target {
			end--
		} else if arr[begin]+arr[end] < target {
			begin++
		} else {
			return []int{begin, end}
		}
	}
	return []int{-1, -1}
}

func TestPairWithTargetSum(t *testing.T) {
	fmt.Println(PairWithTargetSum([]int{1, 2, 3, 4, 6}, 6)) // [1,3]
	fmt.Println(PairWithTargetSum([]int{2, 5, 9, 11}, 11))  // [0,2]
}

// Remove Duplicates
// 备注：不能使用额外的空间
// 第一次没有做出来标记
func RemoveDuplicates(arr []int) int {
	newArrIndex := 0
	next := 0
	for next < len(arr) {
		if arr[newArrIndex] == arr[next] {
			next += 1
		} else {
			newArrIndex += 1
			arr[newArrIndex] = arr[next]
			next += 1
		}
	}
	return newArrIndex + 1
}

func TestRemoveDuplicates(t *testing.T) {
	fmt.Println(RemoveDuplicates([]int{2, 3, 3, 3, 6, 9, 9})) // 4
	fmt.Println(RemoveDuplicates([]int{2, 2, 2, 11}))         // 2
}

// RemoveElement
// 备注：不能使用额外空间
func RemoveElement(arr []int, target int) int {
	newIndex := 0
	next := 0
	for next < len(arr) {
		if arr[next] == target {
			next++
			continue
		}
		arr[newIndex] = arr[next]
		newIndex++
		next++
	}
	return newIndex
}

func TestRemoveElement(t *testing.T) {
	fmt.Println(RemoveElement([]int{3, 2, 3, 6, 3, 10, 9, 3}, 3)) // 4
	fmt.Println(RemoveElement([]int{2, 11, 2, 2, 1}, 2))          // 2
}

// SquaringSortedArray
func SquaringSortedArray(arr []int) []int {
	left := 0
	right := len(arr) - 1
	newarr := make([]int, len(arr))
	index := len(arr) - 1

	for index > 0 {
		if arr[left]*arr[left] > arr[right]*arr[right] {
			newarr[index] = arr[left] * arr[left]
			index--
			left++
		} else {
			newarr[index] = arr[right] * arr[right]
			index--
			right--
		}
	}

	return newarr
}

func TestSquaringSortedArray(t *testing.T) {
	fmt.Println(SquaringSortedArray([]int{-2, -1, 0, 2, 3})) // [0, 1, 4, 4, 9]
	fmt.Println(SquaringSortedArray([]int{-3, -1, 0, 1, 2})) // [0, 1, 1, 4, 9]
}

// TripletSumToZero
// 同一个题见leetcode https://leetcode-cn.com/problems/3sum/submissions/
// 再次做时间从 375ms 提高到了 40ms，而且代码量也减少了
// 本题关键是如何精准的去重
func TripletSumToZero(arr []int) [][]int {
	result := [][]int{}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for cur, a := range arr {
		if cur != 0 && arr[cur] == arr[cur-1] {
			continue
		}
		i := cur + 1
		j := len(arr) - 1
		for i < j {
			if arr[i]+arr[j] < -a {
				i++
				for i < j && arr[i] == arr[i-1] {
					i++
				}
			} else if arr[i]+arr[j] > -a {
				j--
				for i < j && arr[j] == arr[j+1] {
					j--
				}
			} else {
				result = append(result, []int{a, arr[i], arr[j]})
				i++
				for i < j && arr[i] == arr[i-1] {
					i++
				}
			}
		}
	}

	return result
}

func TestTripletSumToZero(t *testing.T) {
	fmt.Println(TripletSumToZero([]int{-3, 0, 1, 2, -1, 1, -2})) // [-3, 1, 2], [-2, 0, 2], [-2, 1, 1], [-1, 0, 1]
	fmt.Println(TripletSumToZero([]int{-5, 2, -1, -2, 3}))       // [-5, 2, 3], [-2, -1, 3]
	fmt.Println(TripletSumToZero([]int{-1, 0, 1, 0}))            // [-1,0,1]
}

// TripletSumCloseToTarget
// 这个题是上一个题的变式，主要是能证明 curDiff 和 i，j 增减的关系
func TripletSumCloseToTarget(arr []int, target int) []int {
	result := []int{}
	diff := math.MaxInt64

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for cur, a := range arr {
		if cur != 0 && arr[cur] == arr[cur-1] {
			continue
		}
		i := cur + 1
		j := len(arr) - 1
		for i < j {
			sumOfThree := arr[i] + arr[j] + a
			curDiff := (sumOfThree - target)
			if curDiff == 0 {
				result = []int{a, arr[i], arr[j]}
				return result
			}
			if curDiff*curDiff < diff {
				diff = curDiff * curDiff
				result = []int{a, arr[i], arr[j]}
			}

			if curDiff < 0 {
				i++
				for i < j && arr[i] == arr[i-1] {
					i++
				}
			} else {
				j--
				for i < j && arr[j] == arr[j+1] {
					j--
				}
			}
		}
	}

	return result
}

func TestTripletSumCloseToTarget(t *testing.T) {
	fmt.Println(TripletSumCloseToTarget([]int{-2, 0, 1, 2}, 2))  // [-2 1 2]
	fmt.Println(TripletSumCloseToTarget([]int{-3, -1, 1, 2}, 1)) // [-3 1 2]
	fmt.Println(TripletSumCloseToTarget([]int{1, 0, 1, 1}, 100)) // [1 1 1]
}

// TripletsWithSmallerSum
func TripletsWithSmallerSum(arr []int, target int) int {
	result := 0

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for cur, a := range arr {
		if cur != 0 && arr[cur] == arr[cur-1] {
			continue
		}
		i := cur + 1
		j := len(arr) - 1
		for i < j {
			sumOfThree := arr[i] + arr[j] + a
			if sumOfThree >= target {
				j--
				for i < j && arr[j] == arr[j+1] {
					j--
				}
				continue
			}
			result += j - i
			i++
			for i < j && arr[i] == arr[i-1] {
				i++
			}
		}
	}

	return result
}

func TestTripletsWithSmallerSum(t *testing.T) {
	fmt.Println(TripletsWithSmallerSum([]int{-1, 0, 2, 3}, 3))    // 2
	fmt.Println(TripletsWithSmallerSum([]int{-1, 4, 2, 1, 3}, 5)) // 4
}

// Subarrays With Product Less than a Target
// 奇怪的是，本题更倾向于一个slice windows的题型，而不是two points的题型
// 也可以认为 slice windows也就是two points的一种变式
func SubarraysWithProductLessThanATarget(array []int, target int) [][]int {
	result := [][]int{}
	left := 0
	product := 1
	for right := range array {
		product *= array[right]

		for product >= target && left < len(array) {
			product /= array[left]
			left++
		}

		sub := []int{}
		for i := right; i >= left; i-- {
			newSub := make([]int, len(sub))
			copy(newSub, sub)
			newSub = append(newSub, array[i])
			result = append(result, newSub)
			sub = newSub
		}

	}
	return result
}

func TestSubarraysWithProductLessThanATarget(t *testing.T) {
	fmt.Println(SubarraysWithProductLessThanATarget([]int{2, 5, 3, 10}, 30)) // [[2], [5], [2, 5], [3], [5, 3], [10]]
	fmt.Println(SubarraysWithProductLessThanATarget([]int{8, 2, 6, 5}, 50))  // [[8], [2], [8, 2], [6], [2, 6], [5], [6, 5]]
}

// Dutch National Flag Problem
func DutchNationalFlagProblem(array []int) {
	low, cur, high := 0, 0, len(array)-1

	for cur < high {
		switch array[cur] {
		case 0:
			if cur == low {
				cur++
			}
			for array[low] == 0 {
				low += 1
				continue
			}
			array[low], array[cur] = array[cur], array[low]
		case 2:
			for array[high] == 2 {
				high--
				continue
			}
			array[high], array[cur] = array[cur], array[high]
		case 1:
			cur++
		}
	}

	return
}

func TestDutchNationalFlagProblem(t *testing.T) {
	array := []int{1, 0, 2, 1, 0}
	DutchNationalFlagProblem(array)
	fmt.Println(array) // 0 0 1 1 2

	array2 := []int{2, 2, 0, 1, 2, 0}
	DutchNationalFlagProblem(array2)
	fmt.Println(array2) // 0 0 1 2 2 2
}

// QuadrupleSumToTarget
// 四数之和
// leetcode 相同题：https://leetcode-cn.com/problems/4sum/
// 需要注意的是对于相同a，b，i，j 的去重
// i，j 去重需要考虑，i 始终要满足小于 j ： [0,0,0,0],0 的测试用例
func QuadrupleSumToTarget(arr []int, target int) [][]int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	result := [][]int{}
	for a := range arr {
		if a != 0 && arr[a] == arr[a-1] {
			continue
		}
		for b := a + 1; b < len(arr); b++ {
			if b != a+1 && arr[b] == arr[b-1] {
				continue
			}
			i := b + 1
			j := len(arr) - 1
			for i < j {
				sum := arr[a] + arr[b] + arr[i] + arr[j]
				switch {
				case sum > target:
					j--
					for arr[j] == arr[j+1] && i < j {
						j--
					}
				case sum < target:
					i++
					for arr[i] == arr[i-1] && i < j {
						i++
					}
				default:
					result = append(result, []int{arr[a], arr[b], arr[i], arr[j]})
					i++
					for arr[i] == arr[i-1] && i < j {
						i++
					}
					j--
					for arr[j] == arr[j+1] && i < j {
						j--
					}
				}
			}
		}
	}
	return result
}

func TestQuadrupleSumToTarget(t *testing.T) {
	fmt.Println(QuadrupleSumToTarget([]int{1, 0, -1, 0, -2, 2}, 0)) // [-1,  0, 0, 1],[-2, -1, 1, 2],[-2,  0, 0, 2]
}

// backspaceCompare
// 比较含有退格字符的两个字符串
// leetcode 对应题： https://leetcode-cn.com/problems/backspace-string-compare/
func back(str string) string {
	arr := []byte(str)
	n, cur := 0, 0
	for cur < len(arr) {
		if arr[cur] == '#' {
			if n > 0 {
				n--
			}
			cur++
			continue
		}
		if n == cur {
			n++
			cur++
			continue
		}
		arr[n] = arr[cur]
		n++
		cur++
	}
	return string(arr[:n])
}
func backspaceCompare(S string, T string) bool {
	return back(S) == back(T)
}

func TestBackSpaceCompare(t *testing.T) {
	fmt.Println(backspaceCompare("xy#z", "xzz#")) // true
	fmt.Println(backspaceCompare("xy#z", "xyz#")) // false
	fmt.Println(backspaceCompare("x##z", "x#z#")) // false
}

// Minimum Window Sort
// O(n*log^n)
func MinimumWindowSort1(nums []int) int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Slice(sortedNums, func(i, j int) bool {
		return sortedNums[i] < sortedNums[j]
	})

	i, j := 0, len(nums)-1
	for i < len(nums)-1 {
		if sortedNums[i] != nums[i] {
			break
		}
		i++
	}

	for j > i {
		if sortedNums[j] != nums[j] {
			break
		}
		j--
	}

	if i == j {
		return 0
	}

	return j - i + 1
}

// O(n)
// 这道题我学到一个非常重要的算法思想： 最终的成果可以是一步一步叠加计算而来的
// 本题中，一开始的过程只是得到的最终结果的一个子串，但是基于这个子串不断扩展expand，最终可以得到正确的结果
func MinimumWindowSort2(nums []int) int {
	low, high := 0, len(nums)-1
	for low < len(nums)-1 && nums[low] <= nums[low+1] {
		low++
	}

	if low == len(nums)-1 {
		return 0
	}

	for high > 0 && nums[high] >= nums[high-1] {
		high -= 1
	}

	// find max and min
	min := math.MaxInt64
	max := math.MinInt64

	for k := low; k <= high; k++ {
		if nums[k] > max {
			max = nums[k]
		}
		if nums[k] < min {
			min = nums[k]
		}
	}

	for low > 0 && nums[low-1] > min {
		low -= 1
	}

	for high < len(nums)-1 && nums[high+1] < max {
		high += 1
	}

	return high - low + 1
}

func TestMinimumWindowSort(t *testing.T) {
	fmt.Println(MinimumWindowSort1([]int{1, 2, 5, 3, 7, 10, 9, 12})) // 5
	fmt.Println(MinimumWindowSort1([]int{1, 3, 2, 0, -1, 7, 10}))    // 5
	fmt.Println(MinimumWindowSort1([]int{1, 2, 3}))                  // 0
	fmt.Println(MinimumWindowSort1([]int{3, 2, 1}))                  // 3

	fmt.Println(MinimumWindowSort2([]int{1, 2, 5, 3, 7, 10, 9, 12})) // 5
	fmt.Println(MinimumWindowSort2([]int{1, 3, 2, 0, -1, 7, 10}))    // 5
	fmt.Println(MinimumWindowSort2([]int{1, 2, 3}))                  // 0
	fmt.Println(MinimumWindowSort2([]int{3, 2, 1}))                  // 3
}

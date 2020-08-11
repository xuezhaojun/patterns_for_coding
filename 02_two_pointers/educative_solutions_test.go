// 双指针算法，用于找到一个【有序的】序列中，满足特定要求的几个元素(一对，三个，或者多个)
// 比如： 给定一个数字序列和一个target，找到序列中满足 a+b=target 的两个元素
package two_pointers

import (
	"fmt"
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

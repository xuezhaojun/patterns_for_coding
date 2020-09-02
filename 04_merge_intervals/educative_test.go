package merge_intervals

import (
	"sort"
	"testing"
)

// 合并区间
// https://leetcode-cn.com/problems/merge-intervals/
// 本题按照建议解法： 1 按照start进行排序 2 遍历intervals，如果发现两个intervals属于可以合并的情况，则进行合并
// 此处的解法是在一开始设想的：这个解法没有考虑倒 [0,0][1,4] 这样特殊的情况（一个区间首尾相同的情况），所以最终是没有通过leetcode的测试(但是可以通过educative的测试)
// 时间复杂度也是 O(n*lgn)
func mergeIntervals(intervals [][]int) [][]int {
	result := [][]int{}

	all := []int{}
	status := make(map[int]int) // 1: start 2: overlap 3: end

	// init
	for _, interval := range intervals {
		all = append(all, interval...)
		start := interval[0]
		end := interval[1]
		if _, ok := status[start]; !ok {
			status[start] = 1
		} else {
			if status[start] != 1 {
				status[start] = 2
			}
		}

		if _, ok := status[end]; !ok {
			status[end] = 3
		} else {
			if status[end] != 3 {
				status[end] = 2
			}
		}
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i] < all[j]
	})

	var start, end int
	for start <= len(all)-1 {
		if end == len(all)-1 || (status[all[end]] == 3 && status[all[end+1]] == 1) {
			result = append(result, []int{
				all[start], all[end],
			})
			end = end + 1
			start = end
		}
		end++
	}

	return result
}

func TestMergeIntervals(t *testing.T) {
	t.Log(mergeIntervals([][]int{{1, 4}, {4, 5}}))                    // [[1,5]]
	t.Log(mergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})) // [[1,6],[8,10],[15,8]]
}

// 插入区间
// https://leetcode-cn.com/problems/insert-interval/
// 方案1的空间都不好：时间80%，空间10%; 但是基本思路是和educative相同的
func insertIntervalas(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}

	var i int

	for i < len(intervals) {
		// init
		interval := intervals[i]
		end1 := interval[1]
		start2 := newInterval[0]

		if start2 <= end1 {
			break
		}

		result = append(result, interval)
		i++
	}

	for i < len(intervals) {
		// init
		interval := intervals[i]
		start1 := interval[0]
		end1 := interval[1]
		start2 := newInterval[0]
		end2 := newInterval[1]

		if start1 > end2 {
			break
		}

		newStart, newEnd := merge(start1, end1, start2, end2)
		newInterval = []int{newStart, newEnd}
		i++
	}

	result = append(result, newInterval)
	result = append(result, intervals[i:]...)

	return result
}

func merge(start1, end1, start2, end2 int) (start, end int) {
	// start
	if start1 < start2 {
		start = start1
	} else {
		start = start2
	}
	// end
	if end1 < end2 {
		end = end2
	} else {
		end = end1
	}
	return
}

func TestInsertIntervals(t *testing.T) {
	t.Log(insertIntervalas([][]int{{1, 3}, {6, 9}}, []int{2, 5}))                            // [[1,5],[6,9]]
	t.Log(insertIntervalas([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})) // [[1,2],[3,10],[12,16]]
	t.Log(insertIntervalas([][]int{}, []int{4, 8}))                                          // [[4,8]]
}

// 区间列表的交集
// leetcode: https://leetcode-cn.com/problems/interval-list-intersections/
// 两个列表都是【已排序的】
func intervalIntersection(A [][]int, B [][]int) [][]int {
	result := [][]int{}
	return result
}

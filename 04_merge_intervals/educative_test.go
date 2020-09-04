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
// 时间70% 空间7%
func intervalIntersection(A [][]int, B [][]int) [][]int {
	result := [][]int{}
	var i, j int
	for i < len(A) && j < len(B) {
		intervalA := A[i]
		intervalB := B[j]
		if overlap(intervalA[0], intervalA[1], intervalB[0], intervalB[1]) {
			start, end := getintersection(intervalA[0], intervalA[1], intervalB[0], intervalB[1])
			result = append(result, []int{start, end})
		}
		if intervalA[1] <= intervalB[1] {
			i++
		} else {
			j++
		}
	}
	return result
}

// 方案2为优化，选择了更简洁的方式来表达intersection这个状态
// 时间复杂度可以优化倒 93%
func intervalIntersection2(A [][]int, B [][]int) [][]int {
	result := [][]int{}
	var i, j int
	for i < len(A) && j < len(B) {
		a := A[i]
		b := B[j]

		// 【重要】如果判断两个区间是否overlap的快捷思路
		var lo, hi int
		if a[0] < b[0] {
			lo = b[0]
		} else {
			lo = a[0]
		}
		if a[1] > b[1] {
			hi = b[1]
		} else {
			hi = a[1]
		}
		if lo <= hi {
			result = append(result, []int{lo, hi})
		}

		if a[1] <= b[1] {
			i++
		} else {
			j++
		}
	}
	return result
}

func overlap(start1, end1, start2, end2 int) bool {
	if (start1 <= start2 && start2 <= end1) || (start2 <= start1 && start1 <= end2) {
		return true
	}
	return false
}

func getintersection(start1, end1, start2, end2 int) (start, end int) {
	if start1 < start2 {
		start = start2
	} else {
		start = start1
	}

	if end1 < end2 {
		end = end1
	} else {
		end = end2
	}

	return
}

// Conflicting Appointments
// 这个题按照merge intervals的模式可以快速完成，主要在如何判断两个interval是否覆盖的问题上
func AttendToAllAppointments(intervals [][]int) bool {
	// 首先讲intervals的按照开始时间排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	if len(intervals) == 1 {
		return true
	}

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}

	return true
}

func TestAttendToAllAppointments(t *testing.T) {
	t.Log(AttendToAllAppointments([][]int{{1, 4}, {2, 5}, {7, 9}}))  // false
	t.Log(AttendToAllAppointments([][]int{{6, 7}, {2, 4}, {8, 12}})) // true
	t.Log(AttendToAllAppointments([][]int{{4, 5}, {2, 3}, {3, 6}}))  // false
}

// Minimum Meeting Rooms 最小数量的会议间
func MinimumMeetingRooms(meetings [][]int) int {
	// 对meetings进行排序
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	// 统计有多少交集
	maxRooms := 1
	curRooms := 1
	curMeeting := meetings[0]
	for i := 1; i < len(meetings); i++ {
		if curMeeting[1] > meetings[i][0] {
			curRooms++
			if curRooms > maxRooms {
				maxRooms = curRooms
			}
			curMeeting = []int{max(curMeeting[0], meetings[i][0]), min(curMeeting[1], meetings[i][1])}
		} else {
			curRooms = 1
			curMeeting = meetings[i]
		}
	}

	return maxRooms
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

func TestMinimunMeetingRooms(t *testing.T) {
	t.Log(MinimumMeetingRooms([][]int{{1, 4}, {2, 5}, {7, 9}}))         // 2
	t.Log(MinimumMeetingRooms([][]int{{6, 7}, {2, 4}, {8, 12}}))        // 1
	t.Log(MinimumMeetingRooms([][]int{{1, 4}, {2, 3}, {3, 6}}))         // 2
	t.Log(MinimumMeetingRooms([][]int{{4, 5}, {2, 3}, {2, 4}, {3, 5}})) // 2
}

// Maximum CPU Load
// max cpu load 基本是上一个题的直观改变，只要把room一次加一个改为一次加cpu load即可
func MaximumCPULoad(jobs [][]int) (maxcpuload int) {
	// 对jobs进行排序
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] < jobs[j][0]
	})

	curJob := jobs[0]
	curLoad := jobs[0][2]

	for i := 1; i < len(jobs); i++ {
		if curJob[1] > jobs[i][0] {
			curLoad += jobs[i][2]
			curJob = []int{max(curJob[0], jobs[i][0]), min(curJob[1], jobs[i][1]), curLoad}
		} else {
			curJob = jobs[i]
			curLoad = jobs[i][2]
		}

		if curLoad > maxcpuload {
			maxcpuload = curLoad
		}
	}

	return
}

func TestMaximumCPULoad(t *testing.T) {
	t.Log(MaximumCPULoad([][]int{{1, 4, 3}, {2, 5, 4}, {7, 9, 6}}))     // 7
	t.Log(MaximumCPULoad([][]int{{6, 7, 10}, {2, 4, 11}, {8, 12, 15}})) // 15
	t.Log(MaximumCPULoad([][]int{{1, 4, 2}, {2, 4, 1}, {3, 6, 5}}))     // 8
}

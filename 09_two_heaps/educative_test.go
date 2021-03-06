package two_heaps

import (
	"container/heap"
	"testing"
)

// Find the Median of a Number Stream
// https://leetcode-cn.com/problems/find-median-from-data-stream/
type MedianFinder struct {
	maxHeap *MaxHeap
	minHeap *MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		maxHeap: new(MaxHeap),
		minHeap: new(MinHeap),
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 判断num是否大于minHeap的最小值
	// 如果大于，则加入minHeap
	// 如果小于，则加入maxHeap
	if this.minHeap.Len() > 0 && num > (*this.minHeap)[0] {
		heap.Push(this.minHeap, num)
	} else {
		heap.Push(this.maxHeap, num)
	}

	// 加入后进行rebalance，两个heap的元素数量最多相差1
	// 否则就从多的里拿出一个来，放到少的里面
	if this.minHeap.Len()-this.maxHeap.Len() > 1 {
		e := heap.Pop(this.minHeap).(int)
		heap.Push(this.maxHeap, e)
	} else if this.maxHeap.Len()-this.minHeap.Len() > 1 {
		e := heap.Pop(this.maxHeap).(int)
		heap.Push(this.minHeap, e)
	}

}

func (this *MedianFinder) FindMedian() float64 {
	// 如果两个相等，则求平均
	// 如果minHeap大于maxHeap 则是 minHeap的min值
	// 如果maxHeap大于minHeap 则是 maxHeap的max值
	minHeapLen := this.minHeap.Len()
	maxHeapLen := this.maxHeap.Len()
	switch {
	case minHeapLen > maxHeapLen:
		return float64((*this.minHeap)[0])
	case maxHeapLen > minHeapLen:
		return float64((*this.maxHeap)[0])
	default:
		return float64((*this.minHeap)[0]+(*this.maxHeap)[0]) / 2
	}
}

// An MinHeap is a min-heap of ints.
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// An MaxHeap is a max-heap of ints.
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Sliding Window Median
// https://leetcode-cn.com/problems/sliding-window-median/
// 右边移动进来的一直按照two heaps进行添加即可，关键左边删除的节点，如何快速的从two heaps中出去
// 由于元素是可以重复的，所有我们的valied需要记录实际的index
func medianSlidingWindow(nums []int, k int) []float64 {
	minHeap := new(MinHeap)
	maxHeap := new(MaxHeap)
	invailed := make(map[int]struct{})

	// 首先是入队k个元素
	for i := 0; i < k; i++ {
		heap.Push(minHeap, nums[i])
	}

	// 从minHeap中取出k/2个元素，放入maxHeap中进行rebalance
	halfK := k / 2
	for i := 0; i < halfK; i++ {
		e := heap.Pop(minHeap).(int)
		heap.Push(maxHeap, e)
	}

	// 此时如何k = 5，那么minHeap中有3个元素，maxHeap有2个元素
	balance := minHeap.Len() - maxHeap.Len()
	rawbalance := balance

	// 操作剩下的元素
	leftIndex := 0
	result := []float64{}
	for i := k; i <= len(nums); i++ {
		// 求出median加入结果
		if balance == 0 {
			minT := (*minHeap)[0]
			maxT := (*maxHeap)[0]
			result = append(result, float64(minT+maxT)/2)
		} else {
			result = append(result, float64((*minHeap)[0]))
		}

		if i == len(nums) {
			break
		}

		// 延迟删除左侧元素
		// 修改balance
		// 【重要】结合大小堆的性格，如果一个left大于minT,就是minHeap中的元素
		left := nums[leftIndex]
		invailed[left] = struct{}{}
		leftIndex++
		if left >= (*minHeap)[0] {
			balance -= 1
		} else {
			balance += 1
		}

		// 加入右侧最新的元素
		// 如果right大于minHeap的最小值，则将其加入minHeap
		// 如果right小于minHeap的最大值，则将其加入maxHeap
		right := nums[i]
		if right > (*minHeap)[0] {
			heap.Push(minHeap, right)
			balance += 1
		} else {
			heap.Push(maxHeap, right)
			balance -= 1
		}

		// rebalance
		// rebalance之后，两个堆顶也不一定都是有效的
	LOOP:
		for {
			switch {
			case balance > rawbalance: // min多
				minT := (*minHeap)[0]
				heap.Pop(minHeap)
				heap.Push(maxHeap, minT)
				balance -= 2
			case balance < rawbalance: // max多
				maxT := (*maxHeap)[0]
				heap.Pop(maxHeap)
				heap.Push(minHeap, maxT)
				balance += 2
			default:
				break LOOP
			}
		}

		// 保证堆顶有效
		for {
			minT := (*minHeap)[0]
			if _, ok := invailed[minT]; ok {
				heap.Pop(minHeap)
				delete(invailed, minT)
			} else {
				break
			}
		}

		for {
			if maxHeap.Len() == 0 {
				break
			}
			maxT := (*maxHeap)[0]
			if _, ok := invailed[maxT]; ok {
				heap.Pop(maxHeap)
				delete(invailed, maxT)
			} else {
				break
			}
		}
	}

	return result
}

func TestMedianSlidingWindow(t *testing.T) {
	result := medianSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	// result := medianSlidingWindow([]int{1, 2}, 1)
	// result := medianSlidingWindow([]int{1, 1, 1, 1}, 2)
	t.Log(result)
}

// Maximize Capital
// https://leetcode-cn.com/problems/ipo/solution/
// 本题比较烦的是min-heap和max-heap需要重新定义
func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	ch := new(capitalHeap)
	ph := new(profitHeap)

	// 构成project
	var projects []project
	for i := 0; i < len(Profits); i++ {
		projects = append(projects, project{
			index:   i,
			profit:  Profits[i],
			capital: Capital[i],
		})
	}

	// 首先将projects按照capital加入min-heap
	for _, p := range projects {
		heap.Push(ch, p)
	}

	// 取k次
	for i := 0; i < k; i++ {
		for {
			if ch.Len() == 0 {
				break
			}
			minC := (*ch)[0].capital
			if minC <= W {
				e := heap.Pop(ch).(project)
				heap.Push(ph, e)
			} else {
				break
			}
		}
		// 取堆顶
		if ph.Len() == 0 {
			break
		}
		p := (*ph)[0].profit
		W += p
		heap.Pop(ph)
	}

	return W
}

type project struct {
	index   int
	profit  int
	capital int
}

type capitalHeap []project // min-heap
type profitHeap []project  // max-heap

func (h capitalHeap) Len() int {
	return len(h)
}

func (h capitalHeap) Less(i, j int) bool {
	return h[i].capital < h[j].capital
}

func (h capitalHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *capitalHeap) Push(x interface{}) {
	*h = append(*h, x.(project))
}

func (h *capitalHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h profitHeap) Len() int {
	return len(h)
}

func (h profitHeap) Less(i, j int) bool {
	return h[i].profit > h[j].profit
}

func (h profitHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *profitHeap) Push(x interface{}) {
	*h = append(*h, x.(project))
}

func (h *profitHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Next Interval
// https://leetcode-cn.com/problems/find-right-interval/
func findRightInterval(intervals [][]int) []int {
	// 首先构造两个min-heap
	startHeap := new(intervalHeap)
	endHeap := new(intervalHeap)
	record := make(map[int]int)

	// 填入结果
	for index, interval := range intervals {
		start := interval[0]
		end := interval[1]
		heap.Push(startHeap, in{
			index: index,
			num:   start,
		})
		heap.Push(endHeap, in{
			index: index,
			num:   end,
		})
	}

	for {
		startMin := (*startHeap)[0]
		endMin := (*endHeap)[0]

		if endMin.num > startMin.num {
			heap.Pop(startHeap)
		} else {
			record[endMin.index] = startMin.index
			heap.Pop(endHeap)
		}

		if startHeap.Len() == 0 || endHeap.Len() == 0 {
			break
		}
	}

	var result []int
	for index, _ := range intervals {
		if next, ok := record[index]; ok {
			result = append(result, next)
		} else {
			result = append(result, -1)
		}
	}

	return result
}

type in struct {
	index int
	num   int
}

// intervalHeap is a min-heap
type intervalHeap []in

func (h intervalHeap) Len() int {
	return len(h)
}

func (h intervalHeap) Less(i, j int) bool {
	return h[i].num < h[j].num
}

func (h intervalHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intervalHeap) Push(x interface{}) {
	*h = append(*h, x.(in))
}

func (h *intervalHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func TestFindRightInterval(t *testing.T) {
	findRightInterval([][]int{
		{3, 4},
		{1, 5},
		{4, 6},
	})
}

package two_heaps

import "container/heap"

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
